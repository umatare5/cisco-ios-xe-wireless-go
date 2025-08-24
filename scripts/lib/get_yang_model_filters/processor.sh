#!/usr/bin/env bash

# YANG Model Filter Extraction - Processing Functions
# Provides processing functions for extracting filterable fields from YANG data

# Load endpoint-specific processors
# shellcheck source=/dev/null
source "${BASH_SOURCE[0]%/*}/wlan_cfg.sh"
# shellcheck source=/dev/null
source "${BASH_SOURCE[0]%/*}/site_cfg.sh"
# shellcheck source=/dev/null
source "${BASH_SOURCE[0]%/*}/ap_cfg.sh"
# shellcheck source=/dev/null
source "${BASH_SOURCE[0]%/*}/client_oper.sh"
# shellcheck source=/dev/null
source "${BASH_SOURCE[0]%/*}/ap_oper.sh"
# shellcheck source=/dev/null
source "${BASH_SOURCE[0]%/*}/ap_global_oper.sh"
# shellcheck source=/dev/null
source "${BASH_SOURCE[0]%/*}/wlan_global_oper.sh"
# shellcheck source=/dev/null
source "${BASH_SOURCE[0]%/*}/client_global_oper.sh"

extract_filterable_fields() {
    local file_path="$1"
    local search_key="$2"
    local output_format="$3"

    # Clean JSON from HTTP response headers
    local clean_file="./tmp/yang_clean_$$.json"
    sed -n '/^{/,$p' "$file_path" > "$clean_file"

    # Determine the model type based on search_key
    if [[ "$search_key" == "site-cfg-data" ]]; then
        if [[ "$output_format" == "json" ]]; then
            extract_site_cfg_filters_as_json "$clean_file"
            return
        fi
        extract_site_cfg_filters_as_table "$clean_file"
        return
    fi

    if [[ "$search_key" == "ap-cfg-data" ]]; then
        if [[ "$output_format" == "json" ]]; then
            extract_ap_cfg_filters_as_json "$clean_file"
            return
        fi
        extract_ap_cfg_filters_as_table "$clean_file"
        return
    fi

    if [[ "$search_key" == "client-oper-data" ]]; then
        if [[ "$output_format" == "json" ]]; then
            extract_client_oper_filters_as_json "$clean_file"
            return
        fi
        extract_client_oper_filters_as_table "$clean_file"
        return
    fi

    if [[ "$search_key" == "access-point-oper-data" ]]; then
        if [[ "$output_format" == "json" ]]; then
            extract_ap_oper_filters_as_json "$clean_file"
            return
        fi
        extract_ap_oper_filters_as_table "$clean_file"
        return
    fi

    if [[ "$search_key" == "ap-global-oper-data" ]]; then
        if [[ "$output_format" == "json" ]]; then
            extract_ap_global_oper_filters_as_json "$clean_file"
            return
        fi
        extract_ap_global_oper_filters_as_table "$clean_file"
        return
    fi

    if [[ "$search_key" == "wlan-global-oper-data" ]]; then
        if [[ "$output_format" == "json" ]]; then
            extract_wlan_global_oper_filters_as_json "$clean_file"
            return
        fi
        extract_wlan_global_oper_filters_as_table "$clean_file"
        return
    fi

    if [[ "$search_key" == "client-global-oper-data" ]]; then
        if [[ "$output_format" == "json" ]]; then
            extract_client_global_oper_filters_as_json "$clean_file"
            return
        fi
        extract_client_global_oper_filters_as_table "$clean_file"
        return
    fi

    if [[ "$search_key" == "general-cfg-data" ]]; then
        info "Model 'Cisco-IOS-XE-wireless-general-cfg' does not support filtering operations"
    fi

    # Default to wlan-cfg processing
    if [[ "$output_format" == "json" ]]; then
        extract_wlan_cfg_filters_as_json "$clean_file"
        return
    fi

    extract_wlan_cfg_filters_as_table "$clean_file"
}
