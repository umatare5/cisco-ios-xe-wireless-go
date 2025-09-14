package rogue_test

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/pkg/testutil"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/rogue"
)

func TestRogueServiceUnit_Constructor_Success(t *testing.T) {
	service := rogue.NewService(nil)
	if service.Client() != nil {
		t.Error("Expected nil client service")
	}

	// Test with valid client
	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{
		"test": `{"data": {}}`,
	}))
	defer mockServer.Close()

	client := testutil.NewTestClient(mockServer)
	service = rogue.NewService(client.Core().(*core.Client))
	if service.Client() == nil {
		t.Error("Expected service to have client, got nil")
	}
}

func TestRogueServiceUnit_GetOperations_MockSuccess(t *testing.T) {
	// Using real WNC rogue data structure with actual MAC addresses from live environment
	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{
		"Cisco-IOS-XE-wireless-rogue-oper:rogue-oper-data": `{
			"Cisco-IOS-XE-wireless-rogue-oper:rogue-oper-data": {
				"rogue-stats": {
					"total-count": 12,
					"unclassified-count": 12,
					"alert-count": 12,
					"malicious-count": 0,
					"friendly-count": 0
				},
				"rogue-data": [
					{
						"rogue-address": "00:25:36:57:ed:cb",
						"rogue-class-type": "rogue-classtype-unclassified",
						"rogue-mode": "rogue-state-alert",
						"rogue-containment-level": 0,
						"contained": false,
						"rogue-first-timestamp": "2025-09-10T16:27:41.521656+00:00",
						"rogue-last-timestamp": "2025-09-10T16:54:41.506309+00:00",
						"max-detected-rssi": -67,
						"ssid-max-rssi": "rt500k-57ed8b-3"
					},
					{
						"rogue-address": "08:10:86:bf:07:e3",
						"rogue-class-type": "rogue-classtype-unclassified",
						"rogue-mode": "rogue-state-alert",
						"rogue-containment-level": 0,
						"contained": false,
						"rogue-first-timestamp": "2025-09-09T10:55:03.573664+00:00",
						"rogue-last-timestamp": "2025-09-10T17:04:41.502126+00:00",
						"max-detected-rssi": -56,
						"ssid-max-rssi": "aterm-b5acbb-g"
					}
				],
				"rogue-client-data": [
					{
						"rogue-client-address": "2a:c5:50:5d:6b:9c",
						"rogue-client-bssid": "1c:61:b4:10:0e:7f",
						"rogue-client-state": "rogue-state-alert",
						"rogue-client-containment-level": 0,
						"contained": false,
						"rogue-client-first-timestamp": "2025-09-10T16:32:41.518791+00:00",
						"rogue-client-last-timestamp": "2025-09-10T17:05:41.501647+00:00"
					}
				]
			}
		}`,
	}))
	defer mockServer.Close()

	client := testutil.NewTestClient(mockServer)
	service := rogue.NewService(client.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test GetOperational
	result, err := service.GetOperational(ctx)
	if err != nil {
		t.Errorf("GetOperational failed: %v", err)
		return
	}

	// Verify result structure based on live WNC data
	if result == nil {
		t.Error("GetOperational returned nil result")
		return
	}

	t.Logf("GetOperational returned valid rogue data with live WNC structure")
}

func TestRogueServiceUnit_ListOperations_MockSuccess(t *testing.T) {
	// Using real WNC rogue data structure
	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{
		"Cisco-IOS-XE-wireless-rogue-oper:rogue-oper-data/rogue-data": `{
			"Cisco-IOS-XE-wireless-rogue-oper:rogue-data": [
				{
					"rogue-address": "00:25:36:57:ed:cb",
					"rogue-class-type": "rogue-classtype-unclassified",
					"rogue-mode": "rogue-state-alert",
					"rogue-containment-level": 0,
					"contained": false,
					"rogue-first-timestamp": "2025-09-10T16:27:41.521656+00:00",
					"rogue-last-timestamp": "2025-09-10T16:54:41.506309+00:00",
					"max-detected-rssi": -67,
					"ssid-max-rssi": "rt500k-57ed8b-3"
				}
			]
		}`,
		"Cisco-IOS-XE-wireless-rogue-oper:rogue-oper-data/rogue-client-data": `{
			"Cisco-IOS-XE-wireless-rogue-oper:rogue-client-data": [
				{
					"rogue-client-address": "2a:c5:50:5d:6b:9c",
					"rogue-client-bssid": "1c:61:b4:10:0e:7f",
					"rogue-client-state": "rogue-state-alert",
					"rogue-client-containment-level": 0,
					"contained": false,
					"rogue-client-first-timestamp": "2025-09-10T16:32:41.518791+00:00",
					"rogue-client-last-timestamp": "2025-09-10T17:05:41.501647+00:00"
				}
			]
		}`,
		"Cisco-IOS-XE-wireless-rogue-oper:rogue-oper-data/rogue-stats": `{
			"Cisco-IOS-XE-wireless-rogue-oper:rogue-stats": {
				"total-count": 12,
				"unclassified-count": 12,
				"alert-count": 12,
				"malicious-count": 0,
				"friendly-count": 0
			}
		}`,
	}))
	defer mockServer.Close()

	client := testutil.NewTestClient(mockServer)
	service := rogue.NewService(client.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test ListRogues
	rogues, err := service.ListRogues(ctx)
	if err != nil {
		t.Errorf("ListRogues failed: %v", err)
		return
	}

	if rogues == nil {
		t.Error("ListRogues returned nil result")
		return
	}

	// Test ListRogueClients
	clients, err := service.ListRogueClients(ctx)
	if err != nil {
		t.Errorf("ListRogueClients failed: %v", err)
		return
	}

	if clients == nil {
		t.Error("ListRogueClients returned nil result")
		return
	}

	// Test GetStats
	stats, err := service.GetStats(ctx)
	if err != nil {
		t.Errorf("GetStats failed: %v", err)
		return
	}

	if stats == nil {
		t.Error("GetStats returned nil result")
		return
	}

	t.Logf("List operations returned valid rogue data with live WNC structure")
}

func TestRogueServiceUnit_GetByMACOperations_MockSuccess(t *testing.T) {
	// Using real WNC rogue data structure
	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{
		"Cisco-IOS-XE-wireless-rogue-oper:rogue-oper-data/rogue-data=00:25:36:57:ed:cb": `{
			"Cisco-IOS-XE-wireless-rogue-oper:rogue-data": [
				{
					"rogue-address": "00:25:36:57:ed:cb",
					"rogue-class-type": "rogue-classtype-unclassified",
					"rogue-mode": "rogue-state-alert",
					"rogue-containment-level": 0,
					"contained": false,
					"rogue-first-timestamp": "2025-09-10T16:27:41.521656+00:00",
					"rogue-last-timestamp": "2025-09-10T16:54:41.506309+00:00",
					"max-detected-rssi": -67,
					"ssid-max-rssi": "rt500k-57ed8b-3"
				}
			]
		}`,
		"Cisco-IOS-XE-wireless-rogue-oper:rogue-oper-data/rogue-client-data=2a:c5:50:5d:6b:9c": `{
			"Cisco-IOS-XE-wireless-rogue-oper:rogue-client-data": [
				{
					"rogue-client-address": "2a:c5:50:5d:6b:9c",
					"rogue-client-bssid": "1c:61:b4:10:0e:7f",
					"rogue-client-state": "rogue-state-alert",
					"rogue-client-containment-level": 0,
					"contained": false,
					"rogue-client-first-timestamp": "2025-09-10T16:32:41.518791+00:00",
					"rogue-client-last-timestamp": "2025-09-10T17:05:41.501647+00:00"
				}
			]
		}`,
	}))
	defer mockServer.Close()

	client := testutil.NewTestClient(mockServer)
	service := rogue.NewService(client.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test GetRogueByMAC with valid MAC
	rogueData, err := service.GetRogueByMAC(ctx, "00:25:36:57:ed:cb")
	if err != nil {
		t.Errorf("GetRogueByMAC failed: %v", err)
		return
	}

	if rogueData == nil {
		t.Error("GetRogueByMAC returned nil result")
		return
	}

	// Test GetRogueClientByMAC with valid MAC
	clientData, err := service.GetRogueClientByMAC(ctx, "2a:c5:50:5d:6b:9c")
	if err != nil {
		t.Errorf("GetRogueClientByMAC failed: %v", err)
		return
	}

	if clientData == nil {
		t.Error("GetRogueClientByMAC returned nil result")
		return
	}

	t.Logf("Get by MAC operations returned valid rogue data with live WNC structure")
}

func TestRogueServiceUnit_ValidationErrors_EmptyMAC(t *testing.T) {
	service := rogue.NewService(nil)
	ctx := testutil.TestContext(t)

	// Test GetRogueByMAC with empty MAC
	_, err := service.GetRogueByMAC(ctx, "")
	if err == nil {
		t.Error("Expected error with empty MAC for GetRogueByMAC, got nil")
	}

	// Test GetRogueClientByMAC with empty MAC
	_, err = service.GetRogueClientByMAC(ctx, "")
	if err == nil {
		t.Error("Expected error with empty MAC for GetRogueClientByMAC, got nil")
	}
}

func TestRogueServiceUnit_ErrorHandling_NilClient(t *testing.T) {
	service := rogue.NewService(nil)
	ctx := testutil.TestContext(t)

	// Test GetOperational with nil client
	_, err := service.GetOperational(ctx)
	if err == nil {
		t.Error("Expected error with nil client for GetOperational, got nil")
	}

	// Test ListRogues with nil client
	_, err = service.ListRogues(ctx)
	if err == nil {
		t.Error("Expected error with nil client for ListRogues, got nil")
	}

	// Test ListRogueClients with nil client
	_, err = service.ListRogueClients(ctx)
	if err == nil {
		t.Error("Expected error with nil client for ListRogueClients, got nil")
	}

	// Test GetStats with nil client
	_, err = service.GetStats(ctx)
	if err == nil {
		t.Error("Expected error with nil client for GetStats, got nil")
	}

	// Test GetRogueByMAC with nil client
	_, err = service.GetRogueByMAC(ctx, "00:25:36:57:ed:cb")
	if err == nil {
		t.Error("Expected error with nil client for GetRogueByMAC, got nil")
	}

	// Test GetRogueClientByMAC with nil client
	_, err = service.GetRogueClientByMAC(ctx, "2a:c5:50:5d:6b:9c")
	if err == nil {
		t.Error("Expected error with nil client for GetRogueClientByMAC, got nil")
	}
}

func TestRogueServiceUnit_AliasOperations_MockSuccess(t *testing.T) {
	// Test alias methods for integration test compatibility
	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{
		"Cisco-IOS-XE-wireless-rogue-oper:rogue-oper-data": `{
			"Cisco-IOS-XE-wireless-rogue-oper:rogue-oper-data": {
				"rogue-stats": {
					"total-count": 12,
					"unclassified-count": 12,
					"alert-count": 12,
					"malicious-count": 0,
					"friendly-count": 0
				}
			}
		}`,
		"Cisco-IOS-XE-wireless-rogue-oper:rogue-oper-data/rogue-data": `{
			"Cisco-IOS-XE-wireless-rogue-oper:rogue-data": [
				{
					"rogue-address": "00:25:36:57:ed:cb",
					"rogue-class-type": "rogue-classtype-unclassified",
					"rogue-mode": "rogue-state-alert",
					"rogue-containment-level": 0,
					"contained": false
				}
			]
		}`,
		"Cisco-IOS-XE-wireless-rogue-oper:rogue-oper-data/rogue-client-data": `{
			"Cisco-IOS-XE-wireless-rogue-oper:rogue-client-data": [
				{
					"rogue-client-address": "2a:c5:50:5d:6b:9c",
					"rogue-client-bssid": "1c:61:b4:10:0e:7f",
					"rogue-client-state": "rogue-state-alert"
				}
			]
		}`,
		"Cisco-IOS-XE-wireless-rogue-oper:rogue-oper-data/rogue-stats": `{
			"Cisco-IOS-XE-wireless-rogue-oper:rogue-stats": {
				"total-count": 12,
				"unclassified-count": 12,
				"alert-count": 12,
				"malicious-count": 0,
				"friendly-count": 0
			}
		}`,
		"Cisco-IOS-XE-wireless-rogue-oper:rogue-oper-data/rogue-data=00:25:36:57:ed:cb": `{
			"Cisco-IOS-XE-wireless-rogue-oper:rogue-data": [
				{
					"rogue-address": "00:25:36:57:ed:cb",
					"rogue-class-type": "rogue-classtype-unclassified",
					"rogue-mode": "rogue-state-alert",
					"rogue-containment-level": 0,
					"contained": false
				}
			]
		}`,
		"Cisco-IOS-XE-wireless-rogue-oper:rogue-oper-data/rogue-client-data=2a:c5:50:5d:6b:9c": `{
			"Cisco-IOS-XE-wireless-rogue-oper:rogue-client-data": [
				{
					"rogue-client-address": "2a:c5:50:5d:6b:9c",
					"rogue-client-bssid": "1c:61:b4:10:0e:7f",
					"rogue-client-state": "rogue-state-alert"
		}
	]
}`,
	}))
	defer mockServer.Close()

	client := testutil.NewTestClient(mockServer)
	service := rogue.NewService(client.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test GetOperClientData (alias for ListRogueClients)
	clientData, err := service.GetOperClientData(ctx)
	if err != nil {
		t.Errorf("GetOperClientData failed: %v", err)
		return
	}

	if clientData == nil {
		t.Error("GetOperClientData returned nil result")
		return
	}

	// Test GetOperData (alias for ListRogues)
	operData, err := service.GetOperData(ctx)
	if err != nil {
		t.Errorf("GetOperData failed: %v", err)
		return
	}

	if operData == nil {
		t.Error("GetOperData returned nil result")
		return
	}

	// Test GetOperStats (alias for GetStats)
	operStats, err := service.GetOperStats(ctx)
	if err != nil {
		t.Errorf("GetOperStats failed: %v", err)
		return
	}

	if operStats == nil {
		t.Error("GetOperStats returned nil result")
		return
	}

	// Test GetRLDPStats
	rldpStats, err := service.GetRLDPStats(ctx)
	if err != nil {
		t.Errorf("GetRLDPStats failed: %v", err)
		return
	}

	if rldpStats == nil {
		t.Error("GetRLDPStats returned nil result")
		return
	}

	// Test GetOperByRogueAddress (alias for GetRogueByMAC)
	rogueByAddr, err := service.GetOperByRogueAddress(ctx, "00:25:36:57:ed:cb")
	if err != nil {
		t.Errorf("GetOperByRogueAddress failed: %v", err)
		return
	}

	if rogueByAddr == nil {
		t.Error("GetOperByRogueAddress returned nil result")
		return
	}

	// Test GetOperByRogueClientAddress (alias for GetRogueClientByMAC)
	clientByAddr, err := service.GetOperByRogueClientAddress(ctx, "2a:c5:50:5d:6b:9c")
	if err != nil {
		t.Errorf("GetOperByRogueClientAddress failed: %v", err)
		return
	}

	if clientByAddr == nil {
		t.Error("GetOperByRogueClientAddress returned nil result")
		return
	}

	t.Logf("Alias operations returned valid rogue data with live WNC structure")
}
