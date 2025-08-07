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
MODULE_DIR="${SCRIPT_DIR}/lib/lint_code"

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
is_auto_fix_enabled() { [[ "${argc_fix:-false}" == "true" ]]; }
is_no_color_enabled() { [[ "${argc_no_color:-0}" == "1" ]]; }
is_command_available() { command -v "${1:-}" >/dev/null 2>&1; }
has_golangci_lint() { is_command_available golangci-lint; }

main() {
    run_lint_operation
}

eval "$(argc --argc-eval "$0" "$@")"
