// Package validation provides test validation functions
package validation

import (
	"strings"
	"testing"
)

// ValidateNilClientError validates that an error occurred due to nil client
func ValidateNilClientError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Fatal("Expected error with nil client, but got nil")
	}
}

// ValidateContextError validates that an error occurred due to context issues (nil/canceled)
func ValidateContextError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Fatal("Expected error with invalid context, but got nil")
	}
	t.Logf("ValidateContextError: correctly returned context error: %v", err)
}

// ValidateExpectedError validates that an expected error containing specific text occurred
func ValidateExpectedError(t *testing.T, err error, expectedContains string) {
	t.Helper()
	if err == nil {
		t.Fatalf("Expected error containing '%s', but got nil", expectedContains)
	}
	if !strings.Contains(err.Error(), expectedContains) {
		t.Fatalf("Expected error containing '%s', but got: %v", expectedContains, err)
	}
}

// ValidateNoError validates that no error occurred
func ValidateNoError(t *testing.T, err error, operation string) {
	t.Helper()
	if err != nil {
		t.Fatalf("Unexpected error during %s: %v", operation, err)
	}
}
