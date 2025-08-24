// Package mobility provides Mobility-specific errors for the Cisco IOS-XE Wireless Network Controller API.
package mobility

import (
	ierrors "github.com/umatare5/cisco-ios-xe-wireless-go/internal/errors"
)

// EntityTypeMobility represents the Mobility entity type for error messages
const EntityTypeMobility = "Mobility"

// Common error message shortcuts for backward compatibility
var (
	// ErrOperationFailed is the generic operation failure message template
	ErrOperationFailed = ierrors.ErrOperationFailedTemplate

	// ErrMobilitySpecificOperationFailed is the Mobility operation failure message template
	ErrMobilitySpecificOperationFailed = "failed to %s Mobility %s: %w"
)
