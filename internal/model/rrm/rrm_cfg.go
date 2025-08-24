// Package model provides data models for RRM configuration data.
package model

import "time"

// RrmCfg  represents RRM configuration response data
type RrmCfg struct {
	CiscoIOSXEWirelessRrmCfgRrmCfgData struct {
		RrmPolicies RrmPolicies `json:"rrm-policies"`
	} `json:"Cisco-IOS-XE-wireless-rrm-cfg:rrm-cfg-data"`
}

// RrmCfgRrmPolicies  represents the corresponding data structure.
type RrmCfgRrmPolicies struct {
	RrmPolicies RrmPolicies `json:"Cisco-IOS-XE-wireless-rrm-cfg:rrm-policies"`
}

type RrmPolicies struct {
	RrmPolicy []RrmPolicy `json:"rrm-policy"`
}

type RrmPolicy struct {
	PolicyName string `json:"policy-name"`
	// Additional fields would be defined here
}

// RrmOper  represents RRM operational response data
type RrmOper struct {
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

// RrmGlobalOper  represents RRM global operational response data
type RrmGlobalOper struct {
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

// RrmEmulOper  represents RRM emulation operational response data
type RrmEmulOper struct {
	RrmEmulOperData struct {
		EmulationStats EmulationStats `json:"emulation-stats"`
	} `json:"Cisco-IOS-XE-wireless-rrm-emul-oper:rrm-emul-oper-data"`
}

type EmulationStats struct {
	EmulationEnabled bool      `json:"emulation-enabled"`
	LastEmulation    time.Time `json:"last-emulation"`
	// Additional emulation stats would be defined here
}
