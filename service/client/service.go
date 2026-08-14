package client

import (
	"context"
	"strings"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
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
	result, err := core.Get[CiscoIOSXEWirelessClientOper](ctx, s.Client(), routes.ClientOperPath, opts...)
	if err == nil {
		return result, nil
	}
	if isKnownGetOperationalIssue(err) {
		// Return empty result for IOS-XE 17.18.1 compatibility
		// Known issue: main endpoint may fail intermittently on certain firmware versions
		return &CiscoIOSXEWirelessClientOper{}, nil
	}
	return nil, err
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
		opts...)
}

// GetCommonInfoByMAC retrieves client operational data filtered by MAC address.
func (s Service) GetCommonInfoByMAC(
	ctx context.Context, mac string, opts ...core.GetOption,
) (*CiscoIOSXEWirelessClientOperCommonOperData, error) {
	if mac == "" {
		return nil, core.ErrInvalidConfiguration
	}

	url := s.Client().RESTCONFBuilder().BuildQueryURL(routes.ClientCommonOperDataPath, mac)
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
	if clientMAC == "" || strings.TrimSpace(clientMAC) == "" {
		return nil, core.ErrInvalidConfiguration
	}

	endpoint := s.Client().RESTCONFBuilder().BuildQueryURL(routes.ClientDcInfoPath, clientMAC)
	return core.Get[CiscoIOSXEWirelessClientOperDcInfo](ctx, s.Client(), endpoint, opts...)
}

// ListDot11Info returns 802.11 operational data for clients.
func (s Service) ListDot11Info(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessClientOperDot11OperData, error) {
	result, err := core.Get[CiscoIOSXEWirelessClientOperDot11OperData](
		ctx,
		s.Client(),
		routes.ClientDot11OperDataPath,
		opts...)
	if err == nil {
		return result, nil
	}
	if isKnownDot11OperationalDataIssue(err) {
		// Return empty result for IOS-XE 17.18.1 compatibility
		// Known issue: endpoint may fail even after client connections and 90-second wait period
		return &CiscoIOSXEWirelessClientOperDot11OperData{Dot11OperData: []Dot11OperData{}}, nil
	}
	return nil, err
}

// GetDot11InfoByMAC retrieves 802.11 operational data filtered by MAC address.
func (s Service) GetDot11InfoByMAC(
	ctx context.Context,
	mac string, opts ...core.GetOption,
) (*CiscoIOSXEWirelessClientOperDot11OperData, error) {
	if mac == "" || strings.TrimSpace(mac) == "" {
		return nil, core.ErrInvalidConfiguration
	}

	url := s.Client().RESTCONFBuilder().BuildQueryURL(routes.ClientDot11OperDataPath, mac)

	result, err := core.Get[CiscoIOSXEWirelessClientOperDot11OperData](ctx, s.Client(), url, opts...)
	if err == nil {
		return result, nil
	}
	if isKnownDot11OperationalDataIssue(err) {
		// Return empty result for IOS-XE 17.18.1 compatibility
		// Known issue: endpoint may fail even after client connections and 90-second wait period
		return &CiscoIOSXEWirelessClientOperDot11OperData{Dot11OperData: []Dot11OperData{}}, nil
	}
	return nil, err
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
		opts...)
}

// GetMMIFClientHistoryByMAC retrieves mm-if-client-history for a specific client by MAC address.
func (s Service) GetMMIFClientHistoryByMAC(
	ctx context.Context,
	clientMAC string, opts ...core.GetOption,
) (*CiscoIOSXEWirelessClientOperMmIfClientHistory, error) {
	if clientMAC == "" || strings.TrimSpace(clientMAC) == "" {
		return nil, core.ErrInvalidConfiguration
	}

	endpoint := s.Client().RESTCONFBuilder().BuildQueryURL(routes.ClientMmIfClientHistoryPath, clientMAC)
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
		opts...)
}

// GetMMIFClientStatsByMAC retrieves mm-if-client-stats for a specific client by MAC address.
func (s Service) GetMMIFClientStatsByMAC(
	ctx context.Context,
	clientMAC string, opts ...core.GetOption,
) (*CiscoIOSXEWirelessClientOperMmIfClientStats, error) {
	if clientMAC == "" || strings.TrimSpace(clientMAC) == "" {
		return nil, core.ErrInvalidConfiguration
	}

	endpoint := s.Client().RESTCONFBuilder().BuildQueryURL(routes.ClientMmIfClientStatsPath, clientMAC)
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
		opts...)
}

// GetMobilityInfoByMAC retrieves mobility operational data filtered by MAC address.
func (s Service) GetMobilityInfoByMAC(
	ctx context.Context,
	clientMAC string, opts ...core.GetOption,
) (*CiscoIOSXEWirelessClientOperMobilityOperData, error) {
	if clientMAC == "" || strings.TrimSpace(clientMAC) == "" {
		return nil, core.ErrInvalidConfiguration
	}

	endpoint := s.Client().RESTCONFBuilder().BuildQueryURL(routes.ClientMobilityOperDataPath, clientMAC)
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
	if clientMAC == "" || strings.TrimSpace(clientMAC) == "" {
		return nil, core.ErrInvalidConfiguration
	}

	endpoint := s.Client().RESTCONFBuilder().BuildQueryURL(routes.ClientPolicyDataPath, clientMAC)
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
	if clientMAC == "" || strings.TrimSpace(clientMAC) == "" {
		return nil, core.ErrInvalidConfiguration
	}

	endpoint := s.Client().RESTCONFBuilder().BuildQueryURL(routes.ClientSisfDBMacPath, clientMAC)
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
		opts...)
}

// GetTrafficStatsByMAC retrieves traffic-stats for a specific client by MAC address.
func (s Service) GetTrafficStatsByMAC(
	ctx context.Context,
	clientMAC string, opts ...core.GetOption,
) (*CiscoIOSXEWirelessClientOperTrafficStatsData, error) {
	if clientMAC == "" || strings.TrimSpace(clientMAC) == "" {
		return nil, core.ErrInvalidConfiguration
	}

	endpoint := s.Client().RESTCONFBuilder().BuildQueryURL(routes.ClientTrafficStatsPath, clientMAC)
	return core.Get[CiscoIOSXEWirelessClientOperTrafficStatsData](ctx, s.Client(), endpoint, opts...)
}

// isKnownGetOperationalIssue checks if the error is a known IOS-XE 17.18.1 compatibility issue
// specific to the GetOperational method (main client operational data endpoint).
func isKnownGetOperationalIssue(err error) bool {
	errorMsg := err.Error()
	// Known IOS-XE 17.18.1 GetOperational issues that require empty response fallback
	return strings.Contains(errorMsg, "unexpected EOF")
}

// isKnownDot11OperationalDataIssue checks if the error is a known IOS-XE 17.18.1 compatibility issue
// specific to the Dot11OperationalData methods (802.11 operational data endpoint).
func isKnownDot11OperationalDataIssue(err error) bool {
	errorMsg := err.Error()
	// Known IOS-XE 17.18.1 Dot11OperationalData issues that require empty response fallback
	return strings.Contains(errorMsg, "failed to retrieve table cursor") ||
		strings.Contains(errorMsg, "Process DBAL response failed")
}
