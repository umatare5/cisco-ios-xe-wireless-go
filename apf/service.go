package apf

import (
	"context"
	"net/http"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
	"github.com/umatare5/cisco-ios-xe-wireless-go/wnc"
)

const (
	// ApfCfgBasePath defines the base path for APF configuration endpoints
	ApfCfgBasePath = "Cisco-IOS-XE-wireless-apf-cfg:apf-cfg-data"
	// ApfCfgEndpoint defines the endpoint for APF configuration data
	ApfCfgEndpoint = ApfCfgBasePath
)

// Service provides APF operations.
type Service struct {
	c *wnc.Client
}

// NewService creates a new service instance.
func NewService(client *wnc.Client) Service {
	return Service{c: client}
}

// Cfg returns APF configuration data.
func (s Service) Cfg(ctx context.Context) (*model.ApfCfgResponse, error) {
	var result model.ApfCfgResponse
	return &result, s.c.Do(ctx, http.MethodGet, ApfCfgEndpoint, &result)
}
