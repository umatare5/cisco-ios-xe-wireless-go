#!/usr/bin/env bash

# Artifacts Cleanup Core Functions

set -euo pipefail

# Source bootstrap library
LIB_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
source "${LIB_ROOT}/bootstrap.sh"
init_wnc_basic

show_artifacts_banner() {
    format_artifacts_banner
}

validate_artifacts_environment() {
    local project_root="${argc_project:-.}"

    # Validate project directory
    if ! is_valid_directory "$project_root"; then
        format_artifacts_error "Invalid project directory: $project_root"
        return 1
    fi

    # Check Go installation
    if ! is_command_available "go"; then
        format_artifacts_error "Go is not installed or not in PATH"
        return 1
    fi

    return 0
}

prepare_artifacts_arguments() {
    readonly ARTIFACTS_PROJECT_ROOT="${argc_project:-.}"
    readonly ARTIFACTS_GO_CACHE="${argc_go_cache:-false}"
    readonly ARTIFACTS_GO_MODULES="${argc_go_modules:-false}"
    readonly ARTIFACTS_TEMP_FILES="${argc_temp_files:-false}"
    readonly ARTIFACTS_TEST_FILES="${argc_test_files:-false}"
    readonly ARTIFACTS_ALL="${argc_all:-true}"
    readonly ARTIFACTS_DRY_RUN="${argc_dry_run:-false}"
    readonly ARTIFACTS_VERBOSE="${argc_verbose:-false}"

    # If specific flags are set, disable --all
    if [[ "$ARTIFACTS_GO_CACHE" == "true" || "$ARTIFACTS_GO_MODULES" == "true" || \
          "$ARTIFACTS_TEMP_FILES" == "true" || "$ARTIFACTS_TEST_FILES" == "true" ]]; then
        readonly ARTIFACTS_ALL="false"
    fi
}

get_directory_size() {
    local dir="$1"
    if [[ -d "$dir" ]]; then
        du -sh "$dir" 2>/dev/null | cut -f1 || echo "unknown"
    else
        echo "0B"
    fi
}

clean_go_cache() {
    format_step_message "1" "Cleaning Go build cache..."

    local go_cache
    go_cache="$(go env GOCACHE 2>/dev/null || echo "")"

    if [[ -z "$go_cache" ]]; then
        format_artifacts_warning "Could not determine Go cache directory"
        return 0
    fi

    local cache_size
    cache_size="$(get_directory_size "$go_cache")"

    if [[ "$ARTIFACTS_DRY_RUN" == "true" ]]; then
        format_dry_run_item "Clean Go cache" "$go_cache" "$cache_size"
    else
        if go clean -cache 2>/dev/null; then
            format_artifacts_success "Go build cache cleaned ($cache_size freed)"
        else
            format_artifacts_error "Failed to clean Go build cache"
            return 1
        fi
    fi
}

clean_go_modules() {
    format_step_message "2" "Cleaning Go module cache..."

    local go_modcache
    go_modcache="$(go env GOMODCACHE 2>/dev/null || echo "")"

    if [[ -z "$go_modcache" ]]; then
        format_artifacts_warning "Could not determine Go module cache directory"
        return 0
    fi

    local cache_size
    cache_size="$(get_directory_size "$go_modcache")"

    if [[ "$ARTIFACTS_DRY_RUN" == "true" ]]; then
        format_dry_run_item "Clean Go module cache" "$go_modcache" "$cache_size"
    else
        if go clean -modcache 2>/dev/null; then
            format_artifacts_success "Go module cache cleaned ($cache_size freed)"
        else
            format_artifacts_error "Failed to clean Go module cache"
            return 1
        fi
    fi
}

clean_temp_files() {
    format_step_message "3" "Cleaning temporary files..."

    local temp_dir="$ARTIFACTS_PROJECT_ROOT/tmp"

    if [[ ! -d "$temp_dir" ]]; then
        format_artifacts_info "No temporary directory found: $temp_dir"
        return 0
    fi

    local temp_size
    temp_size="$(get_directory_size "$temp_dir")"

    if [[ "$ARTIFACTS_DRY_RUN" == "true" ]]; then
        format_dry_run_item "Remove directory" "$temp_dir" "$temp_size"
    else
        if rm -rf "$temp_dir" 2>/dev/null; then
            format_artifacts_success "Temporary files cleaned ($temp_size freed)"
        else
            format_artifacts_error "Failed to clean temporary files"
            return 1
        fi
    fi
}

clean_test_files() {
    format_step_message "4" "Cleaning test artifacts..."

    local cleaned_count=0

    # Find and clean .test binaries
    while IFS= read -r -d '' test_binary; do
        local file_size
        file_size="$(du -sh "$test_binary" 2>/dev/null | cut -f1 || echo "unknown")"

        if [[ "$ARTIFACTS_DRY_RUN" == "true" ]]; then
            format_dry_run_item "Remove test binary" "$test_binary" "$file_size"
        else
            if rm -f "$test_binary" 2>/dev/null; then
                format_file_removed "$test_binary" "$file_size"
                ((cleaned_count++))
            fi
        fi
    done < <(find "$ARTIFACTS_PROJECT_ROOT" -name "*.test" -type f -print0 2>/dev/null || true)

    # Find and clean coverage files
    while IFS= read -r -d '' coverage_file; do
        local file_size
        file_size="$(du -sh "$coverage_file" 2>/dev/null | cut -f1 || echo "unknown")"

        if [[ "$ARTIFACTS_DRY_RUN" == "true" ]]; then
            format_dry_run_item "Remove coverage file" "$coverage_file" "$file_size"
        else
            if rm -f "$coverage_file" 2>/dev/null; then
                format_file_removed "$coverage_file" "$file_size"
                ((cleaned_count++))
            fi
        fi
    done < <(find "$ARTIFACTS_PROJECT_ROOT" -name "*.out" -o -name "*.html" -type f -print0 2>/dev/null || true)

    if [[ "$ARTIFACTS_DRY_RUN" != "true" ]]; then
        if [[ $cleaned_count -gt 0 ]]; then
            format_artifacts_success "Test artifacts cleaned ($cleaned_count files)"
        else
            format_artifacts_info "No test artifacts found to clean"
        fi
    fi
}

display_cache_info() {
    format_step_message "ℹ" "Cache information:"

    local go_cache go_modcache
    go_cache="$(go env GOCACHE 2>/dev/null || echo "")"
    go_modcache="$(go env GOMODCACHE 2>/dev/null || echo "")"

    if [[ -n "$go_cache" ]]; then
        format_cache_info "Go Build Cache" "$go_cache" "$(is_valid_directory "$go_cache" && echo "true" || echo "false")"
    fi

    if [[ -n "$go_modcache" ]]; then
        format_cache_info "Go Module Cache" "$go_modcache" "$(is_valid_directory "$go_modcache" && echo "true" || echo "false")"
    fi

    echo ""
}

run_artifacts_operation() {
    show_artifacts_banner

    if ! validate_artifacts_environment; then
        return 1
    fi

    prepare_artifacts_arguments

    cd "$ARTIFACTS_PROJECT_ROOT" || {
        format_artifacts_error "Failed to change to project directory: $ARTIFACTS_PROJECT_ROOT"
        return 1
    }

    if [[ "$ARTIFACTS_VERBOSE" == "true" ]]; then
        display_cache_info
    fi

    # Execute cleanup operations based on flags
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

    if [[ "$ARTIFACTS_DRY_RUN" == "true" ]]; then
        format_step_message "✓" "Dry-run completed (no files were actually removed)"
    else
        format_step_message "✓" "Artifacts cleanup completed successfully"
    fi
}
