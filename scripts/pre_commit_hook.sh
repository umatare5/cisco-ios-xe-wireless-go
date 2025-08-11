#!/usr/bin/env bash

# @meta version 1.0.0
# @meta author "@umatare5"
# @describe Cisco WNC Pre-commit Validation - Validate commits before they are made

# @flag   -v --verbose                Enable verbose output
# @flag      --no-color               Disable colored output

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
readonly SCRIPT_DIR

# Source bootstrap library
source "${SCRIPT_DIR}/lib/bootstrap.sh"

# Initialize libraries and load validation module
init_wnc_libraries "${SCRIPT_DIR}/lib/pre_commit_hook"

# Handle help early (loaded via init_wnc_libraries)
if [[ "${1:-}" == "--help" ]] || [[ "${1:-}" == "-h" ]]; then
    show_pre_commit_help
    exit 0
fi

# Validate required CLI tools
validate_required_cli_tools "standard"

# Entrypoint: run pre-commit validations
main() {
    run_pre_commit_validation
}

eval "$(argc --argc-eval "$0" "$@")"
