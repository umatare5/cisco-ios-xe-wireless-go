// Package afc provides Automated Frequency Coordination cloud operational data functionality for the Cisco Wireless Network Controller API.
package afc

import (
	"context"

	wnc "github.com/umatare5/cisco-xe-wireless-restconf-go"
)

const (
	// AfcCloudOperBasePath defines the base path for AFC cloud operational data endpoints.
	AfcCloudOperBasePath = "/restconf/data/Cisco-IOS-XE-wireless-afc-cloud-oper:afc-cloud-oper-data"
	// AfcCloudOperEndpoint defines the endpoint for AFC cloud operational data.
	AfcCloudOperEndpoint = AfcCloudOperBasePath
	// AfcCloudStatsEndpoint defines the endpoint for AFC cloud statistics.
	AfcCloudStatsEndpoint = AfcCloudOperBasePath + "/afc-cloud-stats"
)

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

// GetAfcCloudOper retrieves AFC cloud operational data.
func GetAfcCloudOper(client *wnc.Client, ctx context.Context) (*AfcCloudOperResponse, error) {
	var data AfcCloudOperResponse
	if err := client.SendAPIRequest(ctx, AfcCloudOperEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetAfcCloudStats retrieves AFC cloud statistics.
func GetAfcCloudStats(client *wnc.Client, ctx context.Context) (*AfcCloudOperAfcCloudStatsResponse, error) {
	var data AfcCloudOperAfcCloudStatsResponse
	if err := client.SendAPIRequest(ctx, AfcCloudStatsEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
