package wnc

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"log/slog"
	"net/http"
	"net/url"
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

// Default request budgets. A request is bounded by all three, and WithTimeout sets only the first.
const (
	// DefaultTimeout is the default whole-request timeout (re-export of core.DefaultTimeout).
	DefaultTimeout = core.DefaultTimeout
	// DefaultResponseHeaderTimeout is the default budget for the response headers, five seconds,
	// which WithTimeout does not lift; raise it with WithResponseHeaderTimeout.
	DefaultResponseHeaderTimeout = core.DefaultResponseHeaderTimeout
	// DefaultTLSHandshakeTimeout is the default budget for the TLS handshake, five seconds, which
	// WithTimeout does not lift; raise it with WithTLSHandshakeTimeout.
	DefaultTLSHandshakeTimeout = core.DefaultTLSHandshakeTimeout
)

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

// Response is what Request returned: the controller's status and the body as it sent it
// (re-export of internal core.Response). An alias rather than a distinct type, so a caller can
// name it in a variable, a struct field or a test double.
type Response = core.Response

// Client represents the unified WNC API client with access to all domain services.
// This provides a single-import approach to accessing all wireless controller functionality.
type Client struct {
	core *core.Client // Core client that handles HTTP communication
}

// NewClient creates a new unified WNC client with the specified host, token, and options.
// This is the main entry point for all wireless controller operations.
//
// host is an authority and nothing else — "wnc1.example.internal" or "192.0.2.10:443". A scheme, a
// path, a query, a fragment, userinfo or an IPv6 zone id is refused with ErrInvalidConfiguration
// rather than concatenated into a URL that reads another node.
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

// WithTimeout sets the whole-request timeout (re-export wrapper). It lifts neither
// DefaultResponseHeaderTimeout nor DefaultTLSHandshakeTimeout, so a caller that raises this
// alone is still capped at five seconds for the headers, which is when a busy controller is
// slowest. Raise those with WithResponseHeaderTimeout and WithTLSHandshakeTimeout.
func WithTimeout(d time.Duration) Option { return core.WithTimeout(d) }

// WithInsecureSkipVerify controls TLS certificate verification (lab/testing only).
func WithInsecureSkipVerify(skip bool) Option { return core.WithInsecureSkipVerify(skip) }

// WithRootCAs verifies the controller's certificate against pool instead of the host's roots
// (re-export wrapper). Prefer it to WithInsecureSkipVerify where the controller presents a
// certificate from a private CA: the certificate is then verified rather than unverified.
func WithRootCAs(pool *x509.CertPool) Option { return core.WithRootCAs(pool) }

// WithClientCertificate presents cert to the controller (re-export wrapper), for a deployment
// that authenticates the client with mTLS as well as with the Authorization header.
func WithClientCertificate(cert tls.Certificate) Option { return core.WithClientCertificate(cert) }

// WithProxy routes requests through the proxy the resolver returns (re-export wrapper).
func WithProxy(fn func(*http.Request) (*url.URL, error)) Option { return core.WithProxy(fn) }

// WithResponseHeaderTimeout bounds the wait for the response headers (re-export wrapper).
func WithResponseHeaderTimeout(d time.Duration) Option { return core.WithResponseHeaderTimeout(d) }

// WithTLSHandshakeTimeout bounds the TLS handshake (re-export wrapper).
func WithTLSHandshakeTimeout(d time.Duration) Option { return core.WithTLSHandshakeTimeout(d) }

// WithLogger sets a custom slog.Logger.
func WithLogger(l *slog.Logger) Option { return core.WithLogger(l) }

// WithUserAgent sets a custom User-Agent header value.
func WithUserAgent(ua string) Option { return core.WithUserAgent(ua) }

// GetOption customizes a single GET request (re-export of internal core.GetOption).
type GetOption = core.GetOption

// DefaultsMode selects the RFC 6243 retrieval mode for WithDefaults.
type DefaultsMode = core.DefaultsMode

const (
	// ReportAll materializes the leaves in force at their schema default.
	ReportAll = core.DefaultsReportAll
	// Explicit returns the leaves a client set, including any set to their schema default.
	Explicit = core.DefaultsExplicit
)

// WithDefaults requests the given with-defaults retrieval mode (RFC 8040 4.8.9).
// Scope it to the container that needs it: on a whole-container read the added
// leaves accumulate across every nested container.
func WithDefaults(mode DefaultsMode) GetOption { return core.WithDefaults(mode) }

// WithFields limits the answer to an RFC 8040 4.8.3 fields expression (re-export wrapper).
// A pruned leaf is absent, and an absent leaf decodes to a zero value, so prune only
// the fields the caller reads.
func WithFields(expression string) GetOption { return core.WithFields(expression) }

// WithDepth limits the answer to an RFC 8040 4.8.2 subtree depth (re-export wrapper).
// A node the limit cuts is absent exactly as a pruned leaf is, so an absent leaf still
// decodes to a zero value; bound the depth to what the caller reads.
func WithDepth(levels int) GetOption { return core.WithDepth(levels) }

// GetData reads a RESTCONF data path this package has no typed accessor for and returns the
// body as received. The /restconf/data prefix is optional and GetOption values apply as they
// do to a typed read.
//
// Three things the body does not say for itself. The response carries exactly one top-level
// key, the module-qualified name of the node requested, so check that key rather than trusting
// a struct tag: a tag naming a key the controller did not send decodes to nothing and reports
// success. A node holding nothing answers with no body, so the slice is non-nil and empty with
// a nil error — check the length before decoding. The path is sent as given, so a caller keying
// into a list escapes the key itself, an unescaped "#" or "?" ending the path early and reading
// a different node without error.
func (c *Client) GetData(ctx context.Context, path string, opts ...GetOption) ([]byte, error) {
	return core.GetRaw(ctx, c.core, path, opts...)
}

// GetDataInto reads a RESTCONF data path this package has no typed accessor for and decodes it
// into T, applying the envelope check every typed accessor gets: the response must carry exactly
// one top-level key, module-qualified and naming the node the path asked for, and T must declare a
// field for that key. GetData leaves all of that to the caller.
//
// T is the envelope type, so it must be a struct whose outermost tag is the module-qualified node
// name — the shape every Cisco…Data type in this module's service packages has. A map or any other
// non-struct is refused, because the check asks whether T can consume the key rather than trusting
// it to. The check is top-level only: a tag below the top naming a node the response does not
// carry still decodes to nothing.
//
// It is a function rather than a method on Client because a generic method, which this toolchain
// does accept, may not be declared in an interface and is invisible to reflect — so a consumer
// could neither put this behind a seam of its own nor reach it by reflection.
func GetDataInto[T any](ctx context.Context, c *Client, path string, opts ...GetOption) (*T, error) {
	if c == nil {
		return core.Get[T](ctx, nil, path, opts...)
	}

	return core.Get[T](ctx, c.core, path, opts...)
}

// The untyped request methods below are this package's escape hatch. They exist because the
// controller's schema moves between releases: a node or an operation this package has no typed
// accessor for still has to be reachable without waiting for one. Each returns the body as the
// controller sent it, and each reports an *APIError the way every typed accessor does.
//
// A payload is marshaled with encoding/json, as every typed write in this package marshals its
// own, so an SDK config struct or a map may be passed directly. A []byte or a json.RawMessage is
// the exception: it is sent as written, having been checked for well-formed JSON, which is how a
// body read with GetData is edited and sent back. A nil payload sends no body and no Content-Type.
//
// A write is answered with the edited node, with an RPC-style output, or with nothing at all, so
// an empty return with a nil error is a success.

// PostData creates a node under a RESTCONF data path (RFC 8040 4.4).
func (c *Client) PostData(ctx context.Context, path string, payload any) ([]byte, error) {
	return core.EditRaw(ctx, c.core, http.MethodPost, path, payload)
}

// PutData replaces a node at a RESTCONF data path (RFC 8040 4.5). A typed struct carrying
// omitempty marshals fewer leaves than it decoded, so replacing a node with one removes the rest.
func (c *Client) PutData(ctx context.Context, path string, payload any) ([]byte, error) {
	return core.EditRaw(ctx, c.core, http.MethodPut, path, payload)
}

// PatchData merges a payload into a node at a RESTCONF data path. This package sends
// application/yang-data+json, so the edit is the plain patch of RFC 8040 4.6.1: a leaf absent from
// the payload is left alone, and no payload deletes anything.
//
// A typed struct cannot clear a leaf, because encoding/json drops a zero field carrying omitempty
// before the payload is built, and its absence then means "leave alone". Send the leaf as bytes to
// set it to its zero, and DeleteData to remove it.
func (c *Client) PatchData(ctx context.Context, path string, payload any) ([]byte, error) {
	return core.EditRaw(ctx, c.core, http.MethodPatch, path, payload)
}

// DeleteData removes a node at a RESTCONF data path (RFC 8040 4.7).
func (c *Client) DeleteData(ctx context.Context, path string) ([]byte, error) {
	return core.EditRaw(ctx, c.core, http.MethodDelete, path, nil)
}

// PostRPC invokes an operation on a RESTCONF operations path (RFC 8040 3.6 and 4.4.2). The path
// is the RPC name, module-qualified as the controller publishes it, with or without the
// /restconf/operations prefix, and the payload is normally an object under a single "input" key.
func (c *Client) PostRPC(ctx context.Context, path string, payload any) ([]byte, error) {
	return core.CallRPCRaw(ctx, c.core, path, payload)
}

// Request performs a request with the given method on the path as the caller wrote it, for
// whatever the verb methods above cannot express: a method RESTCONF gains later, a bodiless probe
// such as HEAD, or a query parameter this package has no option for.
//
// A path already under /restconf/operations is sent to the operations root and anything else to
// the data root, which passes a /restconf/data-prefixed path through and prefixes a bare one.
//
// On the data root the method is sent as given and is checked against neither the path nor the
// payload; the one value rejected is the empty string, which net/http reads as GET. The operations
// root takes POST alone, and another method there is refused rather than replaced: this package
// would send POST regardless, invoking the operation instead of doing what was asked.
//
// This is the one method here that returns the status as well as the body, because it is the one
// with no fixed verb: 201, 204 and an empty 200 all answer with no body, so the body alone cannot
// say whether the node held nothing, was created or was replaced. The Response is non-nil exactly
// when the error is nil, and a status of 400 or above arrives as an *APIError rather than in it.
func (c *Client) Request(ctx context.Context, method, path string, payload any) (*Response, error) {
	return core.RequestRaw(ctx, c.core, method, path, payload)
}

// CloseIdleConnections closes the pooled connections that have no request on them, releasing the
// sockets a long-lived process would otherwise hold open after its last read. A connection in use
// is left alone and the client stays usable afterwards: the next request dials again.
func (c *Client) CloseIdleConnections() {
	c.core.CloseIdleConnections()
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

// AWIPS returns the Automated Wireless Intrusion Prevention System service.
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
// EXPERIMENTAL: Requires IOS-XE 17.15.1+.
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
