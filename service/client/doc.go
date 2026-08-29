// Package client provides wireless client operational operations for Cisco IOS-XE wireless controllers.
//
// This package allows you to monitor wireless client operational data, statistics, and mobility information.
// It provides methods for client monitoring, traffic statistics retrieval, and policy data access across wireless infrastructures.
//
// WARNING: The deauthentication operations drop a client's session and interrupt its traffic.
// The controller answers them the same way whether or not the identifier matched a client.
//
// RESTCONF Endpoints:
// - Operational: /restconf/data/Cisco-IOS-XE-wireless-client-oper:client-oper-data
// - RPC Operations: /restconf/operations/Cisco-IOS-XE-wireless-client-rpc:apf-ms-delete-all
//
// YANG References:
// - Cisco-IOS-XE-wireless-client-oper.yang (17.12.1, 17.15.1, 17.18.1)
// - Cisco-IOS-XE-wireless-client-rpc.yang (17.15.1, 17.18.1) — the RPC is absent on 17.12
package client
