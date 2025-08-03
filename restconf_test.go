package wnc

import (
	"strings"
	"testing"
)

// =============================================================================
// 1. UNIT TESTS (Structure/Type Validation & JSON Serialization/Deserialization)
// =============================================================================

// TestRESTCONFConstants tests RESTCONF and API path constants
func TestRESTCONFConstants(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected string
	}{
		{"RESTCONFPathPrefix", RESTCONFPathPrefix, "/restconf/data"},
		{"RESTCONFModulesPathPrefix", RESTCONFModulesPathPrefix, "/restconf/tailf/modules"},
		{"RESTCONFLibraryQuery", RESTCONFLibraryQuery, "?fields=ietf-yang-library:modules-state/module"},
		{"URLPathSeparator", URLPathSeparator, "/"},
		{"ProtocolHTTP", ProtocolHTTP, "http"},
		{"ProtocolHTTPS", ProtocolHTTPS, "https"},
		{"DefaultProtocol", DefaultProtocol, "https"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.expected {
				t.Errorf("Expected %s to be '%s', got '%s'", tt.name, tt.expected, tt.value)
			}
		})
	}
}

// TestYANGModelConstants tests YANG model validation constants
func TestYANGModelConstants(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected string
	}{
		{"YANGModelPrefix", YANGModelPrefix, "Cisco-IOS-XE-wireless-"},
		{"YANGModelOperSuffix", YANGModelOperSuffix, "-oper"},
		{"YANGModelCfgSuffix", YANGModelCfgSuffix, "-cfg"},
		{"CiscoIOSXEWirelessPrefix", CiscoIOSXEWirelessPrefix, "Cisco-IOS-XE-wireless-"},
		{"OperDataSuffix", OperDataSuffix, "-oper-data"},
		{"CfgDataSuffix", CfgDataSuffix, "-cfg-data"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.expected {
				t.Errorf("Expected %s to be '%s', got '%s'", tt.name, tt.expected, tt.value)
			}
		})
	}
}

// =============================================================================
// 2. TABLE-DRIVEN TEST PATTERNS
// =============================================================================

// TestRESTCONFURLBuilder tests URL builder creation and basic functionality
func TestRESTCONFURLBuilder(t *testing.T) {
	tests := []struct {
		name       string
		protocol   string
		controller string
	}{
		{"HTTPSBuilder", "https", "wnc.example.com"},
		{"HTTPBuilder", "http", "192.168.1.100"},
		{"LocalBuilder", "https", "localhost"},
		{"PortBuilder", "https", "wnc.local:8443"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := NewRESTCONFURLBuilder(tt.protocol, tt.controller)

			if builder == nil {
				t.Fatal("Expected builder to be created")
			}

			if builder.protocol != tt.protocol {
				t.Errorf("Expected protocol '%s', got '%s'", tt.protocol, builder.protocol)
			}

			if builder.controller != tt.controller {
				t.Errorf("Expected controller '%s', got '%s'", tt.controller, builder.controller)
			}
		})
	}
}

// TestBuildBaseURL tests base URL construction
func TestBuildBaseURL(t *testing.T) {
	tests := []struct {
		name       string
		protocol   string
		controller string
		expected   string
	}{
		{"HTTPSBase", "https", "wnc.example.com", "https://wnc.example.com"},
		{"HTTPBase", "http", "192.168.1.100", "http://192.168.1.100"},
		{"WithPort", "https", "wnc.local:8443", "https://wnc.local:8443"},
		{"Localhost", "http", "localhost", "http://localhost"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := NewRESTCONFURLBuilder(tt.protocol, tt.controller)
			result := builder.BuildBaseURL()

			if result != tt.expected {
				t.Errorf("Expected '%s', got '%s'", tt.expected, result)
			}
		})
	}
}

// TestBuildRESTCONFURL tests RESTCONF URL construction
func TestBuildRESTCONFURL(t *testing.T) {
	tests := []struct {
		name         string
		protocol     string
		controller   string
		endpointPath string
		expected     string
	}{
		{"SimpleEndpoint", "https", "wnc.example.com", "/ap-oper", "https://wnc.example.com/restconf/data/ap-oper"},
		{"NoLeadingSlash", "https", "wnc.example.com", "ap-oper", "https://wnc.example.com/restconf/data/ap-oper"},
		{"ComplexEndpoint", "https", "192.168.1.100", "/client-oper-data", "https://192.168.1.100/restconf/data/client-oper-data"},
		{"EmptyEndpoint", "https", "localhost", "", "https://localhost/restconf/data/"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := NewRESTCONFURLBuilder(tt.protocol, tt.controller)
			result := builder.BuildRESTCONFURL(tt.endpointPath)

			if result != tt.expected {
				t.Errorf("Expected '%s', got '%s'", tt.expected, result)
			}
		})
	}
}

// TestBuildYANGLibraryURL tests YANG library URL construction
func TestBuildYANGLibraryURL(t *testing.T) {
	tests := []struct {
		name       string
		protocol   string
		controller string
		expected   string
	}{
		{"HTTPSLibrary", "https", "wnc.example.com", "https://wnc.example.com/restconf/data?fields=ietf-yang-library:modules-state/module"},
		{"HTTPLibrary", "http", "192.168.1.100", "http://192.168.1.100/restconf/data?fields=ietf-yang-library:modules-state/module"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := NewRESTCONFURLBuilder(tt.protocol, tt.controller)
			result := builder.BuildYANGLibraryURL()

			if result != tt.expected {
				t.Errorf("Expected '%s', got '%s'", tt.expected, result)
			}
		})
	}
}

// TestBuildYANGModuleURL tests YANG module URL construction
func TestBuildYANGModuleURL(t *testing.T) {
	tests := []struct {
		name       string
		protocol   string
		controller string
		yangModel  string
		revision   string
		expected   string
	}{
		{"SimpleModule", "https", "wnc.example.com", "Cisco-IOS-XE-wireless-ap-oper", "2024-01-01", "https://wnc.example.com/restconf/tailf/modules/Cisco-IOS-XE-wireless-ap-oper/2024-01-01"},
		{"CfgModule", "https", "192.168.1.100", "Cisco-IOS-XE-wireless-general-cfg", "2023-11-01", "https://192.168.1.100/restconf/tailf/modules/Cisco-IOS-XE-wireless-general-cfg/2023-11-01"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := NewRESTCONFURLBuilder(tt.protocol, tt.controller)
			result := builder.BuildYANGModuleURL(tt.yangModel, tt.revision)

			if result != tt.expected {
				t.Errorf("Expected '%s', got '%s'", tt.expected, result)
			}
		})
	}
}

// =============================================================================
// 3. FAIL-FAST ERROR DETECTION TESTS
// =============================================================================

// TestIsValidProtocol tests protocol validation
func TestIsValidProtocol(t *testing.T) {
	tests := []struct {
		name     string
		protocol string
		expected bool
	}{
		{"ValidHTTPS", "https", true},
		{"ValidHTTP", "http", true},
		{"InvalidFTP", "ftp", false},
		{"InvalidEmpty", "", false},
		{"InvalidCase", "HTTP", false},
		{"InvalidCase2", "HTTPS", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidProtocol(tt.protocol)
			if result != tt.expected {
				t.Errorf("Expected IsValidProtocol('%s') to be %v, got %v", tt.protocol, tt.expected, result)
			}
		})
	}
}

// TestIsValidYANGModel tests YANG model validation
func TestIsValidYANGModel(t *testing.T) {
	tests := []struct {
		name      string
		yangModel string
		expected  bool
	}{
		{"ValidOperModel", "Cisco-IOS-XE-wireless-ap-oper", true},
		{"ValidCfgModel", "Cisco-IOS-XE-wireless-general-cfg", true},
		{"InvalidPrefix", "cisco-ios-xe-wireless-ap-oper", false},
		{"InvalidSuffix", "Cisco-IOS-XE-wireless-ap", false},
		{"EmptyModel", "", false},
		{"WrongPrefix", "Other-IOS-XE-wireless-ap-oper", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidYANGModel(tt.yangModel)
			if result != tt.expected {
				t.Errorf("Expected IsValidYANGModel('%s') to be %v, got %v", tt.yangModel, tt.expected, result)
			}
		})
	}
}

// TestIsValidRevision tests revision validation
func TestIsValidRevision(t *testing.T) {
	tests := []struct {
		name     string
		revision string
		expected bool
	}{
		{"ValidRevision", "2024-01-01", true},
		{"ValidRevision2", "2023-12-31", true},
		{"InvalidLength", "2024-1-1", false},
		{"InvalidFormat", "2024/01/01", false},
		{"InvalidChars", "2024-ab-01", false},
		{"EmptyRevision", "", false},
		{"TooLong", "2024-01-01-extra", false},
		{"NoSeparators", "20240101", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidRevision(tt.revision)
			if result != tt.expected {
				t.Errorf("Expected IsValidRevision('%s') to be %v, got %v", tt.revision, tt.expected, result)
			}
		})
	}
}

// TestNormalizeEndpointPath tests endpoint path normalization
func TestNormalizeEndpointPath(t *testing.T) {
	tests := []struct {
		name         string
		endpointPath string
		expected     string
	}{
		{"WithLeadingSlash", "/ap-oper", "/ap-oper"},
		{"WithoutLeadingSlash", "ap-oper", "/ap-oper"},
		{"EmptyPath", "", "/"},
		{"ComplexPath", "client-oper-data/stats", "/client-oper-data/stats"},
		{"QueryString", "ap-oper?fields=all", "/ap-oper?fields=all"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := NewRESTCONFURLBuilder("https", "test.com")
			result := builder.normalizeEndpointPath(tt.endpointPath)

			if result != tt.expected {
				t.Errorf("Expected normalizeEndpointPath('%s') to be '%s', got '%s'", tt.endpointPath, tt.expected, result)
			}
		})
	}
}

// TestBuildEndpointURL tests endpoint URL delegation
func TestBuildEndpointURL(t *testing.T) {
	builder := NewRESTCONFURLBuilder("https", "wnc.example.com")
	endpoint := "/ap-oper"

	result1 := builder.BuildEndpointURL(endpoint)
	result2 := builder.BuildRESTCONFURL(endpoint)

	if result1 != result2 {
		t.Errorf("Expected BuildEndpointURL to delegate to BuildRESTCONFURL, got different results: '%s' vs '%s'", result1, result2)
	}
}

// TestConstantConsistency tests that constants are consistent
func TestConstantConsistency(t *testing.T) {
	// Test that CiscoIOSXEWirelessPrefix equals YANGModelPrefix
	if CiscoIOSXEWirelessPrefix != YANGModelPrefix {
		t.Errorf("Expected CiscoIOSXEWirelessPrefix (%s) to equal YANGModelPrefix (%s)",
			CiscoIOSXEWirelessPrefix, YANGModelPrefix)
	}

	// Test that OperDataSuffix contains YANGModelOperSuffix
	if !strings.Contains(OperDataSuffix, YANGModelOperSuffix) {
		t.Errorf("Expected OperDataSuffix (%s) to contain YANGModelOperSuffix (%s)",
			OperDataSuffix, YANGModelOperSuffix)
	}

	// Test that CfgDataSuffix contains YANGModelCfgSuffix
	if !strings.Contains(CfgDataSuffix, YANGModelCfgSuffix) {
		t.Errorf("Expected CfgDataSuffix (%s) to contain YANGModelCfgSuffix (%s)",
			CfgDataSuffix, YANGModelCfgSuffix)
	}

	// Test that DefaultProtocol is HTTPS
	if DefaultProtocol != ProtocolHTTPS {
		t.Errorf("Expected DefaultProtocol (%s) to be HTTPS (%s)", DefaultProtocol, ProtocolHTTPS)
	}
}
