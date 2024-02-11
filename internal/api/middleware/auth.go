package middleware

import (
	"main/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("your-secret-key")

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}
		claims, err := utils.ValidateToken(tokenString)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("user", claims) // Set user information to gin context for later use
		c.Next()
	}
}
