#!/usr/bin/env bash

# Cisco WNC Integration Test - Predicate Functions
# Provides predicate functions for integration test environment validation

has_controller_config() {
    [[ -n "${WNC_CONTROLLER:-}" ]]
}

has_token_config() {
    [[ -n "${WNC_ACCESS_TOKEN:-}" ]]
}
