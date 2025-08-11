#!/usr/bin/env bash

# Dependencies Core Functions
# Provides core business logic for managing Go module dependencies

show_dependencies_banner() {
    wnc_banner_dependencies
}

validate_dependencies_environment() {
    local project_root="$1"

    # Validate project directory
    if ! is_valid_directory "$project_root"; then
        error "Invalid project directory: $project_root"
        return 1
    fi

    # Check for go.mod
    if [[ ! -f "$project_root/go.mod" ]]; then
        error "No go.mod found in $project_root"
        info "Run 'go mod init' to initialize a Go module"
        return 1
    fi

    # Check Go installation
    if ! is_command_available "go"; then
        error "Go is not installed or not in PATH"
        return 1
    fi

    # Display Go version
    local go_version
    go_version="$(go version 2>/dev/null | cut -d' ' -f3)"
    info "Using Go version: $go_version"

    return 0
}

# Execute the clean operation for dependencies
execute_dependencies_clean() {
    local clean="$1"

    if [[ "$clean" != "true" ]]; then
        return
    fi

    format_step_message "1" "Cleaning module cache..."

    if go clean -modcache 2>/dev/null; then
        success "Module cache cleaned"
        printf '\n'
        return
    fi

    warn "Failed to clean module cache (continuing anyway)"
    printf '\n'
}

# Check if download-only mode is enabled
is_download_only_enabled() {
    [[ "$1" == "true" ]]
}

execute_dependencies_download() {
    local download_only="$1"
    local verbose="$2"

    format_step_message "2" "Downloading dependencies..."

    if is_download_only_enabled "$download_only"; then
        # Download only, don't build
        if go mod download 2>/dev/null; then
            success "Dependencies downloaded"
            printf '\n'
            return
        fi

        error "Failed to download dependencies"
        return 1
    fi

    # Download and install
    local tidy_args=()

    if [[ "$verbose" == "true" ]]; then
        tidy_args+=("-v")
    fi

    if go mod tidy "${tidy_args[@]+${tidy_args[@]}}" 2>/dev/null; then
        success "Dependencies tidied"
    else
        error "Failed to tidy dependencies"
        return 1
    fi

    if go mod download "${tidy_args[@]+${tidy_args[@]}}" 2>/dev/null; then
        success "Dependencies downloaded"
        printf '\n'
        return
    fi

    error "Failed to download dependencies"
    return 1
}

# Check if update mode is enabled
is_update_enabled() {
    [[ "$1" == "true" ]]
}

# Check if verbose mode is enabled for dependencies
is_dependencies_verbose_enabled() {
    [[ "$1" == "true" ]]
}

execute_dependencies_update() {
    local update="$1"
    local verbose="$2"

    # Skip if update is not requested
    if ! is_update_enabled "$update"; then
        return
    fi

    format_step_message "3" "Updating dependencies..."

    local update_args=()
    if is_dependencies_verbose_enabled "$verbose"; then
        update_args+=("-v")
    fi

    # Update all dependencies to latest versions
    if go get -u "${update_args[@]}" ./... 2>/dev/null; then
        success "Dependencies updated"
    else
        error "Failed to update dependencies"
        return 1
    fi

    # Re-tidy after updates to ensure consistency
    if go mod tidy 2>/dev/null; then
        success "Dependencies re-tidied after update"
        printf '\n'
        return
    fi

    warn "Failed to tidy after update"
    printf '\n'
}

# Check if verification is enabled
is_verification_enabled() {
    [[ "$1" == "true" ]]
}

execute_dependencies_verify() {
    local verify="$1"

    # Skip verification if not requested
    if ! is_verification_enabled "$verify"; then
        return
    fi

    format_step_message "4" "Verifying dependencies..."

    # Verify checksums and integrity of dependencies
    if go mod verify 2>/dev/null; then
        success "Dependencies verified"
        printf '\n'
        return
    fi

    error "Dependency verification failed"
    return 1
}

display_dependencies_summary() {
    format_step_message "✓" "Dependencies management completed"

    # Show direct dependencies count
    local direct_deps go_list_output
    # Count non-empty direct dependencies; use wc -l to avoid non-zero exit codes
    go_list_output="$(go list -mod=readonly -m -f '{{if not .Indirect}}{{.Path}}{{end}}' all 2>/dev/null)"
    direct_deps="$(printf '%s' "$go_list_output" | sed '/^$/d' | wc -l | tr -d ' ')"

    # Show total dependencies count
    local total_deps
    total_deps="$(go list -mod=readonly -m all 2>/dev/null | tail -n +2 | wc -l | tr -d ' ')"

    info "Direct dependencies: $direct_deps"
    info "Total dependencies: $total_deps"

    if [[ "$DEPENDENCIES_VERBOSE" == "true" ]]; then
        printf '\n'
        format_step_message "ℹ" "Direct dependencies:"
        go list -mod=readonly -m -f '{{if not .Indirect}}  {{.Path}} {{.Version}}{{end}}' all 2>/dev/null
    fi

    return 0
}
