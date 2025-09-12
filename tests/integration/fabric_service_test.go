//go:build integration

package integration_test

import (
	"context"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/fabric"
)

// TestFabricServiceIntegration_GetConfigOperations_Success validates Fabric service
// configuration retrieval against live WNC controller.
func TestFabricServiceIntegration_GetConfigOperations_Success(t *testing.T) {
	t.Parallel() // Safe for parallel execution as read-only operations
	suite := client.IntegrationTestSuite{
		Config: client.TestSuiteConfig{
			ServiceName: "Fabric",
			ServiceConstructor: func(client any) any {
				return fabric.NewService(client.(*core.Client))
			},
			UseTimeout: true,
		},
		BasicMethods: []client.IntegrationTestMethod{
			{
				Name: "GetConfig",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(fabric.Service).GetConfig(ctx)
				},
				LogResult: true,
			},
		},
		ValidationTests: []client.ValidationTestMethod{},
	}

	client.RunIntegrationTestSuite(t, suite)
}
