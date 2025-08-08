#!/usr/bin/env bash

# Artifacts Cleanup Output Functions

set -euo pipefail

# Color definitions (will be empty if --no-color is used)
readonly ARTIFACTS_COLOR_RED="${COLOR_RED:-}"
readonly ARTIFACTS_COLOR_GREEN="${COLOR_GREEN:-}"
readonly ARTIFACTS_COLOR_YELLOW="${COLOR_YELLOW:-}"
readonly ARTIFACTS_COLOR_BLUE="${COLOR_BLUE:-}"
readonly ARTIFACTS_COLOR_RESET="${COLOR_RESET:-}"

format_artifacts_error() {
    echo "${ARTIFACTS_COLOR_RED}✗ Cleanup Error: $1${ARTIFACTS_COLOR_RESET}" >&2
}

format_artifacts_success() {
    echo "${ARTIFACTS_COLOR_GREEN}✓ Cleanup Success: $1${ARTIFACTS_COLOR_RESET}"
}

format_artifacts_warning() {
    echo "${ARTIFACTS_COLOR_YELLOW}⚠ Cleanup Warning: $1${ARTIFACTS_COLOR_RESET}" >&2
}

format_artifacts_info() {
    echo "${ARTIFACTS_COLOR_BLUE}ℹ Cleanup Info: $1${ARTIFACTS_COLOR_RESET}"
}

format_artifacts_banner() {
    wnc_banner_artifacts
}

format_step_message() {
    local step="$1"
    local message="$2"
    echo "${ARTIFACTS_COLOR_BLUE}[$step]${ARTIFACTS_COLOR_RESET} $message"
}

format_file_removed() {
    local path="$1"
    local size="${2:-}"

    if [[ -n "$size" ]]; then
        echo "  ${ARTIFACTS_COLOR_GREEN}✓${ARTIFACTS_COLOR_RESET} $path ($size)"
    else
        echo "  ${ARTIFACTS_COLOR_GREEN}✓${ARTIFACTS_COLOR_RESET} $path"
    fi
}

format_dry_run_item() {
    local action="$1"
    local path="$2"
    local size="${3:-}"

    if [[ -n "$size" ]]; then
        echo "  ${ARTIFACTS_COLOR_YELLOW}[DRY-RUN]${ARTIFACTS_COLOR_RESET} $action: $path ($size)"
    else
        echo "  ${ARTIFACTS_COLOR_YELLOW}[DRY-RUN]${ARTIFACTS_COLOR_RESET} $action: $path"
    fi
}

format_cache_info() {
    local cache_type="$1"
    local cache_path="$2"
    local exists="$3"

    if [[ "$exists" == "true" ]]; then
        echo "  ${ARTIFACTS_COLOR_GREEN}✓${ARTIFACTS_COLOR_RESET} $cache_type: $cache_path"
    else
        echo "  ${ARTIFACTS_COLOR_YELLOW}!${ARTIFACTS_COLOR_RESET} $cache_type: $cache_path (not found)"
    fi
}
