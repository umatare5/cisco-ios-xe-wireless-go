// Package rrm provides RRM-specific errors for the Cisco IOS-XE Wireless Network Controller API.
package rrm

import (
	ierrors "github.com/umatare5/cisco-ios-xe-wireless-go/internal/errors"
)

// EntityTypeRRM represents the RRM entity type for error messages
const EntityTypeRRM = "RRM"

// Common error message shortcuts for backward compatibility
var (
	// ErrOperationFailed is the generic operation failure message template
	ErrOperationFailed = ierrors.ErrOperationFailedTemplate

	// ErrRRMSpecificOperationFailed is the RRM operation failure message template
	ErrRRMSpecificOperationFailed = "failed to %s RRM %s: %w"
)
