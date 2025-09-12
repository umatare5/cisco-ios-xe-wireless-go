package mdns

// MdnsOper represents mDNS operational data container.
type MdnsOper struct {
	CiscoIOSXEWirelessMdnsOperMdnsOperData struct {
		MdnsGlobalStats MdnsGlobalStatsData `json:"mdns-global-stats"`
		MdnsWlanStats   []MdnsWlanStat      `json:"mdns-wlan-stats"`
	} `json:"Cisco-IOS-XE-wireless-mdns-oper:mdns-oper-data"`
}

// MdnsGlobalStats represents mDNS global statistics response wrapper.
type MdnsGlobalStats struct {
	MdnsGlobalStats MdnsGlobalStatsData `json:"Cisco-IOS-XE-wireless-mdns-oper:mdns-global-stats"` // Global mDNS statistics data
}

// MdnsWlanStats represents mDNS WLAN statistics response wrapper.
type MdnsWlanStats struct {
	MdnsWlanStats []MdnsWlanStat `json:"Cisco-IOS-XE-wireless-mdns-oper:mdns-wlan-stats"` // Per-WLAN mDNS statistics list
}

// MdnsGlobalStatsData represents global mDNS statistics container.
type MdnsGlobalStatsData struct {
	StatsGlobal   MdnsStats `json:"stats-global"`    // Global mDNS packet statistics
	LastClearTime string    `json:"last-clear-time"` // Statistics reset timestamp
}

// MdnsWlanStat represents per-WLAN mDNS statistics.
type MdnsWlanStat struct {
	WlanID        int       `json:"wlan-id"`         // WLAN identifier
	StatsWlan     MdnsStats `json:"stats-wlan"`      // mDNS packet statistics for WLAN
	LastClearTime string    `json:"last-clear-time"` // Statistics reset timestamp
}

// MdnsStats represents mDNS packet statistics.
type MdnsStats struct {
	PakSent            string `json:"pak-sent"`              // Total mDNS packets sent
	PakSentV4          string `json:"pak-sent-v4"`           // Total IPv4 mDNS packets sent
	PakSentAdvtV4      string `json:"pak-sent-advt-v4"`      // Total IPv4 mDNS advertisement packets sent
	PakSentQueryV4     string `json:"pak-sent-query-v4"`     // Total IPv4 mDNS query packets sent
	PakSentV6          string `json:"pak-sent-v6"`           // Total IPv6 mDNS packets sent
	PakSentAdvtV6      string `json:"pak-sent-advt-v6"`      // Total IPv6 mDNS advertisement packets sent
	PakSentQueryV6     string `json:"pak-sent-query-v6"`     // Total IPv6 mDNS query packets sent
	PakSentMcast       string `json:"pak-sent-mcast"`        // Total mDNS multicast packets sent
	PakSentMcastV4     string `json:"pak-sent-mcast-v4"`     // Total IPv4 mDNS multicast packets sent
	PakSentMcastV6     string `json:"pak-sent-mcast-v6"`     // Total IPv6 mDNS multicast packets sent
	PakReceived        string `json:"pak-received"`          // Total mDNS packets received
	PakReceivedAdvt    string `json:"pak-received-advt"`     // Total mDNS advertisement packets received
	PakReceivedQuery   string `json:"pak-received-query"`    // Total mDNS query packets received
	PakReceivedV4      string `json:"pak-received-v4"`       // Total IPv4 mDNS packets received
	PakReceivedAdvtV4  string `json:"pak-received-advt-v4"`  // Total IPv4 mDNS advertisement packets received
	PakReceivedQueryV4 string `json:"pak-received-query-v4"` // Total IPv4 mDNS query packets received
	PakReceivedV6      string `json:"pak-received-v6"`       // Total IPv6 mDNS packets received
	PakReceivedAdvtV6  string `json:"pak-received-advt-v6"`  // Total IPv6 mDNS advertisement packets received
	PakReceivedQueryV6 string `json:"pak-received-query-v6"` // Total IPv6 mDNS query packets received
	PakDropped         string `json:"pak-dropped"`           // Total mDNS packets dropped
	PtrQuery           string `json:"ptr-query"`             // Total PTR queries
	SrvQuery           string `json:"srv-query"`             // Total SRV queries
	AQuery             string `json:"a-query"`               // Total IPv4 A queries
	AaaaQuery          string `json:"aaaa-query"`            // Total IPv6 AAAA queries
	TxtQuery           string `json:"txt-query"`             // Total TEXT queries
	AnyQuery           string `json:"any-query"`             // Total ANY queries
	OtherQuery         string `json:"other-query"`           // Total OTHER queries
}
