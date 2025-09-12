---
description: "Go Library Development Instructions"
applyTo: "**/*.go,**/go.mod,**/go.sum"
---

# Go Library Development Instructions

GitHub Copilot **MUST** follow these instructions when generating or modifying Go code in this repository.

## Scope & Metadata

- **Last Updated**: 2025-09-13
- **Precedence**: 1. `copilot-instructions.md` (Global) → 2. `go.instructions.md` (Community) → 3. `go-lib-umatare5.instructions.md` (This)
- **Compatibility**: Go 1.25+ cross-platform
- **Style Base**: [Effective Go](https://go.dev/doc/effective_go), [Go Code Review Comments](https://go.dev/wiki/CodeReviewComments)
- **Goal**: Service-accessor–centric stable API, Functional Options, context-first networking, idiomatic Go, and error wrapping.
- **Expert Persona**: Operate as a **Go 1.25+ expert** with deep knowledge of modern Go idioms, patterns, and best practices.

---

## 1. Architecture

- **PK-001 (MUST)** Expose service accessors from a single root client (e.g., `Client.AP()`, `Client.WLAN()`), with services under `service/`.
- **PK-002 (MUST)** Provide Functional Options in primary constructors: `NewClient(controller, token string, opts ...Option)`.
- **PK-003 (MUST)** Place functionality-based packages under `service/*` and shared code under `internal/*`.
- **PK-004 (SHOULD)** Implement thin service receivers that delegate transport to a shared HTTP helper.
- **PK-005 (MUST)** Keep cross-domain logic in `internal/*` and keep each service focused on its own domain boundary.
- **PK-006 (MUST)** Resolve cross-service data dependencies with private helpers while keeping the public API decoupled.
- **PK-007 (SHOULD)** Provide type-safe integration interfaces with internal normalization and contextual errors.
- **PK-008 (MAY)** Offer direct service constructors (e.g., `NewService(*core.Client)`) for tests and advanced use.

---

## 2. Service Structure Patterns

- **SS-001 (MUST)** Structure each service consistently in `service.go` (or logically split files with the same order):

  1. **(Optional)** Local aliases referencing central routes (see `internal/restconf/routes`).
  2. Type definitions (service struct, enums).
  3. Constructor (**return a concrete type**): `func NewService(c *core.Client) *Service`.
  4. Configuration methods (`GetCfg*`).
  5. Operational methods (`GetOper*`, `GetGlobalOper*`).
  6. Administrative methods (`Enable*`, `Disable*`, `Assign*`, `Reload`, etc.).
  7. Internal helpers (private).

- **SS-002 (SHOULD)** Group related methods by basic operations, filtered operations, and field-projection operations.

---

## 3. Testing Framework

### 3.A Test Structure

- **TS-001 (MUST)** Organize unit tests per service using table-driven tests and `t.Run` subtests. Use multiple `*_test.go` files by topic when it improves clarity (e.g., constructor, get-ops, errors, validation).
- **TS-002 (MUST)** Use `pkg/testutil` mock servers (`NewMockServer`, `NewMockErrorServer`, `NewMockServerWithCustomErrors`) with standard library `testing`/`httptest`.
- **TS-003 (MUST)** Simulate IOS-XE version gaps (e.g., 17.18.1+ features such as WAT, URWB, Spaces) with custom handlers returning representative status codes (e.g., 404) for unsupported versions.
- **TS-004 (MUST)** Base mock responses on real WNC data when available and include the source context in comments.
- **TS-005 (SHOULD)** Place integration tests under `tests/integration/{service}_*_test.go` and name them per **TNL-001**.
- **TS-006 (SHOULD)** Store golden files under `testdata/` and read via `os.ReadFile`. Use `embed` only when hermetic builds are required.

### 3.B Modern Go Patterns

- **TS-007 (MUST)** Apply Go 1.21+ built-ins and Go 1.22+ control-flow features **when they improve clarity or performance**.

```go
// Emptiness check
if len(s) == 0 {
    // handle empty
}

// Range-over-int (Go 1.22+) when it reads clearly
for i := range n {
    process(i)
}

// Built-ins (Go 1.21+)
timeout := max(defaultTimeout, 30*time.Second)
clear(m) // when resetting maps/slices
```

- **TS-008 (MUST)** Define generic types as type definitions.

```go
type HandlerMap[T any] map[string]map[string]T
type MockServerHandlers = HandlerMap[string] // simple alias without type parameters
```

---

## 4. Public API Design

- **API-001 (MUST)** Keep the public surface minimal and stable.
- **API-002 (MUST)** Follow SemVer and provide migration notes for breaking changes.
- **API-003 (SHOULD)** Use `Get*By<Filter>` naming **for remote operations** (e.g., `...ByEthernetMAC`, `...ByWlanID`). Avoid `Get` for trivial accessors; prefer `X()`.
- **API-004 (MAY)** Use `*Only` suffix for field-projection reads (e.g., `GetGlobalOperApHistoryOnly`).
- **API-005 (SHOULD)** Use imperative verbs for write/admin ops (`Enable*`, `Disable*`, `Assign*`, `Set*`).
- **API-006 (SHOULD)** Use `GetCfg` for full config, `Get<Subset>Cfg` for subresources, and `GetCfg<Subset>Only` for projections.
- **API-007 (SHOULD)** Order read params as `(ctx, primaryIdentifier, qualifiers...)`.

---

## 5. Context & Error Handling

- **CTX-001 (MUST)** Thread `context.Context` through all network operations and honor deadlines/cancellation.
- **ERR-001 (MUST)** Wrap underlying errors with `%w` and include actionable context.
- **ERR-002 (MUST)** Define typed/sentinel errors in the owning package and expose them for `errors.Is/As`.
- **ERR-003 (SHOULD)** Provide a small central helper for error classification/formatting in `internal/errors`.

---

## 6. HTTP Client & Transport

- **NET-001 (MUST)** Reuse a single `http.Client` and close `resp.Body` along all paths.
- **NET-002 (MUST)** Set explicit timeouts and honor `context.Context`.
- **NET-003 (SHOULD)** Use shared generic helpers (`core.Get[T]`, `core.Post[T]`, `core.Put[T]`, etc.).
- **NET-004 (SHOULD)** Use the client's RPC helper for YANG RPCs.
- **NET-005 (SHOULD)** Normalize identifiers before URL composition (e.g., MAC formatting).

---

## 7. Package & Module Structure

- **PKG-001 (MUST)** Organize by domain under `service/*`; place shared packages under `internal/*` and reusable public helpers under `pkg/*`.

  - `internal/restconf/routes` — **single source of truth** for endpoint constants and builders.
  - `internal/core` — transport helpers, generic HTTP, RPC facade.
  - `internal/validation` — validators and normalizers.
  - `internal/model/*` — wire models with docs.
  - `pkg/testutil` — standard-library-based test utilities.

- **PKG-002 (SHOULD)** Use `tag_*.go` to structure tag operations by concern (CRUD, models, helpers).

- **PKG-003 (MUST)** Apply stutter-free naming (package/type names without redundant prefixes).

---

## 8. Function Design

- **FN-001 (MUST)** Implement service receiver methods when HTTP I/O or `s.client` access is required.
- **FN-002 (MUST)** Implement pure functions for parameter-only, side-effect–free logic.
- **FN-003 (MUST)** Rely on core-layer client validation (`core.Get/Post/...`) and keep service methods free of redundant validation calls.
- **FN-004 (SHOULD)** Normalize primary identifiers before composing URLs.

---

## 9. Constants & Endpoints

- **CONST-001 (MUST)** Keep endpoint routes centralized in `internal/restconf/routes`. Services may import local aliases but **must not** duplicate endpoint strings.
- **END-001 (SHOULD)** Define `<Feature>BasePath` in central routes and compose leaf endpoints from it.
- **END-002 (MUST)** Build field-projection and leaf URLs via centralized helpers in `internal/restconf/routes`.

---

## 10. Documentation

- **DOC-001 (MUST)** Provide a package comment with Overview, Main Features, Usage Example, Known Limitations, Error Handling, and Requirements.
- **DOC-002 (MUST)** List implemented API only in Main Features.
- **DOC-003 (MUST)** Describe unsupported endpoints in Known Limitations.
- **DOC-004 (SHOULD)** Include a minimal example using the root client and a service accessor.
- **DOC-005 (SHOULD)** Start exported identifiers’ comments with the identifier name.

---

## 11. Tag Management

- **TAG-001 (MUST)** Use **TagSetOper** terminology with a standardized CRUD method set.
- **TAG-002 (MUST)** Use **TagName** as the identifier field across tag structs.
- **TAG-003 (MUST)** Place tag operations in logically split files using the `tag_*.go` pattern.
- **TAG-004 (MUST)** Enforce YANG constraints (e.g., tag name length ≤ 32).

---

## 12. Versioning & Modules

- **VER-001 (MUST)** Apply Semantic Import Versioning for v2+ (`module .../v2`) and update import paths accordingly.
- **VER-002 (MUST)** Maintain `go.mod` and run `go mod tidy` in CI.

---

## 13. General Coding Standards

- **GCS-001 (MUST)** Apply **KISS/DRY** consistently and keep public APIs minimal and composable.
- **GCS-002 (MUST)** Prefer **named constants** and **typed enums** for domain values.
- **GCS-003 (MUST)** Use **predicate helpers** (`isX`, `hasX`) for clarity in conditionals.
- **GCS-004 (SHOULD)** Prefer Go 1.21+/1.22+/1.23+ idioms (`any`, `min/max/clear`, per-iteration `for` variables, iterator `range`) **when they improve clarity or performance**. Avoid gratuitous rewrites solely to use new syntax.
- **GCS-005 (MUST)** Prefer **early return** patterns in conditionals and loops to keep blocks shallow.
- **GCS-006 (MUST)** Organize packages with clear boundaries (service, model, transport, testutil) and eliminate cyclic deps.

---

## 14. Advanced API Design Standards

- **AADS-001 (MUST)** Accept `context.Context` as the first parameter in all public methods.
- **AADS-002 (MUST)** Use functional options for optional behavior; keep constructors simple. Make zero-value usable for simple value/config types (e.g., options); network clients **require** constructors when mandatory parameters exist.
- **AADS-003 (MUST)** Return typed errors and wrap with `%w`; compare with `errors.Is/As`.
- **AADS-004 (MUST)** Provide concurrency-safe clients; document goroutine safety where relevant.
- **AADS-005 (SHOULD)** Avoid internal logging. Return rich, wrapped errors. If logging is needed, accept an optional minimal logger interface from callers; default is silent.

---

## 15. Advanced Testing Standards

- **ATS-001 (MUST)** Use small package-local helpers in `_test.go`. Put cross-package test servers/utilities in `pkg/testutil` (stdlib `testing/httptest`).
- **ATS-002 (MUST)** Base mock payloads on **real WNC JSON structures** captured from supported versions.
- **ATS-003 (MUST)** For features introduced in IOS-XE 17.18.1+ (e.g., WAT, URWB, Spaces), include **version-compatibility scenarios** using `MockErrorServer`.
- **ATS-004 (MUST)** Target **≥ 90% unit test coverage per service**, and keep CI fully green.

---

## 16. Test Naming & Layout (Unified)

- **TNL-001 (MUST)** Use standardized function names:

  - `Test{Service}Unit_{Category}_{Scenario}`
  - `Test{Service}Integration_{Category}_{Scenario}`

- **TNL-002 (MUST)** Apply consistent service names: `Client`, `Afc`, `Wat`, `Urwb`, `Spaces` (exact service names in codebase).
- **TNL-003 (MUST)** Use standard categories: `Constructor`, `GetOperations`, `SetOperations`, `ErrorHandling`, `ValidationErrors`.
- **TNL-004 (MUST)** Use standard scenarios: `Success`, `MockSuccess`, `ErrorExpected`, `RealDataSuccess`, `EmptyMAC` (or service-appropriate `EmptyParam`).
- **TNL-005 (MUST)** Follow placement conventions:

  - Unit tests: `service/{service}/service_test.go` (source co-located).
  - Integration tests: `tests/integration/{service}_service_test.go`.

- **TNL-006 (MUST)** Apply the unified pattern exclusively and replace legacy names systematically.

---

## 17. Quality Standards

- **QS-001 (MUST)** Achieve unit test coverage ≥ 90% per service.
- **QS-002 (MUST)** Maintain zero lint violations with Go 1.25+ compliance.
- **QS-003 (MUST)** Validate against real IOS-XE WNC JSON structures where applicable.
- **QS-004 (MUST)** Eliminate anti-patterns proactively and keep PRs focused.
- **QS-005 (MUST)** Apply modern Go idioms consistently.
- **QS-006 (MUST)** Use the standard library exclusively for test assertions and HTTP mocks.
- **QS-007 (SHOULD)** Run `gofmt -s`, gofumpt, `gci`, `go vet`, and staticcheck in CI.

---

## 18. Centralized Routes & Builders — Example

```go
// internal/restconf/routes/wlan.go
package routes

// WLAN Configuration Paths
const (
    // WLANCfgPath provides the path for retrieving WLAN configuration data
    WLANCfgPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data"

    // WLANWlanCfgEntriesPath provides the path for WLAN configuration entries
    WLANWlanCfgEntriesPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/wlan-cfg-entries"

    // WLANWlanCfgEntryPath provides the path template for specific WLAN configuration entry
    WLANWlanCfgEntryPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/wlan-cfg-entries/wlan-cfg-entry"
)
```

```go
// service/wlan/service.go
package wlan

import (
    "context"

    "github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
    "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/wlan"
    "github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
    "github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

type Service struct {
    service.BaseService
}

func NewService(client *core.Client) Service {
    return Service{BaseService: service.NewBaseService(client)}
}

func (s Service) GetProfileConfig(ctx context.Context, profileName string) (*WlanCfgEntry, error) {
    url := s.Client().RestconfBuilder().BuildQueryURL(routes.WLANWlanCfgEntryPath, profileName)
    return core.Get[WlanCfgEntry](ctx, s.Client(), url)
}
```

## 19. Error Policy — Example

```go
var ErrNotFound = errors.New("not found")

func (s *Service) GetByID(ctx context.Context, id string) (*WLAN, error) {
    // ... issue request ...
    if resp.StatusCode == http.StatusNotFound {
        return nil, fmt.Errorf("wlan %q not found: %w", id, ErrNotFound)
    }
    // ...
}
```

## 20. Naming Examples

```go
// Remote read (OK to use Get*)
func (s *Service) GetByID(ctx context.Context, id string) (*WLAN, error)

// Trivial accessor (avoid Get prefix)
func (c *Client) Token() string
```

## 21. Test Helpers — Example

```go
// service/wlan/service_test.go (local helper)
func must(ctx context.Context, t *testing.T, fn func(context.Context) error) {
    t.Helper()
    if err := fn(ctx); err != nil { t.Fatal(err) }
}

// pkg/testutil/mockserver.go (shared)
func NewMockServer(handlers HandlerMap[http.HandlerFunc]) *httptest.Server { /* ... */ }
```
