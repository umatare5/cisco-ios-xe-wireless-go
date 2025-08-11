#!/usr/bin/env bash

# Cisco WNC Linting - Core Operations
# Provides core functionality for code linting operations

run_lint_operation() {
    local project_root="$1"

    show_lint_banner
    show_lint_verbose_info "$project_root"

    if ! validate_lint_environment "$project_root"; then
        return 1
    fi

    if ! prepare_lint_arguments; then
        return 1
    fi

    format_lint_info "Starting code linting..."
    local exit_code=0
    execute_lint_command "$project_root" || exit_code=$?

    display_lint_results "$exit_code" "$project_root"
    return "$exit_code"
}
