#!/usr/bin/env bash
set -euo pipefail

# Cisco WNC Testing Operations - Help Functions
# Provides help and documentation functionality for testing operations

# Source common predicates
source "$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)/../common/common.sh"
source "$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)/../common/argument_parsing.sh"

show_test_banner() {
    local test_type="$1"
    if ! is_no_color_enabled; then
        echo -e "\033[33m╔════════════════════════════════════════╗\033[0m"
        echo -e "\033[33m║       Cisco WNC $test_type Tests       ║\033[0m"
        echo -e "\033[33m║       Go Testing Framework            ║\033[0m"
        echo -e "\033[33m╚════════════════════════════════════════╝\033[0m"
    else
        echo "========================================"
        echo "       Cisco WNC $test_type Tests"
        echo "       Go Testing Framework"
        echo "========================================"
    fi
    echo
}

show_unit_test_help() {
    show_test_banner "Unit"
    cat << 'EOF'
USAGE:
    run_unit_tests.sh [OPTIONS]

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
    run_unit_tests.sh

    # Run with coverage
    run_unit_tests.sh --coverage

    # Run in short mode
    run_unit_tests.sh --short

    # Verbose output with custom timeout
    run_unit_tests.sh --verbose --timeout 60s

EOF
}

show_coverage_test_help() {
    show_test_banner "Coverage"
    cat << 'EOF'
USAGE:
    run_coverage_tests.sh [OPTIONS]

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
    run_coverage_tests.sh

    # Generate HTML report
    run_coverage_tests.sh --html

    # Generate and open HTML report
    run_coverage_tests.sh --html --open

EOF
}

show_integration_test_help() {
    show_test_banner "Integration"
    cat << 'EOF'
USAGE:
    run_integration_tests.sh [OPTIONS]

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
    run_integration_tests.sh

    # Run with extended timeout
    run_integration_tests.sh --timeout 300s

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
