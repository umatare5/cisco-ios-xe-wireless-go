package core

import (
	"context"
	"net/http"
)

// Get is a generic helper reducing boilerplate in service GET methods.
func Get[T any](ctx context.Context, c *Client, endpoint string) (*T, error) { //nolint:revive // intentional generic helper
	var out T
	return &out, c.Do(ctx, http.MethodGet, endpoint, &out)
}
