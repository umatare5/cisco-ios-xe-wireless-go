// Package rrm provides Radio Resource Management domain services for the Cisco Wireless Network Controller API.
package rrm

import (
	"context"
	"net/http"

	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
	"github.com/umatare5/cisco-ios-xe-wireless-go/wnc"
)

// Service provides access to all RRM operations.
type Service struct {
	c *wnc.Client
}

// NewService creates a new RRM service instance.
func NewService(c *wnc.Client) Service {
	return Service{c: c}
}

// Configuration Methods

// Cfg retrieves complete RRM configuration data.
func (s Service) Cfg(ctx context.Context) (*model.RrmCfgResponse, error) {
	var out model.RrmCfgResponse
	return &out, s.c.Do(ctx, http.MethodGet, RrmCfgEndpoint, &out)
}

// Operational Methods

// Oper retrieves RRM operational data.
func (s Service) Oper(ctx context.Context) (*model.RrmOperResponse, error) {
	var out model.RrmOperResponse
	return &out, s.c.Do(ctx, http.MethodGet, RrmOperEndpoint, &out)
}

// GlobalOper retrieves RRM global operational data.
func (s Service) GlobalOper(ctx context.Context) (*model.RrmGlobalOperResponse, error) {
	var out model.RrmGlobalOperResponse
	return &out, s.c.Do(ctx, http.MethodGet, RrmGlobalOperEndpoint, &out)
}

// EmulOper retrieves RRM emulation operational data.
func (s Service) EmulOper(ctx context.Context) (*model.RrmEmulOperResponse, error) {
	var out model.RrmEmulOperResponse
	return &out, s.c.Do(ctx, http.MethodGet, RrmEmulOperEndpoint, &out)
}
