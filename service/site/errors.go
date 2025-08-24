// Package site provides Site-specific errors for the Cisco IOS-XE Wireless Network Controller API.
package site

import (
	ierrors "github.com/umatare5/cisco-ios-xe-wireless-go/internal/errors"
)

// EntityTypeSite represents the Site entity type for error messages
const EntityTypeSite = "Site"

// Common error message shortcuts for backward compatibility
var (
	// ErrOperationFailed is the generic operation failure message template
	ErrOperationFailed = ierrors.ErrOperationFailedTemplate

	// ErrSiteSpecificOperationFailed is the Site operation failure message template
	ErrSiteSpecificOperationFailed = "failed to %s Site %s: %w"
)
