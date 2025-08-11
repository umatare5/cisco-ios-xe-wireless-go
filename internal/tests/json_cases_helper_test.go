package tests

import (
	"encoding/json"
	"strings"
	"testing"
)

// TestStandardJSONTestCases tests the standard JSON test cases
func TestStandardJSONTestCases(t *testing.T) {
	testCases := StandardJSONTestCases("test-module")
	if len(testCases) == 0 {
		t.Error("expected non-empty slice")
	}
	if len(testCases) != 2 {
		t.Errorf("expected 2 test cases, got %d", len(testCases))
	}
	expectedNames := []string{"Test-moduleCfgResponse", "Test-moduleOperResponse"}
	actualNames := []string{testCases[0].Name, testCases[1].Name}
	for _, exp := range expectedNames {
		found := false
		for _, act := range actualNames {
			if act == exp {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("missing %s in %v", exp, actualNames)
		}
	}
	for _, tc := range testCases {
		if tc.JSONData == "" {
			t.Errorf("%s has empty JSON", tc.Name)
		}
		var obj interface{}
		if err := json.Unmarshal([]byte(tc.JSONData), &obj); err != nil {
			t.Errorf("invalid JSON: %v", err)
		}
	}
	for _, module := range []string{"ap", "wlan", "site", "dot11", "dot15"} {
		m := StandardJSONTestCases(module)
		if len(m) != 2 {
			t.Errorf("module %s should return 2 test cases, got %d", module, len(m))
		}
		for _, tc := range m {
			if !strings.Contains(tc.Name, PascalCase(module)) {
				t.Errorf("name %s should contain %s", tc.Name, module)
			}
		}
		for _, tc := range m {
			if !strings.Contains(tc.JSONData, module) {
				t.Errorf("json should contain %s", module)
			}
		}
	}
	if len(StandardJSONTestCases("")) != 2 {
		t.Errorf("empty module should return 2")
	}
	if len(StandardJSONTestCases("a")) != 2 {
		t.Errorf("single char module should return 2")
	}
}

// TestJSONOperations tests basic JSON operations like marshaling and unmarshaling
func TestJSONOperations(t *testing.T) {
	data := map[string]interface{}{
		"string":  "value",
		"number":  123.45,
		"boolean": true,
		"array":   []interface{}{1, 2, 3},
		"object":  map[string]interface{}{"nested": "value"},
	}
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	s := string(b)
	if !strings.Contains(s, "value") {
		t.Error("missing 'value'")
	}
	var u map[string]interface{}
	if err := json.Unmarshal(b, &u); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if u["string"] != data["string"] {
		t.Error("mismatch")
	}
}
