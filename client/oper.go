// Package client provides client operational data functionality for the Cisco Wireless Network Controller API.
package client

import (
	"context"
	"fmt"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
)

// Client operational data API endpoints
const (
	// ClientOperBasePath is the base path for client operational data endpoints
	ClientOperBasePath = "Cisco-IOS-XE-wireless-client-oper:client-oper-data"
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
	// SisfDbMacEndpoint retrieves SISF database MAC information
	SisfDbMacEndpoint = ClientOperBasePath + "/sisf-db-mac"
	// DcInfoEndpoint retrieves discovery client information
	DcInfoEndpoint = ClientOperBasePath + "/dc-info"
)

// GetClientOper retrieves complete client operational data.
func GetClientOper(client *wnc.Client, ctx context.Context) (*model.ClientOperResponse, error) {
	if client == nil {
		return nil, fmt.Errorf("%w: client cannot be nil", wnc.ErrInvalidConfiguration)
	}
	var data model.ClientOperResponse
	if err := client.SendAPIRequest(ctx, ClientOperEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetClientOperCommonOperData retrieves common operational data for wireless clients.
func GetClientOperCommonOperData(client *wnc.Client, ctx context.Context) (*model.ClientOperCommonOperDataResponse, error) {
	if client == nil {
		return nil, fmt.Errorf("%w: client cannot be nil", wnc.ErrInvalidConfiguration)
	}
	var data model.ClientOperCommonOperDataResponse
	if err := client.SendAPIRequest(ctx, CommonOperDataEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetClientOperDot11OperData retrieves 802.11 operational data for wireless clients.
func GetClientOperDot11OperData(client *wnc.Client, ctx context.Context) (*model.ClientOperDot11OperDataResponse, error) {
	if client == nil {
		return nil, fmt.Errorf("%w: client cannot be nil", wnc.ErrInvalidConfiguration)
	}
	var data model.ClientOperDot11OperDataResponse
	if err := client.SendAPIRequest(ctx, Dot11OperDataEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetClientOperMobilityOperData retrieves mobility operational data for wireless clients.
func GetClientOperMobilityOperData(client *wnc.Client, ctx context.Context) (*model.ClientOperMobilityOperDataResponse, error) {
	if client == nil {
		return nil, fmt.Errorf("%w: client cannot be nil", wnc.ErrInvalidConfiguration)
	}
	var data model.ClientOperMobilityOperDataResponse
	if err := client.SendAPIRequest(ctx, MobilityOperDataEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetClientOperMmIfClientStats retrieves mobility manager interface client statistics.
func GetClientOperMmIfClientStats(client *wnc.Client, ctx context.Context) (*model.ClientOperMmIfClientStatsResponse, error) {
	if client == nil {
		return nil, fmt.Errorf("%w: client cannot be nil", wnc.ErrInvalidConfiguration)
	}
	var data model.ClientOperMmIfClientStatsResponse
	if err := client.SendAPIRequest(ctx, MmIfClientStatsEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetClientOperMmIfClientHistory retrieves mobility manager interface client history.
func GetClientOperMmIfClientHistory(client *wnc.Client, ctx context.Context) (*model.ClientOperMmIfClientHistoryResponse, error) {
	if client == nil {
		return nil, fmt.Errorf("%w: client cannot be nil", wnc.ErrInvalidConfiguration)
	}
	var data model.ClientOperMmIfClientHistoryResponse
	if err := client.SendAPIRequest(ctx, MmIfClientHistoryEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetClientOperTrafficStats retrieves traffic statistics for wireless clients.
func GetClientOperTrafficStats(client *wnc.Client, ctx context.Context) (*model.ClientOperTrafficStatsResponse, error) {
	if client == nil {
		return nil, fmt.Errorf("%w: client cannot be nil", wnc.ErrInvalidConfiguration)
	}
	var data model.ClientOperTrafficStatsResponse
	if err := client.SendAPIRequest(ctx, TrafficStatsEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetClientOperPolicyData retrieves policy data for wireless clients.
func GetClientOperPolicyData(client *wnc.Client, ctx context.Context) (*model.ClientOperPolicyDataResponse, error) {
	if client == nil {
		return nil, fmt.Errorf("%w: client cannot be nil", wnc.ErrInvalidConfiguration)
	}
	var data model.ClientOperPolicyDataResponse
	if err := client.SendAPIRequest(ctx, PolicyDataEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetClientOperSisfDbMac retrieves SISF database MAC information.
func GetClientOperSisfDbMac(client *wnc.Client, ctx context.Context) (*model.ClientOperSisfDbMacResponse, error) {
	if client == nil {
		return nil, fmt.Errorf("%w: client cannot be nil", wnc.ErrInvalidConfiguration)
	}
	var data model.ClientOperSisfDbMacResponse
	if err := client.SendAPIRequest(ctx, SisfDbMacEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetClientOperDcInfo retrieves discovery client information.
func GetClientOperDcInfo(client *wnc.Client, ctx context.Context) (*model.ClientOperDcInfoResponse, error) {
	if client == nil {
		return nil, fmt.Errorf("%w: client cannot be nil", wnc.ErrInvalidConfiguration)
	}
	var data model.ClientOperDcInfoResponse
	if err := client.SendAPIRequest(ctx, DcInfoEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
