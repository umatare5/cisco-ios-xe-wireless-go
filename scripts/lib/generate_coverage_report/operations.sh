#!/usr/bin/env bash

# Cisco WNC Coverage Report - Helper Functions
# Provides functions to generate HTML coverage reports from Go coverage data files

# Prepare file paths by converting relative to absolute
prepare_coverage_file_paths() {
    local project_root="$1"
    local input_file="$2"
    local output_file="$3"
    local report_file="$4"

    # Convert relative paths to absolute
    if [[ "$input_file" != /* ]]; then
        input_file="$project_root/$input_file"
    fi

    if [[ "$output_file" != /* ]]; then
        output_file="$project_root/$output_file"
    fi

    if [[ "$report_file" != /* ]]; then
        report_file="$project_root/$report_file"
    fi

    printf '%s\n%s\n%s\n' "$input_file" "$output_file" "$report_file"
}

# Validate coverage generation environment and dependencies
validate_coverage_dependencies() {
    local project_root="$1"
    local input_file="$2"
    local output_file="$3"
    local report_file="$4"

    show_coverage_html_banner
    show_coverage_html_verbose_info "$project_root" "$input_file" "$output_file"
    display_coverage_file_info "$input_file" "$output_file"

    if ! validate_coverage_environment "$project_root" "$input_file"; then
        return 1
    fi
    if ! prepare_output_directory "$output_file"; then
        return 1
    fi
    if ! prepare_output_directory "$report_file"; then
        return 1
    fi

    return 0
}

# Execute coverage generation and write artifacts
execute_coverage_generation_process() {
    local project_root="$1"
    local input_file="$2"
    local output_file="$3"
    local report_file="$4"

    progress "Generating HTML coverage report..."
    local exit_code=0
    execute_coverage_html_generation "$project_root" "$input_file" "$output_file" || exit_code=$?

    write_coverprofile_artifact "$exit_code" "$input_file" "$report_file"
    display_coverage_html_results "$exit_code" "$output_file"
    return "$exit_code"
}
