// Package ap provides access point domain services for the Cisco Wireless Network Controller API.
package ap

import (
	"context"
	"net/http"

	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
	"github.com/umatare5/cisco-ios-xe-wireless-go/wnc"
)

// Service provides access to all access point operations.
type Service struct {
	c *wnc.Client
}

// NewService creates a new access point service instance.
func NewService(c *wnc.Client) Service {
	return Service{c: c}
}

// Configuration Methods

// Cfg retrieves complete access point configuration data.
func (s Service) Cfg(ctx context.Context) (*model.ApCfgResponse, error) {
	var out model.ApCfgResponse
	return &out, s.c.Do(ctx, http.MethodGet, ApCfgEndpoint, &out)
}

// TagSourcePriorityConfigs retrieves tag source priority configurations.
func (s Service) TagSourcePriorityConfigs(ctx context.Context) (*model.TagSourcePriorityConfigs, error) {
	var out model.TagSourcePriorityConfigs
	return &out, s.c.Do(ctx, http.MethodGet, TagSourcePriorityConfigsEndpoint, &out)
}

// ApTags retrieves access point tag configurations.
func (s Service) ApTags(ctx context.Context) (*model.ApCfgApTagsResponse, error) {
	var out model.ApCfgApTagsResponse
	return &out, s.c.Do(ctx, http.MethodGet, ApTagsEndpoint, &out)
}

// Operational Methods

// Oper retrieves complete access point operational data.
func (s Service) Oper(ctx context.Context) (*model.ApOperResponse, error) {
	var out model.ApOperResponse
	return &out, s.c.Do(ctx, http.MethodGet, ApOperEndpoint, &out)
}

// RadioNeighbor retrieves access point radio neighbor information.
func (s Service) RadioNeighbor(ctx context.Context) (*model.ApOperApRadioNeighborResponse, error) {
	var out model.ApOperApRadioNeighborResponse
	return &out, s.c.Do(ctx, http.MethodGet, ApRadioNeighborEndpoint, &out)
}

// NameMacMap retrieves the mapping between AP names and MAC addresses.
func (s Service) NameMacMap(ctx context.Context) (*[]model.ApNameMacMap, error) {
	var resp struct {
		Data []model.ApNameMacMap `json:"Cisco-IOS-XE-wireless-access-point-oper:ap-name-mac-map"`
	}
	err := s.c.Do(ctx, http.MethodGet, ApNameMacMapEndpoint, &resp)
	return &resp.Data, err
}

// CapwapData retrieves CAPWAP protocol data.
func (s Service) CapwapData(ctx context.Context) (*[]model.CapwapData, error) {
	var resp struct {
		Data []model.CapwapData `json:"Cisco-IOS-XE-wireless-access-point-oper:capwap-data"`
	}
	err := s.c.Do(ctx, http.MethodGet, CapwapDataEndpoint, &resp)
	return &resp.Data, err
}

// Global Operational Methods

// GlobalOper retrieves complete AP global operational data.
func (s Service) GlobalOper(ctx context.Context) (*model.ApGlobalOperResponse, error) {
	var out model.ApGlobalOperResponse
	return &out, s.c.Do(ctx, http.MethodGet, ApGlobalOperEndpoint, &out)
}

// History retrieves AP history data.
func (s Service) History(ctx context.Context) (*model.ApGlobalOperApHistoryResponse, error) {
	var out model.ApGlobalOperApHistoryResponse
	return &out, s.c.Do(ctx, http.MethodGet, ApHistoryEndpoint, &out)
}

// EwlcApStats retrieves EWLC AP statistics.
func (s Service) EwlcApStats(ctx context.Context) (*model.ApGlobalOperEwlcApStatsResponse, error) {
	var out model.ApGlobalOperEwlcApStatsResponse
	return &out, s.c.Do(ctx, http.MethodGet, EwlcApStatsEndpoint, &out)
}
