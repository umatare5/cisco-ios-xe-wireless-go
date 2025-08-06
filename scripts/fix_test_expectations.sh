#!/bin/bash

# Script to update test expectations to match the new endpoint format
# Updates test files to expect paths without "/restconf/data/" prefix

echo "Updating test expectations for endpoint paths..."

# Find and fix test files with hardcoded "/restconf/data/" expectations
find . -name "*_test.go" -type f -exec grep -l "/restconf/data/" {} \; | while read -r file; do
    echo "Updating $file..."

    # Update hardcoded endpoint expectations in test maps and variables
    sed -i '' 's|"/restconf/data/Cisco-IOS-XE-wireless-\([^"]*\)"|"Cisco-IOS-XE-wireless-\1"|g' "$file"

    # Update hardcoded paths in individual test strings
    sed -i '' 's|"/restconf/data/ietf-yang-library:yang-library"|"ietf-yang-library:yang-library"|g' "$file"

    # Update strings.HasPrefix checks
    sed -i '' 's|strings\.HasPrefix([^,]*, "/restconf/data/")|!strings.HasPrefix(&1, "/restconf/data/")|g' "$file"

    # Update expectation comments that mention paths
    sed -i '' 's|should start with '"'"'/restconf/data/'"'"'|should not start with '"'"'/restconf/data/'"'"'|g' "$file"
done

echo "Test expectation updates completed!"
