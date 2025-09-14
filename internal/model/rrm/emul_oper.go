package rrm

// RrmEmulOper represents RRM emulation operational response data.
type RrmEmulOper struct {
	RrmEmulOperData RrmEmulOperData `json:"Cisco-IOS-XE-wireless-rrm-emul-oper:rrm-emul-oper-data"` // RRM operational data (Live: IOS-XE 17.12.5)
}

// RrmEmulOperData represents RRM emulation operational data container.
type RrmEmulOperData struct {
	RrmFraStats *RrmFraStats `json:"rrm-fra-stats,omitempty"` // RRM flexible radio statistics (Live: IOS-XE 17.12.5)
}
