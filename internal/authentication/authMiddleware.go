package authentication

import (
	"log"
	"strings"

	"github.com/LightJack05/gitea-auto-mirror/internal/config"
	"github.com/gin-gonic/gin"
)

// AuthenticationMiddleware Validates the auth header based on the application config loaded on startup
func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := strings.TrimSpace(c.GetHeader("Authorization"))
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
