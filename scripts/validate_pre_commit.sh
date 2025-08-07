#!/usr/bin/env bash
# @meta version 1.0.0
# @meta author "@umatare5"
# @meta description Cisco WNC Pre-commit Validation - Validate commits before they are made

# @flag   -v --verbose                 Enable verbose output
# @flag      --no-color                Disable colored output

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
MODULE_DIR="${SCRIPT_DIR}/lib/pre_commit_validation"

# Source shared libraries
source "${SCRIPT_DIR}/lib/common/common.sh"
SOURCE_WNC_LIBRARIES "$SCRIPT_DIR"

# Source module-specific libraries
source "${MODULE_DIR}/help.sh"
source "${MODULE_DIR}/output.sh"
source "${MODULE_DIR}/core.sh"

# Validate required CLI tools before proceeding
validate_required_cli_tools "minimal"

# Predicate functions for improved readability
is_verbose_enabled() { [[ "${argc_verbose:-0}" == "1" ]]; }
is_no_color_enabled() { [[ "${argc_no_color:-0}" == "1" ]]; }

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
