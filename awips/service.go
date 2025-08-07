package awips

import (
	"context"
	"net/http"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
)

const (
	// AwipsOperBasePath defines the base path for AWIPS operational data endpoints
	AwipsOperBasePath = "Cisco-IOS-XE-wireless-awips-oper:awips-oper-data"
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

// Oper returns AWIPS operational data.
func (s Service) Oper(ctx context.Context) (*model.AwipsOperResponse, error) {
	var result model.AwipsOperResponse
	return &result, s.c.Do(ctx, http.MethodGet, AwipsOperEndpoint, &result)
}
