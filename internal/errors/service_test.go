package errors

import (
	"errors"
	"fmt"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/helper"
)

func TestErrorsServiceUnit_ServiceOperationError_Success(t *testing.T) {
	originalErr := errors.New("connection timeout")

	t.Run("StandardFormat", func(t *testing.T) {
		err := ServiceOperationError("get", "AWIPS", "operational data", originalErr)
		expected := "failed to get AWIPS operational data: connection timeout"

		helper.AssertErrorMessage(t, err, expected, "ServiceOperationError should format message correctly")
	})

	t.Run("DifferentParameters", func(t *testing.T) {
		err := ServiceOperationError("update", "Client", "configuration", originalErr)
		expected := "failed to update Client configuration: connection timeout"

		helper.AssertErrorMessage(t, err, expected,
			"ServiceOperationError should format message correctly for different parameters")
	})

	t.Run("ErrorUnwrapping", func(t *testing.T) {
		err := ServiceOperationError("delete", "AP", "profile", originalErr)

		if !errors.Is(err, originalErr) {
			helper.AssertTrue(t, errors.Is(err, originalErr), "Expected error to wrap original error")
		}
	})
}

func TestErrorsServiceUnit_SimpleServiceError_Success(t *testing.T) {
	originalErr := errors.New("network timeout")

	t.Run("StandardFormat", func(t *testing.T) {
		err := SimpleServiceError("retrieve configuration", originalErr)
		expected := "failed to retrieve configuration: network timeout"

		helper.AssertErrorMessage(t, err, expected, "SimpleServiceError should format message correctly")
	})

	t.Run("DifferentAction", func(t *testing.T) {
		err := SimpleServiceError("connect to server", originalErr)
		expected := "failed to connect to server: network timeout"

		helper.AssertErrorMessage(t, err, expected,
			"SimpleServiceError should format message correctly for different actions")
	})

	t.Run("ErrorUnwrapping", func(t *testing.T) {
		err := SimpleServiceError("load data", originalErr)

		if !errors.Is(err, originalErr) {
			helper.AssertTrue(t, errors.Is(err, originalErr), "Expected error to wrap original error")
		}
	})
}

func TestErrorsServiceUnit_ValidationError_Success(t *testing.T) {
	t.Run("MACAddress", func(t *testing.T) {
		err := ValidationError("MAC address", "invalid-mac")
		expected := "invalid MAC address: invalid-mac"

		helper.AssertErrorMessage(t, err, expected, "ValidationError should format MAC address message correctly")
	})

	t.Run("WLANId", func(t *testing.T) {
		err := ValidationError("WLAN ID", "-1")
		expected := "invalid WLAN ID: -1"

		helper.AssertErrorMessage(t, err, expected, "ValidationError should format WLAN ID message correctly")
	})

	t.Run("EmptyValue", func(t *testing.T) {
		err := ValidationError("parameter", "")
		expected := "invalid parameter: "

		helper.AssertErrorMessage(t, err, expected, "ValidationError should format empty value message correctly")
	})
}

func TestErrorsServiceUnit_RequiredParameterError_Success(t *testing.T) {
	t.Run("APMacAddress", func(t *testing.T) {
		err := RequiredParameterError("AP MAC address")
		expected := "AP MAC address is required"

		helper.AssertErrorMessage(t, err, expected, "RequiredParameterError should format AP MAC message correctly")
	})

	t.Run("RadioBand", func(t *testing.T) {
		err := RequiredParameterError("radio band")
		expected := "radio band is required"

		helper.AssertErrorMessage(t, err, expected,
			"RequiredParameterError should format radio band message correctly")
	})
}

func TestErrorsServiceUnit_EmptyParameterError_Success(t *testing.T) {
	t.Run("SiteTag", func(t *testing.T) {
		err := EmptyParameterError("site tag")
		expected := "site tag cannot be empty"

		helper.AssertErrorMessage(t, err, expected, "EmptyParameterError should format site tag message correctly")
	})

	t.Run("PolicyTag", func(t *testing.T) {
		err := EmptyParameterError("policy tag")
		expected := "policy tag cannot be empty"

		helper.AssertErrorMessage(t, err, expected, "EmptyParameterError should format policy tag message correctly")
	})
}

func TestErrorsServiceUnit_NotFoundError_Success(t *testing.T) {
	t.Run("APWithMAC", func(t *testing.T) {
		err := NotFoundError("AP", "28:ac:9e:11:48:10")
		expected := "AP with 28:ac:9e:11:48:10 not found"

		helper.AssertErrorMessage(t, err, expected, "NotFoundError should format AP with MAC message correctly")
	})

	t.Run("WLANWithId", func(t *testing.T) {
		err := NotFoundError("WLAN", "5")
		expected := "WLAN with 5 not found"

		helper.AssertErrorMessage(t, err, expected, "NotFoundError should format WLAN with ID message correctly")
	})
}

func TestErrorsServiceUnit_ErrorConstants_Success(t *testing.T) {
	t.Run("ErrorTemplates", func(t *testing.T) {
		operationErr := fmt.Sprintf("failed to %s %s %s", "get", "AP", "data")
		expectedOp := "failed to get AP data"
		helper.AssertStringEquals(t, operationErr, expectedOp, "Operation error template should format correctly")

		simpleErr := fmt.Sprintf("failed to %s", "connect")
		expectedSimple := "failed to connect"
		helper.AssertStringEquals(t, simpleErr, expectedSimple, "Simple error template should format correctly")

		// Test ErrSimpleOperationFailedTemplate
		testErr := errors.New("timeout")
		simpleOpErr := fmt.Errorf(ErrSimpleOperationFailedTemplate, "connect", testErr)
		expectedSimpleOp := "failed to connect: timeout"
		helper.AssertStringEquals(t, simpleOpErr.Error(), expectedSimpleOp,
			"Simple operation failed template should format correctly")

		invalidParam := fmt.Sprintf(ErrInvalidParameterTemplate, "MAC", "invalid")
		expectedInvalid := "invalid MAC: invalid"
		helper.AssertStringEquals(t, invalidParam, expectedInvalid,
			"Invalid parameter template should format correctly")

		notFound := fmt.Sprintf(ErrEntityNotFoundTemplate, "WLAN", "10")
		expectedNotFound := "WLAN with 10 not found"
		helper.AssertStringEquals(t, notFound, expectedNotFound, "Entity not found template should format correctly")

		emptyParam := fmt.Sprintf(ErrEmptyParameterTemplate, "tag")
		expectedEmpty := "tag cannot be empty"
		helper.AssertStringEquals(t, emptyParam, expectedEmpty, "Empty parameter template should format correctly")

		requiredParam := fmt.Sprintf(ErrParameterRequiredTemplate, "MAC address")
		expectedRequired := "MAC address is required"
		helper.AssertStringEquals(t, requiredParam, expectedRequired,
			"Required parameter template should format correctly")
	})

	t.Run("ErrorConstants", func(t *testing.T) {
		// Verify that error constants are not empty
		constants := []string{
			ErrClientNil,
			ErrAtLeastOneTagRequired,
			ErrUnsupportedOperation,
			ErrDataUnavailable,
			ErrConfigurationInvalid,
		}

		for _, constant := range constants {
			if constant == "" {
				helper.AssertStringNotEmpty(t, constant, "Error constant should not be empty")
			}
		}
	})

	t.Run("EntityTypeConstants", func(t *testing.T) {
		// Verify that entity type constants are not empty
		entityTypes := []string{
			EntityTypeAP,
			EntityTypeSite,
			EntityTypeWLAN,
			EntityTypeClient,
			EntityTypeRF,
			EntityTypePolicy,
			EntityTypeTag,
			EntityTypeProfile,
			EntityTypeConfiguration,
			EntityTypeOperationalData,
			EntityTypeGeolocation,
		}

		for _, entityType := range entityTypes {
			if entityType == "" {
				helper.AssertStringNotEmpty(t, entityType, "Entity type constant should not be empty")
			}
		}
	})

	t.Run("ActionConstants", func(t *testing.T) {
		// Verify that action constants are not empty
		actions := []string{
			ActionGet,
			ActionRetrieve,
			ActionSet,
			ActionUpdate,
			ActionCreate,
			ActionDelete,
			ActionEnable,
			ActionDisable,
			ActionAssign,
			ActionValidate,
			ActionResolve,
		}

		for _, action := range actions {
			if action == "" {
				helper.AssertStringNotEmpty(t, action, "Action constant should not be empty")
			}
		}
	})
}
