// Package nmsp provides Network Mobility Services Protocol operational data functionality for the Cisco Wireless Network Controller API.
package nmsp

import (
	"context"

	wnc "github.com/umatare5/cisco-xe-wireless-restconf-go"
)

const (
	// NmspOperBasePath defines the base path for NMSP operational data endpoints.
	NmspOperBasePath = "/restconf/data/Cisco-IOS-XE-wireless-nmsp-oper:nmsp-oper-data"
	// NmspOperEndpoint defines the endpoint for NMSP operational data.
	NmspOperEndpoint = NmspOperBasePath
	// ClientRegistrationEndpoint defines the endpoint for client registration data.
	ClientRegistrationEndpoint = NmspOperBasePath + "/client-registration"
	// CmxConnectionEndpoint defines the endpoint for CMX connection data.
	CmxConnectionEndpoint = NmspOperBasePath + "/cmx-connection"
	// CmxCloudInfoEndpoint defines the endpoint for CMX cloud information.
	CmxCloudInfoEndpoint = NmspOperBasePath + "/cmx-cloud-info"
)

// NmspOperResponse represents the response structure for NMSP operational data.
type NmspOperResponse struct {
	CiscoIOSXEWirelessNmspOperData struct {
		ClientRegistration []ClientRegistration `json:"client-registration"`
		CmxConnection      []CmxConnection      `json:"cmx-connection"`
		CmxCloudInfo       CmxCloudInfo         `json:"cmx-cloud-info"`
	} `json:"Cisco-IOS-XE-wireless-nmsp-oper:nmsp-oper-data"`
}

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

// NmspClientRegistrationResponse represents the response structure for NMSP client registration data.
type NmspClientRegistrationResponse struct {
	ClientRegistration []ClientRegistration `json:"Cisco-IOS-XE-wireless-nmsp-oper:client-registration"`
}

// NmspCmxConnectionResponse represents the response structure for NMSP CMX connection data.
type NmspCmxConnectionResponse struct {
	CmxConnection []CmxConnection `json:"Cisco-IOS-XE-wireless-nmsp-oper:cmx-connection"`
}

type NmspCmxCloudInfoResponse struct {
	CmxCloudInfo CmxCloudInfo `json:"Cisco-IOS-XE-wireless-nmsp-oper:cmx-cloud-info"`
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

// GetNmspOper retrieves NMSP operational data.
func GetNmspOper(client *wnc.Client, ctx context.Context) (*NmspOperResponse, error) {
	var data NmspOperResponse
	if err := client.SendAPIRequest(ctx, NmspOperEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetNmspClientRegistration retrieves NMSP client registration data.
func GetNmspClientRegistration(client *wnc.Client, ctx context.Context) (*NmspClientRegistrationResponse, error) {
	var data NmspClientRegistrationResponse
	if err := client.SendAPIRequest(ctx, ClientRegistrationEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetNmspCmxConnection retrieves NMSP CMX connection data.
func GetNmspCmxConnection(client *wnc.Client, ctx context.Context) (*NmspCmxConnectionResponse, error) {
	var data NmspCmxConnectionResponse
	if err := client.SendAPIRequest(ctx, CmxConnectionEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetNmspCmxCloudInfo retrieves NMSP CMX cloud information.
func GetNmspCmxCloudInfo(client *wnc.Client, ctx context.Context) (*NmspCmxCloudInfoResponse, error) {
	var data NmspCmxCloudInfoResponse
	if err := client.SendAPIRequest(ctx, CmxCloudInfoEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
