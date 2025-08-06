package location

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
	"github.com/umatare5/cisco-ios-xe-wireless-go/wnc"
)

// Service provides access to Location services operations.
type Service struct {
	c *wnc.Client
}

// NewService creates a new Location service instance.
func NewService(client *wnc.Client) *Service {
	return &Service{c: client}
}

// Cfg retrieves Location configuration data.
func (s *Service) Cfg(ctx context.Context) (*model.LocationCfgResponse, error) {
	const endpoint = "Cisco-IOS-XE-wireless-location-cfg:location-cfg-data"

	var result model.LocationCfgResponse
	err := s.c.Do(ctx, "GET", endpoint, &result)
	return &result, err
}
