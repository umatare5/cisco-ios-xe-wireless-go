#!/usr/bin/env bash

# Cisco WNC Output Formatter Library
# Output formatting and display functions

# Source validation for helper functions
# shellcheck source=./validation.sh
source "$(dirname "${BASH_SOURCE[0]}")/validation.sh"

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
    echo "Available YANG Models (Cisco Wireless):"
    echo "======================================"

    # Early return with jq processing if available
    if is_jq_available; then
        local jq_query='."ietf-restconf:data"."ietf-yang-library:modules-state"'
        jq_query="${jq_query}.module[]"
        jq -r "${jq_query} | [.name, .revision] | join(\"/\")" "$temp_file" | \
        grep "Cisco-IOS-XE-wireless" | sort
        return
    fi

    # Fallback message and raw output
    echo "JSON parsing requires jq. Showing raw output:"
    cat "$temp_file"
}

# Format pretty output for YANG model details
format_yang_model_details_pretty() {
    local temp_file="$1"
    echo "YANG Model Details:"
    echo "=================="

    # Check if the file contains YANG module definition
    if grep -q "^module\|^submodule" "$temp_file" 2>/dev/null; then
        echo "Raw YANG Module Definition:"
        echo "---------------------------"
        cat "$temp_file"
        return
    fi

    # Check if it's JSON and try to parse with jq if available
    if is_jq_available && jq empty "$temp_file" 2>/dev/null; then
        jq -r 'keys[]' "$temp_file" 2>/dev/null | sort | while read -r key; do
            echo "- $key"
        done
        return
    fi

    # Fallback to raw output
    echo "Showing raw output:"
    echo "------------------"
    cat "$temp_file"
}

# Display configuration information
display_configuration() {
    local protocol="$1"
    local controller="$2"
    local output_format="$3"

    if [[ "$output_format" != "raw" ]]; then
        echo "Configuration:"
        echo "============="
        echo "Protocol: $protocol"
        echo "Controller: $controller"
        echo "Output Format: $output_format"
        echo
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
        echo "Fetching YANG model details from: $url"
        echo "Protocol: $protocol"
        echo "Controller: $controller"
        echo "YANG Model: $yang_model"
        echo "Revision: $revision"
        echo "Output Format: $output_format"
        echo "Insecure mode: ${insecure_flag:-disabled}"
        echo "Verbose mode: $verbose"
        echo
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
        echo "Fetching YANG statement details from: $url"
        echo "Protocol: $protocol"
        echo "Controller: $controller"
        echo "YANG Model: $yang_model"
        echo "Identifier: $identifier"
        echo "Output Format: $output_format"
        echo "Insecure mode: ${insecure_flag:-disabled}"
        echo "Verbose mode: $verbose"
        echo
    fi
}

# Display YANG statement configuration summary
display_yang_statement_summary() {
    local response_file="$1"
    local requested_identifier="$2"

    echo "Configuration:"
    echo "  Requested Identifier: $requested_identifier"

    # Extract actual identifier from response
    if [[ -f "$response_file" && -s "$response_file" ]]; then
        if is_jq_available; then
            local actual_identifier
            actual_identifier=$(jq -r '.["ietf-yang-library:module"].name \
// empty' "$response_file" 2>/dev/null || echo "")
            if [[ -n "$actual_identifier" && \
"$actual_identifier" != "null" ]]; then
                echo "  Found Module: $actual_identifier"
            fi
        fi
    fi
    echo
}

# Format YANG statement details with pretty printing
format_yang_statement_details_pretty() {
    local response_file="$1"

    if ! is_jq_available; then
        echo "Warning: jq not found. Displaying raw JSON response:"
        cat "$response_file"
        return
    fi

    echo "YANG Statement Details:"

    # Parse and display module information
    if jq -e '.["ietf-yang-library:module"]' "$response_file" \
>/dev/null 2>&1; then
        display_module_information "$response_file"
    else
        display_no_module_information "$response_file"
    fi
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

    echo "${indent}Module Name: $module_name"
    echo "${indent}Revision: $revision"
    echo "${indent}Namespace: $namespace"
    echo "${indent}Conformance Type: $conformance_type"

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
        echo "${indent}Features:"
        local jq_query='.["ietf-yang-library:module"].feature[]? // empty'
        jq -r "$jq_query" "$response_file" 2>/dev/null | \
        while read -r feature; do
            [[ -n "$feature" ]] && echo "${indent}${indent}- $feature"
        done
    fi
}

# Display module submodules if available
display_module_submodules() {
    local response_file="$1"
    local indent="$2"

    if jq -e '.["ietf-yang-library:module"].submodule' "$response_file" \
>/dev/null 2>&1; then
        echo "${indent}Submodules:"
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

    sub_name=$(echo "$submodule" | jq -r '.name // "N/A"' 2>/dev/null)
    sub_revision=$(echo "$submodule" | \
        jq -r '.revision // "N/A"' 2>/dev/null)
    echo "${indent}${indent}- $sub_name (revision: $sub_revision)"
}

# Display message when no module information found
display_no_module_information() {
    local response_file="$1"
    local indent="  "

    echo "${indent}No module information found in response"
    echo "${indent}Raw response:"
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

    echo "Raw Response:"
    echo "============="
    cat "$temp_file"
    echo ""
    echo "Formatted Output:"
    echo "================="
}

# Show completion message
show_completion() {
    local output_format="$1"

    if [[ "$output_format" != "raw" ]]; then
        echo ""
        echo "Operation completed successfully."
    fi
}
