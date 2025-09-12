//go:build integration

package integration_test

import (
	"context"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/rrm"
)

// TestRRMServiceIntegration_GetConfigOperations_Success validates RRM service
// configuration retrieval against live WNC controller.
func TestRRMServiceIntegration_GetConfigOperations_Success(t *testing.T) {
	t.Parallel() // Safe for parallel execution as read-only operations
	suite := client.IntegrationTestSuite{
		Config: client.TestSuiteConfig{
			ServiceName: "RRM",
			ServiceConstructor: func(client any) any {
				return rrm.NewService(client.(*core.Client))
			},
			UseTimeout: true,
		},
		BasicMethods: []client.IntegrationTestMethod{
			{
				Name: "GetConfig",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(rrm.Service).GetConfig(ctx)
				},
			},
		},
		ValidationTests: []client.ValidationTestMethod{},
	}

	client.RunIntegrationTestSuite(t, suite)
}

// TestRRMServiceIntegration_GetOperationalOperations_Success validates RRM service
// operational data retrieval against live WNC controller.
func TestRRMServiceIntegration_GetOperationalOperations_Success(t *testing.T) {
	t.Parallel() // Safe for parallel execution as read-only operations
	suite := client.IntegrationTestSuite{
		Config: client.TestSuiteConfig{
			ServiceName: "RRM",
			ServiceConstructor: func(client any) any {
				return rrm.NewService(client.(*core.Client))
			},
			UseTimeout: true,
		},
		BasicMethods: []client.IntegrationTestMethod{
			{
				Name: "GetOperational",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(rrm.Service).GetOperational(ctx)
				},
				LogResult: true,
			},
			{
				Name: "GetGlobalInfo",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(rrm.Service).GetGlobalInfo(ctx)
				},
				LogResult: true,
			},
			{
				Name: "GetEmulationInfo",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(rrm.Service).GetEmulationInfo(ctx)
				},
				LogResult: true,
			},
		},
		ValidationTests: []client.ValidationTestMethod{},
	}

	client.RunIntegrationTestSuite(t, suite)
}
