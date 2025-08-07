package client

import (
	"context"
	"net/http"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
	"github.com/umatare5/cisco-ios-xe-wireless-go/wnc"
)

// Service provides Client operations.
type Service struct {
	c *wnc.Client
}

// NewService creates a new service instance.
func NewService(c *wnc.Client) Service {
	return Service{c: c}
}

// Operational Data Methods

// Oper returns complete client operational data.
func (s Service) Oper(ctx context.Context) (*model.ClientOperResponse, error) {
	var out model.ClientOperResponse
	return &out, s.c.Do(ctx, http.MethodGet,
		"Cisco-IOS-XE-wireless-client-oper:client-oper-data", &out)
}

// CommonOperData returns common operational data for clients.
func (s Service) CommonOperData(ctx context.Context) (*model.ClientOperCommonOperDataResponse, error) {
	var out model.ClientOperCommonOperDataResponse
	return &out, s.c.Do(ctx, http.MethodGet,
		"Cisco-IOS-XE-wireless-client-oper:client-oper-data/common-oper-data", &out)
}

// Dot11OperData returns 802.11 operational data for clients.
func (s Service) Dot11OperData(ctx context.Context) (*model.ClientOperDot11OperDataResponse, error) {
	var out model.ClientOperDot11OperDataResponse
	return &out, s.c.Do(ctx, http.MethodGet,
		"Cisco-IOS-XE-wireless-client-oper:client-oper-data/dot11-oper-data", &out)
}

// MobilityOperData returns mobility operational data for clients.
func (s Service) MobilityOperData(ctx context.Context) (*model.ClientOperMobilityOperDataResponse, error) {
	var out model.ClientOperMobilityOperDataResponse
	return &out, s.c.Do(ctx, http.MethodGet,
		"Cisco-IOS-XE-wireless-client-oper:client-oper-data/mobility-oper-data", &out)
}

// MmIfClientStats returns mobility manager interface client statistics.
func (s Service) MmIfClientStats(ctx context.Context) (*model.ClientOperMmIfClientStatsResponse, error) {
	var out model.ClientOperMmIfClientStatsResponse
	return &out, s.c.Do(ctx, http.MethodGet,
		"Cisco-IOS-XE-wireless-client-oper:client-oper-data/mm-if-client-stats", &out)
}

// MmIfClientHistory returns mobility manager interface client history.
func (s Service) MmIfClientHistory(ctx context.Context) (*model.ClientOperMmIfClientHistoryResponse, error) {
	var out model.ClientOperMmIfClientHistoryResponse
	return &out, s.c.Do(ctx, http.MethodGet,
		"Cisco-IOS-XE-wireless-client-oper:client-oper-data/mm-if-client-history", &out)
}

// TrafficStats returns client traffic statistics.
func (s Service) TrafficStats(ctx context.Context) (*model.ClientOperTrafficStatsResponse, error) {
	var out model.ClientOperTrafficStatsResponse
	return &out, s.c.Do(ctx, http.MethodGet,
		"Cisco-IOS-XE-wireless-client-oper:client-oper-data/traffic-stats", &out)
}

// PolicyData returns client policy data.
func (s Service) PolicyData(ctx context.Context) (*model.ClientOperPolicyDataResponse, error) {
	var out model.ClientOperPolicyDataResponse
	return &out, s.c.Do(ctx, http.MethodGet,
		"Cisco-IOS-XE-wireless-client-oper:client-oper-data/policy-data", &out)
}

// SisfDBMac returns SISF database MAC information.
func (s Service) SisfDBMac(ctx context.Context) (*model.ClientOperSisfDBMacResponse, error) {
	var out model.ClientOperSisfDBMacResponse
	return &out, s.c.Do(ctx, http.MethodGet,
		"Cisco-IOS-XE-wireless-client-oper:client-oper-data/sisf-db-mac", &out)
}

// DcInfo returns discovery client information.
func (s Service) DcInfo(ctx context.Context) (*model.ClientOperDcInfoResponse, error) {
	var out model.ClientOperDcInfoResponse
	return &out, s.c.Do(ctx, http.MethodGet,
		"Cisco-IOS-XE-wireless-client-oper:client-oper-data/dc-info", &out)
}
