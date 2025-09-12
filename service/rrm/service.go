package rrm

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/rrm"
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
func (s Service) GetConfig(ctx context.Context) (*model.RrmCfg, error) {
	return core.Get[model.RrmCfg](ctx, s.Client(), routes.RRMCfgPath)
}

// GetOperational retrieves RRM operational data.
func (s Service) GetOperational(ctx context.Context) (*model.RrmOper, error) {
	return core.Get[model.RrmOper](ctx, s.Client(), routes.RRMOperPath)
}

// GetGlobalInfo retrieves RRM global operational information.
func (s Service) GetGlobalInfo(ctx context.Context) (*model.RrmGlobalOper, error) {
	return core.Get[model.RrmGlobalOper](ctx, s.Client(), routes.RRMGlobalOperPath)
}

// GetEmulationInfo retrieves RRM emulation operational information.
func (s Service) GetEmulationInfo(ctx context.Context) (*model.RrmEmulOper, error) {
	return core.Get[model.RrmEmulOper](ctx, s.Client(), routes.RRMEmulOperPath)
}
