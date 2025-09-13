//go:build integration

package integration_test

import (
	"context"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/rogue"
	"github.com/umatare5/cisco-ios-xe-wireless-go/tests/testutil/integration"
)

// TestRogueServiceIntegration_GetOperationalOperations_Success validates Rogue service
// operational data retrieval against live WNC controller.
func TestRogueServiceIntegration_GetOperationalOperations_Success(t *testing.T) {
	t.Parallel()

	// Define the test suite configuration
	suite := integration.TestSuite{
		Config: integration.TestSuiteConfig{
			ServiceName: "Rogue",
			ServiceConstructor: func(client any) any {
				return rogue.NewService(client.(*core.Client))
			},
			UseTimeout: true,
		},
		BasicMethods: []integration.TestMethod{
			{
				Name: "GetOperational",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(rogue.Service).GetOperational(ctx)
				},
				LogResult: true,
			},
			{
				Name: "GetOperClientData",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(rogue.Service).GetOperClientData(ctx)
				},
				LogResult: true,
			},
			{
				Name: "GetOperData",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(rogue.Service).GetOperData(ctx)
				},
				LogResult: true,
			},
			{
				Name: "GetRLDPStats",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(rogue.Service).GetRLDPStats(ctx)
				},
				LogResult: true,
			},
			{
				Name: "GetOperStats",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(rogue.Service).GetOperStats(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListRogues",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(rogue.Service).ListRogues(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListRogueClients",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(rogue.Service).ListRogueClients(ctx)
				},
				LogResult: true,
			},
			{
				Name: "GetStats",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(rogue.Service).GetStats(ctx)
				},
				LogResult: true,
			},
		},
		FilterMethods: []integration.TestMethod{
			{
				Name: "GetRogueByMAC",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(rogue.Service).GetRogueByMAC(ctx, integration.TestClientMac())
				},
				LogResult:      true,
				ExpectNotFound: true, // Test MAC may not exist unless WNC_CLIENT_MAC_ADDR is set
			},
			{
				Name: "GetRogueClientByMAC",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(rogue.Service).GetRogueClientByMAC(ctx, "00:11:22:33:44:66")
				},
				ExpectNotFound: true,
			},
			{
				Name: "GetOperByRogueAddress",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(rogue.Service).GetOperByRogueAddress(ctx, integration.TestClientMac())
				},
				LogResult:      true,
				ExpectNotFound: true, // Test rogue address may not exist unless WNC_CLIENT_MAC_ADDR is set
			},
			{
				Name: "GetOperByRogueClientAddress",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(rogue.Service).GetOperByRogueClientAddress(ctx, "00:11:22:33:44:66")
				},
				ExpectNotFound: true,
			},
			// Note: GetOperByClassType and GetOperByContainmentLevel are disabled
			// because these query parameters are not supported by the API (returns HTTP 400)
		},
		ValidationTests: []integration.ValidationTestMethod{
			{
				Name: "GetOperByRogueAddress_EmptyAddress",
				Method: func(ctx context.Context, service any) error {
					_, err := service.(rogue.Service).GetOperByRogueAddress(ctx, "")
					return err
				},
				ExpectedError: true,
				ErrorKeywords: []string{"404", "not found"},
			},
			{
				Name: "GetOperByRogueClientAddress_EmptyAddress",
				Method: func(ctx context.Context, service any) error {
					_, err := service.(rogue.Service).GetOperByRogueClientAddress(ctx, "")
					return err
				},
				ExpectedError: true,
				ErrorKeywords: []string{"404", "not found"},
			},
			// Note: GetOperByClassType and GetOperByContainmentLevel validation tests
			// are also disabled because these query parameters are not supported
		},
	}

	integration.RunTestSuite(t, suite)
}
