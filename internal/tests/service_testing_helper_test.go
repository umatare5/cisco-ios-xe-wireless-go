package tests

import "testing"

func TestServiceTestingTypes(t *testing.T) {
	cfg := ServiceTestConfig{
		ServiceName:    "X",
		TestMethods:    []TestMethod{{Name: "M", Method: func() (interface{}, error) { return nil, nil }}},
		JSONTestCases:  []JSONTestCase{{Name: "J", JSONData: `{"a":1}`}},
		SkipShortTests: true,
	}
	if cfg.ServiceName == "" || len(cfg.TestMethods) == 0 || len(cfg.JSONTestCases) == 0 {
		t.Fatal("invalid cfg")
	}
}
