package main

import (
	"fmt"

	"github.com/LightJack05/gitea-auto-mirror/internal/authentication"
	"github.com/LightJack05/gitea-auto-mirror/internal/config"
	"github.com/LightJack05/gitea-auto-mirror/internal/endpoints"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfigFromEnv()
	router := gin.Default()
	router.POST("/repo_create_hook", endpoints.RepoCreatePost)
	router.Run()
}
