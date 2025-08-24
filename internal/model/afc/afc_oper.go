// Package model provides data structures for AFC (Automated Frequency Coordination) operations.
package model

// AfcOper  represents the AFC operational data.
type AfcOper struct {
	CiscoIOSXEWirelessAfcOperAfcOperData struct {
		EwlcAfcApResp []EwlcAfcApResp `json:"ewlc-afc-ap-resp"`
	} `json:"Cisco-IOS-XE-wireless-afc-oper:afc-oper-data"`
}

// AfcOperEwlcAfcApResp  represents the EWLC AFC AP response data.
type AfcOperEwlcAfcApResp struct {
	EwlcAfcApResp []EwlcAfcApResp `json:"Cisco-IOS-XE-wireless-afc-oper:ewlc-afc-ap-resp"`
}

type EwlcAfcApResp struct {
	ApMac    string `json:"ap-mac"`
	RespData struct {
		RequestID string `json:"request-id"`
		RulesetID string `json:"ruleset-id"`
		RespCode  struct {
			Code             int    `json:"code"`
			Description      string `json:"description"`
			SupplementalInfo string `json:"supplemental-info"`
		} `json:"resp-code"`
		Band20 struct {
			GlobalOperClass int `json:"global-oper-class"`
		} `json:"band20"`
		Band40 struct {
			GlobalOperClass int `json:"global-oper-class"`
		} `json:"band40"`
		Band80 struct {
			GlobalOperClass int `json:"global-oper-class"`
		} `json:"band80"`
		Band160 struct {
			GlobalOperClass int `json:"global-oper-class"`
		} `json:"band160"`
		Band80Plus struct {
			GlobalOperClass int `json:"global-oper-class"`
		} `json:"band80plus"`
		ExpireTime        string `json:"expire-time"`
		RespRcvdTimestamp string `json:"resp-rcvd-timestamp"`
	} `json:"resp-data"`
	Slot int `json:"slot"`
}
