// Package apf provides APF-specific errors for the Cisco IOS-XE Wireless Network Controller API.
package apf

import (
	ierrors "github.com/umatare5/cisco-ios-xe-wireless-go/internal/errors"
)

// EntityTypeAPF represents the APF entity type for error messages
const EntityTypeAPF = "APF"

// Common error message shortcuts for backward compatibility
var (
	// ErrOperationFailed is the generic operation failure message template
	ErrOperationFailed = ierrors.ErrOperationFailedTemplate

	// ErrAPFSpecificOperationFailed is the APF operation failure message template
	ErrAPFSpecificOperationFailed = "failed to %s APF %s: %w"
)
