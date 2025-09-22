package testutil

import (
	"errors"
	"testing"
	"time"
)

func TestTestUtilHelperUnit_AssertBoolEquals_Success(t *testing.T) {
	AssertBoolEquals(t, true, true, "test message")
	AssertBoolEquals(t, false, false, "test message")
}

func TestTestUtilHelperUnit_AssertTrue_Success(t *testing.T) {
	AssertTrue(t, true, "test context")
}

func TestTestUtilHelperUnit_AssertFalse_Success(t *testing.T) {
	AssertFalse(t, false, "test context")
}

func TestTestUtilHelperUnit_AssertStringEquals_Success(t *testing.T) {
	AssertStringEquals(t, "expected", "expected", "test context")
}

func TestTestUtilHelperUnit_AssertIntEquals_Success(t *testing.T) {
	AssertIntEquals(t, 42, 42, "test context")
}

func TestTestUtilHelperUnit_AssertDurationEquals_Success(t *testing.T) {
	duration := 30 * time.Second
	AssertDurationEquals(t, duration, duration, "test context")
}

func TestTestUtilHelperUnit_AssertPointerNil_Success(t *testing.T) {
	AssertPointerNil(t, nil, "test context")
	var nilPtr *string
	AssertPointerNil(t, nilPtr, "test context with nil pointer")
}

func TestTestUtilHelperUnit_AssertClientCreated_Success(t *testing.T) {
	client := "mock client"
	AssertClientCreated(t, client, nil, "test context")
}

func TestTestUtilHelperUnit_AssertErrorContains_Success(t *testing.T) {
	err := errors.New("this is a test error message")
	AssertErrorContains(t, err, "test error", "test context")
}

func TestTestUtilHelperUnit_AssertErrorMessage_Success(t *testing.T) {
	err := errors.New("exact error message")
	AssertErrorMessage(t, err, "exact error message", "test context")
}

func TestTestUtilHelperUnit_AssertNotNil_Success(t *testing.T) {
	value := "not nil"
	AssertNotNil(t, &value, "test context")
}

func TestTestUtilHelperUnit_AssertNil_Success(t *testing.T) {
	AssertNil(t, nil, "test context")
}

func TestTestUtilHelperUnit_AssertStringNotEmpty_Success(t *testing.T) {
	AssertStringNotEmpty(t, "not empty", "test context")
}

func TestTestUtilHelperUnit_AssertClientCreationError_Success(t *testing.T) {
	err := errors.New("client creation error")
	AssertClientCreationError(t, err, "test context")
}

func TestTestUtilHelperUnit_AssertError_Success(t *testing.T) {
	err := errors.New("test error")
	AssertError(t, err, "test context")
}

func TestTestUtilHelperUnit_AssertNoError_Success(t *testing.T) {
	AssertNoError(t, nil, "test context")
}

func TestTestUtilHelperUnit_AssertPointerEquals_Success(t *testing.T) {
	// Test with same pointers
	ptr1 := &testing.T{}
	ptr2 := ptr1
	AssertPointerEquals(t, ptr2, ptr1, "test context")

	// Test with both nil pointers
	AssertPointerEquals(t, nil, nil, "test context")
}

func TestTestUtilHelperUnit_AssertStringContains_Success(t *testing.T) {
	AssertStringContains(t, "hello world", "world", "test context")
	AssertStringContains(t, "testing string", "test", "test context")
}
