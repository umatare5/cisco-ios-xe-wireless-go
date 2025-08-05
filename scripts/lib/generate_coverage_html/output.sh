#!/usr/bin/env bash

# Cisco WNC Coverage HTML Generator - Output Functions
# Handles output formatting and display for coverage HTML operations

format_coverage_error() {
    local message="$1"
    if ! is_no_color_enabled; then
        echo -e "\033[31m✗ Error:\033[0m $message" >&2
    else
        echo "✗ Error: $message" >&2
    fi
}

format_coverage_success() {
    local message="$1"
    if ! is_no_color_enabled; then
        echo -e "\033[32m✓\033[0m $message"
    else
        echo "✓ $message"
    fi
}

format_coverage_warning() {
    local message="$1"
    if ! is_no_color_enabled; then
        echo -e "\033[33m⚠ Warning:\033[0m $message" >&2
    else
        echo "⚠ Warning: $message" >&2
    fi
}

format_coverage_info() {
    local message="$1"
    if ! is_no_color_enabled; then
        echo -e "\033[36mℹ Info:\033[0m $message"
    else
        echo "ℹ Info: $message"
    fi
}

show_coverage_progress() {
    local task="$1"

    if ! is_no_color_enabled; then
        echo -e "\033[34m→\033[0m $task"
    else
        echo "→ $task"
    fi
}

display_coverage_file_info() {
    local input_file="$1"
    local output_file="$2"

    if is_verbose_enabled; then
        echo
        format_coverage_info "Processing coverage data:"
        echo "  Input:  $input_file"
        echo "  Output: $output_file"

        if [[ -f "$input_file" ]]; then
            local file_size
            file_size=$(stat -f%z "$input_file" 2>/dev/null || stat -c%s "$input_file" 2>/dev/null || echo "unknown")
            format_coverage_info "Input file size: $file_size bytes"
        fi
        echo
    fi
}

display_coverage_html_results() {
    local exit_code="$1"
    local output_file="$2"

    echo
    if [[ "$exit_code" -eq 0 ]]; then
        format_coverage_success "HTML coverage report generated successfully"
        format_coverage_info "Report location: $output_file"

        if [[ -f "$output_file" ]]; then
            local file_size
            file_size=$(stat -f%z "$output_file" 2>/dev/null || stat -c%s "$output_file" 2>/dev/null || echo "unknown")
            format_coverage_info "Report size: $file_size bytes"
        fi

        echo
        format_coverage_info "To view the report:"
        echo "  open $output_file"
    else
        format_coverage_error "Failed to generate HTML coverage report"
        format_coverage_info "Check the input file and try again"
    fi
}
