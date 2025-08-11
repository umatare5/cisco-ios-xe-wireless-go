#!/usr/bin/env bash

# @meta version 1.0.0
# @meta author "@umatare5"
# @describe Generate HTML coverage report from coverage data

# @option -p --project <DIR>           Project root directory [default: .]
# @option -i --input <FILE>            Coverage input file [default: ./tmp/coverage.out]
# @option -o --output <FILE>           HTML output file [default: ./coverage/report.html]
# @option    --report <FILE>           Coverprofile artifact file [default: ./coverage/report.out]
# @flag   -v --verbose                Enable verbose output
# @flag      --no-color               Disable colored output

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
readonly SCRIPT_DIR

# Source bootstrap library
source "${SCRIPT_DIR}/lib/bootstrap.sh"

# Initialize libraries and load coverage module
init_wnc_libraries "${SCRIPT_DIR}/lib/generate_coverage_report"

# Validate required CLI tools
validate_required_cli_tools "docs"

# Entrypoint: generate HTML coverage report
main() {
    local project_root="${argc_project:-.}"
    local input_file="${argc_input:-./tmp/coverage.out}"
    local output_file="${argc_output:-./coverage/report.html}"
    local report_file="${argc_report:-./coverage/report.out}"

    run_coverage_html_operation "$project_root" "$input_file" "$output_file" "$report_file"
}

eval "$(argc --argc-eval "$0" "$@")"
