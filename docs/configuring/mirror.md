# Configuring the mirror server

Your mirror server only needs to support 'push to create', to allow you to push a non-existing repo and create it like that. 

## Gitea Mirrors
If you want to mirror to a Gitea server, on your mirror server, set the following environment variables:
```bash
GITEA__repository__ENABLE_PUSH_CREATE_ORG=true
GITEA__repository__ENABLE_PUSH_CREATE_USER=true
```
This will enable push to create for users and organizations.
