package restconf

import (
	"testing"
)

func TestNewBuilder(t *testing.T) {
	protocol := "https"
	controller := "192.168.1.100"

	builder := NewBuilder(protocol, controller)

	if builder == nil {
		t.Fatal("NewBuilder returned nil")
	}

	if builder.protocol != protocol {
		t.Errorf("protocol = %q, want %q", builder.protocol, protocol)
	}

	if builder.controller != controller {
		t.Errorf("controller = %q, want %q", builder.controller, controller)
	}
}

func TestBuildBaseURL(t *testing.T) {
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
			baseURL := builder.BuildBaseURL()

			if baseURL != tt.expected {
				t.Errorf("BuildBaseURL() = %q, want %q", baseURL, tt.expected)
			}
		})
	}
}

func TestBuildRESTCONFURL(t *testing.T) {
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
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			url := builder.BuildRESTCONFURL(tt.endpointPath)

			if url != tt.expected {
				t.Errorf("BuildRESTCONFURL(%q) = %q, want %q", tt.endpointPath, url, tt.expected)
			}
		})
	}
}

func TestNormalizeEndpointPath(t *testing.T) {
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

			if normalized != tt.expected {
				t.Errorf("normalizeEndpointPath(%q) = %q, want %q", tt.path, normalized, tt.expected)
			}
		})
	}
}

func TestBuildYANGLibraryURL(t *testing.T) {
	builder := NewBuilder("https", "core.example.com")
	url := builder.BuildYANGLibraryURL()

	expected := "https://core.example.com/restconf/data?fields=ietf-yang-library:modules-state/module"
	if url != expected {
		t.Errorf("BuildYANGLibraryURL() = %q, want %q", url, expected)
	}
}

func TestBuildYANGModuleURL(t *testing.T) {
	builder := NewBuilder("https", "core.example.com")

	testCases := []struct {
		name      string
		yangModel string
		revision  string
		expected  string
	}{
		{
			"standard module",
			"Cisco-IOS-XE-wireless-afc-oper",
			"2021-07-01",
			"https://core.example.com/restconf/tailf/modules/Cisco-IOS-XE-wireless-afc-oper/2021-07-01",
		},
		{
			"config module",
			"Cisco-IOS-XE-wireless-ap-cfg",
			"2022-03-15",
			"https://core.example.com/restconf/tailf/modules/Cisco-IOS-XE-wireless-ap-cfg/2022-03-15",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			url := builder.BuildYANGModuleURL(tt.yangModel, tt.revision)

			if url != tt.expected {
				t.Errorf("BuildYANGModuleURL(%q, %q) = %q, want %q", tt.yangModel, tt.revision, url, tt.expected)
			}
		})
	}
}

func TestBuildEndpointURL(t *testing.T) {
	builder := NewBuilder("https", "core.example.com")
	endpoint := "/Cisco-IOS-XE-wireless-general-oper:general-oper-data"

	endpointURL := builder.BuildEndpointURL(endpoint)
	restconfURL := builder.BuildRESTCONFURL(endpoint)

	if endpointURL != restconfURL {
		t.Errorf("BuildEndpointURL() = %q, BuildRESTCONFURL() = %q, should be equal", endpointURL, restconfURL)
	}
}

func TestIsValidProtocol(t *testing.T) {
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
			result := IsValidProtocol(tt.protocol)

			if result != tt.expected {
				t.Errorf("IsValidProtocol(%q) = %v, want %v", tt.protocol, result, tt.expected)
			}
		})
	}
}

func TestIsValidYANGModel(t *testing.T) {
	testCases := []struct {
		name      string
		yangModel string
		expected  bool
	}{
		{"valid operational model", "Cisco-IOS-XE-wireless-afc-oper", true},
		{"valid config model", "Cisco-IOS-XE-wireless-ap-cfg", true},
		{"invalid prefix", "Invalid-wireless-afc-oper", false},
		{"invalid suffix", "Cisco-IOS-XE-wireless-afc", false},
		{"empty string", "", false},
		{"just prefix", "Cisco-IOS-XE-wireless-", false},
		{"just suffix", "-oper", false},
		{"wrong prefix case", "cisco-ios-xe-wireless-afc-oper", false},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidYANGModel(tt.yangModel)

			if result != tt.expected {
				t.Errorf("IsValidYANGModel(%q) = %v, want %v", tt.yangModel, result, tt.expected)
			}
		})
	}
}

func TestHasValidYANGPrefix(t *testing.T) {
	testCases := []struct {
		name     string
		model    string
		expected bool
	}{
		{"valid prefix", "Cisco-IOS-XE-wireless-afc-oper", true},
		{"invalid prefix", "Invalid-IOS-XE-wireless-afc-oper", false},
		{"empty string", "", false},
		{"just prefix", "Cisco-IOS-XE-wireless-", true},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := hasValidYANGPrefix(tt.model)

			if result != tt.expected {
				t.Errorf("hasValidYANGPrefix(%q) = %v, want %v", tt.model, result, tt.expected)
			}
		})
	}
}

func TestHasValidYANGSuffix(t *testing.T) {
	testCases := []struct {
		name     string
		model    string
		expected bool
	}{
		{"valid oper suffix", "Cisco-IOS-XE-wireless-afc-oper", true},
		{"valid cfg suffix", "Cisco-IOS-XE-wireless-ap-cfg", true},
		{"invalid suffix", "Cisco-IOS-XE-wireless-afc", false},
		{"empty string", "", false},
		{"just oper suffix", "-oper", true},
		{"just cfg suffix", "-cfg", true},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := hasValidYANGSuffix(tt.model)

			if result != tt.expected {
				t.Errorf("hasValidYANGSuffix(%q) = %v, want %v", tt.model, result, tt.expected)
			}
		})
	}
}

func TestIsValidRevision(t *testing.T) {
	testCases := []struct {
		name     string
		revision string
		expected bool
	}{
		{"valid revision", "2021-07-01", true},
		{"valid revision 2", "2022-12-31", true},
		{"invalid format - too short", "2021-7-1", false},
		{"invalid format - too long", "2021-07-011", false},
		{"invalid separators", "2021/07/01", false},
		{"invalid month", "2021-13-01", true}, // Note: This only checks format (digits), not actual date validity
		{"invalid day", "2021-07-32", true},   // Note: This only checks format (digits), not actual date validity
		{"non-numeric year", "abcd-07-01", false},
		{"non-numeric month", "2021-ab-01", false},
		{"non-numeric day", "2021-07-ab", false},
		{"empty string", "", false},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidRevision(tt.revision)

			if result != tt.expected {
				t.Errorf("IsValidRevision(%q) = %v, want %v", tt.revision, result, tt.expected)
			}
		})
	}
}

func TestHasValidDateFormat(t *testing.T) {
	testCases := []struct {
		name     string
		revision string
		expected bool
	}{
		{"valid format", "2021-07-01", true},
		{"invalid separators", "2021/07/01", false},
		{"no separators", "20210701", false},
		{"wrong separator positions", "21-07-011", false},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := hasValidDateFormat(tt.revision)

			if result != tt.expected {
				t.Errorf("hasValidDateFormat(%q) = %v, want %v", tt.revision, result, tt.expected)
			}
		})
	}
}

func TestHasValidDateComponents(t *testing.T) {
	testCases := []struct {
		name     string
		revision string
		expected bool
	}{
		{"all numeric", "2021-07-01", true},
		{"non-numeric year", "abcd-07-01", false},
		{"non-numeric month", "2021-ab-01", false},
		{"non-numeric day", "2021-07-ab", false},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := hasValidDateComponents(tt.revision)

			if result != tt.expected {
				t.Errorf("hasValidDateComponents(%q) = %v, want %v", tt.revision, result, tt.expected)
			}
		})
	}
}

func TestIsDigits(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{"all digits", "1234", true},
		{"single digit", "5", true},
		{"contains letters", "12a4", false},
		{"contains symbols", "12-4", false},
		{"empty string", "", true}, // Empty string has no non-digits
		{"zero", "0", true},
		{"leading zero", "0123", true},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := isDigits(tt.input)

			if result != tt.expected {
				t.Errorf("isDigits(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestConstants(t *testing.T) {
	// Test RESTCONF constants
	if RESTCONFPathPrefix != "/restconf/data" {
		t.Errorf("RESTCONFPathPrefix = %q, want %q", RESTCONFPathPrefix, "/restconf/data")
	}

	if RESTCONFModulesPathPrefix != "/restconf/tailf/modules" {
		t.Errorf("RESTCONFModulesPathPrefix = %q, want %q", RESTCONFModulesPathPrefix, "/restconf/tailf/modules")
	}

	expectedLibraryQuery := "?fields=ietf-yang-library:modules-state/module"
	if RESTCONFLibraryQuery != expectedLibraryQuery {
		t.Errorf("RESTCONFLibraryQuery = %q, want %q", RESTCONFLibraryQuery, expectedLibraryQuery)
	}

	// Test protocol constants
	if ProtocolHTTP != "http" {
		t.Errorf("ProtocolHTTP = %q, want %q", ProtocolHTTP, "http")
	}

	if ProtocolHTTPS != "https" {
		t.Errorf("ProtocolHTTPS = %q, want %q", ProtocolHTTPS, "https")
	}

	if DefaultProtocol != ProtocolHTTPS {
		t.Errorf("DefaultProtocol = %q, want %q", DefaultProtocol, ProtocolHTTPS)
	}

	// Test YANG model constants
	expectedPrefix := "Cisco-IOS-XE-wireless-"
	if RestconfYANGModelPrefix != expectedPrefix {
		t.Errorf("RestconfYANGModelPrefix = %q, want %q", RestconfYANGModelPrefix, expectedPrefix)
	}

	if RestconfYANGModelOperSuffix != "-oper" {
		t.Errorf("RestconfYANGModelOperSuffix = %q, want %q", RestconfYANGModelOperSuffix, "-oper")
	}

	if RestconfYANGModelCfgSuffix != "-cfg" {
		t.Errorf("RestconfYANGModelCfgSuffix = %q, want %q", RestconfYANGModelCfgSuffix, "-cfg")
	}
}
