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

	// WLANWlanCfgEntryPath defines individual WLAN configuration entry path.
	WLANWlanCfgEntryPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/wlan-cfg-entries/wlan-cfg-entry"

	// WLANWlanPoliciesPath defines WLAN policies path.
	WLANWlanPoliciesPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/wlan-policies"

	// WLANWlanPolicyEntryPath defines individual WLAN policy entry path.
	WLANWlanPolicyEntryPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/wlan-policies/wlan-policy"

	// WLANPolicyListEntriesPath defines policy list entries path.
	WLANPolicyListEntriesPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/policy-list-entries"

	// WLANPolicyListEntryPath defines individual policy list entry path.
	WLANPolicyListEntryPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/policy-list-entries/policy-list-entry"

	// WLANWirelessAaaPolicyConfigsPath defines wireless AAA policy configurations path.
	WLANWirelessAaaPolicyConfigsPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/wireless-aaa-policy-configs"
)

// WLAN Operational Paths.
const (
	// WLANGlobalOperPath defines WLAN global operational data path.
	WLANGlobalOperPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-wlan-global-oper:wlan-global-oper-data"

	// WLANWlanProfilesPath defines WLAN profiles path for operational data.
	WLANWlanProfilesPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-wlan-oper:wlan-oper/global-oper/wlan-profiles"
)

// WLAN Query Paths.
const (
	// WLANPolicyListEntryQueryPath provides the path for querying policy list entry by tag name.
	WLANPolicyListEntryQueryPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/policy-list-entries/policy-list-entry"
)
