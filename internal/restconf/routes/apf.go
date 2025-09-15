package routes

// APF (Application Policy Framework) Configuration Paths
//
// These constants define the RESTCONF API paths for Application Policy Framework
// configuration based on Cisco-IOS-XE-wireless-apf-cfg YANG model.

// APF Configuration Paths.
const (
	// APFCfgPath retrieves complete Application Policy Framework configuration data.
	APFCfgPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-apf-cfg:apf-cfg-data"
)
