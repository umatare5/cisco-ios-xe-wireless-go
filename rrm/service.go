package rrm

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/constants"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
)

const (
	// RrmCfgBasePath defines the base path for RRM configuration endpoints
	RrmCfgBasePath = constants.YANGModelPrefix + "rrm-cfg:rrm-cfg-data"
	// RrmCfgEndpoint retrieves complete RRM configuration data
	RrmCfgEndpoint = RrmCfgBasePath

	// RrmOperBasePath defines the base path for RRM operational data endpoints
	RrmOperBasePath = constants.YANGModelPrefix + "rrm-oper:rrm-oper-data"
	// RrmOperEndpoint retrieves RRM operational data
	RrmOperEndpoint = RrmOperBasePath

	// RrmGlobalOperBasePath defines the base path for RRM global operational data endpoints
	RrmGlobalOperBasePath = constants.YANGModelPrefix + "rrm-global-oper:rrm-global-oper-data"
	// RrmGlobalOperEndpoint retrieves RRM global operational data
	RrmGlobalOperEndpoint = RrmGlobalOperBasePath

	// RrmEmulOperBasePath defines the base path for RRM emulation operational data endpoints
	RrmEmulOperBasePath = constants.YANGModelPrefix + "rrm-emul-oper:rrm-emul-oper-data"
	// RrmEmulOperEndpoint retrieves RRM emulation operational data
	RrmEmulOperEndpoint = RrmEmulOperBasePath
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

// Cfg returns complete RRM configuration data.
func (s Service) Cfg(ctx context.Context) (*model.RrmCfgResponse, error) { return core.Get[model.RrmCfgResponse](ctx, s.c, RrmCfgEndpoint) }

// Operational Methods

// Oper returns RRM operational data.
func (s Service) Oper(ctx context.Context) (*model.RrmOperResponse, error) { return core.Get[model.RrmOperResponse](ctx, s.c, RrmOperEndpoint) }

// GlobalOper returns RRM global operational data.
func (s Service) GlobalOper(ctx context.Context) (*model.RrmGlobalOperResponse, error) {
	return core.Get[model.RrmGlobalOperResponse](ctx, s.c, RrmGlobalOperEndpoint)
}

// EmulOper returns RRM emulation operational data.
func (s Service) EmulOper(ctx context.Context) (*model.RrmEmulOperResponse, error) {
	return core.Get[model.RrmEmulOperResponse](ctx, s.c, RrmEmulOperEndpoint)
}
