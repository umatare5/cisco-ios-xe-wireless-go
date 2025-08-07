package client

import (
	"context"
	"net/http"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/constants"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
)

const (
	// ClientOperBasePath defines the base path for client operational data endpoints
	ClientOperBasePath = constants.YANGModelPrefix + "client-oper:client-oper-data"
	// ClientOperEndpoint retrieves complete client operational data
	ClientOperEndpoint = ClientOperBasePath
	// CommonOperDataEndpoint retrieves common operational data for clients
	CommonOperDataEndpoint = ClientOperBasePath + "/common-oper-data"
	// Dot11OperDataEndpoint retrieves 802.11 operational data for clients
	Dot11OperDataEndpoint = ClientOperBasePath + "/dot11-oper-data"
	// MobilityOperDataEndpoint retrieves mobility operational data for clients
	MobilityOperDataEndpoint = ClientOperBasePath + "/mobility-oper-data"
	// MmIfClientStatsEndpoint retrieves mobility manager interface client statistics
	MmIfClientStatsEndpoint = ClientOperBasePath + "/mm-if-client-stats"
	// MmIfClientHistoryEndpoint retrieves mobility manager interface client history
	MmIfClientHistoryEndpoint = ClientOperBasePath + "/mm-if-client-history"
	// TrafficStatsEndpoint retrieves client traffic statistics
	TrafficStatsEndpoint = ClientOperBasePath + "/traffic-stats"
	// PolicyDataEndpoint retrieves client policy data
	PolicyDataEndpoint = ClientOperBasePath + "/policy-data"
	// SisfDBMacEndpoint retrieves SISF database MAC information
	SisfDBMacEndpoint = ClientOperBasePath + "/sisf-db-mac"
	// DcInfoEndpoint retrieves discovery client information
	DcInfoEndpoint = ClientOperBasePath + "/dc-info"
)

// Service provides Client operations.
type Service struct {
	c *core.Client
}

// NewService creates a new service instance.
func NewService(c *core.Client) Service {
	return Service{c: c}
}

// Operational Data Methods

// Oper returns complete client operational data.
func (s Service) Oper(ctx context.Context) (*model.ClientOperResponse, error) {
	var out model.ClientOperResponse
	return &out, s.c.Do(ctx, http.MethodGet, ClientOperEndpoint, &out)
}

// CommonOperData returns common operational data for clients.
func (s Service) CommonOperData(ctx context.Context) (*model.ClientOperCommonOperDataResponse, error) {
	var out model.ClientOperCommonOperDataResponse
	return &out, s.c.Do(ctx, http.MethodGet, CommonOperDataEndpoint, &out)
}

// Dot11OperData returns 802.11 operational data for clients.
func (s Service) Dot11OperData(ctx context.Context) (*model.ClientOperDot11OperDataResponse, error) {
	var out model.ClientOperDot11OperDataResponse
	return &out, s.c.Do(ctx, http.MethodGet, Dot11OperDataEndpoint, &out)
}

// MobilityOperData returns mobility operational data for clients.
func (s Service) MobilityOperData(ctx context.Context) (*model.ClientOperMobilityOperDataResponse, error) {
	var out model.ClientOperMobilityOperDataResponse
	return &out, s.c.Do(ctx, http.MethodGet, MobilityOperDataEndpoint, &out)
}

// MmIfClientStats returns mobility manager interface client statistics.
func (s Service) MmIfClientStats(ctx context.Context) (*model.ClientOperMmIfClientStatsResponse, error) {
	var out model.ClientOperMmIfClientStatsResponse
	return &out, s.c.Do(ctx, http.MethodGet, MmIfClientStatsEndpoint, &out)
}

// MmIfClientHistory returns mobility manager interface client history.
func (s Service) MmIfClientHistory(ctx context.Context) (*model.ClientOperMmIfClientHistoryResponse, error) {
	var out model.ClientOperMmIfClientHistoryResponse
	return &out, s.c.Do(ctx, http.MethodGet, MmIfClientHistoryEndpoint, &out)
}

// TrafficStats returns client traffic statistics.
func (s Service) TrafficStats(ctx context.Context) (*model.ClientOperTrafficStatsResponse, error) {
	var out model.ClientOperTrafficStatsResponse
	return &out, s.c.Do(ctx, http.MethodGet, TrafficStatsEndpoint, &out)
}

// PolicyData returns client policy data.
func (s Service) PolicyData(ctx context.Context) (*model.ClientOperPolicyDataResponse, error) {
	var out model.ClientOperPolicyDataResponse
	return &out, s.c.Do(ctx, http.MethodGet, PolicyDataEndpoint, &out)
}

// SisfDBMac returns SISF database MAC information.
func (s Service) SisfDBMac(ctx context.Context) (*model.ClientOperSisfDBMacResponse, error) {
	var out model.ClientOperSisfDBMacResponse
	return &out, s.c.Do(ctx, http.MethodGet, SisfDBMacEndpoint, &out)
}

// DcInfo returns discovery client information.
func (s Service) DcInfo(ctx context.Context) (*model.ClientOperDcInfoResponse, error) {
	var out model.ClientOperDcInfoResponse
	return &out, s.c.Do(ctx, http.MethodGet, DcInfoEndpoint, &out)
}
