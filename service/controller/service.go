package controller

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides controller management operations for Cisco IOS-XE Wireless LAN Controller.
type Service struct {
	service.BaseService
}

// NewService creates a new Controller service instance with the provided client.
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}

// GetBootTime retrieves the instant at which the controller last booted.
func (s Service) GetBootTime(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEDeviceHardwareOperBootTime, error) {
	return core.Get[CiscoIOSXEDeviceHardwareOperBootTime](ctx, s.Client(), routes.ControllerBootTimePath, opts...)
}

// Reload restarts the WNC controller causing temporary service interruption.
func (s Service) Reload(ctx context.Context, reason string, force bool) error {
	if strings.TrimSpace(reason) == "" {
		return errors.New(ErrInvalidReloadReason)
	}

	return s.reload(ctx, reason, &force)
}

// ReloadWithReason restarts the WNC controller with the specified reason only.
func (s Service) ReloadWithReason(ctx context.Context, reason string) error {
	if strings.TrimSpace(reason) == "" {
		return errors.New(ErrInvalidReloadReason)
	}

	return s.reload(ctx, reason, nil)
}

// reload is the internal helper function for WNC controller reload operations.
func (s Service) reload(ctx context.Context, reason string, force *bool) error {
	requestBody := WNCReloadRPCPayload{
		Input: WNCReloadRPCInput{
			Reason: reason,
			Force:  force,
		},
	}

	err := core.PostRPCVoid(ctx, s.Client(), routes.ControllerReloadRPC, requestBody)
	if err != nil {
		return fmt.Errorf("controller reload operation failed: %w",
			fmt.Errorf("reload RPC execution failed: %w", err))
	}

	return nil
}

// SaveConfig copies the running configuration to the startup configuration, which has no rollback.
//
// Output.Result is the controller's own account of the save — "Save running-config successful" on
// 17.12.8, 17.15.6 and 17.18.4a alike — and nothing here matches it, because a release wording it
// differently must not turn a completed save into a failure. The save took 2.5 to 3.7 seconds when
// measured, against a five-second response-header budget WithTimeout does not lift; raise it with
// WithResponseHeaderTimeout.
func (s Service) SaveConfig(ctx context.Context) (*SaveConfigRPCOutput, error) {
	return core.PostRPC[SaveConfigRPCOutput](ctx, s.Client(), routes.ControllerSaveConfigRPC, nil)
}
