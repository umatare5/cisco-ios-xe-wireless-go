#!/usr/bin/env bash

# YANG Model Filter Extraction - Site Configuration Functions
# Provides processing functions for extracting filterable fields from site configuration data

# Include guard
if [[ -n "${SITE_CFG_FILTER_LOADED:-}" ]]; then
    return 0
fi
readonly SITE_CFG_FILTER_LOADED=1

# Extract site-cfg filters in JSON format
extract_site_cfg_filters_as_json() {
    local file_path="$1"

    info "Available site-cfg filterable fields (JSON format):"
    jq '
        .["Cisco-IOS-XE-wireless-site-cfg:site-cfg-data"] as $data |
        {
            "ap-cfg-profiles": [
                $data["ap-cfg-profiles"]["ap-cfg-profile"][]? |
                {
                    "profile-name": .["profile-name"],
                    "description": .["description"]?,
                    "hyperlocation": .["hyperlocation"]?,
                    "location-measurement": .["location-measurement"]?,
                    "rouge-detection": .["rouge-detection"]?,
                    "awips-enabled": .["awips-enabled"]?
                }
            ],
            "site-tag-configs": [
                $data["site-tag-configs"]["site-tag-config"][]? |
                {
                    "site-tag-name": .["site-tag-name"],
                    "description": .["description"]?,
                    "flex-profile": .["flex-profile"]?,
                    "ap-join-profile": .["ap-join-profile"]?,
                    "local-site": .["local-site"]?
                }
            ]
        }
    ' "$file_path" 2>/dev/null || {
        log_error "Failed to extract site-cfg filterable fields"
        return 1
    }
}

# Extract site-cfg filters in table format
extract_site_cfg_filters_as_table() {
    local file_path="$1"

    info "Available site-cfg filterable fields:"
    echo

    # Extract AP Configuration Profiles
    success "AP Configuration Profiles (ap-cfg-profiles):"
    printf "%-25s %-40s %-15s %-15s %-15s\n" "Profile Name" "Description" "Hyperlocation" "Location Meas." "Rogue Detection"
    printf "%-25s %-40s %-15s %-15s %-15s\n" "------------" "-----------" "-------------" "--------------" "---------------"

    jq -r '
        .["Cisco-IOS-XE-wireless-site-cfg:site-cfg-data"]["ap-cfg-profiles"]["ap-cfg-profile"][]? |
        [
            (.["profile-name"] // "N/A"),
            (.["description"] // "N/A"),
            (if .["hyperlocation"] == true then "Enabled" elif .["hyperlocation"] == false then "Disabled" else "N/A" end),
            (if .["location-measurement"] == true then "Enabled" elif .["location-measurement"] == false then "Disabled" else "N/A" end),
            (if .["rouge-detection"] == true then "Enabled" elif .["rouge-detection"] == false then "Disabled" else "N/A" end)
        ] |
        @tsv
    ' "$file_path" 2>/dev/null | while IFS=$'\t' read -r profile_name description hyperlocation location_meas rogue_detection; do
        printf "%-25s %-40s %-15s %-15s %-15s\n" "$profile_name" "$description" "$hyperlocation" "$location_meas" "$rogue_detection"
    done
    echo

    # Extract Site Tag Configurations
    success "Site Tag Configurations (site-tag-configs):"
    printf "%-20s %-40s %-20s %-20s %-15s\n" "Site Tag Name" "Description" "Flex Profile" "AP Join Profile" "Local Site"
    printf "%-20s %-40s %-20s %-20s %-15s\n" "-------------" "-----------" "------------" "---------------" "----------"

    jq -r '
        .["Cisco-IOS-XE-wireless-site-cfg:site-cfg-data"]["site-tag-configs"]["site-tag-config"][]? |
        [
            (.["site-tag-name"] // "N/A"),
            (.["description"] // "N/A"),
            (.["flex-profile"] // "N/A"),
            (.["ap-join-profile"] // "N/A"),
            (if .["local-site"] == true then "Yes" elif .["local-site"] == false then "No" else "N/A" end)
        ] |
        @tsv
    ' "$file_path" 2>/dev/null | while IFS=$'\t' read -r site_tag_name description flex_profile ap_join_profile local_site; do
        printf "%-20s %-40s %-20s %-20s %-15s\n" "$site_tag_name" "$description" "$flex_profile" "$ap_join_profile" "$local_site"
    done
    echo

    # Show available site-cfg filter examples
    success "Site-CFG Filter Examples:"
    echo "  ap-profile-name=labo-common     # Filter by AP profile name"
    echo "  profile-name-ap=default-ap-profile # Filter by AP profile name (alternative)"
    echo "  site-tag-name=labo-site-flex    # Filter by site tag name"
    echo "  site-tag=default-site-tag       # Filter by site tag name (alternative)"
    echo
}
