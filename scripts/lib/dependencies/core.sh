#!/usr/bin/env bash

# Dependencies Core Functions

set -euo pipefail

# Source bootstrap library
LIB_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
source "${LIB_ROOT}/bootstrap.sh"
init_wnc_basic

show_dependencies_banner() {
    format_dependencies_banner
}

validate_dependencies_environment() {
    local project_root="${argc_project:-.}"

    # Validate project directory
    if ! is_valid_directory "$project_root"; then
        format_dependencies_error "Invalid project directory: $project_root"
        return 1
    fi

    # Check for go.mod
    if [[ ! -f "$project_root/go.mod" ]]; then
        format_dependencies_error "No go.mod found in $project_root"
        format_dependencies_info "Run 'go mod init' to initialize a Go module"
        return 1
    fi

    # Check Go installation
    if ! is_command_available "go"; then
        format_dependencies_error "Go is not installed or not in PATH"
        return 1
    fi

    # Display Go version
    local go_version
    go_version="$(go version 2>/dev/null | cut -d' ' -f3)"
    format_dependencies_info "Using Go version: $go_version"

    return 0
}

prepare_dependencies_arguments() {
    readonly DEPENDENCIES_PROJECT_ROOT="${argc_project:-.}"
    readonly DEPENDENCIES_CLEAN="${argc_clean:-false}"
    readonly DEPENDENCIES_UPDATE="${argc_update:-false}"
    readonly DEPENDENCIES_DOWNLOAD_ONLY="${argc_download_only:-false}"
    readonly DEPENDENCIES_VERIFY="${argc_verify:-false}"
    readonly DEPENDENCIES_VERBOSE="${argc_verbose:-false}"
}

execute_dependencies_clean() {
    if [[ "$DEPENDENCIES_CLEAN" == "true" ]]; then
        format_step_message "1" "Cleaning module cache..."

        if go clean -modcache 2>/dev/null; then
            format_dependencies_success "Module cache cleaned"
        else
            format_dependencies_warning "Failed to clean module cache (continuing anyway)"
        fi
        echo ""
    fi
}

execute_dependencies_download() {
    format_step_message "2" "Downloading dependencies..."

    if [[ "$DEPENDENCIES_DOWNLOAD_ONLY" == "true" ]]; then
        # Download only, don't build
        if go mod download 2>/dev/null; then
            format_dependencies_success "Dependencies downloaded"
        else
            format_dependencies_error "Failed to download dependencies"
            return 1
        fi
    else
        # Download and install
        local tidy_args=()

        if [[ "$DEPENDENCIES_VERBOSE" == "true" ]]; then
            tidy_args+=("-v")
        fi

        if go mod tidy "${tidy_args[@]+${tidy_args[@]}}" 2>/dev/null; then
            format_dependencies_success "Dependencies tidied"
        else
            format_dependencies_error "Failed to tidy dependencies"
            return 1
        fi

        if go mod download "${tidy_args[@]+${tidy_args[@]}}" 2>/dev/null; then
            format_dependencies_success "Dependencies downloaded"
        else
            format_dependencies_error "Failed to download dependencies"
            return 1
        fi
    fi

    echo ""
}

execute_dependencies_update() {
    if [[ "$DEPENDENCIES_UPDATE" == "true" ]]; then
        format_step_message "3" "Updating dependencies..."

        local update_args=()

        if [[ "$DEPENDENCIES_VERBOSE" == "true" ]]; then
            update_args+=("-v")
        fi

        if go get -u "${update_args[@]}" ./... 2>/dev/null; then
            format_dependencies_success "Dependencies updated"
        else
            format_dependencies_error "Failed to update dependencies"
            return 1
        fi

        # Re-tidy after updates
        if go mod tidy 2>/dev/null; then
            format_dependencies_success "Dependencies re-tidied after update"
        else
            format_dependencies_warning "Failed to tidy after update"
        fi

        echo ""
    fi
}

execute_dependencies_verify() {
    if [[ "$DEPENDENCIES_VERIFY" == "true" ]]; then
        format_step_message "4" "Verifying dependencies..."

        if go mod verify 2>/dev/null; then
            format_dependencies_success "Dependencies verified"
        else
            format_dependencies_error "Dependency verification failed"
            return 1
        fi

        echo ""
    fi
}

display_dependencies_summary() {
    format_step_message "✓" "Dependencies management completed"

    # Show direct dependencies count
    local direct_deps
    direct_deps="$(go list -m -f '{{if not .Indirect}}{{.Path}}{{end}}' all 2>/dev/null | grep -c . || echo "0")"

    # Show total dependencies count
    local total_deps
    total_deps="$(go list -m all 2>/dev/null | tail -n +2 | wc -l | tr -d ' ')"

    format_dependencies_info "Direct dependencies: $direct_deps"
    format_dependencies_info "Total dependencies: $total_deps"

    if [[ "$DEPENDENCIES_VERBOSE" == "true" ]]; then
        echo ""
        format_step_message "ℹ" "Direct dependencies:"
        go list -m -f '{{if not .Indirect}}  {{.Path}} {{.Version}}{{end}}' all 2>/dev/null
    fi
}

run_dependencies_operation() {
    show_dependencies_banner

    if ! validate_dependencies_environment; then
        return 1
    fi

    prepare_dependencies_arguments

    cd "$DEPENDENCIES_PROJECT_ROOT" || {
        format_dependencies_error "Failed to change to project directory: $DEPENDENCIES_PROJECT_ROOT"
        return 1
    }

    # Execute operations in sequence
    execute_dependencies_clean || return 1
    execute_dependencies_download || return 1
    execute_dependencies_update || return 1
    execute_dependencies_verify || return 1

    display_dependencies_summary
}
