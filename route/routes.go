package route

import (
    "time"

    "github.com/gin-gonic/gin"
    "golang.org/x/time/rate"
    "gorm.io/gorm"
    "product-api/config"
    "product-api/controller"
    "product-api/middleware"
    "product-api/repository"
    "product-api/service"
)

func SetupRouter(db *gorm.DB, cfg *config.Config) *gin.Engine {
    r := gin.Default()

    // Middleware global
    r.Use(middleware.CORSMiddleware())

    // Dependency injection
    userRepo := repository.NewUserRepository(db)
    productRepo := repository.NewProductRepository(db)

    authService := service.NewAuthService(userRepo, cfg)
    productService := service.NewProductService(productRepo)

    authController := controller.NewAuthController(authService)
    productController := controller.NewProductController(productService)

    // Rate limiter
    authLimiter := middleware.NewRateLimiter(rate.Every(60*time.Second/3), 3)      // 3 req / 60 detik
    productWriteLimiter := middleware.NewRateLimiter(rate.Every(5*time.Second), 1) // 1 req / 5 detik

    api := r.Group("/api")

    // Route auth
    auth := api.Group("/auth")
    auth.Use(authLimiter.Middleware())
    {
        auth.POST("/register", authController.Register)
        auth.POST("/login", authController.Login)
    }

    // Route produk
    products := api.Group("/products")
    {
        products.GET("", productController.GetAll)
        products.GET("/:id", productController.GetByID)

        // Route yang butuh autentikasi
        protected := products.Group("")
        protected.Use(middleware.AuthMiddleware(authService))
        protected.Use(productWriteLimiter.Middleware())
        {
            protected.POST("", productController.Create)
            protected.PUT("/:id", productController.Update)
            protected.DELETE("/:id", productController.Delete)
        }
    }

    return r
}