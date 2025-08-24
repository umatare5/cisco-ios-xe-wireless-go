// Package awips provides AWIPS-specific errors for the Cisco IOS-XE Wireless Network Controller API.
package awips

import (
	ierrors "github.com/umatare5/cisco-ios-xe-wireless-go/internal/errors"
)

// EntityTypeAWIPS represents the AWIPS entity type for error messages
const EntityTypeAWIPS = "AWIPS"

// Common error message shortcuts for backward compatibility
var (
	// ErrOperationFailed is the generic operation failure message template
	ErrOperationFailed = ierrors.ErrOperationFailedTemplate

	// ErrAWIPSSpecificOperationFailed is the AWIPS operation failure message template
	ErrAWIPSSpecificOperationFailed = "failed to %s AWIPS %s: %w"
)
