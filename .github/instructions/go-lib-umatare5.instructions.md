---
description: "Go Library Development Instructions"
applyTo: "**/*.go,**/go.mod,**/go.sum"
---

# Go Library Development Instructions

GitHub Copilot **MUST** follow these instructions when generating or modifying Go code in this repository.

## Scope & Metadata

- **Last Updated**: 2025-09-14
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
  3. Constructor (**return a concrete type**): `func NewService(c *core.Client) Service` (returns concrete type, not pointer).
  4. Configuration methods (`GetConfig`, `ListTagConfigs`).
  5. Operational methods (`GetOperational`, `GetGlobalOperational`, `ListCAPWAPData`).
  6. Administrative methods (`Enable*`, `Disable*`, `Assign*`, `Reload`, etc.).
  7. Internal helpers (private).

- **SS-002 (SHOULD)** Group related methods by basic operations, filtered operations, and field-projection operations.

---

## 3. Testing Framework

### 3.A Test Structure

- **TS-001 (MUST)** Organize unit tests per service using table-driven tests and `t.Run` subtests. Use multiple `*_test.go` files by topic when it improves clarity (e.g., constructor, get-ops, errors, validation).
- **TS-002 (MUST)** Use `pkg/testutil` unified API (`NewMockServer` with functional options: `WithSuccessResponses`, `WithErrorResponses`, `WithCustomResponse`) with standard library `testing`/`httptest`.
- **TS-003 (MUST)** Simulate IOS-XE version gaps (e.g., 17.18.1+ features such as WAT, URWB, Spaces) with custom handlers returning representative status codes (e.g., 404) for unsupported versions.
- **TS-004 (MUST)** Base mock responses on real WNC data when available and include the source context in comments.
- **TS-005 (SHOULD)** Place integration tests under `tests/integration/{service}_*_test.go` and name them per **TNL-001**. Use `tests/testutil/integration` for integration-specific configuration and setup utilities.
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
type HandlerMap[T any] = map[string]map[string]T
// Used in actual testutil implementation:
// - HandlerMap[func() (int, string)] for RESTCONFServer
// - HandlerMap[string] for endpoint response mappings
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
- **NET-003 (SHOULD)** Use shared generic helpers (`core.Get[T]`, `core.Post[T]`, `core.PostVoid`, `core.PutVoid`, `core.PostRPCVoid`, etc.).
- **NET-004 (SHOULD)** Use the client's RPC helper for YANG RPCs (`core.PostRPCVoid`).
- **NET-005 (SHOULD)** Normalize identifiers before URL composition (e.g., MAC formatting using `validation.NormalizeMACAddress`).
- **NET-006 (MUST)** Handle HTTP 204 (No Content) responses appropriately with specialized response types instead of treating them as errors.

---

## 7. Package & Module Structure

- **PKG-001 (MUST)** Organize by domain under `service/*`; place shared packages under `internal/*` and reusable public helpers under `pkg/*`.

  - `internal/restconf/routes` — **single source of truth** for endpoint constants and builders.
  - `internal/core` — transport helpers, generic HTTP, RPC facade.
  - `internal/validation` — validators and normalizers.
  - `internal/model/*` — wire models with docs.
  - `internal/testutil` — internal assertion functions for unit testing.
  - `pkg/testutil` — public test utilities (mock servers, testing helpers).
  - `tests/testutil/integration` — integration-specific configuration and setup utilities.

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

- **END-001 (MUST)** Keep endpoint routes centralized in `internal/restconf/routes`. Services may import local aliases but **must not** duplicate endpoint strings.
- **END-002 (MUST)** Build field-projection and leaf URLs via centralized helpers in `internal/restconf/routes`.

---

## 10. Documentation

- **DOC-001 (MUST)** Provide a package comment with Overview, Main Features, Usage Example, Known Limitations, Error Handling, and Requirements.
- **DOC-002 (MUST)** List implemented API only in Main Features.
- **DOC-003 (MUST)** Describe unsupported endpoints in Known Limitations.
- **DOC-004 (SHOULD)** Include a minimal example using the root client and a service accessor.
- **DOC-005 (SHOULD)** Start exported identifiers’ comments with the identifier name.
- **DOC-006 (SHOULD)** Keep `doc.go` package documentation concise (≤ 5 lines) unless special circumstances require detailed explanations.

---

## 11. Tag Management

- **TAG-001 (MUST)** Use dedicated tag service types (e.g., `PolicyTagService`, `RFTagService`, `SiteTagService`) with standardized CRUD method set.
- **TAG-002 (MUST)** Use **TagName** as the identifier field across tag structs.
- **TAG-003 (MUST)** Place tag operations in logically split files using the `tag_*.go` pattern.
- **TAG-004 (MUST)** Enforce YANG constraints (e.g., tag name length ≤ 32 characters).
- **TAG-005 (MUST)** Provide direct tag service accessors from main client (e.g., `Client.PolicyTag()`, `Client.RFTag()`).

---

## 12. Versioning & Modules

- **VER-001 (MUST)** Apply Semantic Import Versioning for v2+ (`module .../v2`) and update import paths accordingly.
- **VER-002 (MUST)** Maintain `go.mod` and run `go mod tidy` in CI.

---

## 13. General Coding Standards

- **GCS-001 (MUST)** Apply **KISS/DRY** consistently and keep public APIs minimal and composable.
- **GCS-002 (SHOULD)** Prefer Go 1.21+/1.22+/1.23+ idioms (`any`, `min/max/clear`, per-iteration `for` variables, iterator `range`) **when they improve clarity or performance**. Avoid gratuitous rewrites solely to use new syntax.
- **GCS-003 (MUST)** Prefer **early return** patterns in conditionals and loops to keep blocks shallow.
- **GCS-004 (MUST)** Organize packages with clear boundaries (service, model, transport, testutil) and eliminate cyclic deps.

---

## 14. Advanced API Design Standards

- **AADS-001 (MUST)** Accept `context.Context` as the first parameter in all public methods.
- **AADS-002 (MUST)** Use functional options for optional behavior; keep constructors simple. Make zero-value usable for simple value/config types (e.g., options); network clients **require** constructors when mandatory parameters exist.
- **AADS-003 (MUST)** Return typed errors and wrap with `%w`; compare with `errors.Is/As`.
- **AADS-004 (MUST)** Provide concurrency-safe clients; document goroutine safety where relevant.
- **AADS-005 (SHOULD)** Avoid internal logging. Return rich, wrapped errors. If logging is needed, accept an optional minimal logger interface from callers; default is silent.

---

## 15. Advanced Testing Standards

- **ATS-001 (MUST)** Use small package-local helpers in `_test.go`. Put cross-package test servers in `pkg/testutil` (using `NewRESTCONFSuccessServer`, `NewRESTCONFErrorServer`), internal assertion functions in `internal/testutil`, and integration-specific utilities in `tests/testutil/integration` (all stdlib `testing/httptest`).
- **ATS-002 (MUST)** Base mock payloads on **real WNC JSON structures** captured from supported versions. Use comprehensive endpoint mapping in test responses.
- **ATS-003 (MUST)** For features introduced in IOS-XE 17.18.1+ (e.g., WAT, URWB, Spaces), include **version-compatibility scenarios** using `NewRESTCONFErrorServer`.
- **ATS-004 (MUST)** Target **≥ 90% unit test coverage per service**, and keep CI fully green.
- **ATS-005 (MUST)** Create test methods for Constructor, GetOperations, SetOperations, ErrorHandling, and ValidationErrors categories.

---

## 22. Unified Testing API Standards

- **UTA-001 (MUST)** Use the unified `testutil.NewMockServer` API with functional options exclusively. Legacy functions are removed.
- **UTA-002 (MUST)** Apply functional options pattern for test configuration:
  - `WithSuccessResponses(map[string]string)` for successful GET operations
  - `WithErrorResponses([]string, int)` for error scenarios with specific status codes
  - `WithCustomResponse(string, ResponseConfig)` for complex response configurations
  - `WithTesting(*testing.T)` for enhanced test integration
- **UTA-003 (MUST)** Combine multiple options in single `NewMockServer` call for complex test scenarios.
- **UTA-004 (MUST)** Use `defer mockServer.Close()` immediately after server creation.
- **UTA-005 (SHOULD)** Prefer `WithSuccessResponses` for straightforward mock data, `WithErrorResponses` for error testing, and `WithCustomResponse` for advanced scenarios requiring specific HTTP methods or status codes.
- **UTA-006 (MUST)** Structure mock responses based on real WNC JSON data with appropriate RESTCONF path mapping.

---

## 16. Test Naming & Layout (Unified)

- **TNL-001 (MUST)** Use standardized function names:

  - `Test{Service}Unit_{Category}_{Scenario}`
  - `Test{Service}Integration_{Category}_{Scenario}`

- **TNL-002 (MUST)** Apply consistent service names: `ApServiceUnit`, `WlanServiceUnit`, `ClientServiceUnit`, `AfcServiceUnit`, etc. (exact service names in codebase).
- **TNL-003 (MUST)** Use standard categories: `Constructor`, `GetOperations`, `SetOperations`, `ErrorHandling`, `ValidationErrors`.
- **TNL-004 (MUST)** Use standard scenarios: `Success`, `MockSuccess`, `ErrorExpected`, `RealDataSuccess`, `EmptyMAC` (or service-appropriate `EmptyParam`).
- **TNL-005 (MUST)** Follow placement conventions:

  - Unit tests: `service/{service}/service_test.go` (source co-located).
  - Integration tests: `tests/integration/{service}_service_test.go`.
  - Test utilities: `pkg/testutil/` (public), `internal/testutil/` (internal), `tests/testutil/integration/` (integration-specific).

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
// internal/restconf/routes/ap.go
package routes

// AP Configuration Paths
const (
    // APCfgPath retrieves complete access point configuration data
    APCfgPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-ap-cfg:ap-cfg-data"

    // APTagsPath retrieves access point tag configurations
    APTagsPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-ap-cfg:ap-cfg-data/ap-tags"

    // APTagPath retrieves specific access point tag configuration
    APTagPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-ap-cfg:ap-cfg-data/ap-tag"
)
```

```go
// service/ap/service.go
package ap

import (
    "context"

    "github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
    model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/ap"
    "github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
    "github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

type Service struct {
    service.BaseService
}

func NewService(client *core.Client) Service {
    return Service{BaseService: service.NewBaseService(client)}
}

func (s Service) GetConfig(ctx context.Context) (*model.ApCfg, error) {
    return core.Get[model.ApCfg](ctx, s.Client(), routes.APCfgPath)
}

func (s Service) GetTagConfigByMAC(ctx context.Context, mac string) (*model.ApCfgApTag, error) {
    // Validation and MAC normalization...
    url := s.Client().RestconfBuilder().BuildPathQueryURL(routes.APTagsPath, "ap-tag", normalizedMAC)
    return core.Get[model.ApCfgApTag](ctx, s.Client(), url)
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
// Remote read operations (OK to use Get*)
func (s Service) GetConfig(ctx context.Context) (*model.ApCfg, error)
func (s Service) GetTagConfigByMAC(ctx context.Context, mac string) (*model.ApCfgApTag, error)
func (s Service) GetAPJoinStatsByWTPMAC(ctx context.Context, mac string) (*model.ApGlobalOperApJoinStats, error)

// List operations for collections
func (s Service) ListTagConfigs(ctx context.Context) (*model.ApCfgApTag, error)
func (s Service) ListAPHistory(ctx context.Context) (*model.ApGlobalOperApHistory, error)

// Trivial accessor (avoid Get prefix)
func (c *Client) Core() *core.Client
```

## 21. Test Helpers — Example

````go
// service/ap/service_test.go (local helper using unified API)
func TestApServiceUnit_Constructor_Success(t *testing.T) {
    // Test implementation with unified MockServer API
    mockServer := testutil.NewMockServer(
        testutil.WithSuccessResponses(map[string]string{
            "test-endpoint": `{"status": "success"}`,
        }),
        testutil.WithTesting(t),
    )
    defer mockServer.Close()
    // ... rest of test
}

func TestApServiceUnit_GetOperations_ErrorHandling(t *testing.T) {
    // Error response testing
    mockServer := testutil.NewMockServer(
        testutil.WithErrorResponses([]string{
            "Cisco-IOS-XE-wireless-ap-cfg:ap-cfg-data",
        }, 404),
    )
    defer mockServer.Close()
    // ... test error scenario
}

func TestApServiceUnit_SetOperations_CustomResponse(t *testing.T) {
    // Custom response testing
    mockServer := testutil.NewMockServer(
        testutil.WithCustomResponse("custom-endpoint", testutil.ResponseConfig{
            StatusCode: 202,
            Body:       `{"custom": "response"}`,
            Method:     "POST",
        }),
    )
    defer mockServer.Close()
    // ... test custom response
}

```go
// pkg/testutil/testing.go (unified public test utilities)
func NewMockServer(opts ...MockServerOption) MockServer { /* ... */ }
func WithSuccessResponses(responses map[string]string) MockServerOption { /* ... */ }
func WithErrorResponses(paths []string, statusCode int) MockServerOption { /* ... */ }
func WithCustomResponse(path string, config ResponseConfig) MockServerOption { /* ... */ }
func WithTesting(t *testing.T) MockServerOption { /* ... */ }

// internal/testutil/helper.go (internal assertions)
func AssertStringEquals(t *testing.T, actual, expected, message string) { t.Helper(); /* ... */ }

// tests/testutil/integration/config.go (integration config)
func LoadConfig() (*Config, error) { /* ... */ }
````
