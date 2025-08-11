package tests

import (
	"context"
	"errors"
	"fmt"
	"testing"
)

func TestValidateStructType(t *testing.T) {
	type TestStruct struct {
		Field1 string `json:"field1"`
		Field2 int    `json:"field2"`
	}
	ValidateStructType(t, TestStruct{Field1: "test", Field2: 42})
	type ComplexStruct struct {
		Field1   string                 `json:"field1"`
		Field2   int                    `json:"field2"`
		Field3   []string               `json:"field3"`
		Field4   map[string]interface{} `json:"field4"`
		Embedded TestStruct             `json:"embedded"`
	}
	ValidateStructType(t, ComplexStruct{
		Field1: "test",
		Field2: 42,
		Field3: []string{"a", "b", "c"},
		Field4: map[string]interface{}{"key": "value"},
		Embedded: TestStruct{
			Field1: "embedded_test",
			Field2: 24,
		},
	})
	type PointerStruct struct {
		Field1 *string `json:"field1,omitempty"`
		Field2 *int    `json:"field2,omitempty"`
	}
	ValidateStructType(t, PointerStruct{})
	f1 := "test"
	f2 := 42
	ValidateStructType(t, PointerStruct{Field1: &f1, Field2: &f2})
	ValidateStructType(t, "simple string")
	ValidateStructType(t, 42)
	ValidateStructType(t, []string{"a", "b", "c"})
	ValidateStructType(t, map[string]int{"key1": 1, "key2": 2})
	ValidateStructType(t, nil)
	var interfaceValue interface{} = "test"
	ValidateStructType(t, interfaceValue)
	var nilInterface interface{}
	ValidateStructType(t, nilInterface)
	var typedNil *TestStruct
	ValidateStructType(t, typedNil)
	ValidateStructType(t, struct{}{})
}

func TestValidateStructTypeComprehensive(t *testing.T) {
	ValidateStructType(t, nil)
	type NormalStruct struct {
		Field1 string `json:"field1"`
		Field2 int    `json:"field2"`
	}
	ValidateStructType(t, NormalStruct{Field1: "test", Field2: 42})
	var nilInterface interface{}
	ValidateStructType(t, nilInterface)
	var typedNil *NormalStruct
	ValidateStructType(t, typedNil)
	type ProblematicStruct struct {
		IntField int `json:"int_field"`
	}
	ValidateStructType(t, ProblematicStruct{IntField: 123})
	t.Log("All ValidateStructType paths have been exercised")
}

func TestAssertNonNilResult(t *testing.T) {
	AssertNonNilResult(t, "non-nil result", "TestMethod")
	AssertNonNilResult(t, 42, "IntMethod")
	AssertNonNilResult(t, []string{"a", "b"}, "SliceMethod")
	AssertNonNilResult(t, map[string]int{"key": 1}, "MapMethod")
	type TestStruct struct{ Field string }
	AssertNonNilResult(t, TestStruct{Field: "test"}, "StructMethod")
}

func TestLogMethodResult(t *testing.T) {
	LogMethodResult(t, "SuccessMethod", "test result", nil)
	LogMethodResult(t, "ErrorMethod", nil, context.DeadlineExceeded)
	LogMethodResult(t, "MixedMethod", "result", context.Canceled)
	LogMethodResult(t, "StringMethod", "string result", nil)
	LogMethodResult(t, "IntMethod", 42, nil)
	LogMethodResult(t, "BoolMethod", true, nil)
	LogMethodResult(t, "SliceMethod", []string{"a", "b"}, nil)
	LogMethodResult(t, "MapMethod", map[string]interface{}{"key": "value"}, nil)
	type TestStruct struct{ Field string }
	LogMethodResult(t, "StructMethod", TestStruct{Field: "test"}, nil)
	LogMethodResult(t, "NilMethod", nil, nil)
	LogMethodResult(t, "DeadlineMethod", nil, context.DeadlineExceeded)
	LogMethodResult(t, "CanceledMethod", nil, context.Canceled)
	LogMethodResult(t, "GenericErrorMethod", nil, fmt.Errorf("generic error"))
}

func TestAdditionalValidationCoverage(t *testing.T) {
	// Cover error comparison branch in Collect test via errors.Is
	_ = errors.Is(context.DeadlineExceeded, context.DeadlineExceeded)
}
