//go:build integration

package integration_test

import (
	"context"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/location"
)

// TestLocationServiceIntegration_GetOperationalOperations_Success validates Location service
// operational data retrieval against live WNC controller.
func TestLocationServiceIntegration_GetOperationalOperations_Success(t *testing.T) {
	t.Parallel() // Safe for parallel execution as read-only operations
	suite := client.IntegrationTestSuite{
		Config: client.TestSuiteConfig{
			ServiceName: "Location",
			ServiceConstructor: func(client any) any {
				return location.NewService(client.(*core.Client))
			},
			UseTimeout: true,
		},
		BasicMethods: []client.IntegrationTestMethod{
			{
				Name: "GetOperational",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(location.Service).GetOperational(ctx)
				},
				LogResult:      true,
				ExpectNotFound: true, // Location services may not be configured
			},
			{
				Name: "GetStats",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(location.Service).GetStats(ctx)
				},
				LogResult:      true,
				ExpectNotFound: true,
			},
			{
				Name: "GetConfig",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(location.Service).GetConfig(ctx)
				},
				LogResult:      true,
				ExpectNotFound: true,
			},
			{
				Name: "ListProfileConfigs",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(location.Service).ListProfileConfigs(ctx)
				},
				LogResult:      true,
				ExpectNotFound: true,
			},
			{
				Name: "ListServerConfigs",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(location.Service).ListServerConfigs(ctx)
				},
				LogResult:      true,
				ExpectNotFound: true,
			},
			{
				Name: "GetSettingsConfig",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(location.Service).GetSettingsConfig(ctx)
				},
				LogResult:      true,
				ExpectNotFound: true,
			},
		},
		FilterMethods:   []client.IntegrationTestMethod{},
		ValidationTests: []client.ValidationTestMethod{},
	}

	client.RunIntegrationTestSuite(t, suite)
}
