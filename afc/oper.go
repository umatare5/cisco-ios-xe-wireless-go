// Package afc provides Automated Frequency Coordination operational data functionality for the Cisco Wireless Network Controller API.
package afc

import (
	"context"
	"errors"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
)

const (
	// AfcOperBasePath defines the base path for AFC operational data endpoints.
	AfcOperBasePath = "Cisco-IOS-XE-wireless-afc-oper:afc-oper-data"
	// AfcOperEndpoint defines the endpoint for AFC operational data.
	AfcOperEndpoint = AfcOperBasePath
	// AfcOperEwlcAfcApRespEndpoint defines the endpoint for EWLC AFC AP response data.
	AfcOperEwlcAfcApRespEndpoint = AfcOperBasePath + "/ewlc-afc-ap-resp"
)

// Deprecated: use service-based API instead.
func GetAfcOper(client *wnc.Client, ctx context.Context) (*model.AfcOperResponse, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	service := NewService(client.CoreClient())
	return service.Oper(ctx)
}

// Deprecated: use service-based API instead.
func GetAfcEwlcAfcApResp(client *wnc.Client, ctx context.Context) (*model.AfcOperEwlcAfcApRespResponse, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	service := NewService(client.CoreClient())
	return service.APResp(ctx)
}
