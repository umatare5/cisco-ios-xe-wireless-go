package rrm

// RrmEmulOper represents RRM emulation operational response data.
type RrmEmulOper struct {
	RrmEmulOperData RrmEmulOperData `json:"Cisco-IOS-XE-wireless-rrm-emul-oper:rrm-emul-oper-data"` // RRM emulation operational data container
}

// RrmEmulOperData represents RRM emulation operational data container.
type RrmEmulOperData struct {
	RrmFraStats *RrmFraStats `json:"rrm-fra-stats,omitempty"` // Flexible Radio Assignment statistics
}
