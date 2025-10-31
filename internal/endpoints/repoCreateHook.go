package endpoints

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/LightJack05/gitea-auto-mirror/internal/config"
	"github.com/LightJack05/gitea-auto-mirror/internal/datastructures"

	"github.com/gin-gonic/gin"
)

const repoApiPathFormat string = "%sapi/v1/repos/%s/%s/push_mirrors"

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

	repoApiUrl := buildRepoApiUrl(createEvent)
	requestBodyJson := marshalRequestBody(requestBody)

	if config.GetActiveConfig().AppDebugLogging {
		log.Printf("Sending request to %s with body: %s", repoApiUrl, requestBodyJson)
	}

	err := addMirrorToRepo(requestBodyJson, repoApiUrl)
	if err != nil {
		log.Printf("Modifying repository on upstream failed: %s", err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func marshalRequestBody(createEvent datastructures.RepoCreatePushMirrorBody) string {
	jsonString, err := json.Marshal(createEvent)
	if err != nil {
		panic(err)
	}
	return string(jsonString)
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
		Interval:       config.GetActiveConfig().MirrorSyncInterval,
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

func buildRepoApiUrl(createEvent datastructures.RepoCreateEvent) string {
	return fmt.Sprintf(repoApiPathFormat, config.GetActiveConfig().SourceBaseUrl, createEvent.Repository.Owner.Login, createEvent.Repository.Name)
}

func addMirrorToRepo(requestBodyJson string, repoApiUrl string) error {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("POST", repoApiUrl, bytes.NewBuffer([]byte(requestBodyJson)))
	if err != nil {
		panic(err)
	}

	req.Header.Add("Authorization", "token "+config.GetActiveConfig().SourcePassword)
	req.Header.Add("content-type", "application/json;charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Failed to do request to source URL %s : %s", repoApiUrl, err.Error())
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("Error: Request to %s returned non-ok status code: %s", repoApiUrl, resp.Status)
		log.Printf("Error: Details (if available): %s", responseBody)
		return fmt.Errorf("Upstream server returned non-ok status: %s, body: %s", resp.Status, responseBody)
	}
	return nil
}
