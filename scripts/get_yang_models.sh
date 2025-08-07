#!/usr/bin/env bash
# @meta version 1.0.0
# @meta author "@umatare5"
# @describe List Cisco wireless YANG models from a Wireless Network Controller

# @option -c --controller <HOST>       WNC controller hostname or IP [default: wnc1.example.internal]
# @option -t --token <TOKEN>          Basic auth token (or use WNC_ACCESS_TOKEN env var)
# @option -p --protocol <PROTOCOL>    Protocol: http or https [default: https] [choices: http,https]
# @flag   -k --insecure               Skip TLS certificate verification
# @flag   -v --verbose                Enable verbose output
# @flag      --no-color               Disable colored output

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Source bootstrap library
source "${SCRIPT_DIR}/lib/bootstrap.sh"

# Initialize WNC libraries with YANG operations module
init_wnc_libraries "$SCRIPT_DIR" "${SCRIPT_DIR}/lib/yang_operations"

# Validate required CLI tools before proceeding (including curl for API calls)
validate_required_cli_tools "strict"

# Predicate functions for improved readability using argc validation helpers
is_verbose_enabled() { is_enabled "${argc_verbose:-0}"; }
is_no_color_enabled() { is_enabled "${argc_no_color:-0}"; }
is_insecure_enabled() { is_enabled "${argc_insecure:-0}"; }
is_command_available() { command -v "${1:-}" >/dev/null 2>&1; }

main() {
    run_yang_list_operation
}

eval "$(argc --argc-eval "$0" "$@")"
