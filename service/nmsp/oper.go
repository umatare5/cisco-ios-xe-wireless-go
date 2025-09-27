package nmsp

// CiscoIOSXEWirelessNMSPOper represents the NMSP operational data container.
type CiscoIOSXEWirelessNMSPOper struct {
	CiscoIOSXEWirelessNMSPOperData struct {
		ClientRegistration []ClientRegistration `json:"client-registration"` // NMSP client applications providing services (Live: IOS-XE 17.12.6a)
		CmxConnection      []CmxConnection      `json:"cmx-connection"`      // CMX connection table (Live: IOS-XE 17.12.6a)
		CmxCloudInfo       CmxCloudInfo         `json:"cmx-cloud-info"`      // NMSP services over HTTPS transport info (Live: IOS-XE 17.12.6a)
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
	ClientID int          `json:"client-id"` // NMSP client identifier (Live: IOS-XE 17.12.6a)
	Services NMSPServices `json:"services"`  // NMSP capabilities provided by EWLC app (Live: IOS-XE 17.12.6a)
}

// NMSPServices represents NMSP service configuration and status.
type NMSPServices struct {
	Mask                        string     `json:"mask"`                                    // NMSP subservice bitmask combination (Live: IOS-XE 17.12.6a)
	RSSIMs                      *EmptyType `json:"rssi-ms,omitempty"`                       // RSSI Mobile Station capability (Live: IOS-XE 17.12.6a)
	RSSIRFID                    *EmptyType `json:"rssi-rfid,omitempty"`                     // RSSI RFID Tag capability (Live: IOS-XE 17.12.6a)
	RSSIRogue                   *EmptyType `json:"rssi-rogue,omitempty"`                    // RSSI Rogue capability (Live: IOS-XE 17.12.6a)
	RSSIInterferer              *EmptyType `json:"rssi-interferer,omitempty"`               // RSSI Interferer capability (YANG: IOS-XE 17.12.1)
	RSSIHs                      *EmptyType `json:"rssi-hs,omitempty"`                       // RSSI Handover Client capability (YANG: IOS-XE 17.12.1)
	RSSIMsAssociatedOnly        *EmptyType `json:"rssi-ms-associated-only,omitempty"`       // RSSI Associated Mobile Station only capability (Live: IOS-XE 17.12.6a)
	SpectrumInterferer          *EmptyType `json:"spectrum-interferer,omitempty"`           // Spectrum Interferer capability (Live: IOS-XE 17.12.6a)
	SpectrumAirQuality          *EmptyType `json:"spectrum-air-quality,omitempty"`          // Spectrum Air Quality capability (Live: IOS-XE 17.12.6a)
	SpectrumAggregateInterferer *EmptyType `json:"spectrum-aggregate-interferer,omitempty"` // Spectrum Aggregate Interferer capability (Live: IOS-XE 17.12.6a)
	InfoMs                      *EmptyType `json:"info-ms,omitempty"`                       // Information Mobile Station capability (Live: IOS-XE 17.12.6a)
	InfoRFID                    *EmptyType `json:"info-rfid,omitempty"`                     // Information RFID Tag capability (YANG: IOS-XE 17.12.1)
	InfoRogue                   *EmptyType `json:"info-rogue,omitempty"`                    // Information Rogue capability (Live: IOS-XE 17.12.6a)
	InfoHs                      *EmptyType `json:"info-hs,omitempty"`                       // Information Handover Client capability (YANG: IOS-XE 17.12.1)
	StatsMs                     *EmptyType `json:"stats-ms,omitempty"`                      // Statistics Mobile Station capability (Live: IOS-XE 17.12.6a)
	StatsRFID                   *EmptyType `json:"stats-rfid,omitempty"`                    // Statistics RFID Tag capability (Live: IOS-XE 17.12.6a)
	StatsRogue                  *EmptyType `json:"stats-rogue,omitempty"`                   // Statistics Rogue capability (Live: IOS-XE 17.12.6a)
	Attach                      *EmptyType `json:"attach,omitempty"`                        // Mobile Station Attachment capability (YANG: IOS-XE 17.12.1)
	Location                    *EmptyType `json:"location,omitempty"`                      // Location Service capability (YANG: IOS-XE 17.12.1)
	Fmchs                       *EmptyType `json:"fmchs,omitempty"`                         // Handover Service capability (YANG: IOS-XE 17.12.1)
	ApMonitor                   *EmptyType `json:"ap-monitor,omitempty"`                    // AP Monitor Service capability (Live: IOS-XE 17.12.6a)
	Wips                        *EmptyType `json:"wips,omitempty"`                          // Wireless Intrusion Detection System capability (YANG: IOS-XE 17.12.1)
	OnDemand                    *EmptyType `json:"on-demand,omitempty"`                     // On-Demand Service capability (Live: IOS-XE 17.12.6a)
	ApInfo                      *EmptyType `json:"ap-info,omitempty"`                       // AP Info Service capability (Live: IOS-XE 17.12.6a)
}

// CmxConnection represents individual CMX connection details.
type CmxConnection struct {
	PeerIP        string      `json:"peer-ip"`       // IP address of the CMX (Live: IOS-XE 17.12.6a)
	ConnectionID  string      `json:"connection-id"` // Internal connection ID of the CMX (Live: IOS-XE 17.12.6a)
	Active        bool        `json:"active"`        // Flag indicating whether CMX connection is active (Live: IOS-XE 17.12.6a)
	ConStats      CmxConStats `json:"con-stats"`     // Statistics of messages sent/received (Live: IOS-XE 17.12.6a)
	Subscriptions struct {
		Mask string `json:"mask"` // NMSP subservice bitmask combination (Live: IOS-XE 17.12.6a)
	} `json:"subscriptions"` // Service subscriptions established by given CMX (Live: IOS-XE 17.12.6a)
	Transport string `json:"transport"` // The transport used for this CMX connection (Live: IOS-XE 17.12.6a)
}

// CmxConStats represents CMX connection statistics.
type CmxConStats struct {
	TxMsgCounter        []MsgCounter `json:"tx-msg-counter"`        // Messages transmitted from NMSPd to CMX (Live: IOS-XE 17.12.6a)
	RxMsgCounter        []MsgCounter `json:"rx-msg-counter"`        // Messages received by NMSPd from CMX (Live: IOS-XE 17.12.6a)
	UnsupportedMsgCount string       `json:"unsupported-msg-count"` // Unsupported messages received by NMSP daemon (Live: IOS-XE 17.12.6a)
	TxDataFrames        string       `json:"tx-data-frames"`        // Data frames transferred from NMSP daemon to CMX (Live: IOS-XE 17.12.6a)
	RxDataFrames        string       `json:"rx-data-frames"`        // Data frames received by NMSP daemon from CMX (Live: IOS-XE 17.12.6a)
	Connections         string       `json:"connections"`           // Amount of successful connections (Live: IOS-XE 17.12.6a)
	Disconnections      string       `json:"disconnections"`        // Amount of disconnections (Live: IOS-XE 17.12.6a)
}

// MsgCounter represents message counter statistics.
type MsgCounter struct {
	Counter string `json:"counter"` // Amount of messages sent or received (Live: IOS-XE 17.12.6a)
	MsgID   int    `json:"msg-id"`  // NMSP protocol message identifier (Live: IOS-XE 17.12.6a)
}

// CmxCloudInfo represents CMX cloud information container.
type CmxCloudInfo struct {
	CloudStatus CloudStatus `json:"cloud-status"` // Status of the cloud connection (Live: IOS-XE 17.12.6a)
	CloudStats  CloudStats  `json:"cloud-stats"`  // Statistics of the cloud connection (Live: IOS-XE 17.12.6a)
}

// CloudStatus represents CMX cloud connection status.
type CloudStatus struct {
	IPAddress         string `json:"ip-address"`          // IP Address which CMX cloud server is resolved (Live: IOS-XE 17.12.6a)
	Connectivity      string `json:"connectivity"`        // Enum representing status UP/DOWN of HTTP connection (Live: IOS-XE 17.12.6a)
	ServiceUp         bool   `json:"service-up"`          // True if NMSP connectivity towards CMX cloud is UP (Live: IOS-XE 17.12.6a)
	LastRequestStatus string `json:"last-request-status"` // String representing the last request status (Live: IOS-XE 17.12.6a)
	HeartbeatStatusOk bool   `json:"heartbeat-status-ok"` // True if last heartbeat was successful (Live: IOS-XE 17.12.6a)
}

// CloudStats represents CMX cloud connection statistics.
type CloudStats struct {
	TxDataframes     int `json:"tx-dataframes"`     // Number of data frames sent (Live: IOS-XE 17.12.6a)
	RxDataframes     int `json:"rx-dataframes"`     // Number of data frames received (Live: IOS-XE 17.12.6a)
	TxHeartbeatReq   int `json:"tx-heartbeat-req"`  // Number of HTTP heartbeats sent (Live: IOS-XE 17.12.6a)
	HeartbeatTimeout int `json:"heartbeat-timeout"` // Number of heartbeat timeouts (Live: IOS-XE 17.12.6a)
	RxSubscriberReq  int `json:"rx-subscriber-req"` // Number of requests from subscriber (Live: IOS-XE 17.12.6a)
	TxDatabytes      int `json:"tx-databytes"`      // Number of bytes sent (Live: IOS-XE 17.12.6a)
	RxDatabytes      int `json:"rx-databytes"`      // Number of bytes received (Live: IOS-XE 17.12.6a)
	TxHeartbeatFail  int `json:"tx-heartbeat-fail"` // Count of failures sending heartbeat (Live: IOS-XE 17.12.6a)
	RxDataFail       int `json:"rx-data-fail"`      // Number of failures receiving data (Live: IOS-XE 17.12.6a)
	TxDataFail       int `json:"tx-data-fail"`      // Number of failures sending data (Live: IOS-XE 17.12.6a)
}

// EmptyType represents YANG empty type fields appearing as null arrays in RESTCONF JSON.
type EmptyType []interface{} // appears as [null]
