//go:build integration

package integration_test

import (
	"context"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/mesh"
)

// TestMeshServiceIntegration_GetOperationalOperations_Success validates Mesh service
// operational data retrieval against live WNC controller.
func TestMeshServiceIntegration_GetOperationalOperations_Success(t *testing.T) {
	t.Parallel() // Safe for parallel execution as read-only operations
	suite := client.IntegrationTestSuite{
		Config: client.TestSuiteConfig{
			ServiceName: "Mesh",
			ServiceConstructor: func(client any) any {
				return mesh.NewService(client.(*core.Client))
			},
			UseTimeout: true,
		},
		BasicMethods: []client.IntegrationTestMethod{
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
		},
		FilterMethods:   []client.IntegrationTestMethod{},
		ValidationTests: []client.ValidationTestMethod{},
	}

	client.RunIntegrationTestSuite(t, suite)
}
