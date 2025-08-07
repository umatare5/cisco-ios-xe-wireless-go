package cts

import (
	"context"
	"net/http"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
	"github.com/umatare5/cisco-ios-xe-wireless-go/wnc"
)

const (
	// CtsCfgBasePath defines the base path for CTS configuration endpoints
	CtsCfgBasePath = "Cisco-IOS-XE-wireless-cts-cfg:cts-cfg-data"
	// CtsCfgEndpoint defines the endpoint for CTS configuration data
	CtsCfgEndpoint = CtsCfgBasePath
)

// Service provides CTS operations.
type Service struct {
	c *wnc.Client
}

// NewService creates a new service instance.
func NewService(client *wnc.Client) Service {
	return Service{c: client}
}

// Cfg returns CTS configuration data.
func (s Service) Cfg(ctx context.Context) (*model.CtsCfgResponse, error) {
	var result model.CtsCfgResponse
	return &result, s.c.Do(ctx, http.MethodGet, CtsCfgEndpoint, &result)
}
