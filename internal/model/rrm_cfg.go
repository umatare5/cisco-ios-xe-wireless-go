// Package model contains generated response structures for the Cisco WNC API.
package model

import "time"

// RrmCfgResponse represents RRM configuration response data
type RrmCfgResponse struct {
	CiscoIOSXEWirelessRrmCfgRrmCfgData struct {
		RrmPolicies RrmPolicies `json:"rrm-policies"`
	} `json:"Cisco-IOS-XE-wireless-rrm-cfg:rrm-cfg-data"`
}

type RrmPolicies struct {
	RrmPolicy []RrmPolicy `json:"rrm-policy"`
}

type RrmPolicy struct {
	PolicyName string `json:"policy-name"`
	// Additional fields would be defined here
}

// RrmOperResponse represents RRM operational response data
type RrmOperResponse struct {
	CiscoIOSXEWirelessRrmOperRrmOperData struct {
		RrmOperData []RrmOperData `json:"rrm-oper-data"`
	} `json:"Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data"`
}

type RrmOperData struct {
	PolicyName    string    `json:"policy-name"`
	State         string    `json:"state"`
	LastChange    time.Time `json:"last-change"`
	ChannelAssign struct {
		Channel      int       `json:"channel"`
		PowerLevel   int       `json:"power-level"`
		LastAssigned time.Time `json:"last-assigned"`
	} `json:"channel-assign"`
	// Additional operational fields would be defined here
}

// RrmGlobalOperResponse represents RRM global operational response data
type RrmGlobalOperResponse struct {
	RrmGlobalOperData struct {
		RrmStats RrmStats `json:"rrm-stats"`
	} `json:"Cisco-IOS-XE-wireless-rrm-global-oper:rrm-global-oper-data"`
}

type RrmStats struct {
	ChannelChanges   int       `json:"channel-changes"`
	PowerAdjustments int       `json:"power-adjustments"`
	LastOptimization time.Time `json:"last-optimization"`
	// Additional global stats would be defined here
}

// RrmEmulOperResponse represents RRM emulation operational response data
type RrmEmulOperResponse struct {
	RrmEmulOperData struct {
		EmulationStats EmulationStats `json:"emulation-stats"`
	} `json:"Cisco-IOS-XE-wireless-rrm-emul-oper:rrm-emul-oper-data"`
}

type EmulationStats struct {
	EmulationEnabled bool      `json:"emulation-enabled"`
	LastEmulation    time.Time `json:"last-emulation"`
	// Additional emulation stats would be defined here
}
