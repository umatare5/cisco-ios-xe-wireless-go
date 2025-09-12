package awips

// AwipsOper represents the AWIPS operational data.
type AwipsOper struct {
	AwipsOperData AwipsOperData `json:"Cisco-IOS-XE-wireless-awips-oper:awips-oper-data"`
}

// AwipsOperAwipsPerApInfo represents the AWIPS per AP information.
type AwipsOperAwipsPerApInfo struct {
	AwipsPerApInfo []AwipsPerApInfo `json:"Cisco-IOS-XE-wireless-awips-oper:awips-per-ap-info"`
}

// AwipsOperAwipsDwldStatus represents the AWIPS download status.
type AwipsOperAwipsDwldStatus struct {
	AwipsDwldStatus AwipsDwldStatus `json:"Cisco-IOS-XE-wireless-awips-oper:awips-dwld-status"`
}

// AwipsOperData represents complete AWIPS operational data.
type AwipsOperData struct {
	AwipsPerApInfo      []AwipsPerApInfo     `json:"awips-per-ap-info"`
	AwipsDwldStatus     AwipsDwldStatus      `json:"awips-dwld-status"`
	AwipsApDwldStatus   []AwipsApDwldStatus  `json:"awips-ap-dwld-status"`             // (YANG: IOS-XE 17.18.1+)
	AwipsPerSignStats   []AwipsPerSignStats  `json:"awips-per-sign-stats,omitempty"`   // AWIPS statistics per signature (YANG: IOS-XE 17.18.1+)
	AwipsGlobStats      *AwipsGlobStats      `json:"awips-glob-stats,omitempty"`       // AWIPS global statistics (YANG: IOS-XE 17.18.1+)
	AwipsDwldStatusWncd *AwipsDwldStatusWncd `json:"awips-dwld-status-wncd,omitempty"` // AWIPS internal bookkeeping (YANG: IOS-XE 17.18.1+)
}

// AwipsPerApInfo represents AWIPS status and per AP alarm statistics.
type AwipsPerApInfo struct {
	ApMac                 string `json:"ap-mac"`                  // AP MAC address
	AwipsStatus           string `json:"awips-status"`            // AWIPS status
	AlarmCount            string `json:"alarm-count"`             // Alarm counter
	ForensicCaptureStatus string `json:"forensic-capture-status"` // AWIPS forensic capture status
}

// AwipsDwldStatus represents AWIPS file download status.
type AwipsDwldStatus struct {
	LastSuccessTimestamp string `json:"last-success-timestamp"`  // Last file download success timestamp
	LastFailedTimestamp  string `json:"last-failed-timestamp"`   // Last file download failure timestamp
	NumOfFailureAttempts int    `json:"num-of-failure-attempts"` // Number of times file download failed
	LastFailureReason    int    `json:"last-failure-reason"`     // Last failure reason
	WlcVersion           string `json:"wlc-version"`             // WLC version information
	MaxFileVer           int    `json:"max-file-ver"`            // Maximum supported file version information
	LatestFileVersion    int    `json:"latest-file-version"`     // File version information
	DownloadStatus       string `json:"download-status"`         // File download status
	FileHash             string `json:"file-hash"`               // File content hash value
}

// AwipsApDwldStatus represents AWIPS per AP file download status.
type AwipsApDwldStatus struct {
	ApMac       string `json:"ap-mac"`       // AP MAC address
	DwldStatus  string `json:"dwld-status"`  // File download status
	FileVersion int    `json:"file-version"` // File version at AP
	FileHash    string `json:"file-hash"`    // File content hash value
}

// AwipsPerSignStats represents AWIPS statistics per signature.
type AwipsPerSignStats struct {
	SignatureID     int    `json:"signature-id"`     // Signature ID (YANG: IOS-XE 17.18.1+)
	SignatureString string `json:"signature-string"` // Signature description (YANG: IOS-XE 17.18.1+)
	AlarmCounter    int    `json:"alarm-counter"`    // Number of alarms (YANG: IOS-XE 17.18.1+)
}

// AwipsGlobStats represents AWIPS global statistics.
type AwipsGlobStats struct {
	TimeoutInSeconds   int    `json:"timeout-in-seconds"`    // Stats timeout interval in seconds (YANG: IOS-XE 17.18.1+)
	CurrHourTimestamp  string `json:"curr-hour-timestamp"`   // Timestamp of current hour (YANG: IOS-XE 17.18.1+)
	LastClearTimestamp string `json:"last-clear-timestamp"`  // Timestamp of last clear (YANG: IOS-XE 17.18.1+)
	CurrHourAlarmCount int    `json:"curr-hour-alarm-count"` // Current hour alarm count (YANG: IOS-XE 17.18.1+)
}

// AwipsDwldStatusWncd represents AWIPS internal bookkeeping of file download status.
type AwipsDwldStatusWncd struct {
	WlcVersion        string `json:"wlc-version"`         // File version at WLC (YANG: IOS-XE 17.18.1+)
	LatestFileVersion int    `json:"latest-file-version"` // File version of the latest download (YANG: IOS-XE 17.18.1+)
	DownloadStatus    bool   `json:"download-status"`     // File download status (YANG: IOS-XE 17.18.1+)
	FileDirectory     string `json:"file-directory"`      // File directory (YANG: IOS-XE 17.18.1+)
	FileName          string `json:"file-name"`           // File name (YANG: IOS-XE 17.18.1+)
	FileHash          string `json:"file-hash"`           // File hash value (YANG: IOS-XE 17.18.1+)
}
