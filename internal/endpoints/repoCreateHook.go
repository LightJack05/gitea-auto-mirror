package endpoints

import (
	"net/http"
	"regexp"

	"github.com/LightJack05/gitea-auto-mirror/internal/config"
	"github.com/LightJack05/gitea-auto-mirror/internal/datastructures"

	"github.com/gin-gonic/gin"
)

func RepoCreatePost(c *gin.Context) {
	var createEvent datastructures.RepoCreateEvent
	if err := c.BindJSON(&createEvent); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if !shouldModifyRepo(createEvent.Repository.FullName) {
		c.Status(http.StatusNoContent)
		return
	}

	c.Status(http.StatusNotImplemented)
}

func shouldModifyRepo(repoPath string) bool {
	if config.GetActiveConfig().SourceRepoRegExFilter == "" {
		return true
	}

	r := regexp.MustCompile(config.GetActiveConfig().SourceRepoRegExFilter)

	return r.MatchString(repoPath)
}

