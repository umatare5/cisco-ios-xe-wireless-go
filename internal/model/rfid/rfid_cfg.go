// Package model provides data models for RFID configuration data.
package model

// RfidCfg  represents the RFID configuration data.
type RfidCfg struct {
	RfidCfgData RfidCfgData `json:"Cisco-IOS-XE-wireless-rfid-cfg:rfid-cfg-data"`
}

// RfidCfgRfidConfig  represents the RFID configuration.
type RfidCfgRfidConfig struct {
	RfidConfig RfidConfig `json:"Cisco-IOS-XE-wireless-rfid-cfg:rfid-config"`
}

type RfidCfgData struct {
	RfidConfig RfidConfig `json:"rfid-config"`
}

type RfidConfig struct {
	RfidEnabled      bool   `json:"rfid-enabled"`
	DataTimeout      int    `json:"data-timeout"`
	TrackingEnabled  bool   `json:"tracking-enabled"`
	VendorData       string `json:"vendor-data"`
	NotificationAddr string `json:"notification-addr"`
	BloomFilter      bool   `json:"bloom-filter"`
}
