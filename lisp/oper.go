// Package lisp provides LISP (Locator/Identifier Separation Protocol) operational data functionality for the Cisco Wireless Network Controller API.
package lisp

import (
	"context"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
)

const (
	// LispAgentOperBasePath defines the base path for LISP agent operational data endpoints.
	LispAgentOperBasePath = "/restconf/data/Cisco-IOS-XE-wireless-lisp-agent-oper:lisp-agent-oper-data"
	// LispAgentOperEndpoint defines the endpoint for LISP agent operational data.
	LispAgentOperEndpoint = LispAgentOperBasePath
	// LispAgentMemoryStatsEndpoint defines the endpoint for LISP agent memory statistics.
	LispAgentMemoryStatsEndpoint = LispAgentOperBasePath + "/lisp-agent-memory-stats"
	// LispWlcCapabilitiesEndpoint defines the endpoint for LISP WLC capabilities.
	LispWlcCapabilitiesEndpoint = LispAgentOperBasePath + "/lisp-wlc-capabilities"
	// LispApCapabilitiesEndpoint defines the endpoint for LISP AP capabilities.
	LispApCapabilitiesEndpoint = LispAgentOperBasePath + "/lisp-ap-capabilities"
)

// LispAgentOperResponse represents the response structure for LISP agent operational data.
type LispAgentOperResponse struct {
	CiscoIOSXEWirelessLispAgentOperLispAgentOperData struct {
		LispAgentMemoryStats LispAgentMemoryStats `json:"lisp-agent-memory-stats"`
		LispWlcCapabilities  LispWlcCapabilities  `json:"lisp-wlc-capabilities"`
		LispApCapabilities   []LispApCapability   `json:"lisp-ap-capabilities"`
	} `json:"Cisco-IOS-XE-wireless-lisp-agent-oper:lisp-agent-oper-data"`
}

// LispAgentMemoryStatsResponse represents the response structure for LISP agent memory statistics.
type LispAgentMemoryStatsResponse struct {
	LispAgentMemoryStats LispAgentMemoryStats `json:"Cisco-IOS-XE-wireless-lisp-agent-oper:lisp-agent-memory-stats"`
}

// LispWlcCapabilitiesResponse represents the response structure for LISP WLC capabilities.
type LispWlcCapabilitiesResponse struct {
	LispWlcCapabilities LispWlcCapabilities `json:"Cisco-IOS-XE-wireless-lisp-agent-oper:lisp-wlc-capabilities"`
}

// LispApCapabilitiesResponse represents the response structure for LISP AP capabilities.
type LispApCapabilitiesResponse struct {
	LispApCapabilities []LispApCapability `json:"Cisco-IOS-XE-wireless-lisp-agent-oper:lisp-ap-capabilities"`
}

// LispAgentMemoryStats represents LISP agent memory statistics including malloc and free operations.
type LispAgentMemoryStats struct {
	MallocPskBuf        string `json:"malloc-psk-buf"`
	FreePskBuf          string `json:"free-psk-buf"`
	MallocMapRegMsg     string `json:"malloc-map-reg-msg"`
	FreeMapRegMsg       string `json:"free-map-reg-msg"`
	MallocMapReqMsg     string `json:"malloc-map-req-msg"`
	FreeMapReqMsg       string `json:"free-map-req-msg"`
	MallocLispHaNode    string `json:"malloc-lisp-ha-node"`
	FreeLispHaNode      string `json:"free-lisp-ha-node"`
	MallocMapServerCtxt string `json:"malloc-map-server-ctxt"`
	FreeMapServerCtxt   string `json:"free-map-server-ctxt"`
}

// LispWlcCapabilities represents LISP WLC capabilities including fabric support.
type LispWlcCapabilities struct {
	FabricCapable bool `json:"fabric-capable"`
}

// LispApCapability represents LISP AP capability information including type and fabric support.
type LispApCapability struct {
	ApType        int  `json:"ap-type"`
	FabricCapable bool `json:"fabric-capable"`
}

// GetLispAgentOper retrieves LISP agent operational data.
func GetLispAgentOper(client *wnc.Client, ctx context.Context) (*LispAgentOperResponse, error) {
	var data LispAgentOperResponse
	if err := client.SendAPIRequest(ctx, LispAgentOperEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetLispAgentMemoryStats retrieves LISP agent memory statistics.
func GetLispAgentMemoryStats(client *wnc.Client, ctx context.Context) (*LispAgentMemoryStatsResponse, error) {
	var data LispAgentMemoryStatsResponse
	if err := client.SendAPIRequest(ctx, LispAgentMemoryStatsEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetLispWlcCapabilities retrieves LISP WLC capabilities.
func GetLispWlcCapabilities(client *wnc.Client, ctx context.Context) (*LispWlcCapabilitiesResponse, error) {
	var data LispWlcCapabilitiesResponse
	if err := client.SendAPIRequest(ctx, LispWlcCapabilitiesEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetLispApCapabilities retrieves LISP AP capabilities.
func GetLispApCapabilities(client *wnc.Client, ctx context.Context) (*LispApCapabilitiesResponse, error) {
	var data LispApCapabilitiesResponse
	if err := client.SendAPIRequest(ctx, LispApCapabilitiesEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
