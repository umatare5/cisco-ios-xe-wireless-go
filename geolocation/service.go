package geolocation

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/constants"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
)

const (
	// GeolocationOperBasePath defines the base path for geolocation operational endpoints
	GeolocationOperBasePath = constants.YANGModelPrefix + "geolocation-oper:geolocation-oper-data"
	// GeolocationOperEndpoint defines the endpoint for geolocation operational data
	GeolocationOperEndpoint = GeolocationOperBasePath
	// GeolocationOperApGeoLocStatsEndpoint defines the endpoint for AP geolocation statistics
	GeolocationOperApGeoLocStatsEndpoint = GeolocationOperBasePath + "/ap-geo-loc-stats"
)

// Service provides Geolocation operations.
type Service struct {
	c *core.Client
}

// NewService creates a new service instance.
func NewService(c *core.Client) Service {
	return Service{c: c}
}

// GetOper returns geolocation operational data.
// This endpoint provides geographic positioning and location mapping information.
func (s Service) GetOper(ctx context.Context) (*model.GeolocationOperResponse, error) {
	return core.Get[model.GeolocationOperResponse](ctx, s.c, GeolocationOperEndpoint)
}

// GetApGeoLocStats returns AP geolocation statistics.
func (s Service) GetApGeoLocStats(ctx context.Context) (*model.GeolocationOperApGeoLocStatsResponse, error) {
	return core.Get[model.GeolocationOperApGeoLocStatsResponse](ctx, s.c, GeolocationOperApGeoLocStatsEndpoint)
}
