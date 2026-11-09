.PHONY: help build update clean validate

BINARY_NAME := scoopbucket
GO := go
GOFLAGS := -v
VERSION ?= $(shell cat .version/latest 2>/dev/null || echo "latest")
BUILD_DIR := ./bin
VALIDATION_SCRIPT := ./scripts/go-validate.sh

help:
	clear;
	@echo "Makefile usage help:"
	@echo "  make help          - Display this message"
	@echo "  make build         - Build the binary"
	@echo "  make update        - Update module dependencies"
	@echo "  make clean         - Remove build artifacts"
	@echo "  make validate      - Run code and LICENSE validation"
	
build:
	@echo ":: Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	$(GO) build $(GOFLAGS) -ldflags="-X main.version=$(VERSION)" -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/$(BINARY_NAME)

update:
	@echo ":: Updating module dependencies... "
	$(GO) get -u ./... && $(GO) mod tidy
	
clean:
	@echo ":: Cleaning build artifacts..."
	rm -rf $(BUILD_DIR)

validate:
	@bash -c $(VALIDATION_SCRIPT)
	
