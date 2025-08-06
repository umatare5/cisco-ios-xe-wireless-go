package geolocation

import (
	"context"
	"net/http"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
	"github.com/umatare5/cisco-ios-xe-wireless-go/wnc"
)

// Service provides geolocation domain operations.
type Service struct {
	c *wnc.Client
}

// NewService creates a new geolocation service instance.
func NewService(c *wnc.Client) *Service {
	return &Service{c: c}
}

// Oper returns geolocation operational data.
// This endpoint provides geographic positioning and location mapping information.
func (s *Service) Oper(ctx context.Context) (*model.GeolocationOperResponse, error) {
	const endpoint = "Cisco-IOS-XE-wireless-geolocation-oper:geolocation-oper-data"

	var result model.GeolocationOperResponse
	err := s.c.Do(ctx, http.MethodGet, endpoint, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// ApGeoLocStats returns AP geolocation statistics.
func (s *Service) ApGeoLocStats(ctx context.Context) (*model.GeolocationOperApGeoLocStatsResponse, error) {
	const endpoint = "Cisco-IOS-XE-wireless-geolocation-oper:geolocation-oper-data/ap-geo-loc-stats"

	var result model.GeolocationOperApGeoLocStatsResponse
	err := s.c.Do(ctx, http.MethodGet, endpoint, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
