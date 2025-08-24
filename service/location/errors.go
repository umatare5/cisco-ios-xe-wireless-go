// Package location provides Location-specific errors for the Cisco IOS-XE Wireless Network Controller API.
package location

import (
	ierrors "github.com/umatare5/cisco-ios-xe-wireless-go/internal/errors"
)

// EntityTypeLocation represents the Location entity type for error messages
const EntityTypeLocation = "Location"

// Common error message shortcuts for backward compatibility
var (
	// ErrOperationFailed is the generic operation failure message template
	ErrOperationFailed = ierrors.ErrOperationFailedTemplate

	// ErrLocationSpecificOperationFailed is the Location operation failure message template
	ErrLocationSpecificOperationFailed = "failed to %s Location %s: %w"
)
