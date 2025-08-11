#!/usr/bin/env bash

# Cisco WNC Output Formatter Library
# Output formatting and display functions (generic, non-domain-specific)

# Source validation for helper functions
source "$(dirname "${BASH_SOURCE[0]}")/../utils/validation.sh"

# Format JSON output using jq if available
format_json_output() {
    local temp_file="$1"

    # Early return with jq formatting if available
    if is_jq_available; then
        jq '.' "$temp_file"
        return
    fi

    # Fallback to raw output
    cat "$temp_file"
}

# Format pretty output for YANG models list
format_yang_models_pretty() {
    local temp_file="$1"
    # Header for pretty list
    printf "%s\n" "Available YANG Models (Cisco Wireless)"
    printf "%s\n" "--------------------------------------"

    # Early return with jq processing if available
    if is_jq_available; then
        local jq_query='."ietf-restconf:data"."ietf-yang-library:modules-state"'
        jq_query="${jq_query}.module[]"
        jq -r "${jq_query} | [.name, .revision] | join(\"/\")" "$temp_file" | \
        grep "Cisco-IOS-XE-wireless" | sort
        return
    fi

    # Fallback message and raw output
    printf "%s\n" "JSON parsing requires jq. Showing raw output:"
    cat "$temp_file"
}

# Format pretty output for YANG model details
format_yang_model_details_pretty() {
    local temp_file="$1"
    printf "%s\n" "YANG Model Details:"
    printf "%s\n" "------------------"

    # Check if the file contains YANG module definition
    if grep -q "^module\|^submodule" "$temp_file" 2>/dev/null; then
    printf '%s\n' "Raw YANG Module Definition:"
    printf '%s\n' "---------------------------"
        cat "$temp_file"
        return
    fi

    # Check if it's JSON and try to parse with jq if available
    if is_jq_available && jq empty "$temp_file" 2>/dev/null; then
        jq -r 'keys[]' "$temp_file" 2>/dev/null | sort | while read -r key; do
            printf '%s\n' "- $key"
        done
        return
    fi

    # Fallback to raw output
    printf "%s\n" "Showing raw output:"
    printf "%s\n" "------------------"
    cat "$temp_file"
}

# Display configuration information
display_configuration() {
    local protocol="$1"
    local controller="$2"
    local output_format="$3"

    if [[ "$output_format" != "raw" ]]; then
    printf "%s\n" "Configuration:"
    printf "%s\n" "-------------"
    printf "%s\n" "Protocol: $protocol"
    printf "%s\n" "Controller: $controller"
    printf "%s\n\n" "Output Format: $output_format"
    fi
}

# Display YANG model specific configuration
display_yang_model_configuration() {
    local protocol="$1"
    local controller="$2"
    local yang_model="$3"
    local revision="$4"
    local output_format="$5"
    local insecure_flag="$6"
    local verbose="$7"
    local url="$8"

    if [[ "$output_format" != "raw" ]]; then
    printf "%s\n" "Fetching YANG model details from: $url"
    printf "%s\n" "Protocol: $protocol"
    printf "%s\n" "Controller: $controller"
    printf "%s\n" "YANG Model: $yang_model"
    printf "%s\n" "Revision: $revision"
    printf "%s\n" "Output Format: $output_format"
    printf "%s\n" "Insecure mode: ${insecure_flag:-disabled}"
    printf "%s\n\n" "Verbose mode: $verbose"
    fi
}

# Display YANG statement specific configuration
display_yang_statement_configuration() {
    local protocol="$1"
    local controller="$2"
    local yang_model="$3"
    local identifier="$4"
    local output_format="$5"
    local insecure_flag="$6"
    local verbose="$7"
    local url="$8"

    if [[ "$output_format" != "raw" ]]; then
    printf "%s\n" "Fetching YANG statement details from: $url"
    printf "%s\n" "Protocol: $protocol"
    printf "%s\n" "Controller: $controller"
    printf "%s\n" "YANG Model: $yang_model"
    printf "%s\n" "Identifier: $identifier"
    printf "%s\n" "Output Format: $output_format"
    printf "%s\n" "Insecure mode: ${insecure_flag:-disabled}"
    printf "%s\n\n" "Verbose mode: $verbose"
    fi
}

# Display YANG statement configuration summary
display_yang_statement_summary() {
    local response_file="$1"
    local requested_identifier="$2"

    printf "%s\n" "Configuration:"
    printf "%s\n" "  Requested Identifier: $requested_identifier"

    # Extract actual identifier from response
    if [[ -f "$response_file" && -s "$response_file" ]]; then
        if is_jq_available; then
            local actual_identifier jq_filter
            jq_filter='.["ietf-yang-library:module"].name // empty'
            actual_identifier=$(jq -r "$jq_filter" "$response_file" 2>/dev/null || printf "")
            if [[ -n "$actual_identifier" && \
"$actual_identifier" != "null" ]]; then
                printf "%s\n" "  Found Module: $actual_identifier"
            fi
        fi
    fi
    printf '\n'
}

# Check if response contains YANG library module information
has_yang_library_module() {
    local response_file="$1"
    jq -e '.["ietf-yang-library:module"]' "$response_file" >/dev/null 2>&1
}

# Format YANG statement details with pretty printing
format_yang_statement_details_pretty() {
    local response_file="$1"

    if ! is_jq_available; then
        printf '%s\n' "Warning: jq not found. Displaying raw JSON response:"
        cat "$response_file"
        return
    fi

    printf "%s\n" "YANG Statement Details:"

    # Parse and display module information based on response structure
    if has_yang_library_module "$response_file"; then
        display_module_information "$response_file"
        return
    fi

    # Fallback to basic information display
    display_no_module_information "$response_file"
}

# Display module information from response
display_module_information() {
    local response_file="$1"
    local indent="  "

    local module_name revision namespace conformance_type
    module_name=$(extract_jq_field "$response_file" '.name')
    revision=$(extract_jq_field "$response_file" '.revision')
    namespace=$(extract_jq_field "$response_file" '.namespace')
    conformance_type=$(extract_jq_field "$response_file" '."conformance-type"')

    printf '%s\n' "${indent}Module Name: $module_name"
    printf '%s\n' "${indent}Revision: $revision"
    printf '%s\n' "${indent}Namespace: $namespace"
    printf '%s\n' "${indent}Conformance Type: $conformance_type"

    display_module_features "$response_file" "$indent"
    display_module_submodules "$response_file" "$indent"
}

# Extract field from module using jq
extract_jq_field() {
    local response_file="$1"
    local field="$2"
    local jq_query=".\"ietf-yang-library:module\"${field} // \"N/A\""
    jq -r "$jq_query" "$response_file" 2>/dev/null
}

# Display module features if available
display_module_features() {
    local response_file="$1"
    local indent="$2"

    if jq -e '.["ietf-yang-library:module"].feature' "$response_file" \
>/dev/null 2>&1; then
    printf '%s\n' "${indent}Features:"
        local jq_query='.["ietf-yang-library:module"].feature[]? // empty'
        jq -r "$jq_query" "$response_file" 2>/dev/null | \
        while read -r feature; do
            if [[ -n "$feature" ]]; then
                printf '%s\n' "${indent}${indent}- $feature"
            fi
        done
    fi
}

# Display module submodules if available
display_module_submodules() {
    local response_file="$1"
    local indent="$2"

    if jq -e '.["ietf-yang-library:module"].submodule' "$response_file" \
>/dev/null 2>&1; then
    printf '%s\n' "${indent}Submodules:"
        local jq_query='.["ietf-yang-library:module"].submodule[]? // empty'
        jq -c "$jq_query" "$response_file" 2>/dev/null | \
        while read -r submodule; do
            if [[ -n "$submodule" ]]; then
                display_submodule_info "$submodule" "$indent"
            fi
        done
    fi
}

# Display individual submodule information
display_submodule_info() {
    local submodule="$1"
    local indent="$2"
    local sub_name sub_revision

    sub_name=$(printf '%s' "$submodule" | jq -r '.name // "N/A"' 2>/dev/null)
    sub_revision=$(printf '%s' "$submodule" | \
        jq -r '.revision // "N/A"' 2>/dev/null)
    printf '%s\n' "${indent}${indent}- $sub_name (revision: $sub_revision)"
}

# Display message when no module information found
display_no_module_information() {
    local response_file="$1"
    local indent="  "

    printf "%s\n" "${indent}No module information found in response"
    printf "%s\n" "${indent}Raw response:"
    jq '.' "$response_file" 2>/dev/null || cat "$response_file"
}

# Show raw response if verbose mode is enabled
show_raw_response() {
    local temp_file="$1"
    local verbose="$2"
    local output_format="$3"

    # Early return if not verbose or raw format
    if [[ "$verbose" != true || "$output_format" == "raw" ]]; then
        return 0
    fi

    printf "%s\n" "Raw Response:"
    printf "%s\n" "-------------"
    cat "$temp_file"
    printf "\n"
    printf "%s\n" "Formatted Output:"
    printf "%s\n" "-----------------"
}

# Show completion message
show_completion() {
    local output_format="$1"

    if [[ "$output_format" != "raw" ]]; then
    printf "\n"
    printf "%s\n" "Operation completed successfully."
    fi
}
