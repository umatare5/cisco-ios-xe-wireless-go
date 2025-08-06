package model

// AwipsOperResponse represents the response structure for AWIPS operational data.
type AwipsOperResponse struct {
	AwipsOperData AwipsOperData `json:"Cisco-IOS-XE-wireless-awips-oper:awips-oper-data"`
}

// AwipsOperData contains AWIPS operational data
type AwipsOperData struct {
	AwipsPerApInfo  []AwipsPerApInfo `json:"awips-per-ap-info"`
	AwipsDwldStatus AwipsDwldStatus  `json:"awips-dwld-status"`
}

// AwipsPerApInfo represents AWIPS information per AP
type AwipsPerApInfo struct {
	ApMac                 string `json:"ap-mac"`
	AwipsStatus           string `json:"awips-status"`
	AlarmCount            string `json:"alarm-count"`
	ForensicCaptureStatus string `json:"forensic-capture-status"`
}

// AwipsDwldStatus represents AWIPS download status
type AwipsDwldStatus struct {
	LastSuccessTimestamp string `json:"last-success-timestamp"`
	LastFailedTimestamp  string `json:"last-failed-timestamp"`
	NumOfFailureAttempts int    `json:"num-of-failure-attempts"`
}
