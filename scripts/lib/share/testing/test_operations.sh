#!/usr/bin/env bash

# Cisco WNC Testing Library
# Provides unit tests, integration tests, and coverage validation

# Default values (idempotent)
: "${DEFAULT_COVERAGE_THRESHOLD:=90}"
readonly DEFAULT_COVERAGE_THRESHOLD
: "${DEFAULT_TEST_TIMEOUT:=10m}"
readonly DEFAULT_TEST_TIMEOUT
: "${DEFAULT_COVERAGE_FILE:=./tmp/coverage.out}"
readonly DEFAULT_COVERAGE_FILE
: "${DEFAULT_COVERAGE_HTML:=./tmp/coverage.html}"
readonly DEFAULT_COVERAGE_HTML

# Check if gotestsum command is available
has_gotestsum() {
    command_exists gotestsum
}

# Execute unit tests using appropriate test runner
_run_unit_tests_with_tool() {
    local project_root="$1"
    local env_vars=("${@:2:2}")
    local test_cmd_args=("${@:4}")

    if has_gotestsum; then
        printf '%s\n' "Using gotestsum for enhanced test output..."
        (cd "$project_root" && env "${env_vars[@]}" gotestsum --format testname -- "${test_cmd_args[@]}" ./...)
        return
    fi

    printf '%s\n' "gotestsum not found, using go test with verbose output..."
    test_cmd_args+=("-v")
    (cd "$project_root" && env "${env_vars[@]}" go test "${test_cmd_args[@]}" ./...)
}

# Execute unit tests with race detection and environment isolation
run_unit_tests() {
    local project_root="${1:-.}"
    local race_flag="${2:-true}"

    printf '%s\n' "Running unit tests only (no environment variables required)..."

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

    _run_unit_tests_with_tool "$project_root" "${env_vars[@]}" "${test_cmd_args[@]}"
    printf '✓ %s\n' "Unit tests passed"
}

# Check if integration environment is properly configured
has_integration_environment() {
    local controller="${WNC_CONTROLLER:-}"
    local token="${WNC_ACCESS_TOKEN:-}"
    [[ -n "$controller" && -n "$token" ]]
}

# Show integration environment configuration error
_show_integration_env_error() {
    printf '✗ %s\n' "Integration test environment not configured" >&2
    printf '  %s\n' "Set both WNC_CONTROLLER and WNC_ACCESS_TOKEN" >&2
    printf '  %s\n' "Example: export WNC_CONTROLLER=<controller-host-or-ip>" >&2
    printf '  %s\n' "Example: export WNC_ACCESS_TOKEN=<base64-username:password>" >&2
    printf '  %s\n' "Or:      export WNC_CONTROLLER=<controller-host-or-ip>" >&2
    printf '          %s\n' "export WNC_ACCESS_TOKEN=<base64-username:password>" >&2
}

# Check if required environment variables for integration tests are set
validate_integration_environment() {
    if has_integration_environment; then
        printf '✓ %s\n' "Using WNC_CONTROLLER and WNC_ACCESS_TOKEN for integration tests"
        return 0
    fi

    _show_integration_env_error
    return 1
}

# Execute integration tests using appropriate test runner
_run_integration_tests_with_tool() {
    local project_root="$1"
    local test_cmd_args=("${@:2}")

    if has_gotestsum; then
        printf '%s\n' "Using gotestsum for enhanced test output..."
        (cd "$project_root" && gotestsum --format testname -- "${test_cmd_args[@]}" ./...)
        return
    fi

    printf '%s\n' "gotestsum not found, using go test with verbose output..."
    test_cmd_args+=("-v")
    (cd "$project_root" && go test "${test_cmd_args[@]}" ./...)
}

# Run integration tests with environment variable validation
run_integration_tests() {
    local project_root="${1:-.}"
    local race_flag="${2:-true}"

    printf '%s\n' "Running integration tests (requires environment variables)..."

    # Validate environment first
    validate_integration_environment

    local test_cmd_args=()
    test_cmd_args+=("-timeout" "$DEFAULT_TEST_TIMEOUT")

    if [[ "$race_flag" == "true" ]]; then
        test_cmd_args+=("-race")
    fi

    _run_integration_tests_with_tool "$project_root" "${test_cmd_args[@]}"
    printf '✓ %s\n' "Integration tests passed"
}

# Check if coverage meets the required threshold
_coverage_meets_threshold() {
    local coverage_percent="$1"
    local threshold="$2"
    (( $(printf '%s' "$coverage_percent >= $threshold" | bc -l) ))
}

# Show coverage threshold error
_show_coverage_threshold_error() {
    local coverage_percent="$1"
    local threshold="$2"
    printf '✗ Coverage threshold not met: %s%% < %s%%\n' "$coverage_percent" "$threshold" >&2
    printf '  Required coverage: %s%%\n' "$threshold" >&2
}

# Validate coverage against threshold
validate_coverage_threshold() {
    local coverage_file="$1"
    local threshold="$2"

    # Show coverage summary
    local coverage_summary
    coverage_summary=$(go tool cover -func="$coverage_file" | tail -1)
    printf 'Coverage Summary: %s\n' "$coverage_summary"

    # Extract coverage percentage and validate threshold
    local coverage_percent
    coverage_percent=$(printf '%s' "$coverage_summary" | grep -o '[0-9.]*%' | sed 's/%//')

    if _coverage_meets_threshold "$coverage_percent" "$threshold"; then
        printf '✓ Coverage threshold met: %s%% >= %s%%\n' "$coverage_percent" "$threshold"
        return 0
    fi

    _show_coverage_threshold_error "$coverage_percent" "$threshold"
    return 1
}

# Execute coverage tests using appropriate test runner
_run_coverage_tests_with_tool() {
    local project_root="$1"
    local packages="$2"
    local test_cmd_args=("${@:3}")

    if has_gotestsum; then
        printf '%s\n' "Using gotestsum for enhanced test output..."
        (cd "$project_root" && gotestsum --format testname -- "${test_cmd_args[@]}" "$packages")
        return
    fi

    printf '%s\n' "gotestsum not found, using go test with verbose output..."
    test_cmd_args+=("-v")
    (cd "$project_root" && go test "${test_cmd_args[@]}" "$packages")
}

# Execute tests with coverage collection
execute_tests_with_coverage() {
    local project_root="$1"
    local coverage_file="$2"
    local race_flag="$3"

    local test_cmd_args=()
    test_cmd_args+=("-timeout" "$DEFAULT_TEST_TIMEOUT")
    test_cmd_args+=("-coverprofile=$coverage_file")

    if [[ "$race_flag" == "true" ]]; then
        test_cmd_args+=("-race")
    fi

    # Get packages excluding /internal/
    local packages
    packages=$(cd "$project_root" && go list ./... | grep -v '/internal/')

    _run_coverage_tests_with_tool "$project_root" "$packages" "${test_cmd_args[@]}"
}

# Run tests with coverage analysis and threshold validation
run_tests_with_coverage() {
    local project_root="${1:-.}"
    local coverage_file="${2:-$DEFAULT_COVERAGE_FILE}"
    local race_flag="${3:-true}"
    local threshold="${4:-$DEFAULT_COVERAGE_THRESHOLD}"

    printf '%s\n' "Running tests with coverage analysis..."

    # Ensure tmp directory exists
    ensure_tmp_directory "$project_root"

    # Resolve relative path to absolute
    if [[ ! "$coverage_file" = /* ]]; then
        coverage_file="$project_root/$coverage_file"
    fi

    # Execute the tests
    execute_tests_with_coverage "$project_root" "$coverage_file" "$race_flag"

    if [[ -f "$coverage_file" ]]; then
        printf '✓ Coverage report generated: %s\n' "$coverage_file"
        validate_coverage_threshold "$coverage_file" "$threshold"
        return
    fi

    printf '✗ Coverage file not generated: %s\n' "$coverage_file" >&2
    return 1
}

# Generate HTML coverage report from coverage data file
generate_html_coverage() {
    local project_root="${1:-.}"
    local coverage_file="${2:-$DEFAULT_COVERAGE_FILE}"
    local html_file="${3:-$DEFAULT_COVERAGE_HTML}"

    printf '%s\n' "Generating HTML coverage report..."

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
        printf '✗ Coverage file not found: %s\n' "$coverage_file" >&2
        printf '  %s\n' "Run tests with coverage first" >&2
        return 1
    fi

    go tool cover -html="$coverage_file" -o "$html_file"

    if [[ -f "$html_file" ]]; then
        printf '✓ HTML coverage report generated: %s\n' "$html_file"
        return
    fi

    printf '✗ %s\n' "Failed to generate HTML coverage report" >&2
    return 1
}

# Execute both unit and integration tests based on environment availability
run_all_tests() {
    local project_root="${1:-.}"
    local race_flag="${2:-true}"

    printf '%s\n' "Running all tests..."

    # Always run unit tests
    run_unit_tests "$project_root" "$race_flag"

    # Run integration tests if environment is configured
    if validate_integration_environment >/dev/null 2>&1; then
        printf '%s\n' "Integration environment detected, running integration tests..."
        run_integration_tests "$project_root" "$race_flag"
        printf '✓ %s\n' "All available tests passed"
        return
    fi

    printf '%s\n' "Integration environment not configured, skipping integration tests"
    printf '✓ %s\n' "All available tests passed"
}

# Verify basic calculator command availability for coverage calculations
ensure_bc_available() {
    if ! command_exists bc; then
        printf '%s\n' "Warning: 'bc' command not found. Coverage threshold validation may not work properly." >&2
        printf '%s\n' "Install bc with: brew install bc (macOS) or apt-get install bc (Ubuntu)" >&2
    fi
}

# Initialize bc check when library is loaded
ensure_bc_available
