package authentication

// Use this to generate hashes on the CLI:
// bash -c 'salt=$(openssl rand -base64 16); read -s -p "Password: " pw; echo; echo -n "$pw" | argon2 "$salt" -id -t 2 -m 16 -p 1 -l 32 -e'
// Hash will look something like this:
// $argon2id$v=19$m=65536,t=2,p=1$WFowVzArejQrWFNuTWExVlN2UWFpZz09$rM+SL2W85Jw5ZQt/XbxKWL0V5YDfrI+ON9w1yxxGlVk

import (
	"github.com/LightJack05/gitea-auto-mirror/internal/config"
	"github.com/LightJack05/gitea-auto-mirror/internal/crypto"
)

func ValidateRequestAuthHeader(headerContent string) bool {
	if(config.GetActiveConfig().ApiPassword != "") {
		return config.GetActiveConfig().ApiPassword == headerContent;
	}
	return validateAgainstHash(headerContent, config.GetActiveConfig().ApiPasswordHash)
}

func validateAgainstHash(headerContent string, hash *crypto.Argon2idPasswordHash) bool {
	panic("Not Implemented.")
}

