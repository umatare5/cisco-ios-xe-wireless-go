//go:build integration

package integration_test

import (
	"context"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/afc"
	"github.com/umatare5/cisco-ios-xe-wireless-go/tests/testutil/integration"
)

// Test_AfcGetOperational_IntegrationTests runs comprehensive AFC operational data integration tests
func TestAFCServiceIntegration_GetOperationalOperations_Success(t *testing.T) {
	t.Parallel()

	// Define the test suite configuration
	suite := integration.TestSuite{
		Config: integration.TestSuiteConfig{
			ServiceName: "AFC",
			ServiceConstructor: func(client any) any {
				return afc.NewService(client.(*core.Client))
			},
			UseTimeout: true,
		},
		BasicMethods: []integration.TestMethod{
			{
				Name: "GetOperational",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(afc.Service).GetOperational(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListAPResponses",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(afc.Service).ListAPResponses(ctx)
				},
				LogResult: true,
			},
			{
				Name: "GetCloudInfo",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(afc.Service).GetCloudInfo(ctx)
				},
				LogResult: true,
			},
			{
				Name: "GetCloudStats",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(afc.Service).GetCloudStats(ctx)
				},
				LogResult: true,
			},
		},
	}

	// Run the unified test suite
	integration.RunTestSuite(t, suite)
}
