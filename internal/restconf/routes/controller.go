package routes

// Controller Management RPC Operations
//
// These constants define the RESTCONF API paths for controller management
// operations such as system reload, restart, and administrative operations.
// Based on Cisco-IOS-XE-rpc YANG model specifications.

// Controller RPC Operations.
const (
	// ControllerReloadRPC defines the RPC for WNC controller reload operations
	// This operation follows the Cisco-IOS-XE-rpc:reload YANG model specification.
	ControllerReloadRPC = RESTCONFOperationsPath + "/Cisco-IOS-XE-rpc:reload"
)
