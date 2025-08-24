// Package validation provides test validation utilities for the cisco-ios-xe-wireless-go library.
//
// This package contains validation functions for error handling, result validation,
// struct validation, and JSON serialization testing. These utilities support
// the testing framework by providing consistent validation patterns across
// all service tests.
//
// Key components:
//   - Error validation: ValidateNilClientError, ValidateContextError, etc.
//   - Result validation: ValidateNonNilResult, ValidateReflectionResults
//   - Struct validation: ValidateJSONSerialization, ValidateResponseStructure
//   - JSON validation: UnmarshalAndValidateJSON, RunParallelJSONSerializationTests
package validation
