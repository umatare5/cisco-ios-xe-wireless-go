#!/usr/bin/env bash

# Cisco WNC File Utils Library
# File handling and temporary file management functions

# Source constants
# shellcheck source=./constants.sh
source "$(dirname "${BASH_SOURCE[0]}")/constants.sh"

# Setup temporary file with automatic cleanup
setup_temp_file_with_cleanup() {
    local temp_file

    # Create temporary file
    if ! temp_file=$(mktemp); then
        local exit_error
        exit_error="$(get_exit_error)"
        echo "Error: Failed to create temporary file" >&2
        exit "$exit_error"
    fi

    # Setup cleanup trap
    trap 'cleanup_temp_file "$temp_file"' EXIT INT TERM

    echo "$temp_file"
}

# Cleanup temporary file
cleanup_temp_file() {
    local temp_file="$1"

    # Early return if file doesn't exist
    if [[ ! -f "$temp_file" ]]; then
        return 0
    fi

    # Remove temporary file
    rm -f "$temp_file" 2>/dev/null || true
}

# Check if file exists and is readable
is_file_readable() {
    local file_path="$1"
    [[ -f "$file_path" && -r "$file_path" ]]
}

# Check if file exists and is writable
is_file_writable() {
    local file_path="$1"
    [[ -f "$file_path" && -w "$file_path" ]]
}

# Get file size in bytes
get_file_size() {
    local file_path="$1"

    # Early return if file doesn't exist
    if [[ ! -f "$file_path" ]]; then
        echo "0"
        return 1
    fi

    # Get file size using stat or wc as fallback
    if command -v stat >/dev/null 2>&1; then
        stat -f%z "$file_path" 2>/dev/null || \
        stat -c%s "$file_path" 2>/dev/null || \
        wc -c < "$file_path"
    else
        wc -c < "$file_path"
    fi
}
