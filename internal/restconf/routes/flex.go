package routes

// FlexConnect Configuration Paths
//
// These constants define the RESTCONF API paths for FlexConnect configuration
// based on Cisco-IOS-XE-wireless-flex-cfg YANG model.

// FlexConnect Configuration Paths.
const (
	// FlexCfgPath defines the path for FlexConnect configuration data.
	FlexCfgPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-flex-cfg:flex-cfg-data"

	// FlexPolicyEntriesPath defines the path for FlexConnect policy entries.
	FlexPolicyEntriesPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-flex-cfg:flex-cfg-data/flex-policy-entries"
)
