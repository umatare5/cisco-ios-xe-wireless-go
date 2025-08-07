#!/usr/bin/env bash
set -euo pipefail

# Cisco WNC Lint Code - Help Functions
# Provides help and documentation functionality for lint_code operations

# Source bootstrap library
LIB_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
# shellcheck disable=SC1090  # Dynamic source loading
source "${LIB_ROOT}/bootstrap.sh"
init_wnc_basic

show_lint_banner() {
    if ! is_no_color_enabled; then
        echo -e "\033[34m----------------------------------------\033[0m"
        echo -e "\033[34m       Cisco WNC Code Linter           \033[0m"
        echo -e "\033[34m       golangci-lint Integration       \033[0m"
        echo -e "\033[34m----------------------------------------\033[0m"
    else
        echo "----------------------------------------"
        echo "       Cisco WNC Code Linter"
        echo "       golangci-lint Integration"
        echo "----------------------------------------"
    fi
    echo
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
