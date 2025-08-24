// Package validation provides test validation functions
package validation

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"
)

// UnmarshalAndValidateJSON unmarshals JSON string into provided structure and validates
func UnmarshalAndValidateJSON(jsonStr string, target any) error {
	if jsonStr == "" {
		return errors.New("JSON string cannot be empty")
	}

	err := json.Unmarshal([]byte(jsonStr), target)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return nil
}

// MarshalToJSON marshals provided structure to JSON string
func MarshalToJSON(source any) ([]byte, error) {
	if source == nil {
		return nil, errors.New("source cannot be nil")
	}

	data, err := json.Marshal(source)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal to JSON: %w", err)
	}

	return data, nil
}

// RunParallelJSONSerializationTests executes JSON serialization tests for all provided test cases
func RunParallelJSONSerializationTests(t *testing.T, testCases map[string]func() any) {
	t.Helper()

	for name, createFunc := range testCases {
		createFunc := createFunc // capture loop variable
		name := name             // capture loop variable
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Create original object
			original := createFunc()

			// Marshal to JSON
			data, err := json.Marshal(original)
			if err != nil {
				t.Fatalf("Failed to marshal %s: %v", name, err)
			}

			// Create new object and unmarshal
			result := createFunc()
			if err := json.Unmarshal(data, result); err != nil {
				t.Fatalf("Failed to unmarshal %s: %v", name, err)
			}

			// Basic validation - ensure it's not nil
			if result == nil {
				t.Fatalf("%s unmarshal resulted in nil", name)
			}
		})
	}
}
