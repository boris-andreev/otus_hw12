package authmiddleware

import (
	"hw12/internal/utils/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// TokenAuthMiddleware Middleware для проверки JWT токена
func Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/login" {
			c.Next()
			return
		}
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "No token provided"})
			c.Abort()

			return
		}

		tokenString = strings.ReplaceAll(tokenString, "Bearer ", "")

		_, err := jwt.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			c.Abort()

			return
		}

		c.Next()
	}
}
