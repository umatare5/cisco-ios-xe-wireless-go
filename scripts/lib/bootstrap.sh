#!/usr/bin/env bash

# Cisco WNC Scripts Bootstrap Library
# Centralized library loading and initialization for all WNC scripts

# Global paths (check if already set to avoid conflicts)
if [[ -z "${WNC_LIB_ROOT:-}" ]]; then
    WNC_LIB_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
    readonly WNC_LIB_ROOT
fi

# Module-specific library loading - Helper functions (private)
_is_directory_valid() {
    [[ -d "$1" ]]
}

_is_file_exists() {
    [[ -f "$1" ]]
}

# Load all share libraries dynamically
_load_share_libraries() {
    local share_dir="${WNC_LIB_ROOT}/share"

    if ! _is_directory_valid "$share_dir"; then
        return 0
    fi

    # Loop through all directories in share/
    for lib_dir in "$share_dir"/*; do
        if ! _is_directory_valid "$lib_dir"; then
            continue
        fi

        # Load all .sh files in this SDK directory
        for lib_file in "$lib_dir"/*.sh; do
            if _is_file_exists "$lib_file"; then
                source "$lib_file"
            fi
        done
    done
}

_load_priority_scripts() {
    local module_dir="$1"
    local priority_scripts=("predicate.sh" "banner.sh" "help.sh")

    local script_file
    for script_file in "${priority_scripts[@]}"; do
        if _is_file_exists "$module_dir/$script_file"; then
            source "$module_dir/$script_file"
        fi
    done
}

_load_other_scripts() {
    local module_dir="$1"
    local skip_scripts=("predicate.sh" "banner.sh" "help.sh" "core.sh")

    for lib_file in "$module_dir"/*.sh; do
        if ! _is_file_exists "$lib_file"; then
            continue
        fi

        local filename
        filename=$(basename "$lib_file")

        # Check if should skip this script
        local should_skip=false
        local skip_script
        for skip_script in "${skip_scripts[@]}"; do
            if [[ "$filename" == "$skip_script" ]]; then
                should_skip=true
                break
            fi
        done

        if [[ "$should_skip" == false ]]; then
            source "$lib_file"
        fi
    done
}

_load_core_script() {
    local module_dir="$1"

    if _is_file_exists "$module_dir/core.sh"; then
        source "$module_dir/core.sh"
    fi
}

# Module-specific library loading
source_module_libraries() {
    local module_dir="$1"

    # Early return if module directory doesn't exist
    if ! _is_directory_valid "$module_dir"; then
        return 0
    fi

    # Load libraries in specific order
    _load_priority_scripts "$module_dir"
    _load_other_scripts "$module_dir"
    _load_core_script "$module_dir"
}

# Main initialization function
init_wnc_libraries() {
    local module_dir="${1:-}"

    # Load all share libraries dynamically
    _load_share_libraries

    # Load module-specific libraries if provided
    if [[ -n "$module_dir" ]]; then
        source_module_libraries "$module_dir"
    fi
}

# Export for child scripts
export WNC_LIB_ROOT
