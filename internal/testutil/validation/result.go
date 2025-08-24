// Package validation provides test validation functions
package validation

import (
	"testing"
)

// ServiceMethodResult represents the result of a service method call
type ServiceMethodResult struct {
	Response any
	Error    error
}

// ValidateNonNilResult validates that a result is not nil
func ValidateNonNilResult(t *testing.T, result any, methodName string) {
	t.Helper()
	if result == nil {
		t.Errorf("Method %s returned nil result", methodName)
	}
}

// ValidateReflectionResults validates and logs results from ServiceTester reflection
func ValidateReflectionResults(t *testing.T, results map[string]ServiceMethodResult, testType string) {
	t.Helper()

	if len(results) == 0 {
		t.Error("No reflection test results found")
		return
	}

	successCount := 0
	for methodName, result := range results {
		if result.Error != nil {
			switch testType {
			case "unit":
				t.Logf("Method %s returned error: %v", methodName, result.Error)
			case "table-driven":
				t.Logf("%s: Error (may be expected): %v", methodName, result.Error)
			case "integration":
				t.Logf("%s: API error (may be expected): %v", methodName, result.Error)
			default:
				t.Logf("%s error: %v", methodName, result.Error)
			}
		} else {
			successCount++
			switch testType {
			case "unit":
				t.Logf("Method %s returned result", methodName)
			case "table-driven":
				t.Logf("%s: Success", methodName)
			case "integration":
				t.Logf("%s: Success with live data", methodName)
			default:
				t.Logf("%s success", methodName)
			}
		}
	}

	// Log summary based on test type
	switch testType {
	case "unit":
		t.Logf("Unit test completed: tested %d methods via reflection", len(results))
	case "table-driven":
		t.Logf("Table-driven test completed: %d Get methods tested via reflection", len(results))
	case "integration":
		t.Logf("Integration test completed: %d methods tested against live controller", len(results))
		t.Logf("Integration summary: %d/%d methods successful", successCount, len(results))
	default:
		t.Logf("Reflection test completed: %d methods tested", len(results))
	}
}
