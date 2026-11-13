.PHONY: help build install clean update validate bump-version

BINARY_NAME := scoopbucket
GO := go
GOFLAGS := -v
VERSION ?= $(shell cat .version/latest 2>/dev/null || echo "latest")
VALIDATION_SCRIPT := ./scripts/go-validate.sh
BUILD_DIR := ./bin

help:
	clear;
	@echo "Makefile usage help:"
	@echo "  make help          - Display this message"
	@echo "  make build         - Build binary"
	@echo "  make install       - Build/install the binary in GOBIN"
	@echo "  make clean         - Cleanup build artifacts"
	@echo "  make update        - Update module dependencies"
	@echo "  make valiiate      - Run code and LICENSE validation"
	@echo "  make bump-version  - Create new tag/release automatically via workflow"

build:
	@echo ":: Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	$(GO) build $(GOFLAGS) -ldflags="-X main.version=$(VERSION)" -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/$(BINARY_NAME)

install:
	@echo ":: Installing $(BINARY_NAME)..."
	$(GO) install $(GOFLAGS) -ldflags="-X main.version=$(VERSION)" ./cmd/$(BINARY_NAME)

clean:
	@echo ":: Cleaning build artifacts..."
	rm -rf $(BUILD_DIR)

update:
	@echo ":: Updating module dependencies... "
	$(GO) get -u ./... && $(GO) mod tidy
	
validate:
	@bash -c $(VALIDATION_SCRIPT)

bump-version:
	@read -p "Enter new version (current: $(shell cat .version/latest 2>/dev/null || echo latest)): " NEW_VERSION; \
	VERSION_CLEAN=$$(echo "$$NEW_VERSION" | sed 's/^v//'); \
	FULL_TAG="v$$VERSION_CLEAN"; \
	echo "$$VERSION_CLEAN" > .version/latest; \
	git add .version/; \
	git commit -S -m "chore: bump version to $$FULL_TAG"; \
	echo ":: Version bumped to $$FULL_TAG"; \
	echo ":: Don't forget to: git push && git tag $$FULL_TAG && git push --tags"