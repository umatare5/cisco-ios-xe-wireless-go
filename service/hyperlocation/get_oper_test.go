package hyperlocation

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/framework"
)

// Test constants
const (
	testProfileName = "test-profile"
)

func Test_HyperlocationGetOper_StandardTests(t *testing.T) {
	clientSetup := client.OptionalTestClient(t)
	service := NewService(clientSetup)
	nilService := NewService(nil)
	ctx := client.TestContext(t)

	framework.RunStandardGetOperTests(t, framework.GetOperTestSuite{
		ServiceName: "Hyperlocation",
		BasicMethods: []framework.GetOperTestMethod{
			{
				Name: "GetOper",
				Method: func() (any, error) {
					return service.GetOper(ctx)
				},
			},
			{
				Name: "GetOperProfiles",
				Method: func() (any, error) {
					return service.GetOperProfiles(ctx)
				},
			},
		},
		FilterMethods: []framework.GetOperTestMethod{
			{
				Name: "GetOperByName",
				Method: func() (any, error) {
					return service.GetOperByName(ctx, testProfileName)
				},
			},
		},
		ParameterTests: []framework.ErrorFunctionValidationCase{
			{
				Name: "GetOper_NilClient",
				Function: func() (any, error) {
					return nilService.GetOper(ctx)
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperProfiles_NilClient",
				Function: func() (any, error) {
					return nilService.GetOperProfiles(ctx)
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperByName_EmptyName",
				Function: func() (any, error) {
					return service.GetOperByName(ctx, "")
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperByName_NilClient",
				Function: func() (any, error) {
					return nilService.GetOperByName(ctx, testProfileName)
				},
				ExpectedError: true,
			},
		},
		TestConstants: framework.StandardGetOperTestConstants(),
	})
}
