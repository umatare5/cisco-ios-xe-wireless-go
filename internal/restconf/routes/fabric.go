package routes

// Fabric Configuration Paths
//
// These constants define the RESTCONF API paths for Fabric configuration
// based on Cisco-IOS-XE-wireless-fabric-cfg YANG model.

// Fabric Configuration Paths.
const (
	// FabricCfgPath retrieves complete Fabric configuration data.
	FabricCfgPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-fabric-cfg:fabric-cfg-data"

	// FabricPath retrieves Fabric specific configuration.
	FabricPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-fabric-cfg:fabric-cfg-data/fabric"

	// FabricProfilesPath retrieves Fabric profile data.
	FabricProfilesPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-fabric-cfg:fabric-cfg-data/fabric-profiles"

	// FabricControlplaneNamesPath retrieves Fabric controlplane names.
	FabricControlplaneNamesPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-fabric-cfg:fabric-cfg-data/fabric-controlplane-names"
)
