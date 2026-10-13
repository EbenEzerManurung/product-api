package middleware

import (
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"

    "product-api/service"
)

func AuthMiddleware(authService *service.AuthService) gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Header Authorization wajib diisi"})
            c.Abort()
            return
        }

        parts := strings.SplitN(authHeader, " ", 2)
        if len(parts) != 2 || parts[0] != "Bearer" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Format authorization tidak valid"})
            c.Abort()
            return
        }

        claims, err := authService.ValidateToken(parts[1])
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid atau sudah kadaluarsa"})
            c.Abort()
            return
        }

        c.Set("user_id", claims["sub"])
        c.Set("username", claims["username"])
        c.Next()
    }
}