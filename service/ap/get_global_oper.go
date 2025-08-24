// Package ap provides access point functionality for the Cisco IOS-XE Wireless Network Controller API.
package ap

import (
	"context"
	"strconv"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/ap"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
)

// GetGlobalOper retrieves the complete AP global operational data
func (s Service) GetGlobalOper(ctx context.Context) (*model.ApGlobalOper, error) {
	return core.Get[model.ApGlobalOper](ctx, s.Client(), routes.APGlobalOperEndpoint)
}

// GetGlobalOperEwlcApStats retrieves EWLC AP statistics
func (s Service) GetGlobalOperEwlcApStats(ctx context.Context) (*model.ApGlobalOperEwlcApStats, error) {
	return core.Get[model.ApGlobalOperEwlcApStats](ctx, s.Client(), routes.EwlcAPStatsEndpoint)
}

// GetGlobalOperApHistoryByEthernetMAC retrieves AP history data filtered by ethernet MAC address
func (s Service) GetGlobalOperApHistoryByEthernetMAC(
	ctx context.Context,
	ethernetMAC string,
) (*model.ApGlobalOperApHistory, error) {
	if err := s.ValidateClient(); err != nil {
		return nil, err
	}

	url := s.Client().RestconfBuilder().BuildPathQueryURL(routes.APGlobalOperEndpoint, "ap-history", ethernetMAC)
	return core.Get[model.ApGlobalOperApHistory](ctx, s.Client(), url)
}

// GetGlobalOperApJoinStats retrieves AP join statistics filtered by WTP MAC address
func (s Service) GetGlobalOperApJoinStats(
	ctx context.Context, wtpMAC string,
) (*model.ApGlobalOperApJoinStats, error) {
	if err := s.ValidateClient(); err != nil {
		return nil, err
	}

	url := s.Client().RestconfBuilder().BuildPathQueryURL(routes.APGlobalOperEndpoint, "ap-join-stats", wtpMAC)
	return core.Get[model.ApGlobalOperApJoinStats](ctx, s.Client(), url)
}

// GetGlobalOperWlanClientStatsByWlanID retrieves WLAN client statistics filtered by WLAN ID
func (s Service) GetGlobalOperWlanClientStatsByWlanID(
	ctx context.Context,
	wlanID int,
) (*model.ApGlobalOperWlanClientStats, error) {
	if err := s.ValidateClient(); err != nil {
		return nil, err
	}

	url := s.Client().RestconfBuilder().BuildPathQueryURL(
		routes.APGlobalOperEndpoint, "wlan-client-stats", strconv.Itoa(wlanID))
	return core.Get[model.ApGlobalOperWlanClientStats](ctx, s.Client(), url)
}

// GetGlobalOperApLocationStatsByLocation retrieves AP location statistics filtered by location
func (s Service) GetGlobalOperApLocationStatsByLocation(
	ctx context.Context,
	location string,
) (*model.ApGlobalOperApLocationStats, error) {
	if err := s.ValidateClient(); err != nil {
		return nil, err
	}

	url := s.Client().RestconfBuilder().BuildPathQueryURL(routes.APGlobalOperEndpoint, "ap-location-stats", location)
	return core.Get[model.ApGlobalOperApLocationStats](ctx, s.Client(), url)
}

// GetGlobalOperApHistoryOnly retrieves only AP history data using fields parameter
func (s Service) GetGlobalOperApHistoryOnly(ctx context.Context) (*model.ApGlobalOperApHistory, error) {
	if err := s.ValidateClient(); err != nil {
		return nil, err
	}

	url := s.Client().RestconfBuilder().BuildFieldsURL(routes.APGlobalOperEndpoint, "ap-history")
	return core.Get[model.ApGlobalOperApHistory](ctx, s.Client(), url)
}

// GetGlobalOperApJoinStatsOnly retrieves only AP join statistics using fields parameter
func (s Service) GetGlobalOperApJoinStatsOnly(ctx context.Context) (*model.ApGlobalOperApJoinStats, error) {
	if err := s.ValidateClient(); err != nil {
		return nil, err
	}

	url := s.Client().RestconfBuilder().BuildFieldsURL(routes.APGlobalOperEndpoint, "ap-join-stats")
	return core.Get[model.ApGlobalOperApJoinStats](ctx, s.Client(), url)
}

// GetGlobalOperWlanClientStatsOnly retrieves only WLAN client statistics using fields parameter
func (s Service) GetGlobalOperWlanClientStatsOnly(
	ctx context.Context,
) (*model.ApGlobalOperWlanClientStats, error) {
	if err := s.ValidateClient(); err != nil {
		return nil, err
	}

	url := s.Client().RestconfBuilder().BuildFieldsURL(routes.APGlobalOperEndpoint, "wlan-client-stats")
	return core.Get[model.ApGlobalOperWlanClientStats](ctx, s.Client(), url)
}
