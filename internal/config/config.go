package config

import (
	"github.com/LightJack05/gitea-auto-mirror/internal/crypto"
)

// Config The configuration object loaded from environment variables
type Config struct {
	// The base URL of the mirror server
	// Format as http[s]://<hostname>[:<port>]/
	// Example: http://mirror.example.com:3000/
	// Example: https://mirror.example.com/
	MirrorBaseUrl string

	// Whether to append .git to the end of the mirror URL
	// Example: true -> http://mirror.example.com:3000/user/repo.git
	// Example: false -> http://mirror.example.com:3000/user/repo
	MirrorUrlAppendDotGit bool

	// Username for authenticating to the mirror server
	// Used in the created mirror setting on the repository.
	MirrorUsername string

	// Password or Token for authenticating to the mirror server
	// Used in the created mirror setting on the repository.
	MirrorPassword string

	// Whether to verify TLS certificates when connecting to the mirror server
	MirrorVerifyTLS bool

	// The interval between mirror syncs, in Go duration format
	// Will be applied to the repo push mirror entry in Gitea 
	// Example: 10m for 10 minutes, 1h for 1 hour, 2h15m for 2 hours and 15 minutes
	//  Default is empty string, which means the default interval of Gitea is used (8h)
	MirrorSyncInterval string

	// The base URL of the source server
	// Format as http[s]://<hostname>[:<port>]/
	// Example: http://source.example.com:3000/
	// Example: https://source.example.com/
	SourceBaseUrl string

	// A regex filter to apply to source repository full names (owner/repo)
	// Only repositories matching this regex will be mirrored
	// Example: ^myorg/.*$ to only mirror repositories in the "myorg" organization
	// Hint: beware of escaping when setting regexes in environment variables!
	SourceRepoRegExFilter string

	// Whether to verify TLS certificates when connecting to the source server
	SourceVerifyTLS bool

	// Username for authenticating to the source server
	SourceUsername string

	// Password or Token for authenticating to the source server
	SourcePassword string

	// Password hash for authenticating incoming API requests
	// Mutually exclusive with ApiPassword
	ApiPasswordHash *crypto.Argon2idPasswordHash 

	// Plaintext password for authenticating incoming API requests
	// Mutually exclusive with ApiPasswordHash
	ApiPassword string

	// Whether to enable debug logging throughout the application
	AppDebugLogging bool
	// Disable config validation
	DisableConfigCheck bool
	
}
