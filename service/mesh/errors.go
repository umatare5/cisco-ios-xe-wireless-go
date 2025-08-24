// Package mesh provides Mesh-specific errors for the Cisco IOS-XE Wireless Network Controller API.
package mesh

import (
	ierrors "github.com/umatare5/cisco-ios-xe-wireless-go/internal/errors"
)

// EntityTypeMesh represents the Mesh entity type for error messages
const EntityTypeMesh = "Mesh"

// Common error message shortcuts for backward compatibility
var (
	// ErrOperationFailed is the generic operation failure message template
	ErrOperationFailed = ierrors.ErrOperationFailedTemplate

	// ErrMeshSpecificOperationFailed is the Mesh operation failure message template
	ErrMeshSpecificOperationFailed = "failed to %s Mesh %s: %w"
)
