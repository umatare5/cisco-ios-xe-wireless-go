// Package fabric provides SD-Access fabric configuration functionality for the Cisco Wireless Network Controller API.
package fabric

import (
	"context"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
)

const (
	// FabricCfgBasePath defines the base path for fabric configuration endpoints
	FabricCfgBasePath = "/restconf/data/Cisco-IOS-XE-wireless-fabric-cfg:fabric-cfg-data"
	// FabricCfgEndpoint retrieves complete fabric configuration data
	FabricCfgEndpoint = FabricCfgBasePath
	// FabricControlplaneNamesEndpoint retrieves fabric control plane names
	FabricControlplaneNamesEndpoint = FabricCfgBasePath + "/fabric-controlplane-names"
)

// FabricCfgResponse represents the complete fabric configuration response
type FabricCfgResponse struct {
	CiscoIOSXEWirelessFabricCfgFabricCfgData struct {
		Fabric                  struct{}                `json:"fabric"`
		FabricControlplaneNames FabricControlplaneNames `json:"fabric-controlplane-names"`
	} `json:"Cisco-IOS-XE-wireless-fabric-cfg:fabric-cfg-data"`
}

// FabricResponse represents the fabric configuration response
type FabricResponse struct {
	Fabric struct{} `json:"Cisco-IOS-XE-wireless-fabric-cfg:fabric"`
}

// FabricControlplaneNamesResponse represents the fabric control plane names response
type FabricControlplaneNamesResponse struct {
	FabricControlplaneNames FabricControlplaneNames `json:"Cisco-IOS-XE-wireless-fabric-cfg:fabric-controlplane-names"`
}

// FabricControlplaneNames contains fabric control plane name entries
type FabricControlplaneNames struct {
	FabricControlplaneName []FabricControlplaneName `json:"fabric-controlplane-name"`
}

// FabricControlplaneName represents a fabric control plane configuration
type FabricControlplaneName struct {
	ControlPlaneName string `json:"control-plane-name"`    // Control plane name identifier
	Description      string `json:"description,omitempty"` // Optional description
}

// GetFabricCfg retrieves complete fabric configuration data.
func GetFabricCfg(client *wnc.Client, ctx context.Context) (*FabricCfgResponse, error) {
	var data FabricCfgResponse
	if err := client.SendAPIRequest(ctx, FabricCfgEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetFabricControlplaneNames retrieves fabric control plane names.
func GetFabricControlplaneNames(client *wnc.Client, ctx context.Context) (*FabricControlplaneNamesResponse, error) {
	var data FabricControlplaneNamesResponse
	if err := client.SendAPIRequest(ctx, FabricControlplaneNamesEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func GetFabric(client *wnc.Client, ctx context.Context) (*FabricResponse, error) {
	var data FabricResponse
	if err := client.SendAPIRequest(ctx, FabricCfgEndpoint+"/fabric", &data); err != nil {
		return nil, err
	}
	return &data, nil
}
