// Package restconf provides RESTCONF URL building utilities for the WNC client.
package restconf

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// RESTCONF and API path constants
const (
	// RESTCONFPathPrefix is the base path for all RESTCONF API endpoints
	RESTCONFPathPrefix = "/restconf/data"
	// RESTCONFOperationsPathPrefix is the base path for RESTCONF RPC operations
	RESTCONFOperationsPathPrefix = "/restconf/operations"
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
	// YANGModelPrefix is the standard prefix for wireless-related YANG models
	YANGModelPrefix = "Cisco-IOS-XE-wireless-"
	// YANGModelOperSuffix is the suffix for operational YANG models
	YANGModelOperSuffix = "-oper"
	// YANGModelCfgSuffix is the suffix for configuration YANG models
	YANGModelCfgSuffix = "-cfg"
	// OperDataSuffix is the common suffix for operational data endpoints
	OperDataSuffix = YANGModelOperSuffix + "-data"
	// CfgDataSuffix is the common suffix for configuration data endpoints
	CfgDataSuffix = YANGModelCfgSuffix + "-data"

	// Compatibility aliases for existing code
	RestconfYANGModelPrefix     = YANGModelPrefix
	RestconfYANGModelOperSuffix = YANGModelOperSuffix
	RestconfYANGModelCfgSuffix  = YANGModelCfgSuffix
	CiscoIOSXEWirelessPrefix    = YANGModelPrefix
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

// BuildRPCURL constructs a RESTCONF RPC operations URL for the given RPC path
func (b *Builder) BuildRPCURL(rpcPath string) string {
	// RPC paths already contain /restconf/operations prefix, so we build base URL + path
	if strings.HasPrefix(rpcPath, RESTCONFOperationsPathPrefix) {
		return fmt.Sprintf("%s%s", b.BuildBaseURL(), rpcPath)
	}
	// If path doesn't include operations prefix, add it
	normalizedRPCPath := b.normalizeEndpointPath(rpcPath)
	return fmt.Sprintf("%s%s%s", b.BuildBaseURL(), RESTCONFOperationsPathPrefix, normalizedRPCPath)
}

// normalizeEndpointPath ensures endpoint path starts with forward slash
func (b *Builder) normalizeEndpointPath(endpointPath string) string {
	if !strings.HasPrefix(endpointPath, URLPathSeparator) {
		return URLPathSeparator + endpointPath
	}
	return endpointPath
}

// NOTE: BuildYANGLibraryURL, BuildYANGModuleURL, and BuildEndpointURL were removed
// as they were only used in tests and provided no additional value over existing methods.

// BuildFieldsURL constructs URL with fields parameter for selective data retrieval
func (b *Builder) BuildFieldsURL(baseURL, fields string) string {
	if fields == "" {
		return baseURL
	}
	return baseURL + "?fields=" + url.QueryEscape(fields)
}

// BuildPathQueryURL constructs URLs with path and key-value parameters
// Format: endpoint/key=value
func (b *Builder) BuildPathQueryURL(endpoint, key, value string) string {
	return fmt.Sprintf("%s/%s=%s", endpoint, key, value)
}

// BuildQueryURL constructs URLs for queries with key-value parameters
// Format: endpoint=identifier
func (b *Builder) BuildQueryURL(endpoint, identifier string) string {
	return fmt.Sprintf("%s=%s", endpoint, identifier)
}

// convertToString converts interface{} values to strings for URL construction
func convertToString(v interface{}) string {
	switch val := v.(type) {
	case string:
		return val
	case int:
		return strconv.Itoa(val)
	case int64:
		return strconv.FormatInt(val, 10)
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(val)
	default:
		return fmt.Sprintf("%v", val)
	}
}

// BuildCompositeKeyURL constructs URLs with composite key parameters
// Format: endpoint/key=value1,value2,value3...
func (b *Builder) BuildCompositeKeyURL(endpoint, key string, values ...interface{}) string {
	strValues := make([]string, len(values))
	for i, v := range values {
		strValues[i] = convertToString(v)
	}
	return fmt.Sprintf("%s/%s=%s", endpoint, key, strings.Join(strValues, ","))
}

// BuildQueryCompositeURL constructs URLs for queries with composite parameters
// Format: endpoint=value1,value2,value3...
func (b *Builder) BuildQueryCompositeURL(endpoint string, values ...interface{}) string {
	strValues := make([]string, len(values))
	for i, v := range values {
		strValues[i] = convertToString(v)
	}
	return fmt.Sprintf("%s=%s", endpoint, strings.Join(strValues, ","))
}

// BuildFieldsURLMultiple constructs URL with multiple fields parameter for selective data retrieval
func (b *Builder) BuildFieldsURLMultiple(basePath string, fields []string) string {
	if len(fields) == 0 {
		return basePath
	}

	params := url.Values{}
	params.Add("fields", strings.Join(fields, ","))
	return fmt.Sprintf("%s?%s", basePath, params.Encode())
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

	// Check date format and components in one validation pass
	return revisionString[4] == '-' && revisionString[7] == '-' &&
		isDigits(revisionString[0:4]) && isDigits(revisionString[5:7]) && isDigits(revisionString[8:10])
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
