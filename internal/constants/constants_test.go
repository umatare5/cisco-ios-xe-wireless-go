package constants

import (
	"testing"
)

// TestYANGModelPrefixes tests YANG model prefix constants
func TestYANGModelPrefixes(t *testing.T) {
	testCases := []struct {
		name     string
		value    string
		expected string
	}{
		{"YANGModelPrefix", YANGModelPrefix, "Cisco-IOS-XE-wireless-"},
		{"YANGModelPrefixAccess", YANGModelPrefixAccess, "Cisco-IOS-XE-access-"},
		{"YANGModelPrefixSite", YANGModelPrefixSite, "Cisco-IOS-XE-site-"},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.expected {
				t.Errorf("Expected %s to be '%s', got '%s'", tt.name, tt.expected, tt.value)
			}
		})
	}
}

// TestYANGModuleSuffixes tests YANG module suffix constants
func TestYANGModuleSuffixes(t *testing.T) {
	testCases := []struct {
		name     string
		value    string
		expected string
	}{
		{"CfgSuffix", CfgSuffix, "-cfg"},
		{"OperSuffix", OperSuffix, "-oper"},
		{"CfgDataSuffix", CfgDataSuffix, "-cfg-data"},
		{"OperDataSuffix", OperDataSuffix, "-oper-data"},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.expected {
				t.Errorf("Expected %s to be '%s', got '%s'", tt.name, tt.expected, tt.value)
			}
		})
	}
}

// TestYANGModuleNames tests YANG module name constants
func TestYANGModuleNames(t *testing.T) {
	testCases := []struct {
		name     string
		value    string
		expected string
	}{
		{"YangModuleAFC", YangModuleAFC, "afc"},
		{"YangModuleAP", YangModuleAP, "ap"},
		{"YangModuleWLAN", YangModuleWLAN, "wlan"},
		{"YangModuleSite", YangModuleSite, "site"},
		{"YangModuleClient", YangModuleClient, "client"},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.expected {
				t.Errorf("Expected %s to be '%s', got '%s'", tt.name, tt.expected, tt.value)
			}
		})
	}
}

// TestAPServiceYANGModels tests AP service specific YANG model constants
func TestAPServiceYANGModels(t *testing.T) {
	testCases := []struct {
		name     string
		value    string
		expected string
	}{
		{"APCfgModel", APCfgModel, "Cisco-IOS-XE-wireless-ap-cfg"},
		{"APOperModel", APOperModel, "Cisco-IOS-XE-wireless-access-point-oper"},
		{"APGlobalOperModel", APGlobalOperModel, "Cisco-IOS-XE-wireless-ap-global-oper"},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.expected {
				t.Errorf("Expected %s to be '%s', got '%s'", tt.name, tt.expected, tt.value)
			}
		})
	}
}

// TestTestConfiguration tests test configuration constants
func TestConstants(t *testing.T) {
	// Test all configuration constants
	if DefaultTestGoroutines != 10 {
		t.Errorf("Expected DefaultTestGoroutines to be 10, got %d", DefaultTestGoroutines)
	}

	if DefaultTestMethods != 10 {
		t.Errorf("Expected DefaultTestMethods to be 10, got %d", DefaultTestMethods)
	}

	if StandardTestPhases != 4 {
		t.Errorf("Expected StandardTestPhases to be 4, got %d", StandardTestPhases)
	}

	if RogueServiceMethods != 5 {
		t.Errorf("Expected RogueServiceMethods to be 5, got %d", RogueServiceMethods)
	}

	if SingleMethodServices != 1 {
		t.Errorf("Expected SingleMethodServices to be 1, got %d", SingleMethodServices)
	}

	if WLANServiceMethods != 6 {
		t.Errorf("Expected WLANServiceMethods to be 6, got %d", WLANServiceMethods)
	}
}

// TestHTTPConfiguration tests HTTP configuration constants
func TestHTTPConfiguration(t *testing.T) {
	if DefaultHTTPTimeout != 30 {
		t.Errorf("Expected DefaultHTTPTimeout to be 30, got %d", DefaultHTTPTimeout)
	}

	if DefaultMaxRetries != 3 {
		t.Errorf("Expected DefaultMaxRetries to be 3, got %d", DefaultMaxRetries)
	}
}

// TestBuildYangModulePath tests BuildYangModulePath function
func TestBuildYangModulePath(t *testing.T) {
	testCases := []struct {
		name       string
		module     string
		moduleType string
		expected   string
	}{
		{"WLANCfg", "wlan", "cfg", "Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data"},
		{"APOper", "ap", "oper", "Cisco-IOS-XE-wireless-ap-oper:ap-oper-data"},
		{"ClientOper", "client", "oper", "Cisco-IOS-XE-wireless-client-oper:client-oper-data"},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := BuildYangModulePath(tt.module, tt.moduleType)
			if result != tt.expected {
				t.Errorf("BuildYangModulePath(%q, %q) = %q, expected %q", tt.module, tt.moduleType, result, tt.expected)
			}
		})
	}
}

// TestBuildWirelessYangModule tests BuildWirelessYangModule function
func TestBuildWirelessYangModule(t *testing.T) {
	testCases := []struct {
		name       string
		module     string
		moduleType string
		expected   string
	}{
		{"WLANCfg", "wlan", "cfg", "Cisco-IOS-XE-wireless-wlan-cfg"},
		{"APOper", "ap", "oper", "Cisco-IOS-XE-wireless-ap-oper"},
		{"ClientOper", "client", "oper", "Cisco-IOS-XE-wireless-client-oper"},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := BuildWirelessYangModule(tt.module, tt.moduleType)
			if result != tt.expected {
				t.Errorf("BuildWirelessYangModule(%q, %q) = %q, expected %q", tt.module, tt.moduleType, result, tt.expected)
			}
		})
	}
}

// TestBuildYangEndpoint tests BuildYangEndpoint function
func TestBuildYangEndpoint(t *testing.T) {
	testCases := []struct {
		name       string
		module     string
		moduleType string
		endpoint   string
		expected   string
	}{
		{
			"WithEndpoint",
			"wlan", "cfg", "wlan-cfg-entries",
			"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/wlan-cfg-entries",
		},
		{
			"WithoutEndpoint",
			"ap", "oper", "",
			"Cisco-IOS-XE-wireless-ap-oper:ap-oper-data",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := BuildYangEndpoint(tt.module, tt.moduleType, tt.endpoint)
			if result != tt.expected {
				t.Errorf("BuildYangEndpoint(%q, %q, %q) = %q, expected %q", tt.module, tt.moduleType, tt.endpoint, result, tt.expected)
			}
		})
	}
}

// TestBuildAPCfgPath tests BuildAPCfgPath function
func TestBuildAPCfgPath(t *testing.T) {
	testCases := []struct {
		name     string
		endpoint string
		expected string
	}{
		{"WithEndpoint", "ap-tags", "Cisco-IOS-XE-wireless-ap-cfg:ap-cfg-data/ap-tags"},
		{"WithoutEndpoint", "", "Cisco-IOS-XE-wireless-ap-cfg:ap-cfg-data"},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := BuildAPCfgPath(tt.endpoint)
			if result != tt.expected {
				t.Errorf("BuildAPCfgPath(%q) = %q, expected %q", tt.endpoint, result, tt.expected)
			}
		})
	}
}

// TestBuildAPOperPath tests BuildAPOperPath function
func TestBuildAPOperPath(t *testing.T) {
	testCases := []struct {
		name     string
		endpoint string
		expected string
	}{
		{"WithEndpoint", "ap-radio-neighbor", "/Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/ap-radio-neighbor"},
		{"WithoutEndpoint", "", "/Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data"},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := BuildAPOperPath(tt.endpoint)
			if result != tt.expected {
				t.Errorf("BuildAPOperPath(%q) = %q, expected %q", tt.endpoint, result, tt.expected)
			}
		})
	}
}

// TestBuildAPGlobalOperPath tests BuildAPGlobalOperPath function
func TestBuildAPGlobalOperPath(t *testing.T) {
	testCases := []struct {
		name     string
		endpoint string
		expected string
	}{
		{"WithEndpoint", "ap-history", "Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data/ap-history"},
		{"WithoutEndpoint", "", "Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data"},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := BuildAPGlobalOperPath(tt.endpoint)
			if result != tt.expected {
				t.Errorf("BuildAPGlobalOperPath(%q) = %q, expected %q", tt.endpoint, result, tt.expected)
			}
		})
	}
}
