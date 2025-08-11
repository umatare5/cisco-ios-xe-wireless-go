#!/usr/bin/env bash

# Cisco WNC Help - Banner Functions
# Provides banner display functionality for help system

show_help_banner() {
    if command -v wnc_banner_help >/dev/null 2>&1; then
        wnc_banner_help
        return
    fi

    printf '%s\n' "Cisco WNC Development Scripts"
    printf '%s\n' "-------------------------------"
}
