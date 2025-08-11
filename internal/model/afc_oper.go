// Package model contains generated response structures for the Cisco WNC API.
package model

// AfcOperResponse represents the response structure for AFC operational data.
type AfcOperResponse struct {
	CiscoIOSXEWirelessAfcOperAfcOperData struct {
		EwlcAfcApResp []EwlcAfcApResp `json:"ewlc-afc-ap-resp"`
	} `json:"Cisco-IOS-XE-wireless-afc-oper:afc-oper-data"`
}

// AfcOperEwlcAfcApRespResponse represents the response structure for EWLC AFC AP response data.
type AfcOperEwlcAfcApRespResponse struct {
	EwlcAfcApResp []EwlcAfcApResp `json:"Cisco-IOS-XE-wireless-afc-oper:ewlc-afc-ap-resp"`
}

// EwlcAfcApResp represents an EWLC AFC AP response entry containing AP information and response data.
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
