// Package afc provides Automated Frequency Coordination operational data test functionality for the Cisco Wireless Network Controller API.
package afc

import (
	"context"
	"testing"
	"time"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
	testutils "github.com/umatare5/cisco-ios-xe-wireless-go/internal/tests"
)

// =============================================================================
// 1. UNIT TESTS (Structure/Type Validation & JSON Serialization/Deserialization)
// =============================================================================

func TestAfcOperDataStructures(t *testing.T) {
	testCases := []testutils.JSONTestCase{
		{
			Name: "AfcOperResponse",
			JSONData: `{
				"Cisco-IOS-XE-wireless-afc-oper:afc-oper-data": {
					"ewlc-afc-ap-resp": [
						{
							"ap-mac": "aa:bb:cc:dd:ee:ff",
							"resp-data": {
								"request-id": "req-12345",
								"ruleset-id": "rule-67890",
								"resp-code": {
									"code": 0,
									"description": "Success",
									"supplemental-info": "AFC response processed successfully"
								},
								"band20": {
									"global-oper-class": 115
								},
								"band40": {
									"global-oper-class": 116
								},
								"band80": {
									"global-oper-class": 125
								},
								"band160": {
									"global-oper-class": 126
								},
								"band80plus": {
									"global-oper-class": 127
								},
								"expire-time": "2024-12-31T23:59:59.000Z",
								"resp-rcvd-timestamp": "2024-01-01T12:00:00.000Z"
							},
							"slot": 0
						}
					]
				}
			}`,
			Target:     &model.AfcOperResponse{},
			TypeName:   "AfcOperResponse",
			ShouldFail: false,
		},
		{
			Name: "AfcOperEwlcAfcApRespResponse",
			JSONData: `{
				"Cisco-IOS-XE-wireless-afc-oper:ewlc-afc-ap-resp": [
					{
						"ap-mac": "aa:bb:cc:dd:ee:ff",
						"resp-data": {
							"request-id": "req-12345",
							"ruleset-id": "rule-67890",
							"resp-code": {
								"code": 0,
								"description": "Success",
								"supplemental-info": "AFC response processed successfully"
							},
							"band20": {
								"global-oper-class": 115
							},
							"band40": {
								"global-oper-class": 116
							},
							"band80": {
								"global-oper-class": 125
							},
							"band160": {
								"global-oper-class": 126
							},
							"band80plus": {
								"global-oper-class": 127
							},
							"expire-time": "2024-12-31T23:59:59.000Z",
							"resp-rcvd-timestamp": "2024-01-01T12:00:00.000Z"
						},
						"slot": 0
					}
				]
			}`,
			Target:     &model.AfcOperEwlcAfcApRespResponse{},
			TypeName:   "AfcOperEwlcAfcApRespResponse",
			ShouldFail: false,
		},
	}

	testutils.RunJSONTests(t, testCases)
}

// =============================================================================
// 2. ERROR HANDLING TESTS (Nil Client Validation)
// =============================================================================

func TestAfcOperErrorHandling(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	t.Run("GetAfcOperWithNilClient", func(t *testing.T) {
		_, err := GetAfcOper(nil, ctx)
		if err == nil {
			t.Error("Expected error with nil client, got nil")
		}
		if err.Error() != "invalid client configuration: client cannot be nil" {
			t.Errorf("Expected 'client is nil' error, got: %v", err)
		}
	})

	t.Run("GetAfcEwlcAfcApRespWithNilClient", func(t *testing.T) {
		_, err := GetAfcEwlcAfcApResp(nil, ctx)
		if err == nil {
			t.Error("Expected error with nil client, got nil")
		}
		if err.Error() != "invalid client configuration: client cannot be nil" {
			t.Errorf("Expected 'client is nil' error, got: %v", err)
		}
	})
}

// =============================================================================
// 3. INTEGRATION TESTS (Actual API Calls to Live Controller)
// =============================================================================

func TestAfcOperFunctions(t *testing.T) {
	client := testutils.GetTestClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	t.Run("GetAfcOper", func(t *testing.T) {
		data, err := GetAfcOper(client, ctx)
		if err != nil {
			t.Fatalf("GetAfcOper failed: %v", err)
		}

		if data == nil {
			t.Fatal("GetAfcOper returned nil data")
		}

		// Save test data for analysis
		if err := testutils.SaveTestDataToFile("afc_oper_data.json", data); err != nil {
			t.Logf("Warning: Could not save test data: %v", err)
		} else {
			t.Logf("AFC oper data saved to test_data/afc_oper_data.json")
		}

		// Validate endpoint was constructed correctly
		endpoint := AfcOperEndpoint
		if endpoint == "" {
			t.Error("AfcOperEndpoint should not be empty")
		}
	})

	t.Run("GetAfcEwlcAfcApResp", func(t *testing.T) {
		data, err := GetAfcEwlcAfcApResp(client, ctx)
		if err != nil {
			t.Fatalf("GetAfcEwlcAfcApResp failed: %v", err)
		}

		if data == nil {
			t.Fatal("GetAfcEwlcAfcApResp returned nil data")
		}

		// Save test data for analysis
		if err := testutils.SaveTestDataToFile("afc_ewlc_afc_ap_resp_data.json", data); err != nil {
			t.Logf("Warning: Could not save test data: %v", err)
		} else {
			t.Logf("AFC EWLC AP response data saved to test_data/afc_ewlc_afc_ap_resp_data.json")
		}

		// Validate endpoint was constructed correctly
		endpoint := AfcOperEwlcAfcApRespEndpoint
		if endpoint == "" {
			t.Error("AfcOperEwlcAfcApRespEndpoint should not be empty")
		}
	})
}

// =============================================================================
// 4. CONTEXT HANDLING TESTS
// =============================================================================

func TestAfcOperContextHandling(t *testing.T) {
	testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
		_, err := GetAfcOper(client, ctx)
		return err
	})

	testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
		_, err := GetAfcEwlcAfcApResp(client, ctx)
		return err
	})
}

// =============================================================================
// 5. ENDPOINT VALIDATION TESTS
// =============================================================================

func TestAfcOperEndpoints(t *testing.T) {
	// Test base path validation
	t.Run("Validate_AfcOperBasePath", func(t *testing.T) {
		expectedBasePath := "Cisco-IOS-XE-wireless-afc-oper:afc-oper-data"
		if AfcOperBasePath != expectedBasePath {
			t.Errorf("AfcOperBasePath mismatch: expected %s, got %s", expectedBasePath, AfcOperBasePath)
		}
	})

	// Test endpoint validation
	t.Run("Validate_AfcOperEndpoint", func(t *testing.T) {
		if AfcOperEndpoint != AfcOperBasePath {
			t.Errorf("AfcOperEndpoint should equal AfcOperBasePath: expected %s, got %s", AfcOperBasePath, AfcOperEndpoint)
		}
	})

	// Test EWLC AFC AP response endpoint validation
	t.Run("Validate_AfcOperEwlcAfcApRespEndpoint", func(t *testing.T) {
		expectedEndpoint := AfcOperBasePath + "/ewlc-afc-ap-resp"
		if AfcOperEwlcAfcApRespEndpoint != expectedEndpoint {
			t.Errorf("AfcOperEwlcAfcApRespEndpoint mismatch: expected %s, got %s", expectedEndpoint, AfcOperEwlcAfcApRespEndpoint)
		}
	})
}

// =============================================================================
// 6. SERVICE TESTS
// =============================================================================

func TestAFCService(t *testing.T) {
	client := testutils.GetTestClient(t)
	if client == nil {
		t.Skip("Skipping service tests: no test client available")
	}

	ctx := context.Background()
	service := NewService(client.CoreClient())

	// Test operational methods
	t.Run("Service_Oper", func(t *testing.T) {
		testutils.TestServiceMethod(t, func() error {
			_, err := service.Oper(ctx)
			return err
		})
	})

	t.Run("Service_APResp", func(t *testing.T) {
		testutils.TestServiceMethod(t, func() error {
			_, err := service.APResp(ctx)
			return err
		})
	})

	t.Run("Service_CloudOper", func(t *testing.T) {
		testutils.TestServiceMethod(t, func() error {
			_, err := service.CloudOper(ctx)
			return err
		})
	})

	t.Run("Service_CloudStats", func(t *testing.T) {
		testutils.TestServiceMethod(t, func() error {
			_, err := service.CloudStats(ctx)
			return err
		})
	})
}
