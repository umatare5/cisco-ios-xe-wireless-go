// Package wnc provides a unified client for the Cisco Wireless Network Controller API.
package wnc

import (
	"log/slog"
	"time"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/afc"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/ap"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/apf"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/awips"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/ble"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/controller"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/cts"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/dot11"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/dot15"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/fabric"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/flex"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/general"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/geolocation"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/hyperlocation"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/lisp"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/location"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/mcast"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/mdns"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/mesh"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/mobility"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/nmsp"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/radio"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/rf"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/rfid"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/rogue"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/rrm"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/site"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/spaces"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/urwb"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/wat"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/wlan"
)

// DefaultTimeout is the default request timeout (re-export of core.DefaultTimeout).
const DefaultTimeout = core.DefaultTimeout

// Error sentinels re-exported for consumer side error handling with errors.Is.
var (
	ErrAuthenticationFailed = core.ErrAuthenticationFailed
	ErrAccessForbidden      = core.ErrAccessForbidden
	ErrResourceNotFound     = core.ErrResourceNotFound
	ErrInvalidConfiguration = core.ErrInvalidConfiguration
	ErrRequestTimeout       = core.ErrRequestTimeout
)

// APIError is returned for HTTP error responses (type alias to preserve instanceof semantics with errors.As).
type APIError = core.APIError

// Client represents the unified WNC API client with access to all domain services.
// This provides a single-import approach to accessing all wireless controller functionality.
type Client struct {
	core *core.Client // Core client that handles HTTP communication
}

// NewClient creates a new unified WNC client with the specified host, token, and options.
// This is the main entry point for all wireless controller operations.
func NewClient(host, token string, opts ...Option) (*Client, error) {
	coreClient, err := core.New(host, token, opts...)
	if err != nil {
		return nil, err
	}
	return &Client{core: coreClient}, nil
}

// Option is a functional option for configuring the unified client (re-export of internal core.Option).
// This allows end users to supply options without importing the internal/core package.
type Option = core.Option

// WithTimeout sets the request timeout (re-export wrapper).
func WithTimeout(d time.Duration) Option { return core.WithTimeout(d) }

// WithInsecureSkipVerify controls TLS certificate verification (lab/testing only).
func WithInsecureSkipVerify(skip bool) Option { return core.WithInsecureSkipVerify(skip) }

// WithLogger sets a custom slog.Logger.
func WithLogger(l *slog.Logger) Option { return core.WithLogger(l) }

// WithUserAgent sets a custom User-Agent header value.
func WithUserAgent(ua string) Option { return core.WithUserAgent(ua) }

// Core returns the underlying core.Client for advanced use cases.
// This should typically not be needed for normal usage.
func (c *Client) Core() *core.Client {
	return c.core
}

// Domain service accessors - each returns a service instance for the respective domain

// AFC returns the Automated Frequency Coordination service.
func (c *Client) AFC() afc.Service {
	return afc.NewService(c.core)
}

// AP returns the Access Point service.
func (c *Client) AP() ap.Service {
	return ap.NewService(c.core)
}

// APF returns the Application Policy Framework service.
func (c *Client) APF() apf.Service {
	return apf.NewService(c.core)
}

// AWIPS returns the Advanced Weather Interactive Processing System service.
func (c *Client) AWIPS() awips.Service {
	return awips.NewService(c.core)
}

// BLE returns the Bluetooth Low Energy service.
func (c *Client) BLE() ble.Service {
	return ble.NewService(c.core)
}

// Client returns the wireless client service.
func (c *Client) Client() client.Service {
	return client.NewService(c.core)
}

// Controller returns the controller management service.
func (c *Client) Controller() controller.Service {
	return controller.NewService(c.core)
}

// CTS returns the Cisco TrustSec service.
func (c *Client) CTS() cts.Service {
	return cts.NewService(c.core)
}

// Dot11 returns the 802.11 wireless standard service.
func (c *Client) Dot11() dot11.Service {
	return dot11.NewService(c.core)
}

// Dot15 returns the 802.15 standard service.
func (c *Client) Dot15() dot15.Service {
	return dot15.NewService(c.core)
}

// Fabric returns the Fabric service.
func (c *Client) Fabric() fabric.Service {
	return fabric.NewService(c.core)
}

// Flex returns the FlexConnect service.
func (c *Client) Flex() flex.Service {
	return flex.NewService(c.core)
}

// General returns the general controller service.
func (c *Client) General() general.Service {
	return general.NewService(c.core)
}

// Geolocation returns the geolocation service.
func (c *Client) Geolocation() geolocation.Service {
	return geolocation.NewService(c.core)
}

// Hyperlocation returns the hyperlocation service.
func (c *Client) Hyperlocation() hyperlocation.Service {
	return hyperlocation.NewService(c.core)
}

// LISP returns the LISP service.
func (c *Client) LISP() lisp.Service {
	return lisp.NewService(c.core)
}

// Location returns the location services service.
func (c *Client) Location() location.Service {
	return location.NewService(c.core)
}

// Mcast returns the multicast service.
func (c *Client) Mcast() mcast.Service {
	return mcast.NewService(c.core)
}

// MDNS returns the multicast DNS service.
func (c *Client) MDNS() mdns.Service {
	return mdns.NewService(c.core)
}

// Mesh returns the mesh networking service.
func (c *Client) Mesh() mesh.Service {
	return mesh.NewService(c.core)
}

// Mobility returns the mobility management service.
func (c *Client) Mobility() mobility.Service {
	return mobility.NewService(c.core)
}

// NMSP returns the Network Mobility Services Protocol service.
func (c *Client) NMSP() nmsp.Service {
	return nmsp.NewService(c.core)
}

// Radio returns the radio management service.
func (c *Client) Radio() radio.Service {
	return radio.NewService(c.core)
}

// RF returns the Radio Frequency management service.
func (c *Client) RF() rf.Service {
	return rf.NewService(c.core)
}

// RFID returns the RFID service.
func (c *Client) RFID() rfid.Service {
	return rfid.NewService(c.core)
}

// Rogue returns the rogue access point detection service.
func (c *Client) Rogue() rogue.Service {
	return rogue.NewService(c.core)
}

// RRM returns the Radio Resource Management service.
func (c *Client) RRM() rrm.Service {
	return rrm.NewService(c.core)
}

// Site returns the site management service.
func (c *Client) Site() site.Service {
	return site.NewService(c.core)
}

// Spaces returns the Cisco Spaces integration service.
// EXPERIMENTAL: Requires IOS-XE 17.18.1+.
func (c *Client) Spaces() spaces.Service {
	return spaces.NewService(c.core)
}

// URWB returns the Ultra Reliable Wireless Backhaul service.
// EXPERIMENTAL: Requires IOS-XE 17.18.1+.
func (c *Client) URWB() urwb.Service {
	return urwb.NewService(c.core)
}

// WAT returns the Wireless Application Templates service.
// EXPERIMENTAL: Requires IOS-XE 17.18.1+.
func (c *Client) WAT() wat.Service {
	return wat.NewService(c.core)
}

// WLAN returns the WLAN configuration service.
func (c *Client) WLAN() wlan.Service {
	return wlan.NewService(c.core)
}

// Tag service accessors - provide direct access to tag management services

// PolicyTag returns the Policy Tag service for policy tag management operations.
// This provides direct access to policy tag CRUD operations without going through WLAN service.
func (c *Client) PolicyTag() *wlan.PolicyTagService {
	return wlan.NewPolicyTagService(c.core)
}

// RFTag returns the RF Tag service for RF tag management operations.
// This provides direct access to RF tag CRUD operations without going through RF service.
func (c *Client) RFTag() *rf.RFTagService {
	return rf.NewRFTagService(c.core)
}

// SiteTag returns the Site Tag service for site tag management operations.
// This provides direct access to site tag CRUD operations without going through Site service.
func (c *Client) SiteTag() *site.SiteTagService {
	return site.NewSiteTagService(c.core)
}
