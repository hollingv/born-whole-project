.PHONY: help init build test clean version-preview release-preview release docker-build test-links test-links-get-involved test-unit test-env-vars test-integ-local
.SILENT:

APP_NAME       = bwctl
SITE_NAME      := $(shell grep 'const siteName' src/cmd/bwctl/config.go | sed 's/.*"\(.*\)".*/\1/')
TARGET_DIR     = dist
GO_VERSION     = 1.26.1
COG_VERSION    = 7.0.0
LYCHEE_VERSION = 0.15.1
DOCKER_IMAGE   = born-whole-project
DRY_RUN       ?= true
DEBUG         ?= false

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
	@echo "$(SITE_NAME)"
	@echo "(version: $(APP_TAG))"
	@echo ""
	@awk 'BEGIN {FS = ":.*?## "} \
		/^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } \
		/^[a-zA-Z0-9-]+:.*?## / { printf "  \033[36m%-28s\033[0m %s\n", $$1, $$2 }' \
		$(MAKEFILE_LIST)
	@echo ""

##@ Basic

init: ## Install any tools and dependencies (run once after cloning)
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
	# --- Docker ---
	@echo "[ INFO ] Checking for Docker..."
	@if ! command -v docker > /dev/null 2>&1; then \
		echo "[ INFO ] Installing Docker from official repository..."; \
		sudo apt-get update -q; \
		sudo apt-get install -y ca-certificates curl; \
		sudo install -m 0755 -d /etc/apt/keyrings; \
		sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc; \
		sudo chmod a+r /etc/apt/keyrings/docker.asc; \
		echo "deb [arch=$$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu $$(. /etc/os-release && echo $${UBUNTU_CODENAME:-$$VERSION_CODENAME}) stable" \
			| sudo tee /etc/apt/sources.list.d/docker.list > /dev/null; \
		sudo apt-get update -q; \
		sudo apt-get install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin; \
		echo "[ INFO ] Docker installed: $$(docker --version)"; \
	else \
		echo "[ INFO ] Docker already installed: $$(docker --version)"; \
	fi
	@echo "[ INFO ] Ensuring Docker daemon is enabled and running..."
	@sudo systemctl enable docker > /dev/null 2>&1
	@sudo systemctl start docker > /dev/null 2>&1
	@echo "[ INFO ] Docker daemon status: $$(sudo systemctl is-active docker)"
	@if ! groups | grep -q docker; then \
		echo "[ INFO ] Adding $$USER to docker group..."; \
		sudo usermod -aG docker $$USER; \
		echo "[ INFO ] Docker group configured — run 'newgrp docker' or log out and back in for it to take effect"; \
	else \
		echo "[ INFO ] User already in docker group"; \
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

clean: ## Remove built binaries
	@echo "Cleaning up..."
	rm -rf $(TARGET_DIR)/

##@ Testing

test-unit: build ## Run the go unit tests
	$(GO_BIN) test ./src/cmd/$(APP_NAME)/... ./src/server/...

test-env-vars: ## Check the status of the environment variables used for testing
	./$(TARGET_DIR)/$(APP_NAME) status --set-exit-code=true

test-integ-local: build test-env-vars test-links ## Start server, run integration tests, stop server
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

test-links-get-involved: build docker-build ## Check external links on the Get Involved page using lychee
	docker run --rm \
		-v $(PWD)/$(TARGET_DIR):/dist:ro \
		$(DOCKER_IMAGE) \
		lychee \
			--exclude 'https://fonts.googleapis.com' \
			--exclude 'https://fonts.gstatic.com' \
			'/dist/get-involved.html'

test-links: build docker-build test-links-get-involved ## Check all links in the built site using lychee
	docker run --rm \
		-v $(PWD)/$(TARGET_DIR):/dist:ro \
		$(DOCKER_IMAGE) \
		lychee --offline '/dist/**/*.html'

##@ Advanced

docker-build: ## Build the development Docker image
	docker build --build-arg LYCHEE_VERSION=$(LYCHEE_VERSION) -t $(DOCKER_IMAGE) .

app-tag: ## Print the current APP_TAG
	@echo "$(APP_TAG)"

whitepaper: ## Generate whitepaper PDF from README-whitepaper.md
	@echo "[ INFO ] Building whitepaper..."
	docker run --rm -v $(PWD):/data pandoc/latex README-whitepaper.md -o zx.pdf
	@echo "[ INFO ] Generated zx.pdf"

##@ Release

version-preview: ## Show the next semantic version based on commits since last tag
	@printf "Current version:       %s\n" "$(APP_TAG)"
	@printf "Next version would be: %s\n" "$$(cog bump --dry-run --auto)"

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
