#!/usr/bin/env bash

# Cisco WNC Common Library Loader
# Centralized library loading for all WNC scripts

# Function to source all required WNC libraries
source_wnc_libraries() {
    local script_dir="$1"

    # Array of required libraries in dependency order
    local libraries=(
        "constants.sh"
        "validation.sh"
        "dependencies.sh"
        "authentication.sh"
        "http_client.sh"
        "output_formatter.sh"
        "file_utils.sh"
    )

    # Source each library
    for lib in "${libraries[@]}"; do
        local lib_path="${script_dir}/lib/${lib}"
        if [[ -f "$lib_path" ]]; then
            # shellcheck source=/dev/null
            source "$lib_path"
        else
            echo "Error: Cannot find library file: $lib_path" >&2
            exit 1
        fi
    done
}
