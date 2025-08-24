// Package errors provides common error definitions and utilities for Cisco IOS-XE Wireless Network Controller API services.
//
// This package standardizes error handling patterns across all domain-specific
// service implementations, providing consistent error messaging, templates,
// and utility functions for error creation and handling.
//
// # Components
//
// The package includes:
//   - Common error message templates for standardized formatting
//   - Entity type constants for consistent error context
//   - Action constants for consistent operation naming
//   - Utility functions for error creation and validation
//
// # Usage
//
// Services should use the provided templates and constants for consistent error messaging:
//
//	import ierrors "github.com/umatare5/cisco-ios-xe-wireless-go/internal/errors"
//
//	err := fmt.Errorf(ierrors.ErrOperationFailedTemplate,
//		ierrors.ActionGet, ierrors.EntityTypeAP, apMAC, underlyingErr)
//
// # Design Principles
//
// The error infrastructure follows these principles:
//   - Consistency: All services use the same error message patterns
//   - Context: Errors include sufficient context for debugging
//   - Composability: Templates can be combined for complex scenarios
//   - Maintainability: Centralized error definitions reduce duplication
package errors
