#!/usr/bin/env bash

BLUE='\033[0;34m'
RED='\033[0;31m'
NC='\033[0m'

logger_error() {
	echo -e "${RED}[ERR]${NC} $1" >&2
}

validate_license() {
	if ! command -v "go-licenses" >/dev/null 2>&1; then
		go install github.com/google/go-licenses/v2@latest
	fi

	go-licenses check ./... \
		--disallowed_types=forbidden,restricted \
		--ignore=golang.org/x/sys

	echo ""
}

if [[ ! -f "./go.mod" ]]; then
	logger_error ":: Not a Go module directory, exiting..."
	exit 1
fi

if ! command -v "go" >/dev/null 2>&1; then
	logger_error "command not found: go"
	exit 1
fi

if [[ -f "./LICENSE" ]]; then
	validate_license
fi

if command -v "golangci-lint" >/dev/null 2>&1; then
	golangci-lint run --fix
fi
