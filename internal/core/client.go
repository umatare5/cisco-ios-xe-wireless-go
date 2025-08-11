// Package core provides the core client for the Cisco Wireless Network Controller API.
// This package implements the three-layer architecture with Core, Domain Service, and Generated Type separation.
package core

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/httpx"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/validation"
)

// Default timeout constant
const (
	// DefaultTimeout is the default timeout for API requests
	DefaultTimeout = 60 * time.Second
)

// Client represents the core WNC API client with connection pooling and structured logging
type Client struct {
	httpClient *http.Client      // Reused HTTP client with connection pool
	rest       *restconf.Builder // RESTCONF URL builder
	logger     *slog.Logger      // Structured logger
	token      string            // Access token for authorization
}

// Option represents a functional option for configuring the Client
type Option func(*Client) error

// WithTimeout sets the timeout duration for HTTP requests
func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) error {
		if timeout <= 0 {
			return fmt.Errorf("timeout must be positive, got %v", timeout)
		}
		c.httpClient.Timeout = timeout
		return nil
	}
}

// WithInsecureSkipVerify configures TLS certificate verification
func WithInsecureSkipVerify(skip bool) Option {
	return func(c *Client) error {
		// Create new transport with updated TLS settings
		transport := httpx.NewTransport(skip)
		c.httpClient.Transport = transport
		return nil
	}
}

// WithLogger sets a custom logger for the client
func WithLogger(logger *slog.Logger) Option {
	return func(c *Client) error {
		if logger == nil {
			return errors.New("logger cannot be nil")
		}
		c.logger = logger
		return nil
	}
}

// WithUserAgent sets a custom User-Agent header
func WithUserAgent(userAgent string) Option {
	return func(c *Client) error {
		// This will be handled in the headers when making requests
		// For now, we store it in the client context (not implemented yet)
		return nil
	}
}

// New creates a new WNC client with the specified host, token, and options
func New(host, token string, opts ...Option) (*Client, error) {
	// Validate inputs using existing validation functions
	if !validation.IsValidController(host) {
		return nil, fmt.Errorf("invalid controller address: %s", host)
	}
	if !validation.IsValidAccessToken(token) {
		return nil, errors.New("invalid access token")
	} // Create HTTP transport with default settings
	transport := httpx.NewTransport(false) // Default to secure

	// Create HTTP client with transport
	httpClient := &http.Client{
		Transport: transport,
		Timeout:   DefaultTimeout,
	}

	// Create RESTCONF URL builder
	restBuilder := restconf.NewBuilder(restconf.DefaultProtocol, host)

	// Create client with defaults
	client := &Client{
		httpClient: httpClient,
		rest:       restBuilder,
		logger:     slog.Default(),
		token:      token,
	}

	// Apply options
	for _, opt := range opts {
		if err := opt(client); err != nil {
			return nil, fmt.Errorf("failed to apply option: %w", err)
		}
	}

	return client, nil
}

// Do performs a generic HTTP request to the specified path and unmarshals the response into out
// Do executes an HTTP request and unmarshals the response
func (c *Client) Do(ctx context.Context, method, path string, out any) error {
	if err := c.validateDoParameters(ctx, out); err != nil {
		return err
	}

	req, err := c.createRequest(ctx, method, path)
	if err != nil {
		return err
	}

	resp, err := c.executeRequest(req)
	if err != nil {
		return err
	}
	defer c.closeResponseBody(resp)

	body, err := c.readResponseBody(resp)
	if err != nil {
		return err
	}

	if err := c.checkHTTPErrors(resp, body); err != nil {
		return err
	}

	return c.unmarshalResponse(body, out, path)
}

// validateDoParameters validates input parameters for the Do method
func (c *Client) validateDoParameters(ctx context.Context, out any) error {
	if c == nil {
		return errors.New("client cannot be nil")
	}
	if ctx == nil {
		return errors.New("context cannot be nil")
	}
	if out == nil {
		return errors.New("output parameter cannot be nil")
	}
	return nil
}

// createRequest creates and configures an HTTP request
func (c *Client) createRequest(ctx context.Context, method, path string) (*http.Request, error) {
	url := c.rest.BuildRESTCONFURL(path)

	req, err := http.NewRequestWithContext(ctx, method, url, http.NoBody)
	if err != nil {
		c.logger.Error("Failed to create HTTP request", "error", err, "url", url)
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header = httpx.DefaultHeaders(c.token)
	c.logger.Debug("Sending API request", "method", method, "url", url)
	return req, nil
}

// executeRequest executes the HTTP request
func (c *Client) executeRequest(req *http.Request) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		c.logger.Error("HTTP request failed", "error", err, "url", req.URL.String())
		return nil, fmt.Errorf("request failed: %w", err)
	}
	return resp, nil
}

// closeResponseBody safely closes the response body with error logging
func (c *Client) closeResponseBody(resp *http.Response) {
	if closeErr := resp.Body.Close(); closeErr != nil {
		c.logger.Error("Failed to close response body", "error", closeErr)
	}
}

// readResponseBody reads the complete response body
func (c *Client) readResponseBody(resp *http.Response) ([]byte, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.logger.Error("Failed to read response body", "error", err)
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	c.logger.Debug("Received API response", "status", resp.StatusCode, "content_length", len(body))
	return body, nil
}

// checkHTTPErrors validates HTTP status codes and returns appropriate errors
func (c *Client) checkHTTPErrors(resp *http.Response, body []byte) error {
	if resp.StatusCode >= 400 {
		c.logger.Error("HTTP error response", "status", resp.StatusCode, "body", string(body))
		return &APIError{
			StatusCode: resp.StatusCode,
			Message:    string(body),
			Body:       body,
		}
	}
	return nil
}

// unmarshalResponse unmarshals the JSON response into the output parameter
func (c *Client) unmarshalResponse(body []byte, out any, path string) error {
	if err := json.Unmarshal(body, out); err != nil {
		c.logger.Error("Failed to unmarshal JSON response", "error", err, "body_length", len(body))
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.Debug("Successfully processed API response", "path", path)
	return nil
}

// Domain service interfaces - these will be implemented by the respective packages

// AFCService defines the interface for Automated Frequency Coordination operations
type AFCService interface {
	// Methods will be added as the afc package is refactored
}

// APService defines the interface for Access Point operations
type APService interface {
	// Methods will be added as the ap package is refactored
}

// ClientService defines the interface for wireless client operations
type ClientService interface {
	// Methods will be added as the client package is refactored
}

// GeneralService defines the interface for general controller operations
type GeneralService interface {
	// Methods will be added as the general package is refactored
}

// RRMService defines the interface for Radio Resource Management operations
type RRMService interface {
	// Methods will be added as the rrm package is refactored
}

// WLANService defines the interface for WLAN operations
type WLANService interface {
	// Methods will be added as the wlan package is refactored
}

// RogueService defines the interface for rogue access point detection operations
type RogueService interface {
	// Methods will be added as the rogue package is refactored
}

// NMSPService defines the interface for Network Mobility Services Protocol operations
type NMSPService interface {
	// Methods will be added as the nmsp package is refactored
}

// HyperlocationService defines the interface for hyperlocation operations
type HyperlocationService interface {
	// Methods will be added as the hyperlocation package is refactored
}

// MDNSService defines the interface for multicast DNS operations
type MDNSService interface {
	// Methods will be added as the mdns package is refactored
}

// GeolocationService defines the interface for geolocation operations
type GeolocationService interface {
	// Methods will be added as the geolocation package is refactored
}

// McastService defines the interface for multicast operations
type McastService interface {
	// Methods will be added as the mcast package is refactored
}

// APFService defines the interface for Application Policy Framework operations
type APFService interface {
	// Methods will be added as the apf package is refactored
}

// AWIPSService defines the interface for Advanced Weather Interactive Processing System operations
type AWIPSService interface {
	// Methods will be added as the awips package is refactored
}

// BLEService defines the interface for Bluetooth Low Energy operations
type BLEService interface {
	// Methods will be added as the ble package is refactored
}

// CTSService defines the interface for Cisco TrustSec operations
type CTSService interface {
	// Methods will be added as the cts package is refactored
}

// Dot11Service defines the interface for 802.11 wireless standard operations
type Dot11Service interface {
	// Methods will be added as the dot11 package is refactored
}

// Dot15Service defines the interface for 802.15 standard operations
type Dot15Service interface {
	// Methods will be added as the dot15 package is refactored
}

// FabricService defines the interface for Fabric operations
type FabricService interface {
	// Methods will be added as the fabric package is refactored
}

// FlexService defines the interface for FlexConnect operations
type FlexService interface {
	// Methods will be added as the flex package is refactored
}

// LocationService defines the interface for Location services operations
type LocationService interface {
	// Methods will be added as the location package is refactored
}

// RadioService defines the interface for Radio operations
type RadioService interface {
	// Methods will be added as the radio package is refactored
}

// RFService defines the interface for Radio Frequency operations
type RFService interface {
	// Methods will be added as the rf package is refactored
}

// RFIDService defines the interface for RFID operations
type RFIDService interface {
	// Methods will be added as the rfid package is refactored
}

// MobilityService defines the interface for Mobility operations
type MobilityService interface {
	// Methods will be added as the mobility package is refactored
}

// MeshService defines the interface for Mesh operations
type MeshService interface {
	// Methods will be added as the mesh package is refactored
}

// SiteService defines the interface for Site operations
type SiteService interface {
	// Methods will be added as the site package is refactored
}

// LISPService defines the interface for LISP operations
type LISPService interface {
	// Methods will be added as the lisp package is refactored
}

// Domain service accessors - these create service instances that use the client's Do() method

// AFC returns an AFC service instance
func (c *Client) AFC() AFCService {
	// Import the AFC service from the afc package to avoid circular dependencies
	// For now, we return nil to maintain compatibility
	return nil // This will be implemented with a wrapper or interface approach
}

// AP returns an Access Point service instance
func (c *Client) AP() APService {
	// This will be implemented when AP package is refactored to use the new client
	return nil // Placeholder
}

// Client returns a wireless client service instance
func (c *Client) Client() ClientService {
	// This will be implemented when client package is refactored to use the new client
	return nil // Placeholder
}

// General returns a general controller service instance
func (c *Client) General() GeneralService {
	// This will be implemented when general package is refactored to use the new client
	return nil // Placeholder
}

// RRM returns a Radio Resource Management service instance
func (c *Client) RRM() RRMService {
	// This will be implemented when RRM package is refactored to use the new client
	return nil // Placeholder
}

// WLAN returns a WLAN service instance
func (c *Client) WLAN() WLANService {
	// This will be implemented when WLAN package is refactored to use the new client
	return nil // Placeholder
}

// Rogue returns a rogue access point detection service instance
func (c *Client) Rogue() RogueService {
	// This will be implemented when rogue package is refactored to use the new client
	return nil // Placeholder
}

// NMSP returns a Network Mobility Services Protocol service instance
func (c *Client) NMSP() NMSPService {
	// This will be implemented when nmsp package is refactored to use the new client
	return nil // Placeholder
}

// Hyperlocation returns a hyperlocation service instance
func (c *Client) Hyperlocation() HyperlocationService {
	// This will be implemented when hyperlocation package is refactored to use the new client
	return nil // Placeholder
}

// MDNS returns a multicast DNS service instance
func (c *Client) MDNS() MDNSService {
	// This will be implemented when mdns package is refactored to use the new client
	return nil // Placeholder
}

// Geolocation returns a geolocation service instance
func (c *Client) Geolocation() GeolocationService {
	// This will be implemented when geolocation package is refactored to use the new client
	return nil // Placeholder
}

// Mcast returns a multicast service instance
func (c *Client) Mcast() McastService {
	// This will be implemented when mcast package is refactored to use the new client
	return nil // Placeholder
}

// APF returns an APF service instance
func (c *Client) APF() APFService {
	// This will be implemented when apf package is refactored to use the new client
	return nil // Placeholder
}

// AWIPS returns an AWIPS service instance
func (c *Client) AWIPS() AWIPSService {
	// This will be implemented when awips package is refactored to use the new client
	return nil // Placeholder
}

// BLE returns a BLE service instance
func (c *Client) BLE() BLEService {
	// This will be implemented when ble package is refactored to use the new client
	return nil // Placeholder
}

// CTS returns a CTS service instance
func (c *Client) CTS() CTSService {
	// This will be implemented when cts package is refactored to use the new client
	return nil // Placeholder
}

// Dot11 returns a 802.11 service instance
func (c *Client) Dot11() Dot11Service {
	// This will be implemented when dot11 package is refactored to use the new client
	return nil // Placeholder
}

// Dot15 returns a 802.15 service instance
func (c *Client) Dot15() Dot15Service {
	// This will be implemented when dot15 package is refactored to use the new client
	return nil // Placeholder
}

// Fabric returns a Fabric service instance
func (c *Client) Fabric() FabricService {
	// This will be implemented when fabric package is refactored to use the new client
	return nil // Placeholder
}

// Flex returns a FlexConnect service instance
func (c *Client) Flex() FlexService {
	// This will be implemented when flex package is refactored to use the new client
	return nil // Placeholder
}

// Location returns a Location service instance
func (c *Client) Location() LocationService {
	// This will be implemented when location package is refactored to use the new client
	return nil // Placeholder
}

// Radio returns a Radio service instance
func (c *Client) Radio() RadioService {
	// This will be implemented when radio package is refactored to use the new client
	return nil // Placeholder
}

// RF returns a Radio Frequency service instance
func (c *Client) RF() RFService {
	// This will be implemented when rf package is refactored to use the new client
	return nil // Placeholder
}

// RFID returns an RFID service instance
func (c *Client) RFID() RFIDService {
	// This will be implemented when rfid package is refactored to use the new client
	return nil // Placeholder
}

// Mobility returns a Mobility service instance
func (c *Client) Mobility() MobilityService {
	// This will be implemented when mobility package is refactored to use the new client
	return nil // Placeholder
}

// Mesh returns a Mesh service instance
func (c *Client) Mesh() MeshService {
	// This will be implemented when mesh package is refactored to use the new client
	return nil // Placeholder
}

// Site returns a Site service instance
func (c *Client) Site() SiteService {
	// This will be implemented when site package is refactored to use the new client
	return nil // Placeholder
}

// LISP returns a LISP service instance
func (c *Client) LISP() LISPService {
	// This will be implemented when lisp package is refactored to use the new client
	return nil // Placeholder
}
