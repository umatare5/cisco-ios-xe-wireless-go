// Package constants provides shared constants used across the Cisco IOS-XE Wireless Go SDK.
package constants

import (
	"strings"
	"testing"
	"time"
)

// TestYANGModelPrefixes tests YANG model prefix constants
func TestYANGModelPrefixes(t *testing.T) {
	t.Run("YANGModelPrefix", func(t *testing.T) {
		expected := "Cisco-IOS-XE-wireless-"
		if YANGModelPrefix != expected {
			t.Errorf("YANGModelPrefix = %q, want %q", YANGModelPrefix, expected)
		}
	})

	t.Run("YANGModelPrefixAccess", func(t *testing.T) {
		expected := "Cisco-IOS-XE-access-"
		if YANGModelPrefixAccess != expected {
			t.Errorf("YANGModelPrefixAccess = %q, want %q", YANGModelPrefixAccess, expected)
		}
	})

	t.Run("YANGModelPrefixSite", func(t *testing.T) {
		expected := "Cisco-IOS-XE-site-"
		if YANGModelPrefixSite != expected {
			t.Errorf("YANGModelPrefixSite = %q, want %q", YANGModelPrefixSite, expected)
		}
	})
}

// TestYANGModuleSuffixes tests YANG module suffix constants
func TestYANGModuleSuffixes(t *testing.T) {
	t.Run("CfgSuffix", func(t *testing.T) {
		expected := "-cfg"
		if CfgSuffix != expected {
			t.Errorf("CfgSuffix = %q, want %q", CfgSuffix, expected)
		}
	})

	t.Run("OperSuffix", func(t *testing.T) {
		expected := "-oper"
		if OperSuffix != expected {
			t.Errorf("OperSuffix = %q, want %q", OperSuffix, expected)
		}
	})

	t.Run("CfgDataSuffix", func(t *testing.T) {
		expected := "-cfg-data"
		if CfgDataSuffix != expected {
			t.Errorf("CfgDataSuffix = %q, want %q", CfgDataSuffix, expected)
		}
	})

	t.Run("OperDataSuffix", func(t *testing.T) {
		expected := "-oper-data"
		if OperDataSuffix != expected {
			t.Errorf("OperDataSuffix = %q, want %q", OperDataSuffix, expected)
		}
	})
}

// TestYANGModuleNames tests YANG module name constants
func TestYANGModuleNames(t *testing.T) {
	t.Run("YangModuleAFC", func(t *testing.T) {
		if YangModuleAFC != "afc" {
			t.Errorf("YangModuleAFC = %q, want %q", YangModuleAFC, "afc")
		}
	})

	t.Run("YangModuleAP", func(t *testing.T) {
		if YangModuleAP != "ap" {
			t.Errorf("YangModuleAP = %q, want %q", YangModuleAP, "ap")
		}
	})

	t.Run("YangModuleWLAN", func(t *testing.T) {
		if YangModuleWLAN != "wlan" {
			t.Errorf("YangModuleWLAN = %q, want %q", YangModuleWLAN, "wlan")
		}
	})

	t.Run("YangModuleSite", func(t *testing.T) {
		if YangModuleSite != "site" {
			t.Errorf("YangModuleSite = %q, want %q", YangModuleSite, "site")
		}
	})

	t.Run("YangModuleClient", func(t *testing.T) {
		if YangModuleClient != "client" {
			t.Errorf("YangModuleClient = %q, want %q", YangModuleClient, "client")
		}
	})
}

// TestAPServiceYANGModels tests AP service YANG model constants
func TestAPServiceYANGModels(t *testing.T) {
	t.Run("APCfgModel", func(t *testing.T) {
		expected := YANGModelPrefix + "ap-cfg"
		if APCfgModel != expected {
			t.Errorf("APCfgModel = %q, want %q", APCfgModel, expected)
		}
	})

	t.Run("APOperModel", func(t *testing.T) {
		expected := YANGModelPrefix + "access-point-oper"
		if APOperModel != expected {
			t.Errorf("APOperModel = %q, want %q", APOperModel, expected)
		}
	})

	t.Run("APGlobalOperModel", func(t *testing.T) {
		expected := YANGModelPrefix + "ap-global-oper"
		if APGlobalOperModel != expected {
			t.Errorf("APGlobalOperModel = %q, want %q", APGlobalOperModel, expected)
		}
	})
}

// TestConstants tests various constant values
func TestConstants(t *testing.T) {
	// Test that we have non-empty constants
	if YANGModelPrefix == "" {
		t.Error("YANGModelPrefix should not be empty")
	}
	if CfgSuffix == "" {
		t.Error("CfgSuffix should not be empty")
	}
	if OperSuffix == "" {
		t.Error("OperSuffix should not be empty")
	}
}

// TestHTTPConfiguration tests HTTP-related constants
func TestHTTPConfiguration(t *testing.T) {
	if DefaultHTTPTimeout != 30 {
		t.Errorf("DefaultHTTPTimeout = %d, want %d", DefaultHTTPTimeout, 30)
	}
	if DefaultMaxRetries != 3 {
		t.Errorf("DefaultMaxRetries = %d, want %d", DefaultMaxRetries, 3)
	}
}

// TestNetworkConstants tests network and protocol constants
func TestNetworkConstants(t *testing.T) {
	t.Run("NetworkTimeoutSeconds", func(t *testing.T) {
		if NetworkTimeoutSeconds != 60 {
			t.Errorf("NetworkTimeoutSeconds = %d, want %d", NetworkTimeoutSeconds, 60)
		}
	})

	t.Run("HTTPSScheme", func(t *testing.T) {
		if HTTPSScheme != "https" {
			t.Errorf("HTTPSScheme = %q, want %q", HTTPSScheme, "https")
		}
	})

	t.Run("HTTPScheme", func(t *testing.T) {
		if HTTPScheme != "http" {
			t.Errorf("HTTPScheme = %q, want %q", HTTPScheme, "http")
		}
	})

	t.Run("URLSchemeSeparator", func(t *testing.T) {
		if URLSchemeSeparator != "://" {
			t.Errorf("URLSchemeSeparator = %q, want %q", URLSchemeSeparator, "://")
		}
	})
}

// TestTimeoutConstants tests timeout constant values
func TestTimeoutConstants(t *testing.T) {
	t.Run("QuickTimeoutSeconds", func(t *testing.T) {
		if QuickTimeoutSeconds != 5 {
			t.Errorf("QuickTimeoutSeconds = %d, want %d", QuickTimeoutSeconds, 5)
		}
	})

	t.Run("StandardTimeoutSeconds", func(t *testing.T) {
		if StandardTimeoutSeconds != NetworkTimeoutSeconds {
			t.Errorf("StandardTimeoutSeconds = %d, want %d", StandardTimeoutSeconds, NetworkTimeoutSeconds)
		}
	})

	t.Run("ExtendedTimeoutSeconds", func(t *testing.T) {
		if ExtendedTimeoutSeconds != 90 {
			t.Errorf("ExtendedTimeoutSeconds = %d, want %d", ExtendedTimeoutSeconds, 90)
		}
	})

	t.Run("ComprehensiveTimeoutSeconds", func(t *testing.T) {
		if ComprehensiveTimeoutSeconds != 150 {
			t.Errorf("ComprehensiveTimeoutSeconds = %d, want %d", ComprehensiveTimeoutSeconds, 150)
		}
	})

	t.Run("MicroTimeoutMicroseconds", func(t *testing.T) {
		if MicroTimeoutMicroseconds != 1 {
			t.Errorf("MicroTimeoutMicroseconds = %d, want %d", MicroTimeoutMicroseconds, 1)
		}
	})
}

// TestTimeoutDurationConstants tests timeout duration values
func TestTimeoutDurationConstants(t *testing.T) {
	t.Run("QuickTimeout", func(t *testing.T) {
		expected := time.Duration(QuickTimeoutSeconds) * time.Second
		if QuickTimeout != expected {
			t.Errorf("QuickTimeout = %v, want %v", QuickTimeout, expected)
		}
	})

	t.Run("StandardTimeout", func(t *testing.T) {
		expected := time.Duration(StandardTimeoutSeconds) * time.Second
		if StandardTimeout != expected {
			t.Errorf("StandardTimeout = %v, want %v", StandardTimeout, expected)
		}
	})

	t.Run("ExtendedTimeout", func(t *testing.T) {
		expected := time.Duration(ExtendedTimeoutSeconds) * time.Second
		if ExtendedTimeout != expected {
			t.Errorf("ExtendedTimeout = %v, want %v", ExtendedTimeout, expected)
		}
	})

	t.Run("ComprehensiveTimeout", func(t *testing.T) {
		expected := time.Duration(ComprehensiveTimeoutSeconds) * time.Second
		if ComprehensiveTimeout != expected {
			t.Errorf("ComprehensiveTimeout = %v, want %v", ComprehensiveTimeout, expected)
		}
	})

	t.Run("MicroTimeout", func(t *testing.T) {
		expected := time.Duration(MicroTimeoutMicroseconds) * time.Microsecond
		if MicroTimeout != expected {
			t.Errorf("MicroTimeout = %v, want %v", MicroTimeout, expected)
		}
	})
}

// TestEnvironmentVariableConstants tests environment variable constants
func TestEnvironmentVariableConstants(t *testing.T) {
	t.Run("EnvVarController", func(t *testing.T) {
		if EnvVarController != "WNC_CONTROLLER" {
			t.Errorf("EnvVarController = %q, want %q", EnvVarController, "WNC_CONTROLLER")
		}
	})

	t.Run("EnvVarAccessToken", func(t *testing.T) {
		if EnvVarAccessToken != "WNC_ACCESS_TOKEN" {
			t.Errorf("EnvVarAccessToken = %q, want %q", EnvVarAccessToken, "WNC_ACCESS_TOKEN")
		}
	})
}

// TestDefaultValues tests default value constants
// No implicit default controller is provided; controller must be specified explicitly via configuration.

// TestDocumentationConstants tests documentation example constants
func TestDocumentationConstants(t *testing.T) {
	t.Run("ExampleControllerIPAddress", func(t *testing.T) {
		if ExampleControllerIPAddress != "192.168.1.100" {
			t.Errorf("ExampleControllerIPAddress = %q, want %q", ExampleControllerIPAddress, "192.168.1.100")
		}
	})

	t.Run("ExampleControllerHostname", func(t *testing.T) {
		if ExampleControllerHostname != "core.example.local" {
			t.Errorf("ExampleControllerHostname = %q, want %q", ExampleControllerHostname, "core.example.local")
		}
	})

	t.Run("ExampleAccessToken", func(t *testing.T) {
		if ExampleAccessToken != "your-token" {
			t.Errorf("ExampleAccessToken = %q, want %q", ExampleAccessToken, "your-token")
		}
	})

	t.Run("ExampleTestHostname", func(t *testing.T) {
		if ExampleTestHostname != "test.local" {
			t.Errorf("ExampleTestHostname = %q, want %q", ExampleTestHostname, "test.local")
		}
	})

	t.Run("ExampleTimeoutSeconds", func(t *testing.T) {
		if ExampleTimeoutSeconds != 20 {
			t.Errorf("ExampleTimeoutSeconds = %d, want %d", ExampleTimeoutSeconds, 20)
		}
	})
}

// TestTestConstants tests test-related constants
func TestTestConstants(t *testing.T) {
	t.Run("TestAccessTokenValue", func(t *testing.T) {
		if TestAccessTokenValue != "dGVzdDp0ZXN0" {
			t.Errorf("TestAccessTokenValue = %q, want %q", TestAccessTokenValue, "dGVzdDp0ZXN0")
		}
	})

	t.Run("TestTimestamp", func(t *testing.T) {
		if TestTimestamp != "2024-01-01T00:00:00.000Z" {
			t.Errorf("TestTimestamp = %q, want %q", TestTimestamp, "2024-01-01T00:00:00.000Z")
		}
	})

	t.Run("TestAPName", func(t *testing.T) {
		if TestAPName != "test-ap-01" {
			t.Errorf("TestAPName = %q, want %q", TestAPName, "test-ap-01")
		}
	})
}

// TestBuildYangModulePath tests BuildYangModulePath function
func TestBuildYangModulePath(t *testing.T) {
	t.Run("WLANCfg", func(t *testing.T) {
		result := BuildYangModulePath("wlan", "cfg")
		expected := "Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data"
		if result != expected {
			t.Errorf("BuildYangModulePath(\"wlan\", \"cfg\") = %q, want %q", result, expected)
		}
	})

	t.Run("APOper", func(t *testing.T) {
		result := BuildYangModulePath("ap", "oper")
		expected := "Cisco-IOS-XE-wireless-ap-oper:ap-oper-data"
		if result != expected {
			t.Errorf("BuildYangModulePath(\"ap\", \"oper\") = %q, want %q", result, expected)
		}
	})

	t.Run("ClientOper", func(t *testing.T) {
		result := BuildYangModulePath("client", "oper")
		expected := "Cisco-IOS-XE-wireless-client-oper:client-oper-data"
		if result != expected {
			t.Errorf("BuildYangModulePath(\"client\", \"oper\") = %q, want %q", result, expected)
		}
	})
}

// TestBuildWirelessYangModule tests BuildWirelessYangModule function
func TestBuildWirelessYangModule(t *testing.T) {
	t.Run("WLANCfg", func(t *testing.T) {
		result := BuildWirelessYangModule("wlan", "cfg")
		expected := "Cisco-IOS-XE-wireless-wlan-cfg"
		if result != expected {
			t.Errorf("BuildWirelessYangModule(\"wlan\", \"cfg\") = %q, want %q", result, expected)
		}
	})

	t.Run("APOper", func(t *testing.T) {
		result := BuildWirelessYangModule("ap", "oper")
		expected := "Cisco-IOS-XE-wireless-ap-oper"
		if result != expected {
			t.Errorf("BuildWirelessYangModule(\"ap\", \"oper\") = %q, want %q", result, expected)
		}
	})

	t.Run("ClientOper", func(t *testing.T) {
		result := BuildWirelessYangModule("client", "oper")
		expected := "Cisco-IOS-XE-wireless-client-oper"
		if result != expected {
			t.Errorf("BuildWirelessYangModule(\"client\", \"oper\") = %q, want %q", result, expected)
		}
	})
}

// TestBuildYangEndpoint tests BuildYangEndpoint function
func TestBuildYangEndpoint(t *testing.T) {
	t.Run("WithEndpoint", func(t *testing.T) {
		result := BuildYangEndpoint("wlan", "cfg", "profiles")
		expected := "Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/profiles"
		if result != expected {
			t.Errorf("BuildYangEndpoint(\"wlan\", \"cfg\", \"profiles\") = %q, want %q", result, expected)
		}
	})

	t.Run("WithoutEndpoint", func(t *testing.T) {
		result := BuildYangEndpoint("wlan", "cfg", "")
		expected := "Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data"
		if result != expected {
			t.Errorf("BuildYangEndpoint(\"wlan\", \"cfg\", \"\") = %q, want %q", result, expected)
		}
	})
}

// TestBuildAPCfgPath tests BuildAPCfgPath function
func TestBuildAPCfgPath(t *testing.T) {
	t.Run("WithEndpoint", func(t *testing.T) {
		result := BuildAPCfgPath("profiles")
		expected := "Cisco-IOS-XE-wireless-ap-cfg:ap-cfg-data/profiles"
		if result != expected {
			t.Errorf("BuildAPCfgPath(\"profiles\") = %q, want %q", result, expected)
		}
	})

	t.Run("WithoutEndpoint", func(t *testing.T) {
		result := BuildAPCfgPath("")
		expected := "Cisco-IOS-XE-wireless-ap-cfg:ap-cfg-data"
		if result != expected {
			t.Errorf("BuildAPCfgPath(\"\") = %q, want %q", result, expected)
		}
	})
}

// TestBuildAPOperPath tests BuildAPOperPath function
func TestBuildAPOperPath(t *testing.T) {
	t.Run("WithEndpoint", func(t *testing.T) {
		result := BuildAPOperPath("status")
		expected := "/Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/status"
		if result != expected {
			t.Errorf("BuildAPOperPath(\"status\") = %q, want %q", result, expected)
		}
	})

	t.Run("WithoutEndpoint", func(t *testing.T) {
		result := BuildAPOperPath("")
		expected := "/Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data"
		if result != expected {
			t.Errorf("BuildAPOperPath(\"\") = %q, want %q", result, expected)
		}
	})
}

// TestBuildAPGlobalOperPath tests BuildAPGlobalOperPath function
func TestBuildAPGlobalOperPath(t *testing.T) {
	t.Run("WithEndpoint", func(t *testing.T) {
		result := BuildAPGlobalOperPath("summary")
		expected := "Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data/summary"
		if result != expected {
			t.Errorf("BuildAPGlobalOperPath(\"summary\") = %q, want %q", result, expected)
		}
	})

	t.Run("WithoutEndpoint", func(t *testing.T) {
		result := BuildAPGlobalOperPath("")
		expected := "Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data"
		if result != expected {
			t.Errorf("BuildAPGlobalOperPath(\"\") = %q, want %q", result, expected)
		}
	})
}

// TestAllModuleNames tests that all YANG module names are non-empty and lowercase
func TestAllModuleNames(t *testing.T) {
	modules := map[string]string{
		"YangModuleAFC":           YangModuleAFC,
		"YangModuleAP":            YangModuleAP,
		"YangModuleAPF":           YangModuleAPF,
		"YangModuleAWIPS":         YangModuleAWIPS,
		"YangModuleBLE":           YangModuleBLE,
		"YangModuleClient":        YangModuleClient,
		"YangModuleCTS":           YangModuleCTS,
		"YangModuleDot11":         YangModuleDot11,
		"YangModuleDot15":         YangModuleDot15,
		"YangModuleFabric":        YangModuleFabric,
		"YangModuleFlex":          YangModuleFlex,
		"YangModuleGeneral":       YangModuleGeneral,
		"YangModuleGeolocation":   YangModuleGeolocation,
		"YangModuleHyperlocation": YangModuleHyperlocation,
		"YangModuleLISP":          YangModuleLISP,
		"YangModuleLocation":      YangModuleLocation,
		"YangModuleMcast":         YangModuleMcast,
		"YangModuleMDNS":          YangModuleMDNS,
		"YangModuleMesh":          YangModuleMesh,
		"YangModuleMobility":      YangModuleMobility,
		"YangModuleNMSP":          YangModuleNMSP,
		"YangModuleRadio":         YangModuleRadio,
		"YangModuleRF":            YangModuleRF,
		"YangModuleRFID":          YangModuleRFID,
		"YangModuleRogue":         YangModuleRogue,
		"YangModuleRRM":           YangModuleRRM,
		"YangModuleSite":          YangModuleSite,
		"YangModuleWLAN":          YangModuleWLAN,
	}

	for name, value := range modules {
		if value == "" {
			t.Errorf("%s should not be empty", name)
		}
		if strings.ToLower(value) != value {
			t.Errorf("%s = %q should be lowercase", name, value)
		}
	}
}

// TestTestConfigurationConstants tests test configuration constants
func TestTestConfigurationConstants(t *testing.T) {
	if DefaultTestMethods != 10 {
		t.Errorf("DefaultTestMethods = %d, want %d", DefaultTestMethods, 10)
	}
	if StandardTestPhases != 4 {
		t.Errorf("StandardTestPhases = %d, want %d", StandardTestPhases, 4)
	}
	if DefaultTestGoroutines != 10 {
		t.Errorf("DefaultTestGoroutines = %d, want %d", DefaultTestGoroutines, 10)
	}
	if RogueServiceMethods != 5 {
		t.Errorf("RogueServiceMethods = %d, want %d", RogueServiceMethods, 5)
	}
	if SingleMethodServices != 1 {
		t.Errorf("SingleMethodServices = %d, want %d", SingleMethodServices, 1)
	}
	if WLANServiceMethods != 6 {
		t.Errorf("WLANServiceMethods = %d, want %d", WLANServiceMethods, 6)
	}
}
