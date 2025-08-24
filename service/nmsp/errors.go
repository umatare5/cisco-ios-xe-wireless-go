// Package nmsp provides NMSP-specific errors for the Cisco IOS-XE Wireless Network Controller API.
package nmsp

import (
	ierrors "github.com/umatare5/cisco-ios-xe-wireless-go/internal/errors"
)

// EntityTypeNMSP represents the NMSP entity type for error messages
const EntityTypeNMSP = "NMSP"

// Common error message shortcuts for backward compatibility
var (
	// ErrOperationFailed is the generic operation failure message template
	ErrOperationFailed = ierrors.ErrOperationFailedTemplate

	// ErrNMSPSpecificOperationFailed is the NMSP operation failure message template
	ErrNMSPSpecificOperationFailed = "failed to %s NMSP %s: %w"
)
