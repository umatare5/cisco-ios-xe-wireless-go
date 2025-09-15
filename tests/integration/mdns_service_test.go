//go:build integration

package integration_test

import (
	"context"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/mdns"
	"github.com/umatare5/cisco-ios-xe-wireless-go/tests/testutil/integration"
)

// TestMDNSServiceIntegration_GetOperationalOperations_Success validates mDNS service
// operational data retrieval against live WNC controller.
func TestMDNSServiceIntegration_GetOperationalOperations_Success(t *testing.T) {
	t.Parallel()

	// Define the test suite configuration
	suite := integration.TestSuite{
		Config: integration.TestSuiteConfig{
			ServiceName: "MDNS",
			ServiceConstructor: func(client any) any {
				return mdns.NewService(client.(*core.Client))
			},
			UseTimeout: true,
		},
		BasicMethods: []integration.TestMethod{
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
		FilterMethods:   []integration.TestMethod{},
		ValidationTests: []integration.ValidationTestMethod{},
	}

	integration.RunTestSuite(t, suite)
}
