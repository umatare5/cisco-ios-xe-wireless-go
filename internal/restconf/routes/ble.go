package routes

// BLE (Bluetooth Low Energy) LTX Operational Paths
//
// These constants define the RESTCONF API paths for BLE LTX operational
// data based on Cisco-IOS-XE-wireless-ble-ltx-oper YANG model.

// BLE LTX Operational Paths.
const (
	// BLELtxOperPath retrieves complete BLE LTX operational data.
	BLELtxOperPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-ble-ltx-oper:ble-ltx-oper-data"

	// BLELtxApPath retrieves BLE LTX AP operational data.
	BLELtxApPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-ble-ltx-oper:ble-ltx-oper-data/ble-ltx-ap"

	// BLELtxApAntennaPath retrieves BLE LTX AP antenna operational data.
	BLELtxApAntennaPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-ble-ltx-oper:ble-ltx-oper-data/ble-ltx-ap-antenna"
)
