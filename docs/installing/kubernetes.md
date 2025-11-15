# Kubernetes
Since Gitea Auto Mirror is stateless, it can be easily deployed into Kubernetes using a Deployment.

Apply this manifest using `kubectl apply -f` or via your favorite CI/CD tool.
```yaml
# NOTE: You probably want to use a secret operator or sealed secret to put this into your cluster.
apiVersion: v1
kind: Secret
metadata:
  name: gitea-auto-mirror-secret
type: Opaque
stringData:
  GITEA_AUTO_MIRROR_MIRROR_USERNAME: "user"
  GITEA_AUTO_MIRROR_MIRROR_PASSWORD: "password"
  GITEA_AUTO_MIRROR_SOURCE_USERNAME: "user"
  GITEA_AUTO_MIRROR_SOURCE_PASSWORD: "password"
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: gitea-auto-mirror
  labels:
    app: gitea-auto-mirror
spec:
  replicas: 3
  selector:
    matchLabels:
      app: gitea-auto-mirror
  template:
    metadata:
      labels:
        app: gitea-auto-mirror
    spec:
      containers:
        - name: gitea-auto-mirror
          image: ghcr.io/lightjack05/gitea-auto-mirror:latest
          ports:
            - containerPort: 8080
          env:
            - name: GITEA_AUTO_MIRROR_MIRROR_BASE_URL
              value: ""
            - name: GITEA_AUTO_MIRROR_MIRROR_URL_APPEND_DOT_GIT
              value: "false"
            - name: GITEA_AUTO_MIRROR_MIRROR_VERIFY_TLS
              value: "true"
            - name: GITEA_AUTO_MIRROR_SOURCE_BASE_URL
              value: ""
            - name: GITEA_AUTO_MIRROR_SOURCE_REPO_REGEX_FILTER
              value: ""
            - name: GITEA_AUTO_MIRROR_SOURCE_VERIFY_TLS
              value: "true"
            - name: GITEA_AUTO_MIRROR_MIRROR_SYNC_INTERVAL
              value: "3h"
            - name: GITEA_AUTO_MIRROR_API_PASSWORD_HASH
              value: ""
            - name: GITEA_AUTO_MIRROR_MIRROR_USERNAME
              valueFrom:
                secretKeyRef:
                  name: gitea-auto-mirror-secret
                  key: GITEA_AUTO_MIRROR_MIRROR_USERNAME
            - name: GITEA_AUTO_MIRROR_MIRROR_PASSWORD
              valueFrom:
                secretKeyRef:
                  name: gitea-auto-mirror-secret
                  key: GITEA_AUTO_MIRROR_MIRROR_PASSWORD
            - name: GITEA_AUTO_MIRROR_SOURCE_USERNAME
              valueFrom:
                secretKeyRef:
                  name: gitea-auto-mirror-secret
                  key: GITEA_AUTO_MIRROR_SOURCE_USERNAME
            - name: GITEA_AUTO_MIRROR_SOURCE_PASSWORD
              valueFrom:
                secretKeyRef:
                  name: gitea-auto-mirror-secret
                  key: GITEA_AUTO_MIRROR_SOURCE_PASSWORD
---
apiVersion: v1
kind: Service
metadata:
  name: gitea-auto-mirror
  labels:
    app: gitea-auto-mirror
spec:
  selector:
    app: gitea-auto-mirror
  ports:
    - protocol: TCP
      port: 8080
      targetPort: 8080
  type: ClusterIP
```
You may also need to create network policies or an ingress if you intend to expose gitea-auto-mirror outside of your cluster.

See the [configuration guide](../configuring/index.md) to find out how to assign the environment variables, and which ones are available.
