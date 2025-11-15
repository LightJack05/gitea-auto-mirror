# Source configuration

## Environment config
You source Gitea server needs to be set up with the proper config variables in order to allow for the gitea-auto-mirror API to be called.
If you intend to configure Gitea via Environment variables, use the following variables:

```bash
# Below you should configure the hostname of the gitea auto mirror instance you deployed or will deploy
GITEA__webhook__ALLOWED_HOST_LIST=gitea-auto-mirror.foo.local
# Allow gitea to connect to other git servers on local networks
GITEA__migrations__ALLOW_LOCALNETWORKS=true
# Allow gitea to connect to your mirror server, in this case gitea-b:3000. Adapt this depending on your mirror server.
GITEA__migrations__ALLOWED_DOMAINS=gitea-b:3000
```

If you intend to use the configuration file for Gitea found in the apps data directory, refer to the [Gitea documentation](https://docs.gitea.com/administration/config-cheat-sheet).

## Setting up the webhook
In order to auto mirror your repositories, gitea needs to send a webhook request to your Gitea Auto Mirror instance. This will be done by using a webhook upon repo creation.

### Configuring at Site level
Configuring at site level will send a webhook request upon any repo create event.
You may want to set up regex filters on Gitea Auto Mirror if you choose this option.

On your gitea instance as the administrator, click your Icon in the top right and navigate to Site Administration>Integrations>Webhooks.
Now, add a system webhook with type "Gitea".
Under target URL, enter the URL of your Gitea Auto Mirror instance. 
Select "POST" as HTTP request method and "application/json".
Leave "Secret" empty.
Under "Trigger On" select "Custom Events..." then under "Repository Events" check "Repository".
Under "Authorization Header" enter the value that you will use to authenticate between gitea and Gitea Auto Mirror.
Finally click "Add Webhook".

### Configuring at Organization level

On your gitea instance, naviate to your organizaiton and then navigate to Settings>Webhooks.
Now, add a webhook with type "Gitea".
Under target URL, enter the URL of your Gitea Auto Mirror instance. 
Select "POST" as HTTP request method and "application/json".
Leave "Secret" empty.
Under "Trigger On" select "Custom Events..." then under "Repository Events" check "Repository".
Under "Authorization Header" enter the value that you will use to authenticate between gitea and Gitea Auto Mirror.
Finally click "Add Webhook".
