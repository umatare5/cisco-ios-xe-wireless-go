package ap

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/framework"
)

// Test constants
const (
	testAPMac       = "aa:bb:cc:dd:ee:ff"
	testNeighborMAC = "bb:cc:dd:ee:ff:aa"
	testWtpName     = "test-wtp"
	testSlotID      = 0
)

func Test_ApGetOper_StandardTests(t *testing.T) {
	testClient := client.OptionalTestClient(t)
	service := NewService(testClient)
	nilService := NewService(nil)
	ctx := client.TestContext(t)

	framework.RunStandardGetOperTests(t, framework.GetOperTestSuite{
		ServiceName: "AP",
		BasicMethods: []framework.GetOperTestMethod{
			{
				Name: "GetOper",
				Method: func() (any, error) {
					return service.GetOper(ctx)
				},
			},
			{
				Name: "GetOperCapwapDataAll",
				Method: func() (any, error) {
					return service.GetOperCapwapDataAll(ctx)
				},
			},
			{
				Name: "GetOperRadioStatusAll",
				Method: func() (any, error) {
					return service.GetOperRadioStatusAll(ctx)
				},
			},
		},
		FilterMethods: []framework.GetOperTestMethod{
			{
				Name: "GetOperCapwapDataByMac",
				Method: func() (any, error) {
					return service.GetOperCapwapDataByMac(ctx, testAPMac)
				},
			},
			{
				Name: "GetOperNameMacMapByWtpName",
				Method: func() (any, error) {
					return service.GetOperNameMacMapByWtpName(ctx, testWtpName)
				},
			},
			{
				Name: "GetOperRadioStatusBySlot",
				Method: func() (any, error) {
					return service.GetOperRadioStatusBySlot(ctx, testAPMac, testSlotID)
				},
			},
			{
				Name: "GetOperApRadioNeighborBySlotBssid",
				Method: func() (any, error) {
					return service.GetOperApRadioNeighborBySlotBssid(ctx, testAPMac, testSlotID, testNeighborMAC)
				},
			},
			{
				Name: "GetOperApIoxOperData",
				Method: func() (any, error) {
					return service.GetOperApIoxOperData(ctx, testAPMac)
				},
			},
			{
				Name: "GetOperApImageActiveLocationOnly",
				Method: func() (any, error) {
					return service.GetOperApImageActiveLocationOnly(ctx)
				},
			},
			{
				Name: "GetOperApImagePrepareLocationOnly",
				Method: func() (any, error) {
					return service.GetOperApImagePrepareLocationOnly(ctx)
				},
			},
			{
				Name: "GetOperApPwrInfoOnly",
				Method: func() (any, error) {
					return service.GetOperApPwrInfoOnly(ctx)
				},
			},
			{
				Name: "GetOperApSensorStatusOnly",
				Method: func() (any, error) {
					return service.GetOperApSensorStatusOnly(ctx)
				},
			},
			{
				Name: "GetOperCapwapPktsOnly",
				Method: func() (any, error) {
					return service.GetOperCapwapPktsOnly(ctx)
				},
			},
		},
		ParameterTests: []framework.ErrorFunctionValidationCase{
			// Nil client tests
			{
				Name: "GetOper_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetOper(ctx)
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperCapwapDataAll_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetOperCapwapDataAll(ctx)
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperRadioStatusAll_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetOperRadioStatusAll(ctx)
					return nil, err
				},
				ExpectedError: true,
			},
			// Parameter validation tests
			{
				Name: "GetOperCapwapDataByMac_EmptyMAC",
				Function: func() (any, error) {
					_, err := service.GetOperCapwapDataByMac(ctx, "")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperNameMacMapByWtpName_EmptyName",
				Function: func() (any, error) {
					_, err := service.GetOperNameMacMapByWtpName(ctx, "")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperRadioStatusBySlot_EmptyMAC",
				Function: func() (any, error) {
					_, err := service.GetOperRadioStatusBySlot(ctx, "", testSlotID)
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperRadioStatusBySlot_NegativeSlot",
				Function: func() (any, error) {
					_, err := service.GetOperRadioStatusBySlot(ctx, testAPMac, -1)
					return nil, err
				},
				ExpectedError: true,
			},
		},
		TestConstants: framework.StandardGetOperTestConstants(),
	})
}
