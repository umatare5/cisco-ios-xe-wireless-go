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

# Initialize WNC libraries with testing module
init_wnc_libraries "${SCRIPT_DIR}" "${SCRIPT_DIR}/lib/testing"

# Validate required CLI tools before proceeding
validate_required_cli_tools "standard"

# Predicate functions for improved readability using argc validation helpers
is_verbose_enabled() { is_true "${argc_verbose:-false}"; }
is_no_color_enabled() { is_true "${argc_no_color:-false}"; }
is_check_env_only() { is_true "${argc_check_env_only:-false}"; }
is_command_available() { command -v "${1:-}" >/dev/null 2>&1; }

main() {
    run_integration_test_operation
}

eval "$(argc --argc-eval "$0" "$@")"
