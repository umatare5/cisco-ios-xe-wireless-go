#!/usr/bin/env bash

# Cisco WNC Coverage HTML Generator - Help Functions
# Provides help and documentation functionality for coverage HTML generation

show_coverage_html_banner() {
    wnc_banner_coverage
}

show_coverage_html_help() {
    show_coverage_html_banner
    cat << 'EOF'
USAGE:
    generate_coverage_report.sh [OPTIONS]

DESCRIPTION:
    Generates an HTML coverage report from Go coverage data files.
    Creates interactive HTML reports for visualizing test coverage.

OPTIONS:
    -p, --project <DIR>      Project root directory [default: .]
    -i, --input <FILE>       Coverage input file [default: ./tmp/coverage.out]
    -o, --output <FILE>      HTML output file [default: ./tmp/coverage.html]
    -v, --verbose            Enable verbose output
    -h, --help               Show this help message

EXAMPLES:
    # Generate basic HTML report
    generate_coverage_report.sh

    # Specify custom input and output
    generate_coverage_report.sh -i coverage.out -o report.html

    # Generate for specific project
    generate_coverage_report.sh --project /path/to/project

REQUIREMENTS:
    - Go toolchain must be installed
    - Coverage data file must exist (run coverage tests first)

WORKFLOW:
    1. Run coverage tests: ./scripts/test_unit.sh --coverage or ./scripts/test_integration.sh --coverage
    2. Generate HTML report: ./scripts/generate_coverage_report.sh
    3. Open report in browser: open ./tmp/coverage.html

EOF
}

show_coverage_html_verbose_info() {
    local project="$1"
    local input_file="$2"
    local output_file="$3"

    is_verbose_enabled || return 0

    printf '%s\n' "Coverage HTML Generation Configuration:"
    printf '%s\n' "  Project: ${project:-.}"
    printf '%s\n' "  Input file: ${input_file:-./tmp/coverage.out}"
    printf '%s\n' "  Output file: ${output_file:-./tmp/coverage.html}"
    printf '%s\n' "  Verbose: enabled"
    printf '\n'
}
