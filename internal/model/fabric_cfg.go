package model

// FabricCfgResponse represents the response structure for Fabric configuration data.
type FabricCfgResponse struct {
	FabricCfgData FabricCfgData `json:"Cisco-IOS-XE-wireless-fabric-cfg:fabric-cfg-data"`
}

// FabricCfgData contains Fabric configuration data
type FabricCfgData struct {
	FabricConfig FabricConfig `json:"fabric-config"`
}

// FabricConfig represents SD-Access fabric configuration
type FabricConfig struct {
	FabricEnabled       bool   `json:"fabric-enabled"`
	ControlPlaneAddress string `json:"control-plane-address"`
	VirtualNetworkName  string `json:"virtual-network-name"`
	SubnetID            string `json:"subnet-id"`
	BorderNode          bool   `json:"border-node"`
}
