#!/usr/bin/env bash

# Cisco WNC Coverage Report - Core Operations
# Provides core functionality for HTML coverage report generation

run_coverage_html_operation() {
    local project_root="$1"
    local input_file="$2"
    local output_file="$3"
    local report_file="$4"

    # Prepare and standardize file paths for coverage processing
    local file_paths
    file_paths="$(prepare_coverage_file_paths "$project_root" "$input_file" "$output_file" "$report_file")"
    input_file="$(echo "$file_paths" | sed -n '1p')"
    output_file="$(echo "$file_paths" | sed -n '2p')"
    report_file="$(echo "$file_paths" | sed -n '3p')"

    # Validate environment and required dependencies
    if ! validate_coverage_dependencies "$project_root" "$input_file" "$output_file" "$report_file"; then
        return 1
    fi

    # Execute coverage generation process
    execute_coverage_generation_process "$project_root" "$input_file" "$output_file" "$report_file"
}
