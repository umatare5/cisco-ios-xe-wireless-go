#!/usr/bin/env bash

# Cisco WNC File Utils - File and temporary management helpers
# Provides functions for creating, cleaning up, and checking files

# Create a temporary file (caller must handle cleanup)
setup_temp_file() {
    local temp_file

    # Create temporary file
    if ! temp_file=$(mktemp); then
        local exit_error
        exit_error="$(get_exit_error)"
        error "Failed to create temporary file"
        exit "$exit_error"
    fi

    printf '%s\n' "$temp_file"
}

# Install cleanup trap (for use by entry scripts)
setup_temp_file_cleanup() {
    local temp_file="$1"
    trap 'cleanup_temp_file "$temp_file"' EXIT INT TERM
}

# Remove temporary file (used by trap)
cleanup_temp_file() {
    local temp_file="$1"

    # Early return if file doesn't exist
    if [[ ! -f "$temp_file" ]]; then
        return 0
    fi

    # Remove temporary file
    rm -f "$temp_file" 2>/dev/null || true
}

# True if file exists and is readable
is_file_readable() {
    local file_path="$1"
    [[ -f "$file_path" && -r "$file_path" ]]
}

# True if file exists and is writable
is_file_writable() {
    local file_path="$1"
    [[ -f "$file_path" && -w "$file_path" ]]
}

# Get file size in bytes
get_file_size() {
    local file_path="$1"

    # Early return if file doesn't exist
    if [[ ! -f "$file_path" ]]; then
    printf '%s\n' "0"
        return 1
    fi

    # Get file size using stat or wc as fallback
    if command -v stat >/dev/null 2>&1; then
        stat -f%z "$file_path" 2>/dev/null || \
        stat -c%s "$file_path" 2>/dev/null || \
        wc -c < "$file_path"
        return
    fi

    wc -c < "$file_path"
}
