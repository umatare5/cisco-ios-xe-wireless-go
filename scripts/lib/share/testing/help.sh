#!/usr/bin/env bash

# Cisco WNC Testing Operations - Help Functions
# Provides help and documentation functionality for testing operations

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
    local project="$2"
    local timeout="$3"

    is_verbose_enabled || return 0

    printf '%s\n' "$test_type Test Configuration:"
    printf '%s\n' "  Project: ${project:-.}"
    printf '%s\n' "  Timeout: ${timeout:-default}"
    printf '%s\n' "  Short mode: $(is_short_mode_enabled && printf enabled || printf disabled)"
    printf '%s\n' "  Verbose: enabled"
    printf '\n'
}
