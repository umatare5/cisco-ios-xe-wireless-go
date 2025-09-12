#!/usr/bin/env bash

# Cisco WNC Help - Help Content Functions
# Provides help content display for development scripts

show_help_content() {
    cat << 'EOF'

USAGE:
    make <target>                   # Use Makefile targets (recommended)
    ./scripts/<script>.sh [options] # Use scripts directly

COMMON DEVELOPMENT TARGETS:
    help                Show this help message
    clean               Clean build artifacts and temporary files
    deps                Install development dependencies
    lint                Run code linting tools
    build               Verify build compilation
    test-unit           Run unit tests only
    test-integration    Run integration tests (requires environment)
    test-coverage       Run tests with coverage analysis

YANG MODEL DEVELOPMENT:
    yang-list           List all available YANG models
    yang-model          Get YANG model details (MODEL=model-name)
    yang-statement      Get YANG statement details (MODEL=model-name STATEMENT=statement-name)

ENVIRONMENT VARIABLES:
    WNC_CONTROLLER      Controller hostname/IP for integration tests
    WNC_ACCESS_TOKEN    Base64 encoded credentials for integration tests

EXAMPLES:
    # Basic development workflow
    make deps               # Install dependencies
    make lint               # Check code quality
    make test-unit          # Run unit tests
    make test-unit-coverage # Run unit tests with coverage
    make build              # Verify compilation

    # YANG development
    make yang-list                                    # List models
    make yang-model MODEL=wireless-access-point      # Get model details
    make yang-statement MODEL=wireless-client STATEMENT=active # Get statement details

    # Integration testing (requires environment setup)
    export WNC_CONTROLLER="<controller-host-or-ip>"
    export WNC_ACCESS_TOKEN="<base64-username:password>"
    make test-integration

SCRIPT DETAILS:
    For specific script options and advanced usage:
    ./scripts/<script_name>.sh --help

    Available scripts:
    - clean_artifacts.sh      Clean build artifacts
    - install_dependencies.sh Install Go dependencies
    - lint.sh                Run golangci-lint
    - test_unit.sh           Run unit tests (supports --coverage)
    - test_integration.sh    Run integration tests (supports --coverage)
    - pre_commit_hook.sh     Pre-commit validation hook
    - get_yang_models.sh     List YANG models
    - get_yang_model_details.sh Get model details
    - get_yang_statement_details.sh Get statement details

PROJECT STRUCTURE:
    scripts/                Script directory
    +-- lib/               Shared libraries
    |   +-- bootstrap.sh   Bootstrap library loader
    |   +-- coverage/      Coverage report functions
    |   +-- dependencies/  Dependency management
    |   +-- output/        Output formatting utilities
    |   +-- testing/       Test utilities
    |   +-- utils/         Utility functions
    |   +-- validation/    Git commit validation
EOF
}
