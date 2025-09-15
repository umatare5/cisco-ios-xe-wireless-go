#!/usr/bin/env bash

# Artifacts Cleanup Core Functions
# Provides core business logic for artifacts cleanup operations

show_artifacts_banner() {
    wnc_banner_artifacts
}

# Clean up backup files from the repository
clean_backup_files() {
    find . -name "*.tmp*" -type f -print0 | xargs -0 rm -f
    find . -name "*.temp*" -type f -print0 | xargs -0 rm -f
    find . -name "*.bak*" -type f -print0 | xargs -0 rm -f
    find . -name "*.backup*" -type f -print0 | xargs -0 rm -f
    return 0
}

validate_artifacts_environment() {
    local project_root="$1"

    # Validate project directory
    if ! is_valid_directory "$project_root"; then
        error "Invalid project directory: $project_root"
        return 1
    fi

    # Check Go installation
    if ! is_command_available "go"; then
        error "Go is not installed or not in PATH"
        return 1
    fi

    return 0
}

# Check if directory exists for size calculation
has_directory() {
    local dir="$1"
    [[ -d "$dir" ]]
}

# Get size of directory or return default for missing directory
get_directory_size() {
    local dir="$1"

    if has_directory "$dir"; then
        du -sh "$dir" 2>/dev/null | cut -f1 || printf '%s\n' "unknown"
        return
    fi

    printf '%s\n' "0B"
}

# Check if dry run mode is enabled
is_dry_run_enabled() {
    [[ "$ARTIFACTS_DRY_RUN" == "true" ]]
}

# Check if Go cache directory is available
has_go_cache_directory() {
    local go_cache="$1"
    [[ -n "$go_cache" ]]
}

# Execute cache cleaning operation
_execute_cache_clean() {
    local cache_size="$1"

    if go clean -cache 2>/dev/null; then
        success "Go build cache cleaned ($cache_size freed)"
        return
    fi

    error "Failed to clean Go build cache"
    return 1
}

# Clean Go build cache with dry run support
clean_go_cache() {
    format_step_message "1" "Cleaning Go build cache..."

    local go_cache
    go_cache="$(go env GOCACHE 2>/dev/null || printf '')"

    if ! has_go_cache_directory "$go_cache"; then
        warn "Could not determine Go cache directory"
        return 0
    fi

    local cache_size
    cache_size="$(get_directory_size "$go_cache")"

    if is_dry_run_enabled; then
        format_dry_run_item "Clean Go cache" "$go_cache" "$cache_size"
        return
    fi

    _execute_cache_clean "$cache_size"
}

# Check if Go module cache directory is available
has_go_modcache_directory() {
    local go_modcache="$1"
    [[ -n "$go_modcache" ]]
}

# Execute module cache cleaning operation
_execute_modcache_clean() {
    local cache_size="$1"

    if go clean -modcache 2>/dev/null; then
        success "Go module cache cleaned ($cache_size freed)"
        return
    fi

    error "Failed to clean Go module cache"
    return 1
}

# Clean Go module cache with dry run support
clean_go_modules() {
    format_step_message "2" "Cleaning Go module cache..."

    local go_modcache
    go_modcache="$(go env GOMODCACHE 2>/dev/null || printf '')"

    if ! has_go_modcache_directory "$go_modcache"; then
        warn "Could not determine Go module cache directory"
        return 0
    fi

    local cache_size
    cache_size="$(get_directory_size "$go_modcache")"

    if is_dry_run_enabled; then
        format_dry_run_item "Clean Go module cache" "$go_modcache" "$cache_size"
        return
    fi

    _execute_modcache_clean "$cache_size"
}

clean_temp_files() {
    format_step_message "3" "Cleaning temporary files..."

    local temp_dir="$ARTIFACTS_PROJECT_ROOT/tmp"

    if [[ ! -d "$temp_dir" ]]; then
        info "No temporary directory found: $temp_dir"
        return 0
    fi

    local temp_size
    temp_size="$(get_directory_size "$temp_dir")"

    if [[ "$ARTIFACTS_DRY_RUN" == "true" ]]; then
        format_dry_run_item "Remove directory" "$temp_dir" "$temp_size"
        return
    fi

    if rm -rf "$temp_dir" 2>/dev/null; then
        success "Temporary files cleaned ($temp_size freed)"
        return
    fi

    error "Failed to clean temporary files"
    return 1
}

# Clean files matching a pattern with progress reporting
_clean_files_by_pattern() {
    local pattern="$1"
    local description="$2"
    local cleaned_count_var="$3"

    while IFS= read -r -d '' file_path; do
        local file_size
        file_size="$(du -sh "$file_path" 2>/dev/null | cut -f1 || printf 'unknown')"

        if [[ "$ARTIFACTS_DRY_RUN" == "true" ]]; then
            format_dry_run_item "Remove $description" "$file_path" "$file_size"
            continue
        fi

        if rm -f "$file_path" 2>/dev/null; then
            format_file_removed "$file_path" "$file_size"
            eval "((${cleaned_count_var}++))"
        fi
    done < <(find "$ARTIFACTS_PROJECT_ROOT" "$pattern" -type f -print0 2>/dev/null || true)
}

# Clean all test artifacts (binaries and coverage files)
clean_test_files() {
    format_step_message "4" "Cleaning test artifacts..."

    local cleaned_count=0

    # Find and clean .test binaries
    _clean_files_by_pattern "-name '*.test'" "test binary" "cleaned_count"

    # Find and clean coverage files
    _clean_files_by_pattern "-name '*.out' -o -name '*.html'" "coverage file" "cleaned_count"

    if [[ "$ARTIFACTS_DRY_RUN" == "true" ]]; then
        return
    fi

    if [[ $cleaned_count -gt 0 ]]; then
        success "Test artifacts cleaned ($cleaned_count files)"
        return
    fi

    info "No test artifacts found to clean"
}

display_cache_info() {
    format_step_message "ℹ" "Cache information:"

    local go_cache go_modcache
    go_cache="$(go env GOCACHE 2>/dev/null || printf '')"
    go_modcache="$(go env GOMODCACHE 2>/dev/null || printf '')"

    if [[ -n "$go_cache" ]]; then
    format_cache_info "Go Build Cache" "$go_cache" "$(is_valid_directory "$go_cache" && printf true || printf false)"
    fi

    if [[ -n "$go_modcache" ]]; then
        local go_modcache_valid
        go_modcache_valid="$(is_valid_directory "$go_modcache" && printf true || printf false)"
        format_cache_info "Go Module Cache" "$go_modcache" "$go_modcache_valid"
    fi

    printf '\n'
}
