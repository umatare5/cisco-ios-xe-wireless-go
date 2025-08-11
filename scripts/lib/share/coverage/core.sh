#!/usr/bin/env bash

# Cisco WNC Coverage HTML Generator - Core Functions
# Provides core business logic for generating HTML coverage reports

# Validate coverage generation environment and prerequisites
validate_coverage_environment() {
    local project_root="$1"
    local input_file="$2"

    # Validate project directory
    if ! validate_project_directory "$project_root"; then
        error "Invalid project directory: $project_root"
        return 1
    fi

    # Check if input coverage file exists
    if ! file_exists "$input_file"; then
        error "Coverage file not found: $input_file"
        info "Run coverage tests first: ./scripts/test_unit.sh --coverage or ./scripts/test_integration.sh --coverage"
        return 1
    fi

    # Check if Go toolchain is available
    if ! is_command_available go; then
        error "Go toolchain is not installed or not in PATH"
        return 1
    fi

    # Verify the coverage file format
    if ! head -1 "$input_file" | grep -q "mode:"; then
        warn "Input file may not be a valid Go coverage file"
        info "Expected coverage file format: mode: set|count|atomic"
    fi

    return 0
}

# Create and validate output directory for coverage report
prepare_output_directory() {
    local output_file="$1"
    local output_dir

    output_dir="$(dirname "$output_file")"

    # Create output directory if it doesn't exist
    if ! dir_exists "$output_dir"; then
        if ! mkdir -p "$output_dir"; then
            error "Failed to create output directory: $output_dir"
            return 1
        fi
        is_verbose_enabled && info "Created output directory: $output_dir"
    fi

    # Check if output directory is writable
    if [[ ! -w "$output_dir" ]]; then
        error "Output directory is not writable: $output_dir"
        return 1
    fi

    return 0
}

# Generate HTML coverage report using Go toolchain
execute_coverage_html_generation() {
    local project_root="$1"
    local input_file="$2"
    local output_file="$3"

    # Change to project directory
    local original_pwd="$PWD"
    cd "$project_root" || {
        error "Failed to change to project directory: $project_root"
        return 1
    }

    is_verbose_enabled && info "Generating HTML coverage report..."

    # Generate HTML coverage report using go tool cover
    local exit_code=0
    go tool cover -html="$input_file" -o "$output_file" || exit_code=$?

    # Return to original directory
    cd "$original_pwd" || {
        warn "Failed to return to original directory"
    }

    return "$exit_code"
}

# Copy raw coverprofile for downstream tooling consumption
write_coverprofile_artifact() {
    local exit_code="$1" input_file="$2" report_file="$3"

    # Only proceed if HTML generation succeeded
    if [[ "$exit_code" -ne 0 ]]; then
        return 0
    fi

    # Copy raw coverprofile (atomic|set|count) for octocov consumption
    if ! cp "$input_file" "$report_file" 2>/dev/null; then
        warn "Failed to write coverprofile to $report_file"
        return 0
    fi

    # Validate header for downstream tools
    if ! head -1 "$report_file" | grep -q '^mode:'; then
        warn "report.out missing mode header; octocov may fail"
    fi

    is_verbose_enabled && info "Wrote coverprofile: $report_file"
    return 0
}
