// Package helper provides testing utilities and assertion functions for.
package helper

import (
	"reflect"
	"strings"
	"testing"
	"time"
)

// AssertBoolEquals validates that the actual boolean value matches the expected boolean value.
func AssertBoolEquals(t *testing.T, actual, expected bool, message string) {
	t.Helper()

	if actual != expected {
		t.Errorf("%s: Expected %t, got %t", message, expected, actual)
	}
}

// AssertTrue validates that a boolean value is true.
func AssertTrue(t *testing.T, actual bool, context string) {
	t.Helper()

	if !actual {
		t.Errorf("%s: Expected true, got false", context)
	}
}

// AssertFalse validates that a boolean value is false.
func AssertFalse(t *testing.T, actual bool, context string) {
	t.Helper()

	if actual {
		t.Errorf("%s: Expected false, got true", context)
	}
}

// AssertPointerNil validates that a pointer is nil.
func AssertPointerNil(t *testing.T, actual any, context string) {
	t.Helper()

	if actual == nil {
		return // Success - pointer is nil
	}

	// Use reflection to check if the pointer value is nil
	rv := reflect.ValueOf(actual)
	if rv.Kind() == reflect.Ptr && rv.IsNil() {
		return // Success - pointer value is nil
	}

	t.Errorf("%s: Expected pointer to be nil, got %v", context, actual)
}

// AssertClientCreated validates that a client was successfully created.
func AssertClientCreated(t *testing.T, client any, err error, context string) {
	t.Helper()

	if err != nil {
		t.Fatalf("%s: Expected successful client creation, got error: %v", context, err)
	}
	if client == nil {
		t.Fatalf("%s: Expected non-nil client", context)
	}
}

// AssertClientCreationError validates that client creation failed with an error.
func AssertClientCreationError(t *testing.T, err error, context string) {
	t.Helper()

	if err == nil {
		t.Fatalf("%s: Expected error but got none", context)
	}
}

// AssertError validates that an error occurred.
func AssertError(t *testing.T, err error, context string) {
	t.Helper()

	if err == nil {
		t.Fatalf("%s: Expected error but got none", context)
	}
}

// AssertNoError validates that no error occurred.
func AssertNoError(t *testing.T, err error, context string) {
	t.Helper()

	if err != nil {
		t.Fatalf("%s: Expected no error, got: %v", context, err)
	}
}

// AssertErrorContains validates that an error occurred and contains specific text.
func AssertErrorContains(t *testing.T, err error, expectedText, context string) {
	t.Helper()

	if err == nil {
		t.Fatalf("%s: Expected error but got none", context)
	}
	if !strings.Contains(err.Error(), expectedText) {
		t.Fatalf("%s: Expected error containing '%s', got: %v", context, expectedText, err)
	}
}

// AssertNotNil validates that a value is not nil.
func AssertNotNil(t *testing.T, value any, context string) {
	t.Helper()

	if value == nil {
		t.Fatalf("%s: Expected non-nil value", context)
	}
}

// AssertNil validates that a value is nil.
func AssertNil(t *testing.T, value any, context string) {
	t.Helper()

	if value != nil {
		t.Fatalf("%s: Expected nil value, got: %v", context, value)
	}
}

// RequireEnvironmentVars validates that all required environment variables are set.
func RequireEnvironmentVars(t *testing.T, vars map[string]string) {
	t.Helper()

	for name, value := range vars {
		if value == "" {
			t.Skipf("%s environment variable must be set for integration tests", name)
		}
	}
}

// AssertStringEquals validates that two strings are equal.
func AssertStringEquals(t *testing.T, actual, expected, context string) {
	t.Helper()

	if actual != expected {
		t.Errorf("%s: Expected '%s', got '%s'", context, expected, actual)
	}
}

// AssertStringNotEmpty validates that a string is not empty.
func AssertStringNotEmpty(t *testing.T, value, context string) {
	t.Helper()

	if value == "" {
		t.Error(context + ": Expected non-empty string")
	}
}

// AssertErrorMessage validates that an error has the expected message.
func AssertErrorMessage(t *testing.T, err error, expected, context string) {
	t.Helper()

	if err == nil {
		t.Fatalf("%s: Expected error but got none", context)
	}
	if err.Error() != expected {
		t.Errorf("%s: Expected error message '%s', got '%s'", context, expected, err.Error())
	}
}

// AssertIntEquals validates that two integer values are equal.
func AssertIntEquals(t *testing.T, actual, expected int, context string) {
	t.Helper()

	if actual != expected {
		t.Errorf("%s: Expected count %d, got %d", context, expected, actual)
	}
}

// AssertStringContains validates that a string contains specific text.
func AssertStringContains(t *testing.T, text, expectedSubstring, context string) {
	t.Helper()

	if !strings.Contains(text, expectedSubstring) {
		t.Errorf("%s: Expected text to contain '%s', got '%s'", context, expectedSubstring, text)
	}
}

// AssertDurationEquals validates that two time.Duration values are equal.
func AssertDurationEquals(t *testing.T, actual, expected time.Duration, context string) {
	t.Helper()

	if actual != expected {
		t.Errorf("%s: Expected duration %v, got %v", context, expected, actual)
	}
}

// AssertPointerEquals validates that two pointers are equal.
func AssertPointerEquals(t *testing.T, actual, expected any, context string) {
	t.Helper()

	if actual != expected {
		t.Errorf("%s: Expected pointer %v, got %v", context, expected, actual)
	}
}
