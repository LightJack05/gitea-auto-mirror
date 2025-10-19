FROM golang:alpine AS build

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o /app/gitea-auto-mirror

FROM alpine:latest

WORKDIR /app
COPY --from=build /app/gitea-auto-mirror .

ENTRYPOINT ["./gitea-auto-mirror"]

