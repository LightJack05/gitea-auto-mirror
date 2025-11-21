# Configuration Variables 

Gitea Auto Mirror is configured via environment variables.

## Configuring the Mirror

| Variable | Required | Default | Type | Description | Example |
| --------------- | --------------- | --------------- | --------------- | --------------- | --------------- |
| GITEA_AUTO_MIRROR_MIRROR_BASE_URL | yes | N/A | URL (aka string) | The URL of the upstream mirror | http://mirror.example.com:3000/ <br> https://mirror.example.com/ |
| GITEA_AUTO_MIRROR_MIRROR_URL_APPEND_DOT_GIT | no | Bool | false | Whether to append a .git to the repository URL in the mirror entry | true <br> false |
| GITEA_AUTO_MIRROR_MIRROR_USERNAME | yes | N/A | string | The mirror server username to use for mirroring | username |
| GITEA_AUTO_MIRROR_MIRROR_PASSWORD | yes | N/A | string | The mirror server password or token to use for mirroring | Password123 |
| GITEA_AUTO_MIRROR_MIRROR_VERIFY_TLS | no | true | Bool | Whether to verify TLS certificates on the remote | false |
| GITEA_AUTO_MIRROR_MIRROR_SYNC_INTERVAL | no | 8h | Go time format (aka string) | The interval between syncs if no push is performed | 3h10m |


## Source Configuration
| Variable | Required | Default | Type | Description | Example |
| --------------- | --------------- | --------------- | --------------- | --------------- | --------------- |
| GITEA_AUTO_MIRROR_SOURCE_BASE_URL     | yes | N/A | string | The URL of the source Gitea server | http://source.example.com:3000/ |
| GITEA_AUTO_MIRROR_SOURCE_VERIFY_TLS   | no  | true| string | Whether to verify TLS certificates for the source server | true <br> false |
| GITEA_AUTO_MIRROR_SOURCE_USERNAME     | yes | N/A | string | The username for the source Gitea server | username |
| GITEA_AUTO_MIRROR_SOURCE_PASSWORD     | yes | N/A | string | The password for the source Gitea server | Password123 |

## Authentication Settings

| Variable | Required | Default | Type | Description | Example |
| --------------- | --------------- | --------------- | --------------- | --------------- | --------------- |
| GITEA_AUTO_MIRROR_API_PASSWORD_HASH   | no  | ""  | string | The hash to use for password authentication between the source server and Gitea Auto Mirror (See [Configuring Authentication](authentication.md)) | $argon2id$v=19$m=65536,t=2,p=1$aVdKZ3B3djJXcDIydnlKZjJ0L3RWUT09$k8ZJITqnD6n8C9tavbZX8rv6pO6mbvIi/Lpzt8V0ZuY |
| GITEA_AUTO_MIRROR_API_PASSWORD        | no  | ""  | string | The plaintext password to use for authentication between the source server and Gitea Auto Mirror (See [Configuring Authentication](authentication.md)) | Password123 |

## Repository Filtering

| Variable | Required | Default | Type | Description | Example |
| --------------- | --------------- | --------------- | --------------- | --------------- | --------------- |
| GITEA_AUTO_MIRROR_SOURCE_REPO_REGEX_FILTER | no | "" | RegEx (aka string) | A regex filter to apply to repository names that should be included in filtering, matches repository names only (owner/repo) | ^myorg/.*$ |

!!! info "Hint"
    Beware of escaping when setting RegExes through environment variables.

## Debug options

| Variable | Required | Default | Type | Description | Example |
| --------------- | --------------- | --------------- | --------------- | --------------- | --------------- |
| GITEA_AUTO_MIRROR_APP_DEBUG_LOGGING | no | false | Boolean | Enable advanced logging in Gitea Auto Mirror | true |
| GITEA_AUTO_MIRROR_DISABLE_CONFIG_CHECK | no | false | Boolean | Disable advanced config checks for debugging purposes | true |
