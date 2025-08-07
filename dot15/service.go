package dot15

import (
	"context"
	"net/http"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
	"github.com/umatare5/cisco-ios-xe-wireless-go/wnc"
)

const (
	// Dot15CfgBasePath defines the base path for 802.15 configuration endpoints
	Dot15CfgBasePath = "Cisco-IOS-XE-wireless-dot15-cfg:dot15-cfg-data"
	// Dot15CfgEndpoint defines the endpoint for 802.15 configuration data
	Dot15CfgEndpoint = Dot15CfgBasePath
)

// Service provides 802.15 operations.
type Service struct {
	c *wnc.Client
}

// NewService creates a new service instance.
func NewService(client *wnc.Client) Service {
	return Service{c: client}
}

// Cfg returns 802.15 configuration data.
func (s Service) Cfg(ctx context.Context) (*model.Dot15CfgResponse, error) {
	var result model.Dot15CfgResponse
	return &result, s.c.Do(ctx, http.MethodGet, Dot15CfgEndpoint, &result)
}
