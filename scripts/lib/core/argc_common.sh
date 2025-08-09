#!/usr/bin/env bash

# Cisco WNC argc Common Library - argc-specific shared functions

set -euo pipefail

# Common argc-based validation and setup
validate_argc_yang_environment() {
    # shellcheck disable=SC2154
    local token="${argc_token:-${WNC_ACCESS_TOKEN:-}}"

    [[ -n "$token" ]] || {
        show_error "No authentication token provided. Use --token or set WNC_ACCESS_TOKEN"
        exit 1
    }

    command_exists curl || { show_error "curl is required"; exit 1; }
    command_exists jq || { show_error "jq is required"; exit 1; }

    export VALIDATED_TOKEN="$token"
}

# Build YANG API URL from argc variables
build_argc_yang_url() {
    local endpoint="$1"
    # shellcheck disable=SC2154
    local controller="${argc_controller:-}" # require explicit controller
    [[ -n "$controller" ]] || { show_error "No controller specified. Use --controller or set WNC_CONTROLLER"; exit 1; }
    # shellcheck disable=SC2154
    local protocol="${argc_protocol:-https}"

    echo "${protocol}://${controller}/restconf/data/${endpoint}"
}

# Execute curl with argc-based options
execute_argc_curl() {
    local url="$1"
    local output_format="${2:-pretty}"  # pretty, raw, file
    local output_file="${3:-}"

    # shellcheck disable=SC2154
    local insecure="${argc_insecure:-false}"
    local curl_opts=()
    [[ "$insecure" == "true" ]] && curl_opts+=("-k")

    if [[ "$output_format" == "file" && -n "$output_file" ]]; then
    curl -s -H "Authorization: Basic $VALIDATED_TOKEN" \
             "${curl_opts[@]}" "$url" -o "$output_file"
    else
    curl -s -H "Authorization: Basic $VALIDATED_TOKEN" \
             "${curl_opts[@]}" "$url"
    fi
}

# argc predicate helper functions
is_enabled() {
    [[ "${1:-0}" == "1" ]]
}

is_true() {
    [[ "${1:-false}" == "true" ]]
}

is_valid_directory() {
    [[ -d "$1" ]]
}

is_skip_env_check_enabled() {
    [[ "${argc_skip_env_check:-0}" == "1" ]]
}

is_short_mode_enabled() {
    [[ "${argc_short:-0}" == "1" ]]
}
