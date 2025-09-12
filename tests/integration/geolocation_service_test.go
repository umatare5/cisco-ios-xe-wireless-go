//go:build integration

package integration_test

import (
	"context"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/geolocation"
)

// TestGeolocationServiceIntegration_GetOperationalOperations_Success validates Geolocation service
// operational data retrieval against live WNC controller.
func TestGeolocationServiceIntegration_GetOperationalOperations_Success(t *testing.T) {
	t.Parallel() // Safe for parallel execution as read-only operations
	suite := client.IntegrationTestSuite{
		Config: client.TestSuiteConfig{
			ServiceName: "Geolocation",
			ServiceConstructor: func(client any) any {
				return geolocation.NewService(client.(*core.Client))
			},
			UseTimeout: true,
		},
		BasicMethods: []client.IntegrationTestMethod{
			{
				Name: "GetOperational",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(geolocation.Service).GetOperational(ctx)
				},
			},
			{
				Name: "ListAPGeolocationStats",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(geolocation.Service).ListAPGeolocationStats(ctx)
				},
			},
		},
	}

	client.RunIntegrationTestSuite(t, suite)
}
