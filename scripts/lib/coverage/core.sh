#!/usr/bin/env bash

# Cisco WNC Coverage HTML Generator - Core Functions
# Core business logic for coverage HTML generation operations

set -euo pipefail

# Source bootstrap library
LIB_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
source "${LIB_ROOT}/bootstrap.sh"
init_wnc_basic

validate_coverage_environment() {
    local project_root="$1"
    local input_file="$2"

    # Validate project directory
    if ! validate_project_directory "$project_root"; then
        format_coverage_error "Invalid project directory: $project_root"
        return 1
    fi

    # Check if input coverage file exists
    if ! file_exists "$input_file"; then
        format_coverage_error "Coverage file not found: $input_file"
        format_coverage_info "Run coverage tests first: ./scripts/test_coverage.sh"
        return 1
    fi

    # Check if Go toolchain is available
    if ! is_command_available go; then
        format_coverage_error "Go toolchain is not installed or not in PATH"
        return 1
    fi

    # Verify the coverage file format
    if ! head -1 "$input_file" | grep -q "mode:"; then
        format_coverage_warning "Input file may not be a valid Go coverage file"
        format_coverage_info "Expected coverage file format: mode: set|count|atomic"
    fi

    return 0
}

prepare_output_directory() {
    local output_file="$1"
    local output_dir

    output_dir="$(dirname "$output_file")"

    # Create output directory if it doesn't exist
    if ! dir_exists "$output_dir"; then
        if ! mkdir -p "$output_dir"; then
            format_coverage_error "Failed to create output directory: $output_dir"
            return 1
        fi
        is_verbose_enabled && format_coverage_info "Created output directory: $output_dir"
    fi

    # Check if output directory is writable
    if [[ ! -w "$output_dir" ]]; then
        format_coverage_error "Output directory is not writable: $output_dir"
        return 1
    fi

    return 0
}

execute_coverage_html_generation() {
    local project_root="$1"
    local input_file="$2"
    local output_file="$3"

    # Change to project directory
    local original_pwd="$PWD"
    cd "$project_root" || {
        format_coverage_error "Failed to change to project directory: $project_root"
        return 1
    }

    is_verbose_enabled && format_coverage_info "Generating HTML coverage report..."

    # Generate HTML coverage report using go tool cover
    local exit_code=0
    go tool cover -html="$input_file" -o "$output_file" || exit_code=$?

    # Return to original directory
    cd "$original_pwd" || {
        format_coverage_warning "Failed to return to original directory"
    }

    return "$exit_code"
}

# write_coverprofile_artifact copies the original Go coverprofile to the report file.
# It purposefully avoids deep nesting for readability and emits warnings rather than failing the run.
write_coverprofile_artifact() {
    local exit_code="$1" input_file="$2" report_file="$3"

    # Only proceed if HTML generation succeeded
    [[ "$exit_code" -ne 0 ]] && return 0

    # Copy raw coverprofile (atomic|set|count) for octocov consumption
    if ! cp "$input_file" "$report_file" 2>/dev/null; then
        format_coverage_warning "Failed to write coverprofile to $report_file"
        return 0
    fi

    # Validate header for downstream tools
    if ! head -1 "$report_file" | grep -q '^mode:'; then
        format_coverage_warning "report.out missing mode header; octocov may fail"
    fi

    is_verbose_enabled && format_coverage_info "Wrote coverprofile: $report_file"
    return 0
}

run_coverage_html_operation() {
    local project_root="${argc_project:-.}"
    local input_file="${argc_input:-./tmp/coverage.out}"
    local output_file="${argc_output:-./coverage/coverage.html}"
    local report_file="${argc_report:-coverage/report.out}"

    # Convert relative paths to absolute paths
    [[ "$input_file" != /* ]] && input_file="$project_root/$input_file"
    [[ "$output_file" != /* ]] && output_file="$project_root/$output_file"
    [[ "$report_file" != /* ]] && report_file="$project_root/$report_file"

    # Show banner and info
    show_coverage_html_banner
    show_coverage_html_verbose_info

    # Display file information
    display_coverage_file_info "$input_file" "$output_file"

    # Validate environment
    if ! validate_coverage_environment "$project_root" "$input_file"; then
        return 1
    fi

    # Prepare output directories
    if ! prepare_output_directory "$output_file"; then
        return 1
    fi
    if ! prepare_output_directory "$report_file"; then
        return 1
    fi

    # Execute HTML generation
    show_coverage_progress "Generating HTML coverage report..."
    local exit_code=0
    execute_coverage_html_generation "$project_root" "$input_file" "$output_file" || exit_code=$?

    # Produce coverprofile artifact (octocov expects raw coverprofile format, not func summary)
    write_coverprofile_artifact "$exit_code" "$input_file" "$report_file"

    # Display results (HTML primary output)
    display_coverage_html_results "$exit_code" "$output_file"

    return "$exit_code"
}
