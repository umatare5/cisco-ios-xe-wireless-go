// Package wnc provides RESTCONF API configuration constants and URL building utilities for the Cisco Wireless Network Controller.
package wnc

import (
	"fmt"
	"strings"
)

// RESTCONF and API path constants
const (
	// RESTCONFPathPrefix is the base path for all RESTCONF API endpoints
	RESTCONFPathPrefix = "/restconf/data"

	// RESTCONFModulesPathPrefix is the base path for YANG module queries
	RESTCONFModulesPathPrefix = "/restconf/tailf/modules"

	// RESTCONFLibraryQuery is the query string for YANG library modules
	RESTCONFLibraryQuery = "?fields=ietf-yang-library:modules-state/module"
)

// URL construction constants
const (
	// URLPathSeparator defines the path separator in URLs
	URLPathSeparator = "/"
)

// Protocol constants
const (
	// ProtocolHTTP represents HTTP protocol
	ProtocolHTTP = "http"

	// ProtocolHTTPS represents HTTPS protocol
	ProtocolHTTPS = "https"

	// DefaultProtocol is the default protocol for connections
	DefaultProtocol = ProtocolHTTPS
)

// YANG model validation constants
const (
	// YANGModelPrefix is the expected prefix for Cisco wireless YANG models
	YANGModelPrefix = "Cisco-IOS-XE-wireless-"

	// YANGModelOperSuffix is the suffix for operational YANG models
	YANGModelOperSuffix = "-oper"

	// YANGModelCfgSuffix is the suffix for configuration YANG models
	YANGModelCfgSuffix = "-cfg"
)

// Common YANG model patterns
const (
	// CiscoIOSXEWirelessPrefix is the common prefix for all wireless YANG models
	CiscoIOSXEWirelessPrefix = YANGModelPrefix

	// OperDataSuffix is the common suffix for operational data endpoints
	OperDataSuffix = YANGModelOperSuffix + "-data"

	// CfgDataSuffix is the common suffix for configuration data endpoints
	CfgDataSuffix = YANGModelCfgSuffix + "-data"
)

// RESTCONFURLBuilder provides utility functions for building WNC RESTCONF API URLs
type RESTCONFURLBuilder struct {
	protocol   string
	controller string
}

// NewRESTCONFURLBuilder creates a new RESTCONF URL builder for the specified controller
func NewRESTCONFURLBuilder(protocol, controller string) *RESTCONFURLBuilder {
	return &RESTCONFURLBuilder{
		protocol:   protocol,
		controller: controller,
	}
}

// BuildBaseURL constructs the base URL for the controller
func (u *RESTCONFURLBuilder) BuildBaseURL() string {
	return fmt.Sprintf("%s://%s", u.protocol, u.controller)
}

// BuildRESTCONFURL constructs a RESTCONF data URL for the given endpoint path
func (u *RESTCONFURLBuilder) BuildRESTCONFURL(endpointPath string) string {
	normalizedEndpointPath := u.normalizeEndpointPath(endpointPath)
	return fmt.Sprintf("%s%s%s", u.BuildBaseURL(), RESTCONFPathPrefix, normalizedEndpointPath)
}

// normalizeEndpointPath ensures endpoint path starts with forward slash
func (u *RESTCONFURLBuilder) normalizeEndpointPath(endpointPath string) string {
	if !strings.HasPrefix(endpointPath, URLPathSeparator) {
		return URLPathSeparator + endpointPath
	}
	return endpointPath
}

// BuildYANGLibraryURL constructs the URL for querying YANG library modules
func (u *RESTCONFURLBuilder) BuildYANGLibraryURL() string {
	return fmt.Sprintf("%s%s%s", u.BuildBaseURL(), RESTCONFPathPrefix, RESTCONFLibraryQuery)
}

// BuildYANGModuleURL constructs the URL for getting details of a specific YANG module
func (u *RESTCONFURLBuilder) BuildYANGModuleURL(yangModel, revision string) string {
	return fmt.Sprintf("%s%s/%s/%s", u.BuildBaseURL(), RESTCONFModulesPathPrefix, yangModel, revision)
}

// BuildEndpointURL is a convenience method that delegates to BuildRESTCONFURL
func (u *RESTCONFURLBuilder) BuildEndpointURL(endpoint string) string {
	return u.BuildRESTCONFURL(endpoint)
}

// Validation functions for URL components

// IsValidProtocol checks if the protocol is supported
func IsValidProtocol(protocol string) bool {
	return protocol == ProtocolHTTP || protocol == ProtocolHTTPS
}

// IsValidYANGModel checks if the YANG model name follows Cisco wireless conventions
func IsValidYANGModel(yangModelName string) bool {
	return hasValidYANGPrefix(yangModelName) && hasValidYANGSuffix(yangModelName)
}

// hasValidYANGPrefix checks if model name has required Cisco prefix
func hasValidYANGPrefix(yangModelName string) bool {
	return strings.HasPrefix(yangModelName, YANGModelPrefix)
}

// hasValidYANGSuffix checks if model name has valid operational or configuration suffix
func hasValidYANGSuffix(yangModelName string) bool {
	return strings.HasSuffix(yangModelName, YANGModelOperSuffix) ||
		strings.HasSuffix(yangModelName, YANGModelCfgSuffix)
}

// IsValidRevision checks if the revision follows YYYY-MM-DD format
func IsValidRevision(revisionString string) bool {
	// Early return for invalid length
	if len(revisionString) != 10 {
		return false
	}

	return hasValidDateFormat(revisionString) && hasValidDateComponents(revisionString)
}

// hasValidDateFormat checks for correct date separator positions
func hasValidDateFormat(revisionString string) bool {
	return revisionString[4] == '-' && revisionString[7] == '-'
}

// hasValidDateComponents checks if year, month, and day components are numeric
func hasValidDateComponents(revisionString string) bool {
	yearComponent := revisionString[0:4]
	monthComponent := revisionString[5:7]
	dayComponent := revisionString[8:10]

	return isDigits(yearComponent) && isDigits(monthComponent) && isDigits(dayComponent)
}

// isDigits checks if a string contains only digits
func isDigits(digitString string) bool {
	for _, digitRune := range digitString {
		if digitRune < '0' || digitRune > '9' {
			return false
		}
	}
	return true
}
