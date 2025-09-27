package geolocation

import (
	"context"
	"strings"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/validation"
)

// Service provides geolocation tracking operations for Cisco IOS-XE Wireless LAN Controller.
type Service struct {
	service.BaseService
}

// NewService creates a new Geolocation service instance with the provided client.
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}

// GetOperational retrieves geolocation operational data from the controller.
func (s Service) GetOperational(ctx context.Context) (*CiscoIOSXEWirelessGeolocationOper, error) {
	return core.Get[CiscoIOSXEWirelessGeolocationOper](ctx, s.Client(), routes.GeolocationOperPath)
}

// ListAPGeolocationStats retrieves AP geolocation statistics.
func (s Service) ListAPGeolocationStats(ctx context.Context) (*CiscoIOSXEWirelessGeolocationOperApGeoLocStats, error) {
	return core.Get[CiscoIOSXEWirelessGeolocationOperApGeoLocStats](
		ctx,
		s.Client(),
		routes.GeolocationApGeoLocStatsPath,
	)
}

// ListAPGeolocationData retrieves AP geolocation data using CiscoIOSXEWirelessGeolocationOperApGeoLocData wrapper.
func (s Service) ListAPGeolocationData(ctx context.Context) (*CiscoIOSXEWirelessGeolocationOperApGeoLocData, error) {
	return core.Get[CiscoIOSXEWirelessGeolocationOperApGeoLocData](ctx, s.Client(), routes.GeolocationApGeoLocDataPath)
}

// GetAPGeolocationDataByMAC retrieves AP geolocation data for a specific AP by MAC address.
func (s Service) GetAPGeolocationDataByMAC(ctx context.Context, apMAC string) (*ApGeoLocData, error) {
	if apMAC == "" || strings.TrimSpace(apMAC) == "" {
		return nil, core.ErrResourceNotFound
	}

	// Validate and normalize MAC address
	if err := validation.ValidateMACAddress(apMAC); err != nil {
		return nil, err
	}

	normalizedMAC, err := validation.NormalizeMACAddress(apMAC)
	if err != nil {
		return nil, err
	}

	url := s.Client().RESTCONFBuilder().BuildQueryURL(routes.GeolocationApGeoLocDataPath, normalizedMAC)
	return core.Get[ApGeoLocData](ctx, s.Client(), url)
}
