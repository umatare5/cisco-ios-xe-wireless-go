// Package testutils provides common JSON testing utilities for tests.
package testutils

import (
	"encoding/json"
	"testing"
)

// TestJSONUnmarshal provides a standardized way to test JSON unmarshaling.
func TestJSONUnmarshal(t *testing.T, jsonData string, target interface{}, expectedTypeName string) {
	t.Helper()

	err := json.Unmarshal([]byte(jsonData), target)
	if err != nil {
		t.Fatalf("Failed to unmarshal %s: %v", expectedTypeName, err)
	}

	// Verify that target is not nil after unmarshaling
	if target == nil {
		t.Errorf("Unmarshaled %s should not be nil", expectedTypeName)
	}
}

// TestJSONUnmarshalError tests that JSON unmarshaling fails as expected.
func TestJSONUnmarshalError(t *testing.T, invalidJSON string, target interface{}, expectedTypeName string) {
	t.Helper()

	err := json.Unmarshal([]byte(invalidJSON), target)
	if err == nil {
		t.Errorf("Expected error when unmarshaling invalid JSON to %s, but got nil", expectedTypeName)
	}
}

// JSONTestCase represents a single JSON test case.
type JSONTestCase struct {
	Name       string
	JSONData   string
	Target     interface{}
	TypeName   string
	ShouldFail bool
}

// RunJSONTests executes a table of JSON test cases.
func RunJSONTests(t *testing.T, testCases []JSONTestCase) {
	t.Helper()

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Helper()

			if tc.ShouldFail {
				TestJSONUnmarshalError(t, tc.JSONData, tc.Target, tc.TypeName)
			} else {
				TestJSONUnmarshal(t, tc.JSONData, tc.Target, tc.TypeName)
			}
		})
	}
}

// ValidateJSONStructFields validates that all expected fields are present after unmarshaling.
func ValidateJSONStructFields(t *testing.T, structName string, validationFunc func() error) {
	t.Helper()

	if err := validationFunc(); err != nil {
		t.Errorf("Field validation failed for %s: %v", structName, err)
	}
}
