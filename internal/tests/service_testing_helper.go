package tests

// ServiceTestConfig holds configuration for service testing
type ServiceTestConfig struct {
	ServiceName    string
	TestMethods    []TestMethod
	JSONTestCases  []JSONTestCase
	SkipShortTests bool
}

// TestMethod represents a service method to test
type TestMethod struct {
	Name   string
	Method func() (interface{}, error)
}

// JSONTestCase represents a JSON serialization test case
type JSONTestCase struct {
	Name     string
	JSONData string
}
