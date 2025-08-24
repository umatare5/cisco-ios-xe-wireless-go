// Package model provides type definitions for Cisco IOS-XE wireless controller operations.
package model

// mDNS Operational Response Types

// MdnsOper  represents the structure for mDNS operational data.
type MdnsOper struct {
	CiscoIOSXEWirelessMdnsOperMdnsOperData struct {
		MdnsGlobalStats MdnsGlobalStatsData `json:"mdns-global-stats"`
		MdnsWlanStats   []MdnsWlanStat      `json:"mdns-wlan-stats"`
	} `json:"Cisco-IOS-XE-wireless-mdns-oper:mdns-oper-data"`
}

// MdnsGlobalStats  represents the structure for mDNS global statistics.
type MdnsGlobalStats struct {
	MdnsGlobalStats MdnsGlobalStatsData `json:"Cisco-IOS-XE-wireless-mdns-oper:mdns-global-stats"`
}

// MdnsWlanStats  represents the structure for mDNS WLAN statistics.
type MdnsWlanStats struct {
	MdnsWlanStats []MdnsWlanStat `json:"Cisco-IOS-XE-wireless-mdns-oper:mdns-wlan-stats"`
}

// mDNS Supporting Types

type MdnsGlobalStatsData struct {
	StatsGlobal   MdnsStats `json:"stats-global"`
	LastClearTime string    `json:"last-clear-time"`
}

type MdnsWlanStat struct {
	WlanID        int       `json:"wlan-id"`
	StatsWlan     MdnsStats `json:"stats-wlan"`
	LastClearTime string    `json:"last-clear-time"`
}

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
