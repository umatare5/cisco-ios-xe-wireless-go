#!/usr/bin/env bash

# Unified Banner Helper
# Provides a single function to render consistent banners across scripts.

# Compute banner width based on longest line (capped to 60 for consistency)
_wnc_banner_width() {
    local max=0
    for line in "$@"; do
        local len=${#line}
        (( len > max )) && max=$len
    done
    # Add padding spaces (2 on each side inside border)
    (( max < 38 )) && max=38  # minimum width for aesthetic baseline
    (( max > 56 )) && max=56  # cap to avoid overly wide boxes
    printf '%s\n' "$max"
}

# Pad a single line centered within target width
_wnc_banner_center_line() {
    local text="$1" width="$2"
    local len=${#text}
    local space=$(( width - len ))
    local left=$(( space / 2 ))
    local right=$(( space - left ))
    printf '%*s%s%*s' "$left" '' "$text" "$right" ''
}

# Public banner function
wnc_banner() {
    local primary="$1"
    local secondary="${2:-}" # optional second line

    local lines=("$primary")
    if [[ -n "$secondary" ]]; then
        lines+=("$secondary")
    fi

    local width
    width=$(_wnc_banner_width "${lines[@]}")

    local border
    border=$(printf '%*s' "$width" '' | tr ' ' '=')

    # Handle no-color mode with plain text output
    if is_no_color_enabled; then
        printf '%s\n' "$border"
        for l in "${lines[@]}"; do
            _wnc_banner_center_line "$l" "$width"
            printf '\n'
        done
        printf '%s\n' "$border"
        printf '\n'
        return
    fi

    # Render colored banner with unified blue theme
    local color="$COLOR_BLUE"
    local reset="$COLOR_RESET"
    printf '%b\n' "${color}${border}${reset}"
    for l in "${lines[@]}"; do
        printf '%b' "${color}"
        _wnc_banner_center_line "$l" "$width"
        printf '%b\n' "${reset}"
    done
    printf '%b\n' "${color}${border}${reset}"
    printf '\n'
}

# Convenience specialized wrappers (can be extended later without changing callers)
wnc_banner_lint() { wnc_banner "Cisco WNC Code Linter" "golangci-lint Integration"; }
wnc_banner_tests() { local kind="$1"; wnc_banner "Cisco WNC ${kind} Tests" "Go Testing Framework"; }
wnc_banner_coverage() { wnc_banner "Coverage HTML Generator" "Go Tool Cover Integration"; }
wnc_banner_yang() { wnc_banner "Cisco WNC YANG Operations" "RESTCONF API Integration"; }
wnc_banner_dependencies() { wnc_banner "Cisco WNC Dependencies" "Module Management"; }
wnc_banner_artifacts() { wnc_banner "Cisco WNC Artifacts" "Cleanup Utility"; }
wnc_banner_pre_commit() { wnc_banner "Pre-commit Validation" "Branch Protection"; }

export -f wnc_banner \
  wnc_banner_lint \
  wnc_banner_tests \
  wnc_banner_coverage \
  wnc_banner_yang \
  wnc_banner_dependencies \
  wnc_banner_artifacts \
  wnc_banner_pre_commit
