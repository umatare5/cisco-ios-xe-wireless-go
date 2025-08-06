package model

// BleOperResponse represents the response structure for BLE operational data.
type BleOperResponse struct {
	BleOperData BleOperData `json:"Cisco-IOS-XE-wireless-ble-oper:ble-oper-data"`
}

// BleOperData contains BLE operational data
type BleOperData struct {
	BleApInfoList []BleApInfo `json:"ble-ap-info-list"`
	BleStats      BleStats    `json:"ble-stats"`
}

// BleApInfo represents BLE AP information
type BleApInfo struct {
	ApMac       string `json:"ap-mac"`
	BleEnabled  bool   `json:"ble-enabled"`
	BeaconCount int    `json:"beacon-count"`
}

// BleStats represents BLE statistics
type BleStats struct {
	TotalAps      int `json:"total-aps"`
	EnabledAps    int `json:"enabled-aps"`
	TotalBeacons  int `json:"total-beacons"`
	ActiveBeacons int `json:"active-beacons"`
}
