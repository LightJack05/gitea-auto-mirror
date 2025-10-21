package main

import (
	"github.com/LightJack05/gitea-auto-mirror/internal/endpoints"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/repo_create_hook", endpoints.RepoCreatePost)
	router.Run()
}
