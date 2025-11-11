# Docker compose

Installation with docker compose can be done on standard servers normally with docker. 
Install Docker compose following docker's [guide](https://docs.docker.com/compose/install/linux/#install-using-the-repository).

## Compose file
Place a file called `docker-compose.yaml` somewhere on your server with these contents:
```yaml
services:
  gitea-auto-mirror:
    image: ghcr.io/lightjack05/gitea-auto-mirror:latest
    container_name: gitea-auto-mirror
    environment:
      - GITEA_AUTO_MIRROR_MIRROR_BASE_URL=
      - GITEA_AUTO_MIRROR_MIRROR_URL_APPEND_DOT_GIT=false
      - GITEA_AUTO_MIRROR_MIRROR_VERIFY_TLS=true
      - GITEA_AUTO_MIRROR_SOURCE_BASE_URL=
      - GITEA_AUTO_MIRROR_SOURCE_REPO_REGEX_FILTER=
      - GITEA_AUTO_MIRROR_SOURCE_VERIFY_TLS=true
      - GITEA_AUTO_MIRROR_MIRROR_SYNC_INTERVAL=3h
      - GITEA_AUTO_MIRROR_MIRROR_USERNAME=user
      - GITEA_AUTO_MIRROR_MIRROR_PASSWORD=password
      - GITEA_AUTO_MIRROR_SOURCE_USERNAME=user
      - GITEA_AUTO_MIRROR_SOURCE_PASSWORD=password
      - GITEA_AUTO_MIRROR_API_PASSWORD_HASH=
    ports:
      - 8080:8080
```

## Configure Gitea Auto Mirror
Configure Gitea Auto Mirror as per the [configuration guide](../configuring/index.md).

## Start the container
Start the docker container via `docker compose up -d`. You can verify if the container is up and running using `docker ps` and `docker logs gitea-auto-mirror`.
