package lisp

import (
	"context"
	"net/http"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
)

const (
	// LispOperBasePath defines the base path for LISP operational data endpoints
	LispOperBasePath = "Cisco-IOS-XE-wireless-lisp-oper:lisp-oper-data"
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

// Oper returns LISP operational data.
func (s Service) Oper(ctx context.Context) (*model.LispOperResponse, error) {
	var result model.LispOperResponse
	return &result, s.c.Do(ctx, http.MethodGet, LispOperEndpoint, &result)
}
