package helpers

import (
	"context"
	"errors"
	"reflect"
	"strings"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/collectors"
)

// ServiceTester provides utilities for comprehensive service testing using reflection
type ServiceTester struct {
	service any
	ctx     context.Context
}

// NewServiceTester creates a new service tester for comprehensive testing
func NewServiceTester(service any, ctx context.Context) *ServiceTester {
	return &ServiceTester{
		service: service,
		ctx:     ctx,
	}
}

// CollectAllGetMethods automatically discovers and tests all Get* methods on the service
func (st *ServiceTester) CollectAllGetMethods(testParams map[string][]any) *collectors.GenericTestDataCollector {
	collector := collectors.NewGenericTestDataCollector()

	serviceValue := reflect.ValueOf(st.service)
	serviceType := reflect.TypeOf(st.service)

	// Find all methods that start with "Get"
	done := make(chan struct{}, serviceType.NumMethod())
	methodCount := 0

	for i := 0; i < serviceType.NumMethod(); i++ {
		method := serviceType.Method(i)

		// Skip non-Get methods
		if !strings.HasPrefix(method.Name, "Get") {
			continue
		}

		methodCount++

		// Get method value
		methodValue := serviceValue.Method(i)
		methodType := method.Type

		// Prepare arguments
		args := st.prepareArgs(method.Name, methodType, testParams)

		// Execute method in goroutine for concurrent testing
		go func(methodName string, methodVal reflect.Value, arguments []reflect.Value) {
			defer func() { done <- struct{}{} }()

			// Call the method
			results := methodVal.Call(arguments)

			// Extract result and error
			var response any
			var err error

			if len(results) == 2 {
				if !results[0].IsNil() {
					response = results[0].Interface()
				}
				if !results[1].IsNil() {
					if errVal, ok := results[1].Interface().(error); ok {
						err = errVal
					}
				}
			} else if len(results) == 1 {
				if results[0].Type().Implements(reflect.TypeOf((*error)(nil)).Elem()) {
					if !results[0].IsNil() {
						if errVal, ok := results[0].Interface().(error); ok {
							err = errVal
						}
					}
				} else {
					response = results[0].Interface()
				}
			}

			collector.Collect(methodName, response, err)
		}(method.Name, methodValue, args)
	}

	// Wait for all goroutines to complete
	for i := 0; i < methodCount; i++ {
		<-done
	}

	return collector
}

// prepareArgs prepares method arguments based on method signature and test parameters
func (st *ServiceTester) prepareArgs(
	methodName string,
	methodType reflect.Type,
	testParams map[string][]any,
) []reflect.Value {
	numIn := methodType.NumIn()
	args := make([]reflect.Value, numIn)

	// First argument is receiver (skip it)
	argIndex := 1

	// Second argument is typically context
	if argIndex < numIn && methodType.In(argIndex).String() == "context.Context" {
		args[argIndex] = reflect.ValueOf(st.ctx)
		argIndex++
	}

	// Fill remaining arguments with test parameters or zero values
	for i := argIndex; i < numIn; i++ {
		paramType := methodType.In(i)

		// Try to find specific parameters for this method
		if params, exists := testParams[methodName]; exists && len(params) > (i-argIndex) {
			args[i] = reflect.ValueOf(params[i-argIndex])
		} else {
			// Use zero value for the parameter type
			args[i] = reflect.Zero(paramType)
		}
	}

	return args
}

// TestAllMethods tests all methods on a service using reflection
func (st *ServiceTester) TestAllMethods() map[string]collectors.ServiceMethodResult {
	results := make(map[string]collectors.ServiceMethodResult)

	serviceValue := reflect.ValueOf(st.service)
	serviceType := reflect.TypeOf(st.service)

	for i := 0; i < serviceType.NumMethod(); i++ {
		method := serviceType.Method(i)
		methodValue := serviceValue.Method(i)

		// Prepare arguments (basic preparation)
		args := st.prepareBasicArgs(method.Type)

		// Call method and collect result
		func(methodName string) {
			defer func() {
				if r := recover(); r != nil {
					results[methodName] = collectors.ServiceMethodResult{
						Response: nil,
						Error:    errors.New("method panicked"),
					}
				}
			}()

			callResults := methodValue.Call(args)

			var response any
			var err error

			if len(callResults) >= 1 {
				if callResults[len(callResults)-1].Type().Implements(reflect.TypeOf((*error)(nil)).Elem()) {
					if !callResults[len(callResults)-1].IsNil() {
						if errVal, ok := callResults[len(callResults)-1].Interface().(error); ok {
							err = errVal
						}
					}
				}

				if len(callResults) == 2 && !callResults[0].IsNil() {
					response = callResults[0].Interface()
				}
			}

			results[methodName] = collectors.ServiceMethodResult{
				Response: response,
				Error:    err,
			}
		}(method.Name)
	}

	return results
}

// prepareBasicArgs prepares basic arguments for method calls
func (st *ServiceTester) prepareBasicArgs(methodType reflect.Type) []reflect.Value {
	numIn := methodType.NumIn()
	args := make([]reflect.Value, numIn)

	for i := 1; i < numIn; i++ { // Skip receiver
		paramType := methodType.In(i)

		// Handle common types
		switch paramType.String() {
		case "context.Context":
			args[i] = reflect.ValueOf(st.ctx)
		case "string":
			args[i] = reflect.ValueOf("test-value")
		case "int":
			args[i] = reflect.ValueOf(0)
		case "bool":
			args[i] = reflect.ValueOf(false)
		default:
			args[i] = reflect.Zero(paramType)
		}
	}

	return args
}

// GetMethodSignature returns the signature of a method as a string
func GetMethodSignature(methodType reflect.Type) string {
	var signature strings.Builder
	signature.WriteString("(")

	for i := 1; i < methodType.NumIn(); i++ { // Skip receiver
		if i > 1 {
			signature.WriteString(", ")
		}
		signature.WriteString(methodType.In(i).String())
	}

	signature.WriteString(") (")

	for i := 0; i < methodType.NumOut(); i++ {
		if i > 0 {
			signature.WriteString(", ")
		}
		signature.WriteString(methodType.Out(i).String())
	}

	signature.WriteString(")")
	return signature.String()
}
