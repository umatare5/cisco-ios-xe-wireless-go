#!/usr/bin/env bash

# Cisco WNC Color System
# Provides unified color codes and logging functions across all scripts

# Color definitions (ANSI escape sequences)
# Only define if not already set to avoid readonly conflicts
: "${COLOR_RESET:=\033[0m}"
: "${COLOR_WHITE:=\033[37m}"
: "${COLOR_BLUE:=\033[34m}"     # Info/Primary
: "${COLOR_CYAN:=\033[36m}"     # Info (alternative)
: "${COLOR_YELLOW:=\033[33m}"   # Warning
: "${COLOR_RED:=\033[31m}"      # Error
: "${COLOR_GREEN:=\033[32m}"    # Success

# Check if colors should be disabled
is_no_color_enabled() {
    [[ "${NO_COLOR:-}" == "true" || "${CI:-}" == "true" || ! -t 1 ]]
}

# Core logging function with color support
_log_with_color() {
    local color="$1"
    local symbol="$2"
    local prefix="$3"
    local stream="$4"
    shift 4
    local message="$*"

    if is_no_color_enabled; then
        printf "%s %s\n" "$symbol" "$prefix: $message" >&"$stream"
    else
        printf "%b%s%b %s\n" "$color" "$symbol" "$COLOR_RESET" "$prefix: $message" >&"$stream"
    fi
}

# Unified logging functions
debug() {
    _log_with_color "" "•" "Debug" 1 "$@"
}

info() {
    _log_with_color "$COLOR_CYAN" "ℹ" "Info" 1 "$@"
}

warn() {
    _log_with_color "$COLOR_YELLOW" "⚠" "Warning" 2 "$@"
}

error() {
    _log_with_color "$COLOR_RED" "✗" "Error" 2 "$@"
}

success() {
    _log_with_color "$COLOR_GREEN" "✓" "Success" 1 "$@"
}

# Additional utility functions
progress() {
    _log_with_color "$COLOR_BLUE" "→" "Progress" 1 "$@"
}

# Format step message with consistent styling
format_step_message() {
    local step="$1"
    local message="$2"
    if is_no_color_enabled; then
        printf '[%s] %s\n' "$step" "$message"
    else
        printf '%b[%s]%b %s\n' "$COLOR_BLUE" "$step" "$COLOR_RESET" "$message"
    fi
}

# Export color variables for use in other scripts
export COLOR_RESET COLOR_WHITE COLOR_BLUE COLOR_CYAN COLOR_YELLOW COLOR_RED COLOR_GREEN

# Export for child scripts
export is_no_color_enabled debug info warn error success progress format_step_message
