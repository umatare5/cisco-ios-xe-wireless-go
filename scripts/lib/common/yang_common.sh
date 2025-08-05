#!/usr/bin/env bash

# Cisco WNC YANG Common Library
# Shared functions and utilities for YANG-related operations

# Ensure common library is loaded first
if ! declare -F "get_default_controller" >/dev/null 2>&1; then
    CURRENT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
    source "${CURRENT_DIR}/common.sh"
fi

# YANG-specific constants
readonly YANG_DEFAULT_MODEL="Cisco-IOS-XE-wireless-access-point-oper"
readonly YANG_DEFAULT_REVISION="2023-08-01"
readonly YANG_DEFAULT_IDENTIFIER="access-point-oper-data"

# YANG-specific predicate functions
is_valid_hostname() {
    [[ -n "$1" && "$1" != " " && "$1" =~ ^[a-zA-Z0-9.-]+$ ]]
}

is_valid_token() {
    [[ -n "$1" && ${#1} -ge 4 ]]
}

is_yang_model_valid() {
    [[ -n "$1" && "$1" =~ ^Cisco-IOS-XE-wireless-.* ]]
}

is_yang_revision_valid() {
    local revision="$1"
    [[ "$revision" =~ ^[0-9]{4}-[0-9]{2}-[0-9]{2}$ ]]
}

is_yang_identifier_valid() {
    local identifier="$1"
    [[ -n "$identifier" && ${#identifier} -ge 3 ]]
}

# Common YANG argument validation
validate_yang_model_argument() {
    local model="$1"

    if [[ -z "$model" ]]; then
        show_error "YANG model cannot be empty"
        exit 2
    fi

    if ! is_yang_model_valid "$model"; then
        show_error "Invalid YANG model format: $model"
        show_error "Expected format: Cisco-IOS-XE-wireless-*"
        exit 2
    fi
}

validate_yang_revision_argument() {
    local revision="$1"

    if [[ -n "$revision" ]] && ! is_yang_revision_valid "$revision"; then
        show_error "Invalid YANG revision format: $revision"
        show_error "Expected format: YYYY-MM-DD"
        exit 2
    fi
}

validate_yang_identifier_argument() {
    local identifier="$1"

    if [[ -n "$identifier" ]] && ! is_yang_identifier_valid "$identifier"; then
        show_error "Invalid YANG identifier: $identifier"
        show_error "Identifier must be at least 3 characters"
        exit 2
    fi
}

# Common YANG help display function
show_yang_help() {
    local script_name="$1"
    local description="$2"
    local model_param="${3:-true}"
    local revision_param="${4:-false}"
    local identifier_param="${5:-false}"

    cat << EOF
Usage: $script_name [OPTIONS]

$description

OPTIONS:
    -c, --controller HOSTNAME   WNC controller hostname or IP address
                                (default: $(get_default_controller))
    -t, --token TOKEN          Basic auth token
                                (can also use WNC_ACCESS_TOKEN env var)
    -p, --protocol PROTOCOL    Protocol to use: http or https
                                (default: $(get_default_protocol))
    -k, --insecure             Skip TLS certificate verification
    -h, --help                 Show this help message
EOF

    if [[ "$model_param" == "true" ]]; then
        cat << EOF
    -m, --model MODEL          YANG model name
                                (default: $YANG_DEFAULT_MODEL)
EOF
    fi

    if [[ "$revision_param" == "true" ]]; then
        cat << EOF
    -r, --revision REVISION    YANG model revision
                                (default: $YANG_DEFAULT_REVISION)
EOF
    fi

    if [[ "$identifier_param" == "true" ]]; then
        cat << EOF
    -i, --identifier ID        YANG statement identifier
                                (default: $YANG_DEFAULT_IDENTIFIER)
EOF
    fi

    cat << EOF

ENVIRONMENT VARIABLES:
    WNC_ACCESS_TOKEN       Basic authentication token

EXAMPLES:
    # Generate authentication token first
    wnc generate token -u admin -p password123
    # Export the generated token
    export WNC_ACCESS_TOKEN="dXNlcjpwYXNzd29yZA=="

    # Using environment variable
    $script_name

    # Using command line options
    $script_name -c wnc1.example.internal -t "dXNlcjpwYXNzd29yZA=="

    # Using HTTP instead of HTTPS
    $script_name -p http -c 192.168.1.100

    # Skip certificate verification
    $script_name -k -c wnc.local
EOF

    if [[ "$model_param" == "true" ]]; then
        cat << EOF

    # Specify YANG model
    $script_name -m Cisco-IOS-XE-wireless-client-oper
EOF
    fi
}

# Common YANG argument parsing function
parse_yang_arguments() {
    local script_name="$1"
    local description="$2"
    local model_param="${3:-true}"
    local revision_param="${4:-false}"
    local identifier_param="${5:-false}"
    shift 5

    # Local variables (lowercase for scoped)
    local controller token protocol insecure_flag model revision identifier

    controller="$(get_default_controller)"
    token="${WNC_ACCESS_TOKEN:-}"
    protocol="$(get_default_protocol)"
    insecure_flag=""
    model="$YANG_DEFAULT_MODEL"
    revision="$YANG_DEFAULT_REVISION"
    identifier="$YANG_DEFAULT_IDENTIFIER"

    while [[ $# -gt 0 ]]; do
        case "$1" in
            -c|--controller)
                controller="${2:?--controller requires hostname}"
                shift 2
                ;;
            -t|--token)
                token="${2:?--token requires token value}"
                shift 2
                ;;
            -p|--protocol)
                protocol="${2:?--protocol requires http or https}"
                shift 2
                ;;
            -k|--insecure)
                insecure_flag="--insecure"
                shift
                ;;
            -m|--model)
                if [[ "$model_param" == "true" ]]; then
                    model="${2:?--model requires model name}"
                    shift 2
                else
                    show_error "Unknown option: $1"
                    show_yang_help "$script_name" "$description" \
                        "$model_param" "$revision_param" "$identifier_param"
                    exit 2
                fi
                ;;
            -r|--revision)
                if [[ "$revision_param" == "true" ]]; then
                    revision="${2:?--revision requires revision date}"
                    shift 2
                else
                    show_error "Unknown option: $1"
                    show_yang_help "$script_name" "$description" \
                        "$model_param" "$revision_param" "$identifier_param"
                    exit 2
                fi
                ;;
            -i|--identifier)
                if [[ "$identifier_param" == "true" ]]; then
                    identifier="${2:?--identifier requires identifier name}"
                    shift 2
                else
                    show_error "Unknown option: $1"
                    show_yang_help "$script_name" "$description" \
                        "$model_param" "$revision_param" "$identifier_param"
                    exit 2
                fi
                ;;
            -h|--help)
                show_yang_help "$script_name" "$description" \
                    "$model_param" "$revision_param" "$identifier_param"
                exit 0
                ;;
            *)
                show_error "Unknown option: $1"
                show_yang_help "$script_name" "$description" \
                    "$model_param" "$revision_param" "$identifier_param"
                exit 2
                ;;
        esac
    done

    # Validate arguments using available predicate functions
    if ! is_valid_hostname "$controller"; then
        show_error "Invalid controller hostname: $controller"
        return 1
    fi

    if [[ -n "$token" ]] && ! is_valid_token "$token"; then
        show_error "Invalid access token format"
        return 1
    fi

    if [[ "$model_param" == "true" ]]; then
        if ! is_yang_model_valid "$model"; then
            show_error "Invalid YANG model: $model"
            return 1
        fi
    fi

    if [[ "$revision_param" == "true" ]]; then
        if ! is_yang_revision_valid "$revision"; then
            show_error "Invalid YANG revision: $revision"
            return 1
        fi
    fi

    if [[ "$identifier_param" == "true" ]]; then
        if ! is_yang_identifier_valid "$identifier"; then
            show_error "Invalid YANG identifier: $identifier"
            return 1
        fi
    fi

    # Export parsed values as GLOBAL environment variables (UPPERCASE)
    export WNC_YANG_CONTROLLER="$controller"
    export WNC_YANG_TOKEN="$token"
    export WNC_YANG_PROTOCOL="$protocol"
    export WNC_YANG_INSECURE_FLAG="$insecure_flag"
    export WNC_YANG_MODEL="$model"
    export WNC_YANG_REVISION="$revision"
    export WNC_YANG_IDENTIFIER="$identifier"
}

# Global YANG utility functions (for use across scripts)
get_yang_controller() {
    echo "${WNC_YANG_CONTROLLER:-$(get_default_controller)}"
}

get_yang_token() {
    echo "${WNC_YANG_TOKEN:-${WNC_ACCESS_TOKEN:-}}"
}

get_yang_protocol() {
    echo "${WNC_YANG_PROTOCOL:-$(get_default_protocol)}"
}

get_yang_insecure_flag() {
    echo "${WNC_YANG_INSECURE_FLAG:-}"
}

get_yang_model() {
    echo "${WNC_YANG_MODEL:-$YANG_DEFAULT_MODEL}"
}

get_yang_revision() {
    echo "${WNC_YANG_REVISION:-$YANG_DEFAULT_REVISION}"
}

get_yang_identifier() {
    echo "${WNC_YANG_IDENTIFIER:-$YANG_DEFAULT_IDENTIFIER}"
}

# Common YANG setup functions for script standardization
setup_yang_authentication() {
    setup_authentication "$(get_yang_token)"
}

setup_yang_dependencies() {
    check_dependencies_for_format "pretty"
}

setup_yang_temp_file() {
    setup_temp_file_with_cleanup
}

# Complete YANG setup - combines all common initialization steps
setup_yang_environment() {
    local auth_token temp_file

    auth_token=$(setup_yang_authentication)
    setup_yang_dependencies
    temp_file=$(setup_yang_temp_file)

    # Export for caller
    echo "$auth_token"
    echo "$temp_file"
}

# Standard YANG script initialization
init_yang_script() {
    # Basic script setup first
    set -euo pipefail

    # Get script directory
    local script_dir
    script_dir="$(cd "$(dirname "${BASH_SOURCE[1]}")" && pwd)"
    readonly script_dir

    # Source common first if not already loaded
    if ! declare -F "SOURCE_WNC_LIBRARIES" >/dev/null 2>&1; then
        # shellcheck source=/dev/null
        source "${script_dir}/lib/common.sh"
    fi

    # Load all libraries
    SOURCE_WNC_LIBRARIES "$script_dir"

    # Return script directory
    echo "$script_dir"
}
