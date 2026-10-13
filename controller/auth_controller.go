package controller

import (
    "net/http"

    "github.com/gin-gonic/gin"

    "product-api/service"
)

type AuthController struct {
    authService *service.AuthService
}

func NewAuthController(s *service.AuthService) *AuthController {
    return &AuthController{authService: s}
}


func (h *AuthController) Register(c *gin.Context) {
    var input service.RegisterInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := h.authService.Register(input)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message": "User berhasil didaftarkan",
        "user": gin.H{
            "id":       user.ID,
            "username": user.Username,
        },
    })
}

// Login godoc
// @Summary      Login
// @Description  Autentikasi dan menerima token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        body  body  service.LoginInput  true  "Kredensial login"
// @Success      200   {object}  service.TokenResponse
// @Failure      401   {object}  map[string]string
// @Router       /auth/login [post]
func (h *AuthController) Login(c *gin.Context) {
    var input service.LoginInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    tokens, err := h.authService.Login(input)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, tokens)
}