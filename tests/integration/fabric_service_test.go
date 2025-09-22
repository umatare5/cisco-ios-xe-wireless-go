//go:build integration

package integration_test

import (
	"context"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/fabric"
	"github.com/umatare5/cisco-ios-xe-wireless-go/tests/testutil/integration"
)

// TestFabricServiceIntegration_GetConfigOperations_Success validates Fabric service
// configuration retrieval against live WNC controller.
func TestFabricServiceIntegration_GetConfigOperations_Success(t *testing.T) {
	t.Parallel()

	// Define the test suite configuration
	suite := integration.TestSuite{
		Config: integration.TestSuiteConfig{
			ServiceName: "Fabric",
			ServiceConstructor: func(client any) any {
				return fabric.NewService(client.(*core.Client))
			},
			UseTimeout: true,
		},
		BasicMethods: []integration.TestMethod{
			{
				Name: "GetConfig",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(fabric.Service).GetConfig(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListCfgFabric",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(fabric.Service).ListCfgFabric(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListCfgFabricControlplaneNames",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(fabric.Service).ListCfgFabricControlplaneNames(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListFabricConfig",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(fabric.Service).ListFabricConfig(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListFabricControlplanes",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(fabric.Service).ListFabricControlplanes(ctx)
				},
				LogResult: true,
			},
		},
		ValidationTests: []integration.ValidationTestMethod{},
	}

	integration.RunTestSuite(t, suite)
}
