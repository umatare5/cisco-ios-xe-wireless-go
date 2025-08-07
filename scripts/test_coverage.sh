#!/usr/bin/env bash
# @meta version 1.0.0
# @meta author "@umatare5"
# @describe Coverage Tests - Run Go tests with comprehensive coverage analysis

# @option -p --project <DIR>           Project root directory [default: .]
# @option -o --output <FILE>           Coverage output file [default: ./tmp/coverage.out]
# @flag   -v --verbose                 Enable verbose test output
# @flag   -s --short                   Run tests in short mode (skip long-running tests)
# @option -t --timeout <DURATION>      Test timeout duration [default: 30s]
# @flag      --no-color                Disable colored output

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
MODULE_DIR="${SCRIPT_DIR}/lib/testing"

# Source shared libraries
source "${SCRIPT_DIR}/lib/common/common.sh"

# Initialize all libraries using unified function
init_script_libraries "$SCRIPT_DIR" "$MODULE_DIR"

# Validate required CLI tools before proceeding
validate_required_cli_tools "standard"

# Predicate functions for improved readability using argc validation helpers
is_verbose_enabled() { is_enabled "${argc_verbose:-0}"; }
is_no_color_enabled() { is_enabled "${argc_no_color:-0}"; }
is_short_mode_enabled() { is_true "${argc_short:-false}"; }
is_command_available() { command -v "${1:-}" >/dev/null 2>&1; }

main() {
    run_coverage_test_operation
}

eval "$(argc --argc-eval "$0" "$@")"
