#!/bin/bash
# go run main.go
SCRIPT_PATH="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$SCRIPT_PATH" APP_ENV=development go run main.go