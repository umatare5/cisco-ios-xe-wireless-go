//go:build integration

package integration_test

import (
	"context"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	clientService "github.com/umatare5/cisco-ios-xe-wireless-go/service/client"
)

// Test_ClientGetOperational_IntegrationTests runs comprehensive client operational data integration tests
func TestClientServiceIntegration_GetOperationalOperations_Success(t *testing.T) {
	t.Parallel() // Safe for parallel execution as read-only operations
	// Define the test suite configuration
	suite := client.IntegrationTestSuite{
		Config: client.TestSuiteConfig{
			ServiceName: "Client Operational Data",
			ServiceConstructor: func(httpClient any) any {
				return clientService.NewService(httpClient.(*core.Client))
			},
			UseTimeout: true,
		},
		BasicMethods: []client.IntegrationTestMethod{
			{
				Name: "GetOperational",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(clientService.Service).GetOperational(ctx)
				},
				LogResult:      true,
				ExpectNotFound: true, // Client data may not be available on all controllers
			},
			{
				Name: "ListCommonInfo",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(clientService.Service).ListCommonInfo(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListDot11Info",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(clientService.Service).ListDot11Info(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListDCInfo",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(clientService.Service).ListDCInfo(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListMMIFClientHistory",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(clientService.Service).ListMMIFClientHistory(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListMMIFClientStats",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(clientService.Service).ListMMIFClientStats(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListMobilityInfo",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(clientService.Service).ListMobilityInfo(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListPolicyInfo",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(clientService.Service).ListPolicyInfo(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListSISFDB",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(clientService.Service).ListSISFDB(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListTrafficStats",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(clientService.Service).ListTrafficStats(ctx)
				},
				LogResult: true,
			},
		},
		FilterMethods: []client.IntegrationTestMethod{
			{
				Name: "GetCommonInfoByMAC",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(clientService.Service).GetCommonInfoByMAC(ctx, "00:11:22:33:44:55")
				},
				LogResult:      true,
				ExpectNotFound: true, // Sample MAC may not exist
			},
			{
				Name: "GetDCInfoByMAC",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(clientService.Service).GetDCInfoByMAC(ctx, "00:11:22:33:44:55")
				},
				LogResult:      true,
				ExpectNotFound: true, // Sample MAC may not exist
			},
			{
				Name: "GetDot11InfoByMAC",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(clientService.Service).GetDot11InfoByMAC(ctx, "00:11:22:33:44:55")
				},
				LogResult:      true,
				ExpectNotFound: true, // Sample MAC may not exist
			},
			{
				Name: "GetMMIFClientHistoryByMAC",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(clientService.Service).GetMMIFClientHistoryByMAC(ctx, "00:11:22:33:44:55")
				},
				LogResult:      true,
				ExpectNotFound: true, // Sample MAC may not exist
			},
			{
				Name: "GetMMIFClientStatsByMAC",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(clientService.Service).GetMMIFClientStatsByMAC(ctx, "00:11:22:33:44:55")
				},
				LogResult:      true,
				ExpectNotFound: true, // Sample MAC may not exist
			},
			{
				Name: "GetMobilityInfoByMAC",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(clientService.Service).GetMobilityInfoByMAC(ctx, "00:11:22:33:44:55")
				},
				LogResult:      true,
				ExpectNotFound: true, // Sample MAC may not exist
			},
			{
				Name: "GetTrafficStatsByMAC",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(clientService.Service).GetTrafficStatsByMAC(ctx, "00:11:22:33:44:55")
				},
				LogResult:      true,
				ExpectNotFound: true, // Sample MAC may not exist
			},
			{
				Name: "GetPolicyInfoByMAC",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(clientService.Service).GetPolicyInfoByMAC(ctx, "00:11:22:33:44:55")
				},
				LogResult:      true,
				ExpectNotFound: true, // Sample MAC may not exist
			},
			{
				Name: "GetSISFDBByMAC",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(clientService.Service).GetSISFDBByMAC(ctx, "00:11:22:33:44:55")
				},
				LogResult:      true,
				ExpectNotFound: true, // Sample MAC may not exist
			},
		},
	}

	// Run the unified test suite
	client.RunIntegrationTestSuite(t, suite)
}
