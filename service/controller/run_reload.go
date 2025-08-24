package controller

import (
	"context"
	"errors"
	"fmt"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/controller"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/validation"
)

// Reload restarts the WNC controller causing temporary service interruption.
//
// This operation performs a cold restart of the wireless controller with the specified reason.
// The operation will cause complete service interruption and the controller will become
// temporarily unavailable for all operations during the restart process.
//
// Parameters:
//   - ctx: Context for request lifecycle management
//   - reason: Description for the reload operation (required)
//   - force: Optional flag to force restart even with unsaved config
//
// Returns:
//   - error: Operation error if the reload request fails
//
// Example:
//
//	err := controllerService.Reload(ctx, "Scheduled maintenance", false)
//	if err != nil {
//		log.Printf("Failed to reload controller: %v", err)
//		return
//	}
//	log.Println("Controller reload initiated successfully")
//
// Warning: This operation will cause service interruption. Use with caution.
func (s Service) Reload(ctx context.Context, reason string, force bool) error {
	if validation.IsStringEmpty(reason) {
		return errors.New(ErrInvalidReloadReason)
	}

	return s.reload(ctx, reason, &force)
}

// ReloadWithReason restarts the WNC controller with the specified reason only.
//
// This operation performs a cold restart of the wireless controller with the specified reason
// without forcing restart if there is unsaved configuration. This is the safer variant that
// will not restart if there are pending configuration changes.
//
// Parameters:
//   - ctx: Context for request lifecycle management
//   - reason: Description for the reload operation (required)
//
// Returns:
//   - error: Operation error if the reload request fails
//
// Example:
//
//	err := controllerService.ReloadWithReason(ctx, "Memory leak detected")
//	if err != nil {
//		log.Printf("Failed to reload controller: %v", err)
//		return
//	}
//	log.Println("Controller reload initiated successfully")
//
// Warning: This operation will cause service interruption. Use with caution.
func (s Service) ReloadWithReason(ctx context.Context, reason string) error {
	if validation.IsStringEmpty(reason) {
		return errors.New(ErrInvalidReloadReason)
	}

	return s.reload(ctx, reason, nil)
}

// reload is the internal helper function for WNC controller reload operations
func (s Service) reload(ctx context.Context, reason string, force *bool) error {
	requestBody := model.WNCReloadRPCPayload{
		Input: model.WNCReloadRPCInput{
			Reason: reason,
			Force:  force,
		},
	}

	err := core.PostRPCVoid(ctx, s.Client(), routes.WNCReloadRPC, requestBody)
	if err != nil {
		return fmt.Errorf(ErrWNCReloadFailed, err)
	}

	return nil
}
