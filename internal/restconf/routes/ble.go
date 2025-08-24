package routes

import "github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf"

// BLE (Bluetooth Low Energy) LTX Operational Endpoints
//
// These constants define the RESTCONF API endpoints for BLE LTX operational
// data based on Cisco-IOS-XE-wireless-ble-ltx-oper YANG model.

const (
	// BLELtxOperBasePath defines the base path for BLE LTX operational data endpoints
	BLELtxOperBasePath = restconf.YANGModelPrefix + "ble-ltx-oper:ble-ltx-oper-data"
)

// BLE LTX Operational Endpoints
const (
	// BLELtxOperEndpoint retrieves complete BLE LTX operational data
	BLELtxOperEndpoint = BLELtxOperBasePath
)
