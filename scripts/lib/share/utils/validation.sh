#!/usr/bin/env bash

# Cisco WNC Validation Library
# Provides functions to validate various parameters and formats used in WNC scripts

# Display error message and exit with invalid arguments exit code
exit_with_invalid_args() {
    local message="$1"
    error "$message"
    exit "$(get_exit_invalid_args)"
}

# Validate HTTP protocol format allowing only http or https
validate_protocol() {
    local protocol="$1"
    local http_protocol="http"
    local https_protocol="https"

    # Early return for valid protocols
    if [[ "$protocol" == "$http_protocol" || \
"$protocol" == "$https_protocol" ]]; then
        return 0
    fi

    # Invalid protocol - show error
    exit_with_invalid_args "Invalid protocol: $protocol. Protocol must be 'http' or 'https'"
}

# Validate output format allowing only pretty, json, or raw
validate_output_format() {
    local format="$1"
    local format_pretty="pretty"
    local format_json="json"
    local format_raw="raw"

    # Early return for valid formats
    case "$format" in
        "$format_pretty"|"$format_json"|"$format_raw")
            return 0
            ;;
    esac

    # Invalid format - show error
    # Correct message with full list of allowed values
    exit_with_invalid_args "Invalid output format: $format. Format must be 'pretty', 'json', or 'raw'"
}

# Validate YANG model revision format following YYYY-MM-DD pattern
validate_revision_format() {
    local revision="$1"

    # Early return for valid revision
    if is_valid_revision "$revision"; then
        return 0
    fi

    # Invalid revision - show error
    exit_with_invalid_args "Invalid revision format: $revision. Revision must be in YYYY-MM-DD format"
}

# Validate hostname format ensuring it contains only alphanumeric characters, dots, and hyphens
validate_hostname() {
    local hostname="$1"

    # Early return for empty hostname
    if is_empty "$hostname"; then
        exit_with_invalid_args "Hostname cannot be empty"
    fi

    # Early return for valid hostname
    if is_valid_hostname "$hostname"; then
        return 0
    fi

    # Invalid hostname - show error
    local error_msg="Invalid hostname format: $hostname. "
    error_msg+="Hostname must contain only letters, numbers, dots, and hyphens"
    exit_with_invalid_args "$error_msg"
}

# Validate Basic Auth Token format checking for valid base64 encoding
validate_token() {
    local token="$1"

    # Early return for empty token (valid case)
    if is_empty "$token"; then
        return 0
    fi

    # Early return for valid token
    if is_valid_token_format "$token"; then
        return 0
    fi

    # Invalid token - show error
    exit_with_invalid_args "Invalid Basic Auth Token format. Token should be base64 encoded"
}

# Validate YANG model name format ensuring it follows Cisco-IOS-XE-wireless pattern
validate_yang_model() {
    local model="$1"
    local prefix="Cisco-IOS-XE-wireless-"
    local oper_suffix="-oper"
    local cfg_suffix="-cfg"

    # Early return for empty model
    if is_empty "$model"; then
        exit_with_invalid_args "YANG model name cannot be empty"
    fi

    # Early return for valid model
    if is_valid_yang_model "$model"; then
        return 0
    fi

    # Invalid model - show error
    local examples="${prefix}access-point${oper_suffix}, ${prefix}wlan${cfg_suffix}"
    local error_msg="Invalid YANG model format: $model. "
    error_msg+="Model must start with '$prefix' and end with '$oper_suffix' or '$cfg_suffix'. "
    error_msg+="Examples: $examples"
    exit_with_invalid_args "$error_msg"
}

# Validate identifier format ensuring it contains only letters, numbers, and hyphens
validate_identifier() {
    local identifier="$1"

    # Early return for empty identifier
    if is_empty "$identifier"; then
        exit_with_invalid_args "Identifier cannot be empty"
    fi

    # Early return for valid identifier
    if is_valid_identifier "$identifier"; then
        return 0
    fi

    # Invalid identifier - show error
    exit_with_invalid_args \
      "Invalid identifier format: $identifier. Identifier should contain only letters, numbers, and hyphens"
}

# Validate insecure flag (always returns success as this is a boolean flag)
validate_insecure_flag() {
    # Always valid - this is a boolean flag with no parameter
    return 0
}

# Validate verbose flag (always returns success as this is a boolean flag)
validate_verbose_flag() {
    # Always valid - this is a boolean flag with no parameter
    return 0
}

# Predicate functions for better readability

# Check if a string is empty or contains only whitespace
is_empty() {
    local value="$1"
    [[ -z "$value" ]]
}

# Check if a string is not empty and contains non-whitespace characters
is_not_empty() {
    local value="$1"
    [[ -n "$value" ]]
}

# Check if hostname has valid format
is_valid_hostname() {
    local hostname="$1"
    local hostname_pattern="^[a-zA-Z0-9.-]+$"
    [[ "$hostname" =~ $hostname_pattern ]]
}

# Check if token has valid base64 format
is_valid_token_format() {
    local token="$1"
    local token_pattern="^[A-Za-z0-9+/]+=*$"
    local min_token_length=8

    # Early return for empty token (valid case)
    if is_empty "$token"; then
        return 0
    fi

    # Early return for invalid format or length
    if [[ ! "$token" =~ $token_pattern ]]; then
        return 1
    fi

    # Check minimum length
    [[ ${#token} -ge $min_token_length ]]
}

# Check if a YANG model has valid format
is_valid_yang_model() {
    local model="$1"
    local yang_model_pattern="^Cisco-IOS-XE-wireless-.+-(oper|cfg)$"
    [[ "$model" =~ $yang_model_pattern ]]
}

# Check if a revision has valid format
is_valid_revision() {
    local revision="$1"
    local date_pattern="^[0-9]{4}-[0-9]{2}-[0-9]{2}$"
    [[ "$revision" =~ $date_pattern ]]
}

# Check if an identifier has valid format
is_valid_identifier() {
    local identifier="$1"
    local identifier_pattern="^[a-zA-Z0-9-]+$"
    [[ "$identifier" =~ $identifier_pattern ]]
}

# Check if jq is required for the output format
is_jq_required() {
    local format="$1"
    local format_json="json"
    local format_pretty="pretty"
    [[ "$format" == "$format_json" || "$format" == "$format_pretty" ]]
}

# Check if curl is available
is_curl_available() {
    command -v curl >/dev/null 2>&1
}

# Check if jq is available
is_jq_available() {
    command -v jq >/dev/null 2>&1
}
