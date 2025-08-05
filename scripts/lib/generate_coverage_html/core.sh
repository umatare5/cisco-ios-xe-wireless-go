#!/usr/bin/env bash

# Cisco WNC Coverage HTML Generator - Core Functions
# Core business logic for coverage HTML generation operations

# Source common predicates
source "$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)/../common/common.sh"

validate_coverage_environment() {
    local project_root="$1"
    local input_file="$2"

    # Validate project directory
    if ! validate_project_directory "$project_root"; then
        format_coverage_error "Invalid project directory: $project_root"
        return 1
    fi

    # Check if input coverage file exists
    if ! is_file_exists "$input_file"; then
        format_coverage_error "Coverage file not found: $input_file"
        format_coverage_info "Run coverage tests first: ./scripts/run_coverage_tests.sh"
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
    if ! is_directory_exists "$output_dir"; then
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

run_coverage_html_operation() {
    local project_root="${argc_project:-.}"
    local input_file="${argc_input:-./tmp/coverage.out}"
    local output_file="${argc_output:-./tmp/coverage.html}"

    # Convert relative paths to absolute paths
    [[ "$input_file" != /* ]] && input_file="$project_root/$input_file"
    [[ "$output_file" != /* ]] && output_file="$project_root/$output_file"

    # Show banner and info
    show_coverage_html_banner
    show_coverage_html_verbose_info

    # Display file information
    display_coverage_file_info "$input_file" "$output_file"

    # Validate environment
    if ! validate_coverage_environment "$project_root" "$input_file"; then
        return 1
    fi

    # Prepare output directory
    if ! prepare_output_directory "$output_file"; then
        return 1
    fi

    # Execute HTML generation
    show_coverage_progress "Generating HTML coverage report..."
    local exit_code=0
    execute_coverage_html_generation "$project_root" "$input_file" "$output_file" || exit_code=$?

    # Display results
    display_coverage_html_results "$exit_code" "$output_file"

    return "$exit_code"
}
