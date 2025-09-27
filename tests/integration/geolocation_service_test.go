//go:build integration

package integration_test

import (
	"context"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/geolocation"
	"github.com/umatare5/cisco-ios-xe-wireless-go/tests/testutil/integration"
)

// TestGeolocationServiceIntegration_GetOperationalOperations_Success validates Geolocation service
// operational data retrieval against live WNC controller.
func TestGeolocationServiceIntegration_GetOperationalOperations_Success(t *testing.T) {
	t.Parallel()

	// Define the test suite configuration
	suite := integration.TestSuite{
		Config: integration.TestSuiteConfig{
			ServiceName: "Geolocation",
			ServiceConstructor: func(client any) any {
				return geolocation.NewService(client.(*core.Client))
			},
			UseTimeout: true,
		},
		BasicMethods: []integration.TestMethod{
			{
				Name: "GetOperational",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(geolocation.Service).GetOperational(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListAPGeolocationStats",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(geolocation.Service).ListAPGeolocationStats(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListAPGeolocationData",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(geolocation.Service).ListAPGeolocationData(ctx)
				},
				LogResult: true,
			},
		},
	}

	integration.RunTestSuite(t, suite)
}
