package rf

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/framework"
)

func TestService(t *testing.T) {
	testClient := client.OptionalTestClient(t)
	service := NewService(testClient)
	nilService := NewService(nil)
	ctx := client.TestContext(t)

	framework.RunStandardGetOperTests(t, framework.GetOperTestSuite{
		ServiceName: "RF Service",
		BasicMethods: []framework.GetOperTestMethod{
			{
				Name: "GetCfg",
				Method: func() (any, error) {
					return service.GetCfg(ctx)
				},
			},
		},
		ParameterTests: []framework.ErrorFunctionValidationCase{
			{
				Name: "GetCfg with nil client",
				Function: func() (any, error) {
					_, err := nilService.GetCfg(ctx)
					return nil, err
				},
				ExpectedError: true,
			},
		},
		CustomTests: func(t *testing.T) {
			// RF Tag Service operations are tested in separate tag_*_test.go files
			t.Run("RFTagServiceSeparateTests", func(t *testing.T) {
				t.Log("RF Tag Service CRUD operations are tested in dedicated " +
					"tag_create_test.go, tag_set_test.go, tag_get_test.go, " +
					"tag_list_test.go, and tag_delete_test.go files")
			})
		},
	})
}
