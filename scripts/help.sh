#!/usr/bin/env bash

# @meta version 1.0.0
# @meta author "@umatare5"
# @describe Show comprehensive help for Cisco WNC project development

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
readonly SCRIPT_DIR

# Source bootstrap library
source "${SCRIPT_DIR}/lib/bootstrap.sh"

# Load help implementation
init_wnc_libraries "${SCRIPT_DIR}/lib/help"

# Entrypoint: show help
main() {
    run_help_operation
}

main "$@" || exit $?
