package rf

import (
	"context"
	"net/http"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
	"github.com/umatare5/cisco-ios-xe-wireless-go/wnc"
)

const (
	// RfCfgBasePath defines the base path for RF configuration endpoints
	RfCfgBasePath = "Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data"
	// RfCfgEndpoint defines the endpoint for RF configuration data
	RfCfgEndpoint = RfCfgBasePath
)

// Service provides RF operations.
type Service struct {
	c *wnc.Client
}

// NewService creates a new service instance.
func NewService(client *wnc.Client) Service {
	return Service{c: client}
}

// Cfg returns RF configuration data.
func (s Service) Cfg(ctx context.Context) (*model.RfCfgResponse, error) {
	var result model.RfCfgResponse
	return &result, s.c.Do(ctx, http.MethodGet, RfCfgEndpoint, &result)
}
