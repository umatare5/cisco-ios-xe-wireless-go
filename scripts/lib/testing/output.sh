#!/usr/bin/env bash

# Cisco WNC Testing Operations - Output Functions
# Handles output formatting and display for testing operations

format_test_error() {
    local message="$1"
    if ! is_no_color_enabled; then
        echo -e "\033[31m✗ Error:\033[0m $message" >&2
    else
        echo "✗ Error: $message" >&2
    fi
}

format_test_success() {
    local message="$1"
    if ! is_no_color_enabled; then
        echo -e "\033[32m✓\033[0m $message"
    else
        echo "✓ $message"
    fi
}

format_test_warning() {
    local message="$1"
    if ! is_no_color_enabled; then
        echo -e "\033[33m⚠ Warning:\033[0m $message" >&2
    else
        echo "⚠ Warning: $message" >&2
    fi
}

format_test_info() {
    local message="$1"
    if ! is_no_color_enabled; then
        echo -e "\033[36mℹ Info:\033[0m $message"
    else
        echo "ℹ Info: $message"
    fi
}

show_test_progress() {
    local task="$1"

    if ! is_no_color_enabled; then
        echo -e "\033[34m→\033[0m $task"
    else
        echo "→ $task"
    fi
}

display_test_summary() {
    local test_type="$1"
    local exit_code="$2"
    local duration="${3:-unknown}"
    local test_count="${4:-}"

    echo
    echo "========================================="
    if [[ "$exit_code" -eq 0 ]]; then
        format_test_success "$test_type tests completed successfully"
        [[ -n "$test_count" ]] && format_test_info "Tests executed: $test_count"
    else
        format_test_error "$test_type tests failed"
        format_test_info "Check the output above for details"
    fi

    [[ "$duration" != "unknown" ]] && format_test_info "Duration: $duration"
    echo "========================================="
}

display_coverage_summary() {
    local coverage_file="$1"
    local exit_code="$2"

    if [[ "$exit_code" -eq 0 && -f "$coverage_file" ]]; then
        echo
        format_test_info "Coverage report generated: $coverage_file"

        # Extract coverage percentage if available
        if command -v go >/dev/null 2>&1; then
            local coverage_percent
            coverage_percent=$(go tool cover -func="$coverage_file" 2>/dev/null | tail -1 | awk '{print $3}' || echo "unknown")
            if [[ "$coverage_percent" != "unknown" ]]; then
                format_test_info "Total coverage: $coverage_percent"
            fi
        fi
    fi
}
