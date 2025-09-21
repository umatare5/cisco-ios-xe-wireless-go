//go:build integration

package integration_test

import (
	"context"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/ap"
	"github.com/umatare5/cisco-ios-xe-wireless-go/tests/testutil/integration"
)

// TestAPServiceIntegration_GetConfigOperations_Success validates AP service
// configuration retrieval operations against live WNC controller.
func TestAPServiceIntegration_GetConfigOperations_Success(t *testing.T) {
	t.Parallel()

	suite := integration.TestSuite{
		Config: integration.TestSuiteConfig{
			ServiceName: "AP Get Config",
			ServiceConstructor: func(client any) any {
				return ap.NewService(client.(*core.Client))
			},
		},
		BasicMethods: []integration.TestMethod{
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
		FilterMethods: []integration.TestMethod{
			{
				Name: "GetTagConfigByMAC",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(ap.Service).GetTagConfigByMAC(ctx, "28:ac:9e:11:48:10")
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
		ValidationTests: []integration.ValidationTestMethod{
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

	integration.RunTestSuite(t, suite)
}

func TestAPServiceIntegration_GetOperationalOperations_Success(t *testing.T) {
	suite := integration.TestSuite{
		Config: integration.TestSuiteConfig{
			ServiceName: "AP Get Operational",
			ServiceConstructor: func(client any) any {
				return ap.NewService(client.(*core.Client))
			},
		},
		BasicMethods: []integration.TestMethod{
			{
				Name: "GetOperational",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(ap.Service).GetOperational(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListApOperData",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(ap.Service).ListApOperData(ctx)
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
			{
				Name: "ListIotFirmware",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(ap.Service).ListIotFirmware(ctx)
				},
				LogResult: true,
			},
		},
		FilterMethods: []integration.TestMethod{
			{
				Name: "GetCAPWAPDataByWTPMAC",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(ap.Service).GetCAPWAPDataByWTPMAC(ctx, integration.TestAPMac())
				},
				LogResult:      true,
				ExpectNotFound: true, // May not have specific WTP MAC data available
			},
			{
				Name: "ListAPHistoryByEthernetMAC",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(ap.Service).ListAPHistoryByEthernetMAC(ctx, "28:ac:9e:11:48:10")
				},
				LogResult: true,
			},
			{
				Name: "GetAPJoinStatsByWTPMAC",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(ap.Service).GetAPJoinStatsByWTPMAC(ctx, integration.TestAPMac())
				},
				LogResult: true,
			},
			{
				Name: "GetWLANClientStatsByWLANID",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(ap.Service).GetWLANClientStatsByWLANID(ctx, 1)
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
					return service.(ap.Service).GetRadioStatusByWTPMACAndSlot(ctx, integration.TestAPMac(), 0)
				},
				LogResult:      true,
				ExpectNotFound: true, // May not have specific slot data available
			},
		},
		ValidationTests: []integration.ValidationTestMethod{
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

	integration.RunTestSuite(t, suite)
}

// Test_ApGlobalOper_IntegrationTests validates AP service global operational data retrieval
func TestAPServiceIntegration_GlobalOperations_Success(t *testing.T) {
	// Define the test suite configuration
	suite := integration.TestSuite{
		Config: integration.TestSuiteConfig{
			ServiceName: "AP Global Operational",
			ServiceConstructor: func(client any) any {
				return ap.NewService(client.(*core.Client))
			},
			UseTimeout: true,
		},
		BasicMethods: []integration.TestMethod{
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
		FilterMethods:   []integration.TestMethod{},
		ValidationTests: []integration.ValidationTestMethod{},
	}

	// Run the unified test suite
	integration.RunTestSuite(t, suite)
}

// TestAPServiceIntegration_AdvancedFilterOperations_Success validates advanced
// AP service filtering operations with complex parameters.
func TestAPServiceIntegration_AdvancedFilterOperations_Success(t *testing.T) {
	t.Parallel()

	suite := integration.TestSuite{
		Config: integration.TestSuiteConfig{
			ServiceName: "AP Advanced Filter Operations",
			ServiceConstructor: func(client any) any {
				return ap.NewService(client.(*core.Client))
			},
			UseTimeout: true,
		},
		FilterMethods: []integration.TestMethod{
			{
				Name: "GetRadioNeighborByAPMACSlotAndBSSID",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(ap.Service).GetRadioNeighborByAPMACSlotAndBSSID(
						ctx, integration.TestAPMac(), 0, integration.TestAPNeighborBSSID())
				},
				LogResult: true,
				WhenError: skipOnNeighborBSSIDNotFound,
			},
		},
		ValidationTests: []integration.ValidationTestMethod{
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
						ctx, integration.TestAPMac(), 0, "")
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

	integration.RunTestSuite(t, suite)
}

// skipOnNeighborBSSIDNotFound skips the test when neighbor BSSID is not found
func skipOnNeighborBSSIDNotFound(t *testing.T, methodName string, err error) bool {
	// Only handle NotFound errors for neighbor BSSID tests
	if !core.IsNotFoundError(err) {
		return false
	}

	t.Skipf("%s: Neighbor BSSID not found - this may caused as neighbor data changes frequently. "+
		"To update with current data, run: 'go run ./example/list_neighbors/main.go' "+
		"and set WNC_AP_NEIGHBOR_BSSID environment variable. "+
		"Error details: %v", methodName, err)
	return true
}
