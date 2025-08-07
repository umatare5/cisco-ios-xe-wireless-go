#!/usr/bin/env bash
# @meta version 1.0.0
# @meta author "@umatare5"
# @describe Cisco WNC Dependencies Management - Install, update, and clean Go dependencies

# @option -p --project <DIR>           Project root directory [default: .]
# @option    --golangci-lint <VERSION> golangci-lint version [default: latest]
# @option    --gotestsum <VERSION>     gotestsum version [default: latest]
# @flag   -v --verbose                 Enable verbose output
# @flag   -c --clean                   Clean module cache before installing
# @flag   -u --update                  Update all dependencies to latest versions
# @flag      --force                   Force reinstall even if exists
# @flag      --download-only           Download dependencies without installing
# @flag      --verify                  Verify dependencies after installation
# @flag      --no-color                Disable colored output

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Source bootstrap library
source "${SCRIPT_DIR}/lib/bootstrap.sh"

# Initialize WNC libraries with dependencies module
init_wnc_libraries "$SCRIPT_DIR" "${SCRIPT_DIR}/lib/dependencies"

# Predicate functions for improved readability using argc validation helpers
is_verbose_enabled() { is_enabled "${argc_verbose:-0}"; }
is_no_color_enabled() { is_enabled "${argc_no_color:-0}"; }
is_clean_enabled() { is_true "${argc_clean:-false}"; }
is_update_enabled() { is_true "${argc_update:-false}"; }
is_force_enabled() { is_true "${argc_force:-false}"; }
is_download_only_enabled() { is_true "${argc_download_only:-false}"; }
is_verify_enabled() { is_true "${argc_verify:-false}"; }
is_command_available() { command -v "${1:-}" >/dev/null 2>&1; }

main() {
    # Validate required CLI tools before proceeding
    validate_required_cli_tools "standard"

    run_dependencies_operation
}

eval "$(argc --argc-eval "$0" "$@")"
