FROM golang:alpine AS build

WORKDIR /app

COPY go.mod .
RUN go mod download

COPY . .
RUN go build -o /app/gitea-auto-mirror

FROM alpine:latest

WORKDIR /app
COPY LICENSES /licenses
COPY --from=build /app/gitea-auto-mirror .

ENTRYPOINT ["./gitea-auto-mirror"]

