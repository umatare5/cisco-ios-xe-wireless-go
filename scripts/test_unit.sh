#!/usr/bin/env bash
# @meta version 1.0.0
# @meta author "@umatare5"
# @describe Unit Tests - Run Go unit tests with coverage and reporting

# @option -p --project <DIR>           Project root directory [default: .]
# @flag   -v --verbose                Enable verbose test output
# @flag   -s --short                  Run tests in short mode (skip long-running tests)
# @flag   -c --coverage               Generate coverage data
# @option -t --timeout <DURATION>     Test timeout duration [default: 30s]
# @flag      --no-color               Disable colored output

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
readonly SCRIPT_DIR

# Source bootstrap library
source "${SCRIPT_DIR}/lib/bootstrap.sh"

# Initialize WNC libraries with testing module
init_wnc_libraries "${SCRIPT_DIR}" "${SCRIPT_DIR}/lib/testing"

# Validate required CLI tools before proceeding
validate_required_cli_tools "standard"

main() {
    if is_coverage_enabled; then
        # If coverage is requested, use coverage test operation
        argc_output="${argc_output:-./tmp/coverage.out}"
        run_coverage_test_operation
    else
        run_unit_test_operation
    fi
}

eval "$(argc --argc-eval "$0" "$@")"
