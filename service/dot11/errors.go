// Package dot11 provides Dot11-specific errors for the Cisco IOS-XE Wireless Network Controller API.
package dot11

import (
	ierrors "github.com/umatare5/cisco-ios-xe-wireless-go/internal/errors"
)

// EntityTypeDot11 represents the Dot11 entity type for error messages
const EntityTypeDot11 = "Dot11"

// Common error message shortcuts for backward compatibility
var (
	// ErrOperationFailed is the generic operation failure message template
	ErrOperationFailed = ierrors.ErrOperationFailedTemplate

	// ErrDot11SpecificOperationFailed is the Dot11 operation failure message template
	ErrDot11SpecificOperationFailed = "failed to %s Dot11 %s: %w"
)
