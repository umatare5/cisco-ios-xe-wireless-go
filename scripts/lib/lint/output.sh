#!/usr/bin/env bash

# Cisco WNC Lint Code - Output Functions
# Handles output formatting and display for lint operations

set -euo pipefail

format_lint_error() {
    local message="$1"
    if ! is_no_color_enabled; then
        echo -e "\033[31m✗ Error:\033[0m $message" >&2
    else
        echo "✗ Error: $message" >&2
    fi
}

format_lint_success() {
    local message="$1"
    if ! is_no_color_enabled; then
        echo -e "\033[32m✓\033[0m $message"
    else
        echo "✓ $message"
    fi
}

format_lint_warning() {
    local message="$1"
    if ! is_no_color_enabled; then
        echo -e "\033[33m⚠ Warning:\033[0m $message" >&2
    else
        echo "⚠ Warning: $message" >&2
    fi
}

format_lint_info() {
    local message="$1"
    if ! is_no_color_enabled; then
        echo -e "\033[36mℹ Info:\033[0m $message"
    else
        echo "ℹ Info: $message"
    fi
}

show_lint_progress() {
    local current="$1"
    local total="$2"
    local task="$3"

    if ! is_no_color_enabled; then
        echo -e "\033[34m[$current/$total]\033[0m $task"
    else
        echo "[$current/$total] $task"
    fi
}

display_lint_results() {
    local exit_code="$1"
    local project_path="$2"

    echo
    if [[ "$exit_code" -eq 0 ]]; then
        format_lint_success "Code linting completed successfully"
        is_verbose_enabled && format_lint_info "Project: $project_path"
    else
        format_lint_error "Code linting failed with issues"
        format_lint_info "Check the output above for details"
        format_lint_info "Consider using --fix to auto-correct issues"
    fi
}
