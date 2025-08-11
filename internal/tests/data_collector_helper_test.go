package tests

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"testing"
)

func TestNewGenericTestDataCollector(t *testing.T) {
	collector := NewGenericTestDataCollector()
	if collector == nil {
		t.Fatal("nil collector")
	}
	if collector.Results == nil {
		t.Error("Results not initialized")
	}
	if len(collector.Results) != 0 {
		t.Error("Results should be empty")
	}
}

func TestCollect(t *testing.T) {
	collector := NewGenericTestDataCollector()
	collector.Collect("TestMethod", "test response", nil)
	if len(collector.Results) != 1 {
		t.Errorf("expected 1 result, got %d", len(collector.Results))
	}
	res, ok := collector.Results["TestMethod"]
	if !ok {
		t.Error("missing result")
	}
	if res.Response != "test response" {
		t.Errorf("unexpected response: %v", res.Response)
	}
	if res.Error != nil {
		t.Errorf("unexpected error: %v", res.Error)
	}
	collector.Collect("ErrorMethod", nil, context.DeadlineExceeded)
	if len(collector.Results) != 2 {
		t.Errorf("expected 2 results, got %d", len(collector.Results))
	}
	errRes, ok := collector.Results["ErrorMethod"]
	if !ok {
		t.Error("missing error result")
	}
	if errRes.Response != nil {
		t.Errorf("expected nil response, got %v", errRes.Response)
	}
	if !errors.Is(errRes.Error, context.DeadlineExceeded) {
		t.Errorf("expected deadline, got %v", errRes.Error)
	}
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done()
			collector.Collect(fmt.Sprintf("ConcurrentMethod%d", i), i, nil)
		}(i)
	}
	wg.Wait()
	if len(collector.Results) != 12 {
		t.Errorf("expected 12 results, got %d", len(collector.Results))
	}
}
