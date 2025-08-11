#!/usr/bin/env bash

# @meta version 1.0.0
# @meta author "@umatare5"
# @describe List Cisco wireless YANG models from a Wireless Network Controller

# @option -c --controller <HOST>       WNC controller hostname or IP (required unless WNC_CONTROLLER set)
# @option -t --token <TOKEN>          Basic auth token (or use WNC_ACCESS_TOKEN env var)
# @option -p --protocol <PROTOCOL>    Protocol: http or https [default: https] [choices: http,https]
# @flag   -k --insecure               Skip TLS certificate verification
# @flag   -v --verbose                Enable verbose output
# @flag      --no-color               Disable colored output

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
readonly SCRIPT_DIR

# Source bootstrap library
source "${SCRIPT_DIR}/lib/bootstrap.sh"

# Initialize libraries and load YANG module
init_wnc_libraries "${SCRIPT_DIR}/lib/get_yang_models"

# Validate required CLI tools
validate_required_cli_tools "strict"

# Entrypoint: list wireless YANG models
main() {
    # Extract and provide arguments to library
    local controller="${argc_controller:-${WNC_CONTROLLER:-}}"
    local token="${argc_token:-${WNC_ACCESS_TOKEN:-}}"
    local protocol="${argc_protocol:-https}"
    local insecure="${argc_insecure:-false}"
    local format="${argc_format:-json}"
    local raw="${argc_raw:-false}"

    run_yang_list_operation "$controller" "$token" "$protocol" "$insecure" "$format" "$raw"
}

eval "$(argc --argc-eval "$0" "$@")"
