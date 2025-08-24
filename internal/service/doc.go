// Package service provides common service infrastructure for Cisco IOS-XE Wireless Network Controller API services.
package service

// Documentation for the internal service package.
//
// This package provides common infrastructure and utilities that are shared
// across all domain-specific service implementations in the Cisco IOS-XE
// Wireless Network Controller API client library.
//
// # Components
//
// The package includes:
//   - BaseService: Common service structure for embedding
//   - Common error definitions and patterns
//   - Shared validation and utility functions
//
// # Usage
//
// Domain services should embed BaseService to inherit common functionality:
//
//	type Service struct {
//		service.BaseService
//	}
//
//	func NewService(client *core.Client) Service {
//		return Service{
//			BaseService: service.NewBaseService(client),
//		}
//	}
//
// # Design Principles
//
// The service infrastructure follows these principles:
//   - Consistency: All services share common patterns and error handling
//   - Composability: Services can embed and extend base functionality
//   - Validation: Standardized client and configuration validation
//   - Maintainability: Centralized common code reduces duplication
