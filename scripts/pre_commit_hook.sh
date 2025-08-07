#!/usr/bin/env bash
# @meta version 1.0.0
# @meta author "@umatare5"
# @describe Cisco WNC Pre-commit Validation - Validate commits before they are made

# @flag   -v --verbose                Enable verbose output
# @flag      --no-color               Disable colored output

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
readonly SCRIPT_DIR
readonly MODULE_DIR="${SCRIPT_DIR}/lib/validation"

# Source bootstrap library
source "${SCRIPT_DIR}/lib/bootstrap.sh"

# Initialize WNC libraries with validation module
init_wnc_libraries "${SCRIPT_DIR}" "${SCRIPT_DIR}/lib/validation"

# Validate required CLI tools before proceeding
validate_required_cli_tools "minimal"

# Predicate functions for improved readability using argc validation helpers
is_verbose_enabled() { is_true "${argc_verbose:-false}"; }
is_no_color_enabled() { is_true "${argc_no_color:-false}"; }

main() {
    run_pre_commit_validation
}

# Check for help request before argc evaluation
if [[ "${1:-}" == "--help" ]] || [[ "${1:-}" == "-h" ]]; then
    # Load only essential libraries for help
    source "${MODULE_DIR}/help.sh"
    show_pre_commit_help
    exit 0
fi

eval "$(argc --argc-eval "$0" "$@")"
