// Package rfid provides RFID configuration functionality for the Cisco Wireless Network Controller API.
package rfid

import (
	"context"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
)

// WARNING: The RFID configuration endpoints have been reported as "Blank" in testing.
// See docs/missing_endpoints.md for details about endpoint availability.
// These endpoints may not return meaningful data on some WNC versions:
// - Cisco-IOS-XE-wireless-rfid-cfg:rfid-cfg-data
// - Cisco-IOS-XE-wireless-rfid-global-oper:rfid-global-oper-data
// - Cisco-IOS-XE-wireless-rfid-oper:rfid-oper-data

const (
	// RfidCfgBasePath defines the base path for RFID configuration endpoints
	RfidCfgBasePath = "/restconf/data/Cisco-IOS-XE-wireless-rfid-cfg:rfid-cfg-data"
	// RfidCfgEndpoint retrieves complete RFID configuration data
	RfidCfgEndpoint = RfidCfgBasePath
)

// RfidCfgResponse represents the complete RFID configuration response
type RfidCfgResponse struct {
	CiscoIOSXEWirelessRfidCfgData struct {
		Rfid struct{} `json:"rfid"`
	} `json:"Cisco-IOS-XE-wireless-rfid-cfg:rfid-cfg-data"`
}

// RfidResponse represents the RFID configuration response
type RfidResponse struct {
	Rfid struct{} `json:"Cisco-IOS-XE-wireless-rfid-cfg:rfid"`
}

// GetRfidCfg retrieves complete RFID configuration data.
func GetRfidCfg(client *wnc.Client, ctx context.Context) (*RfidCfgResponse, error) {
	var data RfidCfgResponse
	if err := client.SendAPIRequest(ctx, RfidCfgEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
