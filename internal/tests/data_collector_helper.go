package tests

import "sync"

// GenericTestDataCollector provides a generic data collector for service tests
type GenericTestDataCollector struct {
	mu      sync.Mutex
	Results map[string]ServiceMethodResult
}

// ServiceMethodResult holds the result of a service method call
type ServiceMethodResult struct {
	Response interface{}
	Error    error
}

// NewGenericTestDataCollector creates a new generic test data collector
func NewGenericTestDataCollector() *GenericTestDataCollector {
	return &GenericTestDataCollector{
		Results: make(map[string]ServiceMethodResult),
	}
}

// Collect stores the result of a service method call
func (c *GenericTestDataCollector) Collect(methodName string, response interface{}, err error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Results[methodName] = ServiceMethodResult{
		Response: response,
		Error:    err,
	}
}
