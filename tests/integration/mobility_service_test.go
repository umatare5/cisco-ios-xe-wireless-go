//go:build integration

package integration_test

import (
	"context"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/mobility"
	"github.com/umatare5/cisco-ios-xe-wireless-go/tests/testutil/integration"
)

// TestMobilityServiceIntegration_GetOperationalOperations_Success validates Mobility service
// operational data retrieval against live WNC controller.
func TestMobilityServiceIntegration_GetOperationalOperations_Success(t *testing.T) {
	t.Parallel()

	// Define the test suite configuration
	suite := integration.TestSuite{
		Config: integration.TestSuiteConfig{
			ServiceName: "Mobility",
			ServiceConstructor: func(client any) any {
				return mobility.NewService(client.(*core.Client))
			},
			UseTimeout: true,
		},
		BasicMethods: []integration.TestMethod{
			{
				Name: "GetOperational",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(mobility.Service).GetOperational(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListAPCache",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(mobility.Service).ListAPCache(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListAPPeers",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(mobility.Service).ListAPPeers(ctx)
				},
				LogResult: true,
			},
			{
				Name: "GetMMGlobalInfo",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(mobility.Service).GetMMGlobalInfo(ctx)
				},
				LogResult: true,
			},
			{
				Name: "GetMMIFGlobalStats",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(mobility.Service).GetMMIFGlobalStats(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListClients",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(mobility.Service).ListClients(ctx)
				},
				LogResult: true,
			},
			{
				Name: "GetGlobalStats",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(mobility.Service).GetGlobalStats(ctx)
				},
				LogResult: true,
			},
		},
		FilterMethods:   []integration.TestMethod{},
		ValidationTests: []integration.ValidationTestMethod{},
	}

	integration.RunTestSuite(t, suite)
}
