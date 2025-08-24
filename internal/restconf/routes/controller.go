package routes

// Controller Management RPC Endpoints
//
// These constants define the RESTCONF API endpoints for controller management
// operations such as system reload, restart, and administrative operations.
// Based on Cisco-IOS-XE-rpc YANG model specifications.

// Controller RPC Endpoints
const (
	// WNCReloadRPC defines the RPC endpoint for WNC controller reload operations
	// This endpoint follows the Cisco-IOS-XE-rpc:reload YANG model specification
	WNCReloadRPC = "/restconf/operations/Cisco-IOS-XE-rpc:reload"
)
