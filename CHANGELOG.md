# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [v1.5.0] – 2025-08-06

### Added

* Domain Service layer with client.<Domain>().<Method>() pattern
* Service implementations for AFC, AP, Client, General, RRM, WLAN, Rogue, mDNS, NMSP, Hyperlocation, Geolocation, and Mcast domains
* Three-layer architecture separating Core, Domain Service, and Generated Type layers
* `internal/model` package with isolated data structure definitions
* `internal/httpx` and `internal/restconf` packages for HTTP utilities
* Functional options pattern for client configuration (`wnc.New()`, `wnc.WithTimeout()`, etc.)
* Comprehensive package documentation with `doc.go` files
* Service-based examples in README.md

### Changed

* **BREAKING**: Client constructor now uses `wnc.New()` with functional options instead of `wnc.NewClient(config)`
* All domain endpoints now accessible through service pattern (e.g., `client.AFC().Oper(ctx)`)
* Model structs moved from domain packages to `internal/model`
* HTTP client logic consolidated in `internal/httpx/transport.go`
* RESTCONF URL building moved to `internal/restconf/builder.go`

### Deprecated

* Legacy helper functions (`afc.GetAfcOper`, `ap.GetApOper`, etc.) - **removal planned for v2.0.0**
* Large API interfaces (`WirelessControllerAPI`, `AccessPointAPI`, etc.) - **removal planned for v2.0.0**
* Direct `SendAPIRequest` usage - use service methods instead
* `Config` struct and `NewClient(config)` - use `wnc.New()` with functional options

### Removed

* Obsolete HTTP helper methods; logic moved to internal packages
* Direct exposure of internal HTTP implementation details

### Fixed

* Endpoint path duplication issues (`/restconf/data` prefix conflicts)
* Hardware-specific endpoint error handling (proper 404 handling for unavailable features)
* Package import conflicts and circular dependencies
* Test coverage for hardware-specific scenarios

## [v0.1.0] – 2025-01-xx

### Added

* Initial release with basic RESTCONF API support
* Legacy helper functions for major domains
* HTTP client with authentication and TLS support
* Basic error handling and logging
