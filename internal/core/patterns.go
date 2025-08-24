// Package core provides common RESTCONF operation patterns.
package core

import (
	"context"
	"fmt"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/errors"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/validation"
)

// convertStringsToInterfaces converts a slice of strings to a slice of interfaces
func convertStringsToInterfaces(strs []string) []interface{} {
	result := make([]interface{}, len(strs))
	for i, str := range strs {
		result[i] = str
	}
	return result
}

// ConfigOperations provides common configuration data retrieval patterns.
type ConfigOperations[T any] struct {
	client   *Client
	basePath string
}

// NewConfigOperations creates a new configuration operations instance.
func NewConfigOperations[T any](client *Client, basePath string) *ConfigOperations[T] {
	return &ConfigOperations[T]{
		client:   client,
		basePath: basePath,
	}
}

// GetAll retrieves complete configuration data.
func (ops *ConfigOperations[T]) GetAll(ctx context.Context) (*T, error) {
	return Get[T](ctx, ops.client, ops.basePath)
}

// GetByID retrieves configuration data filtered by a specific ID.
func (ops *ConfigOperations[T]) GetByID(ctx context.Context, listName, id string) (*T, error) {
	url := ops.client.RestconfBuilder().BuildPathQueryURL(ops.basePath, listName, id)
	return Get[T](ctx, ops.client, url)
}

// GetOnlyFields retrieves only specific fields from configuration data.
func (ops *ConfigOperations[T]) GetOnlyFields(ctx context.Context, fields string) (*T, error) {
	url := ops.client.RestconfBuilder().BuildFieldsURL(ops.basePath, fields)
	return Get[T](ctx, ops.client, url)
}

// GetSubRes retrieves a specific subresource from configuration data.
func (ops *ConfigOperations[T]) GetSubRes(ctx context.Context, subresource string) (*T, error) {
	url := fmt.Sprintf("%s/%s", ops.basePath, subresource)
	return Get[T](ctx, ops.client, url)
}

// GetByCompositeKey retrieves configuration data filtered by composite key (e.g., MAC,slot).
func (ops *ConfigOperations[T]) GetByCompositeKey(
	ctx context.Context, listName string, keys ...string,
) (*T, error) {
	if len(keys) == 0 {
		return nil, errors.RequiredParameterError("keys")
	}
	url := ops.client.RestconfBuilder().BuildCompositeKeyURL(
		ops.basePath, listName, convertStringsToInterfaces(keys)...)
	return Get[T](ctx, ops.client, url)
}

// GetByMAC retrieves configuration data filtered by MAC address with validation.
func (ops *ConfigOperations[T]) GetByMAC(ctx context.Context, listName, mac string) (*T, error) {
	if err := validation.ValidateAPMac(mac); err != nil {
		return nil, errors.SimpleServiceError("invalid MAC address format", err)
	}
	return ops.GetByID(ctx, listName, mac)
}

// OperationalOperations provides common operational data retrieval patterns.
type OperationalOperations[T any] struct {
	client   *Client
	basePath string
}

// NewOperationalOperations creates a new operational operations instance.
func NewOperationalOperations[T any](client *Client, basePath string) *OperationalOperations[T] {
	return &OperationalOperations[T]{
		client:   client,
		basePath: basePath,
	}
}

// GetAll retrieves complete operational data.
func (ops *OperationalOperations[T]) GetAll(ctx context.Context) (*T, error) {
	return Get[T](ctx, ops.client, ops.basePath)
}

// GetByID retrieves operational data filtered by a specific ID.
func (ops *OperationalOperations[T]) GetByID(ctx context.Context, listName, id string) (*T, error) {
	url := ops.client.RestconfBuilder().BuildPathQueryURL(ops.basePath, listName, id)
	return Get[T](ctx, ops.client, url)
}

// GetByMAC retrieves operational data filtered by MAC address with validation.
func (ops *OperationalOperations[T]) GetByMAC(ctx context.Context, listName, mac string) (*T, error) {
	if err := validation.ValidateAPMac(mac); err != nil {
		return nil, errors.SimpleServiceError("invalid MAC address format", err)
	}
	return ops.GetByID(ctx, listName, mac)
}

// GetByCompositeKey retrieves operational data filtered by composite key (e.g., MAC,slot).
func (ops *OperationalOperations[T]) GetByCompositeKey(
	ctx context.Context, listName string, keys ...string,
) (*T, error) {
	if len(keys) == 0 {
		return nil, errors.RequiredParameterError("keys")
	}
	url := ops.client.RestconfBuilder().BuildCompositeKeyURL(
		ops.basePath, listName, convertStringsToInterfaces(keys)...)
	return Get[T](ctx, ops.client, url)
}

// GetOnlyFields retrieves only specific fields from operational data.
func (ops *OperationalOperations[T]) GetOnlyFields(ctx context.Context, fields string) (*T, error) {
	url := ops.client.RestconfBuilder().BuildFieldsURL(ops.basePath, fields)
	return Get[T](ctx, ops.client, url)
}
