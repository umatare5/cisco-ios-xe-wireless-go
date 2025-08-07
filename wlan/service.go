// Package wlan provides WLAN domain services for the Cisco Wireless Network Controller API.
package wlan

import (
	"context"
	"net/http"

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

// Cfg returns complete WLAN configuration data.
func (s Service) Cfg(ctx context.Context) (*model.WlanCfgResponse, error) {
	var out model.WlanCfgResponse
	return &out, s.c.Do(ctx, http.MethodGet, WlanCfgEndpoint, &out)
}

// CfgEntries returns WLAN configuration entries.
func (s Service) CfgEntries(ctx context.Context) (*model.WlanCfgEntriesResponse, error) {
	var out model.WlanCfgEntriesResponse
	return &out, s.c.Do(ctx, http.MethodGet, WlanCfgEntriesEndpoint, &out)
}

// Policies returns WLAN policies.
func (s Service) Policies(ctx context.Context) (*model.WlanPoliciesResponse, error) {
	var out model.WlanPoliciesResponse
	return &out, s.c.Do(ctx, http.MethodGet, WlanPoliciesEndpoint, &out)
}

// PolicyListEntries returns policy list entries.
func (s Service) PolicyListEntries(ctx context.Context) (*model.PolicyListEntriesResponse, error) {
	var out model.PolicyListEntriesResponse
	return &out, s.c.Do(ctx, http.MethodGet, PolicyListEntriesEndpoint, &out)
}

// WirelessAaaPolicyConfigs returns wireless AAA policy configurations.
func (s Service) WirelessAaaPolicyConfigs(ctx context.Context) (*model.WirelessAaaPolicyConfigsResponse, error) {
	var out model.WirelessAaaPolicyConfigsResponse
	return &out, s.c.Do(ctx, http.MethodGet, WirelessAaaPolicyConfigsEndpoint, &out)
}

// Operational Methods

// GlobalOper returns WLAN global operational data.
func (s Service) GlobalOper(ctx context.Context) (*model.WlanGlobalOperResponse, error) {
	var out model.WlanGlobalOperResponse
	return &out, s.c.Do(ctx, http.MethodGet, WlanGlobalOperDataEndpoint, &out)
}
