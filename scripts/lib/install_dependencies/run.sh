#!/usr/bin/env bash

# Cisco WNC Dependencies - Run dependency management operations
# Provides functions to install, update, and verify project dependencies

# Install/update/verify project dependencies
run_dependencies_operation() {
  local project_root="$1"
  local clean="$2"
  local update="$3"
  local download_only="$4"
  local verify="$5"
  local verbose="$6"

  # Set global verbose flag for dependencies functions
  export DEPENDENCIES_VERBOSE="${verbose}"

  # Display operation information and validate environment
  show_dependencies_banner
  if ! validate_dependencies_environment "$project_root"; then
    return 1
  fi

  # Change to project directory for dependency operations
  cd "$project_root" || {
    error "Failed to change to project directory: $project_root"
    return 1
  }

  # Execute dependency management operations in sequence
  execute_dependencies_clean "$clean" || return 1
  execute_dependencies_download "$download_only" "$verbose" || return 1
  execute_dependencies_update "$update" "$verbose" || return 1
  execute_dependencies_verify "$verify" || return 1

  # Display final operation summary
  display_dependencies_summary
  return 0
}
