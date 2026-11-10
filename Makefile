.PHONY: help update validate

BINARY_NAME := scoopbucket
GO := go
# GOFLAGS := -v
# VERSION ?= $(shell cat .version/latest 2>/dev/null || echo "latest")
VALIDATION_SCRIPT := ./scripts/go-validate.sh

help:
	clear;
	@echo "Makefile usage help:"
	@echo "  make help          - Display this message"
	@echo "  make update        - Update module dependencies"
	@echo "  make validate      - Run code and LICENSE validation"
	
update:
	@echo ":: Updating module dependencies... "
	$(GO) get -u ./... && $(GO) mod tidy
	
validate:
	@bash -c $(VALIDATION_SCRIPT)
