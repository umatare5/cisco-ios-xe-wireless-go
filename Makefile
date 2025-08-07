# Makefile for cisco-ios-xe-wireless-go Go library package
#
# This Makefile provides direct access to specialized build and test scripts
# located in the scripts/ directory for focused development tasks.
#
# For comprehensive help, use: make help
# For specific script options, use: ./scripts/<script_name>.sh --help

.PHONY: help clean deps lint test-unit test-integration test-coverage \
        test-coverage-html build yang-list yang-model yang-statement

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
	@./scripts/lint_code.sh

# Run unit tests only
test-unit:
	@./scripts/test_unit.sh

# Run integration tests (requires environment variables)
test-integration:
	@./scripts/test_integration.sh

# Run tests with coverage analysis
test-coverage:
	@./scripts/test_coverage.sh

# Generate HTML coverage report
test-coverage-html:
	@./scripts/generate_coverage_html.sh

# Verify build compilation
build:
	@go build ./...

# YANG Model Development Tools
# List all available YANG models
yang-list:
	@./scripts/list_yang_models.sh

# Get YANG model details from controller
# (usage: make yang-model MODEL=model-name)
yang-model:
	@./scripts/get_yang_model_details.sh $(MODEL)

# Get YANG statement details from controller
# (usage: make yang-statement MODEL=model-name STATEMENT=statement-name)
yang-statement:
	@./scripts/get_yang_statement_details.sh $(MODEL) $(STATEMENT)
