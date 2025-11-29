# Gitea Auto Mirror

A web API that automatically creates push-mirror entries when a new repository is created on a Gitea instance, mirroring its contents to another git server.

## Overview

Gitea Auto Mirror listens for webhook events triggered by a Gitea server and responds by creating a push mirror entry. When a repository is created, the webhook sends a request to Gitea Auto Mirror, which checks if the repository matches its RegEx filter. If it matches, a POST request is sent back to Gitea to configure the push mirror, and from there Gitea handles pushing changes to the mirror server.

## Documentation

For detailed setup instructions, configuration options, and usage guides, visit the full documentation:

📚 **[https://lightjack05.github.io/gitea-auto-mirror/](https://lightjack05.github.io/gitea-auto-mirror/)**

## Features

- Automatic push mirror creation for new repositories
- RegEx-based repository filtering
- Configurable sync intervals
- Docker container deployment
- Support for Kubernetes and Docker Compose installations
- TLS verification options for secure connections

## Quick Start

Gitea Auto Mirror is available as a Docker container from [GitHub Container Registry](https://github.com/LightJack05/gitea-auto-mirror/pkgs/container/gitea-auto-mirror).

```bash
docker pull ghcr.io/lightjack05/gitea-auto-mirror:latest
```

For installation and configuration details, see the [documentation](https://lightjack05.github.io/gitea-auto-mirror/).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
