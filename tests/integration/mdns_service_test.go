//go:build integration

package integration_test

import (
	"context"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/mdns"
)

// TestMDNSServiceIntegration_GetOperationalOperations_Success validates mDNS service
// operational data retrieval against live WNC controller.
func TestMDNSServiceIntegration_GetOperationalOperations_Success(t *testing.T) {
	t.Parallel() // Safe for parallel execution as read-only operations
	suite := client.IntegrationTestSuite{
		Config: client.TestSuiteConfig{
			ServiceName: "MDNS",
			ServiceConstructor: func(client any) any {
				return mdns.NewService(client.(*core.Client))
			},
			UseTimeout: true,
		},
		BasicMethods: []client.IntegrationTestMethod{
			{
				Name: "GetOperational",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(mdns.Service).GetOperational(ctx)
				},
				LogResult:      true,
				ExpectNotFound: true, // mDNS may not be enabled
			},
			{
				Name: "GetGlobalStats",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(mdns.Service).GetGlobalStats(ctx)
				},
				LogResult:      true,
				ExpectNotFound: true,
			},
			{
				Name: "ListWLANStats",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(mdns.Service).ListWLANStats(ctx)
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
