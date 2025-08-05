#!/usr/bin/env bash

# Cisco WNC YANG Operations - Output Functions
# Handles output formatting and display for YANG operations

format_yang_error() {
    local message="$1"
    if ! is_no_color_enabled; then
        echo -e "\033[31m✗ Error:\033[0m $message" >&2
    else
        echo "✗ Error: $message" >&2
    fi
}

format_yang_success() {
    local message="$1"
    if ! is_no_color_enabled; then
        echo -e "\033[32m✓\033[0m $message"
    else
        echo "✓ $message"
    fi
}

format_yang_warning() {
    local message="$1"
    if ! is_no_color_enabled; then
        echo -e "\033[33m⚠ Warning:\033[0m $message" >&2
    else
        echo "⚠ Warning: $message" >&2
    fi
}

format_yang_info() {
    local message="$1"
    if ! is_no_color_enabled; then
        echo -e "\033[36mℹ Info:\033[0m $message"
    else
        echo "ℹ Info: $message"
    fi
}

show_yang_progress() {
    local task="$1"

    if ! is_no_color_enabled; then
        echo -e "\033[34m→\033[0m $task"
    else
        echo "→ $task"
    fi
}

display_yang_request_info() {
    local url="$1"
    local method="${2:-GET}"

    if is_verbose_enabled; then
        echo
        format_yang_info "RESTCONF Request Details:"
        echo "  Method: $method"
        echo "  URL: $url"
        echo "  Format: ${argc_format:-json}"
        echo
    fi
}

display_yang_response_info() {
    local status_code="$1"
    local response_size="${2:-unknown}"

    if is_verbose_enabled; then
        echo
        format_yang_info "RESTCONF Response:"
        echo "  Status: $status_code"
        echo "  Size: $response_size bytes"
        echo
    fi
}

format_yang_output() {
    local content="$1"
    local format="${argc_format:-json}"
    local raw_output="${argc_raw:-false}"

    # If raw output is requested, just output as-is
    if [[ "$raw_output" == "true" ]]; then
        echo "$content"
        return 0
    fi

    # Format JSON output
    if [[ "$format" == "json" ]]; then
        if command -v jq >/dev/null 2>&1; then
            echo "$content" | jq '.'
        else
            echo "$content"
        fi
    else
        # XML or other formats - output as-is for now
        echo "$content"
    fi
}

display_yang_operation_results() {
    local exit_code="$1"
    local operation="$2"
    local target="${3:-}"

    echo
    if [[ "$exit_code" -eq 0 ]]; then
        format_yang_success "$operation completed successfully"
        [[ -n "$target" ]] && format_yang_info "Target: $target"
    else
        format_yang_error "$operation failed"
        format_yang_info "Check the output above for details"
        format_yang_info "Verify controller connectivity and credentials"
    fi
}
