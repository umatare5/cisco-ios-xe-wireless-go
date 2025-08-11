package tests

import (
	"fmt"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/constants"
)

// StandardJSONTestCases provides standard JSON test cases for most services
func StandardJSONTestCases(yangModule string) []JSONTestCase {
	return []JSONTestCase{
		{
			Name: PascalCase(yangModule) + "CfgResponse",
			JSONData: fmt.Sprintf(`{
				"%s%s-cfg:%s-cfg-data": {
					"test-data": "value"
				}
			}`, constants.YANGModelPrefix, yangModule, yangModule),
		},
		{
			Name: PascalCase(yangModule) + "OperResponse",
			JSONData: fmt.Sprintf(`{
				"%s%s-oper:%s-oper-data": {
					"test-data": "value"
				}
			}`, constants.YANGModelPrefix, yangModule, yangModule),
		},
	}
}
