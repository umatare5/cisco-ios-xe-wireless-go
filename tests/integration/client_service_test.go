//go:build integration

package integration_test

import (
	"context"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	clientService "github.com/umatare5/cisco-ios-xe-wireless-go/service/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/tests/testutil/integration"
)

// Test_ClientGetOperational_IntegrationTests runs comprehensive client operational data integration tests
func TestClientServiceIntegration_GetOperationalOperations_Success(t *testing.T) {
	// Run tests as sequentially because client oper-data endpoints contain large datasets
	// that cause resource contention when accessed concurrently during bulk test execution.
	// t.Parallel()

	// Define the test suite configuration
	suite := integration.TestSuite{
		Config: integration.TestSuiteConfig{
			ServiceName: "Client Operational Data",
			ServiceConstructor: func(httpClient any) any {
				return clientService.NewService(httpClient.(*core.Client))
			},
			UseTimeout: true,
		},
		BasicMethods: []integration.TestMethod{
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
		FilterMethods: []integration.TestMethod{
			{
				Name: "GetCommonInfoByMAC",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(clientService.Service).GetCommonInfoByMAC(ctx, integration.TestClientMac())
				},
				LogResult:      true,
				ExpectNotFound: true, // Test MAC may not exist unless WNC_CLIENT_MAC_ADDR is set
			},
			{
				Name: "GetDCInfoByMAC",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(clientService.Service).GetDCInfoByMAC(ctx, integration.TestClientMac())
				},
				LogResult:      true,
				ExpectNotFound: true, // Test MAC may not exist unless WNC_CLIENT_MAC_ADDR is set
			},
			{
				Name: "GetDot11InfoByMAC",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(clientService.Service).GetDot11InfoByMAC(ctx, integration.TestClientMac())
				},
				LogResult:      true,
				ExpectNotFound: true, // Test MAC may not exist unless WNC_CLIENT_MAC_ADDR is set
			},
			{
				Name: "GetMMIFClientHistoryByMAC",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(clientService.Service).GetMMIFClientHistoryByMAC(ctx, integration.TestClientMac())
				},
				LogResult:      true,
				ExpectNotFound: true, // Test MAC may not exist unless WNC_CLIENT_MAC_ADDR is set
			},
			{
				Name: "GetMMIFClientStatsByMAC",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(clientService.Service).GetMMIFClientStatsByMAC(ctx, integration.TestClientMac())
				},
				LogResult:      true,
				ExpectNotFound: true, // Test MAC may not exist unless WNC_CLIENT_MAC_ADDR is set
			},
			{
				Name: "GetMobilityInfoByMAC",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(clientService.Service).GetMobilityInfoByMAC(ctx, integration.TestClientMac())
				},
				LogResult:      true,
				ExpectNotFound: true, // Test MAC may not exist unless WNC_CLIENT_MAC_ADDR is set
			},
			{
				Name: "GetTrafficStatsByMAC",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(clientService.Service).GetTrafficStatsByMAC(ctx, integration.TestClientMac())
				},
				LogResult:      true,
				ExpectNotFound: true, // Test MAC may not exist unless WNC_CLIENT_MAC_ADDR is set
			},
			{
				Name: "GetPolicyInfoByMAC",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(clientService.Service).GetPolicyInfoByMAC(ctx, integration.TestClientMac())
				},
				LogResult:      true,
				ExpectNotFound: true, // Test MAC may not exist unless WNC_CLIENT_MAC_ADDR is set
			},
			{
				Name: "GetSISFDBByMAC",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(clientService.Service).GetSISFDBByMAC(ctx, integration.TestClientMac())
				},
				LogResult:      true,
				ExpectNotFound: true, // Test MAC may not exist unless WNC_CLIENT_MAC_ADDR is set
			},
		},
	}

	// Run the unified test suite
	integration.RunTestSuite(t, suite)
}
