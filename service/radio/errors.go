// Package radio provides Radio-specific errors for the Cisco IOS-XE Wireless Network Controller API.
package radio

import (
	ierrors "github.com/umatare5/cisco-ios-xe-wireless-go/internal/errors"
)

// EntityTypeRadio represents the Radio entity type for error messages
const EntityTypeRadio = "Radio"

// Common error message shortcuts for backward compatibility
var (
	// ErrOperationFailed is the generic operation failure message template
	ErrOperationFailed = ierrors.ErrOperationFailedTemplate

	// ErrRadioSpecificOperationFailed is the Radio operation failure message template
	ErrRadioSpecificOperationFailed = "failed to %s Radio %s: %w"
)
