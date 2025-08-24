// Package afc provides AFC-specific errors for the Cisco IOS-XE Wireless Network Controller API.
package afc

import (
	ierrors "github.com/umatare5/cisco-ios-xe-wireless-go/internal/errors"
)

// EntityTypeAFC represents the AFC entity type for error messages
const EntityTypeAFC = "AFC"

// Common error message shortcuts for backward compatibility
var (
	// ErrOperationFailed is the generic operation failure message template
	ErrOperationFailed = ierrors.ErrOperationFailedTemplate

	// ErrAFCSpecificOperationFailed is the AFC operation failure message template
	ErrAFCSpecificOperationFailed = "failed to %s AFC data for %s: %w"
)
