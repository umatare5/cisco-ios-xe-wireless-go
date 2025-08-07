#!/usr/bin/env bash

# Cisco WNC Testing Library
# Functions for running various types of tests with coverage analysis

# Prevent double sourcing
if [[ -n "${WNC_TESTING_LOADED:-}" ]]; then
    return 0
fi
readonly WNC_TESTING_LOADED=1

set -euo pipefail

# Source required libraries
SCRIPT_LIB_DIR="$(dirname "${BASH_SOURCE[0]}")"
readonly SCRIPT_LIB_DIR

source "${SCRIPT_LIB_DIR}/build_tools.sh"

# Default values (if not already defined)
if [[ -z "${DEFAULT_COVERAGE_THRESHOLD:-}" ]]; then
    readonly DEFAULT_COVERAGE_THRESHOLD=90
fi
if [[ -z "${DEFAULT_TEST_TIMEOUT:-}" ]]; then
    readonly DEFAULT_TEST_TIMEOUT="10m"
fi
if [[ -z "${DEFAULT_COVERAGE_FILE:-}" ]]; then
    readonly DEFAULT_COVERAGE_FILE="./tmp/coverage.out"
fi
if [[ -z "${DEFAULT_COVERAGE_HTML:-}" ]]; then
    readonly DEFAULT_COVERAGE_HTML="./tmp/coverage.html"
fi

# Function to run unit tests only (no environment variables required)
run_unit_tests() {
    local project_root="${1:-.}"
    local race_flag="${2:-true}"

    echo "Running unit tests only (no environment variables required)..."

    # Build test command arguments
    local test_cmd_args=()
    test_cmd_args+=("-timeout" "$DEFAULT_TEST_TIMEOUT")

    if [[ "$race_flag" == "true" ]]; then
        test_cmd_args+=("-race")
    fi

    # Clear environment variables to ensure unit test isolation
    local env_vars=(
        "WNC_CONTROLLER="
        "WNC_ACCESS_TOKEN="
    )

    # Use gotestsum for better test output
    if command_exists gotestsum; then
        echo "Using gotestsum for enhanced test output..."
        (cd "$project_root" && env "${env_vars[@]}" gotestsum --format testname -- "${test_cmd_args[@]}" ./...)
    else
        echo "gotestsum not found, using go test with verbose output..."
        test_cmd_args+=("-v")
        (cd "$project_root" && env "${env_vars[@]}" gotestsum --format testname -- "${test_cmd_args[@]}" ./...)
    fi

    echo "✓ Unit tests passed"
}

# Function to validate integration test environment
validate_integration_environment() {
    local controller="${WNC_CONTROLLER:-}"
    local token="${WNC_ACCESS_TOKEN:-}"

    # Check if we have both WNC_CONTROLLER and WNC_ACCESS_TOKEN
    if [[ -n "$controller" && -n "$token" ]]; then
        echo "✓ Using WNC_CONTROLLER and WNC_ACCESS_TOKEN for integration tests"
        return 0
    else
        echo "✗ Integration test environment not configured" >&2
        echo "  Set both WNC_CONTROLLER and WNC_ACCESS_TOKEN" >&2
        echo "  Example: export WNC_CONTROLLER=wnc1.example.internal" >&2
        echo "  Example: export WNC_ACCESS_TOKEN=YWRtaW46Y3l0WU43WVh4M2swc3piUnVhb1V1ZUx6" >&2
        echo "  Or:      export WNC_CONTROLLER=wnc1.example.internal" >&2
        echo "          export WNC_ACCESS_TOKEN=YWRtaW46Y3l0WU43WVh4M2swc3piUnVhb1V1ZUx6" >&2
        return 1
    fi
}

# Function to run integration tests
run_integration_tests() {
    local project_root="${1:-.}"
    local race_flag="${2:-true}"

    echo "Running integration tests (requires environment variables)..."

    # Validate environment first
    validate_integration_environment

    local test_cmd_args=()
    test_cmd_args+=("-timeout" "$DEFAULT_TEST_TIMEOUT")

    if [[ "$race_flag" == "true" ]]; then
        test_cmd_args+=("-race")
    fi

    if command_exists gotestsum; then
        echo "Using gotestsum for enhanced test output..."
        (cd "$project_root" && gotestsum --format testname -- "${test_cmd_args[@]}" ./...)
    else
        echo "gotestsum not found, using go test with verbose output..."
        test_cmd_args+=("-v")
        (cd "$project_root" && gotestsum --format testname -- "${test_cmd_args[@]}" ./...)
    fi

    echo "✓ Integration tests passed"
}

# Function to run tests with coverage
run_tests_with_coverage() {
    local project_root="${1:-.}"
    local coverage_file="${2:-$DEFAULT_COVERAGE_FILE}"
    local race_flag="${3:-true}"
    local threshold="${4:-$DEFAULT_COVERAGE_THRESHOLD}"

    echo "Running tests with coverage analysis..."

    # Ensure tmp directory exists
    ensure_tmp_directory "$project_root"

    # Resolve relative path to absolute
    if [[ ! "$coverage_file" = /* ]]; then
        coverage_file="$project_root/$coverage_file"
    fi

    local test_cmd_args=()
    test_cmd_args+=("-timeout" "$DEFAULT_TEST_TIMEOUT")
    test_cmd_args+=("-coverprofile=$coverage_file")

    if [[ "$race_flag" == "true" ]]; then
        test_cmd_args+=("-race")
    fi

    # Get packages excluding /internal/
    local packages
    packages=$(cd "$project_root" && go list ./... | grep -v '/internal/')

    if command_exists gotestsum; then
        echo "Using gotestsum for enhanced test output..."
        # shellcheck disable=SC2086
        (cd "$project_root" && gotestsum --format testname -- "${test_cmd_args[@]}" $packages)
    else
        echo "gotestsum not found, using go test with verbose output..."
        test_cmd_args+=("-v")
        # shellcheck disable=SC2086
        (cd "$project_root" && gotestsum --format testname -- "${test_cmd_args[@]}" $packages)
    fi

    if [[ -f "$coverage_file" ]]; then
        echo "✓ Coverage report generated: $coverage_file"

        # Show coverage summary
        local coverage_summary
        coverage_summary=$(go tool cover -func="$coverage_file" | tail -1)
        echo "Coverage Summary: $coverage_summary"

        # Extract coverage percentage and validate threshold
        local coverage_percent
        coverage_percent=$(echo "$coverage_summary" | grep -o '[0-9.]*%' | sed 's/%//')

        if (( $(echo "$coverage_percent >= $threshold" | bc -l) )); then
            echo "✓ Coverage threshold met: ${coverage_percent}% >= ${threshold}%"
        else
            echo "✗ Coverage threshold not met: ${coverage_percent}% < ${threshold}%" >&2
            echo "  Required coverage: ${threshold}%" >&2
            return 1
        fi
    else
        echo "✗ Coverage file not generated: $coverage_file" >&2
        return 1
    fi
}

# Function to generate HTML coverage report
generate_html_coverage() {
    local project_root="${1:-.}"
    local coverage_file="${2:-$DEFAULT_COVERAGE_FILE}"
    local html_file="${3:-$DEFAULT_COVERAGE_HTML}"

    echo "Generating HTML coverage report..."

    # Ensure tmp directory exists
    ensure_tmp_directory "$project_root"

    # Resolve relative paths to absolute
    if [[ ! "$coverage_file" = /* ]]; then
        coverage_file="$project_root/$coverage_file"
    fi
    if [[ ! "$html_file" = /* ]]; then
        html_file="$project_root/$html_file"
    fi

    if [[ ! -f "$coverage_file" ]]; then
        echo "✗ Coverage file not found: $coverage_file" >&2
        echo "  Run tests with coverage first" >&2
        return 1
    fi

    go tool cover -html="$coverage_file" -o "$html_file"

    if [[ -f "$html_file" ]]; then
        echo "✓ HTML coverage report generated: $html_file"
    else
        echo "✗ Failed to generate HTML coverage report" >&2
        return 1
    fi
}

# Function to run all tests (unit + integration if environment is configured)
run_all_tests() {
    local project_root="${1:-.}"
    local race_flag="${2:-true}"

    echo "Running all tests..."

    # Always run unit tests
    run_unit_tests "$project_root" "$race_flag"

    # Run integration tests if environment is configured
    if validate_integration_environment >/dev/null 2>&1; then
        echo "Integration environment detected, running integration tests..."
        run_integration_tests "$project_root" "$race_flag"
    else
        echo "Integration environment not configured, skipping integration tests"
    fi

    echo "✓ All available tests passed"
}

# Function to check if bc (basic calculator) is available for coverage comparison
ensure_bc_available() {
    if ! command_exists bc; then
        echo "Warning: 'bc' command not found. Coverage threshold validation may not work properly." >&2
        echo "Install bc with: brew install bc (macOS) or apt-get install bc (Ubuntu)" >&2
    fi
}

# Initialize bc check when library is loaded
ensure_bc_available
