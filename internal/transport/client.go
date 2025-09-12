// Package transport provides HTTP client utilities and request creation functions.
package transport

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf"
)

// RequestBuilder provides HTTP request creation utilities.
type RequestBuilder struct {
	restBuilder *restconf.Builder
	token       string
	logger      *slog.Logger
}

// NewRequestBuilder creates a new RequestBuilder instance.
func NewRequestBuilder(restBuilder *restconf.Builder, token string, logger *slog.Logger) *RequestBuilder {
	return &RequestBuilder{
		restBuilder: restBuilder,
		token:       token,
		logger:      logger,
	}
}

// CreateRequest creates and configures an HTTP request.
func (rb *RequestBuilder) CreateRequest(ctx context.Context, method, path string) (*http.Request, error) {
	if rb.restBuilder == nil {
		return nil, errors.New("RESTCONF builder is not properly initialized")
	}

	url := rb.restBuilder.BuildDataURL(path)

	req, err := http.NewRequestWithContext(ctx, method, url, http.NoBody)
	if err != nil {
		rb.logger.Error("Failed to create HTTP request", "error", err, "url", url)
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header = DefaultHeaders(rb.token)
	rb.logger.Debug("Sending API request", "method", method, "url", url)
	return req, nil
}

// CreateRequestWithPayload creates and configures an HTTP request with a JSON payload.
func (rb *RequestBuilder) CreateRequestWithPayload(
	ctx context.Context,
	method, path string,
	payload any,
) (*http.Request, error) {
	if rb.restBuilder == nil {
		return nil, errors.New("RESTCONF builder is not properly initialized")
	}

	// Handle nil payload case early - reuse CreateRequest logic
	if payload == nil {
		return rb.CreateRequest(ctx, method, path)
	}

	url := rb.restBuilder.BuildDataURL(path)

	// Marshal payload and create request
	return rb.createRequestWithJSONPayload(ctx, method, url, payload, "")
}

// CreateRPCRequestWithPayload creates and configures an HTTP RPC request with a JSON payload.
func (rb *RequestBuilder) CreateRPCRequestWithPayload(
	ctx context.Context,
	method, rpcPath string,
	payload any,
) (*http.Request, error) {
	if rb.restBuilder == nil {
		return nil, errors.New("RESTCONF builder is not properly initialized")
	}

	// Handle nil payload case early
	if payload == nil {
		url := rb.restBuilder.BuildOperationsURL(rpcPath)
		req, err := http.NewRequestWithContext(ctx, method, url, http.NoBody)
		if err != nil {
			rb.logger.Error("Failed to create HTTP RPC request", "error", err, "url", url)
			return nil, fmt.Errorf("failed to create RPC request: %w", err)
		}
		req.Header = DefaultHeaders(rb.token)
		rb.logger.Debug("Sending RPC request", "method", method, "url", url)
		return req, nil
	}

	url := rb.restBuilder.BuildOperationsURL(rpcPath)
	// Marshal payload and create RPC request
	return rb.createRequestWithJSONPayload(ctx, method, url, payload, "RPC")
}

// ExecuteRequest executes an HTTP request and handles basic error conditions.
func (rb *RequestBuilder) ExecuteRequest(httpClient *http.Client, req *http.Request) (*http.Response, error) {
	if req == nil {
		return nil, errors.New("request cannot be nil")
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		rb.logger.Error("HTTP request failed", "error", err, "url", req.URL.String())
		return nil, fmt.Errorf("request failed: %w", err)
	}
	return resp, nil
}

// createRequestWithJSONPayload creates an HTTP request with JSON payload and proper headers.
func (rb *RequestBuilder) createRequestWithJSONPayload(
	ctx context.Context,
	method, url string,
	payload any,
	logType string,
) (*http.Request, error) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	body := bytes.NewReader(jsonData)
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		rb.logger.Error("Failed to create HTTP "+logType+" request", "error", err, "url", url)
		return nil, fmt.Errorf("failed to create "+logType+" request: %w", err)
	}

	req.Header = DefaultHeaders(rb.token)
	req.Header.Set("Content-Type", HTTPHeaderValueYANGData)
	rb.logger.Debug("Sending "+logType+" request", "method", method, "url", url)
	return req, nil
}
