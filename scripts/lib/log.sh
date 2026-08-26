#!/usr/bin/env bash

# Terminal output for the scripts/ commands: colour decision, log lines, step counters and
# the banner box. The only place in this tree that spells an ANSI escape.

# Idempotent so a second source cannot fail on readonly, and overridable for a capture run.
: "${COLOR_RESET:=\033[0m}"
: "${COLOR_BLUE:=\033[34m}"
: "${COLOR_CYAN:=\033[36m}"
: "${COLOR_YELLOW:=\033[33m}"
: "${COLOR_RED:=\033[31m}"
: "${COLOR_GREEN:=\033[32m}"

# NO_COLOR and CI are honoured as the literal "true"; a non-terminal stdout disables colour
# on its own, which is what makes every captured run plain.
is_no_color_enabled() {
    [[ "${NO_COLOR:-}" == "true" || "${CI:-}" == "true" || ! -t 1 ]]
}

_log_with_color() {
    local -r color="$1" symbol="$2" prefix="$3" stream="$4"
    shift 4

    if is_no_color_enabled; then
        printf '%s %s\n' "$symbol" "$prefix: $*" >&"$stream"
        return
    fi
    printf '%b%s%b %s\n' "$color" "$symbol" "$COLOR_RESET" "$prefix: $*" >&"$stream"
}

info() { _log_with_color "$COLOR_CYAN" 'ℹ' 'Info' 1 "$@"; }
success() { _log_with_color "$COLOR_GREEN" '✓' 'Success' 1 "$@"; }
progress() { _log_with_color "$COLOR_BLUE" '→' 'Progress' 1 "$@"; }
warn() { _log_with_color "$COLOR_YELLOW" '⚠' 'Warning' 2 "$@"; }
error() { _log_with_color "$COLOR_RED" '✗' 'Error' 2 "$@"; }

# fail <reason> [note...] — report and return 1, so a guard is one line at the call site.
# It returns rather than exits: set -e in the entry script carries it to the exit status.
fail() {
    error "$1"
    shift
    local note
    for note in "$@"; do info "$note"; done
    return 1
}

format_step_message() {
    local -r step="$1" message="$2"

    if is_no_color_enabled; then
        printf '[%s] %s\n' "$step" "$message"
        return
    fi
    printf '%b[%s]%b %s\n' "$COLOR_BLUE" "$step" "$COLOR_RESET" "$message"
}

# Centre one line inside the box.
_banner_line() {
    local -r text="$1" width="$2"
    local -r space=$((width - ${#text})) left=$(((width - ${#text}) / 2))
    printf '%*s%s%*s' "$left" '' "$text" "$((space - left))" ''
}

# Every box is this wide so they all read the same; the longest title in the tree is 27.
readonly BANNER_WIDTH=38

_banner_body() {
    local -r color="$1" reset="$2"
    shift 2
    local l
    for l in "$@"; do
        printf '%b' "$color"
        _banner_line "$l" "$BANNER_WIDTH"
        printf '%b\n' "$reset"
    done
}

# wnc_banner <primary> [secondary]
wnc_banner() {
    local lines=("$1")
    [[ -z "${2:-}" ]] || lines+=("$2")

    local border color='' reset=''
    border=$(printf '%*s' "$BANNER_WIDTH" '' | tr ' ' '=')
    is_no_color_enabled || { color="$COLOR_BLUE"; reset="$COLOR_RESET"; }

    printf '%b\n' "${color}${border}${reset}"
    _banner_body "$color" "$reset" "${lines[@]}"
    printf '%b\n' "${color}${border}${reset}"
    printf '\n'
}
