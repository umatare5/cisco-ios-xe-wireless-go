// Package hyperlocation provides hyperlocation operational data functionality for the Cisco Wireless Network Controller API.
package hyperlocation

import (
	"context"
	"errors"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
)

const (
	// HyperlocationOperBasePath defines the base path for hyperlocation operational data endpoints.
	HyperlocationOperBasePath = "Cisco-IOS-XE-wireless-hyperlocation-oper:hyperlocation-oper-data"
	// HyperlocationOperEndpoint defines the endpoint for hyperlocation operational data.
	HyperlocationOperEndpoint = HyperlocationOperBasePath
	// HyperlocationProfilesEndpoint defines the endpoint for hyperlocation profiles.
	HyperlocationProfilesEndpoint = HyperlocationOperBasePath + "/ewlc-hyperlocation-profile"
)

// Type aliases for backward compatibility - will be removed in v2.0.0
type (
	// Deprecated: Use model.HyperlocationOperResponse instead. Will be removed in v2.0.0.
	HyperlocationOperResponse = model.HyperlocationOperResponse
	// Deprecated: Use model.HyperlocationProfilesResponse instead. Will be removed in v2.0.0.
	HyperlocationProfilesResponse = model.HyperlocationProfilesResponse
	// Deprecated: Use model.EwlcHyperlocationProfile instead. Will be removed in v2.0.0.
	EwlcHyperlocationProfile = model.EwlcHyperlocationProfile
)

// Deprecated: Use hyperlocation.NewService(client).Oper(ctx) instead. Will be removed in v2.0.0.
func GetHyperlocationOper(client *wnc.Client, ctx context.Context) (*model.HyperlocationOperResponse, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	return NewService(client).Oper(ctx)
}

// Deprecated: Use hyperlocation.NewService(client).Profiles(ctx) instead. Will be removed in v2.0.0.
func GetHyperlocationProfiles(client *wnc.Client, ctx context.Context) (*model.HyperlocationProfilesResponse, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	return NewService(client).Profiles(ctx)
}
