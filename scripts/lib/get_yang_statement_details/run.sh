#!/usr/bin/env bash

# Cisco WNC YANG Statement Details - Core Operations
# Provides core functionality for YANG statement information retrieval

run_yang_get_statement_operation() {
    local yang_model_name="$1"
    local statement_name="$2"
    local controller_host="$3"
    local access_token="$4"
    local request_protocol="$5"
    local insecure_mode="$6"
    local output_format="$7"
    # Note: raw parameter not needed as post_process_statement_output handles formatting

    show_yang_banner
    show_yang_verbose_info "$controller_host" "$request_protocol" "$output_format"

    if [[ -z "$controller_host" || -z "$access_token" ]]; then
        error "Controller and token required"
        return 1
    fi

    local output_file="./tmp/yang_stmt_$$.json"
    mkdir -p ./tmp
    progress "Fetching YANG statement details for: $yang_model_name/$statement_name"

    local insecure_flag=""
    if [[ "$insecure_mode" == "1" ]]; then
        insecure_flag="-k"
    fi

    local statement_url
    statement_url="$(build_yang_statement_url "$request_protocol" "$controller_host" \
                     "$yang_model_name" "$statement_name")"
    execute_curl_request "$statement_url" "$access_token" "$insecure_flag" "$output_file"
    post_process_statement_output "$output_file" "$statement_name" "$output_format"
    display_yang_operation_results 0 "YANG statement retrieval" "$yang_model_name/$statement_name"
}
