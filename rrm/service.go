package rrm

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/constants"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
)

const (
	// RRMCfgBasePath defines the base path for RRM configuration endpoints
	RRMCfgBasePath = constants.YANGModelPrefix + "rrm-cfg:rrm-cfg-data"
	// RRMCfgEndpoint retrieves complete RRM configuration data
	RRMCfgEndpoint = RRMCfgBasePath

	// RRMOperBasePath defines the base path for RRM operational data endpoints
	RRMOperBasePath = constants.YANGModelPrefix + "rrm-oper:rrm-oper-data"
	// RRMOperEndpoint retrieves RRM operational data
	RRMOperEndpoint = RRMOperBasePath

	// RRMGlobalOperBasePath defines the base path for RRM global operational data endpoints
	RRMGlobalOperBasePath = constants.YANGModelPrefix + "rrm-global-oper:rrm-global-oper-data"
	// RRMGlobalOperEndpoint retrieves RRM global operational data
	RRMGlobalOperEndpoint = RRMGlobalOperBasePath

	// RRMEmulOperBasePath defines the base path for RRM emulation operational data endpoints
	RRMEmulOperBasePath = constants.YANGModelPrefix + "rrm-emul-oper:rrm-emul-oper-data"
	// RRMEmulOperEndpoint retrieves RRM emulation operational data
	RRMEmulOperEndpoint = RRMEmulOperBasePath
)

// Service provides RRM operations.
type Service struct {
	c *core.Client
}

// NewService creates a new service instance.
func NewService(c *core.Client) Service {
	return Service{c: c}
}

// Configuration Methods

// GetCfg returns complete RRM configuration data.
func (s Service) GetCfg(ctx context.Context) (*model.RrmCfgResponse, error) {
	return core.Get[model.RrmCfgResponse](ctx, s.c, RRMCfgEndpoint)
}

// Operational Methods

// GetOper returns RRM operational data.
func (s Service) GetOper(ctx context.Context) (*model.RrmOperResponse, error) {
	return core.Get[model.RrmOperResponse](ctx, s.c, RRMOperEndpoint)
}

// GetGlobalOper returns RRM global operational data.
func (s Service) GetGlobalOper(ctx context.Context) (*model.RrmGlobalOperResponse, error) {
	return core.Get[model.RrmGlobalOperResponse](ctx, s.c, RRMGlobalOperEndpoint)
}

// GetEmulOper returns RRM emulation operational data.
func (s Service) GetEmulOper(ctx context.Context) (*model.RrmEmulOperResponse, error) {
	return core.Get[model.RrmEmulOperResponse](ctx, s.c, RRMEmulOperEndpoint)
}
