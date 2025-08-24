// Package fabric provides Fabric-specific errors for the Cisco IOS-XE Wireless Network Controller API.
package fabric

import (
	ierrors "github.com/umatare5/cisco-ios-xe-wireless-go/internal/errors"
)

// EntityTypeFabric represents the Fabric entity type for error messages
const EntityTypeFabric = "Fabric"

// Common error message shortcuts for backward compatibility
var (
	// ErrOperationFailed is the generic operation failure message template
	ErrOperationFailed = ierrors.ErrOperationFailedTemplate

	// ErrFabricSpecificOperationFailed is the Fabric operation failure message template
	ErrFabricSpecificOperationFailed = "failed to %s Fabric %s: %w"
)
