#!/usr/bin/env bash

# Cisco WNC YANG - Output functions
# Provides functions to format and display YANG operation messages

# Check if operation has target information
has_target() { [[ -n "${1:-}" ]]; }

# Check if operation was successful
is_operation_successful() { [[ "${1:-1}" -eq 0 ]]; }

# Display RESTCONF request information in verbose mode
display_yang_request_info() {
    local url="$1"
    local method="${2:-GET}"
    local format="${3:-json}"

    is_verbose_enabled || return 0
    printf '\n'
    info "RESTCONF Request Details:"
    printf '%s\n' "  Method: $method"
    printf '%s\n' "  URL: $url"
    printf '%s\n' "  Format: $format"
    printf '\n'
}

# Display RESTCONF response information in verbose mode
display_yang_response_info() {
    local status_code="$1"; local response_size="${2:-unknown}"
    is_verbose_enabled || return 0
    printf '\n'
    info "RESTCONF Response:"
    printf '%s\n' "  Status: $status_code"
    printf '%s\n' "  Size: $response_size bytes"
    printf '\n'
}

# Print raw content without formatting
_print_raw() { printf '%s' "$1" || true; }

# Check if content appears to be valid JSON
_looks_like_json() {
    local content="$1"
    local first_char
    first_char=$(printf '%s' "$content" | sed -e 's/^[[:space:]]*//' | head -c 1 || true)
    [[ "$first_char" == "{" || "$first_char" == "[" ]]
}

# Pretty print JSON from stdin using jq if available
_pretty_json_stdin() {
    if command -v jq >/dev/null 2>&1; then
        jq '.' || true
        return
    fi

    cat || true
}

# Pretty print JSON from file using jq if available
_pretty_json_file() {
    local file_path="$1"

    if command -v jq >/dev/null 2>&1; then
        jq '.' "$file_path" || cat "$file_path"
        return
    fi

    cat "$file_path" || true
}

# Format YANG output content with optional JSON formatting
format_yang_output() {
    local content="$1"
    local format="${2:-json}"
    local raw="${3:-false}"

    if [[ "$raw" == "true" ]]; then
        _print_raw "$content"
        return 0
    fi

    if [[ "$format" == json ]]; then
        if _looks_like_json "$content"; then
            printf '%s' "$content" | _pretty_json_stdin
            return
        fi

        _print_raw "$content"
        return
    fi

    _print_raw "$content"
}

# Format YANG output from file to avoid large argument expansion issues
format_yang_output_file() {
    local file_path="$1"
    local output_format="${2:-json}"
    local raw_output="${3:-false}"

    if [[ "$raw_output" == "true" ]]; then
        cat "$file_path" || true
        return 0
    fi

    if [[ "$output_format" != "json" ]]; then
        cat "$file_path" || true
        return 0
    fi

    local first_char
    first_char=$(sed -e 's/^[[:space:]]*//' "$file_path" | head -c 1 || true)
    if [[ "$first_char" == "{" || "$first_char" == "[" ]]; then
        _pretty_json_file "$file_path"
        return
    fi

    cat "$file_path"
}

# Display results of YANG operation with success/failure status
display_yang_operation_results() {
    local exit_code="$1"
    local operation="$2"
    local target="${3:-}"

    printf '\n'

    # Handle failure case with early return
    if ! is_operation_successful "$exit_code"; then
        error "$operation failed"
        info "Check the output above for details"
        info "Verify controller connectivity and credentials"
        return 0
    fi

    # Handle success case
    success "$operation completed successfully" || true
    if has_target "$target"; then
        info "Target: $target"
    fi
    return 0
}
