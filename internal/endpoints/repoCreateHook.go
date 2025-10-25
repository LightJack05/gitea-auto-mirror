package endpoints

import (
	"encoding/json"
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

	requestBody := createPushMirrorRequestBody(createEvent.Repository.FullName)

	requestJson, err := json.Marshal(requestBody)
	if err != nil {
		log.Println("ERROR: Could not marshal request Go struct to JSON string. Is the application configured correctly?")
	}
	//TODO: Send http request
	log.Println(string(requestJson))

	c.Status(http.StatusNotImplemented)
}

func shouldModifyRepo(repoPath string) bool {
	if config.GetActiveConfig().SourceRepoRegExFilter == "" {
		return true
	}

	r := regexp.MustCompile(config.GetActiveConfig().SourceRepoRegExFilter)

	return r.MatchString(repoPath)
}

func createPushMirrorRequestBody(repoPath string) datastructures.RepoCreatePushMirrorBody {
	requestBody := datastructures.RepoCreatePushMirrorBody{
		//TODO: Add a config option for this
		Interval:       "8h0m0s",
		RemoteAddress:  config.GetActiveConfig().MirrorBaseUrl + repoPath,
		RemoteUsername: config.GetActiveConfig().MirrorUsername,
		RemotePassword: config.GetActiveConfig().MirrorPassword,
		SyncOnCommit:   true,
	}
	if config.GetActiveConfig().MirrorUrlAppendDotGit {
		requestBody.RemoteAddress = requestBody.RemoteAddress + string(".git")
	}
	return requestBody
}
