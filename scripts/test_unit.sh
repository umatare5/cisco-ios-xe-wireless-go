#!/usr/bin/env bash

# Runs the unit tests. Takes no arguments: the repository root is where this script lives,
# and lib/gotest.sh owns the timeout. Coverage and the HTML report are test_coverage.sh.

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
    gotest_run unit "$PROJECT_ROOT"
}

main "$@"
