#!/usr/bin/env bash

# WLAN Global Operational Filters - Extraction Functions
# Provides functions to extract filterable fields from WLAN global operational data

# Include guard
if [[ -n "${WLAN_GLOBAL_OPER_FILTER_LOADED:-}" ]]; then
    return 0
fi
readonly WLAN_GLOBAL_OPER_FILTER_LOADED=1

# Extract WLAN global operational filters as JSON
extract_wlan_global_oper_filters_as_json() {
    local file_path="$1"

    jq '{
        "wlan_info": {
            "wlan-profile": [.["Cisco-IOS-XE-wireless-wlan-global-oper:wlan-global-oper-data"]["wlan-info"][]."wlan-profile" | select(. != null)],
            "curr-clients-count": [.["Cisco-IOS-XE-wireless-wlan-global-oper:wlan-global-oper-data"]["wlan-info"][]."curr-clients-count" | select(. != null)],
            "per-wlan-max-client-syslog": [.["Cisco-IOS-XE-wireless-wlan-global-oper:wlan-global-oper-data"]["wlan-info"][]."per-wlan-max-client-syslog" | select(. != null)]
        }
    }' "$file_path"
}

# Extract WLAN global operational filters as table
extract_wlan_global_oper_filters_as_table() {
    local file_path="$1"

    echo "✓ Success: WLAN Global Operational Data (wlan-global-oper-data):"

    # Extract WLAN info data
    local wlan_profiles client_counts syslog_enabled
    wlan_profiles=$(jq -r '.["Cisco-IOS-XE-wireless-wlan-global-oper:wlan-global-oper-data"]["wlan-info"][]."wlan-profile" | select(. != null)' "$file_path" 2>/dev/null)
    client_counts=$(jq -r '.["Cisco-IOS-XE-wireless-wlan-global-oper:wlan-global-oper-data"]["wlan-info"][]."curr-clients-count" | select(. != null)' "$file_path" 2>/dev/null)
    syslog_enabled=$(jq -r '.["Cisco-IOS-XE-wireless-wlan-global-oper:wlan-global-oper-data"]["wlan-info"][]."per-wlan-max-client-syslog" | select(. != null)' "$file_path" 2>/dev/null)

    printf "%-20s %-15s %-25s\n" "WLAN Profile" "Current Clients" "Max Client Syslog Enabled"
    printf "%-20s %-15s %-25s\n" "------------" "---------------" "-------------------------"

    # Display sample data from each WLAN
    if [[ -n "$wlan_profiles" ]]; then
        local profile count syslog
        while IFS= read -r profile; do
            count=$(echo "$client_counts" | head -1)
            syslog=$(echo "$syslog_enabled" | head -1)

            printf "%-20s %-15s %-25s\n" \
                "${profile:-N/A}" "${count:-0}" "${syslog:-false}"

            # Remove the first line for next iteration
            client_counts=$(echo "$client_counts" | tail -n +2)
            syslog_enabled=$(echo "$syslog_enabled" | tail -n +2)
        done <<< "$wlan_profiles"
    fi

    echo ""
    echo "✓ Success: WLAN-Global-Oper Filter Examples:"
    echo "  wlan-profile=labo2                           # Filter by WLAN profile name"
    echo "  profile-name=labo1                       # Filter by profile name (alternate)"
    echo "  curr-clients-count=16                       # Filter by exact client count"
    echo "  min-clients-count=10                        # Filter by minimum client count"
    echo "  max-clients-count=5                         # Filter by maximum client count"
    echo "  per-wlan-max-client-syslog=true             # Filter by max client syslog status"
    echo "  has-clients=true                            # Filter WLANs with active clients (> 0)"
    echo "  has-clients=false                           # Filter WLANs with no clients (= 0)"
    echo ""
}
