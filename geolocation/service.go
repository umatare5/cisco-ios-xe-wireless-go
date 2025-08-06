package geolocation

import (
	"context"
	"net/http"

	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
	"github.com/umatare5/cisco-ios-xe-wireless-go/wnc"
)

// Service provides Geolocation domain operations for wireless controller geolocation data.
type Service struct {
	c *wnc.Client
}

// NewService creates a new Geolocation service instance.
func NewService(c *wnc.Client) Service {
	return Service{c: c}
}

// Operational Data Methods

// Oper returns geolocation operational data.
func (s Service) Oper(ctx context.Context) (*model.GeolocationOperResponse, error) {
	var out model.GeolocationOperResponse
	return &out, s.c.Do(ctx, http.MethodGet,
		"Cisco-IOS-XE-wireless-geolocation-oper:geolocation-oper-data", &out)
}

// ApGeoLocStats returns AP geolocation statistics.
func (s Service) ApGeoLocStats(ctx context.Context) (*model.GeolocationOperApGeoLocStatsResponse, error) {
	var out model.GeolocationOperApGeoLocStatsResponse
	return &out, s.c.Do(ctx, http.MethodGet,
		"Cisco-IOS-XE-wireless-geolocation-oper:geolocation-oper-data/ap-geo-loc-stats", &out)
}
