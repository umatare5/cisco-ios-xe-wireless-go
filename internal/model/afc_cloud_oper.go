package model

// AfcCloudOperResponse represents the response structure for AFC cloud operational data.
type AfcCloudOperResponse struct {
	CiscoIOSXEWirelessAfcCloudOperAfcCloudOperData struct {
		AfcCloudStats AfcCloudStats `json:"afc-cloud-stats"`
	} `json:"Cisco-IOS-XE-wireless-afc-cloud-oper:afc-cloud-oper-data"`
}

// AfcCloudOperAfcCloudStatsResponse represents the response structure for AFC cloud statistics.
type AfcCloudOperAfcCloudStatsResponse struct {
	AfcCloudStats AfcCloudStats `json:"Cisco-IOS-XE-wireless-afc-cloud-stats"`
}

// AfcCloudStats represents AFC cloud statistics including message counts, RTT metrics, and health check information.
type AfcCloudStats struct {
	NumAfcAp      int    `json:"num-afc-ap"`
	AfcMsgSent    string `json:"afc-msg-sent"`
	AfcMsgRcvd    string `json:"afc-msg-rcvd"`
	AfcMsgErr     string `json:"afc-msg-err"`
	AfcMsgPending int    `json:"afc-msg-pending"`
	LastMsgSent   struct {
		RequestID    string `json:"request-id"`
		ApMac        string `json:"ap-mac"`
		MsgTimestamp string `json:"msg-timestamp"`
	} `json:"last-msg-sent"`
	LastMsgRcvd struct {
		RequestID    string `json:"request-id"`
		ApMac        string `json:"ap-mac"`
		MsgTimestamp string `json:"msg-timestamp"`
	} `json:"last-msg-rcvd"`
	MinMsgRtt   string `json:"min-msg-rtt"`
	MaxMsgRtt   string `json:"max-msg-rtt"`
	AvgRtt      string `json:"avg-rtt"`
	Healthcheck struct {
		HcTimestamp         string `json:"hc-timestamp"`
		QueryInProgress     bool   `json:"query-in-progress"`
		CountryNotSupported bool   `json:"country-not-supported"`
		NumHcDown           int    `json:"num-hc-down"`
		HcErrorStatus       struct {
			NotOtpUpgraded bool `json:"not-otp-upgraded"`
		} `json:"hc-error-status"`
	} `json:"healthcheck"`
	Num6GhzAp int `json:"num-6ghz-ap"`
}
