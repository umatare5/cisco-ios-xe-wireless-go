#!/usr/bin/env bash

# AP Global Operational Filters - Extraction Functions
# Provides functions to extract filterable fields from AP global operational data

# Include guard
if [[ -n "${AP_GLOBAL_OPER_FILTER_LOADED:-}" ]]; then
    return 0
fi
readonly AP_GLOBAL_OPER_FILTER_LOADED=1

# Extract AP global operational filters as JSON
extract_ap_global_oper_filters_as_json() {
    local file_path="$1"

    jq '{
        "ap_history": {
            "ap-name": [.["Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data"]["ap-history"][]."ap-name" | select(. != null)],
            "wtp-mac": [.["Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data"]["ap-history"][]."wtp-mac" | select(. != null)],
            "ethernet-mac": [.["Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data"]["ap-history"][]."ethernet-mac" | select(. != null)]
        },
        "ap_join_stats": {
            "wtp-mac": [.["Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data"]["ap-join-stats"][]."wtp-mac" | select(. != null)],
            "ap-name": [.["Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data"]["ap-join-stats"][]."ap-join-info"."ap-name" | select(. != null)],
            "ap-ip-addr": [.["Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data"]["ap-join-stats"][]."ap-join-info"."ap-ip-addr" | select(. != null)],
            "is-joined": [.["Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data"]["ap-join-stats"][]."ap-join-info"."is-joined" | select(. != null)],
            "ap-disconnect-reason": [.["Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data"]["ap-join-stats"][]."ap-disconnect-reason" | select(. != null and . != "")],
            "reboot-reason": [.["Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data"]["ap-join-stats"][]."reboot-reason" | select(. != null and . != "")],
            "last-join-failure-type": [.["Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data"]["ap-join-stats"][]."ap-join-info"."last-join-failure-type" | select(. != null and . != "")],
            "last-config-failure-type": [.["Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data"]["ap-join-stats"][]."ap-join-info"."last-config-failure-type" | select(. != null and . != "")]
        }
    }' "$file_path"
}

# Extract AP global operational filters as table
extract_ap_global_oper_filters_as_table() {
    local file_path="$1"

    echo "✓ Success: AP Global Operational Data (ap-global-oper-data):"

    # Extract AP history and join stats data
    local ap_names wtp_macs ethernet_macs ip_addrs
    ap_names=$(jq -r '.["Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data"]["ap-join-stats"][]."ap-join-info"."ap-name" | select(. != null)' "$file_path" 2>/dev/null | head -5)
    wtp_macs=$(jq -r '.["Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data"]["ap-join-stats"][]."wtp-mac" | select(. != null)' "$file_path" 2>/dev/null | head -5)
    ethernet_macs=$(jq -r '.["Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data"]["ap-join-stats"][]."ap-join-info"."ap-ethernet-mac" | select(. != null)' "$file_path" 2>/dev/null | head -5)
    ip_addrs=$(jq -r '.["Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data"]["ap-join-stats"][]."ap-join-info"."ap-ip-addr" | select(. != null)' "$file_path" 2>/dev/null | head -5)

    printf "%-20s %-20s %-20s %-15s %-10s %-10s %-30s\n" "AP Name" "WTP MAC" "Ethernet MAC" "IP Address" "Joined" "Join Reqs" "Disconnect Reason"
    printf "%-20s %-20s %-20s %-15s %-10s %-10s %-30s\n" "-------" "-------" "-------" "----------" "------" "---------" "-----------------"

    # Collect join stats and status information
    local join_statuses disconnect_reasons join_reqs
    join_statuses=$(jq -r '.["Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data"]["ap-join-stats"][]."ap-join-info"."is-joined" | select(. != null)' "$file_path" 2>/dev/null | head -5)
    disconnect_reasons=$(jq -r '.["Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data"]["ap-join-stats"][]."ap-disconnect-reason" | select(. != null and . != "")' "$file_path" 2>/dev/null | head -5)
    join_reqs=$(jq -r '.["Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data"]["ap-join-stats"][]."ap-join-info"."num-join-req-recvd" | select(. != null)' "$file_path" 2>/dev/null | head -5)

    # Display sample data from each matched AP
    if [[ -n "$ap_names" ]]; then
        echo "$ap_names" | while read -r name; do
            local mac eth_mac ip joined reason reqs
            mac=$(echo "$wtp_macs" | head -1)
            eth_mac=$(echo "$ethernet_macs" | head -1)
            ip=$(echo "$ip_addrs" | head -1)
            joined=$(echo "$join_statuses" | head -1)
            reason=$(echo "$disconnect_reasons" | head -1)
            reqs=$(echo "$join_reqs" | head -1)

            printf "%-20s %-20s %-20s %-15s %-10s %-10s %-30s\n" \
                "${name:-N/A}" "${mac:-N/A}" "${eth_mac:-N/A}" "${ip:-N/A}" "${joined:-N/A}" "${reqs:-N/A}" "${reason:-N/A}"
            break
        done
    fi

    echo ""
    echo "✓ Success: AP-Global-Oper Filter Examples:"
    echo "  ap-name=TEST-AP01                            # Filter by AP name"
    echo "  wtp-mac=28:ac:9e:bb:3c:80                    # Filter by WTP MAC address"
    echo "  ethernet-mac=28:ac:9e:11:48:10               # Filter by Ethernet MAC address"
    echo "  ap-ip-addr=192.168.255.11                    # Filter by AP IP address"
    echo "  is-joined=true                               # Filter by join status (true/false)"
    echo "  ap-disconnect-reason=Wtp reset config cmd sent  # Filter by disconnect reason"
    echo "  reboot-reason=ap-reboot-reason-reboot-cmd    # Filter by reboot reason"
    echo "  last-join-failure-type=jf-none               # Filter by last join failure type"
    echo "  last-config-failure-type=cf-none             # Filter by last config failure type"
    echo ""
}
