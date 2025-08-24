#!/usr/bin/env bash

# Access Point Operational Filters - Extraction Functions
# Provides functions to extract filterable fields from AP operational data

# Include guard
if [[ -n "${AP_OPER_FILTER_LOADED:-}" ]]; then
    return 0
fi
readonly AP_OPER_FILTER_LOADED=1

# Extract AP operational filters as JSON
extract_ap_oper_filters_as_json() {
    local file_path="$1"

    jq '{
        "ap_name_mac_map": {
            "wtp-name": [.["Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data"]["ap-name-mac-map"][]."wtp-name" | select(. != null)],
            "wtp-mac": [.["Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data"]["ap-name-mac-map"][]."wtp-mac" | select(. != null)],
            "eth-mac": [.["Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data"]["ap-name-mac-map"][]."eth-mac" | select(. != null)]
        },
        "oper_data": {
            "wtp-mac": [.["Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data"]["oper-data"][]."wtp-mac" | select(. != null)],
            "radio-id": [.["Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data"]["oper-data"][]."radio-id" | select(. != null)],
            "ap-ip-addr": [.["Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data"]["oper-data"][]."ap-ip-data"."ap-ip-addr" | select(. != null)],
            "primary-controller-name": [.["Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data"]["oper-data"][]."ap-prime-info"."primary-controller-name" | select(. != null and . != "")],
            "power-type": [.["Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data"]["oper-data"][]."ap-pow"."power-type" | select(. != null)]
        },
        "radio_neighbor": {
            "ap-mac": [.["Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data"]["ap-radio-neighbor"][]."ap-mac" | select(. != null)],
            "slot-id": [.["Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data"]["ap-radio-neighbor"][]."slot-id" | select(. != null)],
            "primary-channel": [.["Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data"]["ap-radio-neighbor"][]."primary-channel" | select(. != null)]
        }
    }' "$file_path"
}

# Extract AP operational filters as table
extract_ap_oper_filters_as_table() {
    local file_path="$1"

    echo "✓ Success: Access Point Operational Data (access-point-oper-data):"

    # AP Name-MAC Map filters
    local wtp_names wtp_macs
    wtp_names=$(jq -r '.["Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data"]["ap-name-mac-map"][]."wtp-name" | select(. != null)' "$file_path" 2>/dev/null | head -5)
    wtp_macs=$(jq -r '.["Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data"]["ap-name-mac-map"][]."wtp-mac" | select(. != null)' "$file_path" 2>/dev/null | head -5)

    printf "%-20s %-20s %-15s %-20s %-20s %-15s\n" "AP Name" "WTP MAC" "Radio ID" "IP Address" "Controller" "Power Type"
    printf "%-20s %-20s %-15s %-20s %-20s %-15s\n" "-------" "-------" "--------" "----------" "----------" "----------"

    # Collect operational data
    local ip_addrs controllers power_types radio_ids
    ip_addrs=$(jq -r '.["Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data"]["oper-data"][]."ap-ip-data"."ap-ip-addr" | select(. != null)' "$file_path" 2>/dev/null | head -5)
    controllers=$(jq -r '.["Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data"]["oper-data"][]."ap-prime-info"."primary-controller-name" | select(. != null and . != "")' "$file_path" 2>/dev/null | head -5)
    power_types=$(jq -r '.["Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data"]["oper-data"][]."ap-pow"."power-type" | select(. != null)' "$file_path" 2>/dev/null | head -5)
    radio_ids=$(jq -r '.["Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data"]["oper-data"][]."radio-id" | select(. != null)' "$file_path" 2>/dev/null | head -5)

    # Display sample data from each matched AP
    if [[ -n "$wtp_names" ]]; then
        echo "$wtp_names" | while read -r name; do
            local mac ip controller power radio
            mac=$(echo "$wtp_macs" | head -1)
            ip=$(echo "$ip_addrs" | head -1)
            controller=$(echo "$controllers" | head -1)
            power=$(echo "$power_types" | head -1)
            radio=$(echo "$radio_ids" | head -1)

            printf "%-20s %-20s %-15s %-20s %-20s %-15s\n" \
                "${name:-N/A}" "${mac:-N/A}" "${radio:-N/A}" "${ip:-N/A}" "${controller:-N/A}" "${power:-N/A}"
            break
        done
    fi

    echo ""
    echo "✓ Success: AP-Oper Filter Examples:"
    echo "  wtp-name=TEST-AP01                  # Filter by AP name"
    echo "  wtp-mac=28:ac:9e:bb:3c:80           # Filter by WTP MAC address"
    echo "  eth-mac=28:ac:9e:11:48:10           # Filter by Ethernet MAC address"
    echo "  radio-id=8                          # Filter by radio ID"
    echo "  ap-ip-addr=192.168.255.11           # Filter by AP IP address"
    echo "  primary-controller-name=WNC1        # Filter by primary controller"
    echo "  power-type=pwr-src-poe-plus         # Filter by power source type"
    echo "  slot-id=0                           # Filter by radio slot ID"
    echo "  primary-channel=4                   # Filter by primary channel"
    echo ""
}
