package fabric

import (
	"context"
	"testing"
	"time"

	testutils "github.com/umatare5/cisco-ios-xe-wireless-go/internal/tests"
)

func TestFabricNilClientHandling(t *testing.T) {
	ctx, cancel := testutils.CreateStandardTestContext()
	defer cancel()

	t.Run("GetFabricCfg with nil client", func(t *testing.T) {
		_, err := GetFabricCfg(nil, ctx)
		if err == nil {
			t.Error("Expected error when client is nil")
		}
		if err.Error() != "client is nil" {
			t.Errorf("Expected 'client is nil', got %v", err)
		}
	})

	t.Run("GetFabricControlplaneNames with nil client", func(t *testing.T) {
		_, err := GetFabricControlplaneNames(nil, ctx)
		if err == nil {
			t.Error("Expected error when client is nil")
		}
		if err.Error() != "client is nil" {
			t.Errorf("Expected 'client is nil', got %v", err)
		}
	})

	t.Run("GetFabric with nil client", func(t *testing.T) {
		_, err := GetFabric(nil, ctx)
		if err == nil {
			t.Error("Expected error when client is nil")
		}
		if err.Error() != "client is nil" {
			t.Errorf("Expected 'client is nil', got %v", err)
		}
	})
}

func TestFabricContextHandling(t *testing.T) {
	client := testutils.GetTestClient(t)

	cancelledCtx, cancel := context.WithCancel(context.Background())
	cancel()

	t.Run("GetFabricCfg with cancelled context", func(t *testing.T) {
		_, err := GetFabricCfg(client, cancelledCtx)
		if err == nil {
			t.Error("Expected error with cancelled context")
		}
	})

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 1*time.Nanosecond)
	defer cancel()
	time.Sleep(10 * time.Millisecond)

	t.Run("GetFabricControlplaneNames with timeout context", func(t *testing.T) {
		_, err := GetFabricControlplaneNames(client, timeoutCtx)
		if err == nil {
			t.Error("Expected error with timeout context")
		}
	})
}

func TestFabricConfigurationFunctions(t *testing.T) {
	client := testutils.CreateTestClientFromEnv(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	t.Run("GetFabricCfg", func(t *testing.T) {
		data, err := GetFabricCfg(client, ctx)
		if err != nil {
			t.Fatalf("GetFabricCfg failed: %v", err)
		}
		if data == nil {
			t.Fatal("GetFabricCfg returned nil data")
		}
		testutils.SaveTestDataToFile("fabric_cfg_data.json", data)
	})

	t.Run("GetFabricControlplaneNames", func(t *testing.T) {
		data, err := GetFabricControlplaneNames(client, ctx)
		if err != nil {
			t.Fatalf("GetFabricControlplaneNames failed: %v", err)
		}
		if data == nil {
			t.Fatal("GetFabricControlplaneNames returned nil data")
		}
		testutils.SaveTestDataToFile("fabric_controlplane_names_data.json", data)
	})

	t.Run("GetFabric", func(t *testing.T) {
		data, err := GetFabric(client, ctx)
		if err != nil {
			t.Fatalf("GetFabric failed: %v", err)
		}
		if data == nil {
			t.Fatal("GetFabric returned nil data")
		}
		testutils.SaveTestDataToFile("fabric_data.json", data)
	})
}

func TestFabricConfigurationEndpoints(t *testing.T) {
	t.Run("Validate_FabricCfgBasePath", func(t *testing.T) {
		expectedBasePath := "/restconf/data/Cisco-IOS-XE-wireless-fabric-cfg:fabric-cfg-data"
		testutils.EndpointValidationTest(t, FabricCfgBasePath, expectedBasePath)
	})

	t.Run("Validate_FabricCfgEndpoint", func(t *testing.T) {
		testutils.EndpointValidationTest(t, FabricCfgEndpoint, FabricCfgBasePath)
	})

	t.Run("Validate_FabricControlplaneNamesEndpoint", func(t *testing.T) {
		expectedEndpoint := FabricCfgBasePath + "/fabric-controlplane-names"
		testutils.EndpointValidationTest(t, FabricControlplaneNamesEndpoint, expectedEndpoint)
	})
}

func TestFabricDataStructures(t *testing.T) {
	t.Run("FabricCfgResponse", func(t *testing.T) {
		sampleJSON := `{
			"Cisco-IOS-XE-wireless-fabric-cfg:fabric-cfg-data": {
				"fabric": {},
				"fabric-controlplane-names": {
					"fabric-controlplane-name": []
				}
			}
		}`

		var response FabricCfgResponse
		testutils.TestJSONUnmarshal(t, sampleJSON, &response, "FabricCfgResponse")
	})

	t.Run("FabricControlplaneNamesResponse", func(t *testing.T) {
		sampleJSON := `{
			"Cisco-IOS-XE-wireless-fabric-cfg:fabric-controlplane-names": {
				"fabric-controlplane-name": []
			}
		}`

		var response FabricControlplaneNamesResponse
		testutils.TestJSONUnmarshal(t, sampleJSON, &response, "FabricControlplaneNamesResponse")
	})

	t.Run("FabricResponse", func(t *testing.T) {
		sampleJSON := `{
			"Cisco-IOS-XE-wireless-fabric-cfg:fabric": {}
		}`

		var response FabricResponse
		testutils.TestJSONUnmarshal(t, sampleJSON, &response, "FabricResponse")
	})
}

// Additional test for GetFabric error path coverage
func TestGetFabricErrorPaths(t *testing.T) {
	ctx := context.Background()

	// Test nil client case specifically to ensure complete coverage
	t.Run("GetFabric with nil client", func(t *testing.T) {
		_, err := GetFabric(nil, ctx)
		if err == nil {
			t.Error("Expected error with nil client")
		}
		if err.Error() != "client is nil" {
			t.Errorf("Expected 'client is nil', got '%s'", err.Error())
		}
	})

	// Test with valid client but cancelled context
	t.Run("GetFabric with cancelled context", func(t *testing.T) {
		client := testutils.GetTestClient(t)
		cancelledCtx, cancel := context.WithCancel(context.Background())
		cancel()

		_, err := GetFabric(client, cancelledCtx)
		if err == nil {
			t.Error("Expected error with cancelled context")
		}
	})
}
