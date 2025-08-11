#!/usr/bin/env bash

# Cisco WNC Dependencies Library
# Provides functions to check for required dependencies before script execution

# Check if curl is available
check_curl_dependency() {
    # Early return if curl is available
    if is_curl_available; then
        return 0
    fi

    # curl not found - show error and exit
    local exit_error
    exit_error="$(get_exit_error)"

    error "curl is required but not installed."
    printf '%s\n' "Please install curl and try again." >&2
    printf '\n' >&2
    printf '%s\n' "Installation instructions:" >&2
    printf '%s\n' "  macOS:    brew install curl" >&2
    printf '%s\n' "  Ubuntu:   sudo apt-get install curl" >&2
    printf '%s\n' "  CentOS:   sudo yum install curl" >&2
    exit "$exit_error"
}

# Check if mktemp is available
check_mktemp_dependency() {
    # Early return if mktemp is available
    if command -v mktemp >/dev/null 2>&1; then
        return 0
    fi

    # mktemp not found - show error and exit
    local exit_error
    exit_error="$(get_exit_error)"

    error "mktemp is required but not available."
    printf '%s\n' "Please install coreutils and try again." >&2
    exit "$exit_error"
}

# Check if jq is available (optional but recommended)
check_jq_dependency() {
    # Early return if jq is available
    if is_jq_available; then
        return 0
    fi

    # jq not found - show warning and continue
    printf '%s\n' "Warning: jq is not installed. JSON formatting will be limited." >&2
    printf '%s\n' "For better JSON output, install jq:" >&2
    printf '%s\n' "  macOS:    brew install jq" >&2
    printf '%s\n' "  Ubuntu:   sudo apt-get install jq" >&2
    printf '%s\n' "  CentOS:   sudo yum install jq" >&2
    printf '\n' >&2

    # Continue without jq (non-fatal)
    return 0
}

# Check if grep is available
check_grep_dependency() {
    # Early return if grep is available
    if command -v grep >/dev/null 2>&1; then
        return 0
    fi

    # grep not found - show error and exit
    local exit_error
    exit_error="$(get_exit_error)"

    error "grep is required but not available."
    printf '%s\n' "Please install grep and try again." >&2
    exit "$exit_error"
}

# Check dependencies based on output format requirements
check_dependencies_for_format() {
    local output_format="$1"

    # Always check core dependencies
    check_curl_dependency
    check_mktemp_dependency
    check_grep_dependency

    # Check jq dependency only if required for the output format
    if is_jq_required "$output_format"; then
        check_jq_dependency
    fi
}

# Check all dependencies (for scripts that always need all)
check_all_dependencies() {
    check_curl_dependency
    check_mktemp_dependency
    check_grep_dependency
    check_jq_dependency
}
