package config

import (
	"log"
	"net/url"
	"os"
	"regexp"
)

var activeConfig Config

func GetActiveConfig() Config {
	return activeConfig
}

func LoadConfigFromEnv() {

	activeConfig.MirrorBaseUrl = os.Getenv("GITEA_AUTO_MIRROR_MIRROR_BASE_URL")
	activeConfig.MirrorUrlAppendDotGit = os.Getenv("GITEA_AUTO_MIRROR_MIRROR_URL_APPEND_DOT_GIT") == "true"
	activeConfig.MirrorUsername = os.Getenv("GITEA_AUTO_MIRROR_MIRROR_USERNAME")
	activeConfig.MirrorPassword = os.Getenv("GITEA_AUTO_MIRROR_MIRROR_PASSWORD")
	activeConfig.MirrorVerifyTLS = os.Getenv("GITEA_AUTO_MIRROR_MIRROR_VERIFY_TLS") != "false"
	activeConfig.SourceBaseUrl = os.Getenv("GITEA_AUTO_MIRROR_SOURCE_BASE_URL")
	activeConfig.SourceRepoRegExFilter = os.Getenv("GITEA_AUTO_MIRROR_SOURCE_REPO_REGEX_FILTER")
	activeConfig.SourceVerifyTLS = os.Getenv("GITEA_AUTO_MIRROR_SOURCE_VERIFY_TLS") != "false"
	activeConfig.SourceUsername = os.Getenv("GITEA_AUTO_MIRROR_SOURCE_USERNAME")
	activeConfig.SourcePassword = os.Getenv("GITEA_AUTO_MIRROR_SOURCE_PASSWORD")
	activeConfig.ApiPasswordHash = os.Getenv("GITEA_AUTO_MIRROR_API_PASSWORD_HASH")
	activeConfig.ApiPassword = os.Getenv("GITEA_AUTO_MIRROR_API_PASSWORD")
	activeConfig.AppDebugLogging = os.Getenv("GITEA_AUTO_MIRROR_APP_DEBUG_LOGGING") == "true"
	activeConfig.DisableConfigCheck = os.Getenv("GITEA_AUTO_MIRROR_DISABLE_CONFIG_CHECK") == "true"

	//If there is no trailing slash on the URLs of source and mirror server, add it and log a warning
	if len(activeConfig.MirrorBaseUrl) > 0 && activeConfig.MirrorBaseUrl[len(activeConfig.MirrorBaseUrl)-1] != '/' {
		activeConfig.MirrorBaseUrl += "/"
		log.Println("WARNING: Added trailing slash to GITEA_AUTO_MIRROR_MIRROR_BASE_URL")
	}
	if len(activeConfig.SourceBaseUrl) > 0 && activeConfig.SourceBaseUrl[len(activeConfig.SourceBaseUrl)-1] != '/' {
		activeConfig.SourceBaseUrl += "/"
		log.Println("WARNING: Added trailing slash to GITEA_AUTO_MIRROR_SOURCE_BASE_URL")
	}

	log.Printf(`Loaded values from Environment:
	GITEA_AUTO_MIRROR_MIRROR_BASE_URL=%s
	GITEA_AUTO_MIRROR_MIRROR_URL_APPEND_DOT_GIT=%t
	GITEA_AUTO_MIRROR_MIRROR_USERNAME=%s
	GITEA_AUTO_MIRROR_MIRROR_PASSWORD=%s
	GITEA_AUTO_MIRROR_MIRROR_VERIFY_TLS=%t
	GITEA_AUTO_MIRROR_SOURCE_BASE_URL=%s
	GITEA_AUTO_MIRROR_SOURCE_REPO_REGEX_FILTER=%s
	GITEA_AUTO_MIRROR_SOURCE_VERIFY_TLS=%t
	GITEA_AUTO_MIRROR_SOURCE_USERNAME=%s
	GITEA_AUTO_MIRROR_SOURCE_PASSWORD=%s
	GITEA_AUTO_MIRROR_API_PASSWORD_HASH=%s
	GITEA_AUTO_MIRROR_API_PASSWORD=%s
	GITEA_AUTO_MIRROR_APP_DEBUG_LOGGING=%t
	GITEA_AUTO_MIRROR_DISABLE_CONFIG_CHECK=%t
	`,
		activeConfig.MirrorBaseUrl,
		activeConfig.MirrorUrlAppendDotGit,
		"**redacted**",
		"**redacted**",
		activeConfig.MirrorVerifyTLS,
		activeConfig.SourceBaseUrl,
		activeConfig.SourceRepoRegExFilter,
		activeConfig.SourceVerifyTLS,
		"**redacted**",
		"**redacted**",
		activeConfig.ApiPasswordHash,
		activeConfig.ApiPassword,
		activeConfig.AppDebugLogging,
		activeConfig.DisableConfigCheck,
	)

	if activeConfig.DisableConfigCheck {
		log.Println("WARNING: Configuration validation is disabled!")
		return
	}
	ValidateConfig(activeConfig)
}

func ValidateConfig(config Config) {
	log.Println("Begining config validation...")
	// URL validation
	mirrorBaseURL, err := url.ParseRequestURI(config.MirrorBaseUrl)
	if err != nil || (mirrorBaseURL.Scheme != "http" && mirrorBaseURL.Scheme != "https") {
		panic("Invalid GITEA_AUTO_MIRROR_MIRROR_BASE_URL in configuration: " + config.MirrorBaseUrl)
	}

	sourceBaseURL, err := url.ParseRequestURI(config.SourceBaseUrl)
	if err != nil || (sourceBaseURL.Scheme != "http" && sourceBaseURL.Scheme != "https") {
		panic("Invalid GITEA_AUTO_MIRROR_SOURCE_BASE_URL in configuration: " + config.SourceBaseUrl)
	}

	// API Password validation
	if config.ApiPassword == "" && config.ApiPasswordHash == "" {
		log.Println("WARNING: There is no API password or hash set. Continuing with unsecured API endpoint...")
	}

	if config.ApiPassword != "" && config.ApiPasswordHash != "" {
		panic("GITEA_AUTO_MIRROR_API_PASSWORD and GITEA_AUTO_MIRROR_API_PASSWORD_HASH are mutually exclusive!")
	}

	// Source credentials validation
	if config.SourceUsername == "" || config.SourcePassword == "" {
		panic("GITEA_AUTO_MIRROR_SOURCE_USERNAME and GITEA_AUTO_MIRROR_SOURCE_PASSWORD must be set!")
	}

	// Mirror credentials validation
	if config.MirrorUsername == "" || config.MirrorPassword == "" {
		panic("GITEA_AUTO_MIRROR_MIRROR_USERNAME and GITEA_AUTO_MIRROR_MIRROR_PASSWORD must be set!")
	}

	// Validate regex
	if !(config.SourceRepoRegExFilter == "") {
		_, err := regexp.Compile(config.SourceRepoRegExFilter)
		if err != nil {
			panic("RegEx in GITEA_AUTO_MIRROR_SOURCE_REPO_REGEX_FILTER must compile. Invalid value: " + err.Error())
		}
	}

	log.Println("Configuration validation complete.")

}
