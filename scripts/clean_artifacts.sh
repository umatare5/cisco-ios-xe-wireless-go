#!/usr/bin/env bash

# @meta version 1.0.0
# @meta author "@umatare5"
# @describe Cisco WNC Artifacts Cleanup - Clean build artifacts, temporary files, and caches

# @option -p --project <DIR>          Project root directory [default: .]
# @flag   -v --verbose                Enable verbose output
# @flag   -f --force                  Force removal without confirmation
# @flag      --go-cache               Clean Go build cache
# @flag      --go-modules             Clean Go module cache
# @flag      --temp-files             Clean temporary files (./tmp)
# @flag      --test-files             Clean test artifacts (.test binaries, coverage files)
# @flag      --all                    Clean all artifacts [default: true]
# @flag      --dry-run                Show what would be cleaned without actually cleaning
# @flag      --no-color               Disable colored output

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
readonly SCRIPT_DIR

# Source bootstrap library
source "${SCRIPT_DIR}/lib/bootstrap.sh"

# Initialize libraries and load artifacts module
init_wnc_libraries "${SCRIPT_DIR}/lib/clean_artifacts"

# Load per-entry library for artifacts cleanup

# Validate required CLI tools
validate_required_cli_tools "minimal"

# Entrypoint: run cleanup operation
main() {
    local project_root="${argc_project:-.}"
    local go_cache="${argc_go_cache:-false}"
    local go_modules="${argc_go_modules:-false}"
    local temp_files="${argc_temp_files:-false}"
    local test_files="${argc_test_files:-false}"
    local all="${argc_all:-true}"
    local dry_run="${argc_dry_run:-false}"
    local verbose="${argc_verbose:-false}"

    run_artifacts_operation \
        "$project_root" "$go_cache" "$go_modules" \
        "$temp_files" "$test_files" "$all" "$dry_run" "$verbose"
}

eval "$(argc --argc-eval "$0" "$@")"
