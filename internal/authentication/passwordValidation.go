package authentication

// Use this to generate hashes on the CLI:
// bash -c 'salt=$(openssl rand -base64 16); read -s -p "Password: " pw; echo; echo -n "$pw" | argon2 "$salt" -id -t 2 -m 16 -p 1 -l 32 -e'

import (
	"github.com/LightJack05/gitea-auto-mirror/internal/config"
	//"golang.org/x/crypto/argon2"
)

func ValidateRequestAuthHeader(headerContent string) bool {
	if(config.GetActiveConfig().ApiPassword != "") {
		return config.GetActiveConfig().ApiPassword == headerContent;
	}
	return validateAgainstHash(headerContent, config.GetActiveConfig().ApiPasswordHash)
}

func validateAgainstHash(headerContent string, hash string) bool {
	panic("Not implemented")
}


