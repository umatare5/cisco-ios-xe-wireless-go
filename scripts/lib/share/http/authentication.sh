#!/usr/bin/env bash

# Cisco WNC Auth - Token handling and validation
# Provides functions to handle authentication tokens, including validation and setup

# Get authentication token with fallback hierarchy
_get_authentication_token() {
    local provided_token="$1"

    # Use provided token if available
    if is_not_empty "$provided_token"; then
        printf '%s\n' "$provided_token"
        return 0
    fi

    # Fallback to environment variable
    if is_not_empty "${WNC_ACCESS_TOKEN:-}"; then
        printf '%s\n' "$WNC_ACCESS_TOKEN"
        return 0
    fi

    # No token available
    return 1
}

# Show authentication error and exit
_show_authentication_error() {
    local exit_auth_error
    exit_auth_error="$(get_exit_auth_error)"

    error "No authentication token provided"
    printf '%s\n' "Either:" >&2
    printf '%s\n' "  1. Use -t option: $0 -t 'your-token'" >&2
    printf '%s\n' "  2. Set environment variable: export WNC_ACCESS_TOKEN='your-token'" >&2
    printf '\n' >&2
    printf '%s\n' "Generate a token and export WNC_ACCESS_TOKEN (Bearer token)" >&2
    exit "$exit_auth_error"
}

# Setup authentication token with fallback to environment variable
setup_authentication() {
    local provided_token="$1"
    local token

    # Try to get token with fallback hierarchy
    if ! token=$(_get_authentication_token "$provided_token"); then
        _show_authentication_error
    fi

    # Verify token format before returning
    verify_authentication_token "$token"
    printf '%s\n' "$token"
}

# Show token format error and exit
_show_token_format_error() {
    local exit_auth_error
    exit_auth_error="$(get_exit_auth_error)"

    error "Authentication token appears to be invalid"
    printf '%s\n' "Please verify your token format (should be base64 encoded)" >&2
    printf '%s\n' "Generate a new Bearer token and export WNC_ACCESS_TOKEN" >&2
    exit "$exit_auth_error"
}

# Verify authentication token format
verify_authentication_token() {
    local token="$1"

    # Early return for valid token
    if is_valid_token_format "$token"; then
        return 0
    fi

    # Invalid token format - show error and exit
    _show_token_format_error
}
