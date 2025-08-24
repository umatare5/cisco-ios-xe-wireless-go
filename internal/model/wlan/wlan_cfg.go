// Package model provides data models for WLAN configuration data.
package model

import "time"

// WlanCfg  represents the complete WLAN configuration
type WlanCfg struct {
	WlanCfgData *WlanCfgData `json:"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data"`
}

// WlanCfgWlanCfgEntries  represents the WLAN configuration entries.
type WlanCfgWlanCfgEntries struct {
	WlanCfgEntries *WlanCfgEntries `json:"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-entries,omitempty"`
}

// WlanCfgWlanPolicies  represents the WLAN policies.
type WlanCfgWlanPolicies struct {
	WlanPolicies *WlanPolicies `json:"Cisco-IOS-XE-wireless-wlan-cfg:wlan-policies,omitempty"`
}

// WlanCfgPolicyListEntries  represents the policy list entries.
type WlanCfgPolicyListEntries struct {
	PolicyListEntries *PolicyListEntries `json:"Cisco-IOS-XE-wireless-wlan-cfg:policy-list-entries,omitempty"`
}

// WlanCfgWirelessAaaPolicyConfigs  represents the wireless AAA policy configurations.
type WlanCfgWirelessAaaPolicyConfigs struct {
	WirelessAaaPolicyConfigs *WirelessAaaPolicyConfigs `json:"Cisco-IOS-XE-wireless-wlan-cfg:wireless-aaa-policy-configs,omitempty"`
}

type WlanCfgData struct {
	WlanCfgEntries           *WlanCfgEntries           `json:"wlan-cfg-entries,omitempty"`
	WlanPolicies             *WlanPolicies             `json:"wlan-policies,omitempty"`
	PolicyListEntries        *PolicyListEntries        `json:"policy-list-entries,omitempty"`
	WirelessAaaPolicyConfigs *WirelessAaaPolicyConfigs `json:"wireless-aaa-policy-configs,omitempty"`
}

// WlanCfgEntries  represents the WLAN configuration entries response
type WlanCfgEntries struct {
	WlanCfgEntry []WlanCfgEntry `json:"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-entries"`
}

// WlanCfgEntry  represents a single WLAN configuration entry response - REMOVED
// This type was causing recursive definition issues

// WlanPolicies  represents the WLAN policies configuration response
type WlanPolicies struct {
	WlanPolicy []WlanPolicy `json:"Cisco-IOS-XE-wireless-wlan-cfg:wlan-policies"`
}

// PolicyListEntries  represents the policy list entries response
type PolicyListEntries struct {
	PolicyListEntry []PolicyListEntry `json:"policy-list-entry"`
}

// WirelessAaaPolicyConfigs  represents the wireless AAA policy configurations response
type WirelessAaaPolicyConfigs struct {
	WirelessAaaPolicyConfig []WirelessAaaPolicyConfig `json:"Cisco-IOS-XE-wireless-wlan-cfg:wireless-aaa-policy-configs"`
}

type WlanCfgEntry struct {
	WlanID                 int                `json:"wlan-id"`
	ProfileName            string             `json:"profile-name"`
	AuthKeyMgmtPsk         bool               `json:"auth-key-mgmt-psk,omitempty"`
	AuthKeyMgmtDot1x       bool               `json:"auth-key-mgmt-dot1x,omitempty"`
	AuthKeyMgmtDot1xSha256 bool               `json:"auth-key-mgmt-dot1x-sha256,omitempty"`
	PSK                    string             `json:"psk,omitempty"`
	PSKType                string             `json:"psk-type,omitempty"`
	FTMode                 string             `json:"ft-mode,omitempty"`
	PMFOptions             string             `json:"pmf-options,omitempty"`
	WPA2Enabled            bool               `json:"wpa2-enabled,omitempty"`
	WPA3Enabled            bool               `json:"wpa3-enabled,omitempty"`
	LoadBalance            bool               `json:"load-balance,omitempty"`
	Wlan11kNeighList       bool               `json:"wlan-11k-neigh-list,omitempty"`
	MulticastBufferValue   int                `json:"multicast-buffer-value,omitempty"`
	ApfVapIDData           *ApfVapIDData      `json:"apf-vap-id-data,omitempty"`
	ApfVap80211vData       *ApfVap80211vData  `json:"apf-vap-802-11v-data,omitempty"`
	MdnsSDMode             string             `json:"mdns-sd-mode,omitempty"`
	WlanRadioPolicies      *WlanRadioPolicies `json:"wlan-radio-policies,omitempty"`
	ClientSteering         bool               `json:"client-steering,omitempty"`
	WepKeyIndex            int                `json:"wep-key-index,omitempty"`
}

type ApfVapIDData struct {
	SSID       string `json:"ssid"`
	WlanStatus bool   `json:"wlan-status"`
}

type ApfVap80211vData struct {
	Dot11vDms bool `json:"dot11v-dms"`
}

type WlanRadioPolicies struct {
	WlanRadioPolicy []WlanRadioPolicy `json:"wlan-radio-policy"`
}

type WlanRadioPolicy struct {
	Band string `json:"band"`
}

type WlanPolicy struct {
	PolicyName string `json:"policy-name"`
	// Additional fields would be defined here
}

type PolicyListEntry struct {
	TagName      string        `json:"tag-name,omitempty"`
	Description  string        `json:"description,omitempty"`
	WLANPolicies *WLANPolicies `json:"wlan-policies,omitempty"`
}

// WLANPolicies  represents the container for WLAN policy mappings
type WLANPolicies struct {
	WLANPolicy []WLANPolicyMap `json:"wlan-policy,omitempty"`
}

// WLANPolicyMap  represents a WLAN to policy profile mapping
type WLANPolicyMap struct {
	WLANProfileName   string `json:"wlan-profile-name"`
	PolicyProfileName string `json:"policy-profile-name"`
}

type WirelessAaaPolicyConfig struct {
	PolicyName string `json:"policy-name"`
	// Additional fields would be defined here
}

// WlanGlobalOper represents WLAN global operational data structure.
type WlanGlobalOper struct {
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
