package afc

// AfcCloudOper represents AFC cloud operational data structure.
type AfcCloudOper struct {
	CiscoIOSXEWirelessAfcCloudOperAfcCloudOperData struct {
		AfcCloudStats AfcCloudStats `json:"afc-cloud-stats"` // AFC cloud statistics (Live: IOS-XE 17.12.5)
	} `json:"Cisco-IOS-XE-wireless-afc-cloud-oper:afc-cloud-oper-data"` // AFC cloud operational data (Live: IOS-XE 17.12.5)
}

// AfcCloudOperAfcCloudStats represents AFC cloud statistics data container.
type AfcCloudOperAfcCloudStats struct {
	AfcCloudStats AfcCloudStats `json:"Cisco-IOS-XE-wireless-afc-cloud-oper:afc-cloud-stats"`
}

// AfcCloudStats represents AFC cloud service statistics and monitoring data.
type AfcCloudStats struct {
	NumAfcAp      int    `json:"num-afc-ap"`                // Number of APs requiring AFC service (Live: IOS-XE 17.12.5)
	AfcMsgSent    string `json:"afc-msg-sent,omitempty"`    // Messages sent to AFC (Live: IOS-XE 17.12.5)
	AfcMsgRcvd    string `json:"afc-msg-rcvd,omitempty"`    // Successful messages received from AFC (Live: IOS-XE 17.12.5)
	AfcMsgErr     string `json:"afc-msg-err,omitempty"`     // Errored AFC messages (Live: IOS-XE 17.12.5)
	AfcMsgPending int    `json:"afc-msg-pending,omitempty"` // Pending AFC messages (Live: IOS-XE 17.12.5)
	LastMsgSent   *struct {
		RequestID    string `json:"request-id"`    // Request ID (Live: IOS-XE 17.12.5)
		ApMac        string `json:"ap-mac"`        // AP MAC address (Live: IOS-XE 17.12.5)
		MsgTimestamp string `json:"msg-timestamp"` // Timestamp (Live: IOS-XE 17.12.5)
	} `json:"last-msg-sent,omitempty"`
	LastMsgRcvd *struct {
		RequestID    string `json:"request-id"`    // Request ID (Live: IOS-XE 17.12.5)
		ApMac        string `json:"ap-mac"`        // AP MAC address (Live: IOS-XE 17.12.5)
		MsgTimestamp string `json:"msg-timestamp"` // Timestamp (Live: IOS-XE 17.12.5)
	} `json:"last-msg-rcvd,omitempty"`
	MinMsgRtt   string `json:"min-msg-rtt,omitempty"` // Minimum response time (Live: IOS-XE 17.12.5)
	MaxMsgRtt   string `json:"max-msg-rtt,omitempty"` // Maximum response time (Live: IOS-XE 17.12.5)
	AvgRtt      string `json:"avg-rtt,omitempty"`     // Average response time (Live: IOS-XE 17.12.5)
	Healthcheck struct {
		HcTimestamp         string                     `json:"hc-timestamp,omitempty"`      // Cloud health check timestamp (Live: IOS-XE 17.12.5)
		QueryInProgress     bool                       `json:"query-in-progress,omitempty"` // Cloud health check query in progress (Live: IOS-XE 17.12.5)
		CountryNotSupported bool                       `json:"country-not-supported"`       // Country not supported by AFC (Live: IOS-XE 17.12.5)
		NumHcDown           int                        `json:"num-hc-down,omitempty"`       // Number of times cloud health check failed (Live: IOS-XE 17.12.5)
		HcErrorStatus       *AfcHealthcheckErrorStatus `json:"hc-error-status,omitempty"`   // AFC cloud health check error (Live: IOS-XE 17.12.5)
		CloudHcOk           *bool                      `json:"cloud-hc-ok,omitempty"`       // Cloud health check success status (YANG: IOS-XE 17.18.1)
		CloudHcUnknown      *bool                      `json:"cloud-hc-unknown,omitempty"`  // Cloud health check unknown status (YANG: IOS-XE 17.18.1)
	} `json:"healthcheck"`
	Num6GhzAp int `json:"num-6ghz-ap,omitempty"` // Number of APs with 6GHz radio (Live: IOS-XE 17.12.5)
}

// AfcHealthcheckErrorStatus represents AFC health check error status details.
type AfcHealthcheckErrorStatus struct {
	NotOtpUpgraded     *bool            `json:"not-otp-upgraded,omitempty"`     // Device is not OTP upgraded (YANG: IOS-XE 17.18.1)
	HTTPConError       *AfcHTTPConError `json:"httpcon-error,omitempty"`        // HTTP connection error codes (YANG: IOS-XE 17.18.1)
	FwStatus           string           `json:"fw-status,omitempty"`            // AFC provider status (YANG: IOS-XE 17.18.1)
	NoValidToken       *bool            `json:"no-valid-token,omitempty"`       // Device does not have a valid token (YANG: IOS-XE 17.18.1)
	DeviceNotOnboarded *bool            `json:"device-not-onboarded,omitempty"` // Device is not on boarded (YANG: IOS-XE 17.18.1)
	ErrorUnknown       *bool            `json:"error-unknown,omitempty"`        // Error status is unknown (YANG: IOS-XE 17.18.1)
}

// AfcHTTPConError represents AFC HTTP connection error details.
type AfcHTTPConError struct {
	HTTPErrorCode    *uint32 `json:"http-error-code,omitempty"`    // HTTP error code (YANG: IOS-XE 17.18.1)
	HTTPConErrorCode *uint32 `json:"httpcon-error-code,omitempty"` // HTTP connection error (YANG: IOS-XE 17.18.1)
}
