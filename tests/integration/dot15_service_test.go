//go:build integration

package integration_test

import (
	"context"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/dot15"
)

// TestDot15ServiceIntegration_GetConfigOperations_Success validates 802.15.4 service
// configuration retrieval against live WNC controller.
func TestDot15ServiceIntegration_GetConfigOperations_Success(t *testing.T) {
	t.Parallel() // Safe for parallel execution as read-only operations
	suite := client.IntegrationTestSuite{
		Config: client.TestSuiteConfig{
			ServiceName: "DOT15",
			ServiceConstructor: func(client any) any {
				return dot15.NewService(client.(*core.Client))
			},
			UseTimeout: true,
		},
		BasicMethods: []client.IntegrationTestMethod{
			{
				Name: "GetConfig",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(dot15.Service).GetConfig(ctx)
				},
				LogResult:      true,
				ExpectNotFound: true, // 802.15 may not be configured
			},
		},
	}

	client.RunIntegrationTestSuite(t, suite)
}
