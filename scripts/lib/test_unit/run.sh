#!/usr/bin/env bash

# Cisco WNC Unit Test Runner
# Provides functions to run unit tests with prepared arguments

# Run unit tests with prepared args
run_unit_test_operation() {
    local project_root="$1"
    local timeout="$2"

    # Display test information and configuration
    show_test_banner "Unit"
    show_test_verbose_info "Unit" "$project_root" "$timeout"

    # Validate test environment before proceeding
    if ! validate_test_environment "$project_root" "unit"; then
        return 1
    fi

    # Prepare test arguments and execute unit tests
    local test_args
    prepare_test_arguments "unit" "$timeout" test_args
    progress "Starting unit tests..."

    # Execute tests and capture exit code
    local exit_code=0
    execute_test_command "$project_root" "unit" "$test_args" || exit_code=$?

    # Display final test results summary
    display_test_summary "Unit" "$exit_code" "${TEST_DURATION:-unknown}"
    return "$exit_code"
}
