// Package ble provides BLE-specific errors for the Cisco IOS-XE Wireless Network Controller API.
package ble

import (
	ierrors "github.com/umatare5/cisco-ios-xe-wireless-go/internal/errors"
)

// EntityTypeBLE represents the BLE entity type for error messages
const EntityTypeBLE = "BLE"

// Common error message shortcuts for backward compatibility
var (
	// ErrOperationFailed is the generic operation failure message template
	ErrOperationFailed = ierrors.ErrOperationFailedTemplate

	// ErrBLESpecificOperationFailed is the BLE operation failure message template
	ErrBLESpecificOperationFailed = "failed to %s BLE %s: %w"
)
