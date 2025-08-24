// Package rf provides RF-specific errors for the Cisco IOS-XE Wireless Network Controller API.
package rf

import (
	ierrors "github.com/umatare5/cisco-ios-xe-wireless-go/internal/errors"
)

// EntityTypeRF represents the RF entity type for error messages
const EntityTypeRF = "RF"

// Common error message shortcuts for backward compatibility
var (
	// ErrOperationFailed is the generic operation failure message template
	ErrOperationFailed = ierrors.ErrOperationFailedTemplate

	// ErrRFSpecificOperationFailed is the RF operation failure message template
	ErrRFSpecificOperationFailed = "failed to %s RF %s: %w"
)
