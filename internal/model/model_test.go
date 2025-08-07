package model

import (
	"encoding/json"
	"testing"
)

// TestEmptyStructsJSONMarshal tests that empty model structs can be marshaled without errors
func TestEmptyStructsJSONMarshal(t *testing.T) {
	emptyStructs := []struct {
		name string
		data interface{}
	}{
		{"AfcOperResponse", &AfcOperResponse{}},
		{"AfcCloudOperResponse", &AfcCloudOperResponse{}},
		{"ApCfgResponse", &ApCfgResponse{}},
	}

	for _, tt := range emptyStructs {
		t.Run(tt.name, func(t *testing.T) {
			// Marshal empty struct
			jsonData, err := json.Marshal(tt.data)
			if err != nil {
				t.Fatalf("Failed to marshal %s: %v", tt.name, err)
			}

			// Should produce valid JSON
			if !json.Valid(jsonData) {
				t.Errorf("Invalid JSON for %s: %s", tt.name, string(jsonData))
			}

			// Basic validation that JSON is not empty
			if len(jsonData) < 2 {
				t.Errorf("JSON output too short for %s: %s", tt.name, string(jsonData))
			}
		})
	}
}

// TestJSONUnmarshalRoundTrip tests that marshaled JSON can be unmarshaled back
func TestJSONUnmarshalRoundTrip(t *testing.T) {
	testCases := []struct {
		name string
		data interface{}
	}{
		{"AfcOperResponse", &AfcOperResponse{}},
		{"AfcCloudOperResponse", &AfcCloudOperResponse{}},
		{"ApCfgResponse", &ApCfgResponse{}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			// Marshal to JSON
			jsonData, err := json.Marshal(tt.data)
			if err != nil {
				t.Fatalf("Failed to marshal %s: %v", tt.name, err)
			}

			// Unmarshal back
			switch tt.name {
			case "AfcOperResponse":
				var unmarshaled AfcOperResponse
				if err := json.Unmarshal(jsonData, &unmarshaled); err != nil {
					t.Fatalf("Failed to unmarshal %s: %v", tt.name, err)
				}
			case "AfcCloudOperResponse":
				var unmarshaled AfcCloudOperResponse
				if err := json.Unmarshal(jsonData, &unmarshaled); err != nil {
					t.Fatalf("Failed to unmarshal %s: %v", tt.name, err)
				}
			case "ApCfgResponse":
				var unmarshaled ApCfgResponse
				if err := json.Unmarshal(jsonData, &unmarshaled); err != nil {
					t.Fatalf("Failed to unmarshal %s: %v", tt.name, err)
				}
			}
		})
	}
}

// TestBasicTypeValidation ensures our model package defines expected types
func TestBasicTypeValidation(t *testing.T) {
	// This test ensures that critical types are defined
	var (
		_ = &AfcOperResponse{}
		_ = &AfcCloudOperResponse{}
		_ = &ApCfgResponse{}
	)

	// If compilation succeeds, the types exist
	t.Log("All expected model types are defined")
}
