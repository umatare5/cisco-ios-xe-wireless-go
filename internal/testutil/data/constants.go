package data

import (
	"fmt"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/helpers"
)

// JSONTestCase represents a JSON serialization test case
type JSONTestCase struct {
	Name     string
	JSONData string
}

// JSONSerializationTestCase defines a test case for JSON serialization/deserialization
type JSONSerializationTestCase struct {
	Name       string
	CreateFunc func() any
}

// StandardJSONTestCases provides standard JSON test cases for most services
func StandardJSONTestCases(yangModule string) []JSONTestCase {
	return []JSONTestCase{
		{
			Name: helpers.PascalCase(yangModule) + "CfgResponse",
			JSONData: fmt.Sprintf(`{
				"%s%s-cfg:%s-cfg-data": {
					"test-data": "value"
				}
			}`, restconf.YANGModelPrefix, yangModule, yangModule),
		},
		{
			Name: helpers.PascalCase(yangModule) + "OperResponse",
			JSONData: fmt.Sprintf(`{
				"%s%s-oper:%s-oper-data": {
					"test-data": "value"
				}
			}`, restconf.YANGModelPrefix, yangModule, yangModule),
		},
	}
}
