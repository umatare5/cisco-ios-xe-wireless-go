package location

import (
	"context"
	"net/http"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
	"github.com/umatare5/cisco-ios-xe-wireless-go/wnc"
)

const (
	// LocationCfgBasePath defines the base path for location configuration endpoints
	LocationCfgBasePath = "Cisco-IOS-XE-wireless-location-cfg:location-cfg-data"
	// LocationCfgEndpoint defines the endpoint for location configuration data
	LocationCfgEndpoint = LocationCfgBasePath
)

// Service provides Location services operations.
type Service struct {
	c *wnc.Client
}

// NewService creates a new service instance.
func NewService(client *wnc.Client) Service {
	return Service{c: client}
}

// Cfg returns Location configuration data.
func (s Service) Cfg(ctx context.Context) (*model.LocationCfgResponse, error) {
	var result model.LocationCfgResponse
	return &result, s.c.Do(ctx, http.MethodGet, LocationCfgEndpoint, &result)
}
