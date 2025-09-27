package mdns

// CiscoIOSXEWirelessMDNSOper represents mDNS operational data container.
type CiscoIOSXEWirelessMDNSOper struct {
	CiscoIOSXEWirelessMDNSOperData struct {
		MDNSGlobalStats CiscoIOSXEWirelessMDNSGlobalStats `json:"mdns-global-stats"` // mDNS global statistics (Live: IOS-XE 17.12.6a)
		MDNSWlanStats   []CiscoIOSXEWirelessMDNSWlanStats `json:"mdns-wlan-stats"`   // mDNS per-WLAN statistics (Live: IOS-XE 17.12.6a)
	} `json:"Cisco-IOS-XE-wireless-mdns-oper:mdns-oper-data"` // mDNS operational data (Live: IOS-XE 17.12.6a)
}

// CiscoIOSXEWirelessMDNSGlobalStats represents mDNS global statistics response wrapper.
type CiscoIOSXEWirelessMDNSGlobalStats struct {
	MDNSGlobalStats MDNSGlobalStatsData `json:"Cisco-IOS-XE-wireless-mdns-oper:mdns-global-stats"`
}

// CiscoIOSXEWirelessMDNSWlanStats represents mDNS WLAN statistics response wrapper.
type CiscoIOSXEWirelessMDNSWlanStats struct {
	MDNSWlanStats []MDNSWlanStat `json:"Cisco-IOS-XE-wireless-mdns-oper:mdns-wlan-stats"`
}

// MDNSGlobalStatsData represents global mDNS statistics container.
type MDNSGlobalStatsData struct {
	StatsGlobal   MDNSStats `json:"stats-global"`    // Global mDNS packet statistics (Live: IOS-XE 17.12.6a)
	LastClearTime string    `json:"last-clear-time"` // mDNS statistics reset timestamp (Live: IOS-XE 17.12.6a)
}

// MDNSWlanStat represents per-WLAN mDNS statistics.
type MDNSWlanStat struct {
	WlanID        int       `json:"wlan-id"`         // WLAN identifier (Live: IOS-XE 17.12.6a)
	StatsWlan     MDNSStats `json:"stats-wlan"`      // mDNS statistics for WLAN (Live: IOS-XE 17.12.6a)
	LastClearTime string    `json:"last-clear-time"` // mDNS statistics reset timestamp (Live: IOS-XE 17.12.6a)
}

// MDNSStats represents mDNS packet statistics.
type MDNSStats struct {
	PakSent            string `json:"pak-sent"`              // Total number of mDNS packets sent (Live: IOS-XE 17.12.6a)
	PakSentV4          string `json:"pak-sent-v4"`           // Total number of IPv4 mDNS packets sent (Live: IOS-XE 17.12.6a)
	PakSentAdvtV4      string `json:"pak-sent-advt-v4"`      // Total number of IPv4 mDNS advertisement packets sent (Live: IOS-XE 17.12.6a)
	PakSentQueryV4     string `json:"pak-sent-query-v4"`     // Total number of IPv4 mDNS query packets sent (Live: IOS-XE 17.12.6a)
	PakSentV6          string `json:"pak-sent-v6"`           // Total number of IPv6 mDNS packets sent (Live: IOS-XE 17.12.6a)
	PakSentAdvtV6      string `json:"pak-sent-advt-v6"`      // Total number of IPv6 mDNS advertisement packets sent (Live: IOS-XE 17.12.6a)
	PakSentQueryV6     string `json:"pak-sent-query-v6"`     // Total number of IPv6 mDNS query packets sent (Live: IOS-XE 17.12.6a)
	PakSentMcast       string `json:"pak-sent-mcast"`        // Total number of mDNS multicast packets sent (Live: IOS-XE 17.12.6a)
	PakSentMcastV4     string `json:"pak-sent-mcast-v4"`     // Total number of IPv4 mDNS multicast packets sent (Live: IOS-XE 17.12.6a)
	PakSentMcastV6     string `json:"pak-sent-mcast-v6"`     // Total number of IPv6 mDNS multicast packets sent (Live: IOS-XE 17.12.6a)
	PakReceived        string `json:"pak-received"`          // Total number of mDNS packets received (Live: IOS-XE 17.12.6a)
	PakReceivedAdvt    string `json:"pak-received-advt"`     // Total number of mDNS advertisement packets received (Live: IOS-XE 17.12.6a)
	PakReceivedQuery   string `json:"pak-received-query"`    // Total number of mDNS query packets received (Live: IOS-XE 17.12.6a)
	PakReceivedV4      string `json:"pak-received-v4"`       // Total number of IPv4 mDNS packets received (Live: IOS-XE 17.12.6a)
	PakReceivedAdvtV4  string `json:"pak-received-advt-v4"`  // Total number of IPv4 mDNS advertisement packets received (Live: IOS-XE 17.12.6a)
	PakReceivedQueryV4 string `json:"pak-received-query-v4"` // Total number of IPv4 mDNS query packets received (Live: IOS-XE 17.12.6a)
	PakReceivedV6      string `json:"pak-received-v6"`       // Total number of IPv6 mDNS packets received (Live: IOS-XE 17.12.6a)
	PakReceivedAdvtV6  string `json:"pak-received-advt-v6"`  // Total number of IPv6 mDNS advertisement packets received (Live: IOS-XE 17.12.6a)
	PakReceivedQueryV6 string `json:"pak-received-query-v6"` // Total number of IPv6 mDNS query packets received (Live: IOS-XE 17.12.6a)
	PakDropped         string `json:"pak-dropped"`           // Total number of mDNS packets dropped (Live: IOS-XE 17.12.6a)
	PtrQuery           string `json:"ptr-query"`             // Total number of PTR queries (Live: IOS-XE 17.12.6a)
	SrvQuery           string `json:"srv-query"`             // Total number of SRV queries (Live: IOS-XE 17.12.6a)
	AQuery             string `json:"a-query"`               // Total number of IPV4 queries (Live: IOS-XE 17.12.6a)
	AaaaQuery          string `json:"aaaa-query"`            // Total number of IPV6 queries (Live: IOS-XE 17.12.6a)
	TxtQuery           string `json:"txt-query"`             // Total number of TEXT queries (Live: IOS-XE 17.12.6a)
	AnyQuery           string `json:"any-query"`             // Total number of ANY queries (Live: IOS-XE 17.12.6a)
	OtherQuery         string `json:"other-query"`           // Total number of OTHER queries (Live: IOS-XE 17.12.6a)
}
