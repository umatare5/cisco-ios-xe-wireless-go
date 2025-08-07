package site

import (
	"context"
	"net/http"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
)

const (
	// SiteOperBasePath defines the base path for site operational endpoints
	SiteOperBasePath = "Cisco-IOS-XE-wireless-site-oper:site-oper-data"
	// SiteOperEndpoint defines the endpoint for site operational data
	SiteOperEndpoint = SiteOperBasePath
)

// Service provides Site operations.
type Service struct {
	c *core.Client
}

// NewService creates a new service instance.
func NewService(client *core.Client) Service {
	return Service{c: client}
}

// Oper returns Site operational data.
func (s Service) Oper(ctx context.Context) (*model.SiteOperResponse, error) {
	var result model.SiteOperResponse
	return &result, s.c.Do(ctx, http.MethodGet, SiteOperEndpoint, &result)
}
