#!/usr/bin/env bash
set -euo pipefail
# Dependencies Output Functions

# Color definitions (will be empty if --no-color is used)
readonly DEP_COLOR_RED="${COLOR_RED:-}"
readonly DEP_COLOR_GREEN="${COLOR_GREEN:-}"
readonly DEP_COLOR_YELLOW="${COLOR_YELLOW:-}"
readonly DEP_COLOR_BLUE="${COLOR_BLUE:-}"
readonly DEP_COLOR_RESET="${COLOR_RESET:-}"

format_dependencies_error() {
    echo "${DEP_COLOR_RED}✗ Dependencies Error: $1${DEP_COLOR_RESET}" >&2
}

format_dependencies_success() {
    echo "${DEP_COLOR_GREEN}✓ Dependencies Success: $1${DEP_COLOR_RESET}"
}

format_dependencies_warning() {
    echo "${DEP_COLOR_YELLOW}⚠ Dependencies Warning: $1${DEP_COLOR_RESET}" >&2
}

format_dependencies_info() {
    echo "${DEP_COLOR_BLUE}ℹ Dependencies Info: $1${DEP_COLOR_RESET}"
}

format_dependencies_banner() {
    if is_no_color_enabled; then
        echo "----------------------------------------------------"
        echo " Cisco WNC Dependencies Management"
        echo "----------------------------------------------------"
    else
        echo "${DEP_COLOR_BLUE}----------------------------------------------------${DEP_COLOR_RESET}"
        echo "${DEP_COLOR_BLUE} Cisco WNC Dependencies Management${DEP_COLOR_RESET}"
        echo "${DEP_COLOR_BLUE}----------------------------------------------------${DEP_COLOR_RESET}"
    fi
}

format_step_message() {
    local step="$1"
    local message="$2"
    echo "${DEP_COLOR_BLUE}[$step]${DEP_COLOR_RESET} $message"
}

format_dependency_status() {
    local package="$1"
    local version="$2"
    local status="$3"

    case "$status" in
        "installed")
            echo "  ${DEP_COLOR_GREEN}✓${DEP_COLOR_RESET} $package ($version)"
            ;;
        "updated")
            echo "  ${DEP_COLOR_YELLOW}↑${DEP_COLOR_RESET} $package ($version)"
            ;;
        "downloading")
            echo "  ${DEP_COLOR_BLUE}↓${DEP_COLOR_RESET} $package ($version)"
            ;;
        *)
            echo "  ${DEP_COLOR_RED}✗${DEP_COLOR_RESET} $package ($version)"
            ;;
    esac
}
