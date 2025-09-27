package afc

// CiscoIOSXEWirelessAFCCloudOper represents AFC cloud operational data structure.
type CiscoIOSXEWirelessAFCCloudOper struct {
	CiscoIOSXEWirelessAFCCloudOperData struct {
		AFCCloudStats AFCCloudStats `json:"afc-cloud-stats"` // AFC cloud statistics (Live: IOS-XE 17.12.6a)
	} `json:"Cisco-IOS-XE-wireless-afc-cloud-oper:afc-cloud-oper-data"` // AFC cloud operational data (Live: IOS-XE 17.12.6a)
}

// CiscoIOSXEWirelessAFCCloudOperAFCCloudStats represents AFC cloud statistics data container.
type CiscoIOSXEWirelessAFCCloudOperAFCCloudStats struct {
	AFCCloudStats AFCCloudStats `json:"Cisco-IOS-XE-wireless-afc-cloud-oper:afc-cloud-stats"`
}

// AFCCloudStats represents AFC cloud service statistics and monitoring data.
type AFCCloudStats struct {
	NumAFCAp      int    `json:"num-afc-ap"`      // Number of APs requiring AFC service (Live: IOS-XE 17.12.6a)
	AFCMsgSent    string `json:"afc-msg-sent"`    // Messages sent to AFC (Live: IOS-XE 17.12.6a)
	AFCMsgRcvd    string `json:"afc-msg-rcvd"`    // Successful messages received from AFC (Live: IOS-XE 17.12.6a)
	AFCMsgErr     string `json:"afc-msg-err"`     // Errored AFC messages (Live: IOS-XE 17.12.6a)
	AFCMsgPending int    `json:"afc-msg-pending"` // Pending AFC messages (Live: IOS-XE 17.12.6a)
	LastMsgSent   struct {
		RequestID    string `json:"request-id"`    // Request ID (Live: IOS-XE 17.12.6a)
		ApMAC        string `json:"ap-mac"`        // AP MAC address (Live: IOS-XE 17.12.6a)
		MsgTimestamp string `json:"msg-timestamp"` // Timestamp (Live: IOS-XE 17.12.6a)
	} `json:"last-msg-sent"`
	LastMsgRcvd struct {
		RequestID    string `json:"request-id"`    // Request ID (Live: IOS-XE 17.12.6a)
		ApMAC        string `json:"ap-mac"`        // AP MAC address (Live: IOS-XE 17.12.6a)
		MsgTimestamp string `json:"msg-timestamp"` // Timestamp (Live: IOS-XE 17.12.6a)
	} `json:"last-msg-rcvd"`
	MinMsgRtt   string `json:"min-msg-rtt"` // Minimum response time (Live: IOS-XE 17.12.6a)
	MaxMsgRtt   string `json:"max-msg-rtt"` // Maximum response time (Live: IOS-XE 17.12.6a)
	AvgRtt      string `json:"avg-rtt"`     // Average response time (Live: IOS-XE 17.12.6a)
	Healthcheck struct {
		HcTimestamp         string                    `json:"hc-timestamp"`          // Cloud health check timestamp (Live: IOS-XE 17.12.6a)
		QueryInProgress     bool                      `json:"query-in-progress"`     // Cloud health check query in progress (Live: IOS-XE 17.12.6a)
		CountryNotSupported bool                      `json:"country-not-supported"` // Country not supported by AFC (Live: IOS-XE 17.12.6a)
		NumHcDown           int                       `json:"num-hc-down"`           // Number of times cloud health check failed (Live: IOS-XE 17.12.6a)
		HcErrorStatus       AFCHealthcheckErrorStatus `json:"hc-error-status"`       // AFC cloud health check error (Live: IOS-XE 17.12.6a)
		// 17.18.1+ features
		CloudHcOk      *bool `json:"cloud-hc-ok,omitempty"`      // Cloud health check success status (YANG: IOS-XE 17.18.1)
		CloudHcUnknown *bool `json:"cloud-hc-unknown,omitempty"` // Cloud health check unknown status (YANG: IOS-XE 17.18.1)
	} `json:"healthcheck"`
	Num6GhzAp int `json:"num-6ghz-ap"` // Number of APs with 6GHz radio (Live: IOS-XE 17.12.6a)
}

// AFCHealthcheckErrorStatus represents AFC health check error status details.
type AFCHealthcheckErrorStatus struct {
	NotOtpUpgraded     bool             `json:"not-otp-upgraded"`               // Device is not OTP upgraded (Live: IOS-XE 17.12.6a)
	HTTPConError       *AFCHTTPConError `json:"httpcon-error,omitempty"`        // HTTP connection error codes (YANG: IOS-XE 17.18.1)
	FwStatus           string           `json:"fw-status,omitempty"`            // AFC provider status (YANG: IOS-XE 17.18.1)
	NoValidToken       *bool            `json:"no-valid-token,omitempty"`       // Device does not have a valid token (YANG: IOS-XE 17.18.1)
	DeviceNotOnboarded *bool            `json:"device-not-onboarded,omitempty"` // Device is not on boarded (YANG: IOS-XE 17.18.1)
	ErrorUnknown       *bool            `json:"error-unknown,omitempty"`        // Error status is unknown (YANG: IOS-XE 17.18.1)
}

// AFCHTTPConError represents AFC HTTP connection error details.
type AFCHTTPConError struct {
	HTTPErrorCode    *uint32 `json:"http-error-code,omitempty"`    // HTTP error code (YANG: IOS-XE 17.18.1)
	HTTPConErrorCode *uint32 `json:"httpcon-error-code,omitempty"` // HTTP connection error (YANG: IOS-XE 17.18.1)
}
