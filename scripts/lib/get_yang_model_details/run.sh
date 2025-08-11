#!/usr/bin/env bash

# Cisco WNC YANG Model Details - Fetch and display YANG model information
# Provides functions to retrieve and display details for specific YANG models

# Fetch details for a specific YANG model with banner and progress display
run_yang_get_model_operation() {
    local yang_model_name="$1"
    local controller_host="$2"
    local access_token="$3"
    local request_protocol="$4"
    local insecure_mode="$5"
    local model_revision="$6"
    local output_format="${7:-json}"
    local raw_output="${8:-false}"

    show_yang_banner
    show_yang_verbose_info "$controller_host" "$request_protocol" "$output_format"

    if [[ -z "$controller_host" || -z "$access_token" ]]; then
        error "Controller and token required"
        return 1
    fi

    local output_file="./tmp/yang_model_$$.json"
    mkdir -p ./tmp
    progress "Fetching YANG model details for: $yang_model_name (rev: $model_revision)"

    local insecure_flag=""
    if [[ "$insecure_mode" == "1" ]]; then
        insecure_flag="-k"
    fi

    local model_details_url
    model_details_url="$(build_yang_model_details_url "$request_protocol" "$controller_host" \
                         "$yang_model_name" "$model_revision")"
    execute_curl_request "$model_details_url" "$access_token" "$insecure_flag" "$output_file"
    format_yang_output_file "$output_file" "$output_format" "$raw_output"
    display_yang_operation_results 0 "YANG model retrieval" "$yang_model_name"
}
