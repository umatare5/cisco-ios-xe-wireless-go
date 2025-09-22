package routes

// 802.11 Configuration Paths
//
// These constants define the RESTCONF API paths for 802.11 standard
// configuration based on Cisco-IOS-XE-wireless-dot11-cfg YANG model.

// 802.11 Configuration Paths.
const (
	// Dot11CfgPath retrieves complete 802.11 configuration data.
	Dot11CfgPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-dot11-cfg:dot11-cfg-data"

	// Dot11ConfiguredCountriesPath retrieves configured countries.
	Dot11ConfiguredCountriesPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-dot11-cfg:dot11-cfg-data/configured-countries"

	// Dot11EntriesPath retrieves 802.11 entries.
	Dot11EntriesPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-dot11-cfg:dot11-cfg-data/dot11-entries"

	// Dot11AcMcsEntriesPath retrieves 802.11ac MCS entries.
	Dot11AcMcsEntriesPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-dot11-cfg:dot11-cfg-data/dot11ac-mcs-entries"
)
