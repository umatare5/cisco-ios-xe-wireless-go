package wlan

// CiscoIOSXEWirelessWlanPolicyListEntriesPayload represents request structure for policy-list-entries endpoint.
type CiscoIOSXEWirelessWlanPolicyListEntriesPayload struct {
	PolicyListEntry PolicyListEntry `json:"Cisco-IOS-XE-wireless-wlan-cfg:policy-list-entry"`
}
