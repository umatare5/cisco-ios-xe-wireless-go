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
MODULE_DIR="${SCRIPT_DIR}/lib/dependencies"

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
is_clean_enabled() { [[ "${argc_clean:-false}" == "true" ]]; }
is_update_enabled() { [[ "${argc_update:-false}" == "true" ]]; }
is_force_enabled() { [[ "${argc_force:-false}" == "true" ]]; }
is_download_only_enabled() { [[ "${argc_download_only:-false}" == "true" ]]; }
is_verify_enabled() { [[ "${argc_verify:-false}" == "true" ]]; }
is_command_available() { command -v "${1:-}" >/dev/null 2>&1; }

main() {
    run_dependencies_operation
}

eval "$(argc --argc-eval "$0" "$@")"
