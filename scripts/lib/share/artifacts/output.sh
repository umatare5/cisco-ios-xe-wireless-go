#!/usr/bin/env bash

# Cisco WNC Artifacts - Output functions
# Provides functions to format and display artifacts messages using unified log_* functions

# Format file removal status using success
format_file_removed() {
    local path="$1"
    local size="${2:-}"

    if [[ -n "$size" ]]; then
        success "Removed: $path ($size)"
        return
    fi

    success "Removed: $path"
}

# Format dry run item using info with [DRY-RUN] prefix
format_dry_run_item() {
    local action="$1"
    local path="$2"
    local size="${3:-}"

    if [[ -n "$size" ]]; then
        info "[DRY-RUN] $action: $path ($size)"
        return
    fi

    info "[DRY-RUN] $action: $path"
}

# Format cache information using appropriate log_* functions
format_cache_info() {
    local cache_type="$1"
    local cache_path="$2"
    local exists="$3"

    if [[ "$exists" == "true" ]]; then
        info "$cache_type: $cache_path"
        return
    fi

    warn "$cache_type: $cache_path (not found)"
}
