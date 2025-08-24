#!/usr/bin/env bash

# YANG Model Filter Extraction - WLAN Configuration Functions
# Provides processing functions for extracting filterable fields from WLAN configuration data

# Include guard
if [[ -n "${WLAN_CFG_FILTER_LOADED:-}" ]]; then
    return 0
fi
readonly WLAN_CFG_FILTER_LOADED=1

# Extract wlan-cfg filters in JSON format
extract_wlan_cfg_filters_as_json() {
    local file_path="$1"

    info "Available filterable fields (JSON format):"
    jq '
        .["Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data"] as $data |
        {
            "wlan-cfg-entries": [
                $data["wlan-cfg-entries"]["wlan-cfg-entry"][]? |
                {
                    "wlan-id": .["wlan-id"],
                    "profile-name": .["profile-name"],
                    "ssid": .["apf-vap-id-data"]["ssid"]?,
                    "wpa2-enabled": .["wpa2-enabled"]?,
                    "wpa3-enabled": .["wpa3-enabled"]?,
                    "auth-key-mgmt-psk": .["auth-key-mgmt-psk"]?,
                    "auth-key-mgmt-dot1x": .["auth-key-mgmt-dot1x"]?,
                    "ft-mode": .["ft-mode"]?,
                    "load-balance": .["load-balance"]?
                }
            ],
            "wlan-policies": [
                $data["wlan-policies"]["wlan-policy"][]? |
                {
                    "policy-profile-name": .["policy-profile-name"],
                    "status": .["status"]?,
                    "interface-name": .["interface-name"]?,
                    "central-switching": .["wlan-switching-policy"]["central-switching"]?,
                    "central-authentication": .["wlan-switching-policy"]["central-authentication"]?,
                    "central-dhcp": .["wlan-switching-policy"]["central-dhcp"]?
                }
            ],
            "policy-list-entries": [
                $data["policy-list-entries"]["policy-list-entry"][]? |
                {
                    "tag-name": .["tag-name"],
                    "description": .["description"]?
                }
            ]
        }
    ' "$file_path" 2>/dev/null || {
        log_error "Failed to extract filterable fields"
        return 1
    }
}

# Extract wlan-cfg filters in table format
extract_wlan_cfg_filters_as_table() {
    local file_path="$1"

    info "Available filterable fields:"
    echo

    # Extract WLAN Configuration Entries
    success "WLAN Configuration Entries (wlan-cfg-entries):"
    printf "%-15s %-20s %-25s %-12s %-12s\n" "WLAN-ID" "Profile Name" "SSID" "WPA2" "WPA3"
    printf "%-15s %-20s %-25s %-12s %-12s\n" "-------" "------------" "----" "----" "----"

    jq -r '
        .["Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data"]["wlan-cfg-entries"]["wlan-cfg-entry"][]? |
        [
            (.["wlan-id"] // "N/A"),
            (.["profile-name"] // "N/A"),
            (.["apf-vap-id-data"]["ssid"] // "N/A"),
            (if .["wpa2-enabled"] == true then "Yes" elif .["wpa2-enabled"] == false then "No" else "N/A" end),
            (if .["wpa3-enabled"] == true then "Yes" elif .["wpa3-enabled"] == false then "No" else "N/A" end)
        ] |
        @tsv
    ' "$file_path" 2>/dev/null | while IFS=$'\t' read -r wlan_id profile_name ssid wpa2 wpa3; do
        printf "%-15s %-20s %-25s %-12s %-12s\n" "$wlan_id" "$profile_name" "$ssid" "$wpa2" "$wpa3"
    done
    echo

    # Extract WLAN Policies
    success "WLAN Policies (wlan-policies):"
    printf "%-25s %-10s %-20s %-15s\n" "Policy Profile Name" "Status" "Interface Name" "Central Switch"
    printf "%-25s %-10s %-20s %-15s\n" "-------------------" "------" "--------------" "--------------"

    jq -r '
        .["Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data"]["wlan-policies"]["wlan-policy"][]? |
        [
            (.["policy-profile-name"] // "N/A"),
            (if .["status"] == true then "Active" elif .["status"] == false then "Inactive" else "N/A" end),
            (.["interface-name"] // "N/A"),
            (if .["wlan-switching-policy"]["central-switching"] == true then "Yes" elif .["wlan-switching-policy"]["central-switching"] == false then "No" else "N/A" end)
        ] |
        @tsv
    ' "$file_path" 2>/dev/null | while IFS=$'\t' read -r policy_name status interface_name central_switch; do
        printf "%-25s %-10s %-20s %-15s\n" "$policy_name" "$status" "$interface_name" "$central_switch"
    done
    echo

    # Extract Policy List Entries
    success "Policy List Entries (policy-list-entries):"
    printf "%-25s %-50s\n" "Tag Name" "Description"
    printf "%-25s %-50s\n" "--------" "-----------"

    jq -r '
        .["Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data"]["policy-list-entries"]["policy-list-entry"][]? |
        [
            (.["tag-name"] // "N/A"),
            (.["description"] // "N/A")
        ] |
        @tsv
    ' "$file_path" 2>/dev/null | while IFS=$'\t' read -r tag_name description; do
        printf "%-25s %-50s\n" "$tag_name" "$description"
    done
    echo

    # Show available filter examples
    success "Filter Examples:"
    echo "  wlan-id=1                    # Filter by WLAN ID"
    echo "  profile-name=labo2            # Filter by profile name"
    echo "  ssid=labo1               # Filter by SSID"
    echo "  wpa3-enabled=true           # Filter by WPA3 enabled"
    echo "  policy-profile-name=default # Filter by policy profile"
    echo "  status=true                 # Filter by active status"
    echo
}
