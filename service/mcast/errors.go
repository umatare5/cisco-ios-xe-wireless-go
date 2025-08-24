// Package mcast provides Mcast-specific errors for the Cisco IOS-XE Wireless Network Controller API.
package mcast

import (
	ierrors "github.com/umatare5/cisco-ios-xe-wireless-go/internal/errors"
)

// EntityTypeMcast represents the Mcast entity type for error messages
const EntityTypeMcast = "Mcast"

// Common error message shortcuts for backward compatibility
var (
	// ErrOperationFailed is the generic operation failure message template
	ErrOperationFailed = ierrors.ErrOperationFailedTemplate

	// ErrMcastSpecificOperationFailed is the Mcast operation failure message template
	ErrMcastSpecificOperationFailed = "failed to %s Mcast %s: %w"
)
