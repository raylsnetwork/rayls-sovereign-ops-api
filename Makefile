.PHONY: all build clean run test docker-dev docker-dev-down docker-dev-logs

# Variables
BINARY_NAME=rayls-ops-api

GOLANGCI_LINT_VERSION=v2.7.1

SRC_DIR=cmd/api
MAIN_FILE=$(SRC_DIR)/main.go
BUILD_DIR=build

# Targets
all: build

build: swagger
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_FILE)

swagger:
	swag init --parseDependency -q -g ./cmd/api/main.go -o ./cmd/api/docs

test:
	go test ./...

clean:
	rm -rf $(BUILD_DIR)

run: build
	./$(BUILD_DIR)/$(BINARY_NAME)

docker-dev:
	docker compose -f docker-compose.dev.yml up --build

docker-dev-down:
	docker compose -f docker-compose.dev.yml down

docker-dev-logs:
	docker compose -f docker-compose.dev.yml logs -f ops-api

# Install all Go linters and code-quality tools used by the project
install-linters:
	@echo "Installing golangci-lint $(GOLANGCI_LINT_VERSION)..."
	@GOBIN=$${GOBIN:-$$(go env GOPATH)/bin}; \
	mkdir -p "$$GOBIN"; \
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh \
		| sh -s -- -b "$$GOBIN" $(GOLANGCI_LINT_VERSION)
	@echo ""
	@echo "  ✓ golangci-lint installed successfully!"

# Install and activate pre-commit so every contributor runs the same checks locally
install-precommit:
	@command -v pre-commit >/dev/null 2>&1 || { \
		echo "Installing pre-commit..."; \
		if command -v pipx >/dev/null 2>&1; then \
			pipx install pre-commit; \
		elif command -v brew >/dev/null 2>&1; then \
			brew install pre-commit; \
		else \
			pip install --user pre-commit; \
		fi \
	}
	pre-commit install
	@echo ""
	@echo "  ✓ pre-commit installed successfully!"

# Full environment setup: installs linters and registers the project's Git hooks
setup-linters: install-linters install-precommit
	@echo ""
	@echo "  ✓ Setup ready to Go! :)"

# Run all linters and formatters checks (usage: make lint [target])
# Examples: make lint . | make lint ./cmd/api | make lint ./cmd/api/core/ports.go
lint:
	@if ! command -v golangci-lint >/dev/null 2>&1; then \
		echo "Error: golangci-lint is not installed"; \
		echo "Run: make install-linters"; \
		exit 1; \
	fi
	@TARGET="$(or $(filter-out $@,$(MAKECMDGOALS)),.)"; \
	if [ -d "$$TARGET" ]; then \
		LINT_TARGET="$$TARGET/..."; \
	else \
		LINT_TARGET="$$TARGET"; \
	fi; \
	if [ "$$TARGET" = "." ]; then \
		LINT_TARGET="./..."; \
		echo "Running golangci-lint on entire codebase..."; \
	else \
		echo "Running golangci-lint on $$TARGET..."; \
	fi; \
	echo "Running formatters..."; \
	golangci-lint fmt $$LINT_TARGET; \
	echo "Running linters..."; \
	golangci-lint run $$LINT_TARGET
