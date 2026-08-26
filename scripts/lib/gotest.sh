#!/usr/bin/env bash

# The gotestsum runs behind test-unit, test-unit-coverage and test-integration. The three
# differ only in argv and in one summary line, so they share one entry point.
# Requires lib/log.sh and lib/env.sh.

# Integration runs need an AP as well as a controller, which the flag help does not say.
_has_integration_env() {
    [[ -n "${WNC_CONTROLLER:-}" && -n "${WNC_ACCESS_TOKEN:-}" && -n "${WNC_AP_MAC_ADDR:-}" ]]
}

_assert_integration_env() {
    _has_integration_env && return 0
    fail "Integration tests require WNC environment variables" \
        "Set WNC_CONTROLLER, WNC_ACCESS_TOKEN, and WNC_AP_MAC_ADDR environment variables" \
        "Optionally set WNC_CLIENT_MAC_ADDR for enhanced client testing"
}

validate_test_environment() {
    local -r project_root="$1" test_type="$2"

    dir_exists "$project_root" || fail "Invalid project directory: $project_root"
    file_exists "$project_root/go.mod" ||
        fail "No go.mod found in $project_root" "This directory doesn't appear to be a Go module"
    command_exists go || fail "Go toolchain is not installed or not in PATH"
    [[ "$test_type" != 'integration' ]] || _assert_integration_env
}

# A unit run is local and quick; an integration run crosses the network to a controller.
_timeout_for() {
    case "$1" in
    unit) printf '30s\n' ;;
    integration) printf '10m\n' ;;
    *) printf '60s\n' ;;
    esac
}

# _test_argv <type> [coverage_file] — one argument per line.
# Integration runs carry the build tag; the others carry -race, matching what CI enables.
_test_argv() {
    local -r test_type="$1" coverage_file="${2:-}"

    printf '%s\n' -v -timeout "$(_timeout_for "$test_type")"
    if [[ "$test_type" == 'integration' ]]; then
        printf '%s\n' -tags=integration
    else
        printf '%s\n' -race
    fi
    [[ -z "$coverage_file" ]] || printf '%s\n' "-coverprofile=$coverage_file" -covermode=atomic
}

_packages_for() {
    if [[ "$1" == 'integration' ]]; then
        printf './tests/integration/...\n'
        return
    fi
    printf './...\n'
}

_summary() {
    local -r label="$1" exit_code="$2" duration="$3"

    printf '\n%s\n' "-----------------------------------------"
    if [[ "$exit_code" -ne 0 ]]; then
        error "$label tests failed"
        info "Check the output above for details"
    else
        success "$label tests completed successfully"
    fi
    info "Duration: $duration"
    printf '%s\n' "-----------------------------------------"
}

_coverage_summary() {
    local -r coverage_file="$1"
    local percent

    file_exists "$coverage_file" || return 0
    printf '\n'
    info "Coverage report generated: $coverage_file"
    percent=$(go tool cover -func="$coverage_file" 2>/dev/null | awk 'END {print $3}' || printf 'unknown')
    [[ "$percent" == 'unknown' ]] || info "Total coverage: $percent"
}

# Capitalise the type for the banner and the summary line.
_label_for() {
    printf '%s%s\n' "$(tr '[:lower:]' '[:upper:]' <<<"${1:0:1}")" "${1:1}"
}

# gotestsum runs from the project root, so a relative coverprofile must be anchored first.
_absolute() {
    if [[ -z "$2" || "$2" == /* ]]; then
        printf '%s\n' "$2"
        return
    fi
    printf '%s/%s\n' "$1" "$2"
}

# ./tmp holds the default coverprofile and is gitignored, so a fresh clone lacks it.
_invoke_gotestsum() {
    local -r project_root="$1" test_type="$2"
    shift 2
    (cd "$project_root" && mkdir -p ./tmp &&
        gotestsum --format testname -- "$@" "$(_packages_for "$test_type")")
}

# gotest_run <unit|coverage|integration> <project_root> [coverage_file]
gotest_run() {
    local -r test_type="$1" project_root="$2"
    local -r coverage_file="$(_absolute "$project_root" "${3:-}")"
    local -r label="$(_label_for "$test_type")"
    local args=() a start exit_code=0

    wnc_banner "Cisco WNC ${label} Tests" "Go Testing Framework"
    validate_test_environment "$project_root" "$test_type"
    while IFS= read -r a; do args+=("$a"); done < <(_test_argv "$test_type" "$coverage_file")

    progress "Starting ${test_type} tests..."
    start=$(date +%s)
    _invoke_gotestsum "$project_root" "$test_type" "${args[@]}" || exit_code=$?
    _summary "$label" "$exit_code" "$(($(date +%s) - start))s"

    [[ -n "$coverage_file" && "$exit_code" -eq 0 ]] && _coverage_summary "$coverage_file"
    return "$exit_code"
}
