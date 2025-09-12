//go:build integration

package integration_test

import (
	"context"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/data"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/ap"
)

// TestAPServiceIntegration_GetConfigOperations_Success validates AP service
// configuration retrieval operations against live WNC controller.
//
// This test verifies that basic configuration operations (GetConfig, ListTagConfigs,
// ListTagSourcePriorityConfigs) return valid data structures and can communicate
// with the WNC API endpoint successfully.
//
// Test Coverage:
//   - Basic configuration retrieval methods
//   - Filter-based configuration operations
//   - Validation of returned data structures
//   - Live WNC controller connectivity
func TestAPServiceIntegration_GetConfigOperations_Success(t *testing.T) {
	t.Parallel() // Safe for parallel execution as read-only operations

	constants := data.StandardTestConstants()

	suite := client.IntegrationTestSuite{
		Config: client.TestSuiteConfig{
			ServiceName: "AP Get Config",
			ServiceConstructor: func(client any) any {
				return ap.NewService(client.(*core.Client))
			},
		},
		BasicMethods: []client.IntegrationTestMethod{
			{
				Name: "GetConfig",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(ap.Service).GetConfig(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListTagConfigs",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(ap.Service).ListTagConfigs(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListTagSourcePriorityConfigs",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(ap.Service).ListTagSourcePriorityConfigs(ctx)
				},
				LogResult: true,
			},
		},
		FilterMethods: []client.IntegrationTestMethod{
			{
				Name: "GetTagConfigByMAC",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(ap.Service).GetTagConfigByMAC(ctx, constants.TestEthMAC)
				},
				LogResult: true,
			},
			{
				Name: "GetTagSourcePriorityConfigByPriority",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(ap.Service).GetTagSourcePriorityConfigByPriority(ctx, 1)
				},
				LogResult: true,
			},
		},
		ValidationTests: []client.ValidationTestMethod{
			{
				Name: "GetTagConfigByMAC_InvalidMAC",
				Method: func(ctx context.Context, service any) error {
					_, err := service.(ap.Service).GetTagConfigByMAC(ctx, "invalid-mac")
					return err
				},
				ExpectedError: true,
				ErrorKeywords: []string{"invalid", "mac"},
			},
			{
				Name: "GetTagConfigByMAC_EmptyMAC",
				Method: func(ctx context.Context, service any) error {
					_, err := service.(ap.Service).GetTagConfigByMAC(ctx, "")
					return err
				},
				ExpectedError: true,
				ErrorKeywords: []string{"empty", "mac"},
			},
		},
	}

	client.RunIntegrationTestSuite(t, suite)
}

func TestAPServiceIntegration_GetOperationalOperations_Success(t *testing.T) {
	constants := data.StandardTestConstants()

	suite := client.IntegrationTestSuite{
		Config: client.TestSuiteConfig{
			ServiceName: "AP Get Operational",
			ServiceConstructor: func(client any) any {
				return ap.NewService(client.(*core.Client))
			},
		},
		BasicMethods: []client.IntegrationTestMethod{
			{
				Name: "GetOperational",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(ap.Service).GetOperational(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListCAPWAPData",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(ap.Service).ListCAPWAPData(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListRadioStatus",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(ap.Service).ListRadioStatus(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListNameMACMaps",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(ap.Service).ListNameMACMaps(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListRadioNeighbors",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(ap.Service).ListRadioNeighbors(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListActiveImageLocations",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(ap.Service).ListActiveImageLocations(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListPreparedImageLocations",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(ap.Service).ListPreparedImageLocations(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListPowerInfo",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(ap.Service).ListPowerInfo(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListSensorStatus",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(ap.Service).ListSensorStatus(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListCAPWAPPackets",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(ap.Service).ListCAPWAPPackets(ctx)
				},
				LogResult: true,
			},
		},
		FilterMethods: []client.IntegrationTestMethod{
			{
				Name: "GetCAPWAPDataByWTPMAC",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(ap.Service).GetCAPWAPDataByWTPMAC(ctx, constants.TestAPMac)
				},
				LogResult:      true,
				ExpectNotFound: true, // May not have specific WTP MAC data available
			},
			{
				Name: "ListAPHistoryByEthernetMAC",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(ap.Service).ListAPHistoryByEthernetMAC(ctx, constants.TestEthMAC)
				},
				LogResult: true,
			},
			{
				Name: "GetAPJoinStatsByWTPMAC",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(ap.Service).GetAPJoinStatsByWTPMAC(ctx, constants.TestAPMac)
				},
				LogResult: true,
			},
			{
				Name: "GetWLANClientStatsByWLANID",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(ap.Service).GetWLANClientStatsByWLANID(ctx, constants.TestWlanID)
				},
				LogResult: true,
			},
			{
				Name: "GetNameMACMapByWTPName",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(ap.Service).GetNameMACMapByWTPName(ctx, "TEST-AP01")
				},
				LogResult:      true,
				ExpectNotFound: true, // May not have specific AP name data available
			},
			{
				Name: "GetRadioStatusByWTPMACAndSlot",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(ap.Service).GetRadioStatusByWTPMACAndSlot(ctx, constants.TestAPMac, 0)
				},
				LogResult:      true,
				ExpectNotFound: true, // May not have specific slot data available
			},
		},
		ValidationTests: []client.ValidationTestMethod{
			{
				Name: "GetCAPWAPDataByWTPMAC_EmptyMAC",
				Method: func(ctx context.Context, service any) error {
					_, err := service.(ap.Service).GetCAPWAPDataByWTPMAC(ctx, "")
					return err
				},
				ExpectedError: true,
				ErrorKeywords: []string{"404", "not found"},
			},
		},
	}

	client.RunIntegrationTestSuite(t, suite)
}

// Test_ApGlobalOper_IntegrationTests runs comprehensive AP global operational data integration tests
func TestAPServiceIntegration_GlobalOperations_Success(t *testing.T) {
	// Define the test suite configuration
	suite := client.IntegrationTestSuite{
		Config: client.TestSuiteConfig{
			ServiceName: "AP Global Operational",
			ServiceConstructor: func(client any) any {
				return ap.NewService(client.(*core.Client))
			},
			UseTimeout: true,
		},
		BasicMethods: []client.IntegrationTestMethod{
			{
				Name: "GetGlobalInfo",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(ap.Service).GetGlobalInfo(ctx)
				},
				LogResult: true,
			},
			{
				Name: "GetEWLCAPStats",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(ap.Service).GetEWLCAPStats(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListAPHistory",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(ap.Service).ListAPHistory(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListAPJoinStats",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(ap.Service).ListAPJoinStats(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListWLANClientStats",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(ap.Service).ListWLANClientStats(ctx)
				},
				LogResult: true,
			},
		},
		FilterMethods:   []client.IntegrationTestMethod{},
		ValidationTests: []client.ValidationTestMethod{},
	}

	// Run the unified test suite
	client.RunIntegrationTestSuite(t, suite)
}

// TestAPServiceIntegration_AdvancedFilterOperations_Success validates advanced
// AP service filtering operations with complex parameters.
//
// This test covers advanced filtering operations that require multiple parameters:
//   - Radio neighbor operations with MAC/slot/BSSID filtering
//   - Complex composite key operations
//   - Multi-parameter validation scenarios
//
// Test Coverage:
//   - Complex parameter filtering against live WNC
//   - Composite key parameter validation
//   - Advanced error handling scenarios
//   - Real data structure validation with complex filters
func TestAPServiceIntegration_AdvancedFilterOperations_Success(t *testing.T) {
	constants := data.StandardTestConstants()

	suite := client.IntegrationTestSuite{
		Config: client.TestSuiteConfig{
			ServiceName: "AP Advanced Filter Operations",
			ServiceConstructor: func(client any) any {
				return ap.NewService(client.(*core.Client))
			},
			UseTimeout: true,
		},
		FilterMethods: []client.IntegrationTestMethod{
			{
				Name: "GetRadioNeighborByAPMACSlotAndBSSID",
				Method: func(ctx context.Context, service any) (any, error) {
					// Use actual BSSID from live data: 98:f1:99:c2:03:db (slot 0)
					testBSSID := "98:f1:99:c2:03:db"
					return service.(ap.Service).GetRadioNeighborByAPMACSlotAndBSSID(
						ctx, constants.TestAPMac, 0, testBSSID)
				},
				LogResult: true,
			},
		},
		ValidationTests: []client.ValidationTestMethod{
			{
				Name: "GetRadioNeighborByAPMACSlotAndBSSID_EmptyMAC",
				Method: func(ctx context.Context, service any) error {
					_, err := service.(ap.Service).GetRadioNeighborByAPMACSlotAndBSSID(
						ctx, "", 0, "aa:bb:cc:dd:ee:ff")
					return err
				},
				ExpectedError: true,
				ErrorKeywords: []string{"empty", "mac"},
			},
			{
				Name: "GetRadioNeighborByAPMACSlotAndBSSID_EmptyBSSID",
				Method: func(ctx context.Context, service any) error {
					_, err := service.(ap.Service).GetRadioNeighborByAPMACSlotAndBSSID(
						ctx, constants.TestAPMac, 0, "")
					return err
				},
				ExpectedError: true,
				ErrorKeywords: []string{"empty", "bssid"},
			},
			{
				Name: "GetRadioNeighborByAPMACSlotAndBSSID_InvalidMAC",
				Method: func(ctx context.Context, service any) error {
					_, err := service.(ap.Service).GetRadioNeighborByAPMACSlotAndBSSID(
						ctx, "invalid-mac", 0, "aa:bb:cc:dd:ee:ff")
					return err
				},
				ExpectedError: true,
				ErrorKeywords: []string{"invalid", "mac"},
			},
		},
	}

	client.RunIntegrationTestSuite(t, suite)
}
