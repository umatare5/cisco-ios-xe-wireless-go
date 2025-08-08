package client

import (
	"context"

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
	return core.Get[model.ClientOperResponse](ctx, s.c, ClientOperEndpoint)
}

// CommonOperData returns common operational data for clients.
func (s Service) CommonOperData(ctx context.Context) (*model.ClientOperCommonOperDataResponse, error) {
	return core.Get[model.ClientOperCommonOperDataResponse](ctx, s.c, CommonOperDataEndpoint)
}

// Dot11OperData returns 802.11 operational data for clients.
func (s Service) Dot11OperData(ctx context.Context) (*model.ClientOperDot11OperDataResponse, error) {
	return core.Get[model.ClientOperDot11OperDataResponse](ctx, s.c, Dot11OperDataEndpoint)
}

// MobilityOperData returns mobility operational data for clients.
func (s Service) MobilityOperData(ctx context.Context) (*model.ClientOperMobilityOperDataResponse, error) {
	return core.Get[model.ClientOperMobilityOperDataResponse](ctx, s.c, MobilityOperDataEndpoint)
}

// MmIfClientStats returns mobility manager interface client statistics.
func (s Service) MmIfClientStats(ctx context.Context) (*model.ClientOperMmIfClientStatsResponse, error) {
	return core.Get[model.ClientOperMmIfClientStatsResponse](ctx, s.c, MmIfClientStatsEndpoint)
}

// MmIfClientHistory returns mobility manager interface client history.
func (s Service) MmIfClientHistory(ctx context.Context) (*model.ClientOperMmIfClientHistoryResponse, error) {
	return core.Get[model.ClientOperMmIfClientHistoryResponse](ctx, s.c, MmIfClientHistoryEndpoint)
}

// TrafficStats returns client traffic statistics.
func (s Service) TrafficStats(ctx context.Context) (*model.ClientOperTrafficStatsResponse, error) {
	return core.Get[model.ClientOperTrafficStatsResponse](ctx, s.c, TrafficStatsEndpoint)
}

// PolicyData returns client policy data.
func (s Service) PolicyData(ctx context.Context) (*model.ClientOperPolicyDataResponse, error) {
	return core.Get[model.ClientOperPolicyDataResponse](ctx, s.c, PolicyDataEndpoint)
}

// SisfDBMac returns SISF database MAC information.
func (s Service) SisfDBMac(ctx context.Context) (*model.ClientOperSisfDBMacResponse, error) {
	return core.Get[model.ClientOperSisfDBMacResponse](ctx, s.c, SisfDBMacEndpoint)
}

// DcInfo returns discovery client information.
func (s Service) DcInfo(ctx context.Context) (*model.ClientOperDcInfoResponse, error) {
	return core.Get[model.ClientOperDcInfoResponse](ctx, s.c, DcInfoEndpoint)
}
