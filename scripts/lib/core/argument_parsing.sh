#!/usr/bin/env bash

# Cisco WNC Argument Parsing Library
# Common predicate functions and utilities for argument parsing

set -euo pipefail

# Prevent double sourcing
if [[ -n "${WNC_ARGUMENT_PARSING_LOADED:-}" ]]; then
    return 0
fi
readonly WNC_ARGUMENT_PARSING_LOADED=1

set -euo pipefail

# ===============================
# Basic Validation Predicates
# ===============================

# Check if help is requested
is_help_requested() { [[ "$1" == "-h" || "$1" == "--help" ]]; }

# Check if verbose mode is enabled
is_verbose_enabled() { [[ "${VERBOSE:-false}" == "true" ]]; }

# Check if race detection is enabled
is_race_detection_enabled() { [[ "${RACE_FLAG:-true}" == "true" ]]; }

# Check if HTML coverage report generation is enabled
is_html_enabled() { [[ "${HTML_COVERAGE:-false}" == "true" ]]; }

# Check if browser opening is enabled
is_open_enabled() { [[ "${OPEN_BROWSER:-false}" == "true" ]]; }

# Check if directory exists and is valid
is_directory_valid() { [[ -d "$1" ]]; }

# Check if file exists and is readable
is_file_valid() { [[ -f "$1" && -r "$1" ]]; }

# Check if string is numeric
is_numeric() { [[ "$1" =~ ^[0-9]+$ ]]; }

# ===============================
# Format Validation Predicates
# ===============================

# Check if timeout format is valid (e.g., 10m, 30s, 1h)
is_valid_timeout() { [[ "$1" =~ ^[0-9]+[smh]?$ ]]; }

# Check if package pattern is valid (non-empty, not just spaces)
is_valid_package_pattern() { [[ -n "$1" && "$1" != " " ]]; }

# Check if threshold is a valid percentage (0-100)
is_valid_threshold() { [[ "$1" -ge 0 && "$1" -le 100 ]] 2>/dev/null; }

# Check if URL format is valid (basic check)
is_valid_url() { [[ "$1" =~ ^https?:// ]]; }

# Check if file extension matches expected
is_valid_file_extension() {
    local file="$1" expected_ext="$2"
    [[ "${file##*.}" == "$expected_ext" ]]
}

# ===============================
# Argument Processing Utilities
# ===============================

# Ensure required argument is provided
requires_argument() {
    [[ -n "${2:-}" ]] || {
        show_error "$1 requires an argument"
        exit 2
    }
}

# Show unknown option error
show_unknown_option_error() {
    show_error "Unknown option '$1'"
    echo "Use --help for usage" >&2
    exit 2
}

# ===============================
# Common Argument Validators
# ===============================

# Validate project directory argument
validate_project_directory() {
    local project_root="$1"
    is_directory_valid "$project_root" || {
        show_error "Directory not found: $project_root"
        exit 1
    }
}

# Validate timeout argument
validate_timeout_argument() {
    local timeout="$1"
    is_valid_timeout "$timeout" || {
        show_error "Invalid timeout format: $timeout" \
             "(use format like 10m, 30s, 1h)" >&2
        exit 2
    }
}

# Validate package pattern argument
validate_package_pattern_argument() {
    local package_pattern="$1"
    is_valid_package_pattern "$package_pattern" || {
        show_error "Invalid package pattern: $package_pattern"
        exit 2
    }
}

# Validate threshold argument
validate_threshold_argument() {
    local threshold="$1"
    is_numeric "$threshold" || {
        show_error "Threshold must be numeric: $threshold"
        exit 2
    }
    is_valid_threshold "$threshold" || {
        show_error "Threshold must be between 0 and 100: $threshold"
        exit 2
    }
}

# ===============================
# Argument Parsing Templates
# ===============================

# Standard help option handler
handle_help_option() {
    show_help && exit 0
}

# Standard project option handler
handle_project_option() {
    requires_argument "--project" "${2:-}"
    echo "$2"  # Return the project path
}

# Standard verbose option handler
handle_verbose_option() {
    echo "true"  # Return verbose flag
}

# Standard timeout option handler
handle_timeout_option() {
    requires_argument "--timeout" "${2:-}"
    validate_timeout_argument "$2"
    echo "$2"  # Return the timeout value
}

# Standard package option handler
handle_package_option() {
    requires_argument "--package" "${2:-}"
    validate_package_pattern_argument "$2"
    echo "$2"  # Return the package pattern
}

# Standard threshold option handler
handle_threshold_option() {
    requires_argument "--threshold" "${2:-}"
    validate_threshold_argument "$2"
    echo "$2"  # Return the threshold value
}
