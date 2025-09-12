//go:build integration

package integration_test

import (
	"context"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/radio"
)

// TestRadioServiceIntegration_GetConfigOperations_Success validates Radio service
// configuration retrieval against live WNC controller.
func TestRadioServiceIntegration_GetConfigOperations_Success(t *testing.T) {
	t.Parallel() // Safe for parallel execution as read-only operations
	suite := client.IntegrationTestSuite{
		Config: client.TestSuiteConfig{
			ServiceName: "Radio",
			ServiceConstructor: func(client any) any {
				return radio.NewService(client.(*core.Client))
			},
			UseTimeout: true,
		},
		BasicMethods: []client.IntegrationTestMethod{
			{
				Name: "GetConfig",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(radio.Service).GetConfig(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListProfileConfigs",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(radio.Service).ListProfileConfigs(ctx)
				},
				LogResult: true,
			},
		},
		FilterMethods:   []client.IntegrationTestMethod{},
		ValidationTests: []client.ValidationTestMethod{},
	}

	client.RunIntegrationTestSuite(t, suite)
}
