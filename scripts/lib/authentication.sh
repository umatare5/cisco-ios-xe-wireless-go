#!/usr/bin/env bash

# Cisco WNC Authentication Library
# Authentication token handling and validation functions

# Source constants
# shellcheck source=./constants.sh
source "$(dirname "${BASH_SOURCE[0]}")/constants.sh"

# Setup authentication token with fallback to environment variable
setup_authentication() {
    local provided_token="$1"
    local env_var_access_token="WNC_ACCESS_TOKEN"
    local token=""

    # Use provided token if available
    if is_not_empty "$provided_token"; then
        token="$provided_token"
    # Fallback to environment variable
    elif is_not_empty "${WNC_ACCESS_TOKEN:-}"; then
        token="$WNC_ACCESS_TOKEN"
    else
        # No authentication token available - show error
        local exit_auth_error
        exit_auth_error="$(get_exit_auth_error)"

        echo "Error: No authentication token provided" >&2
        echo "Either:" >&2
        echo "  1. Use -t option: $0 -t 'your-token'" >&2
        echo "  2. Set environment variable: export \\" >&2
        echo "     $env_var_access_token='your-token'" >&2
        echo "" >&2
        echo "Generate a token using: wnc generate token \\" >&2
        echo "  -u <username> -p <password>" >&2
        exit "$exit_auth_error"
    fi

    # Verify token format before returning
    verify_authentication_token "$token"
    echo "$token"
}

# Verify authentication token format
verify_authentication_token() {
    local token="$1"

    # Early return for valid token
    if is_valid_token_format "$token"; then
        return 0
    fi

    # Invalid token format - show error
    local exit_auth_error
    exit_auth_error="$(get_exit_auth_error)"

    echo "Error: Authentication token appears to be invalid" >&2
    echo "Please verify your token format (should be base64 encoded)" >&2
    echo "Generate a new token using: wnc generate token \\" >&2
    echo "  -u <username> -p <password>" >&2
    exit "$exit_auth_error"
}
