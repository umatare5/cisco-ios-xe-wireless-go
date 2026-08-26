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
# Install pre-commit hook to prevent direct commits to main branch
pre-commit-install:
	@ln -sf ../../.githooks/pre-commit .git/hooks/pre-commit
	@echo "✓ Pre-commit hook installed"

# Test pre-commit hook without installing
pre-commit-test:
	@./.githooks/pre-commit

# Uninstall pre-commit hook
pre-commit-uninstall:
	@rm -f .git/hooks/pre-commit
	@echo "✓ Pre-commit hook uninstalled"
