#!/usr/bin/env bash

# Cisco WNC YANG Operations - Core Functions
# Provides core functions for YANG operations via RESTCONF API

# Generate insecure curl flag based on configuration
_insecure_flag() {
    # Early return for secure mode (empty flag)
    if ! is_insecure_enabled; then
        return 0
    fi

    printf '%s\n' "-k"
}

# Check if controller is provided
has_controller() { [[ -n "${1:-}" ]]; }

# Check if token is provided
has_token() { [[ -n "${1:-}" ]]; }

# Validate required environment variables for YANG operations
_require_env() {
    # Early return for missing controller
    if ! has_controller "$1"; then
        error "WNC_CONTROLLER required"
        return 1
    fi

    # Early return for missing token
    if ! has_token "$2"; then
        error "WNC_ACCESS_TOKEN required"
        return 1
    fi
}

# List YANG models from controller via RESTCONF API
_list_models() {
    local controller="$1"
    local token="$2"
    local output_file="$3"
    local protocol="$4"
    local url
    url="$(build_yang_models_url "$protocol" "$controller")"
    execute_curl_request "$url" "$token" "$(_insecure_flag)" "$output_file"
}

# Get specific YANG model details via RESTCONF API
_get_model() {
    local controller="$1"
    local token="$2"
    local model_name="$3"
    local revision="$4"
    local output_file="$5"
    local protocol="$6"
    local url
    url="$(build_yang_model_details_url "$protocol" "$controller" "$model_name" "$revision")"
    execute_curl_request "$url" "$token" "$(_insecure_flag)" "$output_file"
}

# Get YANG statement details via RESTCONF API
_get_statement() {
    local controller="$1"
    local token="$2"
    local model_name="$3"
    local statement_name="$4"
    local output_file="$5"
    local protocol="$6"
    local url
    url="$(build_yang_statement_url "$protocol" "$controller" "$model_name" "$statement_name")"
    execute_curl_request "$url" "$token" "$(_insecure_flag)" "$output_file"
}


# Post-process YANG statement output with optional JSON filtering
post_process_statement_output() {
    local file_path="$1"
    local search_key="$2"
    local output_format="$3"

    # Early return for non-JSON formats or missing jq
    if [[ "${output_format}" != "json" ]] || ! command -v jq >/dev/null 2>&1; then
        grep -A 20 -B 5 "$search_key" "$file_path" || cat "$file_path"
        return 0
    fi

    # Process JSON with jq filtering
    jq --arg search_key "$search_key" \
       '..|objects|to_entries[]|select(.key==$search_key)|{($search_key):.value}' \
       "$file_path" || cat "$file_path"
}
