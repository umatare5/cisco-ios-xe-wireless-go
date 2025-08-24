package routes

import "github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf"

// Fabric Configuration Endpoints
//
// These constants define the RESTCONF API endpoints for Fabric configuration
// based on Cisco-IOS-XE-wireless-fabric-cfg YANG model.

const (
	// FabricCfgBasePath defines the base path for Fabric configuration endpoints
	FabricCfgBasePath = restconf.YANGModelPrefix + "fabric-cfg:fabric-cfg-data"
)

// Fabric Configuration Endpoints
const (
	// FabricCfgEndpoint retrieves complete Fabric configuration data
	FabricCfgEndpoint = FabricCfgBasePath

	// FabricCfgProfilesEndpoint retrieves Fabric profile data
	FabricCfgProfilesEndpoint = FabricCfgBasePath + "/fabric-profiles"

	// FabricCfgControlplaneNamesEndpoint retrieves Fabric controlplane names
	FabricCfgControlplaneNamesEndpoint = FabricCfgBasePath + "/fabric-controlplane-names"
)
