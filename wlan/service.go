// Package wlan provides WLAN domain services for the Cisco Wireless Network Controller API.
package wlan

import (
	"context"
	"net/http"

	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
	"github.com/umatare5/cisco-ios-xe-wireless-go/wnc"
)

// Service provides access to all WLAN operations.
type Service struct {
	c *wnc.Client
}

// NewService creates a new WLAN service instance.
func NewService(c *wnc.Client) Service {
	return Service{c: c}
}

// Configuration Methods

// Cfg retrieves complete WLAN configuration data.
func (s Service) Cfg(ctx context.Context) (*model.WlanCfgResponse, error) {
	var out model.WlanCfgResponse
	return &out, s.c.Do(ctx, http.MethodGet, WlanCfgEndpoint, &out)
}

// CfgEntries retrieves WLAN configuration entries.
func (s Service) CfgEntries(ctx context.Context) (*model.WlanCfgEntriesResponse, error) {
	var out model.WlanCfgEntriesResponse
	return &out, s.c.Do(ctx, http.MethodGet, WlanCfgEntriesEndpoint, &out)
}

// Policies retrieves WLAN policies.
func (s Service) Policies(ctx context.Context) (*model.WlanPoliciesResponse, error) {
	var out model.WlanPoliciesResponse
	return &out, s.c.Do(ctx, http.MethodGet, WlanPoliciesEndpoint, &out)
}

// PolicyListEntries retrieves policy list entries.
func (s Service) PolicyListEntries(ctx context.Context) (*model.PolicyListEntriesResponse, error) {
	var out model.PolicyListEntriesResponse
	return &out, s.c.Do(ctx, http.MethodGet, PolicyListEntriesEndpoint, &out)
}

// WirelessAaaPolicyConfigs retrieves wireless AAA policy configurations.
func (s Service) WirelessAaaPolicyConfigs(ctx context.Context) (*model.WirelessAaaPolicyConfigsResponse, error) {
	var out model.WirelessAaaPolicyConfigsResponse
	return &out, s.c.Do(ctx, http.MethodGet, WirelessAaaPolicyConfigsEndpoint, &out)
}

// Operational Methods

// GlobalOper retrieves WLAN global operational data.
func (s Service) GlobalOper(ctx context.Context) (*model.WlanGlobalOperResponse, error) {
	var out model.WlanGlobalOperResponse
	return &out, s.c.Do(ctx, http.MethodGet, WlanGlobalOperDataEndpoint, &out)
}
