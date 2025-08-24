// Package helpers provides helper functions for tag services
package helpers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	ierrors "github.com/umatare5/cisco-ios-xe-wireless-go/internal/errors"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/validation"
)

// TagCRUDConfig represents configuration for tag CRUD operations
type TagCRUDConfig struct {
	BasePath           string
	ListPath           string
	YANGPrefix         string
	TagNameField       string
	ValidatorFunc      func(string) bool
	ValidationErrorKey string
}

// ParseSiteTagJSONResponse parses JSON response for site tags that can be either array or single object format
func ParseSiteTagJSONResponse(body []byte, yangPrefix, tagName string) (any, error) {
	if len(body) == 0 {
		return nil, nil
	}

	// Parse JSON response
	var result any
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	// Parse the raw JSON response
	jsonBytes, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response for tag %s: %w", tagName, err)
	}

	// Try array format first - this is needed for site tags
	var arrayResp map[string]json.RawMessage
	if err := json.Unmarshal(jsonBytes, &arrayResp); err == nil {
		if itemBytes, exists := arrayResp[yangPrefix]; exists {
			var items []any
			if err := json.Unmarshal(itemBytes, &items); err == nil && len(items) > 0 {
				return items[0], nil
			}
			// If not array, try as single object
			var singleItem any
			if err := json.Unmarshal(itemBytes, &singleItem); err == nil {
				return singleItem, nil
			}
		}
	}

	// If both formats failed, return nil (not found or unrecognizable format)
	return nil, nil
}

// ParsePolicyTagJSONResponse parses JSON response for policy tags that can be either array or single object format
func ParsePolicyTagJSONResponse(body []byte, yangPrefix, tagName string) (any, error) {
	if !isValidBody(body) {
		return nil, nil
	}

	result, err := parseJSONBody(body)
	if err != nil {
		return nil, err
	}

	return extractTagData(result, yangPrefix, tagName)
}

// Private helper functions for JSON processing

// parseJSONBody parses raw JSON bytes into an interface
func parseJSONBody(body []byte) (any, error) {
	var result any
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return result, nil
}

// extractTagData extracts tag data from parsed JSON response
func extractTagData(result any, yangPrefix, tagName string) (any, error) {
	// Parse the raw JSON response
	jsonBytes, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response for tag %s: %w", tagName, err)
	}

	// Try array format first - this is needed for policy tags
	var arrayResp map[string]json.RawMessage
	if err := json.Unmarshal(jsonBytes, &arrayResp); err == nil {
		if itemBytes, exists := arrayResp[yangPrefix]; exists {
			// Try array format
			if item := tryParseAsArray(itemBytes); item != nil {
				return item, nil
			}
			// Try single object format
			if item := tryParseAsSingleObject(itemBytes); item != nil {
				return item, nil
			}
		}
	}

	// If both formats failed, return nil (not found or unrecognizable format)
	return nil, nil
}

// tryParseAsArray attempts to parse as array and returns first item
func tryParseAsArray(itemBytes json.RawMessage) any {
	var items []any
	if err := json.Unmarshal(itemBytes, &items); err == nil && len(items) > 0 {
		return items[0]
	}
	return nil
}

// tryParseAsSingleObject attempts to parse as single object
func tryParseAsSingleObject(itemBytes json.RawMessage) any {
	var singleItem any
	if err := json.Unmarshal(itemBytes, &singleItem); err == nil {
		return singleItem
	}
	return nil
}

// ValidateTagName validates a tag name using the provided validator function
func ValidateTagName(tagName, tagType string, validatorFunc func(string) bool) error {
	// Use validation.IsNonEmptyString for string validation
	if !validation.IsNonEmptyString(tagName) {
		return ierrors.ValidationError(tagType+" tag name", "empty string")
	}

	// Use the specific validator if provided
	if validatorFunc != nil && !validatorFunc(tagName) {
		return ierrors.ValidationError(tagType+" tag name", tagName)
	}
	return nil
}

// BuildTagPayload builds a payload for tag operations with the specified YANG prefix
func BuildTagPayload(config any, yangPrefix string) map[string]any {
	return map[string]any{
		yangPrefix: config,
	}
}

// TagCRUDOperations executes common CRUD operations for tag services
type TagCRUDOperations struct {
	config TagCRUDConfig
	client *core.Client
}

// NewTagCRUDOperations creates a new tag CRUD operations helper
func NewTagCRUDOperations(config TagCRUDConfig, client *core.Client) *TagCRUDOperations {
	return &TagCRUDOperations{
		config: config,
		client: client,
	}
}

// Create creates a new tag configuration
func (t *TagCRUDOperations) Create(ctx context.Context, config any, tagName string) error {
	payload, err := t.validateAndPreparePayload(config, tagName)
	if err != nil {
		return err
	}

	return t.client.Post(ctx, t.config.ListPath, payload)
}

// Update updates an existing tag configuration
func (t *TagCRUDOperations) Update(ctx context.Context, config any, tagName string) error {
	payload, err := t.validateAndPreparePayload(config, tagName)
	if err != nil {
		return err
	}

	tagURL := t.buildTagURL(tagName)
	return core.PatchVoid(ctx, t.client, tagURL, payload)
}

// Delete deletes a tag configuration
func (t *TagCRUDOperations) Delete(ctx context.Context, tagName string) error {
	if err := ValidateTagName(tagName, t.config.ValidationErrorKey, t.config.ValidatorFunc); err != nil {
		return err
	}

	tagURL := t.buildTagURL(tagName)
	return core.Delete(ctx, t.client, tagURL)
}

// GetRaw gets raw JSON response for a tag
func (t *TagCRUDOperations) GetRaw(ctx context.Context, tagName string) ([]byte, error) {
	if err := ValidateTagName(tagName, t.config.ValidationErrorKey, t.config.ValidatorFunc); err != nil {
		return nil, err
	}

	return t.fetchTagData(ctx, tagName)
}

// Get retrieves raw tag configuration data
func (t *TagCRUDOperations) Get(ctx context.Context, tagName string) ([]byte, error) {
	return t.GetRaw(ctx, tagName)
}

// Parse processes raw tag data using the provided parser function
func (t *TagCRUDOperations) Parse(
	body []byte,
	tagName string,
	parseFunc func([]byte, string, string) (any, error),
) (any, error) {
	if !isValidBody(body) {
		return nil, nil
	}

	return parseFunc(body, t.config.YANGPrefix, tagName)
}

// ParseAndGet gets a parsed tag configuration using the provided parser function
// Deprecated: Use Get and Parse separately for better separation of concerns
func (t *TagCRUDOperations) ParseAndGet(
	ctx context.Context,
	tagName string,
	parseFunc func([]byte, string, string) (any, error),
) (any, error) {
	body, err := t.Get(ctx, tagName)
	if err != nil {
		return nil, err
	}

	return t.Parse(body, tagName, parseFunc)
}

// Private predicate functions for better code readability

// isValidBody checks if the body contains valid data using validation package
func isValidBody(body []byte) bool {
	return !validation.IsNil(body) && len(body) > 0
}

// Private helper functions for common operations

// validateAndPreparePayload validates tag name and prepares payload
func (t *TagCRUDOperations) validateAndPreparePayload(config any, tagName string) (map[string]any, error) {
	if err := ValidateTagName(tagName, t.config.ValidationErrorKey, t.config.ValidatorFunc); err != nil {
		return nil, err
	}
	return BuildTagPayload(config, t.config.YANGPrefix), nil
}

// buildTagURL builds URL for specific tag operations
func (t *TagCRUDOperations) buildTagURL(tagName string) string {
	return fmt.Sprintf("%s=%s", t.config.BasePath, url.QueryEscape(tagName))
}

// fetchTagData performs HTTP GET request for tag data with proper error handling
func (t *TagCRUDOperations) fetchTagData(ctx context.Context, tagName string) ([]byte, error) {
	tagURL := t.buildTagURL(tagName)

	body, err := t.client.Do(ctx, "GET", tagURL)
	if validation.HasError(err) {
		return nil, fmt.Errorf("failed to get tag %s: %w", tagName, err)
	}

	return body, nil
}
