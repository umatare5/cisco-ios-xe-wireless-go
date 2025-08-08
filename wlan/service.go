// Package wlan provides WLAN domain services for the Cisco Wireless Network Controller API.
package wlan

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/constants"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
)

const (
	// WlanCfgBasePath defines the base path for WLAN configuration endpoints
	WlanCfgBasePath = constants.YANGModelPrefix + "wlan-cfg:wlan-cfg-data"
	// WlanCfgEndpoint defines complete WLAN configuration data endpoint
	WlanCfgEndpoint = WlanCfgBasePath
	// WlanCfgEntriesEndpoint defines WLAN configuration entries endpoint
	WlanCfgEntriesEndpoint = WlanCfgBasePath + "/wlan-cfg-entries"
	// WlanPoliciesEndpoint defines WLAN policies endpoint
	WlanPoliciesEndpoint = WlanCfgBasePath + "/wlan-policies"
	// PolicyListEntriesEndpoint defines policy list entries endpoint
	PolicyListEntriesEndpoint = WlanCfgBasePath + "/policy-list-entries"
	// WirelessAaaPolicyConfigsEndpoint defines wireless AAA policy configurations endpoint
	WirelessAaaPolicyConfigsEndpoint = WlanCfgBasePath + "/wireless-aaa-policy-configs"

	// WlanGlobalOperBasePath defines the base path for WLAN global operational data endpoints
	WlanGlobalOperBasePath = constants.YANGModelPrefix + "wlan-global-oper:wlan-global-oper-data"
	// WlanGlobalOperDataEndpoint defines WLAN global operational data endpoint
	WlanGlobalOperDataEndpoint = WlanGlobalOperBasePath
)

// Service provides WLAN operations.
type Service struct {
	c *core.Client
}

// NewService creates a new service instance.
func NewService(c *core.Client) Service {
	return Service{c: c}
}

// Configuration Methods

// GetCfg returns complete WLAN configuration data.
func (s Service) GetCfg(ctx context.Context) (*model.WlanCfgResponse, error) {
	return core.Get[model.WlanCfgResponse](ctx, s.c, WlanCfgEndpoint)
}

// GetCfgEntries returns WLAN configuration entries.
func (s Service) GetCfgEntries(ctx context.Context) (*model.WlanCfgEntriesResponse, error) {
	return core.Get[model.WlanCfgEntriesResponse](ctx, s.c, WlanCfgEntriesEndpoint)
}

// GetPolicies returns WLAN policies.
func (s Service) GetPolicies(ctx context.Context) (*model.WlanPoliciesResponse, error) {
	return core.Get[model.WlanPoliciesResponse](ctx, s.c, WlanPoliciesEndpoint)
}

// GetPolicyListEntries returns policy list entries.
func (s Service) GetPolicyListEntries(ctx context.Context) (*model.PolicyListEntriesResponse, error) {
	return core.Get[model.PolicyListEntriesResponse](ctx, s.c, PolicyListEntriesEndpoint)
}

// GetWirelessAaaPolicyConfigs returns wireless AAA policy configurations.
func (s Service) GetWirelessAaaPolicyConfigs(ctx context.Context) (*model.WirelessAaaPolicyConfigsResponse, error) {
	return core.Get[model.WirelessAaaPolicyConfigsResponse](ctx, s.c, WirelessAaaPolicyConfigsEndpoint)
}

// Operational Methods

// GetGlobalOper returns WLAN global operational data.
func (s Service) GetGlobalOper(ctx context.Context) (*model.WlanGlobalOperResponse, error) {
	return core.Get[model.WlanGlobalOperResponse](ctx, s.c, WlanGlobalOperDataEndpoint)
}
