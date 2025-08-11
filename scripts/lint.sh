#!/usr/bin/env bash

# @meta version 1.0.0
# @meta author "@umatare5"
# @describe Cisco WNC Lint Script - Code linting with golangci-lint

# @option -p --project <DIR>           Project root directory [default: .]
# @flag   -v --verbose                Enable verbose output
# @flag      --fix                    Automatically fix issues where possible
# @option    --config <FILE>          Custom golangci-lint config file path
# @flag      --no-color               Disable colored output

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
readonly SCRIPT_DIR

# Source bootstrap library
source "${SCRIPT_DIR}/lib/bootstrap.sh"

# Initialize libraries and load lint module
init_wnc_libraries "${SCRIPT_DIR}/lib/lint"

# Validate required CLI tools
validate_required_cli_tools "standard"

# Entrypoint: run linters
main() {
    local project_root="${argc_project:-.}"

    # Temporarily disable immediate exit to capture status reliably
    set +e
    run_lint_operation "$project_root"
    local status=$?
    set -e
    return $status
}

eval "$(argc --argc-eval "$0" "$@")"
