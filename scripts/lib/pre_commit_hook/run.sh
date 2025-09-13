#!/usr/bin/env bash

# Cisco WNC Pre-commit Hook - Run pre-commit validations
# Provides functions to validate repository state and branch rules before allowing commit

# Load WNC libraries
PRE_COMMIT_SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "${PRE_COMMIT_SCRIPT_DIR}/../bootstrap.sh"
init_wnc_libraries

# Validate repository state and branch rules before allowing commit
run_pre_commit_validation() {
  local branch=""

  # Validate Git environment and get current branch
  if ! validate_git_environment; then
    return 1
  fi
  branch="${CURRENT_BRANCH:-}"

  # Clean up backup files before validation
  if ! cleanup_backup_files; then
    return 1
  fi

  # Display current validation status
  show_validation_status "$branch"

  # Validate branch permissions and staged changes
  if ! validate_branch_permissions "$branch"; then
    return 1
  fi

  if ! validate_staged_changes; then
    return 1
  fi

  # Run comprehensive formatting (gofumpt, goimports)
  if ! run_formats; then
    return 1
  fi

  # Run comprehensive linting (Go, Shell, Markdown)
  if ! run_lints; then
    return 1
  fi

  # Display success message and return
  show_validation_success "$branch"
  printf '\n'
  return 0
}
