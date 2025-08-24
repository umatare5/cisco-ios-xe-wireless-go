// Package flex provides Flex-specific errors for the Cisco IOS-XE Wireless Network Controller API.
package flex

import (
	ierrors "github.com/umatare5/cisco-ios-xe-wireless-go/internal/errors"
)

// EntityTypeFlex represents the Flex entity type for error messages
const EntityTypeFlex = "Flex"

// Common error message shortcuts for backward compatibility
var (
	// ErrOperationFailed is the generic operation failure message template
	ErrOperationFailed = ierrors.ErrOperationFailedTemplate

	// ErrFlexSpecificOperationFailed is the Flex operation failure message template
	ErrFlexSpecificOperationFailed = "failed to %s Flex %s: %w"
)
