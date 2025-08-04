// Package geolocation provides geolocation operational data test functionality for the Cisco Wireless Network Controller API.
package geolocation

import (
	"context"
	"testing"
	"time"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
	testutils "github.com/umatare5/cisco-ios-xe-wireless-go/internal/tests"
)

// =============================================================================
// 1. UNIT TESTS (Structure/Type Validation & JSON Serialization/Deserialization)
// =============================================================================

func TestGeolocationOperDataStructures(t *testing.T) {
	testCases := []testutils.JSONTestCase{
		{
			Name: "GeolocationOperResponse",
			JSONData: `{
				"Cisco-IOS-XE-wireless-geolocation-oper:geolocation-oper-data": {
					"ap-geo-loc-stats": {
						"num-ap-gnss": 5,
						"num-ap-man-height": 3,
						"num-ap-derived": 2,
						"last-derivation-timestamp": "2024-01-01T12:00:00.000Z"
					}
				}
			}`,
			Target:     &GeolocationOperResponse{},
			TypeName:   "GeolocationOperResponse",
			ShouldFail: false,
		},
		{
			Name: "GeolocationOperApGeoLocStatsResponse",
			JSONData: `{
				"Cisco-IOS-XE-wireless-geolocation-oper:ap-geo-loc-stats": {
					"num-ap-gnss": 10,
					"num-ap-man-height": 7,
					"num-ap-derived": 3,
					"last-derivation-timestamp": "2024-01-01T13:00:00.000Z"
				}
			}`,
			Target:     &GeolocationOperApGeoLocStatsResponse{},
			TypeName:   "GeolocationOperApGeoLocStatsResponse",
			ShouldFail: false,
		},
	}

	testutils.RunJSONTests(t, testCases)
}

// =============================================================================
// 2. ERROR HANDLING TESTS (Nil Client Validation)
// =============================================================================

func TestGeolocationOperErrorHandling(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	t.Run("GetGeolocationOperWithNilClient", func(t *testing.T) {
		_, err := GetGeolocationOper(nil, ctx)
		if err == nil {
			t.Error("Expected error with nil client, got nil")
		}
		if err.Error() != "client is nil" {
			t.Errorf("Expected 'client is nil' error, got: %v", err)
		}
	})

	t.Run("GetGeolocationOperApGeoLocStatsWithNilClient", func(t *testing.T) {
		_, err := GetGeolocationOperApGeoLocStats(nil, ctx)
		if err == nil {
			t.Error("Expected error with nil client, got nil")
		}
		if err.Error() != "client is nil" {
			t.Errorf("Expected 'client is nil' error, got: %v", err)
		}
	})
}

// =============================================================================
// 3. INTEGRATION TESTS (Actual API Calls to Live Controller)
// =============================================================================

func TestGeolocationOperFunctions(t *testing.T) {
	client := testutils.GetTestClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	t.Run("GetGeolocationOper", func(t *testing.T) {
		data, err := GetGeolocationOper(client, ctx)
		if err != nil {
			t.Fatalf("GetGeolocationOper failed: %v", err)
		}

		if data == nil {
			t.Fatal("GetGeolocationOper returned nil data")
		}

		// Save test data for analysis
		if err := testutils.SaveTestDataToFile("geolocation_oper_data.json", data); err != nil {
			t.Logf("Warning: Could not save test data: %v", err)
		} else {
			t.Logf("Geolocation oper data saved to test_data/geolocation_oper_data.json")
		}

		// Validate endpoint was constructed correctly
		endpoint := GeolocationOperEndpoint
		if endpoint == "" {
			t.Error("GeolocationOperEndpoint should not be empty")
		}
	})

	t.Run("GetGeolocationOperApGeoLocStats", func(t *testing.T) {
		data, err := GetGeolocationOperApGeoLocStats(client, ctx)
		if err != nil {
			t.Fatalf("GetGeolocationOperApGeoLocStats failed: %v", err)
		}

		if data == nil {
			t.Fatal("GetGeolocationOperApGeoLocStats returned nil data")
		}

		// Save test data for analysis
		if err := testutils.SaveTestDataToFile("geolocation_ap_geo_loc_stats_data.json", data); err != nil {
			t.Logf("Warning: Could not save test data: %v", err)
		} else {
			t.Logf("Geolocation AP geo loc stats data saved to test_data/geolocation_ap_geo_loc_stats_data.json")
		}

		// Validate endpoint was constructed correctly
		endpoint := GeolocationApGeoLocStatsEndpoint
		if endpoint == "" {
			t.Error("GeolocationApGeoLocStatsEndpoint should not be empty")
		}
	})
}

// =============================================================================
// 4. CONTEXT HANDLING TESTS
// =============================================================================

func TestGeolocationOperContextHandling(t *testing.T) {
	testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
		_, err := GetGeolocationOper(client, ctx)
		return err
	})

	testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
		_, err := GetGeolocationOperApGeoLocStats(client, ctx)
		return err
	})
}

// =============================================================================
// 5. ENDPOINT VALIDATION TESTS
// =============================================================================

func TestGeolocationOperEndpoints(t *testing.T) {
	// Test base path validation
	t.Run("Validate_GeolocationOperBasePath", func(t *testing.T) {
		expectedBasePath := "/restconf/data/Cisco-IOS-XE-wireless-geolocation-oper:geolocation-oper-data"
		if GeolocationOperBasePath != expectedBasePath {
			t.Errorf("GeolocationOperBasePath mismatch: expected %s, got %s", expectedBasePath, GeolocationOperBasePath)
		}
	})

	// Test endpoint validation
	t.Run("Validate_GeolocationOperEndpoint", func(t *testing.T) {
		if GeolocationOperEndpoint != GeolocationOperBasePath {
			t.Errorf("GeolocationOperEndpoint should equal GeolocationOperBasePath: expected %s, got %s", GeolocationOperBasePath, GeolocationOperEndpoint)
		}
	})

	// Test AP geo loc stats endpoint validation
	t.Run("Validate_GeolocationApGeoLocStatsEndpoint", func(t *testing.T) {
		expectedEndpoint := GeolocationOperBasePath + "/ap-geo-loc-stats"
		if GeolocationApGeoLocStatsEndpoint != expectedEndpoint {
			t.Errorf("GeolocationApGeoLocStatsEndpoint mismatch: expected %s, got %s", expectedEndpoint, GeolocationApGeoLocStatsEndpoint)
		}
	})
}
