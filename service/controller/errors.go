package controller

import (
	ierrors "github.com/umatare5/cisco-ios-xe-wireless-go/internal/errors"
)

// EntityTypeController represents the Controller entity type for error messages
const EntityTypeController = "Controller"

// Common error message shortcuts for backward compatibility
var (
	// ErrOperationFailed is the generic operation failure message template
	ErrOperationFailed = ierrors.ErrOperationFailedTemplate

	// ErrControllerSpecificOperationFailed is the Controller operation failure message template
	ErrControllerSpecificOperationFailed = "failed to %s Controller %s: %w"

	// ErrWNCReloadFailed is the error message for WNC reload operation failures
	ErrWNCReloadFailed = "failed to reload WNC controller: %w"

	// ErrInvalidReloadReason is the error message for invalid reload reason
	ErrInvalidReloadReason = "reload reason cannot be empty"
)
