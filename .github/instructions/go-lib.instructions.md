---
description: Go Library Development Instructions
applyTo: "**/*.go"
---

# 🧪 GitHub Copilot Agent – Go Library Instructions

> Applies to all `*.go` files. General rules in `general.instructions.md` take precedence.

Primary Objective: Maintain and evolve a **pure Go SDK (library only)** for Cisco 9800 RESTCONF. No embedded CLI, no external deps.

---

## 🧭 Architecture (3 Layers)

| Layer | Location | Responsibility |
|-------|----------|----------------|
| Core | `wnc/` | Client construction, HTTP execution (`Do`), config, errors |
| Domain Services | `ap/`, `radio/`, `general/`, etc. | Thin typed methods calling core client |
| Generated Models | `internal/model/` | YANG-derived structs (do not hand-edit) |

Support packages under `internal/` (e.g. `httpx`, `restconf`, `validation`) are private helpers—keep boundaries clean.

---

## 🧩 Service Pattern

Each domain package exports:

```go
type Service struct { c *wnc.Client }
func NewService(c *wnc.Client) *Service { return &Service{c: c} }
```

Methods:

```go
func (s *Service) GetFoo(ctx context.Context) (*model.Foo, error) {
  const ep = "Cisco-IOS-XE-wireless-<module>:<container>"
  var out model.Foo
  if err := s.c.Do(ctx, http.MethodGet, ep, &out); err != nil { return nil, err }
  return &out, nil
}
```

Rules:

- First param performing I/O: `ctx context.Context`.
- Return pointer to model + `error` (no bools, no unused second values).
- No side effects beyond REST call + decoding.

Accessor placeholders on `*wnc.Client` may temporarily return `nil` until implemented; do not expose unstable experimental APIs.

---

## 🧱 Design Principles

- Minimal public surface; export only what users need.
- No global mutable state; configuration explicit via `wnc.Config` & options.
- Keep functions ≤ ~40 lines (hard cap 50 unless test data assembly). Split early.
- Prefer clarity over cleverness; reject premature generalization.
- Accept small interfaces, return concrete structs.

---

## 🔀 Branch / State Handling

Use simplest construct that fits:

| Case Count | Approach |
|------------|----------|
| 1 | Single `const` |
| 2–5 | `switch` / `if` chain |
| 6+ or growing | `iota` enum + map dispatch |

Avoid over-abstraction; add enum map only when scaling complexity.

---

## ⚙️ Errors

- Never `panic`, `log.Fatal`, or swallow errors.
- Wrap: `fmt.Errorf("context: %w", err)`.
- Nil client check uses: `fmt.Errorf("%w: client cannot be nil", wnc.ErrInvalidConfiguration)` (standardized message path).
- Avoid bare `errors.New` except for static sentinel definitions (unexported unless part of API contract).

---

## 🧪 Testing

| Aspect | Requirement |
|--------|-------------|
| Style | Table-driven subtests (`t.Run`) |
| Context NIL tests | `var nilCtx context.Context` + `//nolint:SA1012` if intentional |
| Coverage | Keep / improve existing; target ≥98% core pkgs, ≥92% overall |
| Paths | Success + error (HTTP errors, decode failures, nil client) |
| Marshal | Round-trip JSON where decoding logic branches |

Other rules:

- No external network in unit tests (use mock server / fixtures).
- Keep fixtures small & explicit (place under `test_data/`).
- Consistent expected error text (see error section).

---

## 🧰 Implementation Hygiene

- Import order: std lib first, blank line, internal packages.
- No unused / TODO code left uncommented—use `// TODO(username): reason` if needed.
- Prefer early returns to reduce nesting.
- Use `const` for RESTCONF path fragments reused ≥2 times.
- Validate inputs (nil client, empty identifiers) fast; return typed / wrapped error.

---

## � RESTCONF Endpoint Construction

Pattern:
`/restconf/data/<YANG-MODULE>:<top-container>/<sub-containers?...>`
Store module/container strings as `const` when reused. Do not concatenate with unvalidated user input.

---

## 🧾 Model Usage

- Only use types from `internal/model/` for wire payloads.
- Do not embed large model structs anonymously in public exported types (avoid API lock‑in).
- Add lightweight adapter structs in service layer if a narrower public shape is needed (subject to user approval before export).

---

## 🧱 Dependency Rules

- Standard library only. If a feature would require third-party dependency, propose first; do not add.

---

## 🔍 Review Checklist (Per Go Change)

1. Public API unchanged or clearly justified
2. Errors wrapped, no lost context
3. Function sizes within limits
4. Tests updated / added; coverage not reduced
5. No dead code, no extraneous comments
6. Paths & constants named clearly

---

## 🛡️ Anti-Patterns to Reject

- Generic over-engineering for single use
- Large switch statements with duplicated bodies (extract helper)
- Silent fallthrough without comment
- Exporting types solely for tests
- Re-exporting internal/model types unnecessarily

---

## 🧪 Minimal Example (Pattern Reference)

```go
package ap

import (
  "context"
  "net/http"
  "fmt"
  "github.com/your/module/wnc"
  "github.com/your/module/internal/model"
)

const apEndpoint = "Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data"

type Service struct { c *wnc.Client }
func NewService(c *wnc.Client) *Service { return &Service{c: c} }

func (s *Service) List(ctx context.Context) (*model.AccessPointOperData, error) {
  if s == nil || s.c == nil { return nil, fmt.Errorf("%w: client cannot be nil", wnc.ErrInvalidConfiguration) }
  var out model.AccessPointOperData
  if err := s.c.Do(ctx, http.MethodGet, apEndpoint, &out); err != nil { return nil, err }
  return &out, nil
}
```

---

## ♻️ Maintenance

Periodically prune obsolete services / placeholders. Keep this doc short; move deep rationale to separate design notes if needed.

---

End of Go Library Instructions.
