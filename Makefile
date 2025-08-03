# Makefile for cisco-ios-xe-wireless-go Go library package

.PHONY: help clean deps lint test-unit test-integration test-coverage test-coverage-html

# Default target
help:
	@echo "Available targets:"
	@echo "  clean            - Clean build artifacts"
	@echo "  deps             - Install development dependencies (including gotestsum)"
	@echo "  lint             - Run linting tools"
	@echo "  test-unit        - Run unit tests only"
	@echo "  test-integration - Run integration tests (requires environment variables)"
	@echo "  test-coverage    - Run tests with coverage analysis"
	@echo "  test-coverage-html - Generate HTML coverage report"

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -f coverage.out
	rm -rf ./tmp
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
	@if ! command -v gotestsum >/dev/null 2>&1; then \
		echo "Installing gotestsum..."; \
		go install gotest.tools/gotestsum@latest; \
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

.PHONY: test-unit
test-unit:
	@echo "Running unit tests only (no environment variables required)..."
	@if command -v gotestsum >/dev/null 2>&1; then \
		WNC_CONTROLLER="" WNC_ACCESS_TOKEN="" gotestsum --format testname -- -race ./...; \
	else \
		echo "gotestsum not found, running go test with verbose output..."; \
		WNC_CONTROLLER="" WNC_ACCESS_TOKEN="" go test -v -race ./...; \
	fi

.PHONY: test-integration
test-integration:
	@echo "Running integration tests (requires WNC_CONTROLLER and WNC_ACCESS_TOKEN)..."
	@if [ -z "$$WNC_CONTROLLER" ] || [ -z "$$WNC_ACCESS_TOKEN" ]; then \
		echo "Error: WNC_CONTROLLER and WNC_ACCESS_TOKEN environment variables must be set"; \
		exit 1; \
	fi
	@if command -v gotestsum >/dev/null 2>&1; then \
		gotestsum --format testname -- -race ./...; \
	else \
		echo "gotestsum not found, running go test with verbose output..."; \
		go test -v -race ./...; \
	fi

.PHONY: test-coverage
test-coverage:
	@echo "Running tests with coverage..."
	@mkdir -p ./tmp
	@if command -v gotestsum >/dev/null 2>&1; then \
		WNC_CONTROLLER="" WNC_ACCESS_TOKEN="" gotestsum --format testname -- -race -coverprofile=./tmp/coverage.out ./...; \
	else \
		echo "gotestsum not found, running go test with verbose output..."; \
		WNC_CONTROLLER="" WNC_ACCESS_TOKEN="" go test -v -race -coverprofile=./tmp/coverage.out ./...; \
	fi
	@if [ -f ./tmp/coverage.out ]; then \
		echo "Coverage report generated at ./tmp/coverage.out"; \
		go tool cover -func=./tmp/coverage.out | tail -1; \
	fi

.PHONY: test-coverage-html
test-coverage-html: test-coverage
	@echo "Generating HTML coverage report..."
	@mkdir -p ./tmp
	@if [ -f ./tmp/coverage.out ]; then \
		go tool cover -html=./tmp/coverage.out -o ./tmp/coverage.html; \
		echo "HTML coverage report generated at ./tmp/coverage.html"; \
	else \
		echo "No coverage file found. Run 'make test-coverage' first."; \
	fi
