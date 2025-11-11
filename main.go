package main

import (
	"github.com/LightJack05/gitea-auto-mirror/internal/authentication"
	"github.com/LightJack05/gitea-auto-mirror/internal/config"
	"github.com/LightJack05/gitea-auto-mirror/internal/endpoints"
	"github.com/LightJack05/gitea-auto-mirror/internal/health"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration from environment variables
	config.LoadConfigFromEnv()
	// Start Gin router
	router := gin.Default()

	// Health check endpoint
	router.GET("/health", health.HealthCheck)

	// Webhook endpoints
	hooks := router.Group("/hooks/")
	hooks.Use(authentication.AuthenticationMiddleware())
	hooks.POST("/repo_create", endpoints.RepoCreatePost)

	router.Run()
}
