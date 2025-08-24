#!/usr/bin/env bash

# YANG Model Filter Extraction - AP Configuration Functions
# Provides processing functions for extracting filterable fields from AP configuration data

# Include guard
if [[ -n "${AP_CFG_FILTER_LOADED:-}" ]]; then
    return 0
fi
readonly AP_CFG_FILTER_LOADED=1

# Extract ap-cfg filters in JSON format
extract_ap_cfg_filters_as_json() {
    local file_path="$1"

    info "Available AP configuration filterable fields (JSON format):"
    jq '
        .["Cisco-IOS-XE-wireless-ap-cfg:ap-cfg-data"] as $data |
        {
            "ap-tags": [
                $data["ap-tags"]["ap-tag"][]? |
                {
                    "ap-mac": .["ap-mac"],
                    "policy-tag": .["policy-tag"]?,
                    "site-tag": .["site-tag"]?,
                    "rf-tag": .["rf-tag"]?
                }
            ],
            "tag-source-priority-configs": [
                $data["tag-source-priority-configs"]["tag-source-priority-config"][]? |
                {
                    "priority": .["priority"],
                    "tag-src": .["tag-src"]?
                }
            ]
        }
    ' "$file_path" 2>/dev/null || {
        log_error "Failed to extract ap-cfg filterable fields"
        return 1
    }
}

# Extract ap-cfg filters in table format
extract_ap_cfg_filters_as_table() {
    local file_path="$1"

    info "Available AP configuration filterable fields:"
    echo

    # Extract AP Tags
    success "AP Tags (ap-tags):"
    printf "%-20s %-20s %-20s %-15s\n" "AP MAC Address" "Policy Tag" "Site Tag" "RF Tag"
    printf "%-20s %-20s %-20s %-15s\n" "--------------" "----------" "--------" "------"

    jq -r '
        .["Cisco-IOS-XE-wireless-ap-cfg:ap-cfg-data"]["ap-tags"]["ap-tag"][]? |
        [
            (.["ap-mac"] // "N/A"),
            (.["policy-tag"] // "N/A"),
            (.["site-tag"] // "N/A"),
            (.["rf-tag"] // "N/A")
        ] |
        @tsv
    ' "$file_path" 2>/dev/null | while IFS=$'\t' read -r ap_mac policy_tag site_tag rf_tag; do
        printf "%-20s %-20s %-20s %-15s\n" "$ap_mac" "$policy_tag" "$site_tag" "$rf_tag"
    done
    echo

    # Extract Tag Source Priority Configurations
    success "Tag Source Priority Configurations (tag-source-priority-configs):"
    printf "%-10s %-30s\n" "Priority" "Tag Source"
    printf "%-10s %-30s\n" "--------" "----------"

    jq -r '
        .["Cisco-IOS-XE-wireless-ap-cfg:ap-cfg-data"]["tag-source-priority-configs"]["tag-source-priority-config"][]? |
        [
            (.["priority"] // "N/A"),
            (.["tag-src"] // "N/A")
        ] |
        @tsv
    ' "$file_path" 2>/dev/null | while IFS=$'\t' read -r priority tag_src; do
        printf "%-10s %-30s\n" "$priority" "$tag_src"
    done
    echo

    # Show available ap-cfg filter examples
    success "AP-CFG Filter Examples:"
    echo "  ap-mac=28:ac:9e:11:48:10       # Filter by AP MAC address"
    echo "  policy-tag=labo-wlan-flex      # Filter by policy tag"
    echo "  site-tag=labo-site-flex        # Filter by site tag"
    echo "  rf-tag=labo-outside            # Filter by RF tag"
    echo "  priority=0                     # Filter by priority"
    echo "  tag-src=tag-source-static      # Filter by tag source"
    echo
}
