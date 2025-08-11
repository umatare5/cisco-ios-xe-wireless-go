#!/usr/bin/env bash

# Cisco WNC Validation - Output functions for pre-commit checks
# Provides functions to display messages with colored output for pre-commit validation

# Display pre-commit validation banner with appropriate formatting
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

# Get current validation mode description
_get_validation_mode() {
    if is_no_color_enabled; then
        printf '%s\n' "no-color"
        return
    fi

    printf '%s\n' "color"
}

# Display current validation status and branch information
show_validation_status() {
    local current_branch="$1"

    show_pre_commit_banner
    printf '\n'
    info "Current branch: $current_branch"

    if is_verbose_enabled; then
        local mode
        mode=$(_get_validation_mode)
        info "Validation mode: $mode"
        info "Working directory: $(pwd)"
    fi
}

# Display detailed error message for main branch commit attempts
show_main_branch_error() {
    local branch_name="$1"

    error "Direct commits to the '$branch_name' branch are not allowed"
    printf '\n'
    warn "To maintain code quality and enable proper review processes,"
    warn "all changes must go through feature branches."
    printf '\n'

    # Use help functions for guidance
    show_workflow_guidance
    printf '\n'
    show_emergency_bypass
    printf '\n'
}

# Display validation success message with branch information
show_validation_success() {
    local branch_name="$1"

    success "Pre-commit validation passed"
    info "Proceeding with commit on branch '$branch_name'"

    if is_verbose_enabled; then
        info "All validation checks completed successfully"
    fi
}

# Display warning message when no staged changes are found
show_no_staged_changes() {
    warn "No staged changes found"
    info "Use 'git add <files>' to stage changes before committing"

    if is_verbose_enabled; then
        info "To see current status: git status"
        info "To see unstaged changes: git diff"
    fi
}
