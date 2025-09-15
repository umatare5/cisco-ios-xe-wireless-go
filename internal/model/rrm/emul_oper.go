package rrm

// RRMEmulOper represents RRM emulation operational response data.
type RRMEmulOper struct {
	RRMEmulOperData RRMEmulOperData `json:"Cisco-IOS-XE-wireless-rrm-emul-oper:rrm-emul-oper-data"` // RRM operational data (Live: IOS-XE 17.12.5)
}

// RRMEmulOperData represents RRM emulation operational data container.
type RRMEmulOperData struct {
	RRMFraStats *RRMFraStats `json:"rrm-fra-stats,omitempty"` // RRM flexible radio statistics (Live: IOS-XE 17.12.5)
}
