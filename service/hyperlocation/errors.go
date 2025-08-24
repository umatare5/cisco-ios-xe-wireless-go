// Package hyperlocation provides Hyperlocation-specific errors for the Cisco IOS-XE Wireless Network Controller API.
package hyperlocation

import (
	ierrors "github.com/umatare5/cisco-ios-xe-wireless-go/internal/errors"
)

// EntityTypeHyperlocation represents the Hyperlocation entity type for error messages
const EntityTypeHyperlocation = "Hyperlocation"

// Common error message shortcuts for backward compatibility
var (
	// ErrOperationFailed is the generic operation failure message template
	ErrOperationFailed = ierrors.ErrOperationFailedTemplate

	// ErrHyperlocationSpecificOperationFailed is the Hyperlocation operation failure message template
	ErrHyperlocationSpecificOperationFailed = "failed to %s Hyperlocation %s: %w"
)
