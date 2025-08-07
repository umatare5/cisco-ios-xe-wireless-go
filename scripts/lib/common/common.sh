#!/usr/bin/env bash

# Cisco WNC Common Library Loader
# Centralized library loading for all WNC scripts

# Common predicate functions
is_valid_directory() {
    [[ -d "$1" ]]
}

is_verbose_enabled() {
    is_enabled "${argc_verbose:-0}"
}

is_no_color_enabled() {
    is_enabled "${argc_no_color:-0}"
}

is_insecure_enabled() {
    is_enabled "${argc_insecure:-0}"
}

validate_project_directory() {
    local dir="$1"
    is_valid_directory "$dir"
}

has_golangci_lint() {
    command -v golangci-lint >/dev/null 2>&1
}

# Function to source all required WNC libraries
source_wnc_libraries() {
    local script_dir="$1"

    # Array of required libraries in dependency order
    local libraries=(
        "constants.sh"
        "cli_validation.sh"
        "validation.sh"
        "argument_parsing.sh"
        "dependencies.sh"
        "authentication.sh"
        "http_client.sh"
        "output_formatter.sh"
        "file_utils.sh"
        "build_tools.sh"
        "testing.sh"
        "yang_common.sh"
    )

    # Source each library
    for lib in "${libraries[@]}"; do
        local lib_path="${script_dir}/lib/common/${lib}"
        if [[ -f "$lib_path" ]]; then
            # shellcheck source=/dev/null
            source "$lib_path"
        else
            show_error "Cannot find library file: $lib_path"
            exit 1
        fi
    done
}

# Standard script initialization function
INIT_SCRIPT_ENVIRONMENT() {
    local script_name="${1:-unknown}"
    local description="${2:-No description provided}"

    # Standard script setup
    set -euo pipefail

    # Get script directory (consistent across all scripts)
    local script_dir
    script_dir="$(cd "$(dirname "${BASH_SOURCE[1]}")" && pwd)"
    readonly script_dir

    # Load all WNC libraries
    source_wnc_libraries "$script_dir"

    # Set global variables for caller
    export WNC_SCRIPT_DIR="$script_dir"
    export WNC_SCRIPT_NAME="$script_name"
    export WNC_SCRIPT_DESCRIPTION="$description"

    # Return script directory for caller convenience
    echo "$script_dir"
}

# Lightweight script initialization with minimal setup
INIT_SIMPLE_SCRIPT() {
    # Basic setup only
    set -euo pipefail

    # Get and return script directory
    local script_dir
    script_dir="$(cd "$(dirname "${BASH_SOURCE[1]}")" && pwd)"
    readonly script_dir

    echo "$script_dir"
}

# Function to source module-specific libraries
source_module_libraries() {
    local module_dir="$1"

    # Standard module libraries in dependency order
    local libraries=(
        "help.sh"
        "output.sh"
        "core.sh"
    )

    # Source each library if it exists
    for lib in "${libraries[@]}"; do
        local lib_path="${module_dir}/${lib}"
        if [[ -f "$lib_path" ]]; then
            # shellcheck source=/dev/null
            source "$lib_path"
        fi
    done
}

# Unified script initialization function
init_script_libraries() {
    local script_dir="$1"
    local module_dir="$2"

    # Load all WNC common libraries
    source_wnc_libraries "$script_dir"

    # Load module-specific libraries
    source_module_libraries "$module_dir"
}

# argc predicate functions for clean conditional logic
is_true() { [[ "${1:-false}" == "true" ]]; }
is_false() { [[ "${1:-false}" == "false" ]]; }
is_one() { [[ "${1:-0}" == "1" ]]; }
is_zero() { [[ "${1:-0}" == "0" ]]; }
is_enabled() { is_one "$1"; }
is_disabled() { is_zero "$1"; }
