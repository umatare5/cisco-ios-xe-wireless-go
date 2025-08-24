package mcast

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/framework"
)

func Test_McastGetOper_StandardTests(t *testing.T) {
	testClient := client.OptionalTestClient(t)
	service := NewService(testClient)
	nilService := NewService(nil)
	ctx := client.TestContext(t)

	framework.RunStandardGetOperTests(t, framework.GetOperTestSuite{
		ServiceName: "Mcast",
		BasicMethods: []framework.GetOperTestMethod{
			{
				Name: "GetOper",
				Method: func() (any, error) {
					return service.GetOper(ctx)
				},
			},
			{
				Name: "GetOperFlexMediastream",
				Method: func() (any, error) {
					return service.GetOperFlexMediastream(ctx)
				},
			},
			{
				Name: "GetOperVlanL2Mgid",
				Method: func() (any, error) {
					return service.GetOperVlanL2Mgid(ctx)
				},
			},
		},
		FilterMethods: []framework.GetOperTestMethod{
			{
				Name: "GetOperByClientMAC",
				Method: func() (any, error) {
					return service.GetOperByClientMAC(ctx, "aa:bb:cc:dd:ee:ff")
				},
			},
			{
				Name: "GetOperByVlanIndex",
				Method: func() (any, error) {
					return service.GetOperByVlanIndex(ctx, "100")
				},
			},
		},
		ParameterTests: []framework.ErrorFunctionValidationCase{
			{
				Name: "GetOper_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetOper(ctx)
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperFlexMediastream_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetOperFlexMediastream(ctx)
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperVlanL2Mgid_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetOperVlanL2Mgid(ctx)
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperByClientMAC_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetOperByClientMAC(ctx, "aa:bb:cc:dd:ee:ff")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperByClientMAC_EmptyParam",
				Function: func() (any, error) {
					_, err := service.GetOperByClientMAC(ctx, "")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperByClientMAC_InvalidMAC",
				Function: func() (any, error) {
					_, err := service.GetOperByClientMAC(ctx, "invalid-mac")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperByClientMAC_WhitespaceParam",
				Function: func() (any, error) {
					_, err := service.GetOperByClientMAC(ctx, "   ")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperByVlanIndex_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetOperByVlanIndex(ctx, "100")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperByVlanIndex_EmptyParam",
				Function: func() (any, error) {
					_, err := service.GetOperByVlanIndex(ctx, "")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperByVlanIndex_WhitespaceParam",
				Function: func() (any, error) {
					_, err := service.GetOperByVlanIndex(ctx, "   ")
					return nil, err
				},
				ExpectedError: true,
			},
		},
		TestConstants: framework.StandardGetOperTestConstants(),
	})
}

// Test_McastOperOps_Unit tests the operOps method to ensure coverage
func Test_McastOperOps_Unit(t *testing.T) {
	testClient := client.OptionalTestClient(t)
	service := NewService(testClient)

	// Test operOps method
	ops := service.operOps()
	if ops == nil {
		t.Error("operOps() returned nil")
	}
}
