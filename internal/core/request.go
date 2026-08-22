package core

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	ierrors "github.com/umatare5/cisco-ios-xe-wireless-go/internal/errors"
)

// Generic HTTP Operation Functions for Service Layer
// These functions provide a consistent interface for HTTP operations across all services.

// Get is a generic helper reducing boilerplate in service GET methods.
//
// The response is validated before it is decoded: a body whose sole top-level key is not the
// node the endpoint asked for, or for which T declares no field, is an error rather than a
// zero value of T. A read answered with no body is a successful read of a node that holds
// nothing: T comes back at its zero value, which for a list node is a nil slice.
func Get[T any](ctx context.Context, c *Client, endpoint string, opts ...GetOption) (*T, error) {
	if c == nil {
		return nil, errors.New(ierrors.ErrClientNil)
	}

	endpoint = applyGetOptions(endpoint, opts)

	body, err := c.Do(ctx, http.MethodGet, endpoint)
	if err != nil {
		return nil, err
	}

	if len(body) == 0 {
		var out T
		return &out, nil
	}

	return decodeSoleKey[T](body, endpoint)
}

// GetRaw performs a GET and returns the response body as the controller sent it.
//
// This is the seam behind the root client's GetData: it applies GetOption values the
// same way Get does, and leaves the envelope check to the caller.
//
// A read answered with no body is a successful read of a node that holds nothing: the
// returned slice is non-nil and empty, and the error is nil.
func GetRaw(ctx context.Context, c *Client, endpoint string, opts ...GetOption) ([]byte, error) {
	if c == nil {
		return nil, errors.New(ierrors.ErrClientNil)
	}

	return c.Do(ctx, http.MethodGet, applyGetOptions(endpoint, opts))
}

// decode unmarshals a write response body into T, treating an empty body as a zero T.
//
// A write is answered with the changed node, an RPC output, or nothing at all, so there is no
// single top-level key to hold against the endpoint; the read path validates in decodeSoleKey.
func decode[T any](body []byte) (*T, error) {
	var out T
	if len(body) == 0 {
		return &out, nil
	}

	if err := json.Unmarshal(body, &out); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &out, nil
}

// Post is a generic helper for sending POST requests with payload.
func Post[T any](ctx context.Context, c *Client, endpoint string, payload any) (*T, error) {
	if c == nil {
		return nil, errors.New(ierrors.ErrClientNil)
	}

	body, err := c.DoWithPayload(ctx, http.MethodPost, endpoint, payload)
	if err != nil {
		return nil, err
	}

	return decode[T](body)
}

// PostVoid is a generic helper for POST operations without expecting a response body.
func PostVoid(ctx context.Context, c *Client, endpoint string, payload any) error {
	if c == nil {
		return errors.New(ierrors.ErrClientNil)
	}
	_, err := c.DoWithPayload(ctx, http.MethodPost, endpoint, payload)
	return err
}

// PostRPCVoid is a generic helper for RPC POST operations without expecting a response body.
func PostRPCVoid(ctx context.Context, c *Client, rpcEndpoint string, payload any) error {
	if c == nil {
		return errors.New(ierrors.ErrClientNil)
	}
	_, err := c.DoRPCWithPayload(ctx, http.MethodPost, rpcEndpoint, payload)
	return err
}

// Put is a generic helper for PUT operations that expect a response body.
func Put[T any](ctx context.Context, c *Client, endpoint string, payload any) (*T, error) {
	if c == nil {
		return nil, errors.New(ierrors.ErrClientNil)
	}

	body, err := c.DoWithPayload(ctx, http.MethodPut, endpoint, payload)
	if err != nil {
		return nil, err
	}

	return decode[T](body)
}

// PutVoid is a generic helper for PUT operations without expecting a response body.
func PutVoid(ctx context.Context, c *Client, endpoint string, payload any) error {
	if c == nil {
		return errors.New(ierrors.ErrClientNil)
	}
	_, err := c.DoWithPayload(ctx, http.MethodPut, endpoint, payload)
	return err
}

// Patch is a generic helper for PATCH operations that expect a response body.
func Patch[T any](ctx context.Context, c *Client, endpoint string, payload any) (*T, error) {
	if c == nil {
		return nil, errors.New(ierrors.ErrClientNil)
	}

	body, err := c.DoWithPayload(ctx, http.MethodPatch, endpoint, payload)
	if err != nil {
		return nil, err
	}

	return decode[T](body)
}

// PatchVoid is a generic helper for PATCH operations without expecting a response body.
func PatchVoid(ctx context.Context, c *Client, endpoint string, payload any) error {
	if c == nil {
		return errors.New(ierrors.ErrClientNil)
	}
	_, err := c.DoWithPayload(ctx, http.MethodPatch, endpoint, payload)
	return err
}

// Delete is a generic helper for DELETE operations without expecting a response body.
func Delete(ctx context.Context, c *Client, endpoint string) error {
	if c == nil {
		return errors.New(ierrors.ErrClientNil)
	}
	_, err := c.Do(ctx, http.MethodDelete, endpoint)
	return err
}
