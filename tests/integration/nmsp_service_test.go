//go:build integration

package integration_test

import (
	"context"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/nmsp"
	"github.com/umatare5/cisco-ios-xe-wireless-go/tests/testutil/integration"
)

// TestNMSPServiceIntegration_GetOperationalOperations_Success validates NMSP service
// operational data retrieval against live WNC controller.
func TestNMSPServiceIntegration_GetOperationalOperations_Success(t *testing.T) {
	t.Parallel()

	// Define the test suite configuration
	suite := integration.TestSuite{
		Config: integration.TestSuiteConfig{
			ServiceName: "NMSP",
			ServiceConstructor: func(client any) any {
				return nmsp.NewService(client.(*core.Client))
			},
			UseTimeout: true,
		},
		BasicMethods: []integration.TestMethod{
			{
				Name: "GetOperational",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(nmsp.Service).GetOperational(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListClientRegistrations",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(nmsp.Service).ListClientRegistrations(ctx)
				},
				LogResult: true,
			},
			{
				Name: "GetCMXConnectionInfo",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(nmsp.Service).GetCMXConnectionInfo(ctx)
				},
				LogResult: true,
			},
			{
				Name: "GetCMXCloudInfo",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(nmsp.Service).GetCMXCloudInfo(ctx)
				},
				LogResult: true,
			},
		},
		FilterMethods:   []integration.TestMethod{},
		ValidationTests: []integration.ValidationTestMethod{},
	}

	integration.RunTestSuite(t, suite)
}
