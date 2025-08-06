// Package nmsp provides Network Mobility Services Protocol operational data functionality for the Cisco Wireless Network Controller API.
package nmsp

import (
	"context"
	"errors"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
)

const (
	// NmspOperBasePath defines the base path for NMSP operational data endpoints.
	NmspOperBasePath = "Cisco-IOS-XE-wireless-nmsp-oper:nmsp-oper-data"
	// NmspOperEndpoint defines the endpoint for NMSP operational data.
	NmspOperEndpoint = NmspOperBasePath
	// ClientRegistrationEndpoint defines the endpoint for client registration data.
	ClientRegistrationEndpoint = NmspOperBasePath + "/client-registration"
	// CmxConnectionEndpoint defines the endpoint for CMX connection data.
	CmxConnectionEndpoint = NmspOperBasePath + "/cmx-connection"
	// CmxCloudInfoEndpoint defines the endpoint for CMX cloud information.
	CmxCloudInfoEndpoint = NmspOperBasePath + "/cmx-cloud-info"
)

// Type aliases for backward compatibility - will be removed in v2.0.0
type (
	// Deprecated: Use model.NmspOperResponse instead. Will be removed in v2.0.0.
	NmspOperResponse = model.NmspOperResponse
	// Deprecated: Use model.ClientRegistration instead. Will be removed in v2.0.0.
	ClientRegistration = model.ClientRegistration
	// Deprecated: Use model.NmspServices instead. Will be removed in v2.0.0.
	NmspServices = model.NmspServices
	// Deprecated: Use model.NmspClientRegistrationResponse instead. Will be removed in v2.0.0.
	NmspClientRegistrationResponse = model.NmspClientRegistrationResponse
	// Deprecated: Use model.NmspCmxConnectionResponse instead. Will be removed in v2.0.0.
	NmspCmxConnectionResponse = model.NmspCmxConnectionResponse
	// Deprecated: Use model.NmspCmxCloudInfoResponse instead. Will be removed in v2.0.0.
	NmspCmxCloudInfoResponse = model.NmspCmxCloudInfoResponse
	// Deprecated: Use model.CmxConnection instead. Will be removed in v2.0.0.
	CmxConnection = model.CmxConnection
	// Deprecated: Use model.CmxConStats instead. Will be removed in v2.0.0.
	CmxConStats = model.CmxConStats
	// Deprecated: Use model.MsgCounter instead. Will be removed in v2.0.0.
	MsgCounter = model.MsgCounter
	// Deprecated: Use model.CmxCloudInfo instead. Will be removed in v2.0.0.
	CmxCloudInfo = model.CmxCloudInfo
	// Deprecated: Use model.CloudStatus instead. Will be removed in v2.0.0.
	CloudStatus = model.CloudStatus
	// Deprecated: Use model.CloudStats instead. Will be removed in v2.0.0.
	CloudStats = model.CloudStats
)

// Deprecated: Use nmsp.NewService(client).Oper(ctx) instead. Will be removed in v2.0.0.
func GetNmspOper(client *wnc.Client, ctx context.Context) (*model.NmspOperResponse, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	return NewService(client).Oper(ctx)
}

// Deprecated: Use nmsp.NewService(client).ClientRegistration(ctx) instead. Will be removed in v2.0.0.
func GetNmspClientRegistration(client *wnc.Client, ctx context.Context) (*model.NmspClientRegistrationResponse, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	return NewService(client).ClientRegistration(ctx)
}

// Deprecated: Use nmsp.NewService(client).CmxConnection(ctx) instead. Will be removed in v2.0.0.
func GetNmspCmxConnection(client *wnc.Client, ctx context.Context) (*model.NmspCmxConnectionResponse, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	return NewService(client).CmxConnection(ctx)
}

// Deprecated: Use nmsp.NewService(client).CmxCloudInfo(ctx) instead. Will be removed in v2.0.0.
func GetNmspCmxCloudInfo(client *wnc.Client, ctx context.Context) (*model.NmspCmxCloudInfoResponse, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	return NewService(client).CmxCloudInfo(ctx)
}
