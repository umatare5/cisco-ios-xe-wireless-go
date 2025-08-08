// Package restconf provides RESTCONF URL building utilities for the WNC client.
package restconf

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
	// RestconfYANGModelPrefix is the expected prefix for Cisco wireless YANG models
	RestconfYANGModelPrefix = "Cisco-IOS-XE-wireless-"
	// RestconfYANGModelOperSuffix is the suffix for operational YANG models
	RestconfYANGModelOperSuffix = "-oper"
	// RestconfYANGModelCfgSuffix is the suffix for configuration YANG models
	RestconfYANGModelCfgSuffix = "-cfg"
)

// Common YANG model patterns
const (
	// CiscoIOSXEWirelessPrefix is the common prefix for all wireless YANG models
	CiscoIOSXEWirelessPrefix = RestconfYANGModelPrefix
	// OperDataSuffix is the common suffix for operational data endpoints
	OperDataSuffix = RestconfYANGModelOperSuffix + "-data"
	// CfgDataSuffix is the common suffix for configuration data endpoints
	CfgDataSuffix = RestconfYANGModelCfgSuffix + "-data"
)

// Builder provides utility functions for building WNC RESTCONF API URLs
type Builder struct {
	protocol   string
	controller string
}

// NewBuilder creates a new RESTCONF URL builder for the specified protocol and controller
func NewBuilder(protocol, controller string) *Builder {
	return &Builder{
		protocol:   protocol,
		controller: controller,
	}
}

// BuildBaseURL constructs the base URL for the controller
func (b *Builder) BuildBaseURL() string {
	return fmt.Sprintf("%s://%s", b.protocol, b.controller)
}

// BuildRESTCONFURL constructs a RESTCONF data URL for the given endpoint path
func (b *Builder) BuildRESTCONFURL(endpointPath string) string {
	normalizedEndpointPath := b.normalizeEndpointPath(endpointPath)
	return fmt.Sprintf("%s%s%s", b.BuildBaseURL(), RESTCONFPathPrefix, normalizedEndpointPath)
}

// normalizeEndpointPath ensures endpoint path starts with forward slash
func (b *Builder) normalizeEndpointPath(endpointPath string) string {
	if !strings.HasPrefix(endpointPath, URLPathSeparator) {
		return URLPathSeparator + endpointPath
	}
	return endpointPath
}

// BuildYANGLibraryURL constructs the URL for querying YANG library modules
func (b *Builder) BuildYANGLibraryURL() string {
	return fmt.Sprintf("%s%s%s", b.BuildBaseURL(), RESTCONFPathPrefix, RESTCONFLibraryQuery)
}

// BuildYANGModuleURL constructs the URL for getting details of a specific YANG module
func (b *Builder) BuildYANGModuleURL(yangModel, revision string) string {
	return fmt.Sprintf("%s%s/%s/%s", b.BuildBaseURL(), RESTCONFModulesPathPrefix, yangModel, revision)
}

// BuildEndpointURL is a convenience method that delegates to BuildRESTCONFURL
func (b *Builder) BuildEndpointURL(endpoint string) string {
	return b.BuildRESTCONFURL(endpoint)
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
	return strings.HasPrefix(yangModelName, RestconfYANGModelPrefix)
}

// hasValidYANGSuffix checks if model name has valid operational or configuration suffix
func hasValidYANGSuffix(yangModelName string) bool {
	return strings.HasSuffix(yangModelName, RestconfYANGModelOperSuffix) ||
		strings.HasSuffix(yangModelName, RestconfYANGModelCfgSuffix)
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
