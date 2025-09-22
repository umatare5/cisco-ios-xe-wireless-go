package routes

// RFID (Radio-Frequency Identification) Configuration and Operational Paths
//
// These constants define the RESTCONF API paths for RFID configuration
// and operational data based on Cisco-IOS-XE-wireless-rfid YANG models.

// RFID Configuration Paths.
const (
	// RFIDCfgPath provides the path for RFID configuration data.
	RFIDCfgPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rfid-cfg:rfid-cfg-data"

	// RFIDCfgRFIDConfigPath provides the path for RFID configuration settings.
	RFIDCfgRFIDConfigPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rfid-cfg:rfid-cfg-data/rfid-config"
)

// RFID Operational Paths.
const (
	// RFIDOperPath provides the path for RFID operational data.
	RFIDOperPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rfid-oper:rfid-oper-data"
)

// RFID Global Operational Paths.
const (
	// RFIDGlobalOperPath provides the path for RFID global operational data.
	RFIDGlobalOperPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rfid-global-oper:rfid-global-oper-data"

	// RFIDDataDetailPath provides the path for RFID data detail.
	RFIDDataDetailPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rfid-global-oper:rfid-global-oper-data/rfid-data-detail"

	// RFIDRadioDataPath provides the path for RFID radio data.
	RFIDRadioDataPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rfid-global-oper:rfid-global-oper-data/rfid-radio-data"
)

// RFID Query Paths.
const (
	// RFIDDataDetailQueryPath provides the path for querying RFID data detail by MAC.
	RFIDDataDetailQueryPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rfid-global-oper:rfid-global-oper-data/rfid-data-detail"

	// RFIDDataQueryPath provides the path for querying RFID data by MAC.
	RFIDDataQueryPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rfid-oper:rfid-oper-data/rfid-data"
)
