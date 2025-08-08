package site

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/constants"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
)

const (
	// SiteOperBasePath defines the base path for site operational data endpoints
	SiteOperBasePath = constants.YANGModelPrefix + "site-oper:site-oper-data"
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

// GetOper returns Site operational data.
func (s Service) GetOper(ctx context.Context) (*model.SiteOperResponse, error) {
	return core.Get[model.SiteOperResponse](ctx, s.c, SiteOperEndpoint)
}
