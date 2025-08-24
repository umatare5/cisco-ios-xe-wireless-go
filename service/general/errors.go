// Package general provides General-specific errors for the Cisco IOS-XE Wireless Network Controller API.
package general

import (
	ierrors "github.com/umatare5/cisco-ios-xe-wireless-go/internal/errors"
)

// EntityTypeGeneral represents the General entity type for error messages
const EntityTypeGeneral = "General"

// Common error message shortcuts for backward compatibility
var (
	// ErrOperationFailed is the generic operation failure message template
	ErrOperationFailed = ierrors.ErrOperationFailedTemplate

	// ErrGeneralSpecificOperationFailed is the General operation failure message template
	ErrGeneralSpecificOperationFailed = "failed to %s General %s: %w"
)
