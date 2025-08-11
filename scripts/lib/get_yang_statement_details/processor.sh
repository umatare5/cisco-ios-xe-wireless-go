#!/usr/bin/env bash

# Cisco WNC YANG Statement Details - Processing Functions
# Provides processing functions for YANG statement output handling

post_process_statement_output() {
    local file_path="$1"
    local search_key="$2"
    local output_format="$3"

    if _can_process_json_output "$output_format"; then
        jq --arg search_key "$search_key" \
           '..|objects|to_entries[]|select(.key==$search_key)|{($search_key):.value}' \
           "$file_path" || cat "$file_path"
        return
    fi

    grep -A 20 -B 5 "$search_key" "$file_path" || cat "$file_path"
}
