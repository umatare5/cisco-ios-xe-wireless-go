#!/usr/bin/env bash

# Cisco WNC Integration Test - Core Operations
# Provides core functionality for integration test operations

run_integration_test_operation() {
    local project_root="$1"
    local timeout="$2"

        # Display test information and configuration
    show_test_banner "Integration"
    show_test_verbose_info "Integration" "$project_root" "$timeout"
    show_integration_test_environment

    # Validate test environment before proceeding
    if ! validate_test_environment "$project_root" "integration"; then
        return 1
    fi

    # Prepare test arguments and execute integration tests
    local test_args
    prepare_test_arguments "integration" "$timeout" test_args
    progress "Starting integration tests..."

    # Execute tests and capture exit code
    local exit_code=0
    execute_test_command "$project_root" "integration" "$test_args" || exit_code=$?

    # Display final test results summary
    display_test_summary "Integration" "$exit_code" "${TEST_DURATION:-unknown}"
    return "$exit_code"
}
