#!/usr/bin/env bash

# Cisco WNC Core Testing Library
# Provides core functions for running tests in the Cisco WNC project

# Check if custom timeout is provided
has_custom_timeout() {
    [[ -n "$1" ]]
}

# Add timeout arguments based on test type and custom timeout
_add_timeout_args() {
    local timeout="$1"
    local test_type="$2"
    local args_var_name="$3"

    # Use custom timeout if provided
    if has_custom_timeout "$timeout"; then
        eval "${args_var_name}=\"\${${args_var_name}} -timeout $timeout\""
        return
    fi

    # Apply default timeouts based on test type
    case "$test_type" in
        "unit") eval "${args_var_name}=\"\${${args_var_name}} -timeout 30s\"" ;;
        "coverage") eval "${args_var_name}=\"\${${args_var_name}} -timeout 60s\"" ;;
        "integration") eval "${args_var_name}=\"\${${args_var_name}} -timeout 120s\"" ;;
    esac
}

# Add coverage-specific arguments
# Provides core business logic for running tests in the Cisco WNC project

# Source output functions for display (bootstrap is handled by parent script)
source "$(dirname "${BASH_SOURCE[0]}")/output.sh"
source "$(dirname "${BASH_SOURCE[0]}")/help.sh"

# Check if required integration environment variables are present
_has_required_integration_env() {
    [[ -n "${WNC_CONTROLLER:-}" && -n "${WNC_ACCESS_TOKEN:-}" ]]
}

# Check if test type is integration test
_is_integration_test() {
    local test_type="$1"
    [[ "$test_type" == "integration" ]]
}

# Check if integration test requires environment validation
_should_validate_integration_env() {
    local test_type="$1"
    _is_integration_test "$test_type" && ! is_skip_env_check_enabled
}

# Check if browser should be opened automatically for coverage reports
_should_open_browser_automatically() {
    [[ "${OPEN_BROWSER:-false}" == "true" ]]
}

# Check if browser open command is available
_has_browser_open_command() {
    command -v open >/dev/null 2>&1
}

# Check if browser should be opened and is available
_can_open_browser() {
    _should_open_browser_automatically && _has_browser_open_command
}

validate_test_environment() {
    local project_root="$1"
    local test_type="$2"

    # Validate project directory
    if ! validate_project_directory "$project_root"; then
        error "Invalid project directory: $project_root"
        return 1
    fi

    # Check if it's a Go module
    if [[ ! -f "$project_root/go.mod" ]]; then
        error "No go.mod found in $project_root"
        info "This directory doesn't appear to be a Go module"
        return 1
    fi

    # Check if Go is available
    if ! is_command_available go; then
        error "Go toolchain is not installed or not in PATH"
        return 1
    fi

    # For integration tests, check environment variables if validation enabled
    if _should_validate_integration_env "$test_type"; then
        if ! _has_required_integration_env; then
            error "Integration tests require WNC environment variables"
            info "Set WNC_CONTROLLER and WNC_ACCESS_TOKEN environment variables"
            return 1
        fi
    fi

    return 0
}

# Set appropriate timeout for test type
_set_test_timeout() {
    local test_type="$1"
    local timeout="$2"
    local args_var_name="$3"

    if [[ -n "$timeout" ]]; then
        eval "${args_var_name}=\"\${${args_var_name}} -timeout $timeout\""
        return
    fi

    case "$test_type" in
        "unit") eval "${args_var_name}=\"\${${args_var_name}} -timeout 30s\"" ;;
        "coverage") eval "${args_var_name}=\"\${${args_var_name}} -timeout 60s\"" ;;
        "integration") eval "${args_var_name}=\"\${${args_var_name}} -timeout 120s\"" ;;
    esac
}

# Add coverage-specific arguments
_add_coverage_args() {
    local coverage_file="$1"
    local args_var_name="$2"

    local actual_coverage_file="${coverage_file:-./tmp/coverage.out}"
    eval "${args_var_name}=\"\${${args_var_name}} -coverprofile=$actual_coverage_file\""
    eval "${args_var_name}=\"\${${args_var_name}} -covermode=atomic\""
    # is_verbose_enabled && info "Coverage output: $actual_coverage_file"  # Temporarily disabled
}

# Prepare test command arguments based on test type
prepare_test_arguments() {
    local test_type="$1"
    local timeout="${2:-}"
    local result_var_name="$3"
    local coverage_file="${4:-}"

    # Build argument string
    local args="-v"  # Verbose

    # Set timeout
    _set_test_timeout "$test_type" "$timeout" "args"

    # Short mode
    if is_short_mode_enabled; then
        args="$args -short"
        # is_verbose_enabled && info "Running in short mode (skipping long tests)"  # Temporarily disabled
    fi

    # Coverage-specific options
    if [[ "$test_type" == "coverage" ]]; then
        _add_coverage_args "$coverage_file" "args"
    fi

    # Use eval to assign to caller's variable
    if [[ -n "$result_var_name" ]]; then
        eval "$result_var_name=\"\$args\""
        return 0
    fi

    echo "$args"
    return 0
}

execute_test_command() {
    local project_root="$1"
    local test_type="$2"
    local test_args="$3"

    # Change to project directory
    local original_pwd="$PWD"
    cd "$project_root" || {
        error "Failed to change to project directory: $project_root"
        return 1
    }

    # Create tmp directory if needed
    if [[ ! -d "./tmp" ]]; then
        mkdir -p "./tmp"
    fi

    is_verbose_enabled && info "Running $test_type tests in: $project_root"
    is_verbose_enabled && info "Test arguments: $test_args"

    # Execute tests
    local start_time
    start_time=$(date +%s)

    local exit_code=0
    case "$test_type" in
        "unit"|"coverage")
            eval "gotestsum --format testname -- $test_args ./..." || exit_code=$?
            ;;
        "integration")
            # Integration tests with special tags
            eval "gotestsum --format testname -- $test_args -tags=integration ./..." || exit_code=$?
            ;;
    esac

    local end_time
    end_time=$(date +%s)
    local duration=$((end_time - start_time))

    # Return to original directory
    cd "$original_pwd" || {
        warn "Failed to return to original directory"
    }

    # Store duration for reporting
    TEST_DURATION="${duration}s"

    return "$exit_code"
}


# Generate HTML coverage report if requested
_generate_html_report_if_requested() {
    local coverage_file="$1"
    local exit_code="$2"

    if [[ "$exit_code" -ne 0 ]] || [[ "${HTML_COVERAGE:-false}" != "true" ]]; then
        return 0
    fi

    local html_file="${coverage_file%.out}.html"
    progress "Generating HTML coverage report..."

    if go tool cover -html="$coverage_file" -o "$html_file"; then
        success "HTML report generated: $html_file"

        # Open in browser if requested and available
        if _can_open_browser; then
            open "$html_file"
            info "Opened HTML report in browser"
        fi
        return
    fi

    warn "Failed to generate HTML report"
}

# Execute coverage tests with all setup and reporting
run_coverage_test_operation() {
    local project_root="$1"
    local coverage_file="$2"

    # Convert relative path to absolute if needed
    if [[ "$coverage_file" != /* ]]; then
        coverage_file="$project_root/$coverage_file"
    fi

    # Show banner and info
    show_test_banner "Coverage"
    show_test_verbose_info "Coverage" "$project_root" ""

    # Validate environment
    if ! validate_test_environment "$project_root" "coverage"; then
        return 1
    fi

    # Prepare arguments
    local test_args
    prepare_test_arguments "coverage" "" test_args "$coverage_file"

    # Execute tests
    progress "Starting coverage tests..."
    local exit_code=0
    execute_test_command "$project_root" "coverage" "$test_args" || exit_code=$?

    # Display results
    display_test_summary "Coverage" "$exit_code" "${TEST_DURATION:-unknown}"
    display_coverage_summary "$coverage_file" "$exit_code"

    # Generate HTML report if requested
    _generate_html_report_if_requested "$coverage_file" "$exit_code"

    return "$exit_code"
}
