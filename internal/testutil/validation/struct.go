// Package validation provides test validation functions
package validation

import (
	"encoding/json"
	"reflect"
	"testing"
)

// ValidateJSONSerialization tests JSON marshal/unmarshal round-trip
func ValidateJSONSerialization(t *testing.T, original any) {
	t.Helper()

	// Marshal to JSON
	jsonData, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal to JSON: %v", err)
	}

	// Unmarshal back to same type
	originalType := reflect.TypeOf(original)
	newInstance := reflect.New(originalType.Elem()).Interface()

	err = json.Unmarshal(jsonData, newInstance)
	if err != nil {
		t.Fatalf("Failed to unmarshal from JSON: %v", err)
	}

	// Compare original and unmarshaled
	if !reflect.DeepEqual(original, newInstance) {
		t.Errorf("JSON round-trip failed: original != unmarshaled")
	}
}

// ValidateJSONUnmarshal tests unmarshaling from JSON string to struct
func ValidateJSONUnmarshal(t *testing.T, jsonStr string, target any) {
	t.Helper()

	err := json.Unmarshal([]byte(jsonStr), target)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// Validate that target is not nil after unmarshaling
	targetValue := reflect.ValueOf(target)
	if targetValue.Kind() == reflect.Ptr && targetValue.IsNil() {
		t.Error("Target is nil after unmarshaling")
	}
}

// ValidateResponseStructure validates response structure matches expected YANG model
func ValidateResponseStructure(t *testing.T, response any, structName string) {
	t.Helper()

	if response == nil {
		t.Fatalf("Response cannot be nil for %s", structName)
	}

	responseType := reflect.TypeOf(response)

	// Handle pointer types
	if responseType.Kind() == reflect.Ptr {
		responseType = responseType.Elem()
	}

	if responseType.Name() != structName {
		t.Errorf("Expected struct name %s, got %s", structName, responseType.Name())
	}

	// Validate that it's a proper struct
	if responseType.Kind() != reflect.Struct {
		t.Fatalf("Expected struct, got %v", responseType.Kind())
	}
}
