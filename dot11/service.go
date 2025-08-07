package dot11

import (
	"context"
	"net/http"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
	"github.com/umatare5/cisco-ios-xe-wireless-go/wnc"
)

const (
	// Dot11CfgBasePath defines the base path for 802.11 configuration endpoints
	Dot11CfgBasePath = "Cisco-IOS-XE-wireless-dot11-cfg:dot11-cfg-data"
	// Dot11CfgEndpoint defines the endpoint for 802.11 configuration data
	Dot11CfgEndpoint = Dot11CfgBasePath
)

// Service provides 802.11 operations.
type Service struct {
	c *wnc.Client
}

// NewService creates a new service instance.
func NewService(client *wnc.Client) Service {
	return Service{c: client}
}

// Cfg returns 802.11 configuration data.
func (s Service) Cfg(ctx context.Context) (*model.Dot11CfgResponse, error) {
	var result model.Dot11CfgResponse
	return &result, s.c.Do(ctx, http.MethodGet, Dot11CfgEndpoint, &result)
}
