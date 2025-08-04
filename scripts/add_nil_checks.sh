#!/bin/bash
# Script to add nil client checks to all packages

# Function to add nil client check to a package
add_nil_client_check() {
    local package=$1
    echo "Processing package: $package"

    # Add errors import and nil check to cfg.go if it exists
    if [ -f "$package/cfg.go" ]; then
        echo "  - Processing cfg.go"

        # Check if errors import already exists
        if ! grep -q '"errors"' "$package/cfg.go"; then
            # Add errors import
            sed -i '' '/^import (/a\
	"errors"
' "$package/cfg.go"
        fi

        # Add nil client checks to functions
        # This is a simplified approach - for each function that takes *wnc.Client as first parameter
        sed -i '' '/func.*(\*wnc\.Client.*{/{
n
/var data/i\
	if client == nil {\
		return nil, errors.New("client is nil")\
	}
}' "$package/cfg.go"
    fi

    # Similar for oper.go
    if [ -f "$package/oper.go" ]; then
        echo "  - Processing oper.go"

        if ! grep -q '"errors"' "$package/oper.go"; then
            sed -i '' '/^import (/a\
	"errors"
' "$package/oper.go"
        fi

        sed -i '' '/func.*(\*wnc\.Client.*{/{
n
/var data/i\
	if client == nil {\
		return nil, errors.New("client is nil")\
	}
}' "$package/oper.go"
    fi

    # Similar for other .go files
    for gofile in "$package"/*.go; do
        if [[ "$gofile" == *"_test.go" ]]; then
            continue  # Skip test files
        fi

        if [[ "$gofile" != "$package/cfg.go" ]] && [[ "$gofile" != "$package/oper.go" ]]; then
            if [ -f "$gofile" ]; then
                echo "  - Processing $gofile"

                if ! grep -q '"errors"' "$gofile"; then
                    sed -i '' '/^import (/a\
	"errors"
' "$gofile"
                fi

                sed -i '' '/func.*(\*wnc\.Client.*{/{
n
/var data/i\
	if client == nil {\
		return nil, errors.New("client is nil")\
	}
}' "$gofile"
            fi
        fi
    done
}

# List of packages to process
packages=(
    "dot11"
    "dot15"
    "fabric"
    "flex"
    "location"
    "radio"
    "rf"
    "rfid"
    "site"
)

for pkg in "${packages[@]}"; do
    add_nil_client_check "$pkg"
done

echo "Nil client checks added to all packages"
