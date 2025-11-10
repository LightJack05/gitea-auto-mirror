package authentication

import (
	"log"

	"github.com/LightJack05/gitea-auto-mirror/internal/config"
	"github.com/gin-gonic/gin"
)

func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !ValidateRequestAuthHeader(authHeader) {
			if(config.GetActiveConfig().AppDebugLogging) {
				log.Printf("Unauthorized request: Invalid Authorization header: %s\n", authHeader)
			}
			c.AbortWithStatus(401)
			return
		}
		c.Next()
	}
}
