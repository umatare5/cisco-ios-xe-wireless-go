package rfid

// CiscoIOSXEWirelessRFIDCfg represents RFID configuration data structure.
type CiscoIOSXEWirelessRFIDCfg struct {
	CiscoIOSXEWirelessRFIDCfgData *CiscoIOSXEWirelessRFIDCfgData `json:"Cisco-IOS-XE-wireless-rfid-cfg:rfid-cfg-data"` // RFID configuration data for threshold timer (YANG: IOS-XE 17.12.1)
}

// CiscoIOSXEWirelessRFIDCfgData represents RFID configuration data for threshold timer (YANG: IOS-XE 17.12.1).
type CiscoIOSXEWirelessRFIDCfgData struct {
	RFID *RFIDConfig `json:"rfid,omitempty"` // RFID config attributes container (YANG: IOS-XE 17.12.1)
}

// CiscoIOSXEWirelessRFIDCfgRFID represents the RFID config attributes container.
type CiscoIOSXEWirelessRFIDCfgRFID struct {
	RFID *RFIDConfig `json:"Cisco-IOS-XE-wireless-rfid-cfg:rfid,omitempty"` // RFID config attributes container (YANG: IOS-XE 17.12.1)
}

// RFIDConfig represents RFID configuration attributes.
type RFIDConfig struct {
	RFIDEnableBluesoft  *bool   `json:"rfid-enable-bluesoft,omitempty"`  // Bluesoft tag support flag (YANG: IOS-XE 17.12.1)
	RFIDTimeout         *uint16 `json:"rfid-timeout,omitempty"`          // Stale RFID entry cleanup timeout value (YANG: IOS-XE 17.12.1)
	RFIDEnableCisco     *bool   `json:"rfid-enable-cisco,omitempty"`     // Cisco tag RFID support flag (YANG: IOS-XE 17.12.1)
	RFIDRssiExpiry      *uint16 `json:"rfid-rssi-expiry,omitempty"`      // RFID RSSI value cleanup expiry timeout (YANG: IOS-XE 17.12.1)
	RFIDRssiHalflife    *uint32 `json:"rfid-rssi-halflife,omitempty"`    // RSSI averaging half-life for RFID tags (YANG: IOS-XE 17.12.1)
	RFIDNotifyThreshold *uint32 `json:"rfid-notify-threshold,omitempty"` // LOCP notification threshold for RSSI (YANG: IOS-XE 17.12.1)
}
