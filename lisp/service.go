package lisp

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/constants"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
)

const (
	// LispOperBasePath defines the base path for LISP operational data endpoints
	LispOperBasePath = constants.YANGModelPrefix + "lisp-oper:lisp-oper-data"
	// LispOperEndpoint defines the endpoint for LISP operational data
	LispOperEndpoint = LispOperBasePath
)

// Service provides LISP operations.
type Service struct {
	c *core.Client
}

// NewService creates a new service instance.
func NewService(client *core.Client) Service {
	return Service{c: client}
}

// GetOper returns LISP operational data.
func (s Service) GetOper(ctx context.Context) (*model.LispOperResponse, error) {
	return core.Get[model.LispOperResponse](ctx, s.c, LispOperEndpoint)
}
