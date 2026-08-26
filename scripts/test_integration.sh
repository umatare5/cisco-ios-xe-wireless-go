#!/usr/bin/env bash

# Runs tests/integration against a live Cisco C9800. Needs WNC_CONTROLLER,
# WNC_ACCESS_TOKEN and WNC_AP_MAC_ADDR. Takes no arguments.

set -Eeuo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
readonly SCRIPT_DIR
readonly PROJECT_ROOT="${SCRIPT_DIR%/*}"

# shellcheck source=lib/log.sh
source "${SCRIPT_DIR}/lib/log.sh"
# shellcheck source=lib/env.sh
source "${SCRIPT_DIR}/lib/env.sh"
# shellcheck source=lib/gotest.sh
source "${SCRIPT_DIR}/lib/gotest.sh"

main() {
    validate_required_cli_tools "standard"
    gotest_run integration "$PROJECT_ROOT"
}

main "$@"
