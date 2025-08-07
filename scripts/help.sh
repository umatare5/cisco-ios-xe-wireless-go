#!/usr/bin/env bash
# @meta version 1.0.0
# @meta author "@umatare5"
# @describe Show comprehensive help for Cisco WNC project development

set -euo pipefail

cat << 'EOF'
🔧 Cisco WNC Development Scripts
-------------------------------

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
    test-coverage-html  Generate HTML coverage report

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
    make test-coverage      # Run tests with coverage
    make build              # Verify compilation

    # YANG development
    make yang-list                                    # List models
    make yang-model MODEL=wireless-access-point      # Get model details
    make yang-statement MODEL=wireless-client STATEMENT=active # Get statement details

    # Integration testing (requires environment setup)
    export WNC_CONTROLLER="wnc1.example.internal"
    export WNC_ACCESS_TOKEN="YWRtaW46Y3l0WU43WVh4M2swc3piUnVhb1V1ZUx6"
    make test-integration

SCRIPT DETAILS:
    For specific script options and advanced usage:
    ./scripts/<script_name>.sh --help

    Available scripts:
    - clean_artifacts.sh      Clean build artifacts
    - install_dependencies.sh Install Go dependencies
    - lint.sh                Run golangci-lint
    - test_unit.sh           Run unit tests
    - test_integration.sh    Run integration tests
    - test_coverage.sh       Run coverage tests
    - generate_coverage_report.sh Generate HTML coverage
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
    |   +-- yang/          YANG-specific functions
    +-- *.sh               Entry point scripts

This project uses a modular script architecture with shared libraries
for maintainability and consistency across all development operations.
EOF
