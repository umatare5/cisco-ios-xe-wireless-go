// Package constants provides shared constants used across the Cisco IOS-XE Wireless Go SDK.
package constants

// YANG Model prefixes
const (
	// YANGModelPrefix is the standard prefix for wireless-related YANG models
	YANGModelPrefix = "Cisco-IOS-XE-wireless-"

	// YANGModelPrefixAccess is the prefix for access-related YANG models
	YANGModelPrefixAccess = "Cisco-IOS-XE-access-"

	// YANGModelPrefixSite is the prefix for site-related YANG models
	YANGModelPrefixSite = "Cisco-IOS-XE-site-"
)

// YANG Module Suffixes
const (
	// CfgSuffix represents configuration module suffix
	CfgSuffix = "-cfg"

	// OperSuffix represents operational module suffix
	OperSuffix = "-oper"

	// CfgDataSuffix represents configuration data container suffix
	CfgDataSuffix = "-cfg-data"

	// OperDataSuffix represents operational data container suffix
	OperDataSuffix = "-oper-data"
)

// YANG Module Names (without prefix)
const (
	YangModuleAFC           = "afc"
	YangModuleAP            = "ap"
	YangModuleAPF           = "apf"
	YangModuleAWIPS         = "awips"
	YangModuleBLE           = "ble"
	YangModuleClient        = "client"
	YangModuleCTS           = "cts"
	YangModuleDot11         = "dot11"
	YangModuleDot15         = "dot15"
	YangModuleFabric        = "fabric"
	YangModuleFlex          = "flex"
	YangModuleGeneral       = "general"
	YangModuleGeolocation   = "geolocation"
	YangModuleHyperlocation = "hyperlocation"
	YangModuleLISP          = "lisp"
	YangModuleLocation      = "location"
	YangModuleMcast         = "mcast"
	YangModuleMDNS          = "mdns"
	YangModuleMesh          = "mesh"
	YangModuleMobility      = "mobility"
	YangModuleNMSP          = "nmsp"
	YangModuleRadio         = "radio"
	YangModuleRF            = "rf"
	YangModuleRFID          = "rfid"
	YangModuleRogue         = "rogue"
	YangModuleRRM           = "rrm"
	YangModuleSite          = "site"
	YangModuleWLAN          = "wlan"
)

// AP Service specific YANG models
const (
	APCfgModel        = YANGModelPrefix + "ap-cfg"
	APOperModel       = YANGModelPrefix + "access-point-oper"
	APGlobalOperModel = YANGModelPrefix + "ap-global-oper"
)

// BuildYangModulePath constructs a YANG module path using the common pattern
// Examples:
//
//	BuildYangModulePath("wlan", "cfg") -> "Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data"
//	BuildYangModulePath("ap", "oper") -> "Cisco-IOS-XE-wireless-ap-oper:ap-oper-data"
func BuildYangModulePath(module, moduleType string) string {
	return YANGModelPrefix + module + "-" + moduleType + ":" + module + "-" + moduleType + "-data"
}

// BuildWirelessYangModule constructs a wireless YANG module name
// Examples:
//
//	BuildWirelessYangModule("wlan", "cfg") -> "Cisco-IOS-XE-wireless-wlan-cfg"
//	BuildWirelessYangModule("ap", "oper") -> "Cisco-IOS-XE-wireless-ap-oper"
func BuildWirelessYangModule(module, moduleType string) string {
	return YANGModelPrefix + module + "-" + moduleType
}

// BuildYangEndpoint constructs a YANG endpoint path
// Examples:
//
//	BuildYangEndpoint("wlan", "cfg", "wlan-cfg-entries") -> "Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/wlan-cfg-entries"
func BuildYangEndpoint(module, moduleType, endpoint string) string {
	basePath := BuildYangModulePath(module, moduleType)
	if endpoint == "" {
		return basePath
	}
	return basePath + "/" + endpoint
}

// BuildAPCfgPath builds AP configuration paths
func BuildAPCfgPath(endpoint string) string {
	if endpoint == "" {
		return APCfgModel + ":" + YangModuleAP + CfgDataSuffix
	}
	return APCfgModel + ":" + YangModuleAP + CfgDataSuffix + "/" + endpoint
}

// BuildAPOperPath builds AP operational paths
func BuildAPOperPath(endpoint string) string {
	basePath := "/" + APOperModel + ":" + "access-point" + OperDataSuffix
	if endpoint == "" {
		return basePath
	}
	return basePath + "/" + endpoint
}

// BuildAPGlobalOperPath builds AP global operational paths
func BuildAPGlobalOperPath(endpoint string) string {
	basePath := APGlobalOperModel + ":" + YangModuleAP + "-global" + OperDataSuffix
	if endpoint == "" {
		return basePath
	}
	return basePath + "/" + endpoint
}
