//go:build scenario

// Package scenario provides test utilities for scenario tests.
package scenario

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/tests/testutil/integration"
)

// TagContext holds context for tag scenario tests.
// This is a simplified version focused only on tag scenarios.
type TagContext struct {
	T            *testing.T
	Ctx          context.Context
	Client       *core.Client
	TestTagName  string
	InitialCount int
	TagDeleted   bool
}

// SetupTag initializes context for tag scenario tests.
func SetupTag(t *testing.T, tagNameGenerator func() string) *TagContext {
	setup := integration.SetupTestClient(t)
	if setup.Client == nil {
		t.Skip("Skipping scenario tests: no client available")
	}

	return &TagContext{
		T:           t,
		Ctx:         setup.Context,
		Client:      setup.Client,
		TestTagName: tagNameGenerator(),
	}
}

// Cleanup handles tag deletion if not already deleted.
func (tc *TagContext) Cleanup() {
	if tc.TagDeleted {
		tc.T.Logf("Cleanup: Tag %s already deleted", tc.TestTagName)
		return
	}
	tc.T.Logf("Cleanup: Tag %s will be cleaned up by test", tc.TestTagName)
}

// SetInitialCount executes the count function and stores the result.
func (tc *TagContext) SetInitialCount(countFunc func() (int, error)) error {
	count, err := countFunc()
	if err != nil {
		return fmt.Errorf("failed to get initial count: %w", err)
	}
	tc.InitialCount = count
	tc.T.Logf("Initial tags count: %d", count)
	return nil
}

// ValidateTagCount executes the count function and validates expectations.
func (tc *TagContext) ValidateTagCount(countFunc func() (int, error), expectedCount int) error {
	finalCount, err := countFunc()
	if err != nil {
		return fmt.Errorf("failed to get final count: %w", err)
	}
	tc.T.Logf("Final tags count: %d (expected: %d)", finalCount, expectedCount)

	if finalCount != expectedCount {
		return fmt.Errorf("tag count mismatch: got %d, expected %d", finalCount, expectedCount)
	}

	tc.T.Logf("✅ Tag count validation successful")
	return nil
}

// GenerateTestIdentifier generates a unique identifier for test resources.
func GenerateTestIdentifier() string {
	return fmt.Sprintf("%d", time.Now().Unix()+int64(rand.Intn(1000)))
}

// MarkTagDeleted marks the tag as deleted to skip cleanup.
func (tc *TagContext) MarkTagDeleted() {
	tc.TagDeleted = true
}
