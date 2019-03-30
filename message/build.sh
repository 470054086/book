#!/usr/bin/env bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build cmd/main.go
docker build -t message-service:latest .

