// Package controller provides WNC controller configuration and operation data structures.
package controller

// WNC Controller Reload RPC Payload Structures

// WNCReloadRPCPayload represents the RPC payload for WNC controller reload operations
// This follows the Cisco-IOS-XE-rpc:reload YANG model specification.
type WNCReloadRPCPayload struct {
	Input WNCReloadRPCInput `json:"input"`
}

// WNCReloadRPCInput represents the input parameters for WNC reload RPC
// Based on Cisco-IOS-XE-rpc:reload specification.
type WNCReloadRPCInput struct {
	// Force indicates whether to force a restart even if there is unsaved config
	// This is optional and defaults to false if not specified
	Force *bool `json:"force,omitempty"`

	// Reason provides a description for the reload operation
	// This field is required and should not be empty
	Reason string `json:"reason,omitempty"`
}
