package rogue_test

import (
	"errors"
	"net/http"
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

// TestRogueServiceUnit_ValidationErrors_EmptyMAC pins the sentinel an empty list key
// carries. It is core.ErrResourceNotFound, which core.IsNotFoundError reports true for:
// asking for a rogue record by an address that cannot name one is an absence, not a client
// misconfiguration. The nil client proves the guard runs before any request is built.
func TestRogueServiceUnit_ValidationErrors_EmptyMAC(t *testing.T) {
	service := rogue.NewService(nil)
	ctx := testutil.TestContext(t)

	lookups := map[string]func(string) error{
		"GetRogueByMAC": func(mac string) error {
			_, err := service.GetRogueByMAC(ctx, mac)
			return err
		},
		"GetRogueClientByMAC": func(mac string) error {
			_, err := service.GetRogueClientByMAC(ctx, mac)
			return err
		},
	}

	for name, call := range lookups {
		for _, mac := range []string{"", "   "} {
			err := call(mac)
			if !errors.Is(err, core.ErrResourceNotFound) {
				t.Errorf("%s(%q) error = %v, want core.ErrResourceNotFound", name, mac, err)
			}
			if errors.Is(err, core.ErrInvalidConfiguration) {
				t.Errorf("%s(%q) still reports the pre-v0.6.0 sentinel", name, mac)
			}
		}

		err := call("not-a-mac")
		if err == nil {
			t.Errorf("%s with a malformed address returned no error", name)
			continue
		}
		if core.IsNotFoundError(err) {
			t.Errorf("%s with a malformed address = %v, want a validation error", name, err)
		}
	}
}

// TestRogueServiceUnit_ByMAC_WireForm pins the address form that reaches the controller.
// Both rogue lookups key their list by the lower-case colon form, so a dashed or dotted
// address has to arrive as the same string rather than as a key no list holds.
func TestRogueServiceUnit_ByMAC_WireForm(t *testing.T) {
	const (
		rogueRoute  = "Cisco-IOS-XE-wireless-rogue-oper:rogue-oper-data/rogue-data"
		clientRoute = "Cisco-IOS-XE-wireless-rogue-oper:rogue-oper-data/rogue-client-data"
	)

	server := testutil.NewRESTCONFServer(t)
	defer server.Close()
	server.AddHandler(http.MethodGet, rogueRoute, func() (int, string) {
		return http.StatusOK, `{"Cisco-IOS-XE-wireless-rogue-oper:rogue-data":[]}`
	})
	server.AddHandler(http.MethodGet, clientRoute, func() (int, string) {
		return http.StatusOK, `{"Cisco-IOS-XE-wireless-rogue-oper:rogue-client-data":[]}`
	})

	testClient := testutil.NewTestClient(testutil.NewMockServerFromHTTP(server.Server))
	service := rogue.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	if _, err := service.GetRogueByMAC(ctx, "00-11-22-33-44-55"); err != nil {
		t.Fatalf("GetRogueByMAC unexpected error: %v", err)
	}
	if _, err := service.GetRogueClientByMAC(ctx, "0011.2233.4466"); err != nil {
		t.Fatalf("GetRogueClientByMAC unexpected error: %v", err)
	}

	recorded := server.Requests()
	if len(recorded) != 2 {
		t.Fatalf("Recorded %d requests, want 2", len(recorded))
	}
	if got, want := recorded[0].Path, "/restconf/data/"+rogueRoute+"=00:11:22:33:44:55"; got != want {
		t.Errorf("GetRogueByMAC wire path = %q, want %q", got, want)
	}
	if got, want := recorded[1].Path, "/restconf/data/"+clientRoute+"=00:11:22:33:44:66"; got != want {
		t.Errorf("GetRogueClientByMAC wire path = %q, want %q", got, want)
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

	// Test GetRLDPStats with nil client
	_, err = service.GetRLDPStats(ctx)
	if err == nil {
		t.Error("Expected error with nil client for GetRLDPStats, got nil")
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

// TestRogueServiceUnit_GetRLDPStats_MockSuccess holds GetRLDPStats to the envelope the
// controller answers a container read with. A read typed on the bare RLDPStats leaves every
// leaf at zero, which is indistinguishable from a controller reporting no RLDP activity, so
// every leaf here carries a distinct non-zero value.
func TestRogueServiceUnit_GetRLDPStats_MockSuccess(t *testing.T) {
	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{
		"Cisco-IOS-XE-wireless-rogue-oper:rogue-oper-data/rldp-stats": `{
			"Cisco-IOS-XE-wireless-rogue-oper:rldp-stats": {
				"num-in-progress": 1,
				"num-rldp-started": 2,
				"auth-timeout": 3,
				"assoc-timeout": 4,
				"dhcp-timeout": 5,
				"not-connected": 6,
				"connected": 7,
				"rldp-socket-enabled": true
			}
		}`,
	}))
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := rogue.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	result, err := service.GetRLDPStats(ctx)
	if err != nil {
		t.Fatalf("GetRLDPStats failed: %v", err)
	}

	stats := result.RLDPStats
	if stats == nil {
		t.Fatal("Expected rldp-stats to decode")
	}

	for _, leaf := range []struct {
		name string
		got  int
		want int
	}{
		{"num-in-progress", stats.NumInProgress, 1},
		{"num-rldp-started", stats.NumRLDPStarted, 2},
		{"auth-timeout", stats.AuthTimeout, 3},
		{"assoc-timeout", stats.AssocTimeout, 4},
		{"dhcp-timeout", stats.DHCPTimeout, 5},
		{"not-connected", stats.NotConnected, 6},
		{"connected", stats.Connected, 7},
	} {
		if leaf.got != leaf.want {
			t.Errorf("Expected %s to decode as %d, got %d", leaf.name, leaf.want, leaf.got)
		}
	}

	if !stats.RLDPSocketEnabled {
		t.Error("Expected rldp-socket-enabled to decode as true")
	}
}

// TestRogueServiceUnit_QuotedCounters_MockSuccess tests that the quoted 64-bit rogue counters
// decode alongside their bare siblings in the same container.
func TestRogueServiceUnit_QuotedCounters_MockSuccess(t *testing.T) {
	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{
		"Cisco-IOS-XE-wireless-rogue-oper:rogue-oper-data": `{
			"Cisco-IOS-XE-wireless-rogue-oper:rogue-oper-data": {
				"rogue-stats": {
					"adhoc-count": 3,
					"report-count": "41",
					"iapp-unconnected-client": "7",
					"unconnected-client-report": "8",
					"unconnected-client-count": "9",
					"unconnected-reports-drop": "10",
					"ap-drop-urwb-link": "11"
				}
			}
		}`,
	}))
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := rogue.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	result, err := service.GetOperational(ctx)
	if err != nil {
		t.Fatalf("GetOperational failed: %v", err)
	}

	if result.CiscoIOSXEWirelessRogueOperData == nil {
		t.Fatal("Expected rogue-oper-data to decode")
	}

	stats := result.CiscoIOSXEWirelessRogueOperData.RogueStats
	if stats.AdhocCount != 3 {
		t.Errorf("Expected the bare sibling to survive, got adhoc-count %d", stats.AdhocCount)
	}
	if stats.ReportCount != "41" {
		t.Errorf("Expected report-count to decode, got %q", stats.ReportCount)
	}

	quoted := map[string]string{
		"iapp-unconnected-client":   stats.IappUnconnectedClient,
		"unconnected-client-report": stats.UnconnectedClientReport,
		"unconnected-client-count":  stats.UnconnectedClientCount,
		"unconnected-reports-drop":  stats.UnconnectedReportsDrop,
		"ap-drop-urwb-link":         stats.ApDropURWBLink,
	}
	for leaf, got := range quoted {
		if got == "" {
			t.Errorf("Expected %s to decode", leaf)
		}
	}
}
