#!/usr/bin/env bash

# Cisco WNC Testing - Output and formatting functions
# Provides functions to format and display test results and messages
# (bootstrap handled by parent script)

# Check if test count information is available
has_test_count() { [[ -n "${1:-}" ]]; }

# Check if duration information is meaningful
has_meaningful_duration() { [[ "${1:-unknown}" != "unknown" ]]; }

# Display comprehensive test execution summary
display_test_summary() {
    local test_type="$1"
    local exit_code="$2"
    local duration="${3:-unknown}"
    local test_count="${4:-}"

    printf '\n'
    printf '%s\n' "-----------------------------------------"

    if [[ "$exit_code" -eq 0 ]]; then
        success "$test_type tests completed successfully"

        if has_test_count "$test_count"; then
            info "Tests executed: $test_count"
        fi
    fi

    if [[ "$exit_code" -ne 0 ]]; then
        error "$test_type tests failed"
        info "Check the output above for details"
    fi

    if has_meaningful_duration "$duration"; then
        info "Duration: $duration"
    fi
    printf '%s\n' "-----------------------------------------"
}

display_coverage_summary() {
    local coverage_file="$1"
    local exit_code="$2"

    # Early returns to avoid deep nesting
    if [[ "$exit_code" -ne 0 ]]; then
        return 0
    fi
    if [[ ! -f "$coverage_file" ]]; then
        return 0
    fi

    printf '\n'
    info "Coverage report generated: $coverage_file"

    # Extract coverage percentage if available
    if command -v go >/dev/null 2>&1; then
        local coverage_percent
        # Use END block instead of tail -1 for robustness
        coverage_percent=$(go tool cover -func="$coverage_file" 2>/dev/null | awk 'END {print $3}' || printf 'unknown')
        if [[ "$coverage_percent" != "unknown" ]]; then
            info "Total coverage: $coverage_percent"
        fi
    fi
}
