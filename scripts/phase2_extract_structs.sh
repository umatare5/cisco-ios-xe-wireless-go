#!/bin/bash

# Script to extract struct definitions from domain packages to internal/model

# Create arrays of packages and their corresponding files
declare -a OPER_PACKAGES=("afc" "ap" "awips" "ble" "client" "general" "geolocation" "hyperlocation" "lisp" "mcast" "mdns" "mesh" "mobility" "nmsp" "rrm" "rogue")
declare -a CFG_PACKAGES=("ap" "apf" "cts" "dot11" "dot15" "fabric" "flex" "general" "location" "mesh" "radio" "rf" "rfid" "rrm" "site" "wlan")

echo "Phase 2: Moving AI-generated structs to internal/model"

# Process oper files
for pkg in "${OPER_PACKAGES[@]}"; do
    if [ -f "${pkg}/oper.go" ]; then
        echo "Processing ${pkg}/oper.go..."
        # We'll move these manually for proper extraction
    fi
done

# Process cfg files
for pkg in "${CFG_PACKAGES[@]}"; do
    if [ -f "${pkg}/cfg.go" ]; then
        echo "Processing ${pkg}/cfg.go..."
        # We'll move these manually for proper extraction
    fi
done

echo "Phase 2 processing complete"
