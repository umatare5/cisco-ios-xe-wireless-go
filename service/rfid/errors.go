// Package rfid provides RFID-specific errors for the Cisco IOS-XE Wireless Network Controller API.
package rfid

import (
	ierrors "github.com/umatare5/cisco-ios-xe-wireless-go/internal/errors"
)

// EntityTypeRFID represents the RFID entity type for error messages
const EntityTypeRFID = "RFID"

// Common error message shortcuts for backward compatibility
var (
	// ErrOperationFailed is the generic operation failure message template
	ErrOperationFailed = ierrors.ErrOperationFailedTemplate

	// ErrRFIDSpecificOperationFailed is the RFID operation failure message template
	ErrRFIDSpecificOperationFailed = "failed to %s RFID %s: %w"
)
