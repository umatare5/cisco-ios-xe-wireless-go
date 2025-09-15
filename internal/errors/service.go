package errors

import "fmt"

// ServiceOperationError creates a standardized service operation error.
//
// This function provides consistent error formatting across all services
// using the pattern: "failed to {action} {serviceType} {entity}: {underlying error}"
//
// Parameters:
//   - action: The action being performed (get, set, update, etc.)
//   - serviceType: The service type (AWIPS, AFC, Client, etc.)
//   - entity: The entity or operation target
//   - err: The underlying error
//
// Returns:
//   - error: Formatted error with consistent structure
//
// Example:
//
//	err := ServiceOperationError("get", "AWIPS", "operational data", originalErr)
//	// Result: "failed to get AWIPS operational data: original error"
func ServiceOperationError(action, serviceType, entity string, err error) error {
	return fmt.Errorf(ErrOperationFailedTemplate, action, serviceType, entity, err)
}

// SimpleServiceError creates a simplified service error without entity specification.
//
// This function provides a simpler error format for cases where the entity
// is not specific or is part of the action description.
//
// Parameters:
//   - action: The action being performed
//   - err: The underlying error
//
// Returns:
//   - error: Formatted error with simple structure
//
// Example:
//
//	err := SimpleServiceError("retrieve configuration", originalErr)
//	// Result: "failed to retrieve configuration: original error"
func SimpleServiceError(action string, err error) error {
	return fmt.Errorf(ErrSimpleOperationFailedTemplate, action, err)
}

// ValidationError creates a standardized validation error.
//
// This function provides consistent validation error formatting
// for parameter validation across all services.
//
// Parameters:
//   - parameter: The parameter name being validated
//   - value: The invalid value (as string)
//
// Returns:
//   - error: Formatted validation error
//
// Example:
//
//	err := ValidationError("MAC address", "invalid-mac")
//	// Result: "invalid MAC address: invalid-mac"
func ValidationError(parameter, value string) error {
	return fmt.Errorf(ErrInvalidParameterTemplate, parameter, value)
}

// RequiredParameterError creates a standardized required parameter error.
//
// This function provides consistent error formatting for missing
// required parameters across all services.
//
// Parameters:
//   - parameter: The required parameter name
//
// Returns:
//   - error: Formatted required parameter error
//
// Example:
//
//	err := RequiredParameterError("AP MAC address")
//	// Result: "AP MAC address is required"
func RequiredParameterError(parameter string) error {
	return fmt.Errorf(ErrParameterRequiredTemplate, parameter)
}

// EmptyParameterError creates a standardized empty parameter error.
//
// This function provides consistent error formatting for empty
// parameters that should not be empty.
//
// Parameters:
//   - parameter: The parameter name that is empty
//
// Returns:
//   - error: Formatted empty parameter error
//
// Example:
//
//	err := EmptyParameterError("site tag")
//	// Result: "site tag cannot be empty"
func EmptyParameterError(parameter string) error {
	return fmt.Errorf(ErrEmptyParameterTemplate, parameter)
}

// NotFoundError creates a standardized entity not found error.
//
// This function provides consistent error formatting for cases
// where an entity with specific identifier is not found.
//
// Parameters:
//   - entityType: The type of entity (AP, WLAN, etc.)
//   - identifier: The identifier used to search for the entity
//
// Returns:
//   - error: Formatted not found error
//
// Example:
//
//	err := NotFoundError("AP", "28:ac:9e:11:48:10")
//	// Result: "AP with 28:ac:9e:11:48:10 not found"
func NotFoundError(entityType, identifier string) error {
	return fmt.Errorf(ErrEntityNotFoundTemplate, entityType, identifier)
}
