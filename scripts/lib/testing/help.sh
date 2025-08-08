#!/usr/bin/env bash

# Cisco WNC Testing Operations - Help Functions
# Provides help and documentation functionality for testing operations

set -euo pipefail

# Source bootstrap library
LIB_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
source "${LIB_ROOT}/bootstrap.sh"
init_wnc_basic

show_test_banner() {
    local test_type="$1"
    wnc_banner_tests "$test_type"
}

show_unit_test_help() {
    show_test_banner "Unit"
    cat << 'EOF'
USAGE:
    test_unit.sh [OPTIONS]

DESCRIPTION:
    Runs Go unit tests for the Cisco IOS-XE wireless library.
    Executes tests with proper timeout and output formatting.

OPTIONS:
    -p, --project <DIR>      Project root directory [default: .]
    -v, --verbose            Enable verbose test output
    -s, --short              Run tests in short mode (skip long-running tests)
    -c, --coverage           Generate coverage data
    -t, --timeout <DURATION> Test timeout duration [default: 30s]
    -h, --help               Show this help message

EXAMPLES:
    # Run basic unit tests
    test_unit.sh

    # Run with coverage
    test_unit.sh --coverage

    # Run in short mode
    test_unit.sh --short

    # Verbose output with custom timeout
    test_unit.sh --verbose --timeout 60s

EOF
}

show_coverage_test_help() {
    show_test_banner "Coverage"
    cat << 'EOF'
USAGE:
    test_coverage.sh [OPTIONS]

DESCRIPTION:
    Runs Go tests with coverage analysis and generates coverage reports.
    Produces coverage data files and optional HTML reports.

OPTIONS:
    -p, --project <DIR>      Project root directory [default: .]
    -o, --output <FILE>      Coverage output file [default: ./tmp/coverage.out]
    -v, --verbose            Enable verbose output
    -s, --short              Run tests in short mode
    -t, --timeout <DURATION> Test timeout duration [default: 60s]
        --html               Generate HTML coverage report
        --open               Open HTML report in browser (requires --html)
    -h, --help               Show this help message

EXAMPLES:
    # Basic coverage test
    test_coverage.sh

    # Generate HTML report
    test_coverage.sh --html

    # Generate and open HTML report
    test_coverage.sh --html --open

EOF
}

show_integration_test_help() {
    show_test_banner "Integration"
    cat << 'EOF'
USAGE:
    test_integration.sh [OPTIONS]

DESCRIPTION:
    Runs integration tests against live WNC controllers.
    Requires WNC environment variables for authentication.

OPTIONS:
    -p, --project <DIR>      Project root directory [default: .]
    -v, --verbose            Enable verbose output
    -t, --timeout <DURATION> Test timeout duration [default: 120s]
        --skip-env-check     Skip environment variable validation
    -h, --help               Show this help message

ENVIRONMENT:
    WNC_CONTROLLER           Primary controller hostname
    WNC_ACCESS_TOKEN         Authentication token

EXAMPLES:
    # Run integration tests
    test_integration.sh

    # Run with extended timeout
    test_integration.sh --timeout 300s

EOF
}

show_test_verbose_info() {
    local test_type="$1"
    is_verbose_enabled || return 0

    echo "$test_type Test Configuration:"
    echo "  Project: ${argc_project:-.}"
    echo "  Timeout: ${argc_timeout:-default}"
    echo "  Short mode: $(is_short_mode_enabled && echo "enabled" || echo "disabled")"
    echo "  Verbose: enabled"
    echo
}
