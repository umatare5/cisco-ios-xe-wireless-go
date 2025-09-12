//go:build integration

package integration_test

import (
	"context"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/mcast"
)

// TestMcastServiceIntegration_GetOperationalOperations_Success validates Multicast service
// operational data retrieval against live WNC controller.
func TestMcastServiceIntegration_GetOperationalOperations_Success(t *testing.T) {
	t.Parallel() // Safe for parallel execution as read-only operations
	suite := client.IntegrationTestSuite{
		Config: client.TestSuiteConfig{
			ServiceName: "Mcast",
			ServiceConstructor: func(client any) any {
				return mcast.NewService(client.(*core.Client))
			},
			UseTimeout: true,
		},
		BasicMethods: []client.IntegrationTestMethod{
			{
				Name: "GetOperational",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(mcast.Service).GetOperational(ctx)
				},
				LogResult: true,
			},
			{
				Name: "GetFlexConnectMediastreamClientSummary",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(mcast.Service).GetFlexConnectMediastreamClientSummary(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListVLANL2MGIDs",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(mcast.Service).ListVLANL2MGIDs(ctx)
				},
				LogResult: true,
			},
			{
				Name: "GetFabricMediastreamClientSummary",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(mcast.Service).GetFabricMediastreamClientSummary(ctx)
				},
				LogResult: true,
			},
			{
				Name: "GetMcastMgidInfo",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(mcast.Service).GetMcastMgidInfo(ctx)
				},
				LogResult: true,
			},
			{
				Name: "GetMulticastOperData",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(mcast.Service).GetMulticastOperData(ctx)
				},
				LogResult: true,
			},
		},
		FilterMethods:   []client.IntegrationTestMethod{},
		ValidationTests: []client.ValidationTestMethod{},
	}

	client.RunIntegrationTestSuite(t, suite)
}
