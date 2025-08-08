#!/usr/bin/env bash
# @meta version 1.0.0
# @meta author "@umatare5"
# @describe Cisco WNC Dependencies Management - Install, update, and clean Go dependencies

# @option -p --project <DIR>           Project root directory [default: .]
# @option    --golangci-lint <VERSION> golangci-lint version [default: latest]
# @option    --gotestsum <VERSION>     gotestsum version [default: latest]
# @flag   -v --verbose                Enable verbose output
# @flag   -c --clean                  Clean module cache before installing
# @flag   -u --update                 Update all dependencies to latest versions
# @flag      --force                  Force reinstall even if exists
# @flag      --download-only          Download dependencies without installing
# @flag      --verify                 Verify dependencies after installation
# @flag      --no-color               Disable colored output

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
readonly SCRIPT_DIR

# Source bootstrap library
source "${SCRIPT_DIR}/lib/bootstrap.sh"

# Initialize WNC libraries with dependencies module
init_wnc_libraries "${SCRIPT_DIR}" "${SCRIPT_DIR}/lib/dependencies"

# Validate required CLI tools before proceeding
validate_required_cli_tools "standard"

main() {
    run_dependencies_operation
}

eval "$(argc --argc-eval "$0" "$@")"
