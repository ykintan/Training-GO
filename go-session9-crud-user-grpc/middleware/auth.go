package middleware

import (
	"net/http"
	"training-go/go-session9-crud-user-grpc/config"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, password, ok := c.Request.BasicAuth()
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization basic token is required"})
			c.Abort()
			return
		}

		isValid := (username == config.AuthBasicUsername) && (password == config.AuthBasicPassword)
		if !isValid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
			c.Abort()
			return
		}

		c.Next()
	}
}
