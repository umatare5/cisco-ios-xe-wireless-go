// Package flex provides FlexConnect configuration functionality for the Cisco Wireless Network Controller API.
package flex

import (
	"context"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
)

const (
	// FlexCfgBasePath defines the base path for FlexConnect configuration endpoints
	FlexCfgBasePath = "/restconf/data/Cisco-IOS-XE-wireless-flex-cfg:flex-cfg-data"
	// FlexCfgEndpoint retrieves complete FlexConnect configuration data
	FlexCfgEndpoint = FlexCfgBasePath
	// FlexCfgDataEndpoint retrieves FlexConnect policy entries
	FlexCfgDataEndpoint = FlexCfgBasePath + "/flex-policy-entries"
)

// FlexCfgResponse represents the complete FlexConnect configuration response
type FlexCfgResponse struct {
	CiscoIOSXEWirelessFlexCfgData struct {
		FlexCfgData FlexCfgData `json:"flex-policy-entries"`
	} `json:"Cisco-IOS-XE-wireless-flex-cfg:flex-cfg-data"`
}

// FlexCfgDataResponse represents the FlexConnect policy entries response
type FlexCfgDataResponse struct {
	FlexCfgData FlexCfgData `json:"Cisco-IOS-XE-wireless-flex-cfg:flex-policy-entries"`
}

// FlexCfgData contains FlexConnect policy configuration entries
type FlexCfgData struct {
	FlexPolicyEntry []FlexPolicyEntry `json:"flex-policy-entry"`
}

// FlexPolicyEntry represents a FlexConnect policy configuration
type FlexPolicyEntry struct {
	PolicyName    string `json:"policy-name"`           // FlexConnect policy name
	Description   string `json:"description,omitempty"` // Optional policy description
	IfNameVlanIds *struct {
		IfNameVlanId []struct {
			InterfaceName string `json:"interface-name"` // Interface name
			VlanID        int    `json:"vlan-id"`        // VLAN identifier
		} `json:"if-name-vlan-id"`
	} `json:"if-name-vlan-ids,omitempty"`
}

// GetFlexCfg retrieves complete FlexConnect configuration data.
func GetFlexCfg(client *wnc.Client, ctx context.Context) (*FlexCfgResponse, error) {
	var data FlexCfgResponse
	if err := client.SendAPIRequest(ctx, FlexCfgEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetFlexCfgData retrieves FlexConnect policy entries.
func GetFlexCfgData(client *wnc.Client, ctx context.Context) (*FlexCfgDataResponse, error) {
	var data FlexCfgDataResponse
	if err := client.SendAPIRequest(ctx, FlexCfgDataEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
