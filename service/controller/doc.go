// Package controller provides wireless controller management functionality for the Cisco IOS-XE Wireless Network Controller API.
//
// This package allows you to perform administrative operations on a Cisco Catalyst 9800 Wireless LAN Controller.
// It provides essential system management capabilities including controller restart operations and system-wide maintenance commands,
// and it reads the platform state that belongs to the controller itself rather than to its wireless namespace.
//
// WARNING: The reload operations are destructive and will cause complete wireless network service interruption.
//
// RESTCONF Endpoints:
// - Operational: /restconf/data/Cisco-IOS-XE-device-hardware-oper:device-hardware-data
// - RPC Operations: /restconf/operations/Cisco-IOS-XE-rpc:reload
//
// This package also carries controller-level operations from non-wireless modules such
// as Cisco-IOS-XE-rpc and Cisco-IOS-XE-device-hardware-oper, which no wireless service owns.
//
// YANG References:
// - Cisco-IOS-XE-device-hardware-oper.yang
// - Cisco-IOS-XE-rpc.yang
package controller
