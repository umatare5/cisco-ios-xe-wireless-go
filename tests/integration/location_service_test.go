//go:build integration

package integration_test

import (
	"context"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/location"
	"github.com/umatare5/cisco-ios-xe-wireless-go/tests/testutil/integration"
)

// TestLocationServiceIntegration_GetOperationalOperations_Success validates Location service
// operational data retrieval against live WNC controller.
func TestLocationServiceIntegration_GetOperationalOperations_Success(t *testing.T) {
	t.Parallel()

	// Define the test suite configuration
	suite := integration.TestSuite{
		Config: integration.TestSuiteConfig{
			ServiceName: "Location",
			ServiceConstructor: func(client any) any {
				return location.NewService(client.(*core.Client))
			},
			UseTimeout: true,
		},
		BasicMethods: []integration.TestMethod{
			{
				Name: "GetOperational",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(location.Service).GetOperational(ctx)
				},
				LogResult:      true,
				ExpectNotFound: true, // Not Verified on IOS-XE 17.12.5
			},
			{
				Name: "LocationRSSIMeasurements",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(location.Service).LocationRSSIMeasurements(ctx)
				},
				LogResult:      true,
				ExpectNotFound: true, // Should return empty response for HTTP 204
			},
			{
				Name: "GetConfig",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(location.Service).GetConfig(ctx)
				},
				LogResult:      true,
				ExpectNotFound: false, // GetConfig returns data
			},
			{
				Name: "ListOperatorLocations",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(location.Service).ListOperatorLocations(ctx)
				},
				LogResult:      true,
				ExpectNotFound: false, // Should return empty response for HTTP 204
			},
			{
				Name: "ListNMSPConfig",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(location.Service).ListNMSPConfig(ctx)
				},
				LogResult:      true,
				ExpectNotFound: false, // Should return empty response for HTTP 204
			},
			{
				Name: "GetLocation",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(location.Service).GetLocation(ctx)
				},
				LogResult:      true,
				ExpectNotFound: true,
			},
		},
		FilterMethods:   []integration.TestMethod{},
		ValidationTests: []integration.ValidationTestMethod{},
	}

	integration.RunTestSuite(t, suite)
}
