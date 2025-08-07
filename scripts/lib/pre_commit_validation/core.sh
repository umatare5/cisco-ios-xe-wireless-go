#!/usr/bin/env bash

# Cisco WNC Pre-commit Validation - Core Functions
# Core business logic for pre-commit validation

set -euo pipefail

# Get current branch name
get_current_branch() {
    git rev-parse --abbrev-ref HEAD 2>/dev/null || echo ""
}

# Check if we're in a Git repository
is_git_repository() {
    git rev-parse --git-dir >/dev/null 2>&1
}

# Check if the current branch is main or master
is_main_branch() {
    local current_branch="$1"
    [[ "$current_branch" == "main" ]] || [[ "$current_branch" == "master" ]]
}

# Check if there are any staged changes
has_staged_changes() {
    ! git diff --cached --quiet
}

# Validate Git environment
validate_git_environment() {
    # Verify we're in a git repository
    if ! is_git_repository; then
        printf "\033[31m✗ Error:\033[0m Not in a Git repository\n" >&2
        if is_verbose_enabled; then
            printf "\033[36mℹ Info:\033[0m Current directory: $(pwd)\n"
            printf "\033[36mℹ Info:\033[0m Try running this command from within a Git repository\n"
        fi
        return 1
    fi

    # Get current branch
    local current_branch
    current_branch=$(get_current_branch)

    if [[ -z "$current_branch" ]]; then
        printf "\033[31m✗ Error:\033[0m Unable to determine current branch\n" >&2
        if is_verbose_enabled; then
            printf "\033[36mℹ Info:\033[0m This might indicate a detached HEAD state\n"
            printf "\033[36mℹ Info:\033[0m Try: git checkout main && git checkout -b feature/your-branch\n"
        fi
        return 1
    fi

    # Store branch name for use by caller
    CURRENT_BRANCH="$current_branch"
    return 0
}

# Validate branch permissions
validate_branch_permissions() {
    local current_branch="$1"

    # Check if we're trying to commit directly to main/master
    if is_main_branch "$current_branch"; then
        show_main_branch_error "$current_branch"
        return 1
    fi

    return 0
}

# Validate staged changes
validate_staged_changes() {
    # Verify there are staged changes
    if ! has_staged_changes; then
        show_no_staged_changes
        return 1
    fi

    if is_verbose_enabled; then
        local staged_files
        staged_files=$(git diff --cached --name-only | wc -l | tr -d ' ')
        printf "\033[36mℹ Info:\033[0m Found %s staged file(s)\n" "$staged_files"
    fi

    return 0
}

# Main pre-commit validation logic
run_pre_commit_validation() {
    local current_branch=""

    # Validate Git environment and get current branch
    if ! validate_git_environment; then
        return 1
    fi

    current_branch="$CURRENT_BRANCH"

    # Show status
    show_validation_status "$current_branch"

    # Validate branch permissions
    if ! validate_branch_permissions "$current_branch"; then
        return 1
    fi

    # Validate staged changes
    if ! validate_staged_changes; then
        return 1
    fi

    # Show success
    show_validation_success "$current_branch"
    return 0
}
