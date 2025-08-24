// Package mdns provides MDNS-specific errors for the Cisco IOS-XE Wireless Network Controller API.
package mdns

import (
	ierrors "github.com/umatare5/cisco-ios-xe-wireless-go/internal/errors"
)

// EntityTypeMDNS represents the MDNS entity type for error messages
const EntityTypeMDNS = "MDNS"

// Common error message shortcuts for backward compatibility
var (
	// ErrOperationFailed is the generic operation failure message template
	ErrOperationFailed = ierrors.ErrOperationFailedTemplate

	// ErrMDNSSpecificOperationFailed is the MDNS operation failure message template
	ErrMDNSSpecificOperationFailed = "failed to %s MDNS %s: %w"
)
