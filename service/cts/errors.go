// Package cts provides CTS-specific errors for the Cisco IOS-XE Wireless Network Controller API.
package cts

import (
	ierrors "github.com/umatare5/cisco-ios-xe-wireless-go/internal/errors"
)

// EntityTypeCTS represents the CTS entity type for error messages
const EntityTypeCTS = "CTS"

// Common error message shortcuts for backward compatibility
var (
	// ErrOperationFailed is the generic operation failure message template
	ErrOperationFailed = ierrors.ErrOperationFailedTemplate

	// ErrCTSSpecificOperationFailed is the CTS operation failure message template
	ErrCTSSpecificOperationFailed = "failed to %s CTS %s: %w"
)
