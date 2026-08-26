#!/usr/bin/env bash

# Preflight checks shared by the scripts/ commands: the toolchain a command needs, and the
# directory it was pointed at. Requires lib/log.sh.

command_exists() { command -v "$1" >/dev/null 2>&1; }
dir_exists() { [[ -d "$1" ]]; }
file_exists() { [[ -f "$1" ]]; }

# minimal is what a command needs to bootstrap; standard adds the tools lint and the test
# runners shell out to. Anything else resolves to minimal.
_tools_for_level() {
    case "${1:-standard}" in
    standard) printf 'go golangci-lint gotestsum markdownlint-cli2 shellcheck\n' ;;
    *) printf 'go\n' ;;
    esac
}

_install_hint_for() {
    case "$1" in
    go) printf 'https://golang.org/dl/\n' ;;
    gotestsum) printf 'go install gotest.tools/gotestsum@latest\n' ;;
    golangci-lint)
        printf 'go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest\n'
        ;;
    shellcheck)
        printf 'brew install shellcheck (macOS) or sudo apt-get install shellcheck (Ubuntu)\n'
        ;;
    markdownlint-cli2) printf 'npm install -g markdownlint-cli2\n' ;;
    *) printf 'usually pre-installed\n' ;;
    esac
}

_show_install_hint() {
    local tool
    printf '\n%s\n\n' "Installation instructions for missing CLI tools:" >&2
    for tool in "$@"; do
        printf '  %s: %s\n' "$tool" "$(_install_hint_for "$tool")" >&2
    done
    printf '\n' >&2
}

# _report_tool <name> — ticks the tool off, and fails when it is absent.
_report_tool() {
    if command_exists "$1"; then
        printf '✓ %s\n' "$1" >&2
        return 0
    fi
    printf '✗ %s\n' "$1" >&2
    return 1
}

# validate_required_cli_tools <minimal|standard>
validate_required_cli_tools() {
    local -r level="${1:-standard}"
    local tools missing=() tool

    IFS=' ' read -ra tools <<<"$(_tools_for_level "$level")"
    printf 'Validating CLI tools (level: %s)...\n' "$level" >&2
    for tool in "${tools[@]}"; do
        _report_tool "$tool" || missing+=("$tool")
    done

    printf '\n' >&2
    _report_missing "${#tools[@]}" "${missing[@]+${missing[@]}}"
}

_report_missing() {
    local -r total="$1"
    shift

    if [[ $# -eq 0 ]]; then
        printf '✓ All %d required CLI tools are available\n' "$total" >&2
        return 0
    fi
    printf '✗ %d out of %d CLI tools are missing\n' "$#" "$total" >&2
    _show_install_hint "$@"
    return 1
}
