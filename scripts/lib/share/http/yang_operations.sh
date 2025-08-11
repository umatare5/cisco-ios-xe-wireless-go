#!/usr/bin/env bash

# Cisco WNC YANG Operations Library
# Provides HTTP/RESTCONF operations for YANG models

# YANG defaults (idempotent)
: "${YANG_DEFAULT_MODEL:=Cisco-IOS-XE-wireless-access-point-oper}"
readonly YANG_DEFAULT_MODEL
: "${YANG_DEFAULT_REVISION:=2023-08-01}"
readonly YANG_DEFAULT_REVISION
: "${YANG_DEFAULT_IDENTIFIER:=access-point-oper-data}"
readonly YANG_DEFAULT_IDENTIFIER

# Validate access token format and length
is_valid_token() {
    [[ -n "$1" && ${#1} -ge 4 ]]
}

# Validate YANG model naming convention
is_yang_model_valid() {
    [[ -n "$1" && "$1" =~ ^Cisco-IOS-XE-wireless-.* ]]
}

# Validate YANG revision date format (YYYY-MM-DD)
is_yang_revision_valid() {
    local revision="$1"
    [[ "$revision" =~ ^[0-9]{4}-[0-9]{2}-[0-9]{2}$ ]]
}

# Validate YANG identifier format and minimum length
is_yang_identifier_valid() {
    local identifier="$1"
    [[ -n "$identifier" && ${#identifier} -ge 3 ]]
}

# Validate YANG model argument format and exit on error
validate_yang_model_argument() {
    local model="$1"

    if [[ -z "$model" ]]; then
        error "YANG model cannot be empty"
        exit 2
    fi

    if ! is_yang_model_valid "$model"; then
        error "Invalid YANG model format: $model"
        error "Expected format: Cisco-IOS-XE-wireless-*"
        exit 2
    fi
}

# Validate YANG revision argument format and exit on error
validate_yang_revision_argument() {
    local revision="$1"

    # Early return if revision is empty
    if [[ -z "$revision" ]]; then
        return 0
    fi

    # Check revision validity and exit on error
    if ! is_yang_revision_valid "$revision"; then
        error "Invalid YANG revision format: $revision"
        error "Expected format: YYYY-MM-DD"
        exit 2
    fi
}

# Validate YANG identifier argument format and exit on error
validate_yang_identifier_argument() {
    local identifier="$1"

    # Early return if identifier is empty
    if [[ -z "$identifier" ]]; then
        return 0
    fi

    # Check identifier validity and exit on error
    if ! is_yang_identifier_valid "$identifier"; then
        error "Invalid YANG identifier: $identifier"
        error "Identifier must be at least 3 characters"
        exit 2
    fi
}

# Initialize YANG environment from argc parsed arguments
initialize_yang_from_argc() {
    export WNC_YANG_CONTROLLER="${argc_controller:-$(get_default_controller)}"
    export WNC_YANG_TOKEN="${argc_token:-${WNC_ACCESS_TOKEN:-}}"
    export WNC_YANG_PROTOCOL="${argc_protocol:-$(get_default_protocol)}"
    export WNC_YANG_INSECURE_FLAG="${argc_insecure:+--insecure}"
    export WNC_YANG_MODEL="${argc_model:-$YANG_DEFAULT_MODEL}"
    export WNC_YANG_REVISION="${argc_revision:-$YANG_DEFAULT_REVISION}"
    export WNC_YANG_IDENTIFIER="${argc_identifier:-$YANG_DEFAULT_IDENTIFIER}"
}

# Check if YANG token is provided
_has_yang_token() {
    [[ -n "$WNC_YANG_TOKEN" ]]
}

# Check if YANG model is provided
_has_yang_model() {
    [[ -n "$WNC_YANG_MODEL" ]]
}

# Check if YANG revision is provided
_has_yang_revision() {
    [[ -n "$WNC_YANG_REVISION" ]]
}

# Check if YANG identifier is provided
_has_yang_identifier() {
    [[ -n "$WNC_YANG_IDENTIFIER" ]]
}

# Check if YANG model is provided and valid
_is_yang_model_valid_if_present() {
    ! _has_yang_model || is_yang_model_valid "$WNC_YANG_MODEL"
}

# Check if YANG revision is provided and valid
_is_yang_revision_valid_if_present() {
    ! _has_yang_revision || is_yang_revision_valid "$WNC_YANG_REVISION"
}

# Check if YANG identifier is provided and valid
_is_yang_identifier_valid_if_present() {
    ! _has_yang_identifier || is_yang_identifier_valid "$WNC_YANG_IDENTIFIER"
}

# Validate hostname and token parameters
_validate_yang_connection_params() {
    # Validate controller hostname (required)
    if ! is_valid_hostname "$WNC_YANG_CONTROLLER"; then
        error "Invalid controller hostname: $WNC_YANG_CONTROLLER"
        exit 1
    fi

    # Early return if no token provided (optional parameter)
    if ! _has_yang_token; then
        return 0
    fi

    # Validate token format and exit on error
    if ! is_valid_token "$WNC_YANG_TOKEN"; then
        error "Invalid access token format"
        exit 1
    fi
}

# Validate YANG model parameters
_validate_yang_model_params() {
    # Validate YANG model if provided
    if ! _is_yang_model_valid_if_present; then
        error "Invalid YANG model: $WNC_YANG_MODEL"
        exit 2
    fi

    # Validate YANG revision if provided
    if ! _is_yang_revision_valid_if_present; then
        error "Invalid YANG revision: $WNC_YANG_REVISION"
        exit 2
    fi

    # Validate YANG identifier if provided
    if ! _is_yang_identifier_valid_if_present; then
        error "Invalid YANG identifier: $WNC_YANG_IDENTIFIER"
        exit 2
    fi
}

# Validate YANG environment variables after argc parsing
validate_yang_environment() {
    _validate_yang_connection_params
    _validate_yang_model_params
}

# Get currently configured YANG controller hostname
get_yang_controller() {
    printf '%s\n' "${WNC_YANG_CONTROLLER:-$(get_default_controller)}"
}

# Get currently configured YANG access token
get_yang_token() {
    printf '%s\n' "${WNC_YANG_TOKEN:-${WNC_ACCESS_TOKEN:-}}"
}

# Get currently configured YANG protocol (http/https)
get_yang_protocol() {
    printf '%s\n' "${WNC_YANG_PROTOCOL:-$(get_default_protocol)}"
}

# Get currently configured insecure flag for YANG operations
get_yang_insecure_flag() {
    printf '%s\n' "${WNC_YANG_INSECURE_FLAG:-}"
}

# Get currently configured YANG model name
get_yang_model() {
    printf '%s\n' "${WNC_YANG_MODEL:-$YANG_DEFAULT_MODEL}"
}

# Get currently configured YANG model revision
get_yang_revision() {
    printf '%s\n' "${WNC_YANG_REVISION:-$YANG_DEFAULT_REVISION}"
}

# Get currently configured YANG statement identifier
get_yang_identifier() {
    printf '%s\n' "${WNC_YANG_IDENTIFIER:-$YANG_DEFAULT_IDENTIFIER}"
}

# Setup authentication for YANG operations using configured token
setup_yang_authentication() {
    setup_authentication "$(get_yang_token)"
}

# Setup required dependencies for YANG operations
setup_yang_dependencies() {
    check_dependencies_for_format "pretty"
}

# Complete YANG environment setup for argc-based scripts
setup_yang_environment_from_argc() {
    initialize_yang_from_argc
    validate_yang_environment

    local auth_token temp_file
    auth_token=$(setup_yang_authentication)
    setup_yang_dependencies
    temp_file=$(setup_temp_file)

    # Return setup values for caller
    printf '%s\n' "$auth_token"
    printf '%s\n' "$temp_file"
}
