#!/usr/bin/env bash

# Cisco WNC YANG Statement Details - Predicate Functions
# Provides predicate functions for YANG statement processing

_is_json_output_format() {
    local format="$1"
    [[ "${format}" == "json" ]]
}

_is_jq_available() {
    command -v jq >/dev/null 2>&1
}

_can_process_json_output() {
    command -v jq >/dev/null 2>&1
}
