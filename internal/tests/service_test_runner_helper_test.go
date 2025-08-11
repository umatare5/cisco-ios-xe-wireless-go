package tests

import (
	"fmt"
	"os"
	"testing"
)

func TestRunServiceTests_SuiteA(t *testing.T) { // renamed to avoid redeclare with split
	methods := []TestMethod{
		{Name: "MockMethod1", Method: func() (interface{}, error) {
			return "mock response 1", nil
		}},
		{Name: "MockMethod2", Method: func() (interface{}, error) {
			return "mock response 2", fmt.Errorf("mock error")
		}},
		{Name: "MockMethod3", Method: func() (interface{}, error) {
			return nil, nil
		}},
	}
	jsonCases := []JSONTestCase{
		{Name: "MockJSONTest1", JSONData: `{"test":"data"}`},
		{Name: "MockJSONTest2", JSONData: `{"nested":{"key":"value"},"array":[1,2,3]}`},
	}
	cfg := ServiceTestConfig{
		ServiceName:    "MockService",
		TestMethods:    methods,
		JSONTestCases:  jsonCases,
		SkipShortTests: true,
	}
	RunServiceTests(t, cfg)
	RunServiceTests(t, ServiceTestConfig{ServiceName: "EmptyService"})
	RunServiceTests(t, ServiceTestConfig{
		ServiceName:    "ShortTestService",
		TestMethods:    methods[:1],
		JSONTestCases:  jsonCases[:1],
		SkipShortTests: false,
	})
}

func TestRunServiceTests_IntegrationBranch(t *testing.T) {
	origShort := shortModeCheck
	shortModeCheck = func() bool { return false }
	defer func() { shortModeCheck = origShort }()
	origC, origT := os.Getenv("WNC_CONTROLLER"), os.Getenv("WNC_ACCESS_TOKEN")
	os.Setenv("WNC_CONTROLLER", "dummy-controller")
	os.Setenv("WNC_ACCESS_TOKEN", "dummy-token")
	defer func() {
		if origC == "" {
			os.Unsetenv("WNC_CONTROLLER")
		} else {
			os.Setenv("WNC_CONTROLLER", origC)
		}
		if origT == "" {
			os.Unsetenv("WNC_ACCESS_TOKEN")
		} else {
			os.Setenv("WNC_ACCESS_TOKEN", origT)
		}
	}()
	cfg := ServiceTestConfig{
		ServiceName:    "IntegrationService",
		SkipShortTests: true,
		TestMethods: []TestMethod{{
			Name: "Dummy",
			Method: func() (interface{}, error) {
				return struct{ X int }{1}, nil
			},
		}},
	}
	RunServiceTests(t, cfg)
}

func TestRunServiceTests_ShortModeSkip(t *testing.T) {
	orig := shortModeCheck
	shortModeCheck = func() bool { return true }
	defer func() { shortModeCheck = orig }()
	RunServiceTests(t, ServiceTestConfig{ServiceName: "dummy", SkipShortTests: true})
}

func TestRunServiceTests_NoMethods(t *testing.T) {
	RunServiceTests(t, ServiceTestConfig{ServiceName: "dummy", TestMethods: nil})
}

func TestRunServiceTests_SkipBranches(t *testing.T) {
	orig := shortModeCheck
	shortModeCheck = func() bool { return true }
	defer func() { shortModeCheck = orig }()
	RunServiceTests(t, ServiceTestConfig{ServiceName: "SkipService", SkipShortTests: true})
}

func TestRunServiceTests_NilClientSkip(t *testing.T) {
	origC, origT := os.Getenv("WNC_CONTROLLER"), os.Getenv("WNC_ACCESS_TOKEN")
	os.Unsetenv("WNC_CONTROLLER")
	os.Unsetenv("WNC_ACCESS_TOKEN")
	defer func() {
		if origC != "" {
			os.Setenv("WNC_CONTROLLER", origC)
		}
		if origT != "" {
			os.Setenv("WNC_ACCESS_TOKEN", origT)
		}
	}()
	RunServiceTests(t, ServiceTestConfig{ServiceName: "NilClientService", SkipShortTests: false})
}

func TestRunServiceTests_NoClient(t *testing.T) {
	RunServiceTests(t, ServiceTestConfig{ServiceName: "dummy", SkipShortTests: true})
}

func TestRunServiceTests_WithMethods(t *testing.T) {
	called := false
	RunServiceTests(t, ServiceTestConfig{
		ServiceName: "dummy",
		TestMethods: []TestMethod{{
			Name: "M1",
			Method: func() (interface{}, error) {
				called = true
				return struct{ X int }{1}, nil
			},
		}},
	})
	if !called {
		t.Error("expected method to be called")
	}
}

func TestRunServiceTests_JSONCases(t *testing.T) {
	RunServiceTests(t, ServiceTestConfig{
		ServiceName: "dummy",
		JSONTestCases: []JSONTestCase{{
			Name:     "Simple",
			JSONData: `{"a":1}`,
		}},
	})
}

func TestRunServiceTests_IntegrationShortMode(t *testing.T) {
	orig := shortModeCheck
	shortModeCheck = func() bool { return true }
	defer func() { shortModeCheck = orig }()
	RunServiceTests(t, ServiceTestConfig{
		ServiceName: "dummy",
		TestMethods: []TestMethod{{
			Name:   "M",
			Method: func() (interface{}, error) { return nil, nil },
		}},
		SkipShortTests: true,
	})
}

func TestRunServiceTests_IntegrationNoClient(t *testing.T) {
	orig := shortModeCheck
	shortModeCheck = func() bool { return false }
	defer func() { shortModeCheck = orig }()
	RunServiceTests(t, ServiceTestConfig{
		ServiceName: "dummy",
		TestMethods: []TestMethod{{
			Name:   "M",
			Method: func() (interface{}, error) { return nil, nil },
		}},
	})
}

func TestRunServiceTests_IntegrationExec(t *testing.T) {
	orig := shortModeCheck
	shortModeCheck = func() bool { return false }
	defer func() { shortModeCheck = orig }()
	RunServiceTests(t, ServiceTestConfig{
		ServiceName: "dummy",
		TestMethods: []TestMethod{{
			Name:   "M",
			Method: func() (interface{}, error) { return struct{ Y string }{"val"}, nil },
		}},
	})
}
