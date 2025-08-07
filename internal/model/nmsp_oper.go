// Package model contains generated response structures for the Cisco WNC API.
// This package is part of the three-layer architecture providing Generated Type separation.
package model

// NMSP Operational Response Types

// NmspOperResponse represents the response structure for NMSP operational data.
type NmspOperResponse struct {
	CiscoIOSXEWirelessNmspOperData struct {
		ClientRegistration []ClientRegistration `json:"client-registration"`
		CmxConnection      []CmxConnection      `json:"cmx-connection"`
		CmxCloudInfo       CmxCloudInfo         `json:"cmx-cloud-info"`
	} `json:"Cisco-IOS-XE-wireless-nmsp-oper:nmsp-oper-data"`
}

// NmspClientRegistrationResponse represents the response structure for NMSP client registration data.
type NmspClientRegistrationResponse struct {
	ClientRegistration []ClientRegistration `json:"Cisco-IOS-XE-wireless-nmsp-oper:client-registration"`
}

// NmspCmxConnectionResponse represents the response structure for NMSP CMX connection data.
type NmspCmxConnectionResponse struct {
	CmxConnection []CmxConnection `json:"Cisco-IOS-XE-wireless-nmsp-oper:cmx-connection"`
}

// NmspCmxCloudInfoResponse represents the response structure for NMSP CMX cloud information.
type NmspCmxCloudInfoResponse struct {
	CmxCloudInfo CmxCloudInfo `json:"Cisco-IOS-XE-wireless-nmsp-oper:cmx-cloud-info"`
}

// NMSP Supporting Types

// ClientRegistration represents NMSP client registration information.
type ClientRegistration struct {
	ClientID int          `json:"client-id"`
	Services NmspServices `json:"services"`
}

// NmspServices represents NMSP services configuration including various service types.
type NmspServices struct {
	Mask                        string `json:"mask"`
	RssiMs                      []any  `json:"rssi-ms,omitempty"`
	RssiRfid                    []any  `json:"rssi-rfid,omitempty"`
	RssiRogue                   []any  `json:"rssi-rogue,omitempty"`
	RssiMsAssociatedOnly        []any  `json:"rssi-ms-associated-only,omitempty"`
	SpectrumInterferer          []any  `json:"spectrum-interferer,omitempty"`
	SpectrumAirQuality          []any  `json:"spectrum-air-quality,omitempty"`
	SpectrumAggregateInterferer []any  `json:"spectrum-aggregate-interferer,omitempty"`
	InfoMs                      []any  `json:"info-ms,omitempty"`
	InfoRogue                   []any  `json:"info-rogue,omitempty"`
	StatsMs                     []any  `json:"stats-ms,omitempty"`
	StatsRfid                   []any  `json:"stats-rfid,omitempty"`
	StatsRogue                  []any  `json:"stats-rogue,omitempty"`
	ApMonitor                   []any  `json:"ap-monitor,omitempty"`
	OnDemand                    []any  `json:"on-demand,omitempty"`
	ApInfo                      []any  `json:"ap-info,omitempty"`
}

// CmxConnection represents a CMX connection entry with connection details and statistics.
type CmxConnection struct {
	PeerIP        string      `json:"peer-ip"`
	ConnectionID  string      `json:"connection-id"`
	Active        bool        `json:"active"`
	ConStats      CmxConStats `json:"con-stats"`
	Subscriptions struct {
		Mask string `json:"mask"`
	} `json:"subscriptions"`
	Transport string `json:"transport"`
}

// CmxConStats represents CMX connection statistics including message counters and frame counts.
type CmxConStats struct {
	TxMsgCounter        []MsgCounter `json:"tx-msg-counter"`
	RxMsgCounter        []MsgCounter `json:"rx-msg-counter"`
	UnsupportedMsgCount string       `json:"unsupported-msg-count"`
	TxDataFrames        string       `json:"tx-data-frames"`
	RxDataFrames        string       `json:"rx-data-frames"`
	Connections         string       `json:"connections"`
	Disconnections      string       `json:"disconnections"`
}

// MsgCounter represents a message counter with message ID and count.
type MsgCounter struct {
	Counter string `json:"counter"`
	MsgID   int    `json:"msg-id"`
}

// CmxCloudInfo represents CMX cloud information including status and statistics.
type CmxCloudInfo struct {
	CloudStatus CloudStatus `json:"cloud-status"`
	CloudStats  CloudStats  `json:"cloud-stats"`
}

// CloudStatus represents CMX cloud connectivity status.
type CloudStatus struct {
	IPAddress         string `json:"ip-address"`
	Connectivity      string `json:"connectivity"`
	ServiceUp         bool   `json:"service-up"`
	LastRequestStatus string `json:"last-request-status"`
	HeartbeatStatusOk bool   `json:"heartbeat-status-ok"`
}

// CloudStats represents CMX cloud communication statistics.
type CloudStats struct {
	TxDataframes     int `json:"tx-dataframes"`
	RxDataframes     int `json:"rx-dataframes"`
	TxHeartbeatReq   int `json:"tx-heartbeat-req"`
	HeartbeatTimeout int `json:"heartbeat-timeout"`
	RxSubscriberReq  int `json:"rx-subscriber-req"`
	TxDatabytes      int `json:"tx-databytes"`
	RxDatabytes      int `json:"rx-databytes"`
	TxHeartbeatFail  int `json:"tx-heartbeat-fail"`
	RxDataFail       int `json:"rx-data-fail"`
	TxDataFail       int `json:"tx-data-fail"`
}
