// Package errors provides common error definitions and utilities for Cisco IOS-XE Wireless Network Controller API services.
package errors

// Common error message templates for service operations.
const (
	// ErrOperationFailedTemplate is the generic operation failure message template
	// Usage: fmt.Sprintf(ErrOperationFailedTemplate, action, entity, entityID)
	// Example: fmt.Sprintf(ErrOperationFailedTemplate, "get", "AP", "28:ac:9e:11:48:10")
	//         fmt.Sprintf(ErrOperationFailedTemplate, "retrieve", "WLAN", "configuration")
	ErrOperationFailedTemplate = "failed to %s %s %s: %w"

	// ErrSimpleOperationFailedTemplate is the simple operation failure message template
	// Usage: fmt.Sprintf(ErrSimpleOperationFailedTemplate, action, err)
	// Example: fmt.Sprintf(ErrSimpleOperationFailedTemplate, "get configuration", err).
	ErrSimpleOperationFailedTemplate = "failed to %s: %w"

	// ErrInvalidParameterTemplate is the invalid parameter error message template
	// Usage: fmt.Sprintf(ErrInvalidParameterTemplate, parameterName, value)
	// Example: fmt.Sprintf(ErrInvalidParameterTemplate, "MAC address", "invalid-mac").
	ErrInvalidParameterTemplate = "invalid %s: %s"

	// ErrEntityNotFoundTemplate is the entity not found error message template
	// Usage: fmt.Sprintf(ErrEntityNotFoundTemplate, entityType, identifier)
	// Example: fmt.Sprintf(ErrEntityNotFoundTemplate, "AP", "28:ac:9e:11:48:10").
	ErrEntityNotFoundTemplate = "%s with %s not found"

	// ErrEmptyParameterTemplate is the empty parameter error message template
	// Usage: fmt.Sprintf(ErrEmptyParameterTemplate, parameterName)
	// Example: fmt.Sprintf(ErrEmptyParameterTemplate, "site tag").
	ErrEmptyParameterTemplate = "%s cannot be empty"

	// ErrParameterRequiredTemplate is the required parameter error message template
	// Usage: fmt.Sprintf(ErrParameterRequiredTemplate, parameterName)
	// Example: fmt.Sprintf(ErrParameterRequiredTemplate, "radio band").
	ErrParameterRequiredTemplate = "%s is required"
)

// Common error message constants.
const (
	// ErrClientNil is the error message when client is nil.
	ErrClientNil = "client cannot be nil"

	// ErrAtLeastOneTagRequired is the error message when no tags are specified.
	ErrAtLeastOneTagRequired = "at least one tag must be specified"

	// ErrUnsupportedOperation is the error message for unsupported operations.
	ErrUnsupportedOperation = "unsupported operation"

	// ErrDataUnavailable is the error message when data is not available.
	ErrDataUnavailable = "data not available"

	// ErrConfigurationInvalid is the error message for invalid configuration.
	ErrConfigurationInvalid = "invalid configuration"
)

// Entity type constants for consistent error messaging.
const (
	// EntityTypeAP represents the AP entity type.
	EntityTypeAP = "AP"

	// EntityTypeSite represents the site entity type.
	EntityTypeSite = "site"

	// EntityTypeWLAN represents the WLAN entity type.
	EntityTypeWLAN = "WLAN"

	// EntityTypeClient represents the client entity type.
	EntityTypeClient = "client"

	// EntityTypeRF represents the RF entity type.
	EntityTypeRF = "RF"

	// EntityTypePolicy represents the policy entity type.
	EntityTypePolicy = "policy"

	// EntityTypeTag represents the tag entity type.
	EntityTypeTag = "tag"

	// EntityTypeProfile represents the profile entity type.
	EntityTypeProfile = "profile"

	// EntityTypeConfiguration represents the configuration entity type.
	EntityTypeConfiguration = "configuration"

	// EntityTypeOperationalData represents the operational data entity type.
	EntityTypeOperationalData = "operational data"

	// EntityTypeGeolocation represents the geolocation entity type.
	EntityTypeGeolocation = "geolocation"
)

// Action constants for consistent error messaging.
const (
	// ActionGet represents a get/retrieve action.
	ActionGet = "get"

	// ActionRetrieve represents a retrieve action.
	ActionRetrieve = "retrieve"

	// ActionSet represents a set/update action.
	ActionSet = "set"

	// ActionUpdate represents an update action.
	ActionUpdate = "update"

	// ActionCreate represents a create action.
	ActionCreate = "create"

	// ActionDelete represents a delete action.
	ActionDelete = "delete"

	// ActionEnable represents an enable action.
	ActionEnable = "enable"

	// ActionDisable represents a disable action.
	ActionDisable = "disable"

	// ActionAssign represents an assign action.
	ActionAssign = "assign"

	// ActionValidate represents a validate action.
	ActionValidate = "validate"

	// ActionResolve represents a resolve action.
	ActionResolve = "resolve"
)
