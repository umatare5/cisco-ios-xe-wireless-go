package routes

// WAT (Wireless Application Visibility and Control) Configuration Paths
//
// These constants define the RESTCONF API paths for WAT configuration
// based on Cisco-IOS-XE-wireless-wat-cfg YANG model.

// WAT Configuration Paths.
//
// The module is implemented on the 17.18 and 26.1 trains only, and wat-cfg-data holds a
// single child container, wat-config. Paths for wat-profiles, wat-thousandeyes and
// wat-test-profile were removed: no release declares those nodes.
const (
	// WATCfgPath provides the path for WAT configuration data.
	WATCfgPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-wat-cfg:wat-cfg-data"

	// WATConfigPath provides the path for the WAT configuration container.
	WATConfigPath = WATCfgPath + "/wat-config"

	// WATEnablePath provides the path for WAT enable configuration.
	WATEnablePath = WATConfigPath + "/wat-enable"
)

// WAT Operational Paths.
const (
	// WATOperPath provides the path for WAT operational data.
	//
	// No release measured implements the wat-oper module, so this path is unverified. It is
	// kept because a later train may add it, and it reads data rather than invoking an RPC.
	WATOperPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-wat-oper:wat-oper-data"
)
