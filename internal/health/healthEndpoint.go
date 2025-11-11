package health

import (
	"net/http"

	"github.com/LightJack05/gitea-auto-mirror/internal/config"
	"github.com/gin-gonic/gin"
)

// HealthCheck Handles health check requests
func HealthCheck(c *gin.Context) {

	if config.GetConfigLoaded() == false {
		c.Status(http.StatusServiceUnavailable)
		return
	}
	if (config.GetActiveConfig() == config.Config{}) {
		c.Status(http.StatusServiceUnavailable)
		return
	}

	// All checks passed
	c.Status(http.StatusOK)
}
