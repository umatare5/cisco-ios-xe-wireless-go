package routes

// WAT (Wireless Application Visibility and Control) Configuration Paths
//
// These constants define the RESTCONF API paths for WAT configuration
// based on Cisco-IOS-XE-wireless-wat-cfg YANG model.

// WAT Configuration Paths.
const (
	// WATCfgPath provides the path for WAT configuration data.
	WATCfgPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-wat-cfg:wat-cfg-data"

	// WATProfilesPath provides the path for WAT profiles.
	WATProfilesPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-wat-cfg:wat-cfg-data/wat-profiles"

	// WATEnablePath provides the path for WAT enable configuration.
	WATEnablePath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-wat-cfg:wat-cfg-data/wat-enable"

	// WATThousandeyesPath provides the path for WAT Thousandeyes configuration.
	WATThousandeyesPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-wat-cfg:wat-cfg-data/wat-thousandeyes"

	// WATTestProfilePath provides the path for WAT test profiles.
	WATTestProfilePath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-wat-cfg:wat-cfg-data/wat-test-profile"
)

// WAT Operational Paths.
const (
	// WATOperPath provides the path for WAT operational data.
	WATOperPath = RESTCONFOperationsPath + "/Cisco-IOS-XE-wireless-wat-oper:wat-oper-data"
)
