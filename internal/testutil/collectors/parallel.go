package collectors

import (
	"os"
	"strconv"
	"sync"
	"testing"
)

// TestMethod represents a generic test method for parallel execution
type TestMethod struct {
	Name   string
	Method func() (any, error)
}

// StructValidationCase defines a single structure validation test case
type StructValidationCase struct {
	Name           string
	ResponseFunc   func() any
	ExpectedStruct string
}

// StructValidationDataCollector implements TS-009: DataCollector pattern with sync.Mutex-protected error collection
type StructValidationDataCollector struct {
	mu      sync.Mutex
	results map[string]any
	errors  []error
}

// NewStructValidationDataCollector creates a new DataCollector instance
func NewStructValidationDataCollector() *StructValidationDataCollector {
	return &StructValidationDataCollector{
		results: make(map[string]any),
	}
}

// CollectResult stores validation result in a thread-safe manner
func (dc *StructValidationDataCollector) CollectResult(name string, result any, err error) {
	dc.mu.Lock()
	defer dc.mu.Unlock()

	if err != nil {
		dc.errors = append(dc.errors, err)
	} else {
		dc.results[name] = result
	}
}

// GetAllResults returns collected results and errors
func (dc *StructValidationDataCollector) GetAllResults() (map[string]any, []error) {
	dc.mu.Lock()
	defer dc.mu.Unlock()

	// Return copies to avoid race conditions
	resultsCopy := make(map[string]any)
	for k, v := range dc.results {
		resultsCopy[k] = v
	}

	errorsCopy := make([]error, len(dc.errors))
	copy(errorsCopy, dc.errors)

	return resultsCopy, errorsCopy
}

// GenericTestDataCollector provides a generic data collector for service tests
type GenericTestDataCollector struct {
	mu      sync.Mutex
	Results map[string]ServiceMethodResult
	Errors  []error
}

// ServiceMethodResult holds the result of a service method call
type ServiceMethodResult struct {
	Response any
	Error    error
}

// NewGenericTestDataCollector creates a new generic test data collector
func NewGenericTestDataCollector() *GenericTestDataCollector {
	return &GenericTestDataCollector{
		Results: make(map[string]ServiceMethodResult),
		Errors:  make([]error, 0),
	}
}

// Collect stores the result of a service method call
func (c *GenericTestDataCollector) Collect(methodName string, response any, err error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Results[methodName] = ServiceMethodResult{
		Response: response,
		Error:    err,
	}
	// Also collect the error if it exists
	if err != nil {
		c.Errors = append(c.Errors, err)
	}
}

// CollectConcurrently executes all provided methods concurrently and collects their results
func (c *GenericTestDataCollector) CollectConcurrently(methods []TestMethod) {
	var wg sync.WaitGroup
	for _, method := range methods {
		wg.Add(1)
		go func(m TestMethod) {
			defer wg.Done()
			response, err := m.Method()
			c.Collect(m.Name, response, err)
		}(method)
	}
	wg.Wait()
}

// GetOptimalConcurrencyLevel determines optimal concurrency level for parallel tests
func GetOptimalConcurrencyLevel() int {
	if concurrencyStr := os.Getenv("TESTUTIL_CONCURRENCY"); concurrencyStr != "" {
		if level, err := strconv.Atoi(concurrencyStr); err == nil && level > 0 {
			return level
		}
	}
	// Default to 4 for balanced test execution
	return 4
}

// RunConcurrentValidationWithDataCollector executes validation functions concurrently with data collection
func RunConcurrentValidationWithDataCollector(
	t *testing.T,
	validationFunctions map[string]func() any,
) *StructValidationDataCollector {
	t.Helper()

	collector := NewStructValidationDataCollector()
	concurrency := GetOptimalConcurrencyLevel()

	// Channel-based worker pool for controlled concurrency
	work := make(chan StructValidationCase, len(validationFunctions))
	var wg sync.WaitGroup

	// Start workers
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for testCase := range work {
				result := testCase.ResponseFunc()
				collector.CollectResult(testCase.Name, result, nil)
			}
		}()
	}

	// Send work
	for name, fn := range validationFunctions {
		work <- StructValidationCase{
			Name:         name,
			ResponseFunc: fn,
		}
	}
	close(work)

	// Wait for completion
	wg.Wait()

	return collector
}

// CreateStandardValidationFunctions creates standard validation functions from test cases
func CreateStandardValidationFunctions(testCases map[string]func() any) map[string]func() any {
	// Return the test cases as-is since they're already validation functions
	return testCases
}
