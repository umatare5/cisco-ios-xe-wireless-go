// Package afc provides Automated Frequency Coordination operational data functionality for the Cisco Wireless Network Controller API.
package afc

import (
	"context"

	wnc "github.com/umatare5/cisco-xe-wireless-restconf-go"
)

const (
	// AfcOperBasePath defines the base path for AFC operational data endpoints.
	AfcOperBasePath = "/restconf/data/Cisco-IOS-XE-wireless-afc-oper:afc-oper-data"
	// AfcOperEndpoint defines the endpoint for AFC operational data.
	AfcOperEndpoint = AfcOperBasePath
	// AfcOperEwlcAfcApRespEndpoint defines the endpoint for EWLC AFC AP response data.
	AfcOperEwlcAfcApRespEndpoint = AfcOperBasePath + "/ewlc-afc-ap-resp"
)

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

// GetAfcOper retrieves AFC operational data.
func GetAfcOper(client *wnc.Client, ctx context.Context) (*AfcOperResponse, error) {
	var data AfcOperResponse
	if err := client.SendAPIRequest(ctx, AfcOperEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetAfcEwlcAfcApResp retrieves EWLC AFC AP response data.
func GetAfcEwlcAfcApResp(client *wnc.Client, ctx context.Context) (*AfcOperEwlcAfcApRespResponse, error) {
	var data AfcOperEwlcAfcApRespResponse
	if err := client.SendAPIRequest(ctx, AfcOperEwlcAfcApRespEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
