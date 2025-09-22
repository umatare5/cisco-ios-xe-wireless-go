package routes

// BLE (Bluetooth Low Energy) Operational Paths
//
// These constants define the RESTCONF API paths for BLE operational
// data based on Cisco-IOS-XE-wireless-ble YANG models.

// BLE LTX Operational Paths.
const (
	// BLELtxOperPath retrieves complete BLE LTX operational data.
	BLELtxOperPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-ble-ltx-oper:ble-ltx-oper-data"

	// BLELtxApPath retrieves BLE LTX AP operational data.
	BLELtxApPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-ble-ltx-oper:ble-ltx-oper-data/ble-ltx-ap"

	// BLELtxApAntennaPath retrieves BLE LTX AP antenna operational data.
	BLELtxApAntennaPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-ble-ltx-oper:ble-ltx-oper-data/ble-ltx-ap-antenna"
)

// BLE Management Operational Paths.
const (
	// BLEMgmtOperPath retrieves complete BLE management operational data.
	BLEMgmtOperPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-ble-mgmt-oper:ble-mgmt-oper-data"

	// BLEMgmtApPath retrieves BLE management AP operational data.
	BLEMgmtApPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-ble-mgmt-oper:ble-mgmt-oper-data/ble-mgmt-ap"

	// BLEMgmtCmxPath retrieves BLE management CMX operational data.
	BLEMgmtCmxPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-ble-mgmt-oper:ble-mgmt-oper-data/ble-mgmt-cmx"
)
