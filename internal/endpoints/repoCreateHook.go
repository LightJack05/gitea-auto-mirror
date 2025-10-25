package endpoints

import (
	"log"
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
		if config.GetActiveConfig().AppDebugLogging {
			log.Printf("Request for repo %s has been ignored due to regex filter.", createEvent.Repository.FullName)
		}
		c.Status(http.StatusNoContent)
		return
	}

	if config.GetActiveConfig().AppDebugLogging {
		log.Printf("Processing hook for repo %s", createEvent.Repository.FullName)
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
