#!/usr/bin/env bash

# Cisco WNC Integration Test - Display Functions
# Provides display functions for integration test environment information

show_integration_test_environment() {
    if ! is_verbose_enabled; then
        return 0
    fi

    printf '%s\n' "Environment Configuration:"
    printf '%s\n' "-------------------------"

    if has_controller_config; then
        success "WNC_CONTROLLER: ${WNC_CONTROLLER}"
    fi

    if has_token_config; then
        success "WNC_ACCESS_TOKEN: ***"
    fi

    printf '\n'
}
