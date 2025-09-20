package rrm

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides RRM (Radio Resource Management) operations for Cisco IOS-XE Wireless LAN Controller.
type Service struct {
	service.BaseService
}

// NewService creates a new RRM service instance with the provided client.
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}

// GetConfig retrieves complete RRM configuration data.
func (s Service) GetConfig(ctx context.Context) (*RRMCfg, error) {
	return core.Get[RRMCfg](ctx, s.Client(), routes.RRMCfgPath)
}

// GetOperational retrieves RRM operational data.
func (s Service) GetOperational(ctx context.Context) (*RRMOper, error) {
	return core.Get[RRMOper](ctx, s.Client(), routes.RRMOperPath)
}

// GetGlobalOperational retrieves RRM global operational information.
func (s Service) GetGlobalOperational(ctx context.Context) (*RRMGlobalOper, error) {
	return core.Get[RRMGlobalOper](ctx, s.Client(), routes.RRMGlobalOperPath)
}

// GetEmulationOperational retrieves RRM emulation operational information.
func (s Service) GetEmulationOperational(ctx context.Context) (*RRMEmulOper, error) {
	return core.Get[RRMEmulOper](ctx, s.Client(), routes.RRMEmulOperPath)
}
