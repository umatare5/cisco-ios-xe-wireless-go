#!/usr/bin/env bash
#
# Runs the unit tests with a coverprofile, then renders the HTML report and the artifact
# octocov reads. The three paths are positional so a verification run can redirect them:
#
#   test_coverage.sh [coverprofile] [html] [report]
#
# coverage/report.out is tracked and .octocov.yml reads it for the README badge, so the
# defaults write in place. Do not relocate them.

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

_assert_coverprofile() {
    local -r input_file="$1"

    file_exists "$input_file" || fail "Coverage file not found: $input_file"
    head -1 "$input_file" | grep -q 'mode:' ||
        warn "Input file may not be a valid Go coverage file"
}

_html_size() {
    stat -f%z "$1" 2>/dev/null || stat -c%s "$1" 2>/dev/null || printf 'unknown\n'
}

_announce_report() {
    local -r html_output="$1"

    printf '\n'
    success "HTML coverage report generated successfully"
    info "Report location: $html_output"
    file_exists "$html_output" && info "Report size: $(_html_size "$html_output") bytes"
    printf '\n'
    info "To view the report:"
    printf '%s\n' "  open $html_output"
}

# octocov reads report.out, so a missing mode header would fail the badge silently.
_write_artifact() {
    local -r coverage_file="$1" report_output="$2"

    cp "$coverage_file" "$report_output" 2>/dev/null ||
        warn "Failed to write coverprofile to $report_output"
    head -1 "$report_output" 2>/dev/null | grep -q '^mode:' ||
        warn "report.out missing mode header; octocov may fail"
}

_render_html() {
    local -r coverage_file="$1" html_output="$2"

    mkdir -p "$(dirname "$html_output")" || return 1
    go tool cover -html="$coverage_file" -o "$html_output" && return 0
    printf '\n'
    fail "Failed to generate HTML coverage report" "Check the input file and try again"
}

main() {
    # The three defaults are relative, so stand at the repo root before resolving them.
    cd "$PROJECT_ROOT" || fail "Cannot enter $PROJECT_ROOT"

    local -r coverage_file="${1:-./tmp/coverage.out}"
    local -r html_output="${2:-./coverage/report.html}"
    local -r report_output="${3:-./coverage/report.out}"

    validate_required_cli_tools "standard"
    gotest_run coverage "$PROJECT_ROOT" "$coverage_file"

    progress "Generating HTML coverage report..."
    _assert_coverprofile "$coverage_file"
    _render_html "$coverage_file" "$html_output"
    _write_artifact "$coverage_file" "$report_output"
    _announce_report "$html_output"
    success "Coverage report generated successfully"
}

main "$@"
