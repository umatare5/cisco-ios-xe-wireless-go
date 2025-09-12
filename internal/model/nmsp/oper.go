package nmsp

// NmspOper represents the NMSP operational data container.
type NmspOper struct {
	CiscoIOSXEWirelessNmspOperData struct {
		ClientRegistration []ClientRegistration `json:"client-registration"`
		CmxConnection      []CmxConnection      `json:"cmx-connection"`
		CmxCloudInfo       CmxCloudInfo         `json:"cmx-cloud-info"`
	} `json:"Cisco-IOS-XE-wireless-nmsp-oper:nmsp-oper-data"`
}

// NmspClientRegistration represents the NMSP client registration data container.
type NmspClientRegistration struct {
	ClientRegistration []ClientRegistration `json:"Cisco-IOS-XE-wireless-nmsp-oper:client-registration"`
}

// NmspCmxConnection represents the NMSP CMX connection data container.
type NmspCmxConnection struct {
	CmxConnection []CmxConnection `json:"Cisco-IOS-XE-wireless-nmsp-oper:cmx-connection"`
}

// NmspCmxCloudInfo represents the NMSP CMX cloud information data container.
type NmspCmxCloudInfo struct {
	CmxCloudInfo CmxCloudInfo `json:"Cisco-IOS-XE-wireless-nmsp-oper:cmx-cloud-info"`
}

// NMSP Supporting Types

// EmptyType represents YANG empty type fields appearing as null arrays in RESTCONF JSON.
type EmptyType []interface{}

// ClientRegistration represents individual NMSP client registration record.
type ClientRegistration struct {
	ClientID int          `json:"client-id"` // NMSP client identifier
	Services NmspServices `json:"services"`  // NMSP services configuration
}

// NmspServices represents NMSP service configuration and status.
type NmspServices struct {
	Mask                        string     `json:"mask"`                                    // NMSP services bitmask
	RssiMs                      *EmptyType `json:"rssi-ms,omitempty"`                       // NMSP mobile station RSSI service
	RssiRfid                    *EmptyType `json:"rssi-rfid,omitempty"`                     // NMSP RFID RSSI service
	RssiRogue                   *EmptyType `json:"rssi-rogue,omitempty"`                    // NMSP rogue RSSI service
	RssiInterferer              *EmptyType `json:"rssi-interferer,omitempty"`               // NMSP RSSI interferer service (YANG: IOS-XE 17.12.1+)
	RssiHs                      *EmptyType `json:"rssi-hs,omitempty"`                       // NMSP RSSI hotspot service (YANG: IOS-XE 17.12.1+)
	RssiMsAssociatedOnly        *EmptyType `json:"rssi-ms-associated-only,omitempty"`       // NMSP associated mobile station RSSI service
	SpectrumInterferer          *EmptyType `json:"spectrum-interferer,omitempty"`           // NMSP spectrum interferer service
	SpectrumAirQuality          *EmptyType `json:"spectrum-air-quality,omitempty"`          // NMSP spectrum air quality service
	SpectrumAggregateInterferer *EmptyType `json:"spectrum-aggregate-interferer,omitempty"` // NMSP spectrum aggregate interferer service
	InfoMs                      *EmptyType `json:"info-ms,omitempty"`                       // NMSP mobile station information service
	InfoRfid                    *EmptyType `json:"info-rfid,omitempty"`                     // NMSP RFID information service (YANG: IOS-XE 17.12.1+)
	InfoRogue                   *EmptyType `json:"info-rogue,omitempty"`                    // NMSP rogue information service
	InfoHs                      *EmptyType `json:"info-hs,omitempty"`                       // NMSP hotspot information service (YANG: IOS-XE 17.12.1+)
	StatsMs                     *EmptyType `json:"stats-ms,omitempty"`                      // NMSP mobile station statistics service
	StatsRfid                   *EmptyType `json:"stats-rfid,omitempty"`                    // NMSP RFID statistics service
	StatsRogue                  *EmptyType `json:"stats-rogue,omitempty"`                   // NMSP rogue statistics service
	Attach                      *EmptyType `json:"attach,omitempty"`                        // NMSP attach notification service (YANG: IOS-XE 17.12.1+)
	Location                    *EmptyType `json:"location,omitempty"`                      // NMSP location service (YANG: IOS-XE 17.12.1+)
	Fmchs                       *EmptyType `json:"fmchs,omitempty"`                         // NMSP FMCHS service (YANG: IOS-XE 17.12.1+)
	ApMonitor                   *EmptyType `json:"ap-monitor,omitempty"`                    // NMSP access point monitor service
	Wips                        *EmptyType `json:"wips,omitempty"`                          // NMSP WIPS service (YANG: IOS-XE 17.12.1+)
	OnDemand                    *EmptyType `json:"on-demand,omitempty"`                     // NMSP on-demand service
	ApInfo                      *EmptyType `json:"ap-info,omitempty"`                       // NMSP access point information service
}

// CmxConnection represents individual CMX connection details.
type CmxConnection struct {
	PeerIP        string      `json:"peer-ip"`       // CMX peer IP address
	ConnectionID  string      `json:"connection-id"` // CMX connection identifier
	Active        bool        `json:"active"`        // CMX connection active status
	ConStats      CmxConStats `json:"con-stats"`     // CMX connection statistics
	Subscriptions struct {
		Mask string `json:"mask"` // CMX subscription mask
	} `json:"subscriptions"` // CMX subscription configuration
	Transport string `json:"transport"` // CMX transport protocol
}

// CmxConStats represents CMX connection statistics.
type CmxConStats struct {
	TxMsgCounter        []MsgCounter `json:"tx-msg-counter"`        // Transmitted message counters
	RxMsgCounter        []MsgCounter `json:"rx-msg-counter"`        // Received message counters
	UnsupportedMsgCount string       `json:"unsupported-msg-count"` // Unsupported message count
	TxDataFrames        string       `json:"tx-data-frames"`        // Transmitted data frames count
	RxDataFrames        string       `json:"rx-data-frames"`        // Received data frames count
	Connections         string       `json:"connections"`           // Total connections count
	Disconnections      string       `json:"disconnections"`        // Total disconnections count
}

// MsgCounter represents message counter statistics.
type MsgCounter struct {
	Counter string `json:"counter"` // Message counter value
	MsgID   int    `json:"msg-id"`  // Message identifier
}

// CmxCloudInfo represents CMX cloud information container.
type CmxCloudInfo struct {
	CloudStatus CloudStatus `json:"cloud-status"` // CMX cloud connection status
	CloudStats  CloudStats  `json:"cloud-stats"`  // CMX cloud connection statistics
}

// CloudStatus represents CMX cloud connection status.
type CloudStatus struct {
	IPAddress         string `json:"ip-address"`          // CMX cloud server IP address
	Connectivity      string `json:"connectivity"`        // CMX cloud connectivity status
	ServiceUp         bool   `json:"service-up"`          // CMX cloud service availability
	LastRequestStatus string `json:"last-request-status"` // Last request status
	HeartbeatStatusOk bool   `json:"heartbeat-status-ok"` // Heartbeat status indicator
}

// CloudStats represents CMX cloud connection statistics.
type CloudStats struct {
	TxDataframes     int `json:"tx-dataframes"`     // Transmitted data frames to cloud
	RxDataframes     int `json:"rx-dataframes"`     // Received data frames from cloud
	TxHeartbeatReq   int `json:"tx-heartbeat-req"`  // Transmitted heartbeat requests
	HeartbeatTimeout int `json:"heartbeat-timeout"` // Heartbeat timeout count
	RxSubscriberReq  int `json:"rx-subscriber-req"` // Received subscriber requests
	TxDatabytes      int `json:"tx-databytes"`      // Transmitted data bytes
	RxDatabytes      int `json:"rx-databytes"`      // Received data bytes
	TxHeartbeatFail  int `json:"tx-heartbeat-fail"` // Failed heartbeat transmissions
	RxDataFail       int `json:"rx-data-fail"`      // Failed data receptions
	TxDataFail       int `json:"tx-data-fail"`      // Failed data transmissions
}
