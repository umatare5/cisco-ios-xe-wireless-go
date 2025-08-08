#!/usr/bin/env bash
# @meta version 1.0.0
# @meta author "@umatare5"
# @describe Cisco WNC Artifacts Cleanup - Clean build artifacts, temporary files, and caches

# @option -p --project <DIR>           Project root directory [default: .]
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

# Initialize WNC libraries with artifacts module
init_wnc_libraries "${SCRIPT_DIR}" "${SCRIPT_DIR}/lib/artifacts"

# Validate required CLI tools before proceeding
validate_required_cli_tools "minimal"

# Predicate functions centralized in lib/core/predicates.sh

main() {
    run_artifacts_operation
}

eval "$(argc --argc-eval "$0" "$@")"
