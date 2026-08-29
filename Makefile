# Makefile for cisco-ios-xe-wireless-go SDK package
#
# This Makefile provides direct access to specialized build and test scripts
# located in the scripts/ directory for focused development tasks.
#
# For comprehensive help, use: make help
# Every script takes no arguments; see docs/SCRIPT_REFERENCE.md for details.

.PHONY: help deps lint test-unit test-unit-coverage test-integration \
	build pre-commit-install pre-commit-test pre-commit-uninstall

# Default target
help:
	@./scripts/help.sh

# Install development dependencies
deps:
	@./scripts/install_dependencies.sh

# Run linting tools
lint:
	@./scripts/lint.sh

# Run unit tests only
test-unit:
	@./scripts/test_unit.sh

# Run unit tests with coverage analysis
test-unit-coverage:
	@./scripts/test_coverage.sh

# Run integration tests (requires environment variables)
test-integration:
	@./scripts/test_integration.sh

# Verify build compilation
build:
	@go build ./...

# Pre-commit Hook Management
# Install the pre-commit hooks declared in .pre-commit-config.yaml
pre-commit-install:
	@command -v pre-commit >/dev/null 2>&1 || \
		{ echo "✗ pre-commit not found - see https://pre-commit.com/#install"; exit 1; }
	@pre-commit install --allow-missing-config

# Run every hook across the whole tree without committing
pre-commit-test:
	@pre-commit run --all-files

# Uninstall pre-commit hook
pre-commit-uninstall:
	@pre-commit uninstall
