# Gitea Auto Mirror

Gitea Auto Mirror is a web API that is used to automatically create push-mirror entries in one gitea instance, to then mirror it's repository contents to another git server.

## How it works

Gitea Auto Mirror listens for webhook events triggered by the gitea server, and responds to them by POSTing back to the Gitea instance and creating a push mirror entry.
That means that you set up a webhook on your gitea instance. This webhook will trigger when a repository is created, and send a webrequest to Gitea Auto Mirror, which will check if the repository matches it's RegEx filter. If it does, it will send a POST request back to Gitea, modifying the newly created repository and triggering a push to another git server. From there on out, Gitea itself will handle pushing new changes to the mirror server.

> [!NOTE]
> Gitea Auto Mirror only works with HTTP based mirror servers, as other are [not supported by Gitea](https://docs.gitea.com/usage/repo-mirror#mirror-an-existing-ssh-repository).

## Setup

Refer to the [installation guide](installing/index.md) and [configuration guide](configuring/index.md) to find out how to set up Gitea Auto Mirror.
