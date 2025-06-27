// Package mesh provides mesh networking test functionality for the Cisco Wireless Network Controller API.
package mesh

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil"
)

// =============================================================================
// 1. UNIT TESTS (Structure/Type Validation & JSON Serialization/Deserialization)
// =============================================================================

// getTestClient creates a test client using environment variables
func getTestClient(t *testing.T) *wnc.Client {
	return testutil.CreateTestClientFromEnv(t)
}

// MeshGlobalOperTestDataCollector holds test data for mesh global operation functions
type MeshGlobalOperTestDataCollector struct {
	Data map[string]interface{} `json:"mesh_global_oper_test_data"`
}

// =============================================================================
// 2. INTEGRATION TESTS (API Endpoint Testing with Live Data Validation)
// =============================================================================

func TestMeshGlobalOperGetMeshGlobalOper(t *testing.T) {
	client := getTestClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), testutil.DefaultTestTimeout)
	defer cancel()

	result, err := GetMeshGlobalOper(client, ctx)
	if err != nil {
		t.Fatalf("GetMeshGlobalOper failed: %v", err)
	}

	if result == nil {
		t.Fatal("GetMeshGlobalOper returned nil result")
	}

	// Save result to JSON file
	filename := fmt.Sprintf("mesh_global_oper_data_%d.json", time.Now().Unix())
	if err := testutil.SaveTestDataToFile(filename, result); err != nil {
		t.Logf("Warning: Failed to save data to %s: %v", filename, err)
	} else {
		t.Logf("Data saved to %s", filename)
	}

	t.Logf("GetMeshGlobalOper successful, collected mesh global operational data")
}

func TestMeshGlobalOperGetMeshGlobalStats(t *testing.T) {
	client := getTestClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), testutil.DefaultTestTimeout)
	defer cancel()

	result, err := GetMeshGlobalStats(client, ctx)
	if err != nil {
		t.Fatalf("GetMeshGlobalStats failed: %v", err)
	}

	if result == nil {
		t.Fatal("GetMeshGlobalStats returned nil result")
	}

	// Save result to JSON file
	filename := fmt.Sprintf("mesh_global_stats_data_%d.json", time.Now().Unix())
	if err := testutil.SaveTestDataToFile(filename, result); err != nil {
		t.Logf("Warning: Failed to save data to %s: %v", filename, err)
	} else {
		t.Logf("Data saved to %s", filename)
	}

	t.Logf("GetMeshGlobalStats successful")
}

func TestMeshGlobalOperGetMeshApTreeData(t *testing.T) {
	client := getTestClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), testutil.DefaultTestTimeout)
	defer cancel()

	result, err := GetMeshApTreeData(client, ctx)
	if err != nil {
		t.Fatalf("GetMeshApTreeData failed: %v", err)
	}

	if result == nil {
		t.Fatal("GetMeshApTreeData returned nil result")
	}

	// Save result to JSON file
	filename := fmt.Sprintf("mesh_ap_tree_data_%d.json", time.Now().Unix())
	if err := testutil.SaveTestDataToFile(filename, result); err != nil {
		t.Logf("Warning: Failed to save data to %s: %v", filename, err)
	} else {
		t.Logf("Data saved to %s", filename)
	}

	t.Logf("GetMeshApTreeData successful")
}

func TestMeshGlobalOperCollectAllData(t *testing.T) {
	client := getTestClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), testutil.DefaultTestTimeout)
	defer cancel()

	allData := make(map[string]interface{})

	// Collect data from all mesh global operational endpoints
	tests := []struct {
		name string
		fn   func() (interface{}, error)
	}{
		{"GetMeshGlobalOper", func() (interface{}, error) { return GetMeshGlobalOper(client, ctx) }},
		{"GetMeshGlobalStats", func() (interface{}, error) { return GetMeshGlobalStats(client, ctx) }},
		{"GetMeshApTreeData", func() (interface{}, error) { return GetMeshApTreeData(client, ctx) }},
	}

	for _, test := range tests {
		result, err := test.fn()
		if err != nil {
			t.Logf("Warning: %s failed: %v", test.name, err)
			allData[test.name] = map[string]string{"error": err.Error()}
		} else {
			allData[test.name] = result
			t.Logf("%s successful", test.name)
		}
	}

	// Save all collected data to a comprehensive JSON file
	filename := fmt.Sprintf("mesh_global_oper_comprehensive_data_%d.json", time.Now().Unix())
	if err := testutil.SaveTestDataToFile(filename, allData); err != nil {
		t.Logf("Warning: Failed to save comprehensive data to %s: %v", filename, err)
	} else {
		t.Logf("Comprehensive mesh global operational data saved to %s", filename)
	}
}

// TestMeshGlobalOperDataStructures tests the basic structure of mesh global operational data types
func TestMeshGlobalOperDataStructures(t *testing.T) {
	tests := []struct {
		name     string
		jsonData string
		dataType interface{}
	}{
		{
			name: "MeshGlobalOperResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-mesh-global-oper:mesh-global-oper-data": {
					"mesh-global-stats": {
						"total-mesh-aps": 25,
						"active-mesh-aps": 23,
						"inactive-mesh-aps": 2,
						"mesh-tree-depth": 3,
						"mesh-convergence-time": 120
					    },
					"mesh-ap-tree-data": [
						{
							"ap-mac": "aa:bb:cc:dd:ee:ff",
							"parent-mac": "11:22:33:44:55:66",
							"children-count": 2,
							"hop-count": 1,
							"link-quality": 85
						    }
					]
				    }
			    }`,
			dataType: &MeshGlobalOperResponse{},
		},
		{
			name: "MeshGlobalStatsResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-mesh-global-oper:mesh-global-stats": {
					"total-mesh-aps": 25,
					"active-mesh-aps": 23,
					"inactive-mesh-aps": 2,
					"mesh-tree-depth": 3,
					"mesh-convergence-time": 120
				    }
			    }`,
			dataType: &MeshGlobalStatsResponse{},
		},
		{
			name: "MeshApTreeDataResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-mesh-global-oper:mesh-ap-tree-data": [
					{
						"ap-mac": "aa:bb:cc:dd:ee:ff",
						"parent-mac": "11:22:33:44:55:66",
						"children-count": 2,
						"hop-count": 1,
						"link-quality": 85
					    }
				]
			    }`,
			dataType: &MeshApTreeDataResponse{},
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
