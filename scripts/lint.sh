#!/usr/bin/env bash

# Formats the working tree with goimports and gofumpt, then lints Go, shell and Markdown.
# Takes no arguments. The formatters rewrite files in place and nothing re-stages them.

set -Eeuo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
readonly SCRIPT_DIR
readonly PROJECT_ROOT="${SCRIPT_DIR%/*}"

# shellcheck source=lib/log.sh
source "${SCRIPT_DIR}/lib/log.sh"
# shellcheck source=lib/env.sh
source "${SCRIPT_DIR}/lib/env.sh"

# tmp/ is gitignored and holds this repository's git worktrees, so an unpruned sweep from the
# caller's cwd reaches source files belonging to other checkouts and rewrites them in place.
_repo_files() {
    find . \( -name vendor -o -name .git -o -name tmp \) -prune -o -type f -name "$1" -print
}

# _apply <tool> <pattern> <extra args...> — run one tool over every matching file.
_apply() {
    local -r tool="$1" pattern="$2"
    shift 2
    local files
    files=$(_repo_files "$pattern")
    [[ -n "$files" ]] || return 0
    printf '%s\n' "$files" | xargs "$tool" "$@"
}

# _format <tool> — every formatter here is "rewrite every Go file, or skip".
_format() {
    local -r tool="$1"
    printf '\n--- Go Code Formatting (%s) ---\n' "$tool" >&2

    command_exists "$tool" || {
        printf 'Warning: %s not found, skipping\n' "$tool" >&2
        return 0
    }
    _apply "$tool" '*.go' -w
}

_lint_go() {
    printf '\n--- Go Code Linting (golangci-lint) ---\n' >&2

    if ! command_exists golangci-lint; then
        printf '✗ %s\n' "golangci-lint is required for Go linting" >&2
        return 1
    fi

    golangci-lint run --timeout=10m
}

_lint_shell() {
    printf '\n--- Shell Script Linting (shellcheck) ---\n' >&2

    command_exists shellcheck || {
        printf 'Warning: shellcheck not found, skipping shell script linting\n' >&2
        return 0
    }
    _apply shellcheck '*.sh'
}

_lint_markdown() {
    printf '\n--- Markdown Linting (markdownlint-cli2) ---\n' >&2

    if ! command_exists markdownlint-cli2; then
        printf 'Warning: markdownlint-cli2 not found, skipping markdown linting\n' >&2
        return 0
    fi

    markdownlint-cli2
}

# _phase_summary <label> <exit-code> <success sentence>
_phase_summary() {
    local -r label="$1" exit_code="$2" ok="$3"

    printf '\n--- %s Summary ---\n' "$label" >&2
    if [[ "$exit_code" -ne 0 ]]; then
        printf '%s failed - please fix the issues above\n' "$label" >&2
        printf '\n' >&2
        return
    fi
    printf '%s\n' "$ok" >&2
    printf '\n' >&2
}

_assert_go_module() {
    local -r project_root="$1"

    printf '%s\n' "Validating Go module..."
    if ! file_exists "$project_root/go.mod"; then
        printf '✗ go.mod not found in %s\n' "$project_root" >&2
        return 1
    fi
    (cd "$project_root" && go mod verify)
    printf '✓ %s\n' "Go module validated"
}

_run_formats() {
    local status=0

    printf 'Running formatting operations...\n' >&2
    _format goimports || status=1
    _format gofumpt || status=1
    _phase_summary Formatting "$status" 'All formatting operations completed successfully'
    return "$status"
}

_run_lints() {
    local status=0

    printf 'Running linting operations...\n' >&2
    _lint_go || status=1
    _lint_shell || status=1
    _lint_markdown || status=1
    _phase_summary Linting "$status" 'All linting checks passed successfully'
    return "$status"
}

_announce() {
    [[ "$1" -eq 0 ]] || {
        error "Lint failed for $2"
        return 0
    }
    success "Lint passed for $2"
}

# Entrypoint: format then lint
main() {
    local status=0

    # _repo_files sweeps `find .`, so the caller's cwd must not decide what gets formatted.
    cd "$PROJECT_ROOT" || fail "Cannot enter $PROJECT_ROOT"
    validate_required_cli_tools "standard"

    wnc_banner "Cisco WNC Code Linter" "golangci-lint Integration"
    _assert_go_module "$PROJECT_ROOT" || return 1
    info "Starting code linting..."

    _run_formats || status=1
    _run_lints || status=1

    _announce "$status" .
    return "$status"
}

main "$@"
