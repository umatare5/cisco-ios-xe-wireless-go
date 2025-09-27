package nmsp

// CiscoIOSXEWirelessNMSPOper represents the NMSP operational data container.
type CiscoIOSXEWirelessNMSPOper struct {
	CiscoIOSXEWirelessNMSPOperData struct {
		ClientRegistration []ClientRegistration `json:"client-registration"` // NMSP client applications providing services (Live: IOS-XE 17.12.5)
		CmxConnection      []CmxConnection      `json:"cmx-connection"`      // CMX connection table (Live: IOS-XE 17.12.5)
		CmxCloudInfo       CmxCloudInfo         `json:"cmx-cloud-info"`      // NMSP services over HTTPS transport info (Live: IOS-XE 17.12.5)
	} `json:"Cisco-IOS-XE-wireless-nmsp-oper:nmsp-oper-data"` // Network Mobility Services Protocol (YANG: IOS-XE 17.12.1)
}

// CiscoIOSXEWirelessNMSPClientRegistration represents the NMSP client registration data container.
type CiscoIOSXEWirelessNMSPClientRegistration struct {
	ClientRegistration []ClientRegistration `json:"Cisco-IOS-XE-wireless-nmsp-oper:client-registration"`
}

// CiscoIOSXEWirelessNMSPCmxConnection represents the NMSP CMX connection data container.
type CiscoIOSXEWirelessNMSPCmxConnection struct {
	CmxConnection []CmxConnection `json:"Cisco-IOS-XE-wireless-nmsp-oper:cmx-connection"`
}

// CiscoIOSXEWirelessNMSPCmxCloudInfo represents the NMSP CMX cloud information data container.
type CiscoIOSXEWirelessNMSPCmxCloudInfo struct {
	CmxCloudInfo CmxCloudInfo `json:"Cisco-IOS-XE-wireless-nmsp-oper:cmx-cloud-info"`
}

// ClientRegistration represents individual NMSP client registration record.
type ClientRegistration struct {
	ClientID int          `json:"client-id"` // NMSP client identifier (Live: IOS-XE 17.12.5)
	Services NMSPServices `json:"services"`  // NMSP capabilities provided by EWLC app (Live: IOS-XE 17.12.5)
}

// NMSPServices represents NMSP service configuration and status.
type NMSPServices struct {
	Mask                        string     `json:"mask"`                                    // NMSP subservice bitmask combination (Live: IOS-XE 17.12.5)
	RSSIMs                      *EmptyType `json:"rssi-ms,omitempty"`                       // RSSI Mobile Station capability (Live: IOS-XE 17.12.5)
	RSSIRFID                    *EmptyType `json:"rssi-rfid,omitempty"`                     // RSSI RFID Tag capability (Live: IOS-XE 17.12.5)
	RSSIRogue                   *EmptyType `json:"rssi-rogue,omitempty"`                    // RSSI Rogue capability (Live: IOS-XE 17.12.5)
	RSSIInterferer              *EmptyType `json:"rssi-interferer,omitempty"`               // RSSI Interferer capability (YANG: IOS-XE 17.12.1)
	RSSIHs                      *EmptyType `json:"rssi-hs,omitempty"`                       // RSSI Handover Client capability (YANG: IOS-XE 17.12.1)
	RSSIMsAssociatedOnly        *EmptyType `json:"rssi-ms-associated-only,omitempty"`       // RSSI Associated Mobile Station only capability (Live: IOS-XE 17.12.5)
	SpectrumInterferer          *EmptyType `json:"spectrum-interferer,omitempty"`           // Spectrum Interferer capability (Live: IOS-XE 17.12.5)
	SpectrumAirQuality          *EmptyType `json:"spectrum-air-quality,omitempty"`          // Spectrum Air Quality capability (Live: IOS-XE 17.12.5)
	SpectrumAggregateInterferer *EmptyType `json:"spectrum-aggregate-interferer,omitempty"` // Spectrum Aggregate Interferer capability (Live: IOS-XE 17.12.5)
	InfoMs                      *EmptyType `json:"info-ms,omitempty"`                       // Information Mobile Station capability (Live: IOS-XE 17.12.5)
	InfoRFID                    *EmptyType `json:"info-rfid,omitempty"`                     // Information RFID Tag capability (YANG: IOS-XE 17.12.1)
	InfoRogue                   *EmptyType `json:"info-rogue,omitempty"`                    // Information Rogue capability (Live: IOS-XE 17.12.5)
	InfoHs                      *EmptyType `json:"info-hs,omitempty"`                       // Information Handover Client capability (YANG: IOS-XE 17.12.1)
	StatsMs                     *EmptyType `json:"stats-ms,omitempty"`                      // Statistics Mobile Station capability (Live: IOS-XE 17.12.5)
	StatsRFID                   *EmptyType `json:"stats-rfid,omitempty"`                    // Statistics RFID Tag capability (Live: IOS-XE 17.12.5)
	StatsRogue                  *EmptyType `json:"stats-rogue,omitempty"`                   // Statistics Rogue capability (Live: IOS-XE 17.12.5)
	Attach                      *EmptyType `json:"attach,omitempty"`                        // Mobile Station Attachment capability (YANG: IOS-XE 17.12.1)
	Location                    *EmptyType `json:"location,omitempty"`                      // Location Service capability (YANG: IOS-XE 17.12.1)
	Fmchs                       *EmptyType `json:"fmchs,omitempty"`                         // Handover Service capability (YANG: IOS-XE 17.12.1)
	ApMonitor                   *EmptyType `json:"ap-monitor,omitempty"`                    // AP Monitor Service capability (Live: IOS-XE 17.12.5)
	Wips                        *EmptyType `json:"wips,omitempty"`                          // Wireless Intrusion Detection System capability (YANG: IOS-XE 17.12.1)
	OnDemand                    *EmptyType `json:"on-demand,omitempty"`                     // On-Demand Service capability (Live: IOS-XE 17.12.5)
	ApInfo                      *EmptyType `json:"ap-info,omitempty"`                       // AP Info Service capability (Live: IOS-XE 17.12.5)
}

// CmxConnection represents individual CMX connection details.
type CmxConnection struct {
	PeerIP        string      `json:"peer-ip"`       // IP address of the CMX (Live: IOS-XE 17.12.5)
	ConnectionID  string      `json:"connection-id"` // Internal connection ID of the CMX (Live: IOS-XE 17.12.5)
	Active        bool        `json:"active"`        // Flag indicating whether CMX connection is active (Live: IOS-XE 17.12.5)
	ConStats      CmxConStats `json:"con-stats"`     // Statistics of messages sent/received (Live: IOS-XE 17.12.5)
	Subscriptions struct {
		Mask string `json:"mask"` // NMSP subservice bitmask combination (Live: IOS-XE 17.12.5)
	} `json:"subscriptions"` // Service subscriptions established by given CMX (Live: IOS-XE 17.12.5)
	Transport string `json:"transport"` // The transport used for this CMX connection (Live: IOS-XE 17.12.5)
}

// CmxConStats represents CMX connection statistics.
type CmxConStats struct {
	TxMsgCounter        []MsgCounter `json:"tx-msg-counter"`        // Messages transmitted from NMSPd to CMX (Live: IOS-XE 17.12.5)
	RxMsgCounter        []MsgCounter `json:"rx-msg-counter"`        // Messages received by NMSPd from CMX (Live: IOS-XE 17.12.5)
	UnsupportedMsgCount string       `json:"unsupported-msg-count"` // Unsupported messages received by NMSP daemon (Live: IOS-XE 17.12.5)
	TxDataFrames        string       `json:"tx-data-frames"`        // Data frames transferred from NMSP daemon to CMX (Live: IOS-XE 17.12.5)
	RxDataFrames        string       `json:"rx-data-frames"`        // Data frames received by NMSP daemon from CMX (Live: IOS-XE 17.12.5)
	Connections         string       `json:"connections"`           // Amount of successful connections (Live: IOS-XE 17.12.5)
	Disconnections      string       `json:"disconnections"`        // Amount of disconnections (Live: IOS-XE 17.12.5)
}

// MsgCounter represents message counter statistics.
type MsgCounter struct {
	Counter string `json:"counter"` // Amount of messages sent or received (Live: IOS-XE 17.12.5)
	MsgID   int    `json:"msg-id"`  // NMSP protocol message identifier (Live: IOS-XE 17.12.5)
}

// CmxCloudInfo represents CMX cloud information container.
type CmxCloudInfo struct {
	CloudStatus CloudStatus `json:"cloud-status"` // Status of the cloud connection (Live: IOS-XE 17.12.5)
	CloudStats  CloudStats  `json:"cloud-stats"`  // Statistics of the cloud connection (Live: IOS-XE 17.12.5)
}

// CloudStatus represents CMX cloud connection status.
type CloudStatus struct {
	IPAddress         string `json:"ip-address"`          // IP Address which CMX cloud server is resolved (Live: IOS-XE 17.12.5)
	Connectivity      string `json:"connectivity"`        // Enum representing status UP/DOWN of HTTP connection (Live: IOS-XE 17.12.5)
	ServiceUp         bool   `json:"service-up"`          // True if NMSP connectivity towards CMX cloud is UP (Live: IOS-XE 17.12.5)
	LastRequestStatus string `json:"last-request-status"` // String representing the last request status (Live: IOS-XE 17.12.5)
	HeartbeatStatusOk bool   `json:"heartbeat-status-ok"` // True if last heartbeat was successful (Live: IOS-XE 17.12.5)
}

// CloudStats represents CMX cloud connection statistics.
type CloudStats struct {
	TxDataframes     int `json:"tx-dataframes"`     // Number of data frames sent (Live: IOS-XE 17.12.5)
	RxDataframes     int `json:"rx-dataframes"`     // Number of data frames received (Live: IOS-XE 17.12.5)
	TxHeartbeatReq   int `json:"tx-heartbeat-req"`  // Number of HTTP heartbeats sent (Live: IOS-XE 17.12.5)
	HeartbeatTimeout int `json:"heartbeat-timeout"` // Number of heartbeat timeouts (Live: IOS-XE 17.12.5)
	RxSubscriberReq  int `json:"rx-subscriber-req"` // Number of requests from subscriber (Live: IOS-XE 17.12.5)
	TxDatabytes      int `json:"tx-databytes"`      // Number of bytes sent (Live: IOS-XE 17.12.5)
	RxDatabytes      int `json:"rx-databytes"`      // Number of bytes received (Live: IOS-XE 17.12.5)
	TxHeartbeatFail  int `json:"tx-heartbeat-fail"` // Count of failures sending heartbeat (Live: IOS-XE 17.12.5)
	RxDataFail       int `json:"rx-data-fail"`      // Number of failures receiving data (Live: IOS-XE 17.12.5)
	TxDataFail       int `json:"tx-data-fail"`      // Number of failures sending data (Live: IOS-XE 17.12.5)
}

// EmptyType represents YANG empty type fields appearing as null arrays in RESTCONF JSON.
type EmptyType []interface{} // appears as [null]
