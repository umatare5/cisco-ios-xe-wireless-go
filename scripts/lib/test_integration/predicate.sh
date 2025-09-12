#!/usr/bin/env bash

# Cisco WNC Integration Test - Predicate Functions
# Provides predicate functions for integration test environment validation

has_controller_config() {
    [[ -n "${WNC_CONTROLLER:-}" ]]
}

has_token_config() {
    [[ -n "${WNC_ACCESS_TOKEN:-}" ]]
}

# is_dangerous_test_pattern - Check if test pattern is dangerous and should be excluded
is_dangerous_test_pattern() {
    local test_pattern="$1"

    # Dangerous patterns that cause service interruption - ALWAYS EXCLUDE
    case "$test_pattern" in
        TestWNCReloadIntegration | \
        TestAPReloadIntegration | \
        TestControllerReload* | \
        TestAPReload* | \
        *Controller*Reload*Integration* | \
        *AP*Reload*Integration* | \
        *Reload*Integration*)
            return 0  # Is dangerous
            ;;
        *)
            return 1  # Not dangerous
            ;;
    esac
}

# should_skip_dangerous_test - Always skip dangerous tests regardless of environment
should_skip_dangerous_test() {
    local test_name="$1"

    if is_dangerous_test_pattern "$test_name"; then
        echo "SKIPPING DANGEROUS TEST: $test_name (service interruption risk)"
        return 0  # Should skip
    fi

    return 1  # Should not skip
}
