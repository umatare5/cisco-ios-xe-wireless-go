#!/usr/bin/env bash
# @meta version 1.0.0
# @meta author "@umatare5"
# @describe Cisco WNC Integration Tests - Run Go integration tests against WNC controllers

# @option -p --project <DIR>           Project root directory [default: .]
# @flag   -v --verbose                 Enable verbose test output
# @flag      --race                    Enable race detection [default: true]
# @option -t --timeout <DURATION>      Test timeout [default: 10m]
# @option    --package <PATTERN>       Package pattern to test [default: ./...]
# @flag      --check-env-only          Only check environment without running tests
# @flag      --no-color                Disable colored output

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
MODULE_DIR="${SCRIPT_DIR}/lib/testing"

# Source shared libraries
source "${SCRIPT_DIR}/lib/common/common.sh"
SOURCE_WNC_LIBRARIES "$SCRIPT_DIR"

# Source module-specific libraries
source "${MODULE_DIR}/help.sh"
source "${MODULE_DIR}/output.sh"
source "${MODULE_DIR}/core.sh"

# Predicate functions for improved readability
is_verbose_enabled() { [[ "${argc_verbose:-0}" == "1" ]]; }
is_no_color_enabled() { [[ "${argc_no_color:-0}" == "1" ]]; }
is_check_env_only() { [[ "${argc_check_env_only:-false}" == "true" ]]; }
is_command_available() { command -v "${1:-}" >/dev/null 2>&1; }

main() {
    run_integration_test_operation
}

eval "$(argc --argc-eval "$0" "$@")"
