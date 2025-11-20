# Authentication

Gitea Auto Mirror can and should be configured to require an Authentication Header in the request in order to allow for authentication of the request server.

!!! danger "WARNING"
    Not configuring authentication for Gitea Auto Mirror might allow unauthenticated users to extract data from private repositories, by mirroring them to an attacker-controlled server. It is highly discouraged to run Gitea Auto Mirror with no authentication.

## Authentication via Plain-Text Password
To set up authentication using a plaintext password, simply add the password in plain-text to the GITEA_AUTO_MIRROR_API_PASSWORD environment variable. See [Configuration Variables](variables.md)

On the source server, add the password in the "Authentication Header" option on the configured Webhook.

## Authentication via Argon2id Password Hash
To protect the password and avoid leaking it via the container's env variables, secrets, etc., you can specify a hash instead. This hash will be used to validate incoming requests containing the plaintext password.

### Generate a Hash
On a Linux or POSIX-Compliant machine, make sure you have `bash`, `openssl` and `argon2` installed and in your path.
If you have them, run the following command in your shell:

```bash
bash -c 'salt=$(openssl rand -base64 16); read -s -p "Password: " pw; echo; echo -n "$pw" | argon2 "$salt" -id -t 2 -m 16 -p 1 -l 32 -e'
```

It will prompt you for a password. Enter (or paste) the password you would like to use for API authentication and press enter.

!!! hint "Hint"
    The password will not be echoed as you type.

### Apply the Hash for Authentication

The command will output a hash, which you can then set in the GITEA_AUTO_MIRROR_API_PASSWORD_HASH environment variable.
You should see the Gitea Auto Mirror container log the parsed hash details into the container's log.
