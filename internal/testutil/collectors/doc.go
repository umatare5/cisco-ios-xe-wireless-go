// Package collectors provides data collection utilities for parallel test execution.
//
// This package contains thread-safe data collectors and concurrency management
// utilities for collecting and aggregating test results from parallel test
// execution. It supports both structure validation and generic service method
// testing with controlled concurrency levels.
//
// Key components:
//   - Data Collectors: StructValidationDataCollector, GenericTestDataCollector
//   - Parallel Execution: RunConcurrentValidationWithDataCollector
//   - Concurrency Control: GetOptimalConcurrencyLevel with environment variable support
//   - Test Method Management: TestMethod struct for method execution patterns
package collectors
