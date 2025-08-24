#!/usr/bin/env bash

# YANG Model Filter Extraction - Client Operational Functions
# Provides processing functions for extracting filterable fields from client operational data

# Include guard
if [[ -n "${CLIENT_OPER_FILTER_LOADED:-}" ]]; then
    return 0
fi
readonly CLIENT_OPER_FILTER_LOADED=1

# Extract client-oper filters in JSON format
extract_client_oper_filters_as_json() {
    local file_path="$1"

    info "Available Client operational filterable fields (JSON format):"
    jq '
        .["Cisco-IOS-XE-wireless-client-oper:client-oper-data"] as $data |
        {
            "common-oper-data": [
                $data["common-oper-data"][]? |
                {
                    "client-mac": .["client-mac"],
                    "ap-name": .["ap-name"],
                    "wlan-id": .["wlan-id"],
                    "client-type": .["client-type"],
                    "co-state": .["co-state"],
                    "ms-radio-type": .["ms-radio-type"]?,
                    "username": .["username"]?
                }
            ]
        }
    ' "$file_path" 2>/dev/null || {
        log_error "Failed to extract client-oper filterable fields"
        return 1
    }
}

# Extract client-oper filters in table format
extract_client_oper_filters_as_table() {
    local file_path="$1"

    info "Available Client operational filterable fields:"
    echo

    # Extract Client Operational Data
    success "Client Operational Data (common-oper-data):"
    printf "%-20s %-25s %-10s %-20s %-20s %-25s\n" "Client MAC" "AP Name" "WLAN ID" "Client Type" "State" "Radio Type"
    printf "%-20s %-25s %-10s %-20s %-20s %-25s\n" "----------" "-------" "-------" "-----------" "-----" "----------"

    jq -r '
        .["Cisco-IOS-XE-wireless-client-oper:client-oper-data"]["common-oper-data"][]? |
        [
            (.["client-mac"] // "N/A"),
            (.["ap-name"] // "N/A"),
            (.["wlan-id"] // "N/A" | tostring),
            (.["client-type"] // "N/A"),
            (.["co-state"] // "N/A"),
            (.["ms-radio-type"] // "N/A")
        ] |
        @tsv
    ' "$file_path" 2>/dev/null | while IFS=$'\t' read -r client_mac ap_name wlan_id client_type co_state ms_radio_type; do
        printf "%-20s %-25s %-10s %-20s %-20s %-25s\n" "$client_mac" "$ap_name" "$wlan_id" "$client_type" "$co_state" "$ms_radio_type"
    done
    echo

    # Show available client-oper filter examples
    success "Client-Oper Filter Examples:"
    echo "  client-mac=08:84:9d:92:47:00        # Filter by client MAC address"
    echo "  ap-name=AP9166-02          # Filter by AP name"
    echo "  wlan-id=1                           # Filter by WLAN ID"
    echo "  client-type=dot11-client-normal     # Filter by client type"
    echo "  co-state=client-status-run          # Filter by client state"
    echo "  ms-radio-type=client-dot11n-24-ghz-prot # Filter by radio type"
    echo "  username=user1                      # Filter by username"
    echo
}
