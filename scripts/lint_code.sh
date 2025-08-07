#!/usr/bin/env bash
# @meta version 1.0.0
# @meta author "@umatare5"
# @describe Cisco WNC Lint Script - Code linting with golangci-lint

# @option -p --project <DIR>           Project root directory [default: .]
# @flag   -v --verbose                 Enable verbose output
# @flag      --fix                     Automatically fix issues where possible
# @option    --config <FILE>           Custom golangci-lint config file path
# @flag      --no-color                Disable colored output

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Source bootstrap library
source "${SCRIPT_DIR}/lib/bootstrap.sh"

# Initialize WNC libraries with lint_code module
init_wnc_libraries "$SCRIPT_DIR" "${SCRIPT_DIR}/lib/lint_code"

# Validate required CLI tools before proceeding
validate_required_cli_tools "standard"

# Predicate functions for improved readability using argc validation helpers
is_verbose_enabled() { is_enabled "${argc_verbose:-0}"; }
is_auto_fix_enabled() { is_true "${argc_fix:-false}"; }
is_no_color_enabled() { is_enabled "${argc_no_color:-0}"; }
is_command_available() { command -v "${1:-}" >/dev/null 2>&1; }
has_golangci_lint() { is_command_available golangci-lint; }

main() {
    run_lint_operation
}

eval "$(argc --argc-eval "$0" "$@")"
