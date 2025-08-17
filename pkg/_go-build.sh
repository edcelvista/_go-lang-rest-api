#!/bin/bash
SCRIPT_PATH="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$SCRIPT_PATH" APP_ENV=development go test controller/*

GOOS=linux GOARCH=amd64 go build -o bin/app
GOOS=linux GOARCH=arm64 go build -o bin/app-arm64
# go mod vendor # By default, Go caches dependencies globally. However, you can vendor them inside your project: