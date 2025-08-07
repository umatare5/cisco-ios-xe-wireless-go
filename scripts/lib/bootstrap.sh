#!/usr/bin/env bash

# Cisco WNC Scripts Bootstrap Library
# Centralized library loading and initialization for all WNC scripts

set -euo pipefail

# Global paths (check if already set to avoid conflicts)
if [[ -z "${WNC_LIB_ROOT:-}" ]]; then
    WNC_LIB_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
    readonly WNC_LIB_ROOT
    readonly WNC_CORE_LIB="${WNC_LIB_ROOT}/core"
    readonly WNC_UTILS_LIB="${WNC_LIB_ROOT}/utils"
    readonly WNC_NETWORK_LIB="${WNC_LIB_ROOT}/network"
    readonly WNC_OUTPUT_LIB="${WNC_LIB_ROOT}/output"
fi

# Core library loading
source_core_libraries() {
    source "${WNC_CORE_LIB}/constants.sh"
    source "${WNC_CORE_LIB}/argc_common.sh"
    source "${WNC_CORE_LIB}/argument_parsing.sh"
}

# Utility library loading
source_utils_libraries() {
    source "${WNC_UTILS_LIB}/validation.sh"
    source "${WNC_UTILS_LIB}/file_utils.sh"
    source "${WNC_UTILS_LIB}/cli_validation.sh"
    source "${WNC_UTILS_LIB}/build_tools.sh"
    source "${WNC_UTILS_LIB}/dependencies.sh"
}

# Network library loading
source_network_libraries() {
    source "${WNC_NETWORK_LIB}/http_client.sh"
    source "${WNC_NETWORK_LIB}/authentication.sh"
    source "${WNC_NETWORK_LIB}/yang_common.sh"
}

# Output library loading
source_output_libraries() {
    source "${WNC_OUTPUT_LIB}/output_formatter.sh"
}

# Module-specific library loading
source_module_libraries() {
    local module_dir="$1"

    if [[ -d "$module_dir" ]]; then
        # Source all .sh files in the module directory
        for lib_file in "$module_dir"/*.sh; do
            # shellcheck disable=SC1090  # Dynamic source loading
            [[ -f "$lib_file" ]] && source "$lib_file"
        done
    fi
}

# Main initialization function
init_wnc_libraries() {
    local script_dir="${1:-}"
    local module_dir="${2:-}"

    # Always load core libraries first
    source_core_libraries

    # Load utilities
    source_utils_libraries

    # Load network libraries if needed
    if [[ "${WNC_LOAD_NETWORK:-true}" == "true" ]]; then
        source_network_libraries
    fi

    # Load output libraries
    source_output_libraries

    # Load module-specific libraries if provided
    if [[ -n "$module_dir" ]]; then
        source_module_libraries "$module_dir"
    fi
}

# Compatibility function (deprecated)
init_script_libraries() {
    local script_dir="$1"
    local module_dir="${2:-}"

    # Show deprecation warning if verbose
    if [[ "${argc_verbose:-0}" == "1" ]]; then
        echo "Warning: init_script_libraries is deprecated. Use init_wnc_libraries instead." >&2
    fi

    init_wnc_libraries "$script_dir" "$module_dir"
}

# Lightweight initialization for simple scripts
init_wnc_basic() {
    source_core_libraries
    source_utils_libraries
}

# Network-only initialization
init_wnc_network() {
    source_core_libraries
    source_network_libraries
    source_output_libraries
}

# Export for child scripts
export WNC_LIB_ROOT WNC_CORE_LIB WNC_UTILS_LIB WNC_NETWORK_LIB WNC_OUTPUT_LIB
