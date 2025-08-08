#!/usr/bin/env bash
# Centralized predicate helper functions for WNC scripts
# Loaded by bootstrap.sh if present. Avoid duplication across entry scripts.
set -euo pipefail

# argc flag predicates (flags become 1/0)
is_verbose_enabled() { [[ "${argc_verbose:-0}" == "1" ]]; }
is_no_color_enabled() { [[ "${argc_no_color:-0}" == "1" ]]; }
is_short_mode_enabled() { [[ "${argc_short:-0}" == "1" ]]; }
is_coverage_enabled() { [[ "${argc_coverage:-0}" == "1" ]]; }
is_check_env_only() { [[ "${argc_check_env_only:-0}" == "1" ]]; }
is_fix_enabled() { [[ "${argc_fix:-0}" == "1" ]]; }
is_clean_enabled() { [[ "${argc_clean:-0}" == "1" ]]; }
is_update_enabled() { [[ "${argc_update:-0}" == "1" ]]; }
is_force_enabled() { [[ "${argc_force:-0}" == "1" ]]; }
is_download_only_enabled() { [[ "${argc_download_only:-0}" == "1" ]]; }
is_verify_enabled() { [[ "${argc_verify:-0}" == "1" ]]; }
is_insecure_enabled() { [[ "${argc_insecure:-0}" == "1" ]]; }

# Generic helpers
is_command_available() { command -v "${1:-}" >/dev/null 2>&1; }
file_exists() { [[ -f "${1:-}" ]]; }
dir_exists() { [[ -d "${1:-}" ]]; }

# Backwards compatibility shim (legacy functions some modules may call)
is_true() { [[ "${1:-false}" == "true" || "${1:-0}" == "1" ]]; }
