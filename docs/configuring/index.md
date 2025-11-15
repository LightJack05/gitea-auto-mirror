# Configuration
You will need to configure both gitea and gitea-auto-mirror, as well as your remote for this to work.

## Configuring Gitea (source)
Your source server will need to be configured to trust the webhook URL, you will need to create a webhook entry on organization or site level, and you need to configure credentials for authenticating with the source server.

Detailed source instance configuration is described in [Configuring the source server](source.md).

## Configuring Gitea Auto Mirror
Gitea auto mirror needs to be configured properly to provide a secure communication between the api and the source server, to properly forward the requests and match the correct repositories, and to make sure that gitea auto mirror pushes the correct credentails to the remote.

Detailed configuration options for gitea auto mirror can be found here in [Configuring Gitea Auto Mirror](gitea-auto-mirror.md).

## Configuring the mirror server
The mirror server must be configured to allow for "push to create", which is usually an available option on most git servers.
An example for gitea as mirror can be found under [Configuring the upstream mirror](mirror.md).

