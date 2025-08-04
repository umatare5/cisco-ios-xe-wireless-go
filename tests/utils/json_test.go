package testutils

import (
	"errors"
	"testing"
)

func TestTestJSONUnmarshal(t *testing.T) {
	type TestStruct struct {
		Name string `json:"name"`
		ID   int    `json:"id"`
	}

	jsonData := `{"name":"test","id":123}`
	var result TestStruct

	TestJSONUnmarshal(t, jsonData, &result, "TestStruct")

	if result.Name != "test" {
		t.Errorf("Expected name 'test', got %s", result.Name)
	}

	if result.ID != 123 {
		t.Errorf("Expected ID 123, got %d", result.ID)
	}
}

func TestTestJSONUnmarshalError(t *testing.T) {
	type TestStruct struct {
		Name string `json:"name"`
		ID   int    `json:"id"`
	}

	invalidJSON := `{"name":"test","id":"invalid_number"}`
	var result TestStruct

	TestJSONUnmarshalError(t, invalidJSON, &result, "TestStruct")
}

func TestRunJSONTestsSuccess(t *testing.T) {
	type TestData struct {
		Name string `json:"name"`
		ID   int    `json:"id"`
	}

	tests := []JSONTestCase{
		{
			Name:       "valid JSON",
			JSONData:   `{"name":"test","id":123}`,
			Target:     &TestData{},
			TypeName:   "TestData",
			ShouldFail: false,
		},
		{
			Name:       "empty object",
			JSONData:   `{}`,
			Target:     &TestData{},
			TypeName:   "TestData",
			ShouldFail: false,
		},
	}

	RunJSONTests(t, tests)
}

func TestRunJSONTestsFailure(t *testing.T) {
	type TestData struct {
		Name string `json:"name"`
		ID   int    `json:"id"`
	}

	tests := []JSONTestCase{
		{
			Name:       "invalid JSON",
			JSONData:   `{"name":"test","id":"invalid"}`,
			Target:     &TestData{},
			TypeName:   "TestData",
			ShouldFail: true,
		},
	}

	RunJSONTests(t, tests)
}

func TestValidateJSONStructFields(t *testing.T) {
	type TestStruct struct {
		Name string `json:"name"`
		ID   int    `json:"id"`
	}

	jsonData := `{"name":"test","id":123}`
	var result TestStruct
	TestJSONUnmarshal(t, jsonData, &result, "TestStruct")

	ValidateJSONStructFields(t, "TestStruct", func() error {
		if result.Name == "" {
			return errors.New("name field is empty")
		}
		if result.ID == 0 {
			return errors.New("id field is zero")
		}
		return nil
	})
}

func TestJSONUnmarshalWithComplexStructure(t *testing.T) {
	type NestedStruct struct {
		Inner struct {
			Value string `json:"value"`
		} `json:"inner"`
		Items []string `json:"items"`
	}

	jsonData := `{"inner":{"value":"nested"},"items":["a","b","c"]}`
	var result NestedStruct

	TestJSONUnmarshal(t, jsonData, &result, "NestedStruct")

	if result.Inner.Value != "nested" {
		t.Errorf("Expected inner value 'nested', got %s", result.Inner.Value)
	}

	if len(result.Items) != 3 {
		t.Errorf("Expected 3 items, got %d", len(result.Items))
	}
}

func TestJSONTestCaseStructure(t *testing.T) {
	testCase := JSONTestCase{
		Name:     "simple test",
		JSONData: `{"value":"hello"}`,
		Target: &struct {
			Value string `json:"value"`
		}{},
		TypeName:   "SimpleStruct",
		ShouldFail: false,
	}

	if testCase.Name != "simple test" {
		t.Errorf("Expected name 'simple test', got %s", testCase.Name)
	}

	if testCase.ShouldFail != false {
		t.Errorf("Expected ShouldFail false, got %t", testCase.ShouldFail)
	}
}
