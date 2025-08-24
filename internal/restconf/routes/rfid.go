package routes

import "github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf"

// RFID (Radio-Frequency Identification) Configuration and Operational Endpoints
//
// These constants define the RESTCONF API endpoints for RFID configuration
// and operational data based on Cisco-IOS-XE-wireless-rfid YANG models.

// RFID Base Paths
const (
	// RFIDCfgBasePath defines the base path for RFID configuration endpoints
	RFIDCfgBasePath = restconf.YANGModelPrefix + "rfid-cfg:rfid-cfg-data"

	// RFIDOperBasePath defines the base path for RFID operational endpoints
	RFIDOperBasePath = restconf.YANGModelPrefix + "rfid-oper:rfid-oper-data"

	// RFIDGlobalOperBasePath defines the base path for RFID global operational endpoints
	RFIDGlobalOperBasePath = restconf.YANGModelPrefix + "rfid-global-oper:rfid-global-oper-data"
)

// RFID Configuration Endpoints
const (
	// RFIDCfgEndpoint defines the endpoint for RFID configuration data
	RFIDCfgEndpoint = RFIDCfgBasePath
)

// RFID Operational Endpoints
const (
	// RFIDOperEndpoint defines the endpoint for RFID operational data
	RFIDOperEndpoint = RFIDOperBasePath
)

// RFID Global Operational Endpoints
const (
	// RFIDGlobalOperEndpoint defines the endpoint for RFID global operational data
	RFIDGlobalOperEndpoint = RFIDGlobalOperBasePath

	// RfidDataDetailEndpoint defines the endpoint for RFID data detail
	RfidDataDetailEndpoint = RFIDGlobalOperBasePath + "/rfid-data-detail"

	// RfidRadioDataEndpoint defines the endpoint for RFID radio data
	RfidRadioDataEndpoint = RFIDGlobalOperBasePath + "/rfid-radio-data"
)
