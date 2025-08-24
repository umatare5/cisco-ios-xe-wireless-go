// Package client provides Client-specific errors for the Cisco IOS-XE Wireless Network Controller API.
package client

import (
	ierrors "github.com/umatare5/cisco-ios-xe-wireless-go/internal/errors"
)

// EntityTypeClient represents the Client entity type for error messages
const EntityTypeClient = ierrors.EntityTypeClient

// Common error message shortcuts for backward compatibility
var (
	// ErrOperationFailed is the generic operation failure message template
	ErrOperationFailed = ierrors.ErrOperationFailedTemplate

	// ErrClientSpecificOperationFailed is the Client operation failure message template
	ErrClientSpecificOperationFailed = "failed to %s Client %s: %w"
)
