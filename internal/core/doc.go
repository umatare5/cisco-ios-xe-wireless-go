// Package core provides the foundational HTTP client and transport layer for Cisco IOS-XE Wireless Controller SDK.
//
// Contains the primary Client with connection pooling, generic HTTP helpers (Get[T], Post[T], Put[T]),
// wireless domain types (RadioBand, admin states), and structured error handling (APIError, HTTPError).
// Serves as the central foundation for all service-specific operations via RESTCONF API.
package core
