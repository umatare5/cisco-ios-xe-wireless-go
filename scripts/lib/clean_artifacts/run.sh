#!/usr/bin/env bash

# Cisco WNC Artifacts Cleanup - Core Operations
# Provides core functionality for artifacts cleanup operations

run_artifacts_operation() {
    local project_root="$1"
    local go_cache="$2"
    local go_modules="$3"
    local temp_files="$4"
    local test_files="$5"
    local all="$6"
    local dry_run="$7"
    local verbose="$8"

    # Display operation banner and validate environment
    show_artifacts_banner
    if ! validate_artifacts_environment "$project_root"; then
        return 1
    fi

    # Prepare cleanup arguments and configurations
    prepare_artifacts_arguments \
        "$project_root" "$go_cache" "$go_modules" "$temp_files" "$test_files" "$all" "$dry_run" "$verbose"

    # Change to project directory for cleanup operations
    if ! cd "$ARTIFACTS_PROJECT_ROOT"; then
        error "Failed to change to project directory: $ARTIFACTS_PROJECT_ROOT"
        return 1
    fi

    # Display cache information if verbose mode is enabled
    if [[ "$ARTIFACTS_VERBOSE" == "true" ]]; then
        display_cache_info
    fi

    # Execute cleanup operations and show completion status
    execute_cleanup_operations
    show_completion_message
}
