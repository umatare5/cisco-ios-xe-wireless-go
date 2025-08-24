// Package client provides wireless client operational data operations for the Cisco IOS-XE Wireless Network Controller API.
package client

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/framework"
)

// Test constants
const (
	testClientMAC = "aa:bb:cc:dd:ee:ff"
)

func Test_ClientGetOper_StandardTests(t *testing.T) {
	testClient := client.OptionalTestClient(t)
	service := NewService(testClient)
	nilService := NewService(nil)
	ctx := client.TestContext(t)

	// Define test suite using the new standard pattern
	suite := framework.GetOperTestSuite{
		ServiceName: "Client",
		BasicMethods: []framework.GetOperTestMethod{
			{
				Name: "GetOper",
				Method: func() (any, error) {
					return service.GetOper(ctx)
				},
			},
			{
				Name: "GetOperCommonOperData",
				Method: func() (any, error) {
					return service.GetOperCommonOperData(ctx)
				},
			},
			{
				Name: "GetOperDcInfo",
				Method: func() (any, error) {
					return service.GetOperDcInfo(ctx)
				},
			},
			{
				Name: "GetOperDot11OperData",
				Method: func() (any, error) {
					return service.GetOperDot11OperData(ctx)
				},
			},
			{
				Name: "GetOperMmIfClientHistory",
				Method: func() (any, error) {
					return service.GetOperMmIfClientHistory(ctx)
				},
			},
			{
				Name: "GetOperMmIfClientStats",
				Method: func() (any, error) {
					return service.GetOperMmIfClientStats(ctx)
				},
			},
			{
				Name: "GetOperMobilityOperData",
				Method: func() (any, error) {
					return service.GetOperMobilityOperData(ctx)
				},
			},
			{
				Name: "GetOperPolicyData",
				Method: func() (any, error) {
					return service.GetOperPolicyData(ctx)
				},
			},
			{
				Name: "GetOperSisfDBMac",
				Method: func() (any, error) {
					return service.GetOperSisfDBMac(ctx)
				},
			},
			{
				Name: "GetOperTrafficStats",
				Method: func() (any, error) {
					return service.GetOperTrafficStats(ctx)
				},
			},
		},
		FilterMethods: []framework.GetOperTestMethod{
			{
				Name: "GetOperByClientMac",
				Method: func() (any, error) {
					return service.GetOperByClientMac(ctx, testClientMAC)
				},
			},
			{
				Name: "GetOperByApName",
				Method: func() (any, error) {
					return service.GetOperByApName(ctx, "test-ap")
				},
			},
			{
				Name: "GetOperByClientType",
				Method: func() (any, error) {
					return service.GetOperByClientType(ctx, "wireless")
				},
			},
			{
				Name: "GetOperByCoState",
				Method: func() (any, error) {
					return service.GetOperByCoState(ctx, "associated")
				},
			},
			{
				Name: "GetOperByMsRadioType",
				Method: func() (any, error) {
					return service.GetOperByMsRadioType(ctx, "802.11ac")
				},
			},
			{
				Name: "GetOperByUsername",
				Method: func() (any, error) {
					return service.GetOperByUsername(ctx, "testuser")
				},
			},
			{
				Name: "GetOperByWlanID",
				Method: func() (any, error) {
					return service.GetOperByWlanID(ctx, 1)
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
				Name: "GetOperCommonOperData_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetOperCommonOperData(ctx)
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperDcInfo_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetOperDcInfo(ctx)
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperDot11OperData_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetOperDot11OperData(ctx)
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperMmIfClientHistory_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetOperMmIfClientHistory(ctx)
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperMmIfClientStats_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetOperMmIfClientStats(ctx)
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperMobilityOperData_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetOperMobilityOperData(ctx)
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperPolicyData_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetOperPolicyData(ctx)
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperSisfDBMac_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetOperSisfDBMac(ctx)
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperTrafficStats_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetOperTrafficStats(ctx)
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperByClientMac_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetOperByClientMac(ctx, testClientMAC)
					return nil, err
				},
				ExpectedError: true,
			},
			// Parameter validation tests
			{
				Name: "GetOperByClientMac_EmptyMAC",
				Function: func() (any, error) {
					_, err := service.GetOperByClientMac(ctx, "")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperByClientMac_WhitespaceMAC",
				Function: func() (any, error) {
					_, err := service.GetOperByClientMac(ctx, "   ")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperByApName_EmptyParam",
				Function: func() (any, error) {
					_, err := service.GetOperByApName(ctx, "")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperByClientType_EmptyParam",
				Function: func() (any, error) {
					_, err := service.GetOperByClientType(ctx, "")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperByCoState_EmptyParam",
				Function: func() (any, error) {
					_, err := service.GetOperByCoState(ctx, "")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperByMsRadioType_EmptyParam",
				Function: func() (any, error) {
					_, err := service.GetOperByMsRadioType(ctx, "")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperByUsername_EmptyParam",
				Function: func() (any, error) {
					_, err := service.GetOperByUsername(ctx, "")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperByWlanID_NegativeParam",
				Function: func() (any, error) {
					_, err := service.GetOperByWlanID(ctx, -1)
					return nil, err
				},
				ExpectedError: true,
			},
		},
		TestConstants: framework.StandardGetOperTestConstants(),
	}

	// Run the standardized test suite
	framework.RunStandardGetOperTests(t, suite)
}
