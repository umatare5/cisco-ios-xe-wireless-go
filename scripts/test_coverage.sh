#!/usr/bin/env bash
# @meta version 1.0.0
# @meta author "@umatare5"
# @describe Coverage Tests - Run Go tests with comprehensive coverage analysis

# @option -p --project <DIR>           Project root directory [default: .]
# @option -o --output <FILE>           Coverage output file [default: ./tmp/coverage.out]
# @flag   -v --verbose                Enable verbose test output
# @flag   -s --short                  Run tests in short mode (skip long-running tests)
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
    run_coverage_test_operation
}

eval "$(argc --argc-eval "$0" "$@")"
