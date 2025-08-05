#!/usr/bin/env bash

# Cisco WNC Testing Operations - Core Functions
# Core business logic for testing operations

# Source common predicates
source "$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)/../common/common.sh"

# Global variables
declare -a TEST_ARGS

validate_test_environment() {
    local project_root="$1"
    local test_type="$2"

    # Validate project directory
    if ! validate_project_directory "$project_root"; then
        format_test_error "Invalid project directory: $project_root"
        return 1
    fi

    # Check if it's a Go module
    if [[ ! -f "$project_root/go.mod" ]]; then
        format_test_error "No go.mod found in $project_root"
        format_test_info "This directory doesn't appear to be a Go module"
        return 1
    fi

    # Check if Go is available
    if ! is_command_available go; then
        format_test_error "Go toolchain is not installed or not in PATH"
        return 1
    fi

    # For integration tests, check environment variables
    if [[ "$test_type" == "integration" ]] && ! is_skip_env_check_enabled; then
        if [[ -z "${WNC_CONTROLLER:-}" || -z "${WNC_ACCESS_TOKEN:-}" ]]; then
            format_test_error "Integration tests require WNC environment variables"
            format_test_info "Set WNC_CONTROLLER and WNC_ACCESS_TOKEN environment variables"
            return 1
        fi
    fi

    return 0
}

prepare_test_arguments() {
    local test_type="$1"

    # Create array of test arguments
    TEST_ARGS=()

    # Basic test arguments
    TEST_ARGS+=("-v")  # Verbose

    # Timeout
    if [[ -n "${argc_timeout:-}" ]]; then
        TEST_ARGS+=("-timeout" "${argc_timeout}")
    else
        case "$test_type" in
            "unit") TEST_ARGS+=("-timeout" "30s") ;;
            "coverage") TEST_ARGS+=("-timeout" "60s") ;;
            "integration") TEST_ARGS+=("-timeout" "120s") ;;
        esac
    fi

    # Short mode
    if is_short_mode_enabled; then
        TEST_ARGS+=("-short")
        is_verbose_enabled && format_test_info "Running in short mode (skipping long tests)"
    fi

    # Coverage specific arguments
    if [[ "$test_type" == "coverage" ]]; then
        local coverage_file="${argc_output:-./tmp/coverage.out}"
        TEST_ARGS+=("-coverprofile=$coverage_file")
        TEST_ARGS+=("-covermode=atomic")
        is_verbose_enabled && format_test_info "Coverage output: $coverage_file"
    fi

    return 0
}

execute_test_command() {
    local project_root="$1"
    local test_type="$2"
    local -a test_args=("${@:3}")

    # Change to project directory
    local original_pwd="$PWD"
    cd "$project_root" || {
        format_test_error "Failed to change to project directory: $project_root"
        return 1
    }

    # Create tmp directory if needed
    [[ ! -d "./tmp" ]] && mkdir -p "./tmp"

    is_verbose_enabled && format_test_info "Running $test_type tests in: $project_root"
    is_verbose_enabled && format_test_info "Test arguments: ${test_args[*]}"

    # Execute tests
    local start_time
    start_time=$(date +%s)

    local exit_code=0
    case "$test_type" in
        "unit"|"coverage")
            go test "${test_args[@]}" ./... || exit_code=$?
            ;;
        "integration")
            # Integration tests with special tags
            go test "${test_args[@]}" -tags=integration ./... || exit_code=$?
            ;;
    esac

    local end_time
    end_time=$(date +%s)
    local duration=$((end_time - start_time))

    # Return to original directory
    cd "$original_pwd" || {
        format_test_warning "Failed to return to original directory"
    }

    # Store duration for reporting
    TEST_DURATION="${duration}s"

    return "$exit_code"
}

run_unit_test_operation() {
    local project_root="${argc_project:-.}"

    # Show banner and info
    show_test_banner "Unit"
    show_test_verbose_info "Unit"

    # Validate environment
    if ! validate_test_environment "$project_root" "unit"; then
        return 1
    fi

    # Prepare arguments
    prepare_test_arguments "unit"

    # Execute tests
    show_test_progress "Starting unit tests..."
    local exit_code=0
    execute_test_command "$project_root" "unit" "${TEST_ARGS[@]}" || exit_code=$?

    # Display results
    display_test_summary "Unit" "$exit_code" "${TEST_DURATION:-unknown}"

    return "$exit_code"
}

run_coverage_test_operation() {
    local project_root="${argc_project:-.}"
    local coverage_file="${argc_output:-./tmp/coverage.out}"

    # Convert relative path to absolute
    [[ "$coverage_file" != /* ]] && coverage_file="$project_root/$coverage_file"

    # Show banner and info
    show_test_banner "Coverage"
    show_test_verbose_info "Coverage"

    # Validate environment
    if ! validate_test_environment "$project_root" "coverage"; then
        return 1
    fi

    # Prepare arguments
    prepare_test_arguments "coverage"

    # Execute tests
    show_test_progress "Starting coverage tests..."
    local exit_code=0
    execute_test_command "$project_root" "coverage" "${TEST_ARGS[@]}" || exit_code=$?

    # Display results
    display_test_summary "Coverage" "$exit_code" "${TEST_DURATION:-unknown}"
    display_coverage_summary "$coverage_file" "$exit_code"

    # Generate HTML report if requested
    if [[ "$exit_code" -eq 0 ]] && is_html_enabled; then
        local html_file="${coverage_file%.out}.html"
        show_test_progress "Generating HTML coverage report..."

        if go tool cover -html="$coverage_file" -o "$html_file"; then
            format_test_success "HTML report generated: $html_file"

            # Open in browser if requested
            if is_open_enabled && command -v open >/dev/null 2>&1; then
                open "$html_file"
                format_test_info "Opened HTML report in browser"
            fi
        else
            format_test_warning "Failed to generate HTML report"
        fi
    fi

    return "$exit_code"
}

run_integration_test_operation() {
    local project_root="${argc_project:-.}"

    # Show banner and info
    show_test_banner "Integration"
    show_test_verbose_info "Integration"

    # Display environment info
    if is_verbose_enabled; then
        echo "Environment Configuration:"
        echo "========================="
        if [[ -n "${WNC_CONTROLLER:-}" ]]; then
            format_test_success "WNC_CONTROLLER: ${WNC_CONTROLLER}"
        fi
        if [[ -n "${WNC_ACCESS_TOKEN:-}" ]]; then
            format_test_success "WNC_ACCESS_TOKEN: ***"
        fi
        echo
    fi

    # Validate environment
    if ! validate_test_environment "$project_root" "integration"; then
        return 1
    fi

    # Prepare arguments
    prepare_test_arguments "integration"

    # Execute tests
    show_test_progress "Starting integration tests..."
    local exit_code=0
    execute_test_command "$project_root" "integration" "${TEST_ARGS[@]}" || exit_code=$?

    # Display results
    display_test_summary "Integration" "$exit_code" "${TEST_DURATION:-unknown}"

    return "$exit_code"
}
