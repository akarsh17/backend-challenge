package middleware

import (
	"backend-challenge/config"
	"backend-challenge/pkg/errors"
	response "backend-challenge/pkg/utils"

	"github.com/gin-gonic/gin"
)

func APIKeyAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("api_key")
		expectedKey := config.AppConfig.APIKey
		if apiKey == "" || apiKey != expectedKey {
			response.Error(c, errors.UnauthorizedError("Invalid API key"))
			c.Abort()
			return
		}
		c.Next()
	}
}
