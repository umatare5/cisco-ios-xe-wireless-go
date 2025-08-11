#!/usr/bin/env bash

# Cisco WNC YANG Models - List available YANG models
# Provides functions to list and display available YANG models from a controller

# Execute YANG model listing operation with full banner and progress display
run_yang_list_operation() {
    # Extract arguments from entry script context
    local controller_host="${1:-}"
    local access_token="${2:-}"
    local request_protocol="${3:-https}"
    local insecure_mode="${4:-false}"
    local output_format="${5:-json}"
    local raw_output="${6:-false}"

    show_yang_banner
    show_yang_verbose_info "$controller_host" "$request_protocol" "json"

    if [[ -z "$WNC_CONTROLLER" || -z "$WNC_ACCESS_TOKEN" ]]; then
        error "Controller and token required"
        return 1
    fi

    local output_file="./tmp/yang_list_$$.json"
    mkdir -p ./tmp
    progress "Fetching YANG models list..."

    local insecure_flag=""
    if [[ "$insecure_mode" == "1" ]]; then
        insecure_flag="-k"
    fi

    local models_url
    models_url="$(build_yang_models_url "$request_protocol" "$controller_host")"
    execute_curl_request "$models_url" "$access_token" "$insecure_flag" "$output_file"
    format_yang_output_file "$output_file" "$output_format" "$raw_output"
    display_yang_operation_results 0 "YANG models listing"
}
