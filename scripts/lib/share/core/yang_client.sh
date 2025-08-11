#!/usr/bin/env bash

# Cisco WNC Core - YANG Operations Functions
# Provides YANG API operations and RESTCONF client functions

# Guard against multiple inclusions
if [[ -n "${WNC_YANG_CLIENT_LOADED:-}" ]]; then
    return 0
fi
readonly WNC_YANG_CLIENT_LOADED=1

# ===============================
# URL Construction Functions
# ===============================

# Build YANG API URL with parameters
build_yang_url() {
    local controller="$1"
    local protocol="$2"
    local endpoint="$3"

    if ! _has_controller "$controller"; then
        error "Controller not specified"
        return 1
    fi

    printf '%s\n' "${protocol}://${controller}/restconf/data/${endpoint}"
}

# Build YANG model URL for specific model
build_yang_model_url() {
    local controller="$1"
    local protocol="$2"

    if ! _has_controller "$controller"; then
        error "Controller not specified"
        return 1
    fi

    printf '%s\n' "${protocol}://${controller}/restconf/yang-library-version"
}

# ===============================
# HTTP Client Functions
# ===============================

# Execute curl with provided options
execute_curl() {
    local url="$1"
    local insecure="$2"
    local output_format="${3:-pretty}"
    local output_file="${4:-}"

    local curl_opts=()

    if [[ "$insecure" == "true" ]]; then
        curl_opts+=("-k")
    fi

    if [[ "$output_format" == "file" && -n "$output_file" ]]; then
        curl -s -H "Authorization: Basic $VALIDATED_TOKEN" \
             "${curl_opts[@]}" "$url" -o "$output_file"
        return
    fi

    curl -s -H "Authorization: Basic $VALIDATED_TOKEN" \
         "${curl_opts[@]}" "$url"
}

# Execute authenticated YANG request
execute_yang_request() {
    local controller="$1"
    local endpoint="$2"
    local insecure="${3:-false}"
    local output_file="${4:-}"

    local protocol="https"
    if [[ "$insecure" == "true" ]]; then
        protocol="https"
    fi

    local url
    url=$(build_yang_url "$controller" "$protocol" "$endpoint")

    if [[ -n "$output_file" ]]; then
        execute_curl "$url" "$insecure" "file" "$output_file"
    else
        execute_curl "$url" "$insecure" "pretty"
    fi
}
