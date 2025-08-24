package core

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	internalerrors "github.com/umatare5/cisco-ios-xe-wireless-go/internal/errors"
)

// Generic HTTP Operation Functions for Service Layer
// These functions provide a consistent interface for HTTP operations across all services.

// Get is a generic helper reducing boilerplate in service GET methods.
func Get[T any](ctx context.Context, c *Client, endpoint string) (*T, error) {
	if c == nil {
		return nil, errors.New(internalerrors.ErrClientNil)
	}

	body, err := c.Do(ctx, http.MethodGet, endpoint)
	if err != nil {
		return nil, err
	}

	var out T
	if len(body) > 0 {
		if err := json.Unmarshal(body, &out); err != nil {
			return nil, fmt.Errorf("failed to unmarshal response: %w", err)
		}
	}

	return &out, nil
}

// Post is a generic helper for sending POST requests with payload.
func Post[T any](ctx context.Context, c *Client, endpoint string, payload any) (*T, error) {
	if c == nil {
		return nil, errors.New(internalerrors.ErrClientNil)
	}

	body, err := c.DoWithPayload(ctx, http.MethodPost, endpoint, payload)
	if err != nil {
		return nil, err
	}

	var out T
	if len(body) > 0 {
		if err := json.Unmarshal(body, &out); err != nil {
			return nil, fmt.Errorf("failed to unmarshal response: %w", err)
		}
	}

	return &out, nil
}

// PostVoid is a generic helper for POST operations without expecting a response body.
func PostVoid(ctx context.Context, c *Client, endpoint string, payload any) error {
	if c == nil {
		return errors.New(internalerrors.ErrClientNil)
	}
	_, err := c.DoWithPayload(ctx, http.MethodPost, endpoint, payload)
	return err
}

// PostRPCVoid is a generic helper for RPC POST operations without expecting a response body.
func PostRPCVoid(ctx context.Context, c *Client, rpcEndpoint string, payload any) error {
	if c == nil {
		return errors.New(internalerrors.ErrClientNil)
	}
	_, err := c.DoRPCWithPayload(ctx, http.MethodPost, rpcEndpoint, payload)
	return err
}

// Put is a generic helper for PUT operations that expect a response body.
func Put[T any](ctx context.Context, c *Client, endpoint string, payload any) (*T, error) {
	if c == nil {
		return nil, errors.New(internalerrors.ErrClientNil)
	}

	body, err := c.DoWithPayload(ctx, http.MethodPut, endpoint, payload)
	if err != nil {
		return nil, err
	}

	var out T
	if len(body) > 0 {
		if err := json.Unmarshal(body, &out); err != nil {
			return nil, fmt.Errorf("failed to unmarshal response: %w", err)
		}
	}

	return &out, nil
}

// PutVoid is a generic helper for PUT operations without expecting a response body.
func PutVoid(ctx context.Context, c *Client, endpoint string, payload any) error {
	if c == nil {
		return errors.New(internalerrors.ErrClientNil)
	}
	_, err := c.DoWithPayload(ctx, http.MethodPut, endpoint, payload)
	return err
}

// Patch is a generic helper for PATCH operations that expect a response body.
func Patch[T any](ctx context.Context, c *Client, endpoint string, payload any) (*T, error) {
	if c == nil {
		return nil, errors.New(internalerrors.ErrClientNil)
	}

	body, err := c.DoWithPayload(ctx, http.MethodPatch, endpoint, payload)
	if err != nil {
		return nil, err
	}

	var out T
	if len(body) > 0 {
		if err := json.Unmarshal(body, &out); err != nil {
			return nil, fmt.Errorf("failed to unmarshal response: %w", err)
		}
	}

	return &out, nil
}

// PatchVoid is a generic helper for PATCH operations without expecting a response body.
func PatchVoid(ctx context.Context, c *Client, endpoint string, payload any) error {
	if c == nil {
		return errors.New(internalerrors.ErrClientNil)
	}
	_, err := c.DoWithPayload(ctx, http.MethodPatch, endpoint, payload)
	return err
}

// Delete is a generic helper for DELETE operations without expecting a response body.
func Delete(ctx context.Context, c *Client, endpoint string) error {
	if c == nil {
		return errors.New(internalerrors.ErrClientNil)
	}
	_, err := c.Do(ctx, http.MethodDelete, endpoint)
	return err
}

// Client HTTP Methods
// These methods provide direct access to HTTP operations on the client instance.
// They are primarily used for low-level operations and direct client access.

// Post sends a POST request to the specified endpoint with the given payload
func (c *Client) Post(ctx context.Context, endpoint string, payload any) error {
	_, err := c.DoWithPayload(ctx, http.MethodPost, endpoint, payload)
	return err
}

// Put sends a PUT request to the specified endpoint with the given payload
func (c *Client) Put(ctx context.Context, endpoint string, payload any) error {
	_, err := c.DoWithPayload(ctx, http.MethodPut, endpoint, payload)
	return err
}

// Patch sends a PATCH request to the specified endpoint with the given payload
func (c *Client) Patch(ctx context.Context, endpoint string, payload any) error {
	_, err := c.DoWithPayload(ctx, http.MethodPatch, endpoint, payload)
	return err
}

// Delete sends a DELETE request to the specified endpoint
func (c *Client) Delete(ctx context.Context, endpoint string) error {
	_, err := c.Do(ctx, http.MethodDelete, endpoint)
	return err
}
