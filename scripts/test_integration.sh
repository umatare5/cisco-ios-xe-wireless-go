#!/usr/bin/env bash

# @meta version 1.0.0
# @meta author "@umatare5"
# @describe Integration Tests - Run Go integration tests against WNC controllers

# @option -p --project <DIR>           Project root directory [default: .]
# @flag   -v --verbose                Enable verbose test output
# @flag      --race                   Enable race detection [default: true]
# @option -t --timeout <DURATION>     Test timeout [default: 10m]
# @option    --package <PATTERN>      Package pattern to test [default: ./...]
# @flag      --check-env-only         Only check environment without running tests
# @flag      --no-color               Disable colored output

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
readonly SCRIPT_DIR

# Source bootstrap library
source "${SCRIPT_DIR}/lib/bootstrap.sh"

# Initialize libraries and load testing module
init_wnc_libraries "${SCRIPT_DIR}/lib/test_integration"

# Validate required CLI tools
validate_required_cli_tools "standard"

# Entrypoint: run integration tests
main() {
    local project_root="${argc_project:-.}"
    local timeout="${argc_timeout:-10m}"

    run_integration_test_operation "$project_root" "$timeout"
}

eval "$(argc --argc-eval "$0" "$@")"
