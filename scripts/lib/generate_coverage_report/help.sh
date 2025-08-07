#!/usr/bin/env bash
set -euo pipefail

# Cisco WNC Coverage HTML Generator - Help Functions
# Provides help and documentation functionality for coverage HTML generation

# Source bootstrap library
LIB_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
# shellcheck disable=SC1090  # Dynamic source loading
source "${LIB_ROOT}/bootstrap.sh"
init_wnc_basic

show_coverage_html_banner() {
    if ! is_no_color_enabled; then
        echo -e "\033[32m----------------------------------------\033[0m"
        echo -e "\033[32m      Coverage HTML Generator          \033[0m"
        echo -e "\033[32m      Go Tool Cover Integration        \033[0m"
        echo -e "\033[32m----------------------------------------\033[0m"
    else
        echo "----------------------------------------"
        echo "      Coverage HTML Generator"
        echo "      Go Tool Cover Integration"
        echo "----------------------------------------"
    fi
    echo
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
    1. Run coverage tests: ./scripts/test_coverage.sh
    2. Generate HTML report: ./scripts/generate_coverage_report.sh
    3. Open report in browser: open ./tmp/coverage.html

EOF
}

show_coverage_html_verbose_info() {
    is_verbose_enabled || return 0

    echo "Coverage HTML Generation Configuration:"
    echo "  Project: ${argc_project:-.}"
    echo "  Input file: ${argc_input:-./tmp/coverage.out}"
    echo "  Output file: ${argc_output:-./tmp/coverage.html}"
    echo "  Verbose: enabled"
    echo
}
