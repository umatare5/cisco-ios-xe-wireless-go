//go:build integration

package integration_test

import (
	"context"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/mesh"
	"github.com/umatare5/cisco-ios-xe-wireless-go/tests/testutil/integration"
)

// TestMeshServiceIntegration_GetOperationalOperations_Success validates Mesh service
// operational data retrieval against live WNC controller.
func TestMeshServiceIntegration_GetOperationalOperations_Success(t *testing.T) {
	t.Parallel()

	// Define the test suite configuration
	suite := integration.TestSuite{
		Config: integration.TestSuiteConfig{
			ServiceName: "Mesh",
			ServiceConstructor: func(client any) any {
				return mesh.NewService(client.(*core.Client))
			},
			UseTimeout: true,
		},
		BasicMethods: []integration.TestMethod{
			{
				Name: "GetOperational",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(mesh.Service).GetOperational(ctx)
				},
				LogResult: true,
			},
			{
				Name: "GetOperationalData",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(mesh.Service).GetOperationalData(ctx)
				},
				LogResult: true,
			},
			{
				Name: "GetConfig",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(mesh.Service).GetConfig(ctx)
				},
				LogResult: true,
			},
			{
				Name: "GetGlobalStats",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(mesh.Service).GetGlobalStats(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListApCacInfo",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(mesh.Service).ListApCacInfo(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListApPathInfo",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(mesh.Service).ListApPathInfo(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListApTreeData",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(mesh.Service).ListApTreeData(ctx)
				},
				LogResult: true,
			},
		},
		FilterMethods:   []integration.TestMethod{},
		ValidationTests: []integration.ValidationTestMethod{},
	}

	integration.RunTestSuite(t, suite)
}
