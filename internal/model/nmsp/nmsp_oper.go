// Package model provides type definitions for Cisco IOS-XE wireless controller operations.
package model

// NMSP Operational Response Types

// NmspOper  represents the NMSP operational data.
type NmspOper struct {
	CiscoIOSXEWirelessNmspOperData struct {
		ClientRegistration []ClientRegistration `json:"client-registration"`
		CmxConnection      []CmxConnection      `json:"cmx-connection"`
		CmxCloudInfo       CmxCloudInfo         `json:"cmx-cloud-info"`
	} `json:"Cisco-IOS-XE-wireless-nmsp-oper:nmsp-oper-data"`
}

// NmspClientRegistration  represents the NMSP client registration data.
type NmspClientRegistration struct {
	ClientRegistration []ClientRegistration `json:"Cisco-IOS-XE-wireless-nmsp-oper:client-registration"`
}

// NmspCmxConnection  represents the NMSP CMX connection data.
type NmspCmxConnection struct {
	CmxConnection []CmxConnection `json:"Cisco-IOS-XE-wireless-nmsp-oper:cmx-connection"`
}

// NmspCmxCloudInfo  represents the NMSP CMX cloud information.
type NmspCmxCloudInfo struct {
	CmxCloudInfo CmxCloudInfo `json:"Cisco-IOS-XE-wireless-nmsp-oper:cmx-cloud-info"`
}

// NMSP Supporting Types

type ClientRegistration struct {
	ClientID int          `json:"client-id"`
	Services NmspServices `json:"services"`
}

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

type CmxConStats struct {
	TxMsgCounter        []MsgCounter `json:"tx-msg-counter"`
	RxMsgCounter        []MsgCounter `json:"rx-msg-counter"`
	UnsupportedMsgCount string       `json:"unsupported-msg-count"`
	TxDataFrames        string       `json:"tx-data-frames"`
	RxDataFrames        string       `json:"rx-data-frames"`
	Connections         string       `json:"connections"`
	Disconnections      string       `json:"disconnections"`
}

type MsgCounter struct {
	Counter string `json:"counter"`
	MsgID   int    `json:"msg-id"`
}

type CmxCloudInfo struct {
	CloudStatus CloudStatus `json:"cloud-status"`
	CloudStats  CloudStats  `json:"cloud-stats"`
}

type CloudStatus struct {
	IPAddress         string `json:"ip-address"`
	Connectivity      string `json:"connectivity"`
	ServiceUp         bool   `json:"service-up"`
	LastRequestStatus string `json:"last-request-status"`
	HeartbeatStatusOk bool   `json:"heartbeat-status-ok"`
}

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
