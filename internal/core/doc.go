// Package core provides the foundational HTTP client and transport layer for Cisco IOS-XE Wireless Controller SDK.
//
// Contains the primary Client with connection pooling, generic HTTP helpers (Get[T], Post[T], Put[T]),
// the admin-state spellings the AP write RPCs send, and structured error handling (APIError).
// Serves as the central foundation for all service-specific operations via RESTCONF API.
package core
