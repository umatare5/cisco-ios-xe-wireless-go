#!/usr/bin/env bash

# List YANG Model Filters - Core Operations
# Provides core functionality for extracting filterable fields from YANG data

run_yang_filter_extraction() {
    local yang_model_name="$1"    # Legacy parameter, now determined by statement_name
    local statement_name="$2"
    local controller_host="$3"
    local access_token="$4"
    local request_protocol="$5"
    local insecure_mode="$6"
    local output_format="$7"

    # Suppress shellcheck warning for legacy parameter
    : "${yang_model_name}"

    show_yang_banner
    show_yang_verbose_info "$controller_host" "$request_protocol" "$output_format"

    if [[ -z "$controller_host" || -z "$access_token" ]]; then
        log_error "Controller and token required"
        return 1
    fi

    local output_file="./tmp/yang_filters_$$.json"
    mkdir -p ./tmp

    progress "Fetching YANG data for filter extraction: /$statement_name"

    local insecure_flag=""
    if [[ "$insecure_mode" == "1" ]]; then
        insecure_flag="-k"
    fi

    # Determine the correct YANG module from statement name
    local yang_module
    case "$statement_name" in
        "wlan-cfg-data")
            yang_module="Cisco-IOS-XE-wireless-wlan-cfg"
            ;;
        "site-cfg-data")
            yang_module="Cisco-IOS-XE-wireless-site-cfg"
            ;;
        "ap-cfg-data")
            yang_module="Cisco-IOS-XE-wireless-ap-cfg"
            ;;
        "client-oper-data")
            yang_module="Cisco-IOS-XE-wireless-client-oper"
            ;;
        "access-point-oper-data")
            yang_module="Cisco-IOS-XE-wireless-access-point-oper"
            ;;
        "ap-global-oper-data")
            yang_module="Cisco-IOS-XE-wireless-ap-global-oper"
            ;;
        "wlan-global-oper-data")
            yang_module="Cisco-IOS-XE-wireless-wlan-global-oper"
            ;;
        "client-global-oper-data")
            yang_module="Cisco-IOS-XE-wireless-client-global-oper"
            ;;
        "general-cfg-data")
            yang_module="Cisco-IOS-XE-wireless-general-cfg"
            ;;
        *)
            yang_module="Cisco-IOS-XE-wireless-wlan-cfg"  # default
            ;;
    esac

    local statement_url
    statement_url="$(build_yang_statement_url "$request_protocol" "$controller_host" \
                     "$yang_module" "$statement_name")"

    execute_curl_request "$statement_url" "$access_token" "$insecure_flag" "$output_file"
    extract_filterable_fields "$output_file" "$statement_name" "$output_format"
    display_yang_operation_results 0 "YANG filter extraction" "/$statement_name"
}
