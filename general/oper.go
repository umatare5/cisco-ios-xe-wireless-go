// Package general provides general operational data functionality for the Cisco Wireless Network Controller API.
package general

import (
	"context"
	"fmt"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
)

const (
	// GeneralOperBasePath defines the base path for general operational data endpoints
	GeneralOperBasePath = "/restconf/data/Cisco-IOS-XE-wireless-general-oper:general-oper-data"
	// GeneralOperEndpoint retrieves complete general operational data
	GeneralOperEndpoint = GeneralOperBasePath
	// GeneralOperMgmtIntfDataEndpoint retrieves management interface data
	GeneralOperMgmtIntfDataEndpoint = GeneralOperBasePath + "/mgmt-intf-data"
)

// GeneralOperResponse represents the complete general operational data response
type GeneralOperResponse struct {
	CiscoIOSXEWirelessGeneralOperData struct {
		MgmtIntfData MgmtIntfData `json:"mgmt-intf-data"`
	} `json:"Cisco-IOS-XE-wireless-general-oper:general-oper-data"`
}

// GeneralOperMgmtIntfDataResponse represents the management interface data response
type GeneralOperMgmtIntfDataResponse struct {
	MgmtIntfData MgmtIntfData `json:"Cisco-IOS-XE-wireless-general-oper:mgmt-intf-data"`
}

// MgmtIntfData contains management interface configuration and status information
type MgmtIntfData struct {
	IntfName string `json:"intf-name"` // Interface name
	IntfType string `json:"intf-type"` // Interface type
	IntfID   int    `json:"intf-id"`   // Interface ID
	MgmtIP   string `json:"mgmt-ip"`   // Management IP address
	NetMask  string `json:"net-mask"`  // Network mask
	MgmtMAC  string `json:"mgmt-mac"`  // Management MAC address
}

// GetGeneralOper retrieves general operational data.
func GetGeneralOper(client *wnc.Client, ctx context.Context) (*GeneralOperResponse, error) {
	if client == nil {
		return nil, fmt.Errorf("%w: client cannot be nil", wnc.ErrInvalidConfiguration)
	}
	var data GeneralOperResponse
	if err := client.SendAPIRequest(ctx, GeneralOperEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetGeneralOperMgmtIntfData retrieves management interface operational data.
func GetGeneralOperMgmtIntfData(client *wnc.Client, ctx context.Context) (*GeneralOperMgmtIntfDataResponse, error) {
	if client == nil {
		return nil, fmt.Errorf("%w: client cannot be nil", wnc.ErrInvalidConfiguration)
	}
	var data GeneralOperMgmtIntfDataResponse
	if err := client.SendAPIRequest(ctx, GeneralOperMgmtIntfDataEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
