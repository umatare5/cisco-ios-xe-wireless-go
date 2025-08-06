#!/bin/bash

# Script to fix endpoint path duplication issue across all packages
# Removes "/restconf/data/" prefix from BasePath constants

echo "Fixing endpoint path duplication issue..."

# Define packages that need fixing
packages=(
    "ap/cfg.go"
    "ap/global_oper.go"
    "awips/oper.go"
    "ble/oper.go"
    "dot11/cfg.go"
    "hyperlocation/oper.go"
    "location/cfg.go"
    "mcast/oper.go"
    "mdns/oper.go"
    "mobility/oper.go"
    "nmsp/oper.go"
    "radio/cfg.go"
    "rf/cfg.go"
    "rfid/cfg.go"
    "rogue/oper.go"
    "rrm/cfg.go"
    "rrm/oper.go"
    "rrm/emul_oper.go"
    "rrm/global_oper.go"
    "site/cfg.go"
    "wlan/cfg.go"
    "wlan/global_oper.go"
)

# Fix each file
for file in "${packages[@]}"; do
    if [ -f "$file" ]; then
        echo "Fixing $file..."
        # Use sed to replace /restconf/data/ prefix in BasePath constants
        sed -i '' 's|BasePath = "/restconf/data/\(Cisco-IOS-XE-[^"]*\)"|BasePath = "\1"|g' "$file"
    else
        echo "Warning: $file not found"
    fi
done

echo "Endpoint path fixes completed!"
