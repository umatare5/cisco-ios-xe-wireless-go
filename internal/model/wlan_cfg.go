// Package model contains generated response structures for the Cisco WNC API.
package model

import "time"

// WlanCfgResponse represents the complete WLAN configuration response
type WlanCfgResponse struct {
	CiscoIOSXEWirelessWlanCfgWlanCfgData struct {
		WlanCfgEntries           WlanCfgEntries           `json:"wlan-cfg-entries"`
		WlanPolicies             WlanPolicies             `json:"wlan-policies"`
		PolicyListEntries        PolicyListEntries        `json:"policy-list-entries"`
		WirelessAaaPolicyConfigs WirelessAaaPolicyConfigs `json:"wireless-aaa-policy-configs"`
	} `json:"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data"`
}

// WlanCfgEntriesResponse represents the WLAN configuration entries response
type WlanCfgEntriesResponse struct {
	WlanCfgEntries WlanCfgEntries `json:"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-entries"`
}

// WlanPoliciesResponse represents the WLAN policies configuration response
type WlanPoliciesResponse struct {
	WlanPolicies WlanPolicies `json:"Cisco-IOS-XE-wireless-wlan-cfg:wlan-policies"`
}

// PolicyListEntriesResponse represents the policy list entries response
type PolicyListEntriesResponse struct {
	PolicyListEntries PolicyListEntries `json:"Cisco-IOS-XE-wireless-wlan-cfg:policy-list-entries"`
}

// WirelessAaaPolicyConfigsResponse represents the wireless AAA policy configurations response
type WirelessAaaPolicyConfigsResponse struct {
	WirelessAaaPolicyConfigs WirelessAaaPolicyConfigs `json:"Cisco-IOS-XE-wireless-wlan-cfg:wireless-aaa-policy-configs"`
}

// WlanCfgEntries represents WLAN configuration entries
type WlanCfgEntries struct {
	WlanCfgEntry []WlanCfgEntry `json:"wlan-cfg-entry"`
}

type WlanCfgEntry struct {
	WlanID      int    `json:"wlan-id"`
	ProfileName string `json:"profile-name"`
	Ssid        string `json:"ssid"`
	// Additional fields would be defined here
}

type WlanPolicies struct {
	WlanPolicy []WlanPolicy `json:"wlan-policy"`
}

type WlanPolicy struct {
	PolicyName string `json:"policy-name"`
	// Additional fields would be defined here
}

type PolicyListEntries struct {
	PolicyListEntry []PolicyListEntry `json:"policy-list-entry"`
}

type PolicyListEntry struct {
	PolicyName string `json:"policy-name"`
	// Additional fields would be defined here
}

type WirelessAaaPolicyConfigs struct {
	WirelessAaaPolicyConfig []WirelessAaaPolicyConfig `json:"wireless-aaa-policy-config"`
}

type WirelessAaaPolicyConfig struct {
	PolicyName string `json:"policy-name"`
	// Additional fields would be defined here
}

// WlanGlobalOperResponse represents WLAN global operational response data
type WlanGlobalOperResponse struct {
	WlanGlobalOperData struct {
		WlanOperList []WlanOperData `json:"wlan-oper-list"`
	} `json:"Cisco-IOS-XE-wireless-wlan-global-oper:wlan-global-oper-data"`
}

type WlanOperData struct {
	WlanID      int       `json:"wlan-id"`
	ProfileName string    `json:"profile-name"`
	Ssid        string    `json:"ssid"`
	State       string    `json:"state"`
	LastChange  time.Time `json:"last-change"`
	// Additional operational fields would be defined here
}
