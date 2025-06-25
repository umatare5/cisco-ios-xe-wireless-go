// Package mdns provides multicast DNS operational data functionality for the Cisco Wireless Network Controller API.
package mdns

import (
	"context"

	wnc "github.com/umatare5/cisco-xe-wireless-restconf-go"
)

const (
	// MdnsOperBasePath defines the base path for mDNS operational data endpoints.
	MdnsOperBasePath = "/restconf/data/Cisco-IOS-XE-wireless-mdns-oper:mdns-oper-data"
	// MdnsOperEndpoint defines the endpoint for mDNS operational data.
	MdnsOperEndpoint = MdnsOperBasePath
	// MdnsGlobalStatsEndpoint defines the endpoint for mDNS global statistics.
	MdnsGlobalStatsEndpoint = MdnsOperBasePath + "/mdns-global-stats"
	// MdnsWlanStatsEndpoint defines the endpoint for mDNS WLAN statistics.
	MdnsWlanStatsEndpoint = MdnsOperBasePath + "/mdns-wlan-stats"
)

// MdnsOperResponse represents the response structure for mDNS operational data.
type MdnsOperResponse struct {
	CiscoIOSXEWirelessMdnsOperMdnsOperData struct {
		MdnsGlobalStats MdnsGlobalStats `json:"mdns-global-stats"`
		MdnsWlanStats   []MdnsWlanStat  `json:"mdns-wlan-stats"`
	} `json:"Cisco-IOS-XE-wireless-mdns-oper:mdns-oper-data"`
}

// MdnsGlobalStatsResponse represents the response structure for mDNS global statistics.
type MdnsGlobalStatsResponse struct {
	MdnsGlobalStats MdnsGlobalStats `json:"Cisco-IOS-XE-wireless-mdns-oper:mdns-global-stats"`
}

// MdnsWlanStatsResponse represents the response structure for mDNS WLAN statistics.
type MdnsWlanStatsResponse struct {
	MdnsWlanStats []MdnsWlanStat `json:"Cisco-IOS-XE-wireless-mdns-oper:mdns-wlan-stats"`
}

// MdnsGlobalStats represents global mDNS statistics including counters and timestamps.
type MdnsGlobalStats struct {
	StatsGlobal   MdnsStats `json:"stats-global"`
	LastClearTime string    `json:"last-clear-time"`
}

// MdnsWlanStat represents mDNS statistics for a specific WLAN.
type MdnsWlanStat struct {
	WlanID        int       `json:"wlan-id"`
	StatsWlan     MdnsStats `json:"stats-wlan"`
	LastClearTime string    `json:"last-clear-time"`
}

// MdnsStats represents detailed mDNS statistics including packet counts for IPv4 and IPv6.
type MdnsStats struct {
	PakSent            string `json:"pak-sent"`
	PakSentV4          string `json:"pak-sent-v4"`
	PakSentAdvtV4      string `json:"pak-sent-advt-v4"`
	PakSentQueryV4     string `json:"pak-sent-query-v4"`
	PakSentV6          string `json:"pak-sent-v6"`
	PakSentAdvtV6      string `json:"pak-sent-advt-v6"`
	PakSentQueryV6     string `json:"pak-sent-query-v6"`
	PakSentMcast       string `json:"pak-sent-mcast"`
	PakSentMcastV4     string `json:"pak-sent-mcast-v4"`
	PakSentMcastV6     string `json:"pak-sent-mcast-v6"`
	PakReceived        string `json:"pak-received"`
	PakReceivedAdvt    string `json:"pak-received-advt"`
	PakReceivedQuery   string `json:"pak-received-query"`
	PakReceivedV4      string `json:"pak-received-v4"`
	PakReceivedAdvtV4  string `json:"pak-received-advt-v4"`
	PakReceivedQueryV4 string `json:"pak-received-query-v4"`
	PakReceivedV6      string `json:"pak-received-v6"`
	PakReceivedAdvtV6  string `json:"pak-received-advt-v6"`
	PakReceivedQueryV6 string `json:"pak-received-query-v6"`
	PakDropped         string `json:"pak-dropped"`
	PtrQuery           string `json:"ptr-query"`
	SrvQuery           string `json:"srv-query"`
	AQuery             string `json:"a-query"`
	AaaaQuery          string `json:"aaaa-query"`
	TxtQuery           string `json:"txt-query"`
	AnyQuery           string `json:"any-query"`
	OtherQuery         string `json:"other-query"`
}

// GetMdnsOper retrieves mDNS operational data.
func GetMdnsOper(client *wnc.Client, ctx context.Context) (*MdnsOperResponse, error) {
	var data MdnsOperResponse
	if err := client.SendAPIRequest(ctx, MdnsOperEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetMdnsGlobalStats retrieves mDNS global statistics.
func GetMdnsGlobalStats(client *wnc.Client, ctx context.Context) (*MdnsGlobalStatsResponse, error) {
	var data MdnsGlobalStatsResponse
	if err := client.SendAPIRequest(ctx, MdnsGlobalStatsEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetMdnsWlanStats retrieves mDNS WLAN statistics.
func GetMdnsWlanStats(client *wnc.Client, ctx context.Context) (*MdnsWlanStatsResponse, error) {
	var data MdnsWlanStatsResponse
	if err := client.SendAPIRequest(ctx, MdnsWlanStatsEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
