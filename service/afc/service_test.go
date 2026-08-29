package afc_test

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/pkg/testutil"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/afc"
)

// TestAfcServiceUnit_Constructor_Success tests service constructor functionality.
func TestAfcServiceUnit_Constructor_Success(t *testing.T) {
	t.Run("NewServiceWithValidClient", func(t *testing.T) {
		// Create mock server and test client using public API
		responses := map[string]string{
			"test-endpoint": `{"status": "success"}`,
		}
		mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
		defer mockServer.Close()

		testClient := testutil.NewTestClient(mockServer)
		service := afc.NewService(testClient.Core().(*core.Client))

		// Test that service can be created and has a client
		if service.Client() == nil {
			t.Error("Expected service to have a client")
		}
	})

	t.Run("NewServiceWithNilClient", func(t *testing.T) {
		service := afc.NewService(nil)

		// Service should still be created even with nil client
		if service.Client() != nil {
			t.Error("Expected service with nil client to return nil from Client()")
		}
	})
}

// TestAfcServiceUnit_GetOperations_MockSuccess tests Get operations using mock server
// This is essential for CI environments where actual Cisco controllers are not available.
func TestAfcServiceUnit_GetOperations_MockSuccess(t *testing.T) {
	// Create mock RESTCONF server with AFC endpoints
	responses := map[string]string{
		"Cisco-IOS-XE-wireless-afc-oper:afc-oper-data": `{
			"Cisco-IOS-XE-wireless-afc-oper:afc-oper-data": {
				"ewlc-afc-info": {
					"afc-enable": true,
					"afc-server-url": "https://afc.example.com"
				}
			}
		}`,
		"Cisco-IOS-XE-wireless-afc-oper:afc-oper-data/ewlc-afc-ap-resp": `{
			"Cisco-IOS-XE-wireless-afc-oper:ewlc-afc-ap-resp": [{
				"ap-mac": "aa:bb:cc:dd:ee:01",
				"response-status": "success"
			}]
		}`,
		"Cisco-IOS-XE-wireless-afc-oper:afc-oper-data/ewlc-afc-ap-req": `{
			"Cisco-IOS-XE-wireless-afc-oper:ewlc-afc-ap-req": [{
				"ap-mac": "aa:bb:cc:dd:ee:01",
				"request-status": "pending"
			}]
		}`,
		"Cisco-IOS-XE-wireless-afc-cloud-oper:afc-cloud-oper-data": `{
			"Cisco-IOS-XE-wireless-afc-cloud-oper:afc-cloud-oper-data": {
				"afc-cloud-enable": true
			}
		}`,
		"Cisco-IOS-XE-wireless-afc-cloud-oper:afc-cloud-oper-data/afc-cloud-stats": `{
			"Cisco-IOS-XE-wireless-afc-cloud-oper:afc-cloud-stats": {
				"requests-sent": 100,
				"responses-received": 95
			}
		}`,
	}
	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
	defer mockServer.Close()

	// Create test client configured for the mock server
	testClient := testutil.NewTestClient(mockServer)
	service := afc.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test all AFC Get operations
	t.Run("GetOperational", func(t *testing.T) {
		result, err := service.GetOperational(ctx)
		if err != nil {
			t.Errorf("Expected no error for GetOperational, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for GetOperational, got nil")
		}
	})

	t.Run("ListAPResponses", func(t *testing.T) {
		result, err := service.ListAPResponses(ctx)
		if err != nil {
			t.Errorf("Expected no error for ListAPResponses, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for ListAPResponses, got nil")
		}
	})

	t.Run("ListAPRequests", func(t *testing.T) {
		result, err := service.ListAPRequests(ctx)
		if err != nil {
			t.Errorf("Expected no error for ListAPRequests, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for ListAPRequests, got nil")
		}
	})

	t.Run("GetCloudInfo", func(t *testing.T) {
		result, err := service.GetCloudInfo(ctx)
		if err != nil {
			t.Errorf("Expected no error for GetCloudInfo, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for GetCloudInfo, got nil")
		}
	})

	t.Run("GetCloudStats", func(t *testing.T) {
		result, err := service.GetCloudStats(ctx)
		if err != nil {
			t.Errorf("Expected no error for GetCloudStats, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for GetCloudStats, got nil")
		}
	})
}

// TestAfcServiceUnit_GetOperations_ErrorHandling tests error scenarios using mock server.
func TestAfcServiceUnit_GetOperations_ErrorHandling(t *testing.T) {
	// Create mock server that returns 404 for AFC endpoints
	errorPaths := []string{
		"Cisco-IOS-XE-wireless-afc-oper:afc-oper-data",
		"Cisco-IOS-XE-wireless-afc-oper:afc-oper-data/ewlc-afc-ap-resp",
		"Cisco-IOS-XE-wireless-afc-oper:afc-oper-data/ewlc-afc-ap-req",
		"Cisco-IOS-XE-wireless-afc-cloud-oper:afc-cloud-oper-data",
		"Cisco-IOS-XE-wireless-afc-cloud-oper:afc-cloud-oper-data/afc-cloud-stats",
	}
	mockServer := testutil.NewMockServer(testutil.WithErrorResponses(errorPaths, 404))
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := afc.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	t.Run("GetOperational", func(t *testing.T) {
		_, err := service.GetOperational(ctx)
		if err == nil {
			t.Error("Expected error for GetOperational, got nil")
		}
	})

	t.Run("ListAPResponses", func(t *testing.T) {
		_, err := service.ListAPResponses(ctx)
		if err == nil {
			t.Error("Expected error for ListAPResponses, got nil")
		}
	})

	t.Run("ListAPRequests", func(t *testing.T) {
		_, err := service.ListAPRequests(ctx)
		if err == nil {
			t.Error("Expected error for ListAPRequests, got nil")
		}
	})

	t.Run("GetCloudInfo", func(t *testing.T) {
		_, err := service.GetCloudInfo(ctx)
		if err == nil {
			t.Error("Expected error for GetCloudInfo, got nil")
		}
	})

	t.Run("GetCloudStats", func(t *testing.T) {
		_, err := service.GetCloudStats(ctx)
		if err == nil {
			t.Error("Expected error for GetCloudStats, got nil")
		}
	})
}

// TestAfcServiceUnit_QuotedDecimal64_MockSuccess is the regression barrier for the five afc leaves
// this release retyped from *float64 to *string. RFC 7951 6.1 requires a JSON string for a
// decimal64 leaf, so the float form could not decode one; nothing in this package asserted that
// before, and reverting any of the five leaves would have passed the whole suite.
//
// The quoted uint64 in the same body is here for the same reason.
func TestAfcServiceUnit_QuotedDecimal64_MockSuccess(t *testing.T) {
	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{
		"Cisco-IOS-XE-wireless-afc-oper:afc-oper-data/ewlc-afc-ap-req": `{
			"Cisco-IOS-XE-wireless-afc-oper:ewlc-afc-ap-req": [
				{
					"ap-mac": "aa:bb:cc:dd:ee:01",
					"req-id-sent": "18446744073709551615",
					"req-data": {
						"min-desired-power": "-12.5",
						"location": {
							"ellipse": {
								"center": {"longitude": "139.7671248", "latitude": "35.6812405"},
								"major-axis": 30,
								"minor-axis": 10,
								"orientation": "45.0"
							}
						}
					}
				}
			]
		}`,
	}))
	defer mockServer.Close()

	service := afc.NewService(testutil.NewTestClient(mockServer).Core().(*core.Client))

	result, err := service.ListAPRequests(testutil.TestContext(t))
	if err != nil {
		t.Fatalf("ListAPRequests failed: %v", err)
	}
	if len(result.EwlcAFCApReq) != 1 {
		t.Fatalf("Expected 1 request, got %+v", result.EwlcAFCApReq)
	}

	req := result.EwlcAFCApReq[0]
	if req.ReqIDSent == nil || *req.ReqIDSent != "18446744073709551615" {
		t.Errorf("Expected the quoted uint64 to decode as its string, got %v", req.ReqIDSent)
	}
	if req.ReqData == nil || req.ReqData.MinDesiredPower == nil ||
		*req.ReqData.MinDesiredPower != "-12.5" {
		t.Fatalf("Expected min-desired-power to decode as a string, got %+v", req.ReqData)
	}

	ellipse := req.ReqData.Location.Ellipse
	if ellipse == nil || ellipse.Center == nil {
		t.Fatal("Expected the ellipse and its center to decode")
	}
	if ellipse.Center.Longitude == nil || *ellipse.Center.Longitude != "139.7671248" {
		t.Errorf("Expected longitude to decode as a string, got %v", ellipse.Center.Longitude)
	}
	if ellipse.Center.Latitude == nil || *ellipse.Center.Latitude != "35.6812405" {
		t.Errorf("Expected latitude to decode as a string, got %v", ellipse.Center.Latitude)
	}
	if ellipse.Orientation == nil || *ellipse.Orientation != "45.0" {
		t.Errorf("Expected orientation to decode as a string, got %v", ellipse.Orientation)
	}
}

// TestAfcServiceUnit_QuotedMaxEIRP_MockSuccess covers the fifth retyped decimal64 leaf, which sits
// on the response side rather than the request side.
func TestAfcServiceUnit_QuotedMaxEIRP_MockSuccess(t *testing.T) {
	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{
		"Cisco-IOS-XE-wireless-afc-oper:afc-oper-data/ewlc-afc-ap-resp": `{
			"Cisco-IOS-XE-wireless-afc-oper:ewlc-afc-ap-resp": [
				{
					"ap-mac": "aa:bb:cc:dd:ee:01",
					"slot": 2,
					"resp-data": {
						"request-id": "1",
						"band20": {
							"global-oper-class": 131,
							"channels": [
								{"avail-channel-cfi": 1, "max-eirp": "30.5"}
							]
						}
					}
				}
			]
		}`,
	}))
	defer mockServer.Close()

	service := afc.NewService(testutil.NewTestClient(mockServer).Core().(*core.Client))

	result, err := service.ListAPResponses(testutil.TestContext(t))
	if err != nil {
		t.Fatalf("ListAPResponses failed: %v", err)
	}
	if len(result.EwlcAFCApResp) != 1 {
		t.Fatalf("Expected 1 response, got %+v", result.EwlcAFCApResp)
	}

	channels := result.EwlcAFCApResp[0].RespData.Band20.Channels
	if len(channels) != 1 {
		t.Fatalf("Expected 1 channel, got %+v", channels)
	}
	if channels[0].MaxEIRP == nil || *channels[0].MaxEIRP != "30.5" {
		t.Errorf("Expected max-eirp to decode as a string, got %v", channels[0].MaxEIRP)
	}
}
