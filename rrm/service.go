package rrm

import (
	"context"
	"net/http"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
	"github.com/umatare5/cisco-ios-xe-wireless-go/wnc"
)

const (
	// RrmCfgBasePath defines the base path for RRM configuration endpoints
	RrmCfgBasePath = "Cisco-IOS-XE-wireless-rrm-cfg:rrm-cfg-data"
	// RrmCfgEndpoint retrieves complete RRM configuration data
	RrmCfgEndpoint = RrmCfgBasePath

	// RrmOperBasePath defines the base path for RRM operational data endpoints
	RrmOperBasePath = "Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data"
	// RrmOperEndpoint retrieves RRM operational data
	RrmOperEndpoint = RrmOperBasePath

	// RrmGlobalOperBasePath defines the base path for RRM global operational data endpoints
	RrmGlobalOperBasePath = "Cisco-IOS-XE-wireless-rrm-global-oper:rrm-global-oper-data"
	// RrmGlobalOperEndpoint retrieves RRM global operational data
	RrmGlobalOperEndpoint = RrmGlobalOperBasePath

	// RrmEmulOperBasePath defines the base path for RRM emulation operational data endpoints
	RrmEmulOperBasePath = "Cisco-IOS-XE-wireless-rrm-emul-oper:rrm-emul-oper-data"
	// RrmEmulOperEndpoint retrieves RRM emulation operational data
	RrmEmulOperEndpoint = RrmEmulOperBasePath
)

// Service provides RRM operations.
type Service struct {
	c *wnc.Client
}

// NewService creates a new service instance.
func NewService(c *wnc.Client) Service {
	return Service{c: c}
}

// Configuration Methods

// Cfg returns complete RRM configuration data.
func (s Service) Cfg(ctx context.Context) (*model.RrmCfgResponse, error) {
	var out model.RrmCfgResponse
	return &out, s.c.Do(ctx, http.MethodGet, RrmCfgEndpoint, &out)
}

// Operational Methods

// Oper returns RRM operational data.
func (s Service) Oper(ctx context.Context) (*model.RrmOperResponse, error) {
	var out model.RrmOperResponse
	return &out, s.c.Do(ctx, http.MethodGet, RrmOperEndpoint, &out)
}

// GlobalOper returns RRM global operational data.
func (s Service) GlobalOper(ctx context.Context) (*model.RrmGlobalOperResponse, error) {
	var out model.RrmGlobalOperResponse
	return &out, s.c.Do(ctx, http.MethodGet, RrmGlobalOperEndpoint, &out)
}

// EmulOper returns RRM emulation operational data.
func (s Service) EmulOper(ctx context.Context) (*model.RrmEmulOperResponse, error) {
	var out model.RrmEmulOperResponse
	return &out, s.c.Do(ctx, http.MethodGet, RrmEmulOperEndpoint, &out)
}
