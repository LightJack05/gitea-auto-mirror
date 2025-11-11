package config

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"regexp"

	"github.com/LightJack05/gitea-auto-mirror/internal/crypto"
)

var activeConfig Config
var configLoaded bool = false

// GetConfigLoaded Returns whether the configuration has been loaded
func GetConfigLoaded() bool {
	return configLoaded
}

// GetActiveConfig Returns the active configuration
func GetActiveConfig() Config {
	return activeConfig
}

// LoadConfigFromEnv Loads configuration from environment variables into the activeConfig variable
func LoadConfigFromEnv() {

	activeConfig.MirrorBaseUrl = os.Getenv("GITEA_AUTO_MIRROR_MIRROR_BASE_URL")
	activeConfig.MirrorUrlAppendDotGit = os.Getenv("GITEA_AUTO_MIRROR_MIRROR_URL_APPEND_DOT_GIT") == "true"
	activeConfig.MirrorUsername = os.Getenv("GITEA_AUTO_MIRROR_MIRROR_USERNAME")
	activeConfig.MirrorPassword = os.Getenv("GITEA_AUTO_MIRROR_MIRROR_PASSWORD")
	activeConfig.MirrorVerifyTLS = os.Getenv("GITEA_AUTO_MIRROR_MIRROR_VERIFY_TLS") != "false"
	activeConfig.MirrorSyncInterval = os.Getenv("GITEA_AUTO_MIRROR_MIRROR_SYNC_INTERVAL")
	activeConfig.SourceBaseUrl = os.Getenv("GITEA_AUTO_MIRROR_SOURCE_BASE_URL")
	activeConfig.SourceRepoRegExFilter = os.Getenv("GITEA_AUTO_MIRROR_SOURCE_REPO_REGEX_FILTER")
	activeConfig.SourceVerifyTLS = os.Getenv("GITEA_AUTO_MIRROR_SOURCE_VERIFY_TLS") != "false"
	activeConfig.SourceUsername = os.Getenv("GITEA_AUTO_MIRROR_SOURCE_USERNAME")
	activeConfig.SourcePassword = os.Getenv("GITEA_AUTO_MIRROR_SOURCE_PASSWORD")
	activeConfig.ApiPassword = os.Getenv("GITEA_AUTO_MIRROR_API_PASSWORD")
	activeConfig.AppDebugLogging = os.Getenv("GITEA_AUTO_MIRROR_APP_DEBUG_LOGGING") == "true"
	activeConfig.DisableConfigCheck = os.Getenv("GITEA_AUTO_MIRROR_DISABLE_CONFIG_CHECK") == "true"

	//If no sync interval is set, use Gitea default
	if activeConfig.MirrorSyncInterval == "" {
		activeConfig.MirrorSyncInterval = "8h0m0s" // Gitea default
	}

	//If there is no trailing slash on the URLs of source and mirror server, add it and log a warning
	if len(activeConfig.MirrorBaseUrl) > 0 && activeConfig.MirrorBaseUrl[len(activeConfig.MirrorBaseUrl)-1] != '/' {
		activeConfig.MirrorBaseUrl += "/"
		log.Println("WARNING: Added trailing slash to GITEA_AUTO_MIRROR_MIRROR_BASE_URL")
	}
	if len(activeConfig.SourceBaseUrl) > 0 && activeConfig.SourceBaseUrl[len(activeConfig.SourceBaseUrl)-1] != '/' {
		activeConfig.SourceBaseUrl += "/"
		log.Println("WARNING: Added trailing slash to GITEA_AUTO_MIRROR_SOURCE_BASE_URL")
	}

	// Parse API password hash if set
	passwordHashString := os.Getenv("GITEA_AUTO_MIRROR_API_PASSWORD_HASH")

	if passwordHashString != "" {
		passwordHash, err := crypto.ParseHash(passwordHashString)
		if err != nil {
			log.Panicf("Failed to parse GITEA_AUTO_MIRROR_API_PASSWORD_HASH: %s", err.Error())
		}
		activeConfig.ApiPasswordHash = passwordHash
	} else {
		activeConfig.ApiPasswordHash = nil
	}

	// Log loaded values
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
		"**redacted**",
		"**redacted**",
		activeConfig.AppDebugLogging,
		activeConfig.DisableConfigCheck,
	)

	// Log parsed hash details
	if activeConfig.ApiPasswordHash != nil {
		log.Printf(`Parsed hash from GITEA_AUTO_MIRROR_API_PASSWORD_HASH:
		Version: %d
		Time: %d
		Memory: %d
		Parallelism: %d
		Salt: %x
		Hash: %x
		`,
			activeConfig.ApiPasswordHash.Version,
			activeConfig.ApiPasswordHash.Time,
			activeConfig.ApiPasswordHash.Memory,
			activeConfig.ApiPasswordHash.Parallelism,
			activeConfig.ApiPasswordHash.Salt,
			activeConfig.ApiPasswordHash.Hash,
		)
	}

	if activeConfig.DisableConfigCheck {
		log.Println("WARNING: Configuration validation is disabled!")
		return
	}
	err := ValidateConfig(activeConfig)
	if err != nil {
		panic(fmt.Errorf("Config validation failed: %v", err))
	}
	configLoaded = true
}

func ValidateConfig(config Config) error {
	log.Println("Begining config validation...")
	// Required values validation
	err := validateRequiredParameters(config)
	if err != nil {
		return fmt.Errorf("Parameter validation failed: %v", err)
	}
	// URL validation
	err = validateURLs(config)
	if err != nil {
		return fmt.Errorf("URL validation failed: %v", err)
	}

	// API Password validation
	err = validateAuthValues(config)
	if err != nil {
		return fmt.Errorf("API Authentication credential validation failed: %v", err)
	}

	// Source credentials validation
	err = validateServerCredentials(config)
	if err != nil {
		return fmt.Errorf("Server credential validation failed: %v", err)
	}

	// Validate regex
	err = validateRegEx(config)
	if err != nil {
		return fmt.Errorf("RegEx validation failed: %v", err)
	}

	log.Println("Configuration validation complete.")
	return nil
}

func validateRequiredParameters(config Config) error {
	if config.MirrorBaseUrl == "" {
		return fmt.Errorf("GITEA_AUTO_MIRROR_MIRROR_BASE_URL is required.")
	}
	if config.SourceBaseUrl == "" {
		return fmt.Errorf("GITEA_AUTO_MIRROR_SOURCE_BASE_URL is required.")
	}
	if config.SourceUsername == "" {
		return fmt.Errorf("GITEA_AUTO_MIRROR_SOURCE_USERNAME is required.")
	}
	if config.SourcePassword == "" {
		return fmt.Errorf("GITEA_AUTO_MIRROR_SOURCE_PASSWORD is required.")
	}
	if config.MirrorUsername == "" {
		return fmt.Errorf("GITEA_AUTO_MIRROR_MIRROR_USERNAME is required.")
	}
	if config.MirrorPassword == "" {
		return fmt.Errorf("GITEA_AUTO_MIRROR_MIRROR_PASSWORD is required.")
	}
	return nil
}

func validateRegEx(config Config) error {
	if !(config.SourceRepoRegExFilter == "") {
		_, err := regexp.Compile(config.SourceRepoRegExFilter)
		if err != nil {
			return fmt.Errorf("RegEx in GITEA_AUTO_MIRROR_SOURCE_REPO_REGEX_FILTER must compile. Invalid value: %v", err)
		}
	}
	return nil
}

func validateServerCredentials(config Config) error {
	if config.SourceUsername == "" || config.SourcePassword == "" {
		return fmt.Errorf("GITEA_AUTO_MIRROR_SOURCE_USERNAME and GITEA_AUTO_MIRROR_SOURCE_PASSWORD must be set!")
	}

	// Mirror credentials validation
	if config.MirrorUsername == "" || config.MirrorPassword == "" {
		return fmt.Errorf("GITEA_AUTO_MIRROR_MIRROR_USERNAME and GITEA_AUTO_MIRROR_MIRROR_PASSWORD must be set!")
	}
	return nil
}

func validateAuthValues(config Config) error {
	if config.ApiPassword == "" && config.ApiPasswordHash == nil {
		log.Println("WARNING: There is no API password or hash set. Continuing with unsecured API endpoint...")
	}

	if config.ApiPassword != "" && config.ApiPasswordHash != nil {
		return fmt.Errorf("GITEA_AUTO_MIRROR_API_PASSWORD and GITEA_AUTO_MIRROR_API_PASSWORD_HASH are mutually exclusive!")
	}
	return nil
}

func validateURLs(config Config) error {
	mirrorBaseURL, err := url.ParseRequestURI(config.MirrorBaseUrl)
	if err != nil || (mirrorBaseURL.Scheme != "http" && mirrorBaseURL.Scheme != "https") {
		return fmt.Errorf("Invalid GITEA_AUTO_MIRROR_MIRROR_BASE_URL in configuration: %v", config.MirrorBaseUrl)
	}

	sourceBaseURL, err := url.ParseRequestURI(config.SourceBaseUrl)
	if err != nil || (sourceBaseURL.Scheme != "http" && sourceBaseURL.Scheme != "https") {
		return fmt.Errorf("Invalid GITEA_AUTO_MIRROR_SOURCE_BASE_URL in configuration: %v", config.SourceBaseUrl)
	}
	return nil
}
