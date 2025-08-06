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
	GeneralOperBasePath = "Cisco-IOS-XE-wireless-general-oper:general-oper-data"
	// GeneralOperEndpoint retrieves complete general operational data
	GeneralOperEndpoint = GeneralOperBasePath
	// GeneralOperMgmtIntfDataEndpoint retrieves management interface data
	GeneralOperMgmtIntfDataEndpoint = GeneralOperBasePath + "/mgmt-intf-data"
)

// GetGeneralOper retrieves general operational data.
// Deprecated: Use general.NewService(client.CoreClient()).Oper(ctx) instead.
func GetGeneralOper(client *wnc.Client, ctx context.Context) (*model.GeneralOperResponse, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	service := NewService(client.CoreClient())
	return service.Oper(ctx)
}

// GetGeneralOperMgmtIntfData retrieves management interface operational data.
// GetGeneralOperMgmtIntfData retrieves management interface operational data.
// Deprecated: Use general.NewService(client.CoreClient()).MgmtIntfData(ctx) instead.
func GetGeneralOperMgmtIntfData(client *wnc.Client, ctx context.Context) (*model.GeneralOperMgmtIntfDataResponse, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	service := NewService(client.CoreClient())
	return service.MgmtIntfData(ctx)
}
