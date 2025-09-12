package service

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/helper"
)

func TestServiceBaseUnit_Constructor_Success(t *testing.T) {
	t.Run("NewBaseService", func(t *testing.T) {
		t.Run("WithValidClient", func(t *testing.T) {
			// Create mock client instead of real client
			mockClient, _ := core.New("localhost", "test-token", core.WithInsecureSkipVerify(true))
			service := NewBaseService(mockClient)

			helper.AssertPointerEquals(t, service.client, mockClient, "Expected client to match mockClient")
		})

		t.Run("WithNilClient", func(t *testing.T) {
			service := NewBaseService(nil)

			helper.AssertPointerNil(t, service.client, "Expected client to be nil")
		})
	})

	t.Run("Client", func(t *testing.T) {
		t.Run("ReturnsValidClient", func(t *testing.T) {
			// Create mock client instead of real client
			mockClient, _ := core.New("localhost", "test-token", core.WithInsecureSkipVerify(true))
			service := NewBaseService(mockClient)

			retrievedClient := service.Client()
			helper.AssertPointerEquals(t, retrievedClient, mockClient, "Expected retrieved client to match mockClient")
		})

		t.Run("ReturnsNilClient", func(t *testing.T) {
			service := NewBaseService(nil)

			retrievedClient := service.Client()
			helper.AssertPointerNil(t, retrievedClient, "Expected retrieved client to be nil")
		})
	})

	t.Run("ServiceIntegration", func(t *testing.T) {
		// Test that BaseService can be embedded and used properly
		type TestService struct {
			BaseService
		}

		newTestService := func(client *core.Client) TestService {
			return TestService{
				BaseService: NewBaseService(client),
			}
		}

		t.Run("EmbeddingWithValidClient", func(t *testing.T) {
			// Create mock client instead of real client
			mockClient, _ := core.New("localhost", "test-token", core.WithInsecureSkipVerify(true))
			testService := newTestService(mockClient)

			helper.AssertPointerEquals(t, testService.Client(), mockClient,
				"Expected embedded client to match mockClient")
		})

		t.Run("EmbeddingWithNilClient", func(t *testing.T) {
			testService := newTestService(nil)

			helper.AssertPointerNil(t, testService.Client(), "Expected embedded client to be nil")
		})
	})
}
