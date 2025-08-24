// Package dot15 provides Dot15-specific errors for the Cisco IOS-XE Wireless Network Controller API.
package dot15

import (
	ierrors "github.com/umatare5/cisco-ios-xe-wireless-go/internal/errors"
)

// EntityTypeDot15 represents the Dot15 entity type for error messages
const EntityTypeDot15 = "Dot15"

// Common error message shortcuts for backward compatibility
var (
	// ErrOperationFailed is the generic operation failure message template
	ErrOperationFailed = ierrors.ErrOperationFailedTemplate

	// ErrDot15SpecificOperationFailed is the Dot15 operation failure message template
	ErrDot15SpecificOperationFailed = "failed to %s Dot15 %s: %w"
)
