#!/usr/bin/env bash

# Cisco WNC Lint Code - Core Functions
# Core business logic for code linting operations

# Source common predicates
source "$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)/../common/common.sh"

validate_lint_environment() {
    local project_root="$1"

    # Check if golangci-lint is available
    if ! has_golangci_lint; then
        format_lint_error "golangci-lint is not installed"
        format_lint_info "Install it with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"
        return 1
    fi

    # Validate project directory
    if ! validate_project_directory "$project_root"; then
        format_lint_error "Invalid project directory: $project_root"
        return 1
    fi

    # Check if it's a Go module
    if [[ ! -f "$project_root/go.mod" ]]; then
        format_lint_warning "No go.mod found in $project_root"
        format_lint_info "This may not be a Go module directory"
    fi

    return 0
}

prepare_lint_arguments() {
    local -n args_ref="$1"

    # Add auto-fix if enabled
    if is_auto_fix_enabled; then
        args_ref+=("--fix")
        is_verbose_enabled && format_lint_info "Auto-fix enabled"
    fi

    # Add custom config if specified
    if [[ -n "${argc_config:-}" ]]; then
        if [[ ! -f "${argc_config}" ]]; then
            format_lint_error "Config file not found: ${argc_config}"
            return 1
        fi
        args_ref+=("--config" "${argc_config}")
        is_verbose_enabled && format_lint_info "Using config: ${argc_config}"
    fi

    return 0
}

execute_lint_command() {
    local project_root="$1"
    local -a lint_args=("${@:2}")

    # Change to project directory
    local original_pwd="$PWD"
    cd "$project_root" || {
        format_lint_error "Failed to change to project directory: $project_root"
        return 1
    }

    is_verbose_enabled && format_lint_info "Running golangci-lint in: $project_root"
    is_verbose_enabled && format_lint_info "Arguments: ${lint_args[*]:-none}"

    # Run golangci-lint
    local exit_code=0
    if [[ ${#lint_args[@]} -gt 0 ]]; then
        golangci-lint run "${lint_args[@]}" ./... || exit_code=$?
    else
        golangci-lint run ./... || exit_code=$?
    fi

    # Return to original directory
    cd "$original_pwd" || {
        format_lint_warning "Failed to return to original directory"
    }

    return "$exit_code"
}

run_lint_operation() {
    local project_root="${argc_project:-.}"

    # Show banner and info
    show_lint_banner
    show_lint_verbose_info

    # Validate environment
    if ! validate_lint_environment "$project_root"; then
        return 1
    fi

    # Prepare arguments
    local lint_args=()
    if ! prepare_lint_arguments lint_args; then
        return 1
    fi

    # Execute linting
    format_lint_info "Starting code linting..."
    local exit_code=0
    execute_lint_command "$project_root" "${lint_args[@]}" || exit_code=$?

    # Display results
    display_lint_results "$exit_code" "$project_root"

    return "$exit_code"
}
