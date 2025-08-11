#!/usr/bin/env bash

# Cisco WNC Core - Predicate Functions
# Centralized predicate functions for all WNC operations

# Guard against multiple inclusions
if [[ -n "${WNC_PREDICATES_LOADED:-}" ]]; then
    return 0
fi
readonly WNC_PREDICATES_LOADED=1

# ===============================
# Environment and Configuration Predicates
# ===============================

# Check if verbose mode is enabled
is_verbose_enabled() { [[ "${WNC_VERBOSE:-0}" == "1" || "${VERBOSE:-false}" == "true" ]]; }

# Check if no color mode is enabled
is_no_color_enabled() { [[ "${WNC_NO_COLOR:-0}" == "1" ]]; }

# Check if short mode is enabled (skip long-running tests)
is_short_mode_enabled() { [[ "${WNC_SHORT:-0}" == "1" ]]; }

# Check if coverage is enabled
is_coverage_enabled() { [[ "${WNC_COVERAGE:-0}" == "1" ]]; }

# Check if environment check should be skipped
is_skip_env_check_enabled() { [[ "${WNC_CHECK_ENV_ONLY:-0}" == "1" ]]; }

# Check if environment check only mode is enabled
is_check_env_only() { [[ "${WNC_CHECK_ENV_ONLY:-0}" == "1" ]]; }

# Check if fix mode is enabled
is_fix_enabled() { [[ "${WNC_FIX:-0}" == "1" ]]; }

# Check if clean mode is enabled
is_clean_enabled() { [[ "${WNC_CLEAN:-0}" == "1" ]]; }

# Check if update mode is enabled
is_update_enabled() { [[ "${WNC_UPDATE:-0}" == "1" ]]; }

# Check if force mode is enabled
is_force_enabled() { [[ "${WNC_FORCE:-0}" == "1" ]]; }

# Check if download only mode is enabled
is_download_only_enabled() { [[ "${WNC_DOWNLOAD_ONLY:-0}" == "1" ]]; }

# Check if verify mode is enabled
is_verify_enabled() { [[ "${WNC_VERIFY:-0}" == "1" ]]; }

# Check if insecure mode is enabled (skip TLS verification)
is_insecure_enabled() { [[ "${WNC_INSECURE:-0}" == "1" ]]; }

# ===============================
# Help and CLI Predicates
# ===============================

# Check if help is requested
is_help_requested() { [[ "$1" == "-h" || "$1" == "--help" ]]; }

# Check if race detection is enabled
is_race_detection_enabled() { [[ "${RACE_FLAG:-true}" == "true" ]]; }

# Check if HTML coverage report generation is enabled
is_html_enabled() { [[ "${HTML_COVERAGE:-false}" == "true" ]]; }

# Check if browser opening is enabled
is_open_enabled() { [[ "${OPEN_BROWSER:-false}" == "true" ]]; }

# ===============================
# Authentication and Connection Predicates
# ===============================

# Check if authentication token is provided
_has_auth_token() {
    local token="$1"
    [[ -n "$token" ]]
}

# Check if curl command is available
_has_curl_command() {
    command -v curl >/dev/null 2>&1
}

# Check if controller is specified
_has_controller() {
    local controller="$1"
    [[ -n "$controller" ]]
}

# ===============================
# File System and Validation Predicates
# ===============================

# Check if directory is valid and accessible
is_directory_valid() { [[ -d "$1" ]]; }
is_valid_directory() { [[ -d "$1" ]]; }

# Check if file is valid and readable
is_file_valid() { [[ -f "$1" && -r "$1" ]]; }

# Check if file is readable
is_file_readable() {
    local file_path="$1"
    [[ -f "$file_path" && -r "$file_path" ]]
}

# Check if file is writable
is_file_writable() {
    local file_path="$1"
    [[ -f "$file_path" && -w "$file_path" ]]
}

# Check if file exists
file_exists() { [[ -f "$1" ]]; }

# Check if directory exists
dir_exists() { [[ -d "$1" ]]; }

# Check if command exists
command_exists() { command -v "$1" >/dev/null 2>&1; }
is_command_available() { command -v "$1" >/dev/null 2>&1; }

# Check if CLI tool is available
is_cli_available() {
    local cli_name="${1:-}"
    [[ -n "$cli_name" ]] && command -v "$cli_name" >/dev/null 2>&1
}

# Check if value is empty
is_empty() { [[ -z "${1:-}" ]]; }

# Check if value is not empty
is_not_empty() { [[ -n "${1:-}" ]]; }

# Check if value is numeric
is_numeric() { [[ "$1" =~ ^[0-9]+$ ]]; }

# Check if timeout format is valid
is_valid_timeout() { [[ "$1" =~ ^[0-9]+[smh]?$ ]]; }

# Check if package pattern is valid
is_valid_package_pattern() { [[ -n "$1" && "$1" != " " ]]; }

# Check if threshold value is valid (0-100)
is_valid_threshold() { [[ "$1" -ge 0 && "$1" -le 100 ]] 2>/dev/null; }

# Check if URL format is valid
is_valid_url() { [[ "$1" =~ ^https?:// ]]; }

# Check if file extension is valid
is_valid_file_extension() {
    local file="$1"
    local expected_ext="$2"
    [[ "${file##*.}" == "$expected_ext" ]]
}

# ===============================
# Generic Boolean Helpers
# ===============================

# Check if value is enabled (1 or true)
is_enabled() {
    [[ "${1:-0}" == "1" || "${1:-false}" == "true" ]]
}

# Check if value is true
is_true() {
    [[ "${1:-false}" == "true" ]]
}

# ===============================
# Certificate and Security
# ===============================

# Get CA certificate file path
get_cacert_file() { printf '%s\n' "${WNC_CACERT:-}"; }
