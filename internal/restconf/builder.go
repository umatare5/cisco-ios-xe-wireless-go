package restconf

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
)

// URL construction constants.
const (
	// URLPathSeparator defines the path separator in URLs.
	URLPathSeparator = "/"
)

// Protocol constants.
const (
	// ProtocolHTTP represents HTTP protocol.
	ProtocolHTTP = "http"
	// ProtocolHTTPS represents HTTPS protocol.
	ProtocolHTTPS = "https"
	// DefaultProtocol is the default protocol for connections.
	DefaultProtocol = ProtocolHTTPS
)

// Builder provides utility functions for building WNC RESTCONF API URLs.
type Builder struct {
	protocol   string
	controller string
}

// NewBuilder creates a new RESTCONF URL builder for the specified protocol and controller.
func NewBuilder(protocol, controller string) *Builder {
	return &Builder{
		protocol:   protocol,
		controller: controller,
	}
}

// BuildDataURL constructs a RESTCONF data URL for the given endpoint path.
func (b *Builder) BuildDataURL(endpointPath string) string {
	if strings.HasPrefix(endpointPath, routes.RESTCONFDataPath) {
		return fmt.Sprintf("%s%s", b.buildBaseURL(), endpointPath)
	}
	normalizedDataPath := b.normalizeEndpointPath(endpointPath)
	return fmt.Sprintf("%s%s%s", b.buildBaseURL(), routes.RESTCONFDataPath, normalizedDataPath)
}

// BuildOperationsURL constructs a RESTCONF operations URL for the given RPC path.
func (b *Builder) BuildOperationsURL(rpcPath string) string {
	if strings.HasPrefix(rpcPath, routes.RESTCONFOperationsPath) {
		return fmt.Sprintf("%s%s", b.buildBaseURL(), rpcPath)
	}
	normalizedOperationsPath := b.normalizeEndpointPath(rpcPath)
	return fmt.Sprintf("%s%s%s", b.buildBaseURL(), routes.RESTCONFOperationsPath, normalizedOperationsPath)
}

// BuildQueryURL constructs URLs for queries with key-value parameters
// Format: endpoint=identifier.
//
// url.PathEscape is the escaper RFC 8040 3.5.3 asks for: it encodes the "," that
// separates composite key values and the "/" that separates path segments, and leaves
// ":" "-" "." alone, so a MAC key in any of the three separator forms is unchanged.
func (b *Builder) BuildQueryURL(endpoint, identifier string) string {
	return fmt.Sprintf("%s=%s", endpoint, url.PathEscape(identifier))
}

// BuildQueryCompositeURL constructs URLs for queries with composite parameters
// Format: endpoint=value1,value2,value3...
//
// The joining commas are the key separator and stay literal; a comma inside a value
// is escaped by url.PathEscape so it cannot pass for one.
func (b *Builder) BuildQueryCompositeURL(endpoint string, values ...interface{}) string {
	strValues := make([]string, len(values))
	for i, v := range values {
		switch val := v.(type) {
		case string:
			strValues[i] = url.PathEscape(val)
		case int:
			strValues[i] = strconv.Itoa(val)
		case int64:
			strValues[i] = strconv.FormatInt(val, 10)
		case float64:
			strValues[i] = strconv.FormatFloat(val, 'f', -1, 64)
		case bool:
			strValues[i] = strconv.FormatBool(val)
		default:
			strValues[i] = url.PathEscape(fmt.Sprintf("%v", val))
		}
	}
	return fmt.Sprintf("%s=%s", endpoint, strings.Join(strValues, ","))
}

// buildBaseURL constructs the base URL for the controller.
func (b *Builder) buildBaseURL() string {
	return fmt.Sprintf("%s://%s", b.protocol, b.controller)
}

// normalizeEndpointPath ensures endpoint path starts with forward slash.
func (b *Builder) normalizeEndpointPath(endpointPath string) string {
	if !strings.HasPrefix(endpointPath, URLPathSeparator) {
		return URLPathSeparator + endpointPath
	}
	return endpointPath
}
