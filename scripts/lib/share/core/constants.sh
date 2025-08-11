#!/usr/bin/env bash

# Cisco WNC Core - Constants and Default Values
# Provides constants and default values for WNC operations

# Guard against multiple inclusions
if [[ -n "${WNC_CONSTANTS_LOADED:-}" ]]; then
    return 0
fi
readonly WNC_CONSTANTS_LOADED=1

# ===============================
# Protocol and Network Constants
# ===============================

readonly WNC_DEFAULT_PROTOCOL="https"
readonly WNC_DEFAULT_CONTROLLER=""
readonly WNC_DEFAULT_OUTPUT_FORMAT="pretty"
readonly WNC_DEFAULT_TIMEOUT="30s"

# ===============================
# RESTCONF API Paths
# ===============================

readonly WNC_RESTCONF_DATA_PATH="/restconf/data"
readonly WNC_RESTCONF_MODULES_PATH="/restconf/tailf/modules"

# ===============================
# HTTP Headers
# ===============================

readonly WNC_YANG_ACCEPT_HEADER="application/yang-data+json"

# ===============================
# Exit Codes
# ===============================

readonly WNC_EXIT_SUCCESS=0
readonly WNC_EXIT_ERROR=1
readonly WNC_EXIT_INVALID_ARGS=2
readonly WNC_EXIT_AUTH_ERROR=3
readonly WNC_EXIT_NETWORK_ERROR=4

# ===============================
# Default Messages
# ===============================

readonly WNC_NETWORK_ERROR_MSG="Failed to connect to the controller"

# ===============================
# Getter Functions (for backward compatibility)
# ===============================

get_default_protocol() { printf '%s\n' "$WNC_DEFAULT_PROTOCOL"; }
get_default_controller() { printf '%s\n' "$WNC_DEFAULT_CONTROLLER"; }
get_default_output_format() { printf '%s\n' "$WNC_DEFAULT_OUTPUT_FORMAT"; }
get_default_timeout() { printf '%s\n' "$WNC_DEFAULT_TIMEOUT"; }
get_restconf_data_path() { printf '%s\n' "$WNC_RESTCONF_DATA_PATH"; }
get_restconf_modules_path() { printf '%s\n' "$WNC_RESTCONF_MODULES_PATH"; }
get_yang_accept_header() { printf '%s\n' "$WNC_YANG_ACCEPT_HEADER"; }
get_network_error_message() { printf '%s\n' "$WNC_NETWORK_ERROR_MSG"; }
get_exit_success() { printf '%s\n' "$WNC_EXIT_SUCCESS"; }
get_exit_error() { printf '%s\n' "$WNC_EXIT_ERROR"; }
get_exit_invalid_args() { printf '%s\n' "$WNC_EXIT_INVALID_ARGS"; }
get_exit_auth_error() { printf '%s\n' "$WNC_EXIT_AUTH_ERROR"; }
get_exit_network_error() { printf '%s\n' "$WNC_EXIT_NETWORK_ERROR"; }
