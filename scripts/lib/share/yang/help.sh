#!/usr/bin/env bash

# Cisco WNC YANG Operations - Help Functions
# Provides help and documentation functionality for YANG operations (usage + banners)

# Static help content (kept outside functions to keep functions short)
if [[ -z "${YANG_LIST_HELP_CONTENT:-}" ]]; then
    YANG_LIST_HELP_CONTENT="$(cat <<'EOF'
USAGE:
    list_yang_models.sh [OPTIONS]

DESCRIPTION:
    Lists all available YANG models from a Cisco Wireless Network Controller.
    Retrieves model information via RESTCONF API and displays in various formats.

OPTIONS:
    -c, --controller <HOST>   WNC controller hostname or IP (required unless WNC_CONTROLLER set)
    -t, --token <TOKEN>       Basic auth token (or use WNC_ACCESS_TOKEN env var)
    -p, --protocol <PROTO>    Protocol: http or https [default: https]
    -k, --insecure            Skip TLS certificate verification
    -v, --verbose             Enable verbose output
    -h, --help                Show this help message

EXAMPLES:
    # List models from default controller
    list_yang_models.sh

    # List from specific controller
    list_yang_models.sh --controller 192.168.1.100

    # Use with custom token
    list_yang_models.sh --token "base64_encoded_credentials"

    # Skip SSL verification (development)
    list_yang_models.sh --insecure

ENVIRONMENT:
    WNC_CONTROLLER            WNC controller hostname/IP
    WNC_ACCESS_TOKEN          Authentication token (base64 encoded)

AUTHENTICATION:
    Token should be base64 encoded in format: username:password
    Example: echo -n "admin:password" | base64
EOF
)"
    readonly YANG_LIST_HELP_CONTENT
fi

if [[ -z "${YANG_GET_HELP_CONTENT:-}" ]]; then
    YANG_GET_HELP_CONTENT="$(cat <<'EOF'
USAGE:
    get_yang_model_details.sh [OPTIONS] <MODEL>

DESCRIPTION:
    Retrieves detailed information about a specific YANG model from WNC.
    Returns complete model structure, statements, and metadata.

ARGUMENTS:
    <MODEL>                   YANG model name to retrieve

OPTIONS:
    -c, --controller <HOST>   WNC controller hostname (required unless WNC_CONTROLLER set)
    -t, --token <TOKEN>       Access token for authentication
    -f, --format <FORMAT>     Output format: json or xml [default: json]
    -v, --verbose             Enable verbose output
    -R, --raw                 Output raw response without formatting
    -k, --insecure            Skip SSL certificate verification
    -h, --help                Show this help message

EXAMPLES:
    # Get model details
    get_yang_model_details.sh "Cisco-IOS-XE-wireless-ap-cfg"

    # Get in XML format
    get_yang_model_details.sh --format xml "model-name"

    # Get with verbose output
    get_yang_model_details.sh --verbose "model-name"
EOF
)"
    readonly YANG_GET_HELP_CONTENT
fi

if [[ -z "${YANG_STATEMENT_HELP_CONTENT:-}" ]]; then
    YANG_STATEMENT_HELP_CONTENT="$(cat <<'EOF'
USAGE:
    get_yang_statement_details.sh [OPTIONS] <MODEL> <STATEMENT>

DESCRIPTION:
    Retrieves details for a specific YANG statement within a model.
    Provides statement structure, type information, and constraints.

ARGUMENTS:
    <MODEL>                   YANG model name
    <STATEMENT>               YANG statement/container name

OPTIONS:
    -c, --controller <HOST>   WNC controller hostname (required unless WNC_CONTROLLER set)
    -t, --token <TOKEN>       Access token for authentication
    -f, --format <FORMAT>     Output format: json or xml [default: json]
    -v, --verbose             Enable verbose output
    -k, --insecure            Skip SSL certificate verification
    -h, --help                Show this help message

EXAMPLES:
    # Get statement details
    get_yang_statement_details.sh "model-name" "container-name"

    # Get in XML format
    get_yang_statement_details.sh --format xml "model" "statement"
EOF
)"
    readonly YANG_STATEMENT_HELP_CONTENT
fi

# Display YANG operation banner
show_yang_banner() { wnc_banner_yang; }

# Display help for YANG model listing operations
show_yang_list_help() { show_yang_banner; printf "%s\n" "$YANG_LIST_HELP_CONTENT"; }

# Display help for YANG model details operations
show_yang_get_help() { show_yang_banner; printf "%s\n" "$YANG_GET_HELP_CONTENT"; }

# Display help for YANG statement details operations
show_yang_statement_help() { show_yang_banner; printf "%s\n" "$YANG_STATEMENT_HELP_CONTENT"; }

# Display verbose YANG operation configuration information
show_yang_verbose_info() {
    local controller="$1"
    local protocol="$2"
    local format="$3"

    is_verbose_enabled || return 0
    printf "%s\n" "YANG Operation Configuration:"
    printf "%s\n" "  Controller: ${controller:-${WNC_CONTROLLER:-<unset>}}"
    printf "%s\n" "  Protocol: ${protocol:-https}"
    printf "%s\n" "  Format: ${format:-json}"

    local insecure_status
    if is_insecure_enabled; then
        insecure_status="enabled"
    else
        insecure_status="disabled"
    fi
    printf "%s\n" "  Insecure: $insecure_status"

    printf "%s\n\n" "  Verbose: enabled"
}
