package routes

// WLAN Configuration and Operational Paths
//
// These constants define the RESTCONF API paths for WLAN configuration
// and operational data based on Cisco-IOS-XE-wireless-wlan YANG models.

// WLAN Configuration Paths.
const (
	// WLANCfgPath defines complete WLAN configuration data path.
	WLANCfgPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data"

	// WLANWlanCfgEntriesPath defines WLAN configuration entries path.
	WLANWlanCfgEntriesPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/wlan-cfg-entries"

	// WLANWlanPoliciesPath defines WLAN policies path.
	WLANWlanPoliciesPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/wlan-policies"

	// WLANPolicyListEntriesPath defines policy list entries path.
	WLANPolicyListEntriesPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/policy-list-entries"

	// WLANWirelessAaaPolicyConfigsPath defines wireless AAA policy configurations path.
	WLANWirelessAaaPolicyConfigsPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/wireless-aaa-policy-configs"

	// WLANDot11beProfilesPath defines Wi-Fi 7 / 802.11be profiles path.
	WLANDot11beProfilesPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/dot11be-profiles"
)

// WLAN Operational Paths.
const (
	// WLANGlobalOperPath defines WLAN global operational data path.
	WLANGlobalOperPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-wlan-global-oper:wlan-global-oper-data"

	// WLANWlanInfoPath defines WLAN information path for global operational data.
	WLANWlanInfoPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-wlan-global-oper:wlan-global-oper-data/wlan-info"
)

// WLAN Query Paths.
const (
	// WLANPolicyListEntryQueryPath provides the path for querying policy list entry by tag name.
	WLANPolicyListEntryQueryPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/policy-list-entries/policy-list-entry"
)
