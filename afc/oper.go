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
	AfcOperBasePath = "/restconf/data/Cisco-IOS-XE-wireless-afc-oper:afc-oper-data"
	// AfcOperEndpoint defines the endpoint for AFC operational data.
	AfcOperEndpoint = AfcOperBasePath
	// AfcOperEwlcAfcApRespEndpoint defines the endpoint for EWLC AFC AP response data.
	AfcOperEwlcAfcApRespEndpoint = AfcOperBasePath + "/ewlc-afc-ap-resp"
)

// GetAfcOper retrieves AFC operational data.
func GetAfcOper(client *wnc.Client, ctx context.Context) (*model.AfcOperResponse, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	var data model.AfcOperResponse
	if err := client.SendAPIRequest(ctx, AfcOperEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetAfcEwlcAfcApResp retrieves EWLC AFC AP response data.
func GetAfcEwlcAfcApResp(client *wnc.Client, ctx context.Context) (*model.AfcOperEwlcAfcApRespResponse, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	var data model.AfcOperEwlcAfcApRespResponse
	if err := client.SendAPIRequest(ctx, AfcOperEwlcAfcApRespEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
