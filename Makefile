# Makefile for cisco-ios-xe-wireless-go Go library package
#
# This Makefile provides direct access to specialized build and test scripts
# located in the scripts/ directory for focused development tasks.
#
# For comprehensive help, use: make help
# For specific script options, use: ./scripts/<script_name>.sh --help

.PHONY: help clean deps lint test-unit test-unit-coverage test-integration test-integration-coverage \
	test-coverage-report build yang-list yang-model yang-statement \
        pre-commit-install pre-commit-test pre-commit-uninstall

# Default args (can be overridden: make test-integration INTEGRATION_ARGS="")
INTEGRATION_ARGS ?= --check-env-only

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

# Run unit tests only
test-unit:
	@./scripts/test_unit.sh

# Run unit tests with coverage analysis
test-unit-coverage:
	@./scripts/test_unit.sh --coverage

# Run integration tests (requires environment variables)
test-integration:
	@./scripts/test_integration.sh $(INTEGRATION_ARGS)

# Run integration tests with coverage analysis (requires environment variables)
test-integration-coverage:
	@./scripts/test_integration.sh --coverage

# Generate HTML coverage report
test-coverage-report:
	@./scripts/generate_coverage_report.sh

# Verify build compilation
build:
	@go build ./...

# YANG Model Development Tools
# List all available YANG models
yang-list:
	@./scripts/get_yang_models.sh $(ARGS) || { \
		echo "ℹ YANG list skipped (offline or unreachable controller)"; true; }

# Get YANG model details from controller
# (usage: make yang-model MODEL=model-name)
yang-model:
	@[ -n "$(MODEL)" ] || { echo "ℹ YANG model skipped (MODEL not set)"; exit 0; }
	@./scripts/get_yang_model_details.sh $(ARGS) --model $(MODEL) || { \
		echo "ℹ YANG model skipped (offline or unreachable controller)"; true; }

# Get YANG statement details from controller
# (usage: make yang-statement MODEL=model-name STATEMENT=statement-name)
yang-statement:
	@[ -n "$(MODEL)" ] && [ -n "$(STATEMENT)" ] || { \
		echo "ℹ YANG statement skipped (MODEL and/or STATEMENT not set)"; exit 0; }
	@./scripts/get_yang_statement_details.sh $(ARGS) --model $(MODEL) --statement $(STATEMENT) || { \
		echo "ℹ YANG statement skipped (offline or unreachable controller)"; true; }

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
