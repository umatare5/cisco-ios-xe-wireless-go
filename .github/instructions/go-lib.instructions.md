---
description: Go Library Development Instructions
applyTo: "**/*.go"
---

# GitHub Copilot Agent Mode – Go Library Development Instructions

Copilot **MUST** comply with all instructions described in this document when editing or creating any Go code in this repository.

However, when there are conflicts between this document and `general.instructions.md`, **ALWAYS** prioritize the instructions in `general.instructions.md`.

---

## 🎯 Primary Goal

Contribute to the SDK/library. **DO NOT** build a standalone application.

---

## 🧭 Architecture & Design Principles

- **Three-Layer Architecture:**
  - **Core Layer**: `wnc/` package containing HTTP client and core infrastructure
  - **Domain Service Layer**: Service packages (`afc/`, `ap/`, `general/`, etc.) with business logic
  - **Generated Type Layer**: `internal/model/` package with auto-generated YANG model structs

- **Service-Based Design Pattern:**
  Each domain has a `service.go` file with `NewService(c *wnc.Client)` constructor and typed methods.
  Services use the core client's `Do()` method for all HTTP operations.
  All service methods follow: `func (s *Service) Method(ctx context.Context) (*model.ResponseType, error)`

- **Package Organization:**
  Place core client in `wnc/` package with `New()` constructor.
  Domain services in separate packages with consistent patterns.
  Internal utilities in `internal/` directory (validation, model, httpx, restconf).
  Minimize cross-package dependencies to ensure clear responsibilities.

- **Strict Dependency Injection:**
  Inject dependencies via constructors like `NewService(client *wnc.Client)`.
  **Do not** use global state or singletons for shared state/configuration.
  Pass dependencies explicitly through struct fields, function arguments, or constructors.

- **Clean API Design:**
  Export only intended public types and functions, keeping API minimal and stable.
  Define focused interfaces following Interface Segregation and Dependency Inversion principles.

---

## 🛠️ Go Coding Practices & Style

- **Follow Idiomatic Go:**
  Conform to [Effective Go](https://go.dev/doc/effective_go) and [Go Code Review Comments](https://go.dev/wiki/CodeReviewComments).

- **Style & Linting:**
  Format all code with `gofmt` and ensure it passes `golangci-lint`.
  Static analysis must pass without SA1012 warnings (use `var nilCtx context.Context` instead of `nil` literals in tests).

- **Functions:**
  Keep functions **ideally between 20 to 40 lines**.
  If a function exceeds **50 lines**, **refactor or split** it unless there are clear, justified exceptions (e.g., some tests or initialization code).
  Each function should fit within a single screen for readability and easy maintenance.

- **KISS Principle:**
  Keep implementations simple and avoid unnecessary complexity.

- **DRY Principle:**
  Factor out reusable, unexported helper functions.

- **SOLID Principles:**

  - **Single Responsibility Principle (SRP):** Each type, function, or package should have a distinct responsibility.
  - **Interface Segregation Principle (ISP):** Use small, focused interfaces ("Accept interfaces, return structs").

- **No Third-party Packages:**
  Use only the Go standard library.

- **Additional Practices:**

  - Use clear, explicit, and consistent names.
  - Prefer constants over hardcoded values.
  - Prefer early returns, minimize deep nesting and loops.
  - Only write necessary, non-redundant comments.

---

## 🏷️ Enum and Branch Logic Guidelines

- **For branching logic or representing states/kinds:**

  1. **For a single state/value:**
     Use a constant (`const`) only.
     _Do not define an enum type, use `iota`, or a `map` in this case._

     ```go
     const StatusActive = 1
     ```

  2. **For 2–5 states/branches:**
     Prefer `if-else` or `switch-case` statements for clarity.

     ```go
     switch status {
     case StatusActive:
         // ...
     case StatusInactive:
         // ...
     }
     ```

  3. **For 6+ states/branches, or if growth is expected:**
     Define an enum-like type using `iota` and use a `map` to associate each value with its handler or value.
     This improves scalability and maintainability.

     ```go
     type Status int

     const (
         StatusActive Status = iota
         StatusInactive
         StatusPending
         // ...
     )

     var statusHandlers = map[Status]func(){
         StatusActive:   handleActive,
         StatusInactive: handleInactive,
         StatusPending:  handlePending,
         // ...
     }

     if handler, ok := statusHandlers[status]; ok {
         handler()
     }
     ```

- Even with a smaller number of states, if significant future growth is likely, consider the enum-plus-map pattern from the beginning.
- Choose the approach that maximizes maintainability, clarity, and minimizes risk of errors as the project evolves.

---

## ⚙️ Core API Design

- **Three-Layer Construction Pattern:**
  - Core client: `wnc.New(controller, token, ...options)`
  - Domain services: `service.NewService(client)`
  - Typed methods: `service.Method(ctx)`

- **Configuration:**
  Accept configuration via `wnc.Config` struct, never read from environment/files directly.

- **Context Usage:**
  All functions that perform I/O or network operations must take `context.Context` as their first argument.

- **Error Handling:**
  - Never call `panic` or `log.Fatal`.
  - Always return errors to the caller—wrap with `fmt.Errorf("...: %w", err)`.
  - Define custom error types for actionable API errors.
  - **Standardized Error Patterns:** All client nil validation must use: `fmt.Errorf("%w: client cannot be nil", wnc.ErrInvalidConfiguration)`
  - **Consistency:** Avoid basic `errors.New()` for nil client checks; use wrapped errors for better error handling.

---

## 🧪 Testing Practices

- **Table-Driven Tests:**
  Use table-driven tests with `testing.T` and `t.Run()` for clarity and maintainability.

- **Subtest Isolation:**
  Use `t.Run()` for structured and named subtests.

- **SA1012 Compliance:**
  For nil context tests, use: `var nilCtx context.Context` instead of `nil` literals.
  Add `//nolint:SA1012` when necessary for intentional nil context testing.

- **Test Utilities:**
  Factor out common setup, environment checks, and data helpers.

- **Validate Marshaling:**
  Ensure Go structs can be marshaled/unmarshaled to JSON in accordance with the REST API, covering edge cases.

- **Comprehensive Error Checks:**
  Fail early and clearly report unrecoverable errors.

- **Context Use in Tests:**
  Ensure contexts are properly used and propagated in tests.

- **Collect and Persist Test Data:**
  Structure and save results (e.g., as JSON) for later review.

- **Pre-Run Validation:**
  Validate testing environment and inputs before running tests.

- **High Test Coverage:**
  Target 98% or higher for main codebase packages. Maintain ≥92% total project coverage.

- **Standardized Error Testing:**
  Expect consistent error format: `"invalid client configuration: client cannot be nil"`

- **Mock Server Testing:**
  Use full RESTCONF paths: `/restconf/data/[YANG-MODULE]:[CONTAINER]/[ENDPOINT]`

- **Comprehensive Path Coverage:**
  Test both success and error paths. Add dedicated HTTP error tests.

- **Service Architecture Testing:**
  Test service constructors: `NewService(client)` patterns
  Validate typed method signatures: `Method(ctx) (*model.Type, error)`
  Ensure proper client.Do() usage in service implementations

---

## 🏗️ Service Implementation Pattern

All domain services follow this standardized pattern:

```go
// Service provides [domain] operations using the WNC client
type Service struct {
    c *wnc.Client
}

// NewService creates a new [domain] service instance
func NewService(c *wnc.Client) *Service {
    return &Service{c: c}
}

// Method returns [domain] operational data
func (s *Service) Method(ctx context.Context) (*model.ResponseType, error) {
    const endpoint = "Cisco-IOS-XE-wireless-[module]-[type]:[container]"

    var result model.ResponseType
    err := s.c.Do(ctx, http.MethodGet, endpoint, &result)
    if err != nil {
        return nil, err
    }

    return &result, nil
}
```

**Service Requirements:**
- All services accept `*wnc.Client` in constructor
- Methods use `s.c.Do()` for HTTP operations
- Return typed structs from `internal/model/`
- Include descriptive documentation
- Follow consistent naming: `NewService`, domain-specific method names

**Client Integration:**
Services integrate with core client through accessor methods that return `nil` (placeholder pattern):
```go
// Domain service accessor (returns nil as placeholder)
func (c *Client) ServiceName() ServiceInterface {
    return nil // Placeholder for future service integration
}
```

---
