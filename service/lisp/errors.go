// Package lisp provides LISP-specific errors for the Cisco IOS-XE Wireless Network Controller API.
package lisp

import (
	ierrors "github.com/umatare5/cisco-ios-xe-wireless-go/internal/errors"
)

// EntityTypeLISP represents the LISP entity type for error messages
const EntityTypeLISP = "LISP"

// Common error message shortcuts for backward compatibility
var (
	// ErrOperationFailed is the generic operation failure message template
	ErrOperationFailed = ierrors.ErrOperationFailedTemplate

	// ErrLISPSpecificOperationFailed is the LISP operation failure message template
	ErrLISPSpecificOperationFailed = "failed to %s LISP %s: %w"
)
