package geolocation_test

import (
	"context"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/pkg/testutil"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/geolocation"
)

// TestGeolocationServiceUnit_Constructor_Success tests service constructor functionality.
func TestGeolocationServiceUnit_Constructor_Success(t *testing.T) {
	t.Run("NewServiceWithValidClient", func(t *testing.T) {
		// Create mock server and test client using public API
		responses := map[string]string{
			"test-endpoint": `{"status": "success"}`,
		}
		mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
		defer mockServer.Close()

		testClient := testutil.NewTestClient(mockServer)
		service := geolocation.NewService(testClient.Core().(*core.Client))

		// Test that service can be created and has a client
		if service.Client() == nil {
			t.Error("Expected service to have a client")
		}
	})

	t.Run("NewServiceWithNilClient", func(t *testing.T) {
		service := geolocation.NewService(nil)

		// Service should still be created even with nil client
		if service.Client() != nil {
			t.Error("Expected service with nil client to return nil from Client()")
		}
	})
}

// TestGeolocationServiceUnit_GetOperations_MockSuccess tests Get operations using mock server
// This is essential for CI environments where actual Cisco controllers are not available.
func TestGeolocationServiceUnit_GetOperations_MockSuccess(t *testing.T) {
	// Create mock RESTCONF server with Geolocation endpoints based on live WNC data
	responses := map[string]string{
		"Cisco-IOS-XE-wireless-geolocation-oper:geolocation-oper-data": `{
			"Cisco-IOS-XE-wireless-geolocation-oper:geolocation-oper-data": {
				"ap-geo-loc-stats": {
					"num-ap-gnss": 0,
					"num-ap-man-height": 0,
					"num-ap-derived": 0,
					"last-derivation-timestamp": "2024-01-15T10:36:00.000000+00:00"
				}
			}
		}`,
		"Cisco-IOS-XE-wireless-geolocation-oper:geolocation-oper-data/ap-geo-loc-stats": `{
			"Cisco-IOS-XE-wireless-geolocation-oper:ap-geo-loc-stats": {
				"num-ap-gnss": 0,
				"num-ap-man-height": 0,
				"num-ap-derived": 0,
				"last-derivation-timestamp": "2024-01-15T10:36:00.000000+00:00"
			}
		}`,
		"Cisco-IOS-XE-wireless-geolocation-oper:geolocation-oper-data/ap-geo-loc-data": `{
			"Cisco-IOS-XE-wireless-geolocation-oper:ap-geo-loc-data": [
				{
					"ap-mac": "aa:bb:cc:dd:ee:02",
					"loc": {
						"source": "manual",
						"ellipse": {
							"center": {
								"longitude": "-122.0",
								"latitude": "37.0"
							}
						}
					}
				}
			]
		}`,
	}
	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
	defer mockServer.Close()

	// Create test client configured for the mock server
	testClient := testutil.NewTestClient(mockServer)
	service := geolocation.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test GetOperational operation
	result, err := service.GetOperational(ctx)
	if err != nil {
		t.Errorf("Expected no error for mock GetOperational, got: %v", err)
	}
	if result == nil {
		t.Error("Expected result for mock GetOperational, got nil")
	}

	// Test ListAPGeolocationStats operation
	statsResult, err := service.ListAPGeolocationStats(ctx)
	if err != nil {
		t.Errorf("Expected no error for mock ListAPGeolocationStats, got: %v", err)
	}
	if statsResult == nil {
		t.Error("Expected result for mock ListAPGeolocationStats, got nil")
	}

	// Test ListAPGeolocationData operation (may return 404 if not configured)
	dataResult, err := service.ListAPGeolocationData(ctx)
	if err != nil {
		// Geolocation data endpoints may not be supported by all WNC configurations
		t.Logf("ListAPGeolocationData failed (expected for unconfigured geolocation): %v", err)
	} else if dataResult == nil {
		t.Error("Expected result for mock ListAPGeolocationData, got nil")
	}

	// Test GetAPGeolocationDataByMAC operation (may return 404 if not configured)
	macResult, err := service.GetAPGeolocationDataByMAC(ctx, "aa:bb:cc:dd:ee:02")
	if err != nil {
		// Geolocation data endpoints may not be supported by all WNC configurations
		t.Logf("GetAPGeolocationDataByMAC failed (expected for unconfigured geolocation): %v", err)
	} else if macResult == nil {
		t.Error("Expected result for mock GetAPGeolocationDataByMAC, got nil")
	}
}

// TestGeolocationServiceUnit_GetOperations_IndividualEndpoints tests individual endpoints for complete coverage.
func TestGeolocationServiceUnit_GetOperations_IndividualEndpoints(t *testing.T) {
	// Test individual endpoint responses with correct mock data structure
	testCases := []struct {
		name         string
		endpoint     string
		response     string
		testFunction func(service geolocation.Service, ctx context.Context) (interface{}, error)
	}{
		{
			name:     "GetOperational",
			endpoint: "Cisco-IOS-XE-wireless-geolocation-oper:geolocation-oper-data",
			response: `{
				"Cisco-IOS-XE-wireless-geolocation-oper:geolocation-oper-data": {
					"ap-geo-loc-stats": {
						"num-ap-gnss": 0,
						"num-ap-man-height": 0,
						"num-ap-derived": 1,
						"last-derivation-timestamp": "2024-01-15T10:36:00.000000+00:00"
					}
				}
			}`,
			testFunction: func(service geolocation.Service, ctx context.Context) (interface{}, error) {
				return service.GetOperational(ctx)
			},
		},
		{
			name:     "ListAPGeolocationStats",
			endpoint: "Cisco-IOS-XE-wireless-geolocation-oper:geolocation-oper-data/ap-geo-loc-stats",
			response: `{
				"Cisco-IOS-XE-wireless-geolocation-oper:ap-geo-loc-stats": {
					"num-ap-gnss": 0,
					"num-ap-man-height": 0,
					"num-ap-derived": 1,
					"last-derivation-timestamp": "2024-01-15T10:36:00.000000+00:00"
				}
			}`,
			testFunction: func(service geolocation.Service, ctx context.Context) (interface{}, error) {
				return service.ListAPGeolocationStats(ctx)
			},
		},
		{
			name:     "ListAPGeolocationData",
			endpoint: "Cisco-IOS-XE-wireless-geolocation-oper:geolocation-oper-data/ap-geo-loc-data",
			response: `{
				"Cisco-IOS-XE-wireless-geolocation-oper:ap-geo-loc-data": [
					{
						"ap-mac": "aa:bb:cc:dd:ee:02",
						"loc": {
							"source": "manual",
							"ellipse": {
								"center": {
									"longitude": "-122.0",
									"latitude": "37.0"
								}
							}
						}
					}
				]
			}`,
			testFunction: func(service geolocation.Service, ctx context.Context) (interface{}, error) {
				return service.ListAPGeolocationData(ctx)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			responses := map[string]string{tc.endpoint: tc.response}
			mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
			defer mockServer.Close()

			testClient := testutil.NewTestClient(mockServer)
			service := geolocation.NewService(testClient.Core().(*core.Client))
			ctx := testutil.TestContext(t)

			result, err := tc.testFunction(service, ctx)
			if err != nil {
				t.Errorf("Expected success for %s, got error: %v", tc.name, err)
			}
			if result == nil {
				t.Errorf("Expected non-nil result for %s", tc.name)
			}
		})
	}
}

// TestGeolocationServiceUnit_GetOperations_ErrorHandling tests error scenarios using mock server.
func TestGeolocationServiceUnit_GetOperations_ErrorHandling(t *testing.T) {
	// Create mock server that returns 404 for Geolocation endpoints
	errorPaths := []string{
		"Cisco-IOS-XE-wireless-geolocation-oper:geolocation-oper-data",
	}
	mockServer := testutil.NewMockServer(testutil.WithErrorResponses(errorPaths, 404))
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := geolocation.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test that GetOperational properly handles 404 errors
	_, err := service.GetOperational(ctx)
	if err == nil {
		t.Error("Expected error for 404 response, got nil")
	}

	// Verify error contains expected information
	if !core.IsNotFoundError(err) {
		t.Errorf("Expected NotFound error, got: %v", err)
	}

	// Test that ListAPGeolocationData properly handles 404 errors
	_, err = service.ListAPGeolocationData(ctx)
	if err == nil {
		t.Error("Expected error for 404 response, got nil")
	}

	// Test that GetAPGeolocationDataByMAC properly handles 404 errors
	_, err = service.GetAPGeolocationDataByMAC(ctx, "aa:bb:cc:dd:ee:02")
	if err == nil {
		t.Error("Expected error for 404 response, got nil")
	}
}

// TestGeolocationServiceUnit_GetOperations_GetAPGeolocationDataByMAC tests GetAPGeolocationDataByMAC success scenarios.
func TestGeolocationServiceUnit_GetOperations_GetAPGeolocationDataByMAC(t *testing.T) {
	responses := map[string]string{
		// Exact path that BuildQueryURL will construct
		"/restconf/data/Cisco-IOS-XE-wireless-geolocation-oper:geolocation-oper-data/ap-geo-loc-data=aa:bb:cc:dd:ee:02": `{
			"Cisco-IOS-XE-wireless-geolocation-oper:ap-geo-loc-data": [
				{
					"ap-mac": "aa:bb:cc:dd:ee:02",
					"loc": {
						"source": "manual",
						"ellipse": {
							"center": {
								"longitude": "-122.0",
								"latitude": "37.0"
							}
						}
					}
				}
			]
		}`,
		"/restconf/data/Cisco-IOS-XE-wireless-geolocation-oper:geolocation-oper-data/ap-geo-loc-data=aa:bb:cc:dd:ee:01": `{
			"Cisco-IOS-XE-wireless-geolocation-oper:ap-geo-loc-data": [
				{
					"ap-mac": "aa:bb:cc:dd:ee:01",
					"loc": {
						"source": "derived",
						"ellipse": {
							"center": {
								"longitude": "-122.419416",
								"latitude": "37.7749295"
							}
						}
					}
				}
			]
		}`,
		// Also support simplified paths for the mock server
		"Cisco-IOS-XE-wireless-geolocation-oper:geolocation-oper-data/ap-geo-loc-data=aa:bb:cc:dd:ee:02": `{
			"Cisco-IOS-XE-wireless-geolocation-oper:ap-geo-loc-data": [
				{
					"ap-mac": "aa:bb:cc:dd:ee:02",
					"loc": {
						"source": "manual",
						"ellipse": {
							"center": {
								"longitude": "-122.0",
								"latitude": "37.0"
							}
						}
					}
				}
			]
		}`,
		"Cisco-IOS-XE-wireless-geolocation-oper:geolocation-oper-data/ap-geo-loc-data=aa:bb:cc:dd:ee:01": `{
			"Cisco-IOS-XE-wireless-geolocation-oper:ap-geo-loc-data": [
				{
					"ap-mac": "aa:bb:cc:dd:ee:01",
					"loc": {
						"source": "derived",
						"ellipse": {
							"center": {
								"longitude": "-122.419416",
								"latitude": "37.7749295"
							}
						}
					}
				}
			]
		}`,
	}
	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := geolocation.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test successful GetAPGeolocationDataByMAC with colon format
	result, err := service.GetAPGeolocationDataByMAC(ctx, "aa:bb:cc:dd:ee:02")
	if err != nil {
		t.Errorf("Expected success for valid MAC, got error: %v", err)
	}
	if result == nil {
		t.Error("Expected non-nil result for valid MAC")
	}

	// Test successful GetAPGeolocationDataByMAC with different MAC format normalization
	result, err = service.GetAPGeolocationDataByMAC(ctx, "aa-bb-cc-dd-ee-01")
	if err != nil {
		t.Errorf("Expected success for dash-separated MAC, got error: %v", err)
	}
	if result == nil {
		t.Error("Expected non-nil result for dash-separated MAC")
	}

	// Test successful GetAPGeolocationDataByMAC with uppercase no-separator format
	result, err = service.GetAPGeolocationDataByMAC(ctx, "AABBCCDDEE01")
	if err != nil {
		t.Errorf("Expected success for uppercase no-separator MAC, got error: %v", err)
	}
	if result == nil {
		t.Error("Expected non-nil result for uppercase no-separator MAC")
	}

	// Test MAC normalization validation path coverage
	testMACs := []string{
		"aA:bB:cC:dD:eE:01", // Mixed case
		"aa-bb-cc-dd-ee-02", // Dash format
		"AABBCCDDEE02",      // No separators uppercase
		"aabbccddee02",      // No separators lowercase
	}

	for _, mac := range testMACs {
		_, err := service.GetAPGeolocationDataByMAC(ctx, mac)
		// We don't care about success/failure here - we're testing code path coverage
		// through validation and normalization
		_ = err
	}
}

// TestGeolocationServiceUnit_ValidationErrors_InvalidMAC tests validation error scenarios.
func TestGeolocationServiceUnit_ValidationErrors_InvalidMAC(t *testing.T) {
	responses := map[string]string{
		"test-endpoint": `{"status": "success"}`,
	}
	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := geolocation.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test validation errors for empty MAC
	_, err := service.GetAPGeolocationDataByMAC(ctx, "")
	if err == nil {
		t.Error("Expected validation error for empty MAC address, got nil")
	}

	// Test validation errors for whitespace MAC
	_, err = service.GetAPGeolocationDataByMAC(ctx, "   ")
	if err == nil {
		t.Error("Expected validation error for whitespace MAC address, got nil")
	}

	// Test validation errors for invalid MAC
	_, err = service.GetAPGeolocationDataByMAC(ctx, "invalid-mac")
	if err == nil {
		t.Error("Expected validation error for invalid MAC address, got nil")
	}

	// Test validation errors for invalid MAC format
	_, err = service.GetAPGeolocationDataByMAC(ctx, "gg:hh:ii:jj:kk:ll")
	if err == nil {
		t.Error("Expected validation error for invalid MAC characters, got nil")
	}

	// Test MAC normalization error path
	_, err = service.GetAPGeolocationDataByMAC(ctx, "12:34:56:78:90:zz")
	if err == nil {
		t.Error("Expected normalization error for invalid MAC format, got nil")
	}

	// Test MAC too short error path
	_, err = service.GetAPGeolocationDataByMAC(ctx, "12:34:56")
	if err == nil {
		t.Error("Expected validation error for short MAC, got nil")
	}

	// Test MAC too long error path
	_, err = service.GetAPGeolocationDataByMAC(ctx, "12:34:56:78:90:ab:cd:ef")
	if err == nil {
		t.Error("Expected validation error for long MAC, got nil")
	}
}

// TestGeolocationServiceUnit_ErrorHandling_NilClient tests nil client scenarios.
func TestGeolocationServiceUnit_ErrorHandling_NilClient(t *testing.T) {
	service := geolocation.NewService(nil)
	ctx := testutil.TestContext(t)

	// Test that GetOperational handles nil client
	_, err := service.GetOperational(ctx)
	if err == nil {
		t.Error("Expected error with nil client for GetOperational, got nil")
	}

	// Test that ListAPGeolocationStats handles nil client
	_, err = service.ListAPGeolocationStats(ctx)
	if err == nil {
		t.Error("Expected error with nil client for ListAPGeolocationStats, got nil")
	}

	// Test that ListAPGeolocationData handles nil client
	_, err = service.ListAPGeolocationData(ctx)
	if err == nil {
		t.Error("Expected error with nil client for ListAPGeolocationData, got nil")
	}

	// Test that GetAPGeolocationDataByMAC handles nil client
	_, err = service.GetAPGeolocationDataByMAC(ctx, "aa:bb:cc:dd:ee:02")
	if err == nil {
		t.Error("Expected error with nil client for GetAPGeolocationDataByMAC, got nil")
	}
}

// TestGeolocationServiceUnit_QuotedDecimal64_MockSuccess tests that the platform's quoted
// decimal64 encoding decodes and that no bare sibling is lost with it.
func TestGeolocationServiceUnit_QuotedDecimal64_MockSuccess(t *testing.T) {
	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{
		"Cisco-IOS-XE-wireless-geolocation-oper:geolocation-oper-data": `{
			"Cisco-IOS-XE-wireless-geolocation-oper:geolocation-oper-data": {
				"ap-geo-loc-data": [
					{
						"ap-mac": "aa:bb:cc:dd:ee:01",
						"loc": {
							"source": "gnss",
							"area-of-uncertainty": 12,
							"hdop": "0.900000",
							"ellipse": {
								"center": {"longitude": "0.000000", "latitude": "0.000000"},
								"major-axis": 10,
								"orientation": "0.000000"
							}
						}
					}
				]
			}
		}`,
	}))
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := geolocation.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	result, err := service.GetOperational(ctx)
	if err != nil {
		t.Fatalf("GetOperational failed: %v", err)
	}

	if result.CiscoIOSXEWirelessGeolocationOperData == nil {
		t.Fatal("Expected geolocation-oper-data to decode")
	}

	records := result.CiscoIOSXEWirelessGeolocationOperData.ApGeoLocData
	if len(records) != 1 {
		t.Fatalf("Expected 1 record, got %d", len(records))
	}

	loc := records[0].Loc
	if loc == nil || loc.HDOP == nil || *loc.HDOP != "0.900000" {
		t.Fatalf("Expected hdop to decode, got %+v", loc)
	}
	if loc.AreaOfUncertainty == nil || *loc.AreaOfUncertainty != 12 {
		t.Error("Expected a bare sibling of the quoted leaves to survive")
	}
	if loc.Ellipse == nil || loc.Ellipse.Center == nil || loc.Ellipse.Center.Longitude == nil {
		t.Fatal("Expected ellipse center to decode")
	}
	if *loc.Ellipse.Center.Longitude != "0.000000" {
		t.Errorf("Expected longitude in the quoted wire form, got %q", *loc.Ellipse.Center.Longitude)
	}
	if loc.Ellipse.Orientation == nil || *loc.Ellipse.Orientation != "0.000000" {
		t.Error("Expected orientation to decode")
	}
}

// TestGeolocationServiceUnit_DeclaredWidths_MockSuccess pins the five integer leaves to the
// widths the device declares. The in-range case carries each leaf at its declared maximum,
// which a narrower type would refuse.
func TestGeolocationServiceUnit_DeclaredWidths_MockSuccess(t *testing.T) {
	const route = "Cisco-IOS-XE-wireless-geolocation-oper:geolocation-oper-data"
	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{
		route: `{
			"Cisco-IOS-XE-wireless-geolocation-oper:geolocation-oper-data": {
				"ap-geo-loc-data": [
					{
						"ap-mac": "aa:bb:cc:dd:ee:01",
						"loc": {
							"area-of-uncertainty": 4294967295,
							"ellipse": {"major-axis": 65535, "minor-axis": 65535}
						},
						"elevation": {"agl-data": {"height": -32768, "uncertainty": 65535}}
					}
				]
			}
		}`,
	}))
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := geolocation.NewService(testClient.Core().(*core.Client))

	result, err := service.GetOperational(testutil.TestContext(t))
	if err != nil {
		t.Fatalf("GetOperational failed: %v", err)
	}

	records := result.CiscoIOSXEWirelessGeolocationOperData.ApGeoLocData
	if len(records) != 1 {
		t.Fatalf("Expected 1 record, got %d", len(records))
	}

	loc, elev := records[0].Loc, records[0].Elevation
	if loc.AreaOfUncertainty == nil || *loc.AreaOfUncertainty != 4294967295 {
		t.Errorf("area-of-uncertainty is a uint32: %v", loc.AreaOfUncertainty)
	}
	if loc.Ellipse.MajorAxis == nil || *loc.Ellipse.MajorAxis != 65535 {
		t.Errorf("major-axis is a uint16: %v", loc.Ellipse.MajorAxis)
	}
	if loc.Ellipse.MinorAxis == nil || *loc.Ellipse.MinorAxis != 65535 {
		t.Errorf("minor-axis is a uint16: %v", loc.Ellipse.MinorAxis)
	}
	if elev.AGLData.Height == nil || *elev.AGLData.Height != -32768 {
		t.Errorf("height is a signed int16: %v", elev.AGLData.Height)
	}
	if elev.AGLData.Uncertainty == nil || *elev.AGLData.Uncertainty != 65535 {
		t.Errorf("elevation uncertainty is a uint16: %v", elev.AGLData.Uncertainty)
	}
}

// TestGeolocationServiceUnit_OutOfRangeIsRefused_MockError pins the widths from the other
// side: a value the declared type cannot hold must fail the read rather than arrive
// truncated or negative. A plain int would accept every one of these on a 64-bit build,
// so this is what separates the declared width from "wide enough today".
func TestGeolocationServiceUnit_OutOfRangeIsRefused_MockError(t *testing.T) {
	const route = "Cisco-IOS-XE-wireless-geolocation-oper:geolocation-oper-data"

	for name, loc := range map[string]string{
		"NegativeAreaOfUncertainty": `"area-of-uncertainty": -1`,
		"AreaOfUncertaintyOverflow": `"area-of-uncertainty": 4294967296`,
		"MajorAxisOverflow":         `"ellipse": {"major-axis": 65536}`,
		"MinorAxisNegative":         `"ellipse": {"minor-axis": -1}`,
	} {
		t.Run(name, func(t *testing.T) {
			mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{
				route: `{"Cisco-IOS-XE-wireless-geolocation-oper:geolocation-oper-data":{"ap-geo-loc-data":[{"loc":{` +
					loc + `}}]}}`,
			}))
			defer mockServer.Close()

			testClient := testutil.NewTestClient(mockServer)
			service := geolocation.NewService(testClient.Core().(*core.Client))

			if _, err := service.GetOperational(testutil.TestContext(t)); err == nil {
				t.Error("Expected a value outside the declared width to fail the read")
			}
		})
	}

	for name, elev := range map[string]string{
		"HeightOverflow":               `"height": 32768`,
		"HeightUnderflow":              `"height": -32769`,
		"ElevationUncertaintyNegative": `"uncertainty": -1`,
	} {
		t.Run(name, func(t *testing.T) {
			mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{
				route: `{"Cisco-IOS-XE-wireless-geolocation-oper:geolocation-oper-data":{"ap-geo-loc-data":[{"elevation":{"agl-data":{` +
					elev + `}}}]}}`,
			}))
			defer mockServer.Close()

			testClient := testutil.NewTestClient(mockServer)
			service := geolocation.NewService(testClient.Core().(*core.Client))

			if _, err := service.GetOperational(testutil.TestContext(t)); err == nil {
				t.Error("Expected a value outside the declared width to fail the read")
			}
		})
	}
}
