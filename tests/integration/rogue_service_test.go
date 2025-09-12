//go:build integration

package integration_test

import (
	"context"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/rogue"
)

// TestRogueServiceIntegration_GetOperationalOperations_Success validates Rogue service
// operational data retrieval against live WNC controller.
func TestRogueServiceIntegration_GetOperationalOperations_Success(t *testing.T) {
	t.Parallel() // Safe for parallel execution as read-only operations
	suite := client.IntegrationTestSuite{
		Config: client.TestSuiteConfig{
			ServiceName: "Rogue",
			ServiceConstructor: func(client any) any {
				return rogue.NewService(client.(*core.Client))
			},
			UseTimeout: true,
		},
		BasicMethods: []client.IntegrationTestMethod{
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
		FilterMethods: []client.IntegrationTestMethod{
			{
				Name: "GetRogueByMAC",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(rogue.Service).GetRogueByMAC(ctx, "00:11:22:33:44:55")
				},
				ExpectNotFound: true, // Rogue may not exist
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
					return service.(rogue.Service).GetOperByRogueAddress(ctx, "00:11:22:33:44:55")
				},
				ExpectNotFound: true, // Rogue may not exist
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
		ValidationTests: []client.ValidationTestMethod{
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

	client.RunIntegrationTestSuite(t, suite)
}
