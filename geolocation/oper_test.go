// Package geolocation provides geolocation operational data test functionality for the Cisco Wireless Network Controller API.
package geolocation

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

// GeolocationOperTestDataCollector holds test data for geolocation operation functions
type GeolocationOperTestDataCollector struct {
	Data map[string]interface{} `json:"geolocation_oper_test_data"`
}

// newGeolocationOperTestDataCollector creates a new test data collector
func newGeolocationOperTestDataCollector() *GeolocationOperTestDataCollector {
	return &GeolocationOperTestDataCollector{
		Data: make(map[string]interface{}),
	}
}

// =============================================================================
// 2. INTEGRATION TESTS (API Endpoint Testing with Live Data Validation)
// =============================================================================

func TestGeolocationOperGetGeolocationOper(t *testing.T) {
	client := getTestClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), testutil.DefaultTestTimeout)
	defer cancel()

	result, err := GetGeolocationOper(client, ctx)
	if err != nil {
		t.Fatalf("GetGeolocationOper failed: %v", err)
	}

	if result == nil {
		t.Fatal("GetGeolocationOper returned nil result")
	}

	// Save result to JSON file
	filename := fmt.Sprintf("geolocation_oper_data_%d.json", time.Now().Unix())
	if err := testutil.SaveTestDataToFile(filename, result); err != nil {
		t.Logf("Warning: Failed to save data to %s: %v", filename, err)
	} else {
		t.Logf("Data saved to %s", filename)
	}

	t.Logf("GetGeolocationOper successful, collected geolocation operational data")
}

func TestGeolocationOperGetGeolocationOperApGeoLocStats(t *testing.T) {
	client := getTestClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), testutil.DefaultTestTimeout)
	defer cancel()

	result, err := GetGeolocationOperApGeoLocStats(client, ctx)
	if err != nil {
		t.Fatalf("GetGeolocationOperApGeoLocStats failed: %v", err)
	}

	if result == nil {
		t.Fatal("GetGeolocationOperApGeoLocStats returned nil result")
	}

	// Save result to JSON file
	filename := fmt.Sprintf("geolocation_oper_ap_geo_loc_stats_data_%d.json", time.Now().Unix())
	if err := testutil.SaveTestDataToFile(filename, result); err != nil {
		t.Logf("Warning: Failed to save data to %s: %v", filename, err)
	} else {
		t.Logf("Data saved to %s", filename)
	}

	t.Logf("GetGeolocationOperApGeoLocStats successful")
}

func TestGeolocationOperCollectAllData(t *testing.T) {
	client := getTestClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), testutil.DefaultTestTimeout)
	defer cancel()

	allData := make(map[string]interface{})

	// Collect data from all geolocation operational endpoints
	tests := []struct {
		name string
		fn   func() (interface{}, error)
	}{
		{"GetGeolocationOper", func() (interface{}, error) { return GetGeolocationOper(client, ctx) }},
		{"GetGeolocationOperApGeoLocStats", func() (interface{}, error) { return GetGeolocationOperApGeoLocStats(client, ctx) }},
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
	filename := fmt.Sprintf("geolocation_oper_comprehensive_data_%d.json", time.Now().Unix())
	if err := testutil.SaveTestDataToFile(filename, allData); err != nil {
		t.Logf("Warning: Failed to save comprehensive data to %s: %v", filename, err)
	} else {
		t.Logf("Comprehensive geolocation operational data saved to %s", filename)
	}
}

// TestGeolocationOperDataStructures tests the basic structure of geolocation operational data types
func TestGeolocationOperDataStructures(t *testing.T) {
	tests := []struct {
		name     string
		jsonData string
		dataType interface{}
	}{
		{
			name: "GeolocationOperResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-geolocation-oper:geolocation-oper-data": {
					"ap-geo-loc-stats": {
						"num-ap-gnss": 10,
						"num-ap-man-height": 5,
						"num-ap-derived": 3,
						"num-ap-manual": 15,
						"num-ap-auto": 8,
						"num-ap-invalid": 2,
						"num-ap-total": 33
					    }
				    }
			    }`,
			dataType: &GeolocationOperResponse{},
		},
		{
			name: "GeolocationOperApGeoLocStatsResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-geolocation-oper:ap-geo-loc-stats": {
					"num-ap-gnss": 10,
					"num-ap-man-height": 5,
					"num-ap-derived": 3,
					"num-ap-manual": 15,
					"num-ap-auto": 8,
					"num-ap-invalid": 2,
					"num-ap-total": 33
				    }
			    }`,
			dataType: &GeolocationOperApGeoLocStatsResponse{},
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
