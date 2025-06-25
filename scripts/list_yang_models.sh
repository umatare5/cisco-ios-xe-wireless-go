#!/usr/bin/env bash

# Cisco WNC YANG Model Listing Script
# This script lists Cisco wireless YANG models available on the WNC controller

set -euo pipefail

# Load library functions
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# shellcheck source=./lib/common.sh
source "${SCRIPT_DIR}/lib/common.sh"
source_wnc_libraries "$SCRIPT_DIR"

# Global variables - set default values using constants functions
controller="$(get_default_controller)"
token=""
protocol="$(get_default_protocol)"
insecure_flag=""
temp_file=""

# Display usage information
usage() {
    local default_controller
    local default_protocol
    default_controller="$(get_default_controller)"
    default_protocol="$(get_default_protocol)"

    cat << EOF
Usage: $0 [OPTIONS]

List Cisco wireless YANG models from a Wireless Network Controller

OPTIONS:
    -c, --controller HOSTNAME   WNC controller hostname or IP address \
(default: $default_controller)
    -t, --token TOKEN          Basic auth token (can also use \
WNC_ACCESS_TOKEN env var)
    -p, --protocol PROTOCOL    Protocol to use: http or https \
(default: $default_protocol)
    -k, --insecure             Skip TLS certificate verification
    -h, --help                 Show this help message

ENVIRONMENT VARIABLES:
    WNC_ACCESS_TOKEN       Basic authentication token \
(generate with: wnc generate token)

EXAMPLES:
    # Generate authentication token first
    wnc generate token -u admin -p password123
    # Export the generated token
    export WNC_ACCESS_TOKEN="dXNlcjpwYXNzd29yZA=="

    # Using environment variable
    $0

    # Using command line options
    $0 -c wnc1.example.internal -t "dXNlcjpwYXNzd29yZA=="

    # Using HTTP instead of HTTPS
    $0 -p http -c 192.168.1.100

    # Skip certificate verification
    $0 -k -c wnc.local
EOF
}

# Parse command line arguments
parse_arguments() {
    while [[ $# -gt 0 ]]; do
        case $1 in
            -c|--controller) validate_hostname "$2"; controller="$2"; shift 2 ;;
            -t|--token)      validate_token "$2"; token="$2"; shift 2 ;;
            -p|--protocol)   validate_protocol "$2"; protocol="$2"; shift 2 ;;
            -k|--insecure)   validate_insecure_flag; insecure_flag="--insecure"; shift ;;
            -h|--help)       usage; exit "$(get_exit_success)" ;;
            *)
                echo "Error: Unknown option $1" >&2
                usage
                exit "$(get_exit_invalid_args)"
                ;;
        esac
    done
}

# Main function that orchestrates the entire process
main() {
    parse_arguments "$@"

    # Early authentication check and setup
    local auth_token
    auth_token=$(setup_authentication "$token")

    check_dependencies_for_format "pretty"

    temp_file=$(setup_temp_file_with_cleanup)

    local url
    url=$(build_yang_models_url "$protocol" "$controller")

    display_configuration "$protocol" "$controller" "pretty"

    execute_curl_request "$url" "$auth_token" "$insecure_flag" "$temp_file"

    format_yang_models_pretty "$temp_file"

    show_completion "pretty"
}

# Execute main function with all arguments
main "$@"
