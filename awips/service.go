package awips

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/constants"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
)

const (
	// AwipsOperBasePath defines the base path for AWIPS operational data endpoints
	AwipsOperBasePath = constants.YANGModelPrefix + "awips-oper:awips-oper-data"
	// AwipsOperEndpoint defines the endpoint for AWIPS operational data
	AwipsOperEndpoint = AwipsOperBasePath
)

// Service provides AWIPS operations.
type Service struct {
	c *core.Client
}

// NewService creates a new service instance.
func NewService(client *core.Client) Service {
	return Service{c: client}
}

// GetOper returns AWIPS operational data.
func (s Service) GetOper(ctx context.Context) (*model.AwipsOperResponse, error) {
	return core.Get[model.AwipsOperResponse](ctx, s.c, AwipsOperEndpoint)
}
