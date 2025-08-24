// Package model provides data models for AWIPS operational data.
package model

// AwipsOper  represents the AWIPS operational data.
type AwipsOper struct {
	AwipsOperData AwipsOperData `json:"Cisco-IOS-XE-wireless-awips-oper:awips-oper-data"`
}

// AwipsOperAwipsPerApInfo  represents the AWIPS per AP information.
type AwipsOperAwipsPerApInfo struct {
	AwipsPerApInfo []AwipsPerApInfo `json:"Cisco-IOS-XE-wireless-awips-oper:awips-per-ap-info"`
}

// AwipsOperAwipsDwldStatus  represents the AWIPS download status.
type AwipsOperAwipsDwldStatus struct {
	AwipsDwldStatus AwipsDwldStatus `json:"Cisco-IOS-XE-wireless-awips-oper:awips-dwld-status"`
}

type AwipsOperData struct {
	AwipsPerApInfo  []AwipsPerApInfo `json:"awips-per-ap-info"`
	AwipsDwldStatus AwipsDwldStatus  `json:"awips-dwld-status"`
}

type AwipsPerApInfo struct {
	ApMac                 string `json:"ap-mac"`
	AwipsStatus           string `json:"awips-status"`
	AlarmCount            string `json:"alarm-count"`
	ForensicCaptureStatus string `json:"forensic-capture-status"`
}

type AwipsDwldStatus struct {
	LastSuccessTimestamp string `json:"last-success-timestamp"`
	LastFailedTimestamp  string `json:"last-failed-timestamp"`
	NumOfFailureAttempts int    `json:"num-of-failure-attempts"`
}
