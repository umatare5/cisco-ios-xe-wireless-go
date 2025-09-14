package restconf

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil"
)

func TestRESTCONFBuilderUnit_NewBuilder_Success(t *testing.T) {
	protocol := "https"
	controller := "192.168.1.100"

	builder := NewBuilder(protocol, controller)

	testutil.AssertNotNil(t, builder, "NewBuilder result")

	testutil.AssertStringEquals(t, builder.protocol, protocol, "protocol")
	testutil.AssertStringEquals(t, builder.controller, controller, "controller")
}

func TestRESTCONFBuilderUnit_buildBaseURL_Success(t *testing.T) {
	testCases := []struct {
		name       string
		protocol   string
		controller string
		expected   string
	}{
		{"HTTPS with IP", "https", "192.168.1.100", "https://192.168.1.100"},
		{"HTTP with hostname", "http", "core.example.com", "http://core.example.com"},
		{"HTTPS with port", "https", "core.example.com:8443", "https://core.example.com:8443"},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			builder := NewBuilder(tt.protocol, tt.controller)
			baseURL := builder.buildBaseURL()
			testutil.AssertStringEquals(t, baseURL, tt.expected, "buildBaseURL()")
		})
	}
}

func TestRESTCONFBuilderUnit_BuildDataURL_Success(t *testing.T) {
	builder := NewBuilder("https", "192.168.1.100")

	testCases := []struct {
		name         string
		endpointPath string
		expected     string
	}{
		{
			"endpoint with leading slash",
			"/Cisco-IOS-XE-wireless-afc-oper:afc-oper-data",
			"https://192.168.1.100/restconf/data/Cisco-IOS-XE-wireless-afc-oper:afc-oper-data",
		},
		{
			"endpoint without leading slash",
			"Cisco-IOS-XE-wireless-ap-oper:ap-oper-data",
			"https://192.168.1.100/restconf/data/Cisco-IOS-XE-wireless-ap-oper:ap-oper-data",
		},
		{
			"empty endpoint",
			"",
			"https://192.168.1.100/restconf/data/",
		},
		{
			"complex endpoint",
			"/Cisco-IOS-XE-wireless-general-oper:general-oper-data/interfaces",
			"https://192.168.1.100/restconf/data/Cisco-IOS-XE-wireless-general-oper:general-oper-data/interfaces",
		},
		{
			"endpoint already with restconf data path",
			"/restconf/data/Cisco-IOS-XE-wireless-ap-oper:ap-oper-data",
			"https://192.168.1.100/restconf/data/Cisco-IOS-XE-wireless-ap-oper:ap-oper-data",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			url := builder.BuildDataURL(tt.endpointPath)
			testutil.AssertStringEquals(t, url, tt.expected, "BuildDataURL()")
		})
	}
}

func TestRESTCONFBuilderUnit_NormalizeEndpointPath_Success(t *testing.T) {
	builder := NewBuilder("https", "192.168.1.100")

	testCases := []struct {
		name     string
		path     string
		expected string
	}{
		{"path with leading slash", "/endpoint", "/endpoint"},
		{"path without leading slash", "endpoint", "/endpoint"},
		{"empty path", "", "/"},
		{"path with only slash", "/", "/"},
		{"complex path", "path/to/endpoint", "/path/to/endpoint"},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			normalized := builder.normalizeEndpointPath(tt.path)
			testutil.AssertStringEquals(t, normalized, tt.expected, "normalizeEndpointPath()")
		})
	}
}

func TestRESTCONFBuilderUnit_isValidProtocol_Success(t *testing.T) {
	testCases := []struct {
		protocol string
		expected bool
	}{
		{"http", true},
		{"https", true},
		{"HTTP", false},
		{"HTTPS", false},
		{"ftp", false},
		{"", false},
		{"tcp", false},
	}

	for _, tt := range testCases {
		t.Run(tt.protocol, func(t *testing.T) {
			result := isValidProtocol(tt.protocol)
			testutil.AssertBoolEquals(t, result, tt.expected, "isValidProtocol()")
		})
	}
}

func TestRESTCONFBuilderUnit_Constants_Success(t *testing.T) {
	// Test RESTCONF constants
	testutil.AssertStringEquals(t, routes.RESTCONFDataPath, "/restconf/data", "RESTCONFDataPath")

	// Test protocol constants
	testutil.AssertStringEquals(t, ProtocolHTTP, "http", "ProtocolHTTP")

	testutil.AssertStringEquals(t, ProtocolHTTPS, "https", "ProtocolHTTPS")

	testutil.AssertStringEquals(t, DefaultProtocol, ProtocolHTTPS, "DefaultProtocol")
}

// TestBuildOperationsURL tests RPC URL construction.
func TestRESTCONFBuilderUnit_BuildOperationsURL_Success(t *testing.T) {
	builder := NewBuilder("https", "192.168.1.1")

	tests := []struct {
		name     string
		rpcPath  string
		expected string
	}{
		{
			name:     "RPC path with operations prefix",
			rpcPath:  "/restconf/operations/cisco-wireless:ap-join",
			expected: "https://192.168.1.1/restconf/operations/cisco-wireless:ap-join",
		},
		{
			name:     "RPC path without operations prefix",
			rpcPath:  "cisco-wireless:ap-join",
			expected: "https://192.168.1.1/restconf/operations/cisco-wireless:ap-join",
		},
		{
			name:     "RPC path without leading slash",
			rpcPath:  "cisco-wireless:ap-reload",
			expected: "https://192.168.1.1/restconf/operations/cisco-wireless:ap-reload",
		},
		{
			name:     "RPC path with leading slash",
			rpcPath:  "/cisco-wireless:ap-reload",
			expected: "https://192.168.1.1/restconf/operations/cisco-wireless:ap-reload",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := builder.BuildOperationsURL(tt.rpcPath)
			testutil.AssertStringEquals(t, result, tt.expected, "BuildOperationsURL()")
		})
	}
}

// TestBuildQueryURL tests query URL construction.
func TestRESTCONFBuilderUnit_BuildQueryURL_Success(t *testing.T) {
	builder := NewBuilder("https", "192.168.1.1")

	tests := []struct {
		name       string
		endpoint   string
		identifier string
		expected   string
	}{
		{
			name:       "Simple query",
			endpoint:   "ap-name",
			identifier: "TEST-AP",
			expected:   "ap-name=TEST-AP",
		},
		{
			name:       "MAC query",
			endpoint:   "wtp-mac",
			identifier: "aa:bb:cc:dd:ee:ff",
			expected:   "wtp-mac=aa:bb:cc:dd:ee:ff",
		},
		{
			name:       "Empty identifier",
			endpoint:   "test",
			identifier: "",
			expected:   "test=",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := builder.BuildQueryURL(tt.endpoint, tt.identifier)
			testutil.AssertStringEquals(t, result, tt.expected, "BuildQueryURL()")
		})
	}
}

// TestBuildQueryCompositeURL tests composite query URL construction.
func TestRESTCONFBuilderUnit_BuildQueryCompositeURL_Success(t *testing.T) {
	builder := NewBuilder("https", "192.168.1.1")

	tests := []struct {
		name     string
		endpoint string
		values   []interface{}
		expected string
	}{
		{
			name:     "String values",
			endpoint: "composite-key",
			values:   []interface{}{"mac", "slot"},
			expected: "composite-key=mac,slot",
		},
		{
			name:     "Mixed types",
			endpoint: "query",
			values:   []interface{}{"aa:bb:cc:dd:ee:ff", 0, true},
			expected: "query=aa:bb:cc:dd:ee:ff,0,true",
		},
		{
			name:     "Single value",
			endpoint: "single",
			values:   []interface{}{"value"},
			expected: "single=value",
		},
		{
			name:     "Empty values",
			endpoint: "empty",
			values:   []interface{}{},
			expected: "empty=",
		},
		{
			name:     "Int64 values",
			endpoint: "int64-test",
			values:   []interface{}{int64(123456789)},
			expected: "int64-test=123456789",
		},
		{
			name:     "Float64 values",
			endpoint: "float64-test",
			values:   []interface{}{float64(123.456)},
			expected: "float64-test=123.456",
		},
		{
			name:     "Bool values",
			endpoint: "bool-test",
			values:   []interface{}{true, false},
			expected: "bool-test=true,false",
		},
		{
			name:     "Custom type fallback",
			endpoint: "custom-test",
			values:   []interface{}{struct{ Name string }{"test"}},
			expected: "custom-test={test}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := builder.BuildQueryCompositeURL(tt.endpoint, tt.values...)
			testutil.AssertStringEquals(t, result, tt.expected, "BuildQueryCompositeURL()")
		})
	}
}
