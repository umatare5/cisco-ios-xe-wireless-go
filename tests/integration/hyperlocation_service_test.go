//go:build integration

package integration_test

import (
	"context"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/hyperlocation"
)

// TestHyperlocationServiceIntegration_GetOperationalOperations_Success validates Hyperlocation service
// operational data retrieval against live WNC controller.
func TestHyperlocationServiceIntegration_GetOperationalOperations_Success(t *testing.T) {
	t.Parallel() // Safe for parallel execution as read-only operations
	suite := client.IntegrationTestSuite{
		Config: client.TestSuiteConfig{
			ServiceName: "Hyperlocation",
			ServiceConstructor: func(client any) any {
				return hyperlocation.NewService(client.(*core.Client))
			},
			UseTimeout: true,
		},
		BasicMethods: []client.IntegrationTestMethod{
			{
				Name: "GetOperational",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(hyperlocation.Service).GetOperational(ctx)
				},
			},
			{
				Name: "ListProfiles",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(hyperlocation.Service).ListProfiles(ctx)
				},
			},
		},
		FilterMethods:   []client.IntegrationTestMethod{},
		ValidationTests: []client.ValidationTestMethod{},
	}

	client.RunIntegrationTestSuite(t, suite)
}
