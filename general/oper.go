// Package general provides general operational data functionality for the Cisco Wireless Network Controller API.
package general

import (
	"context"
	"errors"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
)

const (
	// GeneralOperBasePath defines the base path for general operational data endpoints
	GeneralOperBasePath = "/restconf/data/Cisco-IOS-XE-wireless-general-oper:general-oper-data"
	// GeneralOperEndpoint retrieves complete general operational data
	GeneralOperEndpoint = GeneralOperBasePath
	// GeneralOperMgmtIntfDataEndpoint retrieves management interface data
	GeneralOperMgmtIntfDataEndpoint = GeneralOperBasePath + "/mgmt-intf-data"
)

// GetGeneralOper retrieves general operational data.
func GetGeneralOper(client *wnc.Client, ctx context.Context) (*model.GeneralOperResponse, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	var data model.GeneralOperResponse
	if err := client.SendAPIRequest(ctx, GeneralOperEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetGeneralOperMgmtIntfData retrieves management interface operational data.
func GetGeneralOperMgmtIntfData(client *wnc.Client, ctx context.Context) (*model.GeneralOperMgmtIntfDataResponse, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	var data model.GeneralOperMgmtIntfDataResponse
	if err := client.SendAPIRequest(ctx, GeneralOperMgmtIntfDataEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
