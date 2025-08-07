#!/usr/bin/env bash
# @meta version 1.0.0
# @meta author "@umatare5"
# @describe Generate HTML coverage report from coverage data

# @option -p --project <DIR>           Project root directory [default: .]
# @option -i --input <FILE>            Coverage input file [default: ./tmp/coverage.out]
# @option -o --output <FILE>           HTML output file [default: ./tmp/coverage.html]
# @flag   -v --verbose                 Enable verbose output
# @flag      --no-color                Disable colored output

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Source bootstrap library
source "${SCRIPT_DIR}/lib/bootstrap.sh"

# Initialize WNC libraries with generate_coverage_html module
init_wnc_libraries "$SCRIPT_DIR" "${SCRIPT_DIR}/lib/generate_coverage_html"

# Validate required CLI tools before proceeding
validate_required_cli_tools "standard"

# Predicate functions for improved readability using argc validation helpers
is_verbose_enabled() { is_enabled "${argc_verbose:-0}"; }
is_no_color_enabled() { is_enabled "${argc_no_color:-0}"; }
is_file_exists() { [[ -f "${1:-}" ]]; }
is_directory_exists() { [[ -d "${1:-}" ]]; }
is_command_available() { command -v "${1:-}" >/dev/null 2>&1; }

main() {
    run_coverage_html_operation
}

eval "$(argc --argc-eval "$0" "$@")"
