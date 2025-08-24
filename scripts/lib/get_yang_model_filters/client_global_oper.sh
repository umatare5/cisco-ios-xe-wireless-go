#!/usr/bin/env bash

# YANG Model Filter Extraction - Client Global Operational Data Functions
# Provides processing functions for extracting filterable fields from client global operational data

# Include guard
if [[ -n "${CLIENT_GLOBAL_OPER_FILTER_LOADED:-}" ]]; then
    return 0
fi
readonly CLIENT_GLOBAL_OPER_FILTER_LOADED=1

# Extract client-global-oper filters in JSON format
extract_client_global_oper_filters_as_json() {
    local file_path="$1"

    info "Available filterable fields (JSON format):"
    jq '
        .["Cisco-IOS-XE-wireless-client-global-oper:client-global-oper-data"] as $data |
        {
            "sm-device-list": [
                $data["sm-device-count"]["sm-device-list"][]? |
                {
                    "device-type": .["device-type"],
                    "device-count": .["device-count"]
                }
            ],
            "client-latency-stats": {
                "client-state-stats-value": [
                    $data["client-latency-stats"]["client-states-stats"]["client-state-stats-value"][]? |
                    {
                        "avg-client-state-duration": .["avg-client-state-duration"],
                        "total-sessions": .["total-sessions"]
                    }
                ],
                "avg-run-state-latency": $data["client-latency-stats"]["avg-run-state-latency"]?
            },
            "global-stats": {
                "client-stats": $data["client-stats"]?,
                "client-live-stats": $data["client-live-stats"]?,
                "client-global-stats-data": $data["client-global-stats-data"]?,
                "client-dot11-stats": $data["client-dot11-stats"]?,
                "client-exclusion-stats": $data["client-exclusion-stats"]?,
                "dot1x-global-stats": $data["dot1x-global-stats"]?,
                "sm-webauth-stats": $data["sm-webauth-stats"]?
            }
        }
    ' "$file_path" 2>/dev/null || {
        log_error "Failed to extract filterable fields"
        return 1
    }
}

# Extract client-global-oper filters in table format
extract_client_global_oper_filters_as_table() {
    local file_path="$1"

    info "Available filterable fields:"
    echo

    # Extract SM Device List
    success "Smart Monitoring Device List (sm-device-list):"
    printf "%-40s %-15s\n" "Device Type" "Device Count"
    printf "%-40s %-15s\n" "-----------" "------------"

    jq -r '
        .["Cisco-IOS-XE-wireless-client-global-oper:client-global-oper-data"]["sm-device-count"]["sm-device-list"][]? |
        [
            (.["device-type"] // "N/A"),
            (.["device-count"] // "0")
        ] |
        @tsv
    ' "$file_path" 2>/dev/null | while IFS=$'\t' read -r device_type device_count; do
        printf "%-40s %-15s\n" "$device_type" "$device_count"
    done
    echo

    # Extract Client Latency Stats
    success "Client State Statistics (client-latency-stats):"
    printf "%-8s %-25s %-15s\n" "Index" "Avg Duration (ms)" "Total Sessions"
    printf "%-8s %-25s %-15s\n" "-----" "----------------" "--------------"

    local index=0
    jq -r '
        .["Cisco-IOS-XE-wireless-client-global-oper:client-global-oper-data"]["client-latency-stats"]["client-states-stats"]["client-state-stats-value"][]? |
        [
            (.["avg-client-state-duration"] // "0"),
            (.["total-sessions"] // "0")
        ] |
        @tsv
    ' "$file_path" 2>/dev/null | while IFS=$'\t' read -r avg_duration total_sessions; do
        printf "%-8s %-25s %-15s\n" "$((++index))" "$avg_duration" "$total_sessions"
    done
    echo

    # Extract Client Live Stats
    success "Client Live Statistics (client-live-stats):"
    printf "%-25s %-15s\n" "State" "Client Count"
    printf "%-25s %-15s\n" "-----" "------------"

    jq -r '
        .["Cisco-IOS-XE-wireless-client-global-oper:client-global-oper-data"]["client-live-stats"] |
        to_entries[] |
        [
            .key,
            (.value // "0")
        ] |
        @tsv
    ' "$file_path" 2>/dev/null | while IFS=$'\t' read -r state_name client_count; do
        # Convert kebab-case to human readable
        local readable_state
        readable_state=$(echo "$state_name" | sed 's/-/ /g' | sed 's/\b\w/\U&/g')
        printf "%-25s %-15s\n" "$readable_state" "$client_count"
    done
    echo

    # Extract Global Statistics Summary
    success "Available Global Statistics:"
    echo "  client-stats              # Client deletion reasons and statistics"
    echo "  client-live-stats         # Current client state distribution"
    echo "  client-global-stats-data  # Session and mobility statistics"
    echo "  client-dot11-stats        # 802.11 association and authentication stats"
    echo "  client-exclusion-stats    # Client exclusion and disabled counts"
    echo "  dot1x-global-stats        # 802.1x authentication statistics"
    echo "  sm-webauth-stats          # Web authentication statistics"
    echo

    # Show available filter examples
    success "Filter Examples:"
    echo "  device-type=MacBook Pro (14-inch, 2021)  # Filter by specific device type"
    echo "  device-type=iPad Pro 3rd Gen (11 inch)   # Filter by iPad device type"
    echo "  device-type=Unknown                      # Filter by unknown devices"
    echo "  device-count=1                           # Filter by device count"
    echo "  avg-client-state-duration=0              # Filter by latency duration"
    echo "  total-sessions=30102                     # Filter by session count"
    echo "  run-state-clients=20                     # Filter by running clients"
    echo
}
