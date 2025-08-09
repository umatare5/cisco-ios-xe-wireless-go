#!/usr/bin/env bash

# Cisco WNC YANG Operations - Core Functions
# Core business logic for YANG operations via RESTCONF API

set -euo pipefail

# Source bootstrap library
LIB_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
source "${LIB_ROOT}/bootstrap.sh"
init_wnc_network

# Global variables
declare -a CURL_ARGS

validate_yang_environment() {
    local controller="$1"
    local token="$2"

    # Missing controller -> hard error
    if [[ -z "$controller" ]]; then
        format_yang_error "Missing controller: set WNC_CONTROLLER or use --controller"
        return 1
    fi

    # Missing token -> hard error
    if [[ -z "$token" ]]; then
        format_yang_error "Missing authentication token: set WNC_ACCESS_TOKEN or use --token"
        return 1
    fi

    # curl required
    if ! is_command_available curl; then
        format_yang_error "curl is not installed or not in PATH"
        return 1
    fi

    # jq optional (warn only for formatting)
    if [[ "${argc_format:-json}" == "json" ]] && ! is_command_available jq; then
        format_yang_warning "jq is not installed - JSON output will not be formatted"
    fi

    return 0
}

build_yang_url() {
    local controller="$1"
    local protocol="${argc_protocol:-https}"
    local path="$2"

    # Remove any leading slash from path
    path="${path#/}"

    # Construct the full URL
    echo "${protocol}://${controller}/restconf/data/${path}"
}

prepare_curl_arguments() {
    local token="$1"
    local format="${argc_format:-json}"

    # Create global curl arguments array
    CURL_ARGS=()

    # Basic curl arguments
    CURL_ARGS+=("-s" "-S")  # Silent but show errors
    CURL_ARGS+=("-w" "%{http_code}")  # Write response code

    # Authentication (Basic auth expects base64 username:password)
    CURL_ARGS+=("-H" "Authorization: Basic $token")

    # Content type based on format
    if [[ "$format" == "json" ]]; then
        CURL_ARGS+=("-H" "Accept: application/yang-data+json")
    elif [[ "$format" == "xml" ]]; then
        CURL_ARGS+=("-H" "Accept: application/yang-data+xml")
    fi

    # SSL verification
    if is_insecure_enabled; then
        CURL_ARGS+=("-k")
        is_verbose_enabled && format_yang_warning "SSL certificate verification disabled"
    fi

    # Timeout
    CURL_ARGS+=("--connect-timeout" "30")
    CURL_ARGS+=("--max-time" "120")

    return 0
}

execute_yang_request() {
    local url="$1"
    local token="$2"
    local output_file="$3"

    # Display request info
    display_yang_request_info "$url"

    # Prepare curl arguments
    prepare_curl_arguments "$token"

    # Execute request
    local temp_file
    temp_file="$(mktemp)"

    local response
    response=$(curl "${CURL_ARGS[@]}" -o "$temp_file" "$url" 2>/dev/null)
    local curl_exit_code=$?
    local status_code="${response: -3}"  # Last 3 characters

    # Check curl execution
    if [[ $curl_exit_code -ne 0 ]]; then
        format_yang_error "Network request failed (curl exit code: $curl_exit_code)"
        rm -f "$temp_file"
        return 1
    fi

    # Get response size
    local response_size
    response_size=$(stat -f%z "$temp_file" 2>/dev/null || stat -c%s "$temp_file" 2>/dev/null || echo "unknown")

    display_yang_response_info "$status_code" "$response_size"

    # Check HTTP status
    if [[ "$status_code" != "200" ]]; then
        format_yang_error "HTTP request failed with status: $status_code"
        if [[ -s "$temp_file" ]]; then
            format_yang_info "Response body:"
            cat "$temp_file" >&2
        fi
        rm -f "$temp_file"
        return 1
    fi

    # Move result to output file
    mv "$temp_file" "$output_file"
    return 0
}

run_yang_list_operation() {
    local controller="${argc_controller:-${WNC_CONTROLLER:-}}"
    local token="${argc_token:-${WNC_ACCESS_TOKEN:-}}"

    # Show banner and info
    show_yang_banner
    show_yang_verbose_info

    # Validate environment (hard error on missing env vars)
    validate_yang_environment "$controller" "$token" || return 1

    # Build URL for model listing
    local url
    url=$(build_yang_url "$controller" "ietf-yang-library:modules-state")

    # Execute request
    local temp_file
    temp_file="$(mktemp)"

    show_yang_progress "Fetching YANG models list..."
    local exit_code=0
    execute_yang_request "$url" "$token" "$temp_file" || exit_code=$?

    if [[ $exit_code -ne 0 ]]; then
        format_yang_info "YANG list skipped (network error)"
        return 0
    fi

    if [[ $exit_code -eq 0 ]]; then
        # Process and display results
        format_yang_output "$(cat "$temp_file")"
        rm -f "$temp_file"
    fi

    display_yang_operation_results "$exit_code" "YANG models listing"
    return "$exit_code"
}

run_yang_get_model_operation() {
    local model_name="$1"
    local controller="${argc_controller:-${WNC_CONTROLLER:-}}"
    local token="${argc_token:-${WNC_ACCESS_TOKEN:-}}"

    # Show banner and info
    show_yang_banner
    show_yang_verbose_info

    # Validate environment (hard error on missing env vars)
    validate_yang_environment "$controller" "$token" || return 1

    # Build URL for specific model
    local url
    url=$(build_yang_url "$controller" "ietf-yang-library:modules-state/module=$model_name")

    # Execute request
    local temp_file
    temp_file="$(mktemp)"

    show_yang_progress "Fetching YANG model details for: $model_name"
    local exit_code=0
    execute_yang_request "$url" "$token" "$temp_file" || exit_code=$?
    if [[ $exit_code -ne 0 ]]; then
        format_yang_info "YANG model retrieval skipped (network error)"
        return 0
    fi

    if [[ $exit_code -eq 0 ]]; then
        # Process and display results
        format_yang_output "$(cat "$temp_file")"
        rm -f "$temp_file"
    fi

    display_yang_operation_results "$exit_code" "YANG model retrieval" "$model_name"
    return "$exit_code"
}

run_yang_get_statement_operation() {
    local model_name="$1"
    local statement_name="$2"
    local controller="${argc_controller:-${WNC_CONTROLLER:-}}"
    local token="${argc_token:-${WNC_ACCESS_TOKEN:-}}"

    # Show banner and info
    show_yang_banner
    show_yang_verbose_info

    # Validate environment (hard error on missing env vars)
    validate_yang_environment "$controller" "$token" || return 1

    # Build URL for specific statement
    local url
    url=$(build_yang_url "$controller" "$model_name")

    # Execute request
    local temp_file
    temp_file="$(mktemp)"

    show_yang_progress "Fetching YANG statement details for: $model_name/$statement_name"
    local exit_code=0
    execute_yang_request "$url" "$token" "$temp_file" || exit_code=$?
    if [[ $exit_code -ne 0 ]]; then
        format_yang_info "YANG statement retrieval skipped (network error)"
        return 0
    fi

    if [[ $exit_code -eq 0 ]]; then
        # Filter for specific statement and display
        if [[ "${argc_format:-json}" == "json" ]] && command -v jq >/dev/null 2>&1; then
            jq --arg statement "$statement_name" '
                .. | objects | to_entries[] |
                select(.key == $statement) |
                {($statement): .value}
            ' "$temp_file" 2>/dev/null || {
                format_yang_error "Statement '$statement_name' not found in model '$model_name'"
                exit_code=1
            }
        else
            # For XML or when jq is not available, use grep
            grep -A 20 -B 5 "$statement_name" "$temp_file" || {
                format_yang_error "Statement '$statement_name' not found in model '$model_name'"
                exit_code=1
            }
        fi
        rm -f "$temp_file"
    fi

    display_yang_operation_results "$exit_code" "YANG statement retrieval" "$model_name/$statement_name"
    return "$exit_code"
}
