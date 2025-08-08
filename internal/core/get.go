package core

import (
	"context"
	"net/http"
)

// Get is a generic helper reducing boilerplate in service GET methods.
func Get[T any](ctx context.Context, c *Client, endpoint string) (*T, error) { // generic helper for GET operations
	var out T
	return &out, c.Do(ctx, http.MethodGet, endpoint, &out)
}
