//go:build integration

package integration_test

import (
	"context"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/nmsp"
)

// TestNMSPServiceIntegration_GetOperationalOperations_Success validates NMSP service
// operational data retrieval against live WNC controller.
func TestNMSPServiceIntegration_GetOperationalOperations_Success(t *testing.T) {
	t.Parallel() // Safe for parallel execution as read-only operations
	suite := client.IntegrationTestSuite{
		Config: client.TestSuiteConfig{
			ServiceName: "NMSP",
			ServiceConstructor: func(client any) any {
				return nmsp.NewService(client.(*core.Client))
			},
			UseTimeout: true,
		},
		BasicMethods: []client.IntegrationTestMethod{
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
		FilterMethods:   []client.IntegrationTestMethod{},
		ValidationTests: []client.ValidationTestMethod{},
	}

	client.RunIntegrationTestSuite(t, suite)
}
