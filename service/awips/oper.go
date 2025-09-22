package awips

// AWIPSOper represents the AWIPS operational data.
type AWIPSOper struct {
	CiscoIOSXEWirelessAWIPSOperData struct {
		AWIPSPerApInfo      []AWIPSPerApInfo     `json:"awips-per-ap-info"`                // Per AP AWIPS information (Live: IOS-XE 17.12.5)
		AWIPSDwldStatus     AWIPSDwldStatus      `json:"awips-dwld-status"`                // AWIPS file download status (Live: IOS-XE 17.12.5)
		AWIPSApDwldStatus   []AWIPSApDwldStatus  `json:"awips-ap-dwld-status"`             // Per AP AWIPS file download status (Live: IOS-XE 17.12.5)
		AWIPSPerSignStats   []AWIPSPerSignStats  `json:"awips-per-sign-stats,omitempty"`   // AWIPS statistics per signature (YANG: IOS-XE 17.18.1)
		AWIPSGlobStats      *AWIPSGlobStats      `json:"awips-glob-stats,omitempty"`       // AWIPS global statistics (YANG: IOS-XE 17.18.1)
		AWIPSDwldStatusWncd *AWIPSDwldStatusWncd `json:"awips-dwld-status-wncd,omitempty"` // AWIPS internal bookkeeping (YANG: IOS-XE 17.18.1)
	} `json:"Cisco-IOS-XE-wireless-awips-oper:awips-oper-data"` // AWIPS operational data (Live: IOS-XE 17.12.5)
}

// AWIPSOperAWIPSPerApInfo represents the AWIPS per AP information.
type AWIPSOperAWIPSPerApInfo struct {
	AWIPSPerApInfo []AWIPSPerApInfo `json:"Cisco-IOS-XE-wireless-awips-oper:awips-per-ap-info"`
}

// AWIPSOperAWIPSDwldStatus represents the AWIPS download status.
type AWIPSOperAWIPSDwldStatus struct {
	AWIPSDwldStatus AWIPSDwldStatus `json:"Cisco-IOS-XE-wireless-awips-oper:awips-dwld-status"`
}

// AWIPSOperAWIPSApDwldStatus represents the AWIPS per AP download status.
type AWIPSOperAWIPSApDwldStatus struct {
	AWIPSApDwldStatus []AWIPSApDwldStatus `json:"Cisco-IOS-XE-wireless-awips-oper:awips-ap-dwld-status"`
}

// AWIPSOperAWIPSPerSignStats represents the AWIPS per signature statistics.
type AWIPSOperAWIPSPerSignStats struct {
	AWIPSPerSignStats []AWIPSPerSignStats `json:"Cisco-IOS-XE-wireless-awips-oper:awips-per-sign-stats"`
}

// AWIPSOperAWIPSGlobStats represents the AWIPS global statistics.
type AWIPSOperAWIPSGlobStats struct {
	AWIPSGlobStats *AWIPSGlobStats `json:"Cisco-IOS-XE-wireless-awips-oper:awips-glob-stats"`
}

// AWIPSOperAWIPSDwldStatusWncd represents the AWIPS download status for WNCD.
type AWIPSOperAWIPSDwldStatusWncd struct {
	AWIPSDwldStatusWncd *AWIPSDwldStatusWncd `json:"Cisco-IOS-XE-wireless-awips-oper:awips-dwld-status-wncd"`
}

// AWIPSPerApInfo represents AWIPS status and per AP alarm statistics.
type AWIPSPerApInfo struct {
	ApMAC                 string `json:"ap-mac"`                  // AP MAC address (Live: IOS-XE 17.12.5)
	AWIPSStatus           string `json:"awips-status"`            // AWIPS status (Live: IOS-XE 17.12.5)
	AlarmCount            string `json:"alarm-count"`             // Alarm counter (Live: IOS-XE 17.12.5)
	ForensicCaptureStatus string `json:"forensic-capture-status"` // AWIPS forensic capture status (Live: IOS-XE 17.12.5)
}

// AWIPSDwldStatus represents AWIPS file download status.
type AWIPSDwldStatus struct {
	LastSuccessTimestamp string `json:"last-success-timestamp"`  // Last file download success timestamp (Live: IOS-XE 17.12.5)
	LastFailedTimestamp  string `json:"last-failed-timestamp"`   // Last file download failure timestamp (Live: IOS-XE 17.12.5)
	NumOfFailureAttempts int    `json:"num-of-failure-attempts"` // Number of times file download failed (Live: IOS-XE 17.12.5)
	LastFailureReason    int    `json:"last-failure-reason"`     // Last failure reason (Live: IOS-XE 17.12.5)
	WlcVersion           string `json:"wlc-version"`             // WLC version information (Live: IOS-XE 17.12.5)
	MaxFileVer           int    `json:"max-file-ver"`            // Maximum supported file version info (Live: IOS-XE 17.12.5)
	LatestFileVersion    int    `json:"latest-file-version"`     // File version information (Live: IOS-XE 17.12.5)
	DownloadStatus       string `json:"download-status"`         // File download status (Live: IOS-XE 17.12.5)
	FileHash             string `json:"file-hash"`               // File content hash value (Live: IOS-XE 17.12.5)
}

// AWIPSApDwldStatus represents AWIPS per AP file download status.
type AWIPSApDwldStatus struct {
	ApMAC       string `json:"ap-mac"`       // AP MAC address (Live: IOS-XE 17.12.5)
	DwldStatus  string `json:"dwld-status"`  // File download status (Live: IOS-XE 17.12.5)
	FileVersion int    `json:"file-version"` // File version at AP (Live: IOS-XE 17.12.5)
	FileHash    string `json:"file-hash"`    // File content hash value (Live: IOS-XE 17.12.5)
}

// AWIPSPerSignStats represents AWIPS statistics per signature.
type AWIPSPerSignStats struct {
	SignatureID     int    `json:"signature-id"`     // Signature ID (YANG: IOS-XE 17.18.1)
	SignatureString string `json:"signature-string"` // Signature description (YANG: IOS-XE 17.18.1)
	AlarmCounter    int    `json:"alarm-counter"`    // Number of alarms (YANG: IOS-XE 17.18.1)
}

// AWIPSGlobStats represents AWIPS global statistics.
type AWIPSGlobStats struct {
	TimeoutInSeconds   int    `json:"timeout-in-seconds"`    // Stats timeout interval in seconds (YANG: IOS-XE 17.18.1)
	CurrHourTimestamp  string `json:"curr-hour-timestamp"`   // Timestamp of current hour (YANG: IOS-XE 17.18.1)
	LastClearTimestamp string `json:"last-clear-timestamp"`  // Timestamp of last clear (YANG: IOS-XE 17.18.1)
	CurrHourAlarmCount int    `json:"curr-hour-alarm-count"` // Number of alarms in current hour (YANG: IOS-XE 17.18.1)
}

// AWIPSDwldStatusWncd represents AWIPS internal bookkeeping of file download status.
type AWIPSDwldStatusWncd struct {
	WlcVersion        string `json:"wlc-version"`         // File version at WLC (YANG: IOS-XE 17.18.1)
	LatestFileVersion int    `json:"latest-file-version"` // File version of the latest download (YANG: IOS-XE 17.18.1)
	DownloadStatus    bool   `json:"download-status"`     // File download status (YANG: IOS-XE 17.18.1)
	FileDirectory     string `json:"file-directory"`      // File directory (YANG: IOS-XE 17.18.1)
	FileName          string `json:"file-name"`           // File name (YANG: IOS-XE 17.18.1)
	FileHash          string `json:"file-hash"`           // File hash value (YANG: IOS-XE 17.18.1)
}
