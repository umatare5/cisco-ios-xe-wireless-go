// Package afc provides data structures for AFC (Automated Frequency Coordination) operations.
package afc

// AfcCloudOper represents AFC cloud operational data structure.
type AfcCloudOper struct {
	CiscoIOSXEWirelessAfcCloudOperAfcCloudOperData struct {
		AfcCloudStats AfcCloudStats `json:"afc-cloud-stats"`
	} `json:"Cisco-IOS-XE-wireless-afc-cloud-oper:afc-cloud-oper-data"`
}

// AfcCloudOperAfcCloudStats represents AFC cloud statistics data container.
type AfcCloudOperAfcCloudStats struct {
	AfcCloudStats AfcCloudStats `json:"Cisco-IOS-XE-wireless-afc-cloud-oper:afc-cloud-stats"`
}

// AfcCloudStats represents AFC cloud service statistics and monitoring data.
type AfcCloudStats struct {
	NumAfcAp      int    `json:"num-afc-ap"`                // Number of AFC-enabled access points
	AfcMsgSent    string `json:"afc-msg-sent,omitempty"`    // Total AFC messages sent
	AfcMsgRcvd    string `json:"afc-msg-rcvd,omitempty"`    // Total AFC messages received
	AfcMsgErr     string `json:"afc-msg-err,omitempty"`     // Total AFC message errors
	AfcMsgPending int    `json:"afc-msg-pending,omitempty"` // Pending AFC messages count
	LastMsgSent   *struct {
		RequestID    string `json:"request-id"`    // Last sent message request ID
		ApMac        string `json:"ap-mac"`        // AP MAC address for last sent message
		MsgTimestamp string `json:"msg-timestamp"` // Timestamp of last sent message
	} `json:"last-msg-sent,omitempty"`
	LastMsgRcvd *struct {
		RequestID    string `json:"request-id"`    // Last received message request ID
		ApMac        string `json:"ap-mac"`        // AP MAC address for last received message
		MsgTimestamp string `json:"msg-timestamp"` // Timestamp of last received message
	} `json:"last-msg-rcvd,omitempty"`
	MinMsgRtt   string `json:"min-msg-rtt,omitempty"` // Minimum message round-trip time
	MaxMsgRtt   string `json:"max-msg-rtt,omitempty"` // Maximum message round-trip time
	AvgRtt      string `json:"avg-rtt,omitempty"`     // Average round-trip time
	Healthcheck struct {
		HcTimestamp         string                     `json:"hc-timestamp,omitempty"`      // Health check timestamp
		QueryInProgress     bool                       `json:"query-in-progress,omitempty"` // Health check query in progress status
		CountryNotSupported bool                       `json:"country-not-supported"`       // Country not supported flag
		NumHcDown           int                        `json:"num-hc-down,omitempty"`       // Number of health checks down
		HcErrorStatus       *AfcHealthcheckErrorStatus `json:"hc-error-status,omitempty"`   // Health check error status details
		CloudHcOk           *bool                      `json:"cloud-hc-ok,omitempty"`       // Cloud health check success status (YANG: IOS-XE 17.18.1+)
		CloudHcUnknown      *bool                      `json:"cloud-hc-unknown,omitempty"`  // Cloud health check unknown status (YANG: IOS-XE 17.18.1+)
	} `json:"healthcheck"`
	Num6GhzAp int `json:"num-6ghz-ap,omitempty"` // Number of 6GHz capable access points
}

// AfcHealthcheckErrorStatus represents AFC health check error status details.
type AfcHealthcheckErrorStatus struct {
	NotOtpUpgraded     *bool            `json:"not-otp-upgraded,omitempty"`     // OTP upgrade status flag (YANG: IOS-XE 17.18.1+)
	HTTPConError       *AfcHTTPConError `json:"httpcon-error,omitempty"`        // HTTP connection error details (YANG: IOS-XE 17.18.1+)
	FwStatus           string           `json:"fw-status,omitempty"`            // Firmware status (YANG: IOS-XE 17.18.1+)
	NoValidToken       *bool            `json:"no-valid-token,omitempty"`       // Valid token availability flag (YANG: IOS-XE 17.18.1+)
	DeviceNotOnboarded *bool            `json:"device-not-onboarded,omitempty"` // Device onboarding status flag (YANG: IOS-XE 17.18.1+)
	ErrorUnknown       *bool            `json:"error-unknown,omitempty"`        // Unknown error status flag (YANG: IOS-XE 17.18.1+)
}

// AfcHTTPConError represents AFC HTTP connection error details.
type AfcHTTPConError struct {
	HTTPErrorCode    *uint32 `json:"http-error-code,omitempty"`    // HTTP error code (YANG: IOS-XE 17.18.1+)
	HTTPConErrorCode *uint32 `json:"httpcon-error-code,omitempty"` // HTTP connection error code (YANG: IOS-XE 17.18.1+)
}
