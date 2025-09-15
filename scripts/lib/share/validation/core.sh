#!/usr/bin/env bash

# Cisco WNC Pre-commit Validation - Core Functions
# Provides core functions for validating Git pre-commit conditions

# Return current Git branch name or empty string if not available
get_current_branch() {
    git rev-parse --abbrev-ref HEAD 2>/dev/null || printf '%s\n' ""
}

# Check if current directory is within a Git repository
is_git_repository() {
    git rev-parse --git-dir >/dev/null 2>&1
}

# Determine if given branch is a protected main branch
is_main_branch() {
    local current_branch="$1"
    [[ "$current_branch" == "main" ]] || [[ "$current_branch" == "master" ]]
}

# Check if there are any files staged for commit
has_staged_changes() {
    ! git diff --cached --quiet
}

# Validate Git repository environment and set current branch
validate_git_environment() {
    # Verify we're in a git repository
    if ! is_git_repository; then
        error "Not in a Git repository"
        if is_verbose_enabled; then
            info "Current directory: $(pwd)"
            info "Try running this command from within a Git repository"
        fi
        return 1
    fi

    # Get current branch
    local current_branch
    current_branch=$(get_current_branch)

    if [[ -z "$current_branch" ]]; then
        error "Unable to determine current branch"
        if is_verbose_enabled; then
            info "This might indicate a detached HEAD state"
            info "Try: git checkout main && git checkout -b feature/your-branch"
        fi
        return 1
    fi

    # Store branch name for use by caller
    # Export current branch for use by other modules
    CURRENT_BRANCH="$current_branch"
    export CURRENT_BRANCH
    return 0
}

# Check if current branch allows direct commits
validate_branch_permissions() {
    local current_branch="$1"

    # Check if we're trying to commit directly to main/master
    if is_main_branch "$current_branch"; then
        show_main_branch_error "$current_branch"
        return 1
    fi

    return 0
}

# Verify staged changes exist and provide feedback
validate_staged_changes() {
    # Verify there are staged changes
    if ! has_staged_changes; then
        # Emit warning but do not fail overall validation when invoked via Makefile test target.
        # Rationale: 'make pre-commit-test' is often used as a dry-run; absence of staged
        # changes should not produce a failing exit status that disrupts automated validation.
        show_no_staged_changes
        return 0
    fi

    if is_verbose_enabled; then
        local staged_files
        staged_files=$(git diff --cached --name-only | wc -l | tr -d ' ')
        info "Found $staged_files staged file(s)"
    fi

    return 0
}
