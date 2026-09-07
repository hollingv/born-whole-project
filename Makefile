.PHONY: help init build test clean version-preview release-preview release
.SILENT:

APP_NAME    = bwctl
TARGET_DIR  = dist
GO_VERSION  = 1.26.1
COG_VERSION = 7.0.0
DRY_RUN    ?= true

# Resolve Go binary: prefer whatever is in PATH, fall back to the standard install location.
GO_BIN := $(shell command -v go 2>/dev/null || echo /usr/local/go/bin/go)

APP_TAG := $(shell \
	if test -n "$$(git status --short)"; then \
		echo "local-dirty"; \
	else \
		TAG=$$(git describe --tags --match "v*" 2>/dev/null || git rev-parse --short=7 HEAD); \
		BRANCH=$$(git rev-parse --abbrev-ref HEAD); \
		if [ "$$BRANCH" = "main" ]; then \
			echo "$$TAG"; \
		else \
			echo "$$TAG-$$BRANCH"; \
		fi; \
	fi)

help:
	@echo ""
	@echo "Bodily Integrity Commons Website"
	@echo "(version: $(APP_TAG))"
	@echo ""
	@grep -E '^[a-zA-Z0-9-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

init: ## Install Go if not already present (run once after cloning)
	# --- Go ---
	@if [ ! -f "$(GO_BIN)" ]; then \
		echo "[ INFO ] Installing Go $(GO_VERSION)..."; \
		curl -sL "https://go.dev/dl/go$(GO_VERSION).linux-amd64.tar.gz" | sudo tar -xz -C /usr/local; \
		echo "[ INFO ] Go $(GO_VERSION) installed to /usr/local/go"; \
	else \
		echo "[ INFO ] Go already installed: $$($(GO_BIN) version)"; \
	fi
	@if ! command -v go > /dev/null 2>&1; then \
		echo ""; \
		echo "[ INFO ] Add Go to your PATH by adding this line to ~/.bashrc:"; \
		echo "[ INFO ]   export PATH=\$$PATH:/usr/local/go/bin"; \
		echo ""; \
	fi
	# --- Go dependencies ---
	@echo "[ INFO ] Fetching Go dependencies..."
	$(GO_BIN) mod tidy
	# --- cog ---
	@echo "[ INFO ] Checking for cog..."
	@if ! command -v cog > /dev/null 2>&1; then \
		echo "[ INFO ] Installing cog $(COG_VERSION)..."; \
		curl -sL "https://github.com/cocogitto/cocogitto/releases/download/$(COG_VERSION)/cocogitto-$(COG_VERSION)-x86_64-unknown-linux-musl.tar.gz" | tar -xz -C /tmp --strip-components=1; \
		mkdir -p ~/.local/bin; \
		mv /tmp/cog ~/.local/bin/cog; \
		echo "[ INFO ] cog installed: $$(cog --version)"; \
	else \
		echo "[ INFO ] cog already installed: $$(cog --version)"; \
	fi
	# --- git hooks ---
	@echo "[ INFO ] Installing git hooks..."
	cp scripts/hooks/commit-msg.sh .git/hooks/commit-msg
	chmod +x .git/hooks/commit-msg
	cp scripts/hooks/pre-push.sh .git/hooks/pre-push
	chmod +x .git/hooks/pre-push
	@echo "[ INFO ] git hooks installed"

build: ## Build the binary and site
	@echo "[ INFO ] Tidying Go modules..."
	$(GO_BIN) mod tidy
	@echo "Building $(APP_NAME) version $(APP_TAG)..."
	mkdir -p $(TARGET_DIR)
	CGO_ENABLED=0 $(GO_BIN) build -ldflags '-X main.version=$(APP_TAG)' -o $(TARGET_DIR)/$(APP_NAME) ./src/cmd/$(APP_NAME)
	cp -r site/. $(TARGET_DIR)/
	./$(TARGET_DIR)/$(APP_NAME) site --version $(APP_TAG)
	@echo "Built: $(APP_NAME) at version $(APP_TAG)"

test: build test-unit test-integ-local ## Build and run all tests

test-unit: build ## Run the go unit tests
	$(GO_BIN) test ./src/cmd/$(APP_NAME)/... ./src/server/...

test-env-vars: ## Chaeck the status of the environment variables used for testing
	./$(TARGET_DIR)/$(APP_NAME) status --set-exit-code=true

test-integ-local: build test-env-vars ## Start server, run integration tests, stop server
	@go run ./src/server > /dev/null 2>&1 & \
	SERVER_PID=$$!; \
	trap "kill $$SERVER_PID 2>/dev/null; wait $$SERVER_PID 2>/dev/null; kill $$(lsof -t -i:8080) 2>/dev/null" EXIT INT TERM; \
	sleep 2; \
	$(GO_BIN) test ./src/test-integ/... -url http://localhost:8080; \
	TEST_EXIT=$$?; \
	kill $$SERVER_PID 2>/dev/null; \
	wait $$SERVER_PID 2>/dev/null; \
	kill $$(lsof -t -i:8080) 2>/dev/null; \
	if [ $$TEST_EXIT -eq 0 ]; then echo "Tests PASSED"; else echo "Tests FAILED"; fi; \
	exit $$TEST_EXIT

version-preview: ## Show the next semantic version based on commits since last tag
	@echo "Current version: $(APP_TAG)"
	@echo "Next version would be:"
	@cog bump --dry-run --auto

release-preview: ## Dry run showing next version and full changelog without making any changes
	@echo "[INFO] Changelog"
	@cog changelog v0.0.4..HEAD
	@echo "[INFO] Version"
	@echo "[INFO] Current: $(APP_TAG)"
	@echo "[INFO] Next:"
	@cog bump --dry-run --auto
	@echo ""

release: build release-preview ## Perform a full release. Set DRY_RUN=false to perform the actual release.
	@if [ "$(DRY_RUN)" = "true" ]; then \
		echo "[INFO] Dry run complete";\
		echo "[INFO] Run 'make release DRY_RUN=false' to perform the actual release.";\
	else \
		echo "[INFO] Bumping version";\
		cog bump --auto;\
		echo "[INFO] Pushing commits and tags";\
		git push origin HEAD;\
		NEW_TAG=$$(git describe --tags --abbrev=0 --match "v*"); \
		git push origin $$NEW_TAG;\
		echo "";\
		echo "[INFO] Release completed OK. CI will build and test.";\
	fi

whitepaper: ## Generate whitepaper
	@echo "[ INFO ] Building whitepaper..."
	docker run --rm -v $(PWD):/data pandoc/latex README-whitepaper.md -o zx.pdf
	@echo "[ INFO ] Generated zx.pdf"

app-tag: ## Print the current APP_TAG
	@echo "$(APP_TAG)"

clean: ## Remove built binaries
	@echo "Cleaning up..."
	rm -rf $(TARGET_DIR)/
