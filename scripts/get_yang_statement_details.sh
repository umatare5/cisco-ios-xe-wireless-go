#!/usr/bin/env bash

# Cisco WNC YANG Statement Details Script
# This script fetches detailed information about a specific YANG statement
# from the WNC controller

set -euo pipefail

# Load library functions
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# shellcheck source=./lib/common.sh
source "${SCRIPT_DIR}/lib/common.sh"
source_wnc_libraries "$SCRIPT_DIR"

# Default values
readonly DEFAULT_YANG_MODEL="Cisco-IOS-XE-wireless-access-point-oper"
readonly DEFAULT_IDENTIFIER="access-point-oper-data"

# Global variables - set default values using constants functions
controller="$(get_default_controller)"
token=""
protocol="$(get_default_protocol)"
yang_model="$DEFAULT_YANG_MODEL"
identifier="$DEFAULT_IDENTIFIER"
insecure_flag=""
verbose=false
output_format="pretty"
temp_file=""

# Display usage information
usage() {
    local default_protocol
    default_protocol="$(get_default_protocol)"

    cat << EOF
Usage: $0 [OPTIONS]

Fetch detailed information about a specific YANG statement from a Wireless \
Network Controller

OPTIONS:
    -c, --controller HOSTNAME   WNC controller hostname or IP address (required)
    -t, --token TOKEN          Basic auth token (can also use \
WNC_ACCESS_TOKEN env var)
    -p, --protocol PROTOCOL    Protocol to use: http or https \
(default: $default_protocol)
    -m, --model MODEL          YANG model name \
(default: Cisco-IOS-XE-wireless-access-point-oper)
    -i, --id IDENTIFIER        YANG model identifier \
(default: access-point-oper-data)
    -k, --insecure             Skip TLS certificate verification
    -v, --verbose              Show detailed output including raw response
    -f, --format FORMAT        Output format: pretty, json, raw \
(default: pretty)
    -h, --help                 Show this help message

ENVIRONMENT VARIABLES:
    WNC_ACCESS_TOKEN       Basic authentication token \
(generate with: wnc generate token)

NOTES:
    Authentication Token Generation:
    - Use the wnc CLI tool to generate a Basic authentication token:
      wnc generate token -u <username> -p <password>
    - The generated token can be used with the -t option or \
WNC_ACCESS_TOKEN environment variable

    YANG Model and Identifier:
    - The model should be in format like \
'Cisco-IOS-XE-wireless-access-point-oper'
    - The identifier corresponds to the data node within that model
    - Use the list_yang_models.sh script to find available models

EXAMPLES:
    # Generate authentication token first
    wnc generate token -u admin -p password123
    # Export the generated token
    export WNC_ACCESS_TOKEN="dXNlcjpwYXNzd29yZA=="

    # Basic usage with environment variable
    $0 -c wnc1.example.internal

    # Get specific YANG statement details
    $0 -c wnc1.example.internal \
-m "Cisco-IOS-XE-wireless-access-point-oper" -i "access-point-oper-data"

    # Output in JSON format
    $0 -c wnc1.example.internal -f json

    # Raw output for further processing
    $0 -c wnc1.example.internal -f raw
EOF
}

# Parse command line arguments
parse_arguments() {
    while [[ $# -gt 0 ]]; do
        case $1 in
            -c|--controller) validate_hostname "$2"; controller="$2"; shift 2 ;;
            -t|--token)      validate_token "$2"; token="$2"; shift 2 ;;
            -p|--protocol)   validate_protocol "$2"; protocol="$2"; shift 2 ;;
            -m|--model)      validate_yang_model "$2"; yang_model="$2"; shift 2 ;;
            -i|--id)         validate_identifier "$2"; identifier="$2"; shift 2 ;;
            -k|--insecure)   validate_insecure_flag; insecure_flag="--insecure"; shift ;;
            -v|--verbose)    validate_verbose_flag; verbose=true; shift ;;
            -f|--format)     validate_output_format "$2"; output_format="$2"; shift 2 ;;
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

    check_dependencies_for_format "$output_format"

    temp_file=$(setup_temp_file_with_cleanup)

    local url
    url=$(build_yang_statement_url "$protocol" "$controller" "$yang_model" "$identifier")

    display_yang_statement_configuration \
      "$protocol" "$controller" "$yang_model" "$identifier" \
      "$output_format" "$insecure_flag" "$verbose" "$url"

    execute_curl_request "$url" "$auth_token" "$insecure_flag" "$temp_file"

    show_raw_response "$temp_file" "$verbose" "$output_format"

    # Output based on format
    case "$output_format" in
        "raw")
            cat "$temp_file"
            ;;
        "json")
            format_json_output "$temp_file"
            ;;
        "pretty")
            format_yang_statement_details_pretty "$temp_file"
            ;;
    esac

    show_completion "$output_format"
}

# Execute main function with all arguments
main "$@"
