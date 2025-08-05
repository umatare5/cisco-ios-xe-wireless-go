#!/usr/bin/env bash

# Cisco WNC Common Library Loader
# Centralized library loading for all WNC scripts

# Common predicate functions
is_valid_directory() {
    [[ -d "$1" ]]
}

is_verbose_enabled() {
    [[ "${argc_verbose:-0}" == "1" ]]
}

is_no_color_enabled() {
    [[ "${argc_no_color:-0}" == "1" ]]
}

is_insecure_enabled() {
    [[ "${argc_insecure:-0}" == "1" ]]
}

validate_project_directory() {
    local dir="$1"
    is_valid_directory "$dir"
}

has_golangci_lint() {
    command -v golangci-lint >/dev/null 2>&1
}

# Function to source all required WNC libraries
SOURCE_WNC_LIBRARIES() {
    local script_dir="$1"

    # Array of required libraries in dependency order
    local libraries=(
        "constants.sh"
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
    SOURCE_WNC_LIBRARIES "$script_dir"

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
