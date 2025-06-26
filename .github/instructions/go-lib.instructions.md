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

- **Organize by Package-per-Resource:**
  Group related features into distinct packages/directories (e.g., `/client`, `/api`, `/models`).
  Place the main `Client` and `NewClient(...)` in the root package.
  Isolate private/internal logic in the `internal/` directory.
  Minimize dependencies between packages to ensure clear responsibilities and avoid coupling.

- **Strict Dependency Injection:**
  Inject dependencies (such as `*http.Client`) via constructors like `NewClient`.
  **Do not** use global state or singletons for shared state/configuration.
  Always pass dependencies explicitly through struct fields, function arguments, or constructors.
  This enhances testability, composability, and long-term maintainability.

- **Clean API Design:**
  Export only intended public types and functions (start with uppercase), keeping the API minimal and stable.
  Define small, focused interfaces (Interface Segregation, Dependency Inversion).

  - Example:

    ```go
    type Storage interface {
        Save(ctx context.Context, obj MyObj) error
    }
    ```

---

## 🛠️ Go Coding Practices & Style

- **Follow Idiomatic Go:**
  Conform to [Effective Go](https://go.dev/doc/effective_go) and [Go Code Review Comments](https://go.dev/wiki/CodeReviewComments).

- **Style & Linting:**
  Format all code with `gofmt` and ensure it passes `golangci-lint`.

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

- **Construction:**
  Use `NewClient(config Config)` as the entry point to instantiate the client.

- **Configuration:**
  Do not read configuration from environment variables or files. Always accept a configuration struct.

- **Context Usage:**
  All functions that perform I/O or network operations must take `context.Context` as their first argument.

- **Error Handling:**
  - Never call `panic` or `log.Fatal`.
  - Always return errors to the caller—wrap with `fmt.Errorf("...: %w", err)`.
  - Define custom error types for actionable API errors.

---

## 🧪 Testing Practices

- **Table-Driven Tests:**
  Use table-driven tests with `testing.T` and `t.Run()` for clarity and maintainability.

- **Subtest Isolation:**
  Use `t.Run()` for structured and named subtests.

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
  Target 80% or higher code coverage for all packages.

---
