#!/usr/bin/env bash

# Cisco WNC Lint Code - Help Functions
# Provides help and documentation functionality for lint_code operations

set -euo pipefail

# Source bootstrap library
LIB_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
source "${LIB_ROOT}/bootstrap.sh"
init_wnc_basic

show_lint_banner() {
    # Use unified banner helper for consistency
    wnc_banner_lint
}

show_lint_help() {
    show_lint_banner
    cat << 'EOF'
USAGE:
    lint.sh [OPTIONS]

DESCRIPTION:
    Runs golangci-lint to analyze Go code for potential issues, style violations,
    and best practice deviations. Supports automatic fixing and custom configurations.

OPTIONS:
    -p, --project <DIR>      Project root directory [default: .]
    -v, --verbose            Enable verbose output
        --fix                Automatically fix issues where possible
        --config <FILE>      Custom golangci-lint config file path
    -h, --help               Show this help message

EXAMPLES:
    # Basic linting
    lint.sh

    # Lint with auto-fix
    lint.sh --fix

    # Lint specific project directory
    lint.sh --project /path/to/project

    # Use custom config
    lint.sh --config .golangci.custom.yml

REQUIREMENTS:
    - golangci-lint must be installed
    - Valid Go module in target directory

INSTALLATION:
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

EOF
}

show_lint_verbose_info() {
    is_verbose_enabled || return 0

    echo "Lint Configuration:"
    echo "  Project: ${argc_project:-.}"
    echo "  Auto-fix: $(is_auto_fix_enabled && echo "enabled" || echo "disabled")"
    echo "  Config file: ${argc_config:-default}"
    echo "  Verbose: enabled"
    echo
}
