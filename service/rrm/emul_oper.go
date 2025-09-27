package rrm

// CiscoIOSXEWirelessRRMEmulOper represents RRM emulation operational response data.
type CiscoIOSXEWirelessRRMEmulOper struct {
	RRMEmulOperData struct {
		RRMFraStats *RRMFraStats `json:"rrm-fra-stats,omitempty"` // RRM flexible radio statistics (Live: IOS-XE 17.12.5)
	} `json:"Cisco-IOS-XE-wireless-rrm-emul-oper:rrm-emul-oper-data"` // RRM operational data (Live: IOS-XE 17.12.5)
}

// CiscoIOSXEWirelessRRMEmulOperRRMFraStats represents the RRM flexible radio assignment statistics (YANG: IOS-XE 17.12.1).
type CiscoIOSXEWirelessRRMEmulOperRRMFraStats struct {
	RRMFraStats *RRMFraStats `json:"Cisco-IOS-XE-wireless-rrm-emul-oper:rrm-fra-stats"`
}
