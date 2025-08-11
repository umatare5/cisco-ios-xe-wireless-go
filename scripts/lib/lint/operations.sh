#!/usr/bin/env bash

# Cisco WNC Linting - Operations Functions
# Provides operational functions for linting processes

show_lint_verbose_info() {
    local project_root="$1"

    if is_verbose_enabled; then
        info "Linting project: ${project_root:-.}"
        if command -v golangci-lint >/dev/null 2>&1; then
            info "golangci-lint: $(golangci-lint version 2>/dev/null | head -n1)"
        fi
    fi
}

format_lint_info() {
    info "$*"
}

validate_lint_environment() {
    local project_root="${1:-.}"
    validate_go_module "$project_root" || return 1
    return 0
}

prepare_lint_arguments() {
    return 0
}

execute_lint_command() {
    local project_root="${1:-.}"
    run_lints "$project_root"
}

display_lint_results() {
    local exit_code="${1:-0}"
    local project_root="${2:-.}"

    if [[ "$exit_code" -eq 0 ]]; then
                success "Lint passed for $project_root"
        return
    fi

    error "Lint failed for $project_root"
}
