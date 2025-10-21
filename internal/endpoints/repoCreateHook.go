package endpoints

import (
	"net/http"

	"github.com/LightJack05/gitea-auto-mirror/internal/datastructures"

	"github.com/gin-gonic/gin"
)

func RepoCreatePost(c *gin.Context) {
	var createEvent datastructures.Event
	if err := c.BindJSON(&createEvent); err != nil {
		c.Status(http.StatusInternalServerError)
	}

	c.Status(http.StatusNotImplemented)
}
