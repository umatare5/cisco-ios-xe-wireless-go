package geolocation

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
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
func (s Service) GetOperational(ctx context.Context) (*GeolocationOper, error) {
	return core.Get[GeolocationOper](ctx, s.Client(), routes.GeolocationOperPath)
}

// ListAPGeolocationStats retrieves AP geolocation statistics.
func (s Service) ListAPGeolocationStats(ctx context.Context) (*GeolocationOperApGeoLocStats, error) {
	return core.Get[GeolocationOperApGeoLocStats](ctx, s.Client(), routes.GeolocationApGeoLocStatsPath)
}
