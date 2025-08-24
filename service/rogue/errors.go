// Package rogue provides Rogue-specific errors for the Cisco IOS-XE Wireless Network Controller API.
package rogue

import (
	ierrors "github.com/umatare5/cisco-ios-xe-wireless-go/internal/errors"
)

// EntityTypeRogue represents the Rogue entity type for error messages
const EntityTypeRogue = "Rogue"

// Common error message shortcuts for backward compatibility
var (
	// ErrOperationFailed is the generic operation failure message template
	ErrOperationFailed = ierrors.ErrOperationFailedTemplate

	// ErrRogueSpecificOperationFailed is the Rogue operation failure message template
	ErrRogueSpecificOperationFailed = "failed to %s Rogue %s: %w"
)
