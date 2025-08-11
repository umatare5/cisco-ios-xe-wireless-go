---
description: "Go Library Development Instructions"
applyTo: "**/*.go,**/go.mod,**/go.sum"
---

# Go **Library** Development Instructions

## Scope & Metadata

- **Last Updated**: 2025-08-10
- **Precedence**: 1. `copilot-instructions.md` (Global) → 2. `go.instructions.md` (Community) → 3. `go-lib-umatare5.instructions.md` (This)
- **Compatibility**: Depends on Go specifications (cross‑platform).
- **Style Base**: [Effective Go](https://go.dev/doc/effective_go) / [Go Code Review Comments](https://go.dev/wiki/CodeReviewComments)
- **Goal**: Service‑accessor‑centric stable API, Functional Options, context‑first networking, idiomatic Go, and error wrapping.

## 1. Architecture

- PK-001 (**MUST**) Expose **service accessors** from a **single root client** (e.g., `Client.AP()`, `Client.Clients()`, `Client.Rogue()`), and concentrate related operations in each service.
- PK-002 (**MUST**) Use **Functional Options** in primary constructors: `NewClient(controller string, token string, opts ...Option)`.
- PK-003 (**MUST**) Keep **functionality‑based packages** (e.g., `ap/`, `client/`, `wlan/`); place common wire types at the root as appropriate; put private helpers in `internal/`.
- PK-004 (**SHOULD**) Use **thin service structs** with small methods (e.g., `Oper(ctx)`, `Config(ctx, ...)`) that call shared HTTP helpers.
- PK-005 (**MUST NOT**) Do not pull cross‑service domain logic into a service. Share via `internal` utilities only.

## 2. Coding Style

- **STY-001 (MUST) Prioritize human maintainability, especially readability.**
- STY-003 (**SHOULD**) Functions are roughly **20–40 lines**; refactor those exceeding **50 lines** unless well justified.
- STY-004 (**SHOULD**) Split functions for readability and avoid excessive nesting.

## 3. Dependencies & Injection

- DI-001 (**MUST**) Inject external dependencies via **Options** (e.g., `WithHTTPClient(*http.Client)`, `WithLogger(*slog.Logger)`, `WithTimeout(...)`, `WithInsecureSkipVerify(...)`).
- DI-002 (**MUST NOT**) Do not use singletons/mutable globals. Hold state in `Client` or service receivers.
- DI-003 (**SHOULD**) Design small interfaces (e.g., Clock/RetryPolicy) that can be accepted via Options.
- DI-004 (**MUST**) Prefer the standard library first. Justify any new dependency per global policy.

## 4. Public API

- API-001 (**MUST**) Keep the public surface **minimal and stable**, built around service accessors and focused methods.
- API-002 (**MUST**, **v1.0.0+**) Follow **SemVer**; provide migration notes for breaking changes.
- API-003 (**SHOULD**) Place **executable examples** under `examples/` that compile with the current API.

## 5. Context / Concurrency

- CTX-001 (**MUST**) Pass `context.Context` through **all network operations**, respecting deadlines/cancellation.

## 6. Errors

- ERR-002 (**MUST**) Wrap with `%w` and include **actionable context** (operation/endpoint/identifier, etc.).

## 7. Logging

- LOG-001 (**SHOULD**) Use **structured logging** (e.g., `slog`) **only if injected** (via `WithLogger`). **MUST NOT** create global loggers inside the library.

## 8. Configuration

- CFG-001 (**MUST**) Accept configuration via **constructor args + Functional Options**; do **not** read env/files inside the library.
- CFG-002 (**SHOULD**) Validate options on construction and fail fast on invalid configs.
- CFG-003 (**SHOULD**) Provide safe defaults for timeouts/headers/TLS.

## 9. Serialization / Wire Compatibility

- SER-001 (**MUST**) Maintain JSON/YANG wire compatibility; keep struct tags correct (e.g., `json:"field_name,omitempty"`).

## 10. HTTP Client / Transport Hygiene

- NET-001 (**MUST**) Reuse a single `http.Client` on the root client and **close `resp.Body` on all paths**.
- NET-002 (**MUST**) Set explicit timeouts and honor `context` deadlines.
- NET-003 (**SHOULD**) Provide **bounded** retries/backoff (configurable via Options). Do not retry indefinitely.
- NET-004 (**MUST**) Set required headers and never log credentials in errors/logs.

## 11. Testing

- TS-001 (**MUST**) Use **table‑driven tests** and `t.Run` subtests.
- TS-002 (**MUST**) Cover success/failure paths; avoid time‑fragile tests.
- TS-003 (**SHOULD**) Use fakes/mocks via small interfaces; avoid heavyweight frameworks.
- TS-004 (**SHOULD**) Guard integration tests via build tags/environment; skip gracefully when unset.
- TS-005 (**MUST**) Focus on critical paths; target **\~95%** coverage as a guideline.
- TS-006 (**SHOULD**) Ensure examples under `examples/` compile in CI.
- TS-007 (**SHOULD**) Centralize fixtures/helpers under `internal/testutil` to reduce duplication.
