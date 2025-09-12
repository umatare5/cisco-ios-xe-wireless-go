// Package rfid provides data models for RFID configuration data.
package rfid

// RfidCfg represents RFID configuration data.
type RfidCfg struct {
	RfidCfgData RfidCfgData `json:"Cisco-IOS-XE-wireless-rfid-cfg:rfid-cfg-data"`
}

// RfidCfgData represents RFID configuration data container.
type RfidCfgData struct {
	Rfid *RfidConfig `json:"rfid,omitempty"`
}

// RfidConfig represents RFID configuration attributes.
type RfidConfig struct {
	// RfidEnableBluesoft represents whether bluesoft tag supported or not.
	// If enable state then bluesoft tag is supported and if disable state then bluesoft tag not supported.
	RfidEnableBluesoft *bool `json:"rfid-enable-bluesoft,omitempty"` // Bluesoft tag support flag (YANG: IOS-XE 17.12.1+)

	// RfidTimeout represents the timeout value to cleanup stale rfid entries.
	// Range: 60..7200, Default: 1200
	RfidTimeout *uint16 `json:"rfid-timeout,omitempty"` // Stale entry cleanup timeout (YANG: IOS-XE 17.12.1+)

	// RfidEnableCisco represents whether Cisco tag rfid supported or not.
	// If enable state then Cisco tag is supported and if disable state then Cisco tag not supported.
	RfidEnableCisco *bool `json:"rfid-enable-cisco,omitempty"` // Cisco tag support flag (YANG: IOS-XE 17.12.1+)

	// RfidRssiExpiry represents to cleanup rfid rssi when rssi value expires.
	// Range: 5..300, Default: 5
	RfidRssiExpiry *uint16 `json:"rfid-rssi-expiry,omitempty"` // RSSI value expiry timeout (YANG: IOS-XE 17.12.1+)

	// RfidRssiHalflife represents half life when averaging two RSSI readings for RFID tags,
	// timeout value zero indicates that timeout is disabled.
	// Range: 0,1,2,5,10,20,30,60,90,120,180,300, Default: 0
	RfidRssiHalflife *uint32 `json:"rfid-rssi-halflife,omitempty"` // RSSI averaging half-life duration (YANG: IOS-XE 17.12.1+)

	// RfidNotifyThreshold indicates LOCP notification threshold for RSSI measurements.
	// Range: 1..10, Default: 5
	RfidNotifyThreshold *uint32 `json:"rfid-notify-threshold,omitempty"` // LOCP notification threshold (YANG: IOS-XE 17.12.1+)
}
