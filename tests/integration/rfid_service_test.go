//go:build integration

package integration_test

import (
	"context"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/data"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/rfid"
)

// TestRFIDServiceIntegration_GetConfigOperations_Success validates RFID service
// configuration retrieval against live WNC controller.
func TestRFIDServiceIntegration_GetConfigOperations_Success(t *testing.T) {
	t.Parallel() // Safe for parallel execution as read-only operations
	suite := client.IntegrationTestSuite{
		Config: client.TestSuiteConfig{
			ServiceName: "RFID",
			ServiceConstructor: func(client any) any {
				return rfid.NewService(client.(*core.Client))
			},
			UseTimeout: true,
		},
		BasicMethods: []client.IntegrationTestMethod{
			{
				Name: "GetConfig",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(rfid.Service).GetConfig(ctx)
				},
				LogResult: true,
			},
			{
				Name: "GetConfigSettings",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(rfid.Service).GetConfigSettings(ctx)
				},
				ExpectNotFound: true, // RFID config endpoint may not be available
			},
		},
		FilterMethods:   []client.IntegrationTestMethod{},
		ValidationTests: []client.ValidationTestMethod{},
	}

	client.RunIntegrationTestSuite(t, suite)
}

// TestRFIDServiceIntegration_GetOperationalOperations_Success validates RFID service
// operational data retrieval against live WNC controller.
func TestRFIDServiceIntegration_GetOperationalOperations_Success(t *testing.T) {
	t.Parallel() // Safe for parallel execution as read-only operations
	constants := data.StandardTestConstants()

	suite := client.IntegrationTestSuite{
		Config: client.TestSuiteConfig{
			ServiceName: "RFID",
			ServiceConstructor: func(client any) any {
				return rfid.NewService(client.(*core.Client))
			},
			UseTimeout: true,
		},
		BasicMethods: []client.IntegrationTestMethod{
			{
				Name: "GetOperational",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(rfid.Service).GetOperational(ctx)
				},
				LogResult: true,
			},
			{
				Name: "GetGlobalInfo",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(rfid.Service).GetGlobalInfo(ctx)
				},
				LogResult: true,
			},
		},
		FilterMethods: []client.IntegrationTestMethod{
			{
				Name: "GetRadioInfo",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(rfid.Service).GetRadioInfo(ctx, constants.TestAPMac, constants.TestAPMac, constants.TestSlotID)
				},
				ExpectNotFound: true,
			},
			{
				Name: "GetDetailByMAC",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(rfid.Service).GetDetailByMAC(ctx, constants.TestAPMac)
				},
				ExpectNotFound: true,
			},
		},
		ValidationTests: []client.ValidationTestMethod{},
	}

	client.RunIntegrationTestSuite(t, suite)
}
