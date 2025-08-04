// Package afc provides Automated Frequency Coordination cloud operational data test functionality for the Cisco Wireless Network Controller API.
package afc

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

func TestAfcCloudOperDataStructures(t *testing.T) {
	testCases := []testutils.JSONTestCase{
		{
			Name: "AfcCloudOperResponse",
			JSONData: `{
				"Cisco-IOS-XE-wireless-afc-cloud-oper:afc-cloud-oper-data": {
					"afc-cloud-stats": {
						"num-afc-ap": 5,
						"afc-msg-sent": "123",
						"afc-msg-rcvd": "120",
						"afc-msg-err": "3",
						"afc-msg-pending": 2,
						"last-msg-sent": {
							"request-id": "req-12345",
							"ap-mac": "aa:bb:cc:dd:ee:ff",
							"msg-timestamp": "2024-01-01T12:00:00.000Z"
						},
						"last-msg-rcvd": {
							"request-id": "req-12344",
							"ap-mac": "aa:bb:cc:dd:ee:fe",
							"msg-timestamp": "2024-01-01T12:00:01.000Z"
						},
						"min-msg-rtt": "50ms",
						"max-msg-rtt": "500ms",
						"avg-rtt": "150ms",
						"healthcheck": {
							"hc-timestamp": "2024-01-01T12:05:00.000Z",
							"query-in-progress": false,
							"country-not-supported": false,
							"num-hc-down": 0,
							"hc-error-status": {
								"not-otp-upgraded": false
							}
						},
						"num-6ghz-ap": 3
					}
				}
			}`,
			Target:     &AfcCloudOperResponse{},
			TypeName:   "AfcCloudOperResponse",
			ShouldFail: false,
		},
		{
			Name: "AfcCloudOperAfcCloudStatsResponse",
			JSONData: `{
				"Cisco-IOS-XE-wireless-afc-cloud-stats": {
					"num-afc-ap": 5,
					"afc-msg-sent": "123",
					"afc-msg-rcvd": "120",
					"afc-msg-err": "3",
					"afc-msg-pending": 2,
					"last-msg-sent": {
						"request-id": "req-12345",
						"ap-mac": "aa:bb:cc:dd:ee:ff",
						"msg-timestamp": "2024-01-01T12:00:00.000Z"
					},
					"last-msg-rcvd": {
						"request-id": "req-12344",
						"ap-mac": "aa:bb:cc:dd:ee:fe",
						"msg-timestamp": "2024-01-01T12:00:01.000Z"
					},
					"min-msg-rtt": "50ms",
					"max-msg-rtt": "500ms",
					"avg-rtt": "150ms",
					"healthcheck": {
						"hc-timestamp": "2024-01-01T12:05:00.000Z",
						"query-in-progress": false,
						"country-not-supported": false,
						"num-hc-down": 0,
						"hc-error-status": {
							"not-otp-upgraded": false
						}
					},
					"num-6ghz-ap": 3
				}
			}`,
			Target:     &AfcCloudOperAfcCloudStatsResponse{},
			TypeName:   "AfcCloudOperAfcCloudStatsResponse",
			ShouldFail: false,
		},
	}

	testutils.RunJSONTests(t, testCases)
}

// =============================================================================
// 2. ERROR HANDLING TESTS (Nil Client Validation)
// =============================================================================

func TestAfcCloudOperErrorHandling(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	t.Run("GetAfcCloudOperWithNilClient", func(t *testing.T) {
		_, err := GetAfcCloudOper(nil, ctx)
		if err == nil {
			t.Error("Expected error with nil client, got nil")
		}
		if err.Error() != "client is nil" {
			t.Errorf("Expected 'client is nil' error, got: %v", err)
		}
	})

	t.Run("GetAfcCloudStatsWithNilClient", func(t *testing.T) {
		_, err := GetAfcCloudStats(nil, ctx)
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

func TestAfcCloudOperFunctions(t *testing.T) {
	client := testutils.GetTestClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	t.Run("GetAfcCloudOper", func(t *testing.T) {
		data, err := GetAfcCloudOper(client, ctx)
		if err != nil {
			t.Fatalf("GetAfcCloudOper failed: %v", err)
		}

		if data == nil {
			t.Fatal("GetAfcCloudOper returned nil data")
		}

		// Save test data for analysis
		if err := testutils.SaveTestDataToFile("afc_cloud_oper_data.json", data); err != nil {
			t.Logf("Warning: Could not save test data: %v", err)
		} else {
			t.Logf("AFC cloud oper data saved to test_data/afc_cloud_oper_data.json")
		}

		// Validate endpoint was constructed correctly
		endpoint := AfcCloudOperEndpoint
		if endpoint == "" {
			t.Error("AfcCloudOperEndpoint should not be empty")
		}
	})

	t.Run("GetAfcCloudStats", func(t *testing.T) {
		data, err := GetAfcCloudStats(client, ctx)
		if err != nil {
			t.Fatalf("GetAfcCloudStats failed: %v", err)
		}

		if data == nil {
			t.Fatal("GetAfcCloudStats returned nil data")
		}

		// Save test data for analysis
		if err := testutils.SaveTestDataToFile("afc_cloud_stats_data.json", data); err != nil {
			t.Logf("Warning: Could not save test data: %v", err)
		} else {
			t.Logf("AFC cloud stats data saved to test_data/afc_cloud_stats_data.json")
		}

		// Validate endpoint was constructed correctly
		endpoint := AfcCloudStatsEndpoint
		if endpoint == "" {
			t.Error("AfcCloudStatsEndpoint should not be empty")
		}
	})
}

// =============================================================================
// 4. CONTEXT HANDLING TESTS
// =============================================================================

func TestAfcCloudOperContextHandling(t *testing.T) {
	testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
		_, err := GetAfcCloudOper(client, ctx)
		return err
	})

	testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
		_, err := GetAfcCloudStats(client, ctx)
		return err
	})
}

// =============================================================================
// 5. ENDPOINT VALIDATION TESTS
// =============================================================================

func TestAfcCloudOperEndpoints(t *testing.T) {
	// Test base path validation
	t.Run("Validate_AfcCloudOperBasePath", func(t *testing.T) {
		expectedBasePath := "/restconf/data/Cisco-IOS-XE-wireless-afc-cloud-oper:afc-cloud-oper-data"
		if AfcCloudOperBasePath != expectedBasePath {
			t.Errorf("AfcCloudOperBasePath mismatch: expected %s, got %s", expectedBasePath, AfcCloudOperBasePath)
		}
	})

	// Test endpoint validation
	t.Run("Validate_AfcCloudOperEndpoint", func(t *testing.T) {
		if AfcCloudOperEndpoint != AfcCloudOperBasePath {
			t.Errorf("AfcCloudOperEndpoint should equal AfcCloudOperBasePath: expected %s, got %s", AfcCloudOperBasePath, AfcCloudOperEndpoint)
		}
	})

	// Test cloud stats endpoint validation
	t.Run("Validate_AfcCloudStatsEndpoint", func(t *testing.T) {
		expectedEndpoint := AfcCloudOperBasePath + "/afc-cloud-stats"
		if AfcCloudStatsEndpoint != expectedEndpoint {
			t.Errorf("AfcCloudStatsEndpoint mismatch: expected %s, got %s", expectedEndpoint, AfcCloudStatsEndpoint)
		}
	})
}
