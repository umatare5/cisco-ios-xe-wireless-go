package mdns

// MdnsOper represents mDNS operational data container.
type MdnsOper struct {
	CiscoIOSXEWirelessMdnsOperMdnsOperData struct {
		MdnsGlobalStats MdnsGlobalStatsData `json:"mdns-global-stats"` // mDNS global statistics (Live: IOS-XE 17.12.5)
		MdnsWlanStats   []MdnsWlanStat      `json:"mdns-wlan-stats"`   // mDNS statistics per WLAN (Live: IOS-XE 17.12.5)
	} `json:"Cisco-IOS-XE-wireless-mdns-oper:mdns-oper-data"` // mDNS gateway operational data (Live: IOS-XE 17.12.5)
}

// MdnsGlobalStats represents mDNS global statistics response wrapper.
type MdnsGlobalStats struct {
	MdnsGlobalStats MdnsGlobalStatsData `json:"Cisco-IOS-XE-wireless-mdns-oper:mdns-global-stats"`
}

// MdnsWlanStats represents mDNS WLAN statistics response wrapper.
type MdnsWlanStats struct {
	MdnsWlanStats []MdnsWlanStat `json:"Cisco-IOS-XE-wireless-mdns-oper:mdns-wlan-stats"`
}

// MdnsGlobalStatsData represents global mDNS statistics container.
type MdnsGlobalStatsData struct {
	StatsGlobal   MdnsStats `json:"stats-global"`    // Global mDNS packet statistics (Live: IOS-XE 17.12.5)
	LastClearTime string    `json:"last-clear-time"` // mDNS statistics reset timestamp (Live: IOS-XE 17.12.5)
}

// MdnsWlanStat represents per-WLAN mDNS statistics.
type MdnsWlanStat struct {
	WlanID        int       `json:"wlan-id"`         // WLAN identifier (Live: IOS-XE 17.12.5)
	StatsWlan     MdnsStats `json:"stats-wlan"`      // mDNS statistics for WLAN (Live: IOS-XE 17.12.5)
	LastClearTime string    `json:"last-clear-time"` // mDNS statistics reset timestamp (Live: IOS-XE 17.12.5)
}

// MdnsStats represents mDNS packet statistics.
type MdnsStats struct {
	PakSent            string `json:"pak-sent"`              // Total number of mDNS packets sent (Live: IOS-XE 17.12.5)
	PakSentV4          string `json:"pak-sent-v4"`           // Total number of IPv4 mDNS packets sent (Live: IOS-XE 17.12.5)
	PakSentAdvtV4      string `json:"pak-sent-advt-v4"`      // Total number of IPv4 mDNS advertisement packets sent (Live: IOS-XE 17.12.5)
	PakSentQueryV4     string `json:"pak-sent-query-v4"`     // Total number of IPv4 mDNS query packets sent (Live: IOS-XE 17.12.5)
	PakSentV6          string `json:"pak-sent-v6"`           // Total number of IPv6 mDNS packets sent (Live: IOS-XE 17.12.5)
	PakSentAdvtV6      string `json:"pak-sent-advt-v6"`      // Total number of IPv6 mDNS advertisement packets sent (Live: IOS-XE 17.12.5)
	PakSentQueryV6     string `json:"pak-sent-query-v6"`     // Total number of IPv6 mDNS query packets sent (Live: IOS-XE 17.12.5)
	PakSentMcast       string `json:"pak-sent-mcast"`        // Total number of mDNS multicast packets sent (Live: IOS-XE 17.12.5)
	PakSentMcastV4     string `json:"pak-sent-mcast-v4"`     // Total number of IPv4 mDNS multicast packets sent (Live: IOS-XE 17.12.5)
	PakSentMcastV6     string `json:"pak-sent-mcast-v6"`     // Total number of IPv6 mDNS multicast packets sent (Live: IOS-XE 17.12.5)
	PakReceived        string `json:"pak-received"`          // Total number of mDNS packets received (Live: IOS-XE 17.12.5)
	PakReceivedAdvt    string `json:"pak-received-advt"`     // Total number of mDNS advertisement packets received (Live: IOS-XE 17.12.5)
	PakReceivedQuery   string `json:"pak-received-query"`    // Total number of mDNS query packets received (Live: IOS-XE 17.12.5)
	PakReceivedV4      string `json:"pak-received-v4"`       // Total number of IPv4 mDNS packets received (Live: IOS-XE 17.12.5)
	PakReceivedAdvtV4  string `json:"pak-received-advt-v4"`  // Total number of IPv4 mDNS advertisement packets received (Live: IOS-XE 17.12.5)
	PakReceivedQueryV4 string `json:"pak-received-query-v4"` // Total number of IPv4 mDNS query packets received (Live: IOS-XE 17.12.5)
	PakReceivedV6      string `json:"pak-received-v6"`       // Total number of IPv6 mDNS packets received (Live: IOS-XE 17.12.5)
	PakReceivedAdvtV6  string `json:"pak-received-advt-v6"`  // Total number of IPv6 mDNS advertisement packets received (Live: IOS-XE 17.12.5)
	PakReceivedQueryV6 string `json:"pak-received-query-v6"` // Total number of IPv6 mDNS query packets received (Live: IOS-XE 17.12.5)
	PakDropped         string `json:"pak-dropped"`           // Total number of mDNS packets dropped (Live: IOS-XE 17.12.5)
	PtrQuery           string `json:"ptr-query"`             // Total number of PTR queries (Live: IOS-XE 17.12.5)
	SrvQuery           string `json:"srv-query"`             // Total number of SRV queries (Live: IOS-XE 17.12.5)
	AQuery             string `json:"a-query"`               // Total number of IPV4 queries (Live: IOS-XE 17.12.5)
	AaaaQuery          string `json:"aaaa-query"`            // Total number of IPV6 queries (Live: IOS-XE 17.12.5)
	TxtQuery           string `json:"txt-query"`             // Total number of TEXT queries (Live: IOS-XE 17.12.5)
	AnyQuery           string `json:"any-query"`             // Total number of ANY queries (Live: IOS-XE 17.12.5)
	OtherQuery         string `json:"other-query"`           // Total number of OTHER queries (Live: IOS-XE 17.12.5)
}
