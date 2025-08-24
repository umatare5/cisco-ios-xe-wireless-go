// Package geolocation provides Geolocation-specific errors for the Cisco IOS-XE Wireless Network Controller API.
package geolocation

import (
	ierrors "github.com/umatare5/cisco-ios-xe-wireless-go/internal/errors"
)

// EntityTypeGeolocation represents the Geolocation entity type for error messages
const EntityTypeGeolocation = "Geolocation"

// Common error message shortcuts for backward compatibility
var (
	// ErrOperationFailed is the generic operation failure message template
	ErrOperationFailed = ierrors.ErrOperationFailedTemplate

	// ErrGeolocationSpecificOperationFailed is the Geolocation operation failure message template
	ErrGeolocationSpecificOperationFailed = "failed to %s Geolocation %s: %w"
)
