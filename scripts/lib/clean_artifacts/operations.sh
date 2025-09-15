#!/usr/bin/env bash

# Cisco WNC Artifacts Cleanup - Helper Functions
# Provides functions to clean up Go build artifacts, caches, and temporary files

# Collect flags/args into readonly settings
prepare_artifacts_arguments() {
    local project_root="$1"
    local go_cache="$2"
    local go_modules="$3"
    local temp_files="$4"
    local test_files="$5"
    local all="$6"
    local dry_run="$7"
    local verbose="$8"

    # Only set readonly variables if not already set
    if [[ -z "${ARTIFACTS_PROJECT_ROOT:-}" ]]; then
        readonly ARTIFACTS_PROJECT_ROOT="$project_root"
    fi
    if [[ -z "${ARTIFACTS_GO_CACHE:-}" ]]; then
        readonly ARTIFACTS_GO_CACHE="$go_cache"
    fi
    if [[ -z "${ARTIFACTS_GO_MODULES:-}" ]]; then
        readonly ARTIFACTS_GO_MODULES="$go_modules"
    fi
    if [[ -z "${ARTIFACTS_TEMP_FILES:-}" ]]; then
        readonly ARTIFACTS_TEMP_FILES="$temp_files"
    fi
    if [[ -z "${ARTIFACTS_TEST_FILES:-}" ]]; then
        readonly ARTIFACTS_TEST_FILES="$test_files"
    fi
    if [[ -z "${ARTIFACTS_ALL:-}" ]]; then
        readonly ARTIFACTS_ALL="$all"
    fi
    if [[ -z "${ARTIFACTS_DRY_RUN:-}" ]]; then
        readonly ARTIFACTS_DRY_RUN="$dry_run"
    fi
    if [[ -z "${ARTIFACTS_VERBOSE:-}" ]]; then
        readonly ARTIFACTS_VERBOSE="$verbose"
    fi

    export ARTIFACTS_PROJECT_ROOT
    export ARTIFACTS_VERBOSE
}

# Execute cleanup operations based on flags
execute_cleanup_operations() {
    if [[ "$ARTIFACTS_ALL" == "true" || "$ARTIFACTS_GO_CACHE" == "true" ]]; then
        clean_go_cache || return 1
    fi

    if [[ "$ARTIFACTS_ALL" == "true" || "$ARTIFACTS_GO_MODULES" == "true" ]]; then
        clean_go_modules || return 1
    fi

    if [[ "$ARTIFACTS_ALL" == "true" || "$ARTIFACTS_TEMP_FILES" == "true" ]]; then
        clean_temp_files || return 1
    fi

    if [[ "$ARTIFACTS_ALL" == "true" || "$ARTIFACTS_TEST_FILES" == "true" ]]; then
        clean_test_files || return 1
    fi

    if [[ "$ARTIFACTS_ALL" == "true" ]]; then
        clean_backup_files || return 1
    fi
}

# Display completion message based on dry-run flag
show_completion_message() {
    if [[ "$ARTIFACTS_DRY_RUN" == "true" ]]; then
        format_step_message "✓" "Dry-run completed (no files were actually removed)"
        return
    fi

    format_step_message "✓" "Artifacts cleanup completed successfully"
}
