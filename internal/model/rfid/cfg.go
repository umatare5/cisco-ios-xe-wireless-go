package rfid

// RfidCfg represents RFID configuration data structure.
type RfidCfg struct {
	RfidCfgData struct { // RFID configuration data for threshold timer (YANG: IOS-XE 17.12.1)
		Rfid *RfidConfig `json:"rfid,omitempty"` // RFID config attributes container (YANG: IOS-XE 17.12.1)
	} `json:"Cisco-IOS-XE-wireless-rfid-cfg:rfid-cfg-data"`
}

// RfidConfig represents RFID configuration attributes.
type RfidConfig struct {
	RfidEnableBluesoft  *bool   `json:"rfid-enable-bluesoft,omitempty"`  // Bluesoft tag support flag (YANG: IOS-XE 17.12.1)
	RfidTimeout         *uint16 `json:"rfid-timeout,omitempty"`          // Stale RFID entry cleanup timeout value (YANG: IOS-XE 17.12.1)
	RfidEnableCisco     *bool   `json:"rfid-enable-cisco,omitempty"`     // Cisco tag RFID support flag (YANG: IOS-XE 17.12.1)
	RfidRssiExpiry      *uint16 `json:"rfid-rssi-expiry,omitempty"`      // RFID RSSI value cleanup expiry timeout (YANG: IOS-XE 17.12.1)
	RfidRssiHalflife    *uint32 `json:"rfid-rssi-halflife,omitempty"`    // RSSI averaging half-life for RFID tags (YANG: IOS-XE 17.12.1)
	RfidNotifyThreshold *uint32 `json:"rfid-notify-threshold,omitempty"` // LOCP notification threshold for RSSI (YANG: IOS-XE 17.12.1)
}
