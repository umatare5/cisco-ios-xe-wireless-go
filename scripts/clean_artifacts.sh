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

# Initialize all libraries using unified function
init_script_libraries "$SCRIPT_DIR" "$MODULE_DIR"

# Validate required CLI tools before proceeding
validate_required_cli_tools "minimal"

# Predicate functions for improved readability using argc validation helpers
is_verbose_enabled() { is_enabled "${argc_verbose:-0}"; }
is_no_color_enabled() { is_enabled "${argc_no_color:-0}"; }
is_force_enabled() { is_true "${argc_force:-false}"; }
is_go_cache_enabled() { is_true "${argc_go_cache:-false}"; }
is_go_modules_enabled() { is_true "${argc_go_modules:-false}"; }
is_temp_files_enabled() { is_true "${argc_temp_files:-false}"; }
is_test_files_enabled() { is_true "${argc_test_files:-false}"; }
is_all_enabled() { is_true "${argc_all:-true}"; }
is_dry_run_enabled() { is_true "${argc_dry_run:-false}"; }
is_command_available() { command -v "${1:-}" >/dev/null 2>&1; }

main() {
    run_artifacts_operation
}

eval "$(argc --argc-eval "$0" "$@")"
