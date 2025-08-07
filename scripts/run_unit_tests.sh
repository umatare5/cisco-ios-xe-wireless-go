#!/usr/bin/env bash
# @meta version 1.0.0
# @meta author "@umatare5"
# @describe Cisco WNC Unit Tests - Run Go unit tests

# @option -p --project <DIR>           Project root directory [default: .]
# @flag   -v --verbose                 Enable verbose test output
# @flag   -s --short                   Run tests in short mode (skip long-running tests)
# @flag   -c --coverage                Generate coverage data
# @option -t --timeout <DURATION>      Test timeout duration [default: 30s]
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

# Validate required CLI tools before proceeding
validate_required_cli_tools "standard"

# Predicate functions for improved readability
is_verbose_enabled() { [[ "${argc_verbose:-0}" == "1" ]]; }
is_no_color_enabled() { [[ "${argc_no_color:-0}" == "1" ]]; }
is_short_mode_enabled() { [[ "${argc_short:-false}" == "true" ]]; }
is_coverage_enabled() { [[ "${argc_coverage:-false}" == "true" ]]; }
is_command_available() { command -v "${1:-}" >/dev/null 2>&1; }

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
