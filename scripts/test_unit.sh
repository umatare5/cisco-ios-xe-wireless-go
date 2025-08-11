#!/usr/bin/env bash

# @meta version 1.0.0
# @meta author "@umatare5"
# @describe Unit Tests - Run Go unit tests with coverage and reporting

# @option -p --project <DIR>           Project root directory [default: .]
# @flag   -v --verbose                Enable verbose test output
# @flag   -s --short                  Run tests in short mode (skip long-running tests)
# @flag   -c --coverage               Generate coverage data
# @option -o --output <FILE>          Coverage output file [default: ./tmp/coverage.out]
# @option -t --timeout <DURATION>     Test timeout duration [default: 30s]
# @flag      --no-color               Disable colored output

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
readonly SCRIPT_DIR

# Source bootstrap library
source "${SCRIPT_DIR}/lib/bootstrap.sh"

# Initialize libraries and load testing module
init_wnc_libraries "${SCRIPT_DIR}/lib/test_unit"

# Validate required CLI tools
validate_required_cli_tools "standard"

# Entrypoint: run unit tests (or coverage when enabled)
main() {
    local project_root="${argc_project:-.}"
    local timeout="${argc_timeout:-30s}"

    if [[ "${argc_coverage:-0}" == "1" ]]; then
        # If coverage is requested, use coverage test operation
        local coverage_file="${argc_output:-./tmp/coverage.out}"
        run_coverage_test_operation "$project_root" "$coverage_file"
        return
    fi

    run_unit_test_operation "$project_root" "$timeout"
}

eval "$(argc --argc-eval "$0" "$@")"
