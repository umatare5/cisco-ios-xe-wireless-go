#!/usr/bin/env bash

# Cisco WNC Core - Argument Parser Functions
# Provides argument parsing and validation utilities

# Guard against multiple inclusions
if [[ -n "${WNC_ARGUMENT_PARSER_LOADED:-}" ]]; then
    return 0
fi
readonly WNC_ARGUMENT_PARSER_LOADED=1

# ===============================
# Argument Processing Utilities
# ===============================

# Ensure required argument is provided
requires_argument() {
    [[ -n "${2:-}" ]] || {
        error "$1 requires an argument"
        exit 2
    }
}

# Show unknown option error
show_unknown_option_error() {
    error "Unknown option '$1'"
    printf '%s\n' "Use --help for usage" >&2
    exit 2
}

# ===============================
# Common Argument Validators
# ===============================

# Validate project directory argument
validate_project_directory() {
    local project_root="$1"
    is_directory_valid "$project_root" || {
        error "Directory not found: $project_root"
        exit 1
    }
}

# Validate timeout argument
validate_timeout_argument() {
    local timeout="$1"
    is_valid_timeout "$timeout" || {
        error "Invalid timeout format: $timeout (use format like 10m, 30s, 1h)" >&2
        exit 2
    }
}

# Validate package pattern argument
validate_package_pattern_argument() {
    local package_pattern="$1"
    is_valid_package_pattern "$package_pattern" || {
        error "Invalid package pattern: $package_pattern"
        exit 2
    }
}

# Validate threshold argument
validate_threshold_argument() {
    local threshold="$1"
    is_numeric "$threshold" || {
        error "Threshold must be numeric: $threshold"
        exit 2
    }
    is_valid_threshold "$threshold" || {
        error "Threshold must be between 0 and 100: $threshold"
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
    local project_arg="$1"
    requires_argument "--project" "$project_arg"
    validate_project_directory "$project_arg"
    PROJECT_ROOT="$project_arg"
    export PROJECT_ROOT
}

# Standard verbose option handler
handle_verbose_option() {
    WNC_VERBOSE=1
    export WNC_VERBOSE
}

# Standard timeout option handler
handle_timeout_option() {
    local timeout_arg="$1"
    requires_argument "--timeout" "$timeout_arg"
    validate_timeout_argument "$timeout_arg"
    TEST_TIMEOUT="$timeout_arg"
    export TEST_TIMEOUT
}

# Standard coverage option handler
handle_coverage_option() {
    WNC_COVERAGE=1
    export WNC_COVERAGE
}

# Standard output option handler
handle_output_option() {
    local output_arg="$1"
    requires_argument "--output" "$output_arg"
    OUTPUT_FILE="$output_arg"
    export OUTPUT_FILE
}
