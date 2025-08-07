package geolocation

import (
	"context"
	"net/http"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/constants"
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

// Oper returns geolocation operational data.
// This endpoint provides geographic positioning and location mapping information.
func (s Service) Oper(ctx context.Context) (*model.GeolocationOperResponse, error) {
	var result model.GeolocationOperResponse
	return &result, s.c.Do(ctx, http.MethodGet, GeolocationOperEndpoint, &result)
}

// ApGeoLocStats returns AP geolocation statistics.
func (s Service) ApGeoLocStats(ctx context.Context) (*model.GeolocationOperApGeoLocStatsResponse, error) {
	var result model.GeolocationOperApGeoLocStatsResponse
	return &result, s.c.Do(ctx, http.MethodGet, GeolocationOperApGeoLocStatsEndpoint, &result)
}
