#!/usr/bin/env bash
# @meta version 1.0.0
# @meta author "@umatare5"
# @describe Cisco WNC YANG Statement Details - Fetch detailed statement information

# @arg model!                          YANG model name to fetch statement details for
# @option -s --statement <STMT>        Specific statement to fetch (optional)
# @option -c --controller <HOST>       WNC controller hostname/IP
# @option -t --token <TOKEN>           Base64 encoded credentials
# @option -o --output <FILE>           Output file path (default: stdout)
# @flag   -v --verbose                 Enable verbose output
# @flag   -r --raw                     Output raw JSON without formatting
# @flag      --no-color                Disable colored output

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
MODULE_DIR="${SCRIPT_DIR}/lib/yang_operations"

# Source shared libraries
source "${SCRIPT_DIR}/lib/common/common.sh"
SOURCE_WNC_LIBRARIES "$SCRIPT_DIR"

# Source module-specific libraries
source "${MODULE_DIR}/help.sh"
source "${MODULE_DIR}/output.sh"
source "${MODULE_DIR}/core.sh"

# Predicate functions for improved readability
is_verbose_enabled() { [[ "${argc_verbose:-0}" == "1" ]]; }
is_no_color_enabled() { [[ "${argc_no_color:-0}" == "1" ]]; }
is_raw_enabled() { [[ "${argc_raw:-false}" == "true" ]]; }
is_command_available() { command -v "${1:-}" >/dev/null 2>&1; }

main() {
    run_yang_get_statement_operation
}

eval "$(argc --argc-eval "$0" "$@")"
