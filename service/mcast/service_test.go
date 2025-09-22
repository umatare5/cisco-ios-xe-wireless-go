package mcast_test

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/pkg/testutil"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/mcast"
)

// TestMcastServiceUnit_Constructor_Success tests service constructor functionality.
func TestMcastServiceUnit_Constructor_Success(t *testing.T) {
	t.Run("NewServiceWithValidClient", func(t *testing.T) {
		// Create mock server and test client using public API
		responses := map[string]string{
			"test-endpoint": `{"status": "success"}`,
		}
		mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
		defer mockServer.Close()

		testClient := testutil.NewTestClient(mockServer)
		service := mcast.NewService(testClient.Core().(*core.Client))

		// Test that service can be created and has a client
		if service.Client() == nil {
			t.Error("Expected service to have a client")
		}
	})

	t.Run("NewServiceWithNilClient", func(t *testing.T) {
		service := mcast.NewService(nil)

		// Service should still be created even with nil client
		if service.Client() != nil {
			t.Error("Expected service with nil client to return nil from Client()")
		}
	})
}

// TestMcastServiceUnit_GetOperations_MockSuccess tests Get operations using mock server
// with real multicast operational data structure from live Cisco C9800 IOS-XE 17.12.5.
func TestMcastServiceUnit_GetOperations_MockSuccess(t *testing.T) {
	// Create mock RESTCONF server with real multicast operational data structure
	// Based on live WNC data from IOS-XE 17.12.5 environment
	responses := map[string]string{
		// GetOperational - Complete multicast operational data with FlexConnect and VLAN info
		"Cisco-IOS-XE-wireless-mcast-oper:mcast-oper-data": `{
			"Cisco-IOS-XE-wireless-mcast-oper:mcast-oper-data": {
				"flex-mediastream-client-summary": [
					{
						"client-mac": "2a:e3:42:8f:06:c8",
						"vlan-id": 800,
						"flex-mcast-client-group": [
							{
								"mcast-ip": "224.0.0.251",
								"stream-name": "-",
								"ap-mac": "aa:bb:cc:dd:ee:ff",
								"is-direct": false
							},
							{
								"mcast-ip": "ff02::fb",
								"stream-name": "-",
								"ap-mac": "aa:bb:cc:dd:ee:ff",
								"is-direct": false
							}
						]
					}
				],
				"vlan-l2-mgid-op": [
					{
						"vlan-index": 1,
						"is-nonip-multicast-enabled": true,
						"is-broadcast-enable": true
					},
					{
						"vlan-index": 800,
						"is-nonip-multicast-enabled": true,
						"is-broadcast-enable": true
					}
				]
			}
		}`,
		// GetFlexConnectMediastreamClientSummary - FlexConnect mediastream client data
		"Cisco-IOS-XE-wireless-mcast-oper:mcast-oper-data/flex-mediastream-client-summary": `{
			"Cisco-IOS-XE-wireless-mcast-oper:flex-mediastream-client-summary": [
				{
					"client-mac": "2a:e3:42:8f:06:c8",
					"vlan-id": 800,
					"flex-mcast-client-group": [
						{
							"mcast-ip": "224.0.0.251",
							"stream-name": "-",
							"ap-mac": "aa:bb:cc:dd:ee:ff",
							"is-direct": false
						},
						{
							"mcast-ip": "ff02::fb",
							"stream-name": "-",
							"ap-mac": "aa:bb:cc:dd:ee:ff",
							"is-direct": false
						}
					]
				},
				{
					"client-mac": "68:db:f5:0f:84:18",
					"vlan-id": 800,
					"flex-mcast-client-group": [
						{
							"mcast-ip": "224.0.0.251",
							"stream-name": "-",
							"ap-mac": "aa:bb:cc:dd:ee:ff",
							"is-direct": false
						},
						{
							"mcast-ip": "ff02::fb",
							"stream-name": "-",
							"ap-mac": "aa:bb:cc:dd:ee:ff",
							"is-direct": false
						}
					]
				}
			]
		}`,
		// ListVLANL2MGIDs - VLAN Layer 2 multicast group ID operational data
		"Cisco-IOS-XE-wireless-mcast-oper:mcast-oper-data/vlan-l2-mgid-op": `{
			"Cisco-IOS-XE-wireless-mcast-oper:vlan-l2-mgid-op": [
				{
					"vlan-index": 1,
					"is-nonip-multicast-enabled": true,
					"is-broadcast-enable": true
				},
				{
					"vlan-index": 800,
					"is-nonip-multicast-enabled": true,
					"is-broadcast-enable": true
				},
				{
					"vlan-index": 801,
					"is-nonip-multicast-enabled": true,
					"is-broadcast-enable": true
				},
				{
					"vlan-index": 1002,
					"is-nonip-multicast-enabled": true,
					"is-broadcast-enable": true
				}
			]
		}`,

		// GetFabricMediastreamClientSummary - Empty data (not available in test environment)
		"Cisco-IOS-XE-wireless-mcast-oper:mcast-oper-data/fabric-media-stream-client-summary": `{}`,

		// GetMcastMgidInfo - Empty data (not available in test environment)
		"Cisco-IOS-XE-wireless-mcast-oper:mcast-oper-data/mcast-mgid-info": `{}`,

		// GetMulticastOperData - Empty data (not available in test environment)
		"Cisco-IOS-XE-wireless-mcast-oper:mcast-oper-data/multicast-oper-data": `{}`,

		"Cisco-IOS-XE-wireless-mcast-oper:mcast-oper-data/rrc-history-client-record-data": `{
			"Cisco-IOS-XE-wireless-mcast-oper:rrc-history-client-record-data": [
				{
					"user-time-stamp": "2023-09-22T15:30:00Z",
					"client-mac": "aa:bb:cc:dd:ee:ff",
					"decision": "admit",
					"reason-code": 0,
					"stream-name": "test-stream"
				}
			]
		}`,
		"Cisco-IOS-XE-wireless-mcast-oper:mcast-oper-data/rrc-sr-radio-record": `{
			"Cisco-IOS-XE-wireless-mcast-oper:rrc-sr-radio-record": [
				{
					"ap-mac": "11:22:33:44:55:66",
					"slot-id": 0,
					"radio-type": 1,
					"number-of-admitted": 2
				}
			]
		}`,
		"Cisco-IOS-XE-wireless-mcast-oper:mcast-oper-data/rrc-stream-record": `{
			"Cisco-IOS-XE-wireless-mcast-oper:rrc-stream-record": [
				{
					"stream-name-str": "test-stream",
					"group-ip": "224.0.1.1",
					"client-mac": "aa:bb:cc:dd:ee:ff",
					"decision": "admit"
				}
			]
		}`,
		"Cisco-IOS-XE-wireless-mcast-oper:mcast-oper-data/rrc-stream-admit-record": `{
			"Cisco-IOS-XE-wireless-mcast-oper:rrc-stream-admit-record": [
				{
					"last-updated": "2023-09-22T15:30:00Z",
					"client-mac": "aa:bb:cc:dd:ee:ff",
					"dest-ip": "224.0.1.1"
				}
			]
		}`,
		"Cisco-IOS-XE-wireless-mcast-oper:mcast-oper-data/rrc-stream-deny-record": `{
			"Cisco-IOS-XE-wireless-mcast-oper:rrc-stream-deny-record": [
				{
					"last-updated": "2023-09-22T15:30:00Z",
					"client-mac": "aa:bb:cc:dd:ee:ff",
					"dest-ip": "224.0.1.1"
				}
			]
		}`,
	}
	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
	defer mockServer.Close()

	// Create test client configured for the mock server
	testClient := testutil.NewTestClient(mockServer)
	service := mcast.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	t.Run("GetOperational", func(t *testing.T) {
		result, err := service.GetOperational(ctx)
		if err != nil {
			t.Errorf("Expected no error for mock GetOperational, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for mock GetOperational, got nil")
		}
	})

	t.Run("GetFlexConnectMediastreamClientSummary", func(t *testing.T) {
		result, err := service.GetFlexConnectMediastreamClientSummary(ctx)
		if err != nil {
			t.Errorf("Expected no error for mock GetFlexConnectMediastreamClientSummary, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for mock GetFlexConnectMediastreamClientSummary, got nil")
		}
	})

	t.Run("ListVLANL2MGIDs", func(t *testing.T) {
		result, err := service.ListVLANL2MGIDs(ctx)
		if err != nil {
			t.Errorf("Expected no error for mock ListVLANL2MGIDs, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for mock ListVLANL2MGIDs, got nil")
		}
	})

	t.Run("GetFabricMediastreamClientSummary", func(t *testing.T) {
		// This endpoint returns empty data in the test environment
		result, err := service.GetFabricMediastreamClientSummary(ctx)
		if err != nil {
			t.Errorf("Expected no error for mock GetFabricMediastreamClientSummary, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for mock GetFabricMediastreamClientSummary, got nil")
		}
	})

	t.Run("GetMcastMgidInfo", func(t *testing.T) {
		// This endpoint returns empty data in the test environment
		result, err := service.GetMcastMgidInfo(ctx)
		if err != nil {
			t.Errorf("Expected no error for mock GetMcastMgidInfo, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for mock GetMcastMgidInfo, got nil")
		}
	})

	t.Run("GetMulticastOperData", func(t *testing.T) {
		// This endpoint returns empty data in the test environment
		result, err := service.GetMulticastOperData(ctx)
		if err != nil {
			t.Errorf("Expected no error for mock GetMulticastOperData, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for mock GetMulticastOperData, got nil")
		}
	})

	// RRC (Resource Reservation Control) related functions
	t.Run("ListRrcHistoryClientRecordData", func(t *testing.T) {
		result, err := service.ListRrcHistoryClientRecordData(ctx)
		if err != nil {
			t.Errorf("Expected no error for mock ListRrcHistoryClientRecordData, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for mock ListRrcHistoryClientRecordData, got nil")
		}
	})

	t.Run("ListRrcSrRadioRecord", func(t *testing.T) {
		result, err := service.ListRrcSrRadioRecord(ctx)
		if err != nil {
			t.Errorf("Expected no error for mock ListRrcSrRadioRecord, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for mock ListRrcSrRadioRecord, got nil")
		}
	})

	t.Run("ListRrcStreamRecord", func(t *testing.T) {
		result, err := service.ListRrcStreamRecord(ctx)
		if err != nil {
			t.Errorf("Expected no error for mock ListRrcStreamRecord, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for mock ListRrcStreamRecord, got nil")
		}
	})

	t.Run("ListRrcStreamAdmitRecord", func(t *testing.T) {
		result, err := service.ListRrcStreamAdmitRecord(ctx)
		if err != nil {
			t.Errorf("Expected no error for mock ListRrcStreamAdmitRecord, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for mock ListRrcStreamAdmitRecord, got nil")
		}
	})

	t.Run("ListRrcStreamDenyRecord", func(t *testing.T) {
		result, err := service.ListRrcStreamDenyRecord(ctx)
		if err != nil {
			t.Errorf("Expected no error for mock ListRrcStreamDenyRecord, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for mock ListRrcStreamDenyRecord, got nil")
		}
	})
}

// TestMcastServiceUnit_GetOperations_ErrorHandling tests error scenarios using mock server.
func TestMcastServiceUnit_GetOperations_ErrorHandling(t *testing.T) {
	// Create mock server that returns 404 for multicast endpoints
	errorPaths := []string{
		"Cisco-IOS-XE-wireless-mcast-oper:mcast-oper-data",
		"Cisco-IOS-XE-wireless-mcast-oper:mcast-oper-data/flex-mediastream-client-summary",
		"Cisco-IOS-XE-wireless-mcast-oper:mcast-oper-data/vlan-l2-mgid-op",
		"Cisco-IOS-XE-wireless-mcast-oper:mcast-oper-data/fabric-media-stream-client-summary",
		"Cisco-IOS-XE-wireless-mcast-oper:mcast-oper-data/mcast-mgid-info",
		"Cisco-IOS-XE-wireless-mcast-oper:mcast-oper-data/multicast-oper-data",

		// Note: Not Verified on IOS-XE 17.12.5 - may return 404 errors on some controller versions.
		"Cisco-IOS-XE-wireless-mcast-oper:mcast-oper-data/rrc-history-client-record-data",
		"Cisco-IOS-XE-wireless-mcast-oper:mcast-oper-data/rrc-sr-radio-record",
		"Cisco-IOS-XE-wireless-mcast-oper:mcast-oper-data/rrc-stream-record",
		"Cisco-IOS-XE-wireless-mcast-oper:mcast-oper-data/rrc-stream-admit-record",
		"Cisco-IOS-XE-wireless-mcast-oper:mcast-oper-data/rrc-stream-deny-record",
	}
	mockServer := testutil.NewMockServer(testutil.WithErrorResponses(errorPaths, 404))
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := mcast.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	t.Run("GetOperational", func(t *testing.T) {
		_, err := service.GetOperational(ctx)
		if err == nil {
			t.Error("Expected error for 404 response, got nil")
		}
		if !core.IsNotFoundError(err) {
			t.Errorf("Expected NotFound error, got: %v", err)
		}
	})

	t.Run("GetFlexConnectMediastreamClientSummary", func(t *testing.T) {
		_, err := service.GetFlexConnectMediastreamClientSummary(ctx)
		if err == nil {
			t.Error("Expected error for 404 response, got nil")
		}
		if !core.IsNotFoundError(err) {
			t.Errorf("Expected NotFound error, got: %v", err)
		}
	})

	t.Run("ListVLANL2MGIDs", func(t *testing.T) {
		_, err := service.ListVLANL2MGIDs(ctx)
		if err == nil {
			t.Error("Expected error for 404 response, got nil")
		}
		if !core.IsNotFoundError(err) {
			t.Errorf("Expected NotFound error, got: %v", err)
		}
	})

	t.Run("GetFabricMediastreamClientSummary", func(t *testing.T) {
		_, err := service.GetFabricMediastreamClientSummary(ctx)
		if err == nil {
			t.Error("Expected error for 404 response, got nil")
		}
		if !core.IsNotFoundError(err) {
			t.Errorf("Expected NotFound error, got: %v", err)
		}
	})

	t.Run("GetMcastMgidInfo", func(t *testing.T) {
		_, err := service.GetMcastMgidInfo(ctx)
		if err == nil {
			t.Error("Expected error for 404 response, got nil")
		}
		if !core.IsNotFoundError(err) {
			t.Errorf("Expected NotFound error, got: %v", err)
		}
	})

	t.Run("GetMulticastOperData", func(t *testing.T) {
		_, err := service.GetMulticastOperData(ctx)
		if err == nil {
			t.Error("Expected error for 404 response, got nil")
		}
		if !core.IsNotFoundError(err) {
			t.Errorf("Expected NotFound error, got: %v", err)
		}
	})

	// Note: Not Verified on IOS-XE 17.12.5 - may return 404 errors on some controller versions.
	t.Run("ListRrcHistoryClientRecordData_404Error", func(t *testing.T) {
		_, err := service.ListRrcHistoryClientRecordData(ctx)
		if err == nil {
			t.Error("Expected error for 404 response, got nil")
		}
		if !core.IsNotFoundError(err) {
			t.Errorf("Expected NotFound error, got: %v", err)
		}
	})

	t.Run("ListRrcStreamRecord_404Error", func(t *testing.T) {
		_, err := service.ListRrcStreamRecord(ctx)
		if err == nil {
			t.Error("Expected error for 404 response, got nil")
		}
		if !core.IsNotFoundError(err) {
			t.Errorf("Expected NotFound error, got: %v", err)
		}
	})
}

// TestMcastServiceUnit_GetOperations_NilClient tests operations with nil client.
func TestMcastServiceUnit_GetOperations_NilClient(t *testing.T) {
	service := mcast.NewService(nil)
	ctx := testutil.TestContext(t)

	t.Run("GetOperational", func(t *testing.T) {
		_, err := service.GetOperational(ctx)
		if err == nil {
			t.Error("Expected error for nil client, got nil")
		}
	})

	t.Run("GetFlexConnectMediastreamClientSummary", func(t *testing.T) {
		_, err := service.GetFlexConnectMediastreamClientSummary(ctx)
		if err == nil {
			t.Error("Expected error for nil client, got nil")
		}
	})

	t.Run("ListVLANL2MGIDs", func(t *testing.T) {
		_, err := service.ListVLANL2MGIDs(ctx)
		if err == nil {
			t.Error("Expected error for nil client, got nil")
		}
	})

	t.Run("GetFabricMediastreamClientSummary", func(t *testing.T) {
		_, err := service.GetFabricMediastreamClientSummary(ctx)
		if err == nil {
			t.Error("Expected error for nil client, got nil")
		}
	})

	t.Run("GetMcastMgidInfo", func(t *testing.T) {
		_, err := service.GetMcastMgidInfo(ctx)
		if err == nil {
			t.Error("Expected error for nil client, got nil")
		}
	})

	t.Run("GetMulticastOperData", func(t *testing.T) {
		_, err := service.GetMulticastOperData(ctx)
		if err == nil {
			t.Error("Expected error for nil client, got nil")
		}
	})

	// Representative RRC function nil client tests
	t.Run("ListRrcHistoryClientRecordData_NilClient", func(t *testing.T) {
		_, err := service.ListRrcHistoryClientRecordData(ctx)
		if err == nil {
			t.Error("Expected error for nil client, got nil")
		}
	})

	t.Run("ListRrcStreamRecord_NilClient", func(t *testing.T) {
		_, err := service.ListRrcStreamRecord(ctx)
		if err == nil {
			t.Error("Expected error for nil client, got nil")
		}
	})
}
