# Makefile for cisco-ios-xe-wireless-go Go library package
#
# This Makefile provides direct access to specialized build and test scripts
# located in the scripts/ directory for focused development tasks.
#
# For comprehensive help, use: make help
# For specific script options, use: ./scripts/<script_name>.sh --help

.PHONY: help clean deps lint test-unit test-integration test-coverage \
        test-coverage-html build yang-list yang-model yang-statement \
        fetch-yang-model fetch-yang-statement

# Default target
help:
	@./scripts/show_help.sh

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
	@./scripts/run_unit_tests.sh

# Run integration tests (requires environment variables)
test-integration:
	@./scripts/run_integration_tests.sh

# Run tests with coverage analysis
test-coverage:
	@./scripts/run_coverage_tests.sh

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

# Get YANG model details (usage: make yang-model MODEL=model-name)
yang-model:
	@./scripts/get_yang_model_details.sh \
		$(if $(MODEL),--model $(MODEL),)

# Get YANG statement details (usage: make yang-statement MODEL=model-name
# STATEMENT=statement-name)
yang-statement:
	@./scripts/get_yang_statement_details.sh \
		$(if $(MODEL),--model $(MODEL),) \
		$(if $(STATEMENT),--statement $(STATEMENT),)

# Fetch YANG model details from controller
# (usage: make fetch-yang-model MODEL=model-name)
fetch-yang-model:
	@./scripts/fetch_yang_model_details.sh \
		$(if $(MODEL),--model $(MODEL),)

# Fetch YANG statement details from controller
# (usage: make fetch-yang-statement MODEL=model-name STATEMENT=statement-name)
fetch-yang-statement:
	@./scripts/fetch_yang_statement_details.sh \
		$(if $(MODEL),--model $(MODEL),) \
		$(if $(STATEMENT),--statement $(STATEMENT),)
