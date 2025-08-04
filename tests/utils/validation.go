// Package testutils provides common validation utilities for tests.
package testutils

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
)

// ValidateStructFields validates that a struct has expected fields.
func ValidateStructFields(t *testing.T, obj interface{}, expectedFields []string) {
	t.Helper()

	objType := reflect.TypeOf(obj)
	if objType.Kind() == reflect.Ptr {
		objType = objType.Elem()
	}

	for _, field := range expectedFields {
		_, found := objType.FieldByName(field)
		if !found {
			t.Errorf("Expected field '%s' not found in struct", field)
		}
	}
}

// ValidateJSONTags validates that struct fields have proper JSON tags.
func ValidateJSONTags(t *testing.T, obj interface{}) {
	t.Helper()

	objType := reflect.TypeOf(obj)
	if objType.Kind() == reflect.Ptr {
		objType = objType.Elem()
	}

	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)
		jsonTag := field.Tag.Get("json")

		if jsonTag == "" && field.IsExported() {
			t.Errorf("Field '%s' is exported but missing JSON tag", field.Name)
		}
	}
}

// ValidateStringNotEmpty validates that a string is not empty.
func ValidateStringNotEmpty(t *testing.T, value, fieldName string) {
	t.Helper()

	if strings.TrimSpace(value) == "" {
		t.Errorf("Field '%s' should not be empty", fieldName)
	}
}

// ValidateErrorContains validates that an error contains expected substring.
func ValidateErrorContains(t *testing.T, err error, expected string) {
	t.Helper()

	if err == nil {
		t.Errorf("Expected error containing '%s', but got nil", expected)
		return
	}

	if !strings.Contains(err.Error(), expected) {
		t.Errorf("Expected error to contain '%s', got: %v", expected, err)
	}
}

// ValidateNoError validates that no error occurred.
func ValidateNoError(t *testing.T, err error, operation string) {
	t.Helper()

	if err != nil {
		t.Errorf("Unexpected error in %s: %v", operation, err)
	}
}

// RunTableTests executes table-driven tests with consistent structure.
func RunTableTests(t *testing.T, tests []TableTest) {
	t.Helper()

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			t.Helper()
			tt.Test(t)
		})
	}
}

// TableTest represents a single table-driven test case.
type TableTest struct {
	Name string
	Test func(t *testing.T)
}

// EndpointValidationTest provides common endpoint validation testing.
func EndpointValidationTest(t *testing.T, endpoint, expectedPath string) {
	t.Helper()

	if endpoint != expectedPath {
		t.Errorf("Expected endpoint '%s', got '%s'", expectedPath, endpoint)
	}
}

// DataStructureValidationTest validates data structure fields and JSON tags.
func DataStructureValidationTest(t *testing.T, obj interface{}, expectedFields []string) {
	t.Helper()

	ValidateStructFields(t, obj, expectedFields)
	ValidateJSONTags(t, obj)
}

// JSONSerializationTest tests JSON marshaling and unmarshaling.
type JSONSerializationTest struct {
	Name     string
	JSONData string
	DataType interface{}
}

// RunJSONSerializationTests executes JSON serialization tests.
func RunJSONSerializationTests(t *testing.T, tests []JSONSerializationTest) {
	t.Helper()

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			t.Helper()

			// Test unmarshaling
			err := json.Unmarshal([]byte(tt.JSONData), tt.DataType)
			ValidateNoError(t, err, "JSON unmarshaling")

			// Test marshaling back
			_, err = json.Marshal(tt.DataType)
			ValidateNoError(t, err, "JSON marshaling")
		})
	}
}
