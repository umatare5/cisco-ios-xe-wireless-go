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
