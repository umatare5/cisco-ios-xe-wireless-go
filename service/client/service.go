package client

import (
	"context"
	"errors"
	"fmt"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/validation"
)

// Service provides wireless client operations for Cisco IOS-XE Wireless LAN Controller.
type Service struct {
	service.BaseService
}

// NewService creates a new Client service instance with the provided client.
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}

// GetOperational retrieves the complete client operational data.
func (s Service) GetOperational(ctx context.Context, opts ...core.GetOption) (*CiscoIOSXEWirelessClientOper, error) {
	return core.Get[CiscoIOSXEWirelessClientOper](ctx, s.Client(), routes.ClientOperPath, opts...)
}

// ListCommonInfo retrieves common operational data for clients.
func (s Service) ListCommonInfo(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessClientOperCommonOperData, error) {
	return core.Get[CiscoIOSXEWirelessClientOperCommonOperData](
		ctx,
		s.Client(),
		routes.ClientCommonOperDataPath,
		opts...,
	)
}

// GetCommonInfoByMAC retrieves client operational data filtered by MAC address.
func (s Service) GetCommonInfoByMAC(
	ctx context.Context, mac string, opts ...core.GetOption,
) (*CiscoIOSXEWirelessClientOperCommonOperData, error) {
	normalizedMAC, err := service.RequireMACAddress(mac)
	if err != nil {
		return nil, err
	}

	url := s.Client().RESTCONFBuilder().BuildQueryURL(routes.ClientCommonOperDataPath, normalizedMAC)
	return core.Get[CiscoIOSXEWirelessClientOperCommonOperData](ctx, s.Client(), url, opts...)
}

// ListDCInfo retrieves discovery client information.
func (s Service) ListDCInfo(ctx context.Context, opts ...core.GetOption) (*CiscoIOSXEWirelessClientOperDcInfo, error) {
	return core.Get[CiscoIOSXEWirelessClientOperDcInfo](ctx, s.Client(), routes.ClientDcInfoPath, opts...)
}

// GetDCInfoByMAC retrieves dc-info for a specific client by MAC address.
func (s Service) GetDCInfoByMAC(
	ctx context.Context,
	clientMAC string,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessClientOperDcInfo, error) {
	normalizedMAC, err := service.RequireMACAddress(clientMAC)
	if err != nil {
		return nil, err
	}

	endpoint := s.Client().RESTCONFBuilder().BuildQueryURL(routes.ClientDcInfoPath, normalizedMAC)
	return core.Get[CiscoIOSXEWirelessClientOperDcInfo](ctx, s.Client(), endpoint, opts...)
}

// ListDot11Info returns 802.11 operational data for clients.
func (s Service) ListDot11Info(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessClientOperDot11OperData, error) {
	return core.Get[CiscoIOSXEWirelessClientOperDot11OperData](
		ctx,
		s.Client(),
		routes.ClientDot11OperDataPath,
		opts...,
	)
}

// GetDot11InfoByMAC retrieves 802.11 operational data filtered by MAC address.
func (s Service) GetDot11InfoByMAC(
	ctx context.Context,
	mac string, opts ...core.GetOption,
) (*CiscoIOSXEWirelessClientOperDot11OperData, error) {
	normalizedMAC, err := service.RequireMACAddress(mac)
	if err != nil {
		return nil, err
	}

	url := s.Client().RESTCONFBuilder().BuildQueryURL(routes.ClientDot11OperDataPath, normalizedMAC)
	return core.Get[CiscoIOSXEWirelessClientOperDot11OperData](ctx, s.Client(), url, opts...)
}

// ListMMIFClientHistory retrieves mobility manager interface client history.
func (s Service) ListMMIFClientHistory(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessClientOperMmIfClientHistory, error) {
	return core.Get[CiscoIOSXEWirelessClientOperMmIfClientHistory](
		ctx,
		s.Client(),
		routes.ClientMmIfClientHistoryPath,
		opts...,
	)
}

// GetMMIFClientHistoryByMAC retrieves mm-if-client-history for a specific client by MAC address.
func (s Service) GetMMIFClientHistoryByMAC(
	ctx context.Context,
	clientMAC string, opts ...core.GetOption,
) (*CiscoIOSXEWirelessClientOperMmIfClientHistory, error) {
	normalizedMAC, err := service.RequireMACAddress(clientMAC)
	if err != nil {
		return nil, err
	}

	endpoint := s.Client().RESTCONFBuilder().BuildQueryURL(routes.ClientMmIfClientHistoryPath, normalizedMAC)
	return core.Get[CiscoIOSXEWirelessClientOperMmIfClientHistory](ctx, s.Client(), endpoint, opts...)
}

// ListMMIFClientStats retrieves mobility manager interface client statistics.
func (s Service) ListMMIFClientStats(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessClientOperMmIfClientStats, error) {
	return core.Get[CiscoIOSXEWirelessClientOperMmIfClientStats](
		ctx,
		s.Client(),
		routes.ClientMmIfClientStatsPath,
		opts...,
	)
}

// GetMMIFClientStatsByMAC retrieves mm-if-client-stats for a specific client by MAC address.
func (s Service) GetMMIFClientStatsByMAC(
	ctx context.Context,
	clientMAC string, opts ...core.GetOption,
) (*CiscoIOSXEWirelessClientOperMmIfClientStats, error) {
	normalizedMAC, err := service.RequireMACAddress(clientMAC)
	if err != nil {
		return nil, err
	}

	endpoint := s.Client().RESTCONFBuilder().BuildQueryURL(routes.ClientMmIfClientStatsPath, normalizedMAC)
	return core.Get[CiscoIOSXEWirelessClientOperMmIfClientStats](ctx, s.Client(), endpoint, opts...)
}

// ListMobilityInfo retrieves mobility operational data for clients.
func (s Service) ListMobilityInfo(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessClientOperMobilityOperData, error) {
	return core.Get[CiscoIOSXEWirelessClientOperMobilityOperData](
		ctx,
		s.Client(),
		routes.ClientMobilityOperDataPath,
		opts...,
	)
}

// GetMobilityInfoByMAC retrieves mobility operational data filtered by MAC address.
func (s Service) GetMobilityInfoByMAC(
	ctx context.Context,
	clientMAC string, opts ...core.GetOption,
) (*CiscoIOSXEWirelessClientOperMobilityOperData, error) {
	normalizedMAC, err := service.RequireMACAddress(clientMAC)
	if err != nil {
		return nil, err
	}

	endpoint := s.Client().RESTCONFBuilder().BuildQueryURL(routes.ClientMobilityOperDataPath, normalizedMAC)
	return core.Get[CiscoIOSXEWirelessClientOperMobilityOperData](ctx, s.Client(), endpoint, opts...)
}

// ListPolicyInfo retrieves client policy data.
func (s Service) ListPolicyInfo(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessClientOperPolicyData, error) {
	return core.Get[CiscoIOSXEWirelessClientOperPolicyData](ctx, s.Client(), routes.ClientPolicyDataPath, opts...)
}

// GetPolicyInfoByMAC retrieves policy-data for a specific client by MAC address.
func (s Service) GetPolicyInfoByMAC(
	ctx context.Context,
	clientMAC string, opts ...core.GetOption,
) (*CiscoIOSXEWirelessClientOperPolicyData, error) {
	normalizedMAC, err := service.RequireMACAddress(clientMAC)
	if err != nil {
		return nil, err
	}

	endpoint := s.Client().RESTCONFBuilder().BuildQueryURL(routes.ClientPolicyDataPath, normalizedMAC)
	return core.Get[CiscoIOSXEWirelessClientOperPolicyData](ctx, s.Client(), endpoint, opts...)
}

// ListSISFDB retrieves SISF database MAC information.
func (s Service) ListSISFDB(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessClientOperSisfDBMac, error) {
	return core.Get[CiscoIOSXEWirelessClientOperSisfDBMac](ctx, s.Client(), routes.ClientSisfDBMacPath, opts...)
}

// GetSISFDBByMAC retrieves sisf-db-mac for a specific client by MAC address.
func (s Service) GetSISFDBByMAC(
	ctx context.Context,
	clientMAC string,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessClientOperSisfDBMac, error) {
	normalizedMAC, err := service.RequireMACAddress(clientMAC)
	if err != nil {
		return nil, err
	}

	endpoint := s.Client().RESTCONFBuilder().BuildQueryURL(routes.ClientSisfDBMacPath, normalizedMAC)
	return core.Get[CiscoIOSXEWirelessClientOperSisfDBMac](ctx, s.Client(), endpoint, opts...)
}

// ListTrafficStats retrieves client traffic statistics.
func (s Service) ListTrafficStats(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessClientOperTrafficStatsData, error) {
	return core.Get[CiscoIOSXEWirelessClientOperTrafficStatsData](
		ctx,
		s.Client(),
		routes.ClientTrafficStatsPath,
		opts...,
	)
}

// GetTrafficStatsByMAC retrieves traffic-stats for a specific client by MAC address.
func (s Service) GetTrafficStatsByMAC(
	ctx context.Context,
	clientMAC string, opts ...core.GetOption,
) (*CiscoIOSXEWirelessClientOperTrafficStatsData, error) {
	normalizedMAC, err := service.RequireMACAddress(clientMAC)
	if err != nil {
		return nil, err
	}

	endpoint := s.Client().RESTCONFBuilder().BuildQueryURL(routes.ClientTrafficStatsPath, normalizedMAC)
	return core.Get[CiscoIOSXEWirelessClientOperTrafficStatsData](ctx, s.Client(), endpoint, opts...)
}

// DeauthenticateByMAC drops the session of the client with this address.
//
// The By suffix names the arm of the RPC's mandatory choice this fills. The controller answers 204
// whether or not the address matched a client, so a caller that needs to know reads the client
// first; measured on 17.15.6, an address matching nothing and one matching a live client are
// indistinguishable. The client reassociates on its own, so this forces a reassociation rather
// than removing anything, and it leaves the PMK cache intact.
//
// The operation is declared from module revision 2024-03-01, which 17.15 is the first release to
// serve. An earlier release answers 400 with "invalid path" rather than 404.
func (s Service) DeauthenticateByMAC(ctx context.Context, clientMAC string) error {
	normalizedMAC, err := validation.NormalizeMACAddress(clientMAC)
	if err != nil {
		return fmt.Errorf(ErrInvalidClientMACFormat, clientMAC)
	}

	return s.deauthenticate(ctx, ClientDeauthRPCInput{MACAddr: normalizedMAC})
}

// DeauthenticateByIP drops the session of the client holding this address.
//
// The address is resolved within zone 0, which is the only zone the payload names. A client holds
// one IPv4 binding and may hold several IPv6 bindings, so an address names one client while a
// client does not name one address.
func (s Service) DeauthenticateByIP(ctx context.Context, clientIP string) error {
	if !validation.IsNonEmptyString(clientIP) {
		return errors.New(ErrEmptyClientIPAddr)
	}

	return s.deauthenticate(ctx, ClientDeauthRPCInput{IPAddr: clientIP})
}

// DeauthenticateByUsername drops the session of the client authenticated under this username.
//
// A username is not measured to select one client: the lab estate carries one session per
// username, so nothing establishes what the controller does when a username holds several.
func (s Service) DeauthenticateByUsername(ctx context.Context, username string) error {
	if !validation.IsNonEmptyString(username) {
		return errors.New(ErrEmptyClientUsername)
	}

	return s.deauthenticate(ctx, ClientDeauthRPCInput{Username: username})
}

// deauthenticate posts one arm of the deauthentication RPC's choice.
func (s Service) deauthenticate(ctx context.Context, input ClientDeauthRPCInput) error {
	return core.PostRPCVoid(ctx, s.Client(), routes.ClientDeauthRPC, ClientDeauthRPCPayload{Input: input})
}
