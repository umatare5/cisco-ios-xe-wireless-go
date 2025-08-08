# Makefile for cisco-ios-xe-wireless-go Go library package
#
# This Makefile provides direct access to specialized build and test scripts
# located in the scripts/ directory for focused development tasks.
#
# For comprehensive help, use: make help
# For specific script options, use: ./scripts/<script_name>.sh --help

.PHONY: help clean deps lint test-unit test-integration test-coverage \
	test-coverage-report build yang-list yang-model yang-statement \
        pre-commit-install pre-commit-test pre-commit-uninstall

# Default target
help:
	@./scripts/help.sh

# Clean build artifacts
clean:
	@./scripts/clean_artifacts.sh

# Install development dependencies
deps:
	@./scripts/install_dependencies.sh

# Run linting tools
lint:
	@./scripts/lint.sh

# Run unit tests only (lint runs first for code quality)
test-unit: lint
	@./scripts/test_unit.sh

# Run integration tests (requires environment variables; lint runs first)
test-integration: lint
	@./scripts/test_integration.sh

# Run tests with coverage analysis
test-coverage:
	@./scripts/test_coverage.sh

# Generate HTML coverage report
test-coverage-report:
	@./scripts/generate_coverage_report.sh

# Verify build compilation
build:
	@go build ./...

# YANG Model Development Tools
# List all available YANG models
yang-list:
	@./scripts/get_yang_models.sh

# Get YANG model details from controller
# (usage: make yang-model MODEL=model-name)
yang-model:
	@./scripts/get_yang_model_details.sh $(MODEL)

# Get YANG statement details from controller
# (usage: make yang-statement MODEL=model-name STATEMENT=statement-name)
yang-statement:
	@./scripts/get_yang_statement_details.sh $(MODEL) $(STATEMENT)

# Pre-commit Hook Management
# Install pre-commit hook to prevent direct commits to main branch
pre-commit-install:
	@ln -sf ../../scripts/pre_commit_hook.sh .git/hooks/pre-commit
	@echo "✓ Pre-commit hook installed"

# Test pre-commit hook without installing
pre-commit-test:
	@./scripts/pre_commit_hook.sh

# Uninstall pre-commit hook
pre-commit-uninstall:
	@rm -f .git/hooks/pre-commit
	@echo "✓ Pre-commit hook uninstalled"
