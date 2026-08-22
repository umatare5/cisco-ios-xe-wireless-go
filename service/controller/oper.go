package controller

import "time"

// CiscoIOSXEDeviceHardwareOperBootTime represents the controller boot instant.
//
// The module is Cisco-IOS-XE-device-hardware-oper rather than a wireless one: the controller as a
// system is in this service's charter, as Cisco-IOS-XE-rpc:reload is.
//
// The instant is a pointer because a read answered with no body decodes into a zero value of this
// type, and a non-pointer field would make that zero indistinguishable from an instant the
// controller reported.
type CiscoIOSXEDeviceHardwareOperBootTime struct {
	BootTime *time.Time `json:"Cisco-IOS-XE-device-hardware-oper:boot-time,omitempty"` // Instant the controller last booted (Live: IOS-XE 17.12.7a)
}
