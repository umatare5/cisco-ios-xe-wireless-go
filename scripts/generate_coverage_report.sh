#!/usr/bin/env bash
# @meta version 1.0.0
# @meta author "@umatare5"
# @describe Generate HTML coverage report from coverage data

# @option -p --project <DIR>           Project root directory [default: .]
# @option -i --input <FILE>            Coverage input file [default: ./tmp/coverage.out]
# @option -o --output <FILE>           HTML output file [default: ./tmp/coverage.html]
# @flag   -v --verbose                Enable verbose output
# @flag      --no-color               Disable colored output

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
readonly SCRIPT_DIR

# Source bootstrap library
source "${SCRIPT_DIR}/lib/bootstrap.sh"

# Initialize WNC libraries with coverage module
init_wnc_libraries "${SCRIPT_DIR}" "${SCRIPT_DIR}/lib/coverage"

# Validate required CLI tools before proceeding
validate_required_cli_tools "standard"

# Predicate functions centralized in lib/core/predicates.sh

main() {
    run_coverage_html_operation
}

eval "$(argc --argc-eval "$0" "$@")"
