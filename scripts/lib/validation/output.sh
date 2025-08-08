#!/usr/bin/env bash

# Cisco WNC Pre-commit Validation - Output Functions
# Output formatting and display functions

set -euo pipefail

# Helper functions for colored output
show_info() {
    if ! is_no_color_enabled; then
        printf "\033[36mℹ Info:\033[0m %s\n" "$*"
    else
        printf "Info: %s\n" "$*"
    fi
}

show_warning() {
    if ! is_no_color_enabled; then
        printf "\033[33m⚠ Warning:\033[0m %s\n" "$*" >&2
    else
        printf "Warning: %s\n" "$*" >&2
    fi
}

# Banner display
show_pre_commit_banner() {
    # If unified banner helper exists, use it; otherwise fallback.
    if command -v wnc_banner_pre_commit >/dev/null 2>&1; then
        wnc_banner_pre_commit
        return
    fi
    cat <<'EOF'
====================================
 Pre-commit Validation
 Branch Protection System
====================================
EOF
}

# Show validation status
show_validation_status() {
    local current_branch="$1"

    show_pre_commit_banner
    echo
    printf "\033[36mℹ Info:\033[0m Current branch: %s\n" "$current_branch"

    if is_verbose_enabled; then
        printf "\033[36mℹ Info:\033[0m Validation mode: $(is_no_color_enabled && echo "no-color" || echo "color")\n"
        printf "\033[36mℹ Info:\033[0m Working directory: $(pwd)\n"
    fi
}

# Show main branch error with detailed guidance
show_main_branch_error() {
    local branch_name="$1"

    printf "\033[31m✗ Error:\033[0m Direct commits to the '%s' branch are not allowed\n" "$branch_name" >&2
    echo
    printf "\033[33m⚠ Warning:\033[0m To maintain code quality and enable proper review processes,\n" >&2
    printf "\033[33m⚠ Warning:\033[0m all changes must go through feature branches.\n" >&2
    echo

    # Use help functions for guidance
    show_workflow_guidance
    echo
    show_emergency_bypass
    echo
}

# NOTE: Duplicate show_no_staged_changes removed; using the concise version defined later.

# Show success message
show_validation_success() {
    local branch_name="$1"

    printf "\033[32m✓ Success:\033[0m Pre-commit validation passed\n"
    printf "\033[36mℹ Info:\033[0m Proceeding with commit on branch '%s'\n" "$branch_name"

    if is_verbose_enabled; then
        printf "\033[36mℹ Info:\033[0m All validation checks completed successfully\n"
    fi
}

# Show no staged changes warning
show_no_staged_changes() {
    show_warning "No staged changes found"
    show_info "Use 'git add <files>' to stage changes before committing"

    if is_verbose_enabled; then
        show_info "To see current status: git status"
        show_info "To see unstaged changes: git diff"
    fi
}
