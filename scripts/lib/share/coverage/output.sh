#!/usr/bin/env bash

# Cisco WNC Coverage - Output functions for HTML report generation
# Provides functions to format and display coverage report messages

# Display detailed information about coverage input and output files
display_coverage_file_info() {
    local input_file="$1"
    local output_file="$2"

    if is_verbose_enabled; then
    printf '\n'
        info "Processing coverage data:"
    printf '%s\n' "  Input:  $input_file"
    printf '%s\n' "  Output: $output_file"

        if [[ -f "$input_file" ]]; then
            local file_size
            file_size=$(stat -f%z "$input_file" 2>/dev/null || \
                       stat -c%s "$input_file" 2>/dev/null || \
                       printf '%s\n' "unknown")
            info "Input file size: $file_size bytes"
        fi
        printf '\n'
    fi
}

# Display final results and instructions for generated coverage report
display_coverage_html_results() {
    local exit_code="$1"
    local output_file="$2"

    printf '\n'

    if [[ "$exit_code" -eq 0 ]]; then
        success "HTML coverage report generated successfully"
        info "Report location: $output_file"

        if [[ -f "$output_file" ]]; then
            local file_size
            file_size=$(stat -f%z "$output_file" 2>/dev/null || \
                       stat -c%s "$output_file" 2>/dev/null || \
                       printf '%s\n' "unknown")
            info "Report size: $file_size bytes"
        fi

        printf '\n'
        info "To view the report:"
        printf '%s\n' "  open $output_file"
        return
    fi

    error "Failed to generate HTML coverage report"
    info "Check the input file and try again"
}
