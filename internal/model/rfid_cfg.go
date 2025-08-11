package model

// RfidCfgResponse represents the response structure for RFID configuration data.
type RfidCfgResponse struct {
	RfidCfgData RfidCfgData `json:"Cisco-IOS-XE-wireless-rfid-cfg:rfid-cfg-data"`
}

// RfidCfgData contains RFID configuration data
type RfidCfgData struct {
	RfidConfig RfidConfig `json:"rfid-config"`
}

// RfidConfig represents RFID configuration settings
type RfidConfig struct {
	RfidEnabled      bool   `json:"rfid-enabled"`
	DataTimeout      int    `json:"data-timeout"`
	TrackingEnabled  bool   `json:"tracking-enabled"`
	VendorData       string `json:"vendor-data"`
	NotificationAddr string `json:"notification-addr"`
	BloomFilter      bool   `json:"bloom-filter"`
}
