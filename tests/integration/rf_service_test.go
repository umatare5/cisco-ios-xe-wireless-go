//go:build integration

package integration_test

import (
	"context"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/rf"
)

// TestRFServiceIntegration_GetConfigOperations_Success validates RF service
// configuration retrieval against live WNC controller.
func TestRFServiceIntegration_GetConfigOperations_Success(t *testing.T) {
	t.Parallel() // Safe for parallel execution as read-only operations
	suite := client.IntegrationTestSuite{
		Config: client.TestSuiteConfig{
			ServiceName: "RF",
			ServiceConstructor: func(client any) any {
				return rf.NewService(client.(*core.Client))
			},
			UseTimeout: true,
		},
		BasicMethods: []client.IntegrationTestMethod{
			{
				Name: "GetConfig",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(rf.Service).GetConfig(ctx)
				},
				LogResult: true,
			},
		},
	}

	client.RunIntegrationTestSuite(t, suite)
}
