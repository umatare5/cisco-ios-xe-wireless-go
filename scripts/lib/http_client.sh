#!/usr/bin/env bash

# Cisco WNC HTTP Client Library
# HTTP communication and request handling functions

# Source constants
# shellcheck source=./constants.sh
source "$(dirname "${BASH_SOURCE[0]}")/constants.sh"

# Execute curl request with standard WNC API headers
execute_curl_request() {
    local url="$1"
    local token="$2"
    local insecure_flag="$3"
    local output_file="$4"
    local accept_type_yang
    local error_network_failure
    local exit_network_error

    accept_type_yang="$(get_yang_accept_header)"
    error_network_failure="$(get_network_error_message)"
    exit_network_error="$(get_exit_network_error)"

    local curl_args=(-sS -X 'GET')
    [[ -n "$insecure_flag" ]] && curl_args+=("$insecure_flag")

    # Early return for successful curl request
    if curl "${curl_args[@]}" \
        -H "accept: $accept_type_yang" \
        -H "Authorization: Basic $token" \
        "$url" \
        -o "$output_file" 2>/dev/null; then
        return 0
    fi

    # Curl failed - show error
    echo "$error_network_failure" >&2
    echo "Please check your connection settings and authentication token." >&2
    exit "$exit_network_error"
}

# Build base URL for controller
build_base_url() {
    local protocol="$1"
    local controller="$2"
    echo "${protocol}://${controller}"
}

# Build URL for YANG model list endpoint
build_yang_models_url() {
    local protocol="$1"
    local controller="$2"
    local base_url
    local restconf_data_path
    local yang_library_query="?fields=ietf-yang-library:modules-state/module"

    base_url="$(build_base_url "$protocol" "$controller")"
    restconf_data_path="$(get_restconf_data_path)"
    echo "${base_url}${restconf_data_path}${yang_library_query}"
}

# Build URL for YANG model details endpoint
build_yang_model_details_url() {
    local protocol="$1"
    local controller="$2"
    local yang_model="$3"
    local revision="$4"
    local base_url
    local restconf_modules_path

    base_url="$(build_base_url "$protocol" "$controller")"
    restconf_modules_path="$(get_restconf_modules_path)"
    echo "${base_url}${restconf_modules_path}/${yang_model}/${revision}"
}

# Build URL for YANG statement details endpoint
build_yang_statement_url() {
    local protocol="$1"
    local controller="$2"
    local yang_model="$3"
    local identifier="$4"
    local base_url
    local restconf_data_path

    base_url="$(build_base_url "$protocol" "$controller")"
    restconf_data_path="$(get_restconf_data_path)"
    echo "${base_url}${restconf_data_path}/${yang_model}:${identifier}"
}

# Test connection to controller
test_connection() {
    local protocol="$1"
    local controller="$2"
    local token="$3"
    local insecure_flag="$4"
    local base_url
    local restconf_data_path
    local accept_type_yang
    local error_network_failure
    local exit_error

    base_url="$(build_base_url "$protocol" "$controller")"
    restconf_data_path="$(get_restconf_data_path)"
    accept_type_yang="$(get_yang_accept_header)"
    error_network_failure="$(get_network_error_message)"
    exit_error="$(get_exit_error)"

    local test_url="${base_url}${restconf_data_path}"

    # Early return for successful connection test
    if curl -sS -X 'GET' "${insecure_flag}" \
        -H "accept: $accept_type_yang" \
        -H "Authorization: Basic $token" \
        "$test_url" \
        -o /dev/null 2>/dev/null; then
        return 0
    fi

    # Connection failed - show error
    echo "$error_network_failure" >&2
    echo "Please check your controller address and network connectivity." >&2
    exit "$exit_error"
}
