//go:build scenario

package client

import (
	"context"
	"testing"
	"time"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
)

// TagScenarioStep represents the current step in tag scenario testing.
type TagScenarioStep int

const (
	StepInitialList TagScenarioStep = iota + 1
	StepCreate
	StepConfigure
	StepGet
	StepDelete
	StepFinalList
)

// TagScenarioContext holds common context for tag scenario tests.
type TagScenarioContext struct {
	T            *testing.T
	Ctx          context.Context
	Client       *core.Client
	TestTagName  string
	InitialCount int
	TagDeleted   bool // Track if tag was successfully deleted
}

// InitializeTagScenario sets up the common context for tag scenario tests.
func InitializeTagScenario(t *testing.T, tagNameGenerator func() string) (*TagScenarioContext, bool) {
	// Initialize client using integration test utilities
	setup := SetupRequiredClient(t)

	return &TagScenarioContext{
		T:           t,
		Ctx:         setup.Context,
		Client:      setup.Client,
		TestTagName: tagNameGenerator(),
	}, false
}

// LogStep logs the current step being executed.
func (tsc *TagScenarioContext) LogStep(step TagScenarioStep, message string) {
	tsc.T.Logf("Step %d: %s", int(step), message)
}

// LogStepWithTag logs the current step with tag name.
func (tsc *TagScenarioContext) LogStepWithTag(step TagScenarioStep, message string) {
	tsc.T.Logf("Step %d: %s: %s", int(step), message, tsc.TestTagName)
}

// LogSuccess logs a successful operation with tag name.
func (tsc *TagScenarioContext) LogSuccess(operation string) {
	tsc.T.Logf("Successfully %s test tag: %s", operation, tsc.TestTagName)
}

// ValidateTagCount validates that the tag count matches expectations.
func (tsc *TagScenarioContext) ValidateTagCount(finalCount int) {
	logFinalTagCount(tsc.T, finalCount)
	logScenarioCompletion(tsc.T)

	// Verify tag count returned to initial state
	if finalCount != tsc.InitialCount {
		logCountMismatch(tsc.T, finalCount, tsc.InitialCount)
	}
}

// SetInitialCount stores the initial tag count for later validation.
func (tsc *TagScenarioContext) SetInitialCount(count int) {
	tsc.InitialCount = count
	logInitialTagCount(tsc.T, count)
}

// CreateCleanupFunction returns a cleanup function for tag deletion.
func (tsc *TagScenarioContext) CreateCleanupFunction(deleteFunc func(context.Context, string) error) func() {
	return func() {
		if tsc.TagDeleted {
			logCleanupSkipped(tsc.T, tsc.TestTagName)
			return
		}

		logCleanupAttempt(tsc.T, tsc.TestTagName)
		if deleteErr := deleteFunc(tsc.Ctx, tsc.TestTagName); deleteErr != nil {
			logCleanupFailure(tsc.T, tsc.TestTagName, deleteErr)
		} else {
			logCleanupSuccess(tsc.T, tsc.TestTagName)
		}
	}
}

// ValidateConfigNotNil validates that retrieved configuration is not nil.
func ValidateConfigNotNil(t *testing.T, config any, tagName string) {
	t.Helper()
	if config == nil {
		t.Errorf("Tag '%s' configuration should not be nil", tagName)
	}
}

// ValidateCommonTagFields validates common tag fields like name and description.
func ValidateCommonTagFields(t *testing.T, expectedTagName, expectedDescription, actualTagName, actualDescription string) {
	validateTagName(t, expectedTagName, actualTagName)
	validateTagDescription(t, expectedDescription, actualDescription)
}

// ExecuteStep executes a step with error handling and logging.
func (tsc *TagScenarioContext) ExecuteStep(step TagScenarioStep, description string, operation func() error) {
	tsc.LogStepWithTag(step, description)
	if err := operation(); err != nil {
		logStepError(tsc.T, description, tsc.TestTagName, err)
	}
}

// ExecuteStepWithResult executes a step with a result and error handling.
func ExecuteStepWithResult[T any](tsc *TagScenarioContext, step TagScenarioStep, description string, operation func() (T, error)) T {
	tsc.LogStepWithTag(step, description)
	result, err := operation()
	if err != nil {
		logStepError(tsc.T, description, tsc.TestTagName, err)
	}
	return result
}

// ExecuteDeleteStep executes a delete step and marks the tag as deleted on success.
func (tsc *TagScenarioContext) ExecuteDeleteStep(description string, operation func() error) {
	tsc.LogStepWithTag(StepDelete, description)
	if err := operation(); err != nil {
		logStepError(tsc.T, description, tsc.TestTagName, err)
	}
	// Mark tag as deleted on successful deletion
	tsc.TagDeleted = true
}

// GenerateTestTagName creates a unique test tag name with timestamp.
// Format: test-tag-{yyyymmdd-hhmmss}
func GenerateTestTagName() string {
	timestamp := time.Now().Format("20060102-150405")
	return "test-tag-" + timestamp
}

// validateTagName validates that tag names match.
func validateTagName(t *testing.T, expected, actual string) {
	if actual != expected {
		t.Errorf("Tag name mismatch: expected %s, got %s", expected, actual)
	}
}

// validateTagDescription validates that tag descriptions match.
func validateTagDescription(t *testing.T, expected, actual string) {
	if actual != expected {
		t.Errorf("Description mismatch: expected %s, got %s", expected, actual)
	}
}

// logStepError logs step execution errors and fails the test.
func logStepError(t *testing.T, description, tagName string, err error) {
	t.Fatalf("Failed to %s for test tag %s: %v", description, tagName, err)
}

// logInitialTagCount logs the initial tag count.
func logInitialTagCount(t *testing.T, count int) {
	t.Logf("Initial tags count: %d", count)
}

// logFinalTagCount logs the final tag count.
func logFinalTagCount(t *testing.T, count int) {
	t.Logf("Final tags count: %d", count)
}

// logScenarioCompletion logs successful scenario completion.
func logScenarioCompletion(t *testing.T) {
	t.Logf("Tag scenario test completed successfully")
}

// logCountMismatch logs when final count differs from initial count.
func logCountMismatch(t *testing.T, finalCount, initialCount int) {
	t.Logf("Warning: Final tag count (%d) differs from initial count (%d)", finalCount, initialCount)
}

// logCleanupSkipped logs when cleanup is skipped.
func logCleanupSkipped(t *testing.T, tagName string) {
	t.Logf("Cleanup: Test tag %s was already deleted in test, skipping cleanup", tagName)
}

// logCleanupAttempt logs cleanup attempt.
func logCleanupAttempt(t *testing.T, tagName string) {
	t.Logf("Cleanup: Attempting to delete test tag: %s", tagName)
}

// logCleanupFailure logs cleanup failure.
func logCleanupFailure(t *testing.T, tagName string, err error) {
	t.Logf("Warning: Failed to cleanup test tag %s: %v", tagName, err)
}

// logCleanupSuccess logs successful cleanup.
func logCleanupSuccess(t *testing.T, tagName string) {
	t.Logf("Successfully cleaned up test tag: %s", tagName)
}
