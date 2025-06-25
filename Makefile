# Makefile for wnc Go library package

.PHONY: help check clean deps lint test

# Default target
help:
	@echo "Available targets:"
	@echo "  check      - Run tests and linting"
	@echo "  clean      - Clean build artifacts"
	@echo "  deps       - Install development dependencies"
	@echo "  lint       - Run linting tools"
	@echo "  test       - Run all tests"

# Run comprehensive checks
check: test lint
	@echo "All checks completed successfully!"

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -f coverage.out
	cd ../.. && go clean -cache -testcache

# Install development dependencies
deps:
	@echo "Installing development dependencies..."
	@if ! command -v golangci-lint >/dev/null 2>&1; then \
		echo "Installing golangci-lint..."; \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
	fi
	@if ! command -v goreleaser >/dev/null 2>&1; then \
		echo "Installing goreleaser..."; \
		go install github.com/goreleaser/goreleaser@latest; \
	fi
	@echo "Development dependencies installed!"

# Run linting (if tools are available)
lint:
	@echo "Running linting..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		cd ../.. && golangci-lint run ./...; \
	else \
		echo "golangci-lint not found, running go vet instead..."; \
		cd ../.. && go vet ./...; \
	fi

# Run tests
test: test-unit test-integration
	@echo "Running tests..."

.PHONY: test-unit
test-unit:
	@echo "Running unit tests only (no environment variables required)..."
	WNC_CONTROLLER="" WNC_ACCESS_TOKEN="" go test -v -race ./...

.PHONY: test-integration
test-integration:
	@echo "Running integration tests (requires WNC_CONTROLLER and WNC_ACCESS_TOKEN)..."
	@if [ -z "$$WNC_CONTROLLER" ] || [ -z "$$WNC_ACCESS_TOKEN" ]; then \
		echo "Error: WNC_CONTROLLER and WNC_ACCESS_TOKEN environment variables must be set"; \
		exit 1; \
	fi
	go test -v -race ./...
