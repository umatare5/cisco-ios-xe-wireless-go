// Package wlan provides WLAN-specific errors for the Cisco IOS-XE Wireless Network Controller API.
package wlan

import (
	ierrors "github.com/umatare5/cisco-ios-xe-wireless-go/internal/errors"
)

// EntityTypeWLAN represents the WLAN entity type for error messages
const EntityTypeWLAN = "WLAN"

// Common error message shortcuts for backward compatibility
var (
	// ErrOperationFailed is the generic operation failure message template
	ErrOperationFailed = ierrors.ErrOperationFailedTemplate

	// ErrWLANSpecificOperationFailed is the WLAN operation failure message template
	ErrWLANSpecificOperationFailed = "failed to %s WLAN %s: %w"
)
