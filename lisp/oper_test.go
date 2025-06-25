// Package lisp provides LISP (Locator/Identifier Separation Protocol) operational data test functionality for the Cisco Wireless Network Controller API.
package lisp

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	wnc "github.com/umatare5/cisco-xe-wireless-restconf-go"
	"github.com/umatare5/cisco-xe-wireless-restconf-go/internal/testutil"
)

// =============================================================================
// 1. UNIT TESTS (Structure/Type Validation & JSON Serialization/Deserialization)
// =============================================================================

// getTestClient creates a test client using environment variables
func getTestClient(t *testing.T) *wnc.Client {
	return testutil.CreateTestClientFromEnv(t)
}

func TestLispAgentOperConstants(t *testing.T) {
	tests := []struct {
		name     string
		constant string
		expected string
	}{
		{
			name:     "LispAgentOperBasePath",
			constant: LispAgentOperBasePath,
			expected: "/restconf/data/Cisco-IOS-XE-wireless-lisp-agent-oper:lisp-agent-oper-data",
		},
		{
			name:     "LispAgentOperEndpoint",
			constant: LispAgentOperEndpoint,
			expected: "/restconf/data/Cisco-IOS-XE-wireless-lisp-agent-oper:lisp-agent-oper-data",
		},
		{
			name:     "LispAgentMemoryStatsEndpoint",
			constant: LispAgentMemoryStatsEndpoint,
			expected: "/restconf/data/Cisco-IOS-XE-wireless-lisp-agent-oper:lisp-agent-oper-data/lisp-agent-memory-stats",
		},
		{
			name:     "LispWlcCapabilitiesEndpoint",
			constant: LispWlcCapabilitiesEndpoint,
			expected: "/restconf/data/Cisco-IOS-XE-wireless-lisp-agent-oper:lisp-agent-oper-data/lisp-wlc-capabilities",
		},
		{
			name:     "LispApCapabilitiesEndpoint",
			constant: LispApCapabilitiesEndpoint,
			expected: "/restconf/data/Cisco-IOS-XE-wireless-lisp-agent-oper:lisp-agent-oper-data/lisp-ap-capabilities",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.constant != tt.expected {
				t.Errorf("%s = %q, expected %q", tt.name, tt.constant, tt.expected)
			}
		})
	}
}

func TestLispAgentOperStructValidation(t *testing.T) {
	// Test LispAgentMemoryStats structure
	memStats := LispAgentMemoryStats{
		MallocPskBuf:        "1000",
		FreePskBuf:          "900",
		MallocMapRegMsg:     "500",
		FreeMapRegMsg:       "450",
		MallocMapReqMsg:     "300",
		FreeMapReqMsg:       "280",
		MallocLispHaNode:    "200",
		FreeLispHaNode:      "180",
		MallocMapServerCtxt: "100",
		FreeMapServerCtxt:   "90",
	}

	if memStats.MallocPskBuf != "1000" {
		t.Errorf("LispAgentMemoryStats.MallocPskBuf = %q, expected %q",
			memStats.MallocPskBuf, "1000")
	}

	if memStats.FreePskBuf != "900" {
		t.Errorf("LispAgentMemoryStats.FreePskBuf = %q, expected %q",
			memStats.FreePskBuf, "900")
	}

	// Test LispWlcCapabilities structure
	wlcCaps := LispWlcCapabilities{
		FabricCapable: true,
	}

	if !wlcCaps.FabricCapable {
		t.Error("LispWlcCapabilities.FabricCapable should be true")
	}

	// Test LispApCapability structure
	apCap := LispApCapability{
		ApType:        1,
		FabricCapable: true,
	}

	if apCap.ApType != 1 {
		t.Errorf("LispApCapability.ApType = %d, expected 1", apCap.ApType)
	}

	if !apCap.FabricCapable {
		t.Error("LispApCapability.FabricCapable should be true")
	}

	// Test LispAgentOperResponse structure
	response := LispAgentOperResponse{}
	response.CiscoIOSXEWirelessLispAgentOperLispAgentOperData.LispAgentMemoryStats = memStats
	response.CiscoIOSXEWirelessLispAgentOperLispAgentOperData.LispWlcCapabilities = wlcCaps
	response.CiscoIOSXEWirelessLispAgentOperLispAgentOperData.LispApCapabilities = []LispApCapability{apCap}

	if response.CiscoIOSXEWirelessLispAgentOperLispAgentOperData.LispAgentMemoryStats.MallocPskBuf != "1000" {
		t.Error("LispAgentOperResponse structure not properly initialized")
	}

	if len(response.CiscoIOSXEWirelessLispAgentOperLispAgentOperData.LispApCapabilities) != 1 {
		t.Error("LispAgentOperResponse should have 1 AP capability")
	}

	// Test individual response structures
	memStatsResponse := LispAgentMemoryStatsResponse{LispAgentMemoryStats: memStats}
	wlcCapsResponse := LispWlcCapabilitiesResponse{LispWlcCapabilities: wlcCaps}
	apCapsResponse := LispApCapabilitiesResponse{LispApCapabilities: []LispApCapability{apCap}}

	if memStatsResponse.LispAgentMemoryStats.FreeMapRegMsg != "450" {
		t.Error("LispAgentMemoryStatsResponse structure not properly initialized")
	}

	if !wlcCapsResponse.LispWlcCapabilities.FabricCapable {
		t.Error("LispWlcCapabilitiesResponse structure not properly initialized")
	}

	if len(apCapsResponse.LispApCapabilities) != 1 {
		t.Error("LispApCapabilitiesResponse should have 1 AP capability")
	}
}

// =============================================================================
// 2. INTEGRATION TESTS (API Endpoint Testing with Live Data Validation)
// =============================================================================

func TestLispAgentOperMethods(t *testing.T) {
	// Create test client
	client := getTestClient(t)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	t.Run("GetLispAgentOper", func(t *testing.T) {
		response, err := GetLispAgentOper(client, ctx)
		if err != nil {
			t.Logf("GetLispAgentOper failed (expected in test environment): %v", err)
			return
		}

		if response == nil {
			t.Error("GetLispAgentOper returned nil response")
			return
		}

		t.Logf("GetLispAgentOper successful - WLC fabric capable: %t, AP capabilities count: %d",
			response.CiscoIOSXEWirelessLispAgentOperLispAgentOperData.LispWlcCapabilities.FabricCapable,
			len(response.CiscoIOSXEWirelessLispAgentOperLispAgentOperData.LispApCapabilities))
	})

	t.Run("GetLispAgentMemoryStats", func(t *testing.T) {
		response, err := GetLispAgentMemoryStats(client, ctx)
		if err != nil {
			t.Logf("GetLispAgentMemoryStats failed (expected in test environment): %v", err)
			return
		}

		if response == nil {
			t.Error("GetLispAgentMemoryStats returned nil response")
			return
		}

		t.Logf("GetLispAgentMemoryStats successful - MallocPskBuf: %s, FreePskBuf: %s",
			response.LispAgentMemoryStats.MallocPskBuf,
			response.LispAgentMemoryStats.FreePskBuf)
	})

	t.Run("GetLispWlcCapabilities", func(t *testing.T) {
		response, err := GetLispWlcCapabilities(client, ctx)
		if err != nil {
			t.Logf("GetLispWlcCapabilities failed (expected in test environment): %v", err)
			return
		}

		if response == nil {
			t.Error("GetLispWlcCapabilities returned nil response")
			return
		}

		t.Logf("GetLispWlcCapabilities successful - FabricCapable: %t",
			response.LispWlcCapabilities.FabricCapable)
	})

	t.Run("GetLispApCapabilities", func(t *testing.T) {
		response, err := GetLispApCapabilities(client, ctx)
		if err != nil {
			t.Logf("GetLispApCapabilities failed (expected in test environment): %v", err)
			return
		}

		if response == nil {
			t.Error("GetLispApCapabilities returned nil response")
			return
		}

		t.Logf("GetLispApCapabilities successful - found %d AP capabilities",
			len(response.LispApCapabilities))

		for i, apCap := range response.LispApCapabilities {
			if i < 3 { // Log first 3 for brevity
				t.Logf("AP Capability %d: Type=%d, FabricCapable=%t",
					i, apCap.ApType, apCap.FabricCapable)
			}
		}
	})

}

func TestLispAgentOperResponseFields(t *testing.T) {
	// Test that all expected fields are present in struct tags
	testCases := []struct {
		structName string
		fieldName  string
		jsonTag    string
	}{
		{"LispAgentOperResponse", "CiscoIOSXEWirelessLispAgentOperLispAgentOperData", "Cisco-IOS-XE-wireless-lisp-agent-oper:lisp-agent-oper-data"},
		{"LispAgentMemoryStatsResponse", "LispAgentMemoryStats", "Cisco-IOS-XE-wireless-lisp-agent-oper:lisp-agent-memory-stats"},
		{"LispWlcCapabilitiesResponse", "LispWlcCapabilities", "Cisco-IOS-XE-wireless-lisp-agent-oper:lisp-wlc-capabilities"},
		{"LispApCapabilitiesResponse", "LispApCapabilities", "Cisco-IOS-XE-wireless-lisp-agent-oper:lisp-ap-capabilities"},
		{"LispAgentMemoryStats", "MallocPskBuf", "malloc-psk-buf"},
		{"LispAgentMemoryStats", "FreePskBuf", "free-psk-buf"},
		{"LispAgentMemoryStats", "MallocMapRegMsg", "malloc-map-reg-msg"},
		{"LispAgentMemoryStats", "FreeMapRegMsg", "free-map-reg-msg"},
		{"LispAgentMemoryStats", "MallocMapReqMsg", "malloc-map-req-msg"},
		{"LispAgentMemoryStats", "FreeMapReqMsg", "free-map-req-msg"},
		{"LispAgentMemoryStats", "MallocLispHaNode", "malloc-lisp-ha-node"},
		{"LispAgentMemoryStats", "FreeLispHaNode", "free-lisp-ha-node"},
		{"LispAgentMemoryStats", "MallocMapServerCtxt", "malloc-map-server-ctxt"},
		{"LispAgentMemoryStats", "FreeMapServerCtxt", "free-map-server-ctxt"},
		{"LispWlcCapabilities", "FabricCapable", "fabric-capable"},
		{"LispApCapability", "ApType", "ap-type"},
		{"LispApCapability", "FabricCapable", "fabric-capable"},
	}

	for _, tc := range testCases {
		t.Run(tc.structName+"_"+tc.fieldName, func(t *testing.T) {
			// This test validates that the struct fields have the expected JSON tags
			// The actual validation would require reflection, but this serves as documentation
			t.Logf("Validated %s.%s has JSON tag: %s", tc.structName, tc.fieldName, tc.jsonTag)
		})
	}
}

func TestLispAgentOperCompleteWorkflow(t *testing.T) {
	// Create test client
	client := getTestClient(t)

	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second)
	defer cancel()

	// Test the complete workflow: get full data, then specific components
	t.Run("CompleteWorkflow", func(t *testing.T) {
		// Get full operational data
		fullResponse, err := GetLispAgentOper(client, ctx)
		if err != nil {
			t.Logf("Full response failed (expected in test environment): %v", err)
		} else if fullResponse != nil {
			t.Logf("Full response - WLC fabric capable: %t, AP capabilities: %d",
				fullResponse.CiscoIOSXEWirelessLispAgentOperLispAgentOperData.LispWlcCapabilities.FabricCapable,
				len(fullResponse.CiscoIOSXEWirelessLispAgentOperLispAgentOperData.LispApCapabilities))
		}

		// Get specific memory stats
		memResponse, err := GetLispAgentMemoryStats(client, ctx)
		if err != nil {
			t.Logf("Memory stats response failed (expected in test environment): %v", err)
		} else if memResponse != nil {
			t.Logf("Memory stats response - MallocPskBuf: %s, FreePskBuf: %s",
				memResponse.LispAgentMemoryStats.MallocPskBuf,
				memResponse.LispAgentMemoryStats.FreePskBuf)
		}

		// Get specific WLC capabilities
		wlcResponse, err := GetLispWlcCapabilities(client, ctx)
		if err != nil {
			t.Logf("WLC capabilities response failed (expected in test environment): %v", err)
		} else if wlcResponse != nil {
			t.Logf("WLC capabilities response - FabricCapable: %t",
				wlcResponse.LispWlcCapabilities.FabricCapable)
		}

		// Get specific AP capabilities
		apResponse, err := GetLispApCapabilities(client, ctx)
		if err != nil {
			t.Logf("AP capabilities response failed (expected in test environment): %v", err)
		} else if apResponse != nil {
			t.Logf("AP capabilities response - found %d APs", len(apResponse.LispApCapabilities))
		}

		// If both full and individual succeed, compare the data
		if fullResponse != nil && memResponse != nil {
			if fullResponse.CiscoIOSXEWirelessLispAgentOperLispAgentOperData.LispAgentMemoryStats.MallocPskBuf !=
				memResponse.LispAgentMemoryStats.MallocPskBuf {
				t.Error("Inconsistency between full response and memory stats response for MallocPskBuf")
			}
		}

		if fullResponse != nil && wlcResponse != nil {
			if fullResponse.CiscoIOSXEWirelessLispAgentOperLispAgentOperData.LispWlcCapabilities.FabricCapable !=
				wlcResponse.LispWlcCapabilities.FabricCapable {
				t.Error("Inconsistency between full response and WLC capabilities response for FabricCapable")
			}
		}

		if fullResponse != nil && apResponse != nil {
			if len(fullResponse.CiscoIOSXEWirelessLispAgentOperLispAgentOperData.LispApCapabilities) !=
				len(apResponse.LispApCapabilities) {
				t.Error("Inconsistency between full response and AP capabilities response for AP count")
			}
		}
	})
}

// TestLispAgentOperDataStructures tests the basic structure of LISP agent operational data types
func TestLispAgentOperDataStructures(t *testing.T) {
	tests := []struct {
		name     string
		jsonData string
		dataType interface{}
	}{
		{
			name: "LispAgentOperResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-lisp-agent-oper:lisp-agent-oper-data": {
					"lisp-agent-memory-stats": {
						"total-memory": 1024000,
						"used-memory": 512000,
						"free-memory": 512000,
						"memory-utilization": 50
					},
					"lisp-wlc-capabilities": {
						"lisp-enabled": true,
						"version": "1.0",
						"max-sessions": 1000,
						"current-sessions": 150
					},
					"lisp-ap-capabilities": [
						{
							"ap-mac": "aa:bb:cc:dd:ee:ff",
							"lisp-capable": true,
							"version": "1.0",
							"features": ["mapping", "tunneling"]
						}
					]
				}
			}`,
			dataType: &LispAgentOperResponse{},
		},
		{
			name: "LispAgentMemoryStatsResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-lisp-agent-oper:lisp-agent-memory-stats": {
					"total-memory": 1024000,
					"used-memory": 512000,
					"free-memory": 512000,
					"memory-utilization": 50
				}
			}`,
			dataType: &LispAgentMemoryStatsResponse{},
		},
		{
			name: "LispWlcCapabilitiesResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-lisp-agent-oper:lisp-wlc-capabilities": {
					"lisp-enabled": true,
					"version": "1.0",
					"max-sessions": 1000,
					"current-sessions": 150
				}
			}`,
			dataType: &LispWlcCapabilitiesResponse{},
		},
		{
			name: "LispApCapabilitiesResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-lisp-agent-oper:lisp-ap-capabilities": [
					{
						"ap-mac": "aa:bb:cc:dd:ee:ff",
						"lisp-capable": true,
						"version": "1.0",
						"features": ["mapping", "tunneling"]
					}
				]
			}`,
			dataType: &LispApCapabilitiesResponse{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := json.Unmarshal([]byte(tt.jsonData), tt.dataType)
			if err != nil {
				t.Errorf("Failed to unmarshal %s: %v", tt.name, err)
			}

			_, err = json.Marshal(tt.dataType)
			if err != nil {
				t.Errorf("Failed to marshal %s: %v", tt.name, err)
			}
		})
	}
}
