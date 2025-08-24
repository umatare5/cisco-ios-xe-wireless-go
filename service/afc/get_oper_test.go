package afc

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/framework"
)

// Test constants
const (
	testAPMac     = "a4:8c:db:ac:ba:20"
	testRequestID = "test-request-id"
	testSlotID    = 0
)

func Test_AfcGetOper_StandardTests(t *testing.T) {
	testClient := client.OptionalTestClient(t)
	service := NewService(testClient)
	nilService := NewService(nil)
	ctx := client.TestContext(t)

	// Define test suite using the new standard pattern
	suite := framework.GetOperTestSuite{
		ServiceName: "AFC",
		BasicMethods: []framework.GetOperTestMethod{
			{
				Name: "GetOper",
				Method: func() (any, error) {
					return service.GetOper(ctx)
				},
			},
			{
				Name: "GetOperAPResp",
				Method: func() (any, error) {
					return service.GetOperAPResp(ctx)
				},
			},
			{
				Name: "GetOperCloudOper",
				Method: func() (any, error) {
					return service.GetOperCloudOper(ctx)
				},
			},
			{
				Name: "GetOperCloudStats",
				Method: func() (any, error) {
					return service.GetOperCloudStats(ctx)
				},
			},
		},
		FilterMethods: []framework.GetOperTestMethod{
			{
				Name: "GetOperByApMac",
				Method: func() (any, error) {
					return service.GetOperByApMac(ctx, testAPMac)
				},
			},
			{
				Name: "GetOperByApMacAndSlot",
				Method: func() (any, error) {
					return service.GetOperByApMacAndSlot(ctx, testAPMac, testSlotID)
				},
			},
			{
				Name: "GetOperByApMacAndRequestID",
				Method: func() (any, error) {
					return service.GetOperByApMacAndRequestID(ctx, testAPMac, testRequestID)
				},
			},
			{
				Name: "GetOperBySlot",
				Method: func() (any, error) {
					return service.GetOperBySlot(ctx, testAPMac, testSlotID)
				},
			},
			{
				Name: "GetOperByRequestID",
				Method: func() (any, error) {
					return service.GetOperByRequestID(ctx, testAPMac, testRequestID)
				},
			},
		},
		ParameterTests: []framework.ErrorFunctionValidationCase{
			// Nil client tests embedded in the framework
			{
				Name: "GetOper_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetOper(ctx)
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperAPResp_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetOperAPResp(ctx)
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperCloudOper_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetOperCloudOper(ctx)
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperCloudStats_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetOperCloudStats(ctx)
					return nil, err
				},
				ExpectedError: true,
			},
			// Parameter validation tests
			{
				Name: "GetOperByApMac_Empty",
				Function: func() (any, error) {
					_, err := service.GetOperByApMac(ctx, "")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperByApMac_Whitespace",
				Function: func() (any, error) {
					_, err := service.GetOperByApMac(ctx, "   ")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperByApMacAndSlot_EmptyMAC",
				Function: func() (any, error) {
					_, err := service.GetOperByApMacAndSlot(ctx, "", testSlotID)
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperByApMacAndSlot_NegativeSlot",
				Function: func() (any, error) {
					_, err := service.GetOperByApMacAndSlot(ctx, testAPMac, -1)
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperByApMacAndRequestID_EmptyMAC",
				Function: func() (any, error) {
					_, err := service.GetOperByApMacAndRequestID(ctx, "", testRequestID)
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperByApMacAndRequestID_EmptyRequestID",
				Function: func() (any, error) {
					_, err := service.GetOperByApMacAndRequestID(ctx, testAPMac, "")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperBySlot_EmptyMAC",
				Function: func() (any, error) {
					_, err := service.GetOperBySlot(ctx, "", testSlotID)
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperBySlot_NegativeSlot",
				Function: func() (any, error) {
					_, err := service.GetOperBySlot(ctx, testAPMac, -1)
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperByRequestID_EmptyMAC",
				Function: func() (any, error) {
					_, err := service.GetOperByRequestID(ctx, "", testRequestID)
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperByRequestID_EmptyRequestID",
				Function: func() (any, error) {
					_, err := service.GetOperByRequestID(ctx, testAPMac, "")
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
