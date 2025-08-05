#!/usr/bin/env bash
# @meta version 1.0.0
# @meta author "@umatare5"
# @describe Cisco WNC Artifacts Cleanup - Clean build artifacts, temporary files, and caches

# @option -p --project <DIR>           Project root directory [default: .]
# @flag   -v --verbose                 Enable verbose output
# @flag   -f --force                   Force removal without confirmation
# @flag      --go-cache                Clean Go build cache
# @flag      --go-modules              Clean Go module cache
# @flag      --temp-files              Clean temporary files (./tmp)
# @flag      --test-files              Clean test artifacts (.test binaries, coverage files)
# @flag      --all                     Clean all artifacts [default: true]
# @flag      --dry-run                 Show what would be cleaned without actually cleaning
# @flag      --no-color                Disable colored output

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
MODULE_DIR="${SCRIPT_DIR}/lib/artifacts"

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
is_force_enabled() { [[ "${argc_force:-false}" == "true" ]]; }
is_go_cache_enabled() { [[ "${argc_go_cache:-false}" == "true" ]]; }
is_go_modules_enabled() { [[ "${argc_go_modules:-false}" == "true" ]]; }
is_temp_files_enabled() { [[ "${argc_temp_files:-false}" == "true" ]]; }
is_test_files_enabled() { [[ "${argc_test_files:-false}" == "true" ]]; }
is_all_enabled() { [[ "${argc_all:-true}" == "true" ]]; }
is_dry_run_enabled() { [[ "${argc_dry_run:-false}" == "true" ]]; }
is_command_available() { command -v "${1:-}" >/dev/null 2>&1; }

main() {
    run_artifacts_operation
}

eval "$(argc --argc-eval "$0" "$@")"
