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
func (s Service) GetOperational(ctx context.Context) (*ClientOper, error) {
	result, err := core.Get[ClientOper](ctx, s.Client(), routes.ClientOperPath)
	if err == nil {
		return result, nil
	}
	if isKnownGetOperationalIssue(err) {
		// Return empty result for IOS-XE 17.18.1 compatibility
		// Known issue: main endpoint may fail intermittently on certain firmware versions
		return &ClientOper{}, nil
	}
	return nil, err
}

// ListCommonInfo retrieves common operational data for clients.
func (s Service) ListCommonInfo(ctx context.Context) (*ClientOperCommonOperData, error) {
	return core.Get[ClientOperCommonOperData](ctx, s.Client(), routes.ClientCommonOperDataPath)
}

// GetCommonInfoByMAC retrieves client operational data filtered by MAC address.
func (s Service) GetCommonInfoByMAC(
	ctx context.Context, mac string,
) (*ClientOperCommonOperData, error) {
	if mac == "" {
		return nil, core.ErrInvalidConfiguration
	}

	url := s.Client().RESTCONFBuilder().BuildQueryURL(routes.ClientCommonOperDataPath, mac)
	return core.Get[ClientOperCommonOperData](ctx, s.Client(), url)
}

// ListDCInfo retrieves discovery client information.
func (s Service) ListDCInfo(ctx context.Context) (*ClientOperDcInfo, error) {
	return core.Get[ClientOperDcInfo](ctx, s.Client(), routes.ClientDcInfoPath)
}

// GetDCInfoByMAC retrieves dc-info for a specific client by MAC address.
func (s Service) GetDCInfoByMAC(ctx context.Context, clientMAC string) (*ClientOperDcInfo, error) {
	if clientMAC == "" || strings.TrimSpace(clientMAC) == "" {
		return nil, core.ErrInvalidConfiguration
	}

	endpoint := s.Client().RESTCONFBuilder().BuildQueryURL(routes.ClientDcInfoPath, clientMAC)
	return core.Get[ClientOperDcInfo](ctx, s.Client(), endpoint)
}

// ListDot11Info returns 802.11 operational data for clients.
func (s Service) ListDot11Info(ctx context.Context) (*ClientOperDot11OperData, error) {
	result, err := core.Get[ClientOperDot11OperData](ctx, s.Client(), routes.ClientDot11OperDataPath)
	if err == nil {
		return result, nil
	}
	if isKnownDot11OperationalDataIssue(err) {
		// Return empty result for IOS-XE 17.18.1 compatibility
		// Known issue: endpoint may fail even after client connections and 90-second wait period
		return &ClientOperDot11OperData{Dot11OperData: []Dot11OperData{}}, nil
	}
	return nil, err
}

// GetDot11InfoByMAC retrieves 802.11 operational data filtered by MAC address.
func (s Service) GetDot11InfoByMAC(ctx context.Context, mac string) (*ClientOperDot11OperData, error) {
	if mac == "" || strings.TrimSpace(mac) == "" {
		return nil, core.ErrInvalidConfiguration
	}

	url := s.Client().RESTCONFBuilder().BuildQueryURL(routes.ClientDot11OperDataPath, mac)

	result, err := core.Get[ClientOperDot11OperData](ctx, s.Client(), url)
	if err == nil {
		return result, nil
	}
	if isKnownDot11OperationalDataIssue(err) {
		// Return empty result for IOS-XE 17.18.1 compatibility
		// Known issue: endpoint may fail even after client connections and 90-second wait period
		return &ClientOperDot11OperData{Dot11OperData: []Dot11OperData{}}, nil
	}
	return nil, err
}

// ListMMIFClientHistory retrieves mobility manager interface client history.
func (s Service) ListMMIFClientHistory(ctx context.Context) (*ClientOperMmIfClientHistory, error) {
	return core.Get[ClientOperMmIfClientHistory](ctx, s.Client(), routes.ClientMmIfClientHistoryPath)
}

// GetMMIFClientHistoryByMAC retrieves mm-if-client-history for a specific client by MAC address.
func (s Service) GetMMIFClientHistoryByMAC(
	ctx context.Context,
	clientMAC string,
) (*ClientOperMmIfClientHistory, error) {
	if clientMAC == "" || strings.TrimSpace(clientMAC) == "" {
		return nil, core.ErrInvalidConfiguration
	}

	endpoint := s.Client().RESTCONFBuilder().BuildQueryURL(routes.ClientMmIfClientHistoryPath, clientMAC)
	return core.Get[ClientOperMmIfClientHistory](ctx, s.Client(), endpoint)
}

// ListMMIFClientStats retrieves mobility manager interface client statistics.
func (s Service) ListMMIFClientStats(ctx context.Context) (*ClientOperMmIfClientStats, error) {
	return core.Get[ClientOperMmIfClientStats](ctx, s.Client(), routes.ClientMmIfClientStatsPath)
}

// GetMMIFClientStatsByMAC retrieves mm-if-client-stats for a specific client by MAC address.
func (s Service) GetMMIFClientStatsByMAC(
	ctx context.Context,
	clientMAC string,
) (*ClientOperMmIfClientStats, error) {
	if clientMAC == "" || strings.TrimSpace(clientMAC) == "" {
		return nil, core.ErrInvalidConfiguration
	}

	endpoint := s.Client().RESTCONFBuilder().BuildQueryURL(routes.ClientMmIfClientStatsPath, clientMAC)
	return core.Get[ClientOperMmIfClientStats](ctx, s.Client(), endpoint)
}

// ListMobilityInfo retrieves mobility operational data for clients.
func (s Service) ListMobilityInfo(ctx context.Context) (*ClientOperMobilityOperData, error) {
	return core.Get[ClientOperMobilityOperData](ctx, s.Client(), routes.ClientMobilityOperDataPath)
}

// GetMobilityInfoByMAC retrieves mobility-oper-data for a specific client by MAC address.
func (s Service) GetMobilityInfoByMAC(
	ctx context.Context,
	clientMAC string,
) (*ClientOperMobilityOperData, error) {
	if clientMAC == "" || strings.TrimSpace(clientMAC) == "" {
		return nil, core.ErrInvalidConfiguration
	}

	endpoint := s.Client().RESTCONFBuilder().BuildQueryURL(routes.ClientMobilityOperDataPath, clientMAC)
	return core.Get[ClientOperMobilityOperData](ctx, s.Client(), endpoint)
}

// ListPolicyInfo retrieves client policy data.
func (s Service) ListPolicyInfo(ctx context.Context) (*ClientOperPolicyData, error) {
	return core.Get[ClientOperPolicyData](ctx, s.Client(), routes.ClientPolicyDataPath)
}

// GetPolicyInfoByMAC retrieves policy-data for a specific client by MAC address.
func (s Service) GetPolicyInfoByMAC(ctx context.Context, clientMAC string) (*ClientOperPolicyData, error) {
	if clientMAC == "" || strings.TrimSpace(clientMAC) == "" {
		return nil, core.ErrInvalidConfiguration
	}

	endpoint := s.Client().RESTCONFBuilder().BuildQueryURL(routes.ClientPolicyDataPath, clientMAC)
	return core.Get[ClientOperPolicyData](ctx, s.Client(), endpoint)
}

// ListSISFDB retrieves SISF database MAC information.
func (s Service) ListSISFDB(ctx context.Context) (*ClientOperSisfDBMac, error) {
	return core.Get[ClientOperSisfDBMac](ctx, s.Client(), routes.ClientSisfDBMacPath)
}

// GetSISFDBByMAC retrieves sisf-db-mac for a specific client by MAC address.
func (s Service) GetSISFDBByMAC(ctx context.Context, clientMAC string) (*ClientOperSisfDBMac, error) {
	if clientMAC == "" || strings.TrimSpace(clientMAC) == "" {
		return nil, core.ErrInvalidConfiguration
	}

	endpoint := s.Client().RESTCONFBuilder().BuildQueryURL(routes.ClientSisfDBMacPath, clientMAC)
	return core.Get[ClientOperSisfDBMac](ctx, s.Client(), endpoint)
}

// ListTrafficStats retrieves client traffic statistics.
func (s Service) ListTrafficStats(ctx context.Context) (*ClientOperTrafficStats, error) {
	return core.Get[ClientOperTrafficStats](ctx, s.Client(), routes.ClientTrafficStatsPath)
}

// GetTrafficStatsByMAC retrieves traffic-stats for a specific client by MAC address.
func (s Service) GetTrafficStatsByMAC(
	ctx context.Context,
	clientMAC string,
) (*ClientOperTrafficStats, error) {
	if clientMAC == "" || strings.TrimSpace(clientMAC) == "" {
		return nil, core.ErrInvalidConfiguration
	}

	endpoint := s.Client().RESTCONFBuilder().BuildQueryURL(routes.ClientTrafficStatsPath, clientMAC)
	return core.Get[ClientOperTrafficStats](ctx, s.Client(), endpoint)
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
