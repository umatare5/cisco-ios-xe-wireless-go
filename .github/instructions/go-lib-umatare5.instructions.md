---
description: "Go Library Development Instructions"
applyTo: "**/*.go,**/go.mod,**/go.sum"
---

# Go **Library** Development Instructions

GitHub Copilot **MUST** follow these instructions when generating or modifying Go code in this repository.

## Scope & Metadata

- **Last Updated**: 2025-08-23
- **Precedence**: 1. `copilot-instructions.md` (Global) → 2. `go.instructions.md` (Community) → 3. `go-lib-umatare5.instructions.md` (This)
- **Compatibility**: Depends on Go specifications (cross‑platform).
- **Style Base**: [Effective Go](https://go.dev/doc/effective_go) / [Go Code Review Comments](https://go.dev/wiki/CodeReviewComments)
- **Goal**: Service‑accessor‑centric stable API, Functional Options, context‑first networking, idiomatic Go, and error wrapping.

## 1. Architecture

- **PK-001 (MUST)** Expose **service accessors** from a **single root client** (e.g., `Client.AP()`, `Client.WLAN()`). All services live under `service/`.
- **PK-002 (MUST)** Use **Functional Options** in primary constructors: `NewClient(controller, token string, opts ...Option)`.
- **PK-003 (MUST)** Keep **functionality‑based packages** under `service/` (e.g., `service/ap`, `service/wlan`); put shared code under `internal/`.
- **PK-004 (SHOULD)** Use **thin service receivers** that delegate transport to shared HTTP helpers.
- **PK-005 (MUST NOT)** Pull cross‑service domain logic into a single service. Share via `internal/*` only.
- **PK-006 (MUST)** Resolve **cross‑service data dependencies internally** with private helpers; keep the public API decoupled.
- **PK-007 (SHOULD)** Provide **type‑safe integration interfaces** (normalize params internally; attach contextual errors).
- **PK-008 (MAY)** Offer **direct service constructors** (e.g., `NewService(*core.Client)`) primarily for tests/advanced use.

## 2. Coding Style

- **STY-001 (MUST)** Prefer readability and maintainability.
- **STY-002 (SHOULD)** Functions \~20–40 lines; refactor if >50 without cause.
- **STY-003 (SHOULD)** Split functions to avoid deep nesting; use early returns.

## 3. Dependencies & Injection

- **DI-001 (MUST)** Inject externals via **Options** (`WithHTTPClient`, `WithLogger`, `WithTimeout`, `WithInsecureSkipVerify`).
- **DI-002 (MUST NOT)** No singletons/mutable globals. Keep state on `Client`/receivers.
- **DI-003 (SHOULD)** Small interfaces for clocks/retries.
- **DI-004 (MUST)** Prefer stdlib first; justify any new dep.

## 4. Public API

- **API-001 (MUST)** Keep the public surface **minimal and stable**.
- **API-002 (MUST, v1+)** Follow **SemVer**; provide migration notes for breaks.
- **API-003 (SHOULD)** Place **executable examples** under `examples/`.
- **API-004 (SHOULD)** For filtered reads prefer `Get*By<Filter>` (e.g., `...ByEthernetMAC`, `...ByWlanID`, `...ByLocation`). **Exception**: if a single primary identifier (e.g., AP MAC) is canonical, you **may omit** `By<Primary>` (e.g., `GetApTags(ctx, apMAC)`). Be **internally consistent**.
- **API-005 (MAY)** For field‑projections use `*Only` (e.g., `GetGlobalOperApHistoryOnly`). Document projection semantics.
- **API-006 (SHOULD)** For write/admin ops use **imperative verbs** (`Enable*`, `Disable*`, `Assign*`, `Set*`, `Clear*`) with param order `(ctx, primaryIdentifier, qualifiers...)`.
- **API-007 (SHOULD)** Config root patterns: `GetCfg` (full), `Get<Subset>Cfg` (subresource), `GetCfg<Subset>Only` (field projection of root). Keep these distinct.
- **API-008 (SHOULD)** Read operation params order `(ctx, primaryIdentifier, qualifiers...)`.
- **API-009 (MAY)** Use `*All` suffix for unfiltered full‑dataset variants (e.g., `GetOperCapwapDataAll`).
- **API-010 (SHOULD)** Tag assignment API: `AssignSiteTag`, `AssignPolicyTag`, `AssignRFTag` → `(ctx, apMAC, tag)` plus internal multi‑tag orchestrator.
- **API-011 (SHOULD)** Prefer **domain‑typed qualifiers** (e.g., `core.RadioBand`) over raw ints for public APIs. _Legacy slot‑based tests are tolerated but SHOULD migrate._

## 5. Context / Concurrency

- **CTX-001 (MUST)** Thread `context.Context` through all network ops; honor deadlines/cancellation.

## 6. Errors

- **ERR-001 (MUST)** When propagating underlying errors, wrap with `%w` and include actionable context (operation/endpoint/identifier).
- **ERR-002 (MAY)** For local validation/sentinel errors, return typed/sentinel errors without wrapping; document and reuse in tests.
- **ERR-003 (MUST)** If using `fmt.Errorf` with an underlying error, the format string **must contain `%w`**; otherwise don’t pass the underlying error.

## 7. Logging

- **LOG-001 (SHOULD)** Use structured logging only if injected (`WithLogger`). **Never** create globals inside the library.

## 8. Configuration

- **CFG-001 (MUST)** Accept config via constructor args + Options; **do not** read env/files from inside the library.
- **CFG-002 (SHOULD)** Validate options on construction; fail fast.
- **CFG-003 (SHOULD)** Provide safe defaults for timeouts/headers/TLS.

## 9. Serialization / Wire Compatibility

- **SER-001 (MUST)** Keep JSON/YANG tags accurate (e.g., `json:"field_name,omitempty"`).
- **SER-002 (MUST)** Evolve models **gradually**; prefer additive fields; compose specializations where needed.
- **SER-003 (SHOULD)** Align JSON tags with **YANG** names and hierarchy; apply `omitempty` per API.
- **SER-004 (SHOULD)** For GETs, return ergonomic inner models when outer envelopes are transport/meta only; document deviations and be consistent within a family.

## 10. HTTP Client / Transport Hygiene

- **NET-001 (MUST)** Reuse one `http.Client`; **close `resp.Body`** on all paths.
- **NET-002 (MUST)** Set explicit timeouts; honor context.
- **NET-003 (SHOULD)** Use bounded retries/backoff; never infinite.
- **NET-004 (MUST)** Set required headers; never log credentials.
- **NET-005 (SHOULD)** Use shared generic helpers (`core.Get[T]`, `core.Post[T]`, `core.Put[T]`, `core.PutVoid`, `core.Delete[T]`, `core.DeleteVoid`).
- **NET-006 (SHOULD)** Use the client’s **RPC helper** for YANG RPCs (e.g., `s.client.PostRPC(ctx, <RPCConstant>, payload)`). **Do not** POST RPCs via data endpoints.

## 11. Testing

### 11.A Core principles

- **TS-001 (MUST)** Use **table‑driven tests** and `t.Run` subtests.
- **TS-002 (MUST)** Cover success/failure paths; avoid time‑fragile tests.
- **TS-003 (SHOULD)** Use fakes/mocks via small interfaces; avoid heavy frameworks.
- **TS-004 (SHOULD)** Guard integration tests via build tags/env/`testing.Short()`; skip gracefully. **Set/admin ops run sequentially**; read‑only GETs may run in parallel.
- **TS-005 (MUST)** Focus on critical paths; target \~95% coverage (guideline).
- **TS-006 (SHOULD)** Ensure `examples/` compile in CI.
- **TS-007 (SHOULD)** Centralize fixtures/helpers under `internal/testutil`.
- **TS-008 (MUST)** Adopt the **4‑pattern** structure per service test file: (1) Unit/Struct+JSON, (2) Table‑driven Method, (3) Fail‑fast, (4) Integration (guarded). Use standardized runners from `internal/testutil`.
- **TS-009 (SHOULD)** Use a **DataCollector** with `sync.Mutex` for concurrent verification.
- **TS-010 (SHOULD)** Optimize concurrency via `sync.WaitGroup` with bounded parallelism for CI stability.
- **TS-011 (MUST)** Make error collection **non‑blocking**; distinguish expected API errors (e.g., 404) from fatal ones.
- **TS-012 (SHOULD)** Prefer shared parallel runners: `RunParallel*Validation`, `RunParallelIntegrationValidation`, `RunSetOperationStructureTests`, etc.
- **TS-013 (SHOULD)** Use generic test cases like `GenericFunctionValidationCase[T]`.
- **TS-014 (MAY)** For internal helpers, allow exact string equality assertions; for public API, prefer sentinels or substring matches.
- **TS-015 (SHOULD)** Name subtests with **domain/scenario** (`site/valid`, `policy/empty`).
- **TS-015a (MUST)** Test function names reflect **what** is tested (e.g., _StatusCode/Integration/Validation_), not **how** (e.g., _TableDriven_).
- **TS-015b (MUST)** Live integration test names use `test-{service}-{operation}-{yyyymmdd}-{hhmmss}` (≤31 chars) for IOS‑XE compatibility.
- **TS-015c (SHOULD)** Keep naming consistent across services.
- **TS-016 (SHOULD)** Validate **nil client**, **canceled context**, and **nil context** paths.
- **TS-017 (SHOULD)** Use `SetupOptionalClient` / `SetupRequiredClient` helpers; skip if unmet.
- **TS-018 (SHOULD)** For **set/admin ops**, run integration sequentially and log safety notes; treat non‑existent resources as expected errors where applicable.
- **TS-019 (MAY)** Delineate test files with numbered headers (e.g., `// 1-1`, `// 1-2`).
- **TS-020 (SHOULD)** Use clear local constants or `testutil` (priorities, IDs, MACs, radio bands, locations).
- **TS-021 (MAY)** In struct validation, comparing **expected type names** as strings is acceptable via helpers.
- **TS-022 (SHOULD)** Prefer sample factory helpers (`CreateSample*`) from `internal/testutil`.
- **TS-023 (MAY)** Use underscore‑segmented function names to group by area (e.g., `Test_ApGetCfg_Methods`).
- **TS-024 (SHOULD)** In GET integrations, group tests by **Basic/Filtered/FieldsOnly**; parallel OK.
- **TS-025 (SHOULD)** For `By<Filter>` methods, include invalid/edge filter cases and assert non‑fatal handling.
- **TS-026 (SHOULD)** For set/admin methods with numeric/enum qualifiers (slots/priorities/**RadioBand**), include invalid/out‑of‑range cases.
- **TS-027 (SHOULD)** Tag assignment tests must cover invalid MAC format, empty/invalid tags, and “at least one tag required”; verify normalization before transport.
- **TS-028 (SHOULD)** Use centralized skip messages (`SkipIntegrationShortMode`, etc.).
- **TS-029 (SHOULD)** Use centralized timeouts (`testutil.TestOperationTimeout`).
- **TS-030 (MAY)** Use standardized expectation strings from `testutil`.
- **TS-031 (SHOULD)** **Disambiguate MAC types**: use `TestAPMac` (Radio/WTP) for admin/oper; use `TestAPEthernetMac` for cfg when required.
- **TS-032 (MAY)** Provide BSSID samples in dotted/colon forms if accepted; normalize or document.
- **TS-033 (SHOULD)** Prefer **domain‑typed** test constants (e.g., `core.RadioBand`); migrate legacy slot tests.

### 11.B Test naming & organization (additions)

- **Files**: name as `[target_file]_test.go` (e.g., `service.go` → `service_test.go`).
- **Service test functions**: `Test_[ServiceName]Service_[TestType]` (e.g., `TestClientService_Integration`, `TestSiteService_Integration`).
- **CRUD test functions**: `Test_[ServiceName]Service_[Operation]_[TestType]` (e.g., `TestAPService_GetOper_Methods`, `TestRFTagService_CreateRFTag_StatusCode`).

### 11.C Service (`service.go`) test flow

- **FailFast** first: verify early returns for abnormal paths (**NilClient / CtxCancel**).
- **Integration**: run CRUD against a real controller (Create → Get → Set → Delete → List).

### 11.D Standard CRUD test focus (set|get|run)\_\*.go

- **Structural**: struct models are correct.
- **URLConstruction**: URL builders produce expected paths.
- **Methods**: functions behave as expected (happy & error paths).

### 11.E Tag services (`tag_service.go`) — site/rf/wlan shared

- **FailFast** first (NilClient / CtxCancel), then **Integration** CRUD (Create → Get → Set → Delete → List).

### 11.F Tag CRUD files `tag_(set|get|list|create|delete).go`

- **Structural** / **URLConstruction** / **Methods** checks as above.

## 12. Test Consolidation Framework

- **TC-001 (MUST)** Use standardized suites from `internal/testutil` (`CompleteServiceTestSuite`, `MinimalServiceTestSuite`, `SimpleConfigurationTestSuite`, `GetOperTestSuite`).
- **TC-002 (MUST)** Choose suite by service capability; see repo patterns.
- **TC-003 (MUST)** Organize per service: `service_test.go`, `get_cfg_test.go`, `get_oper_test.go`, plus operation‑specific tests.
- **TC-004 (SHOULD)** Apply unified pattern directories (patterns/validation/execution) provided by `internal/testutil`.
- **TC-005 (MUST)** Classify expected vs fatal errors consistently (e.g., 404 for unsupported endpoints).
- **TC-006 (SHOULD)** Integration grouping and execution model as in §11.A.
- **TC-007 (SHOULD)** Centralize test data and factories.
- **TC-008 (MUST)** Use JSON/struct/method/context validation helpers per framework.

## 13. Function Design & Receiver Pattern

- **FN-001 (MUST)** Use **service receiver methods** when HTTP I/O or `s.client` is needed, the method is exported, or logic is stateful/complex.
- **FN-002 (MUST)** Use **pure functions** for parameter‑only, side‑effect‑free logic (validation/transforms/builders).
- **FN-003 (MUST)** **File placement**:

  - Service receiver methods → `<feature>.go` under `service/<domain>`.
  - Pure helpers → **centralized** in `internal/helpers/<domain>.go`.
  - Builders → **prefer** `internal/builders/<domain>.go`.
  - **Compatibility note (updated)**: Where current code keeps builder functions **inside the service package** (e.g., `buildAPCfgApTagData` used by tests), this is **allowed** if the builder is pure, covered by tests, and has no transport logic. New code SHOULD prefer centralized builders.

- **FN-004 (SHOULD)** Decision priority: need for HTTP → need for `s.client` → public API → statefulness → testability.
- **FN-005 (SHOULD)** Validation/builders that depend on external services → receiver methods; otherwise keep pure.
- **FN-006 (MUST)** Fail fast on **nil client**; cover with tests.
- **FN-007 (SHOULD)** Return promptly on **canceled contexts**.
- **FN-008 (SHOULD)** Early range/format validation for numeric/index qualifiers.
- **FN-009 (SHOULD)** Normalize primary identifiers (e.g., AP MAC) **before** composing URLs.
- **FN-010 (SHOULD)** Internal orchestrators may accept pointer qualifiers to signal optionality; validate nil clearly.

## 14. Package Structure & Namespace Management

- **PKG-001 (MUST)** Organize by domain under `service/*`; shared packages under `internal/*`:

  - `internal/restconf` — builders & protocol; `internal/restconf/routes` — endpoint constants.
  - `internal/core` — transport helpers, generic HTTP, RPC facade, URL builders.
  - `internal/helpers` — centralized, pure helpers (ap, site, wlan, mdns, dot11, rfid, rogue, tag, ...).
  - `internal/builders` — domain builders (e.g., `ap.go`).
  - `internal/validation` — format/range validators, normalizers, runtime defaults.
  - `internal/errors` — cross‑domain error strings; service‑specific errors may live under `service/<domain>/errors.go`.
  - `internal/model/*` — wire models with docs.
  - `internal/testutil` — unified test framework.

- **PKG-002 (MUST)** If package names collide with stdlib, use **import aliases** and document boundaries.
- **PKG-003 (SHOULD)** Split broad namespaces by domain.
- **PKG-004 (MUST NOT)** Avoid deep trees; keep `internal/` depth ≤ 2–3.
- **PKG-005 (SHOULD)** Intent‑driven test filenames: `service_test.go`, `get_cfg_test.go`, `get_oper_test.go`, `set_*_test.go`, `tag_*.go`, `*_helper_test.go`.
- **PKG-006 (SHOULD)** Make file scopes obvious; don’t mix unrelated responsibilities.

## 15. Design Principles in Practice

- **DP-001 (MUST)** Apply **SRP** strictly.
- **DP-002 (SHOULD)** Leverage **DI** to maximize testability; prefer pure functions when mocks aren’t needed.
- **DP-003 (SHOULD)** Prefer **early returns**.
- **DP-004 (SHOULD)** Use predicate helpers `is*/has*` for readability.

## 16. Template Constants & Indirection

- **TMP-001 (SHOULD)** Remove “template constants” that add indirection; prefer direct `fmt.Sprintf` for one‑offs.
- **TMP-002 (MUST)** Convert all usages before deletion; update imports and keep formatting consistent.

## 17. Function Ordering & File Organization

- **ORD-001 (MUST)** Order within files: (1) **Public API** → (2) **Core internal** → (3) **Validation** → (4) **Builders/transformers** → (5) **Error helpers**.
- **ORD-002 (SHOULD)** Keep related functions adjacent; align names & params.
- **ORD-003 (MUST)** Apply consistently across files.

## 18. Constants & Separation of Concerns

- **CONST-001 (MUST)** Separate by domain: errors (`errors.go`), payload/wire, endpoints (`internal/restconf/routes`), defaults (`internal/validation`).
- **CONST-002 (SHOULD)** Group by **usage pattern**, not alphabet.
- **CONST-003 (MUST)** Treat `internal/validation` as the **source of truth for runtime defaults**. If `internal/helpers/tag` exposes tag defaults, keep them as **thin aliases only**. Tests should prefer `internal/testutil` for constructing data; assert defaulting behavior against `internal/validation`.

## 19. Endpoint & RPC URL Construction

- **END-001 (SHOULD)** Define `<Feature>BasePath` using `restconf.YANGModelPrefix`; compose leaf endpoints from it.
- **END-002 (MUST)** Build field‑projection URLs via centralized helpers (e.g., `buildCfgFieldsURL` / generic `buildFieldsURL`).
- **END-003 (SHOULD)** For path/query construction with parameter encoding use `core.BuildQueryURL(base, id)`; **normalize identifiers first**.
- **RPC-001 (SHOULD)** Name RPC constants with action suffixes (e.g., `SetAPSlotAdminStateRPC`); colocate with endpoints; RPC payloads end with `RPCPayload`/`RPCInput`.

## 20. Helper Function Patterns & Naming

- **HELP-001 (SHOULD)** Intent‑revealing names: `*All`, `Get*By<Filter>`, primary‑implicit filtered reads, `*Only` for fields; builders `build*URL`; validation `is*/has*`; lookup `find*By*`.
- **HELP-002 (MUST)** Centralized helpers in `internal/helpers/<domain>.go` are **pure** and **stateless**.

## 21. Code Quality Metrics & Validation

- **QM-001 (SHOULD)** Track DRY, cohesion, cycle elimination, testability.
- **QM-002 (MUST)** Validate via compile, tests, lint, and updated docs.
- **QM-003 (MUST)** Detect and consolidate function duplication.
- **QM-004 (SHOULD)** Group helpers by role (builders/filters/finders/predicates).

## 22. Builder Semantics & Defaults

- **BLD-001 (MUST)** Builders are **pure**, deterministically fill missing optional fields using centralized defaults.
- **BLD-002 (MUST)** Test **all‑set**, **partial‑empty**, and **all‑empty** cases (e.g., `buildAPCfgApTagData`).
- **BLD-003 (SHOULD)** Keep builders small/composable; never embed transport logic.

## 23. Package Documentation (GoDoc)

- **DOC-001 (MUST)** Each top‑level package has a package comment with **Overview, Main Features, Usage Example, Known Limitations, Error Handling, Requirements**.
- **DOC-002 (MUST)** List only **implemented** API in **Main Features**.
- **DOC-003 (MUST)** Document unsupported/removed endpoints under **Known Limitations**, including typical **404** behavior.
- **DOC-004 (MUST)** Error Handling section must note that **404 may be expected** for unimplemented endpoints and how tests treat them (non‑fatal vs fatal).
- **DOC-005 (MUST)** Requirements: min controller/version (e.g., IOS‑XE **17.12+**), RESTCONF enabled, auth.
- **DOC-006 (SHOULD)** Minimal example using root client and service accessor with `context.Context`.
- **DOC-007 (SHOULD)** Safety notes for admin/set operations.
- **DOC-008 (MUST)** Use canonical qualifier names (e.g., `EthernetMAC`, `WlanID`, `Location`) consistently; when historical/idiomatic names exist (e.g., `MAC`, `WtpMAC`), keep the package internally consistent and document mapping.
- **DOC-009 (MUST)** Doc method names must match exported code; deprecations noted explicitly.
- **DOC-010 (SHOULD)** Keep functionality lists synchronized with code.
- **DOC-011 (SHOULD)** Examples compile in CI.
- **DOC-012 (MAY)** Include short error‑handling snippets (404/cancellation).
- **DOC-013 (SHOULD)** Document relationships between `*All` and filtered variants.
- **DOC-014 (SHOULD)** For radio admin, document **`core.RadioBand`** usage and mapping.
- **DOC-015 (SHOULD)** Keep Known Limitations aligned with integration results and controller versions.
- **DOC-016 (MUST)** **Do not** contradict exported methods in Known Limitations; explain version‑specific behavior instead.

## 24. Endpoint Development Best Practices (New Endpoint Guide)

- **EDP-001 (SHOULD)** Define clear **Goal** and **Deliverables** (`service.go`, `service_test.go`, `doc.go`) and **Success criteria** (unit/integration passing, verified on real WNC, lint clean).
- **EDP-002 (MUST)** Provide **technical specs upfront**: module names, target RESTCONF paths, resource types (cfg/oper/global‑oper), key filters, RPC vs data classification.
- **EDP-003 (SHOULD)** Phased plan: Foundation → Basic Methods → Filters → RPC → Tests & Docs.
- **EDP-004 (MUST)** Distinguish **data** (`/restconf/data/`) vs **RPC** (`/restconf/operations/`); use proper plumbing (`PostRPC`, `BuildRPCURL`).
- **EDP-005 (MUST)** **Identifier selection (AP domain)**: admin RPCs use **Radio MAC** (`ap-mac`); config queries use the YANG key; always validate/normalize.
- **EDP-006 (SHOULD)** Keep a **verification checklist** (build/tests/lint + real environment checks, MAC type correctness).
- **EDP-007 (MAY)** Provide **sanitized live commands** using env vars (no secrets).
- **EDP-008 (SHOULD)** Capture minimal notes tying filters/RPC inputs to YANG definitions.
- **EDP-009 (MUST)** For filters, prefer RESTCONF list syntax `/list=key` via helpers; never inline unescaped input.
- **EDP-010 (MUST)** Integration‑test RPC operations against a live controller; expect **HTTP 204 No Content** on success; log safety notes.
- **EDP-011 (SHOULD)** Keep `doc.go` limitations in sync with observed behavior.

### Endpoint Instruction Template

_(Omitted here for brevity; use the template block from the Best Practices section when creating a new endpoint service.)_

## 26. Tag Management Standardization (TagSetOper)

- **TAG-001 (MUST)** Use **TagSetOper** terminology; avoid temporal labels.
- **TAG-002 (MUST)** Use **TagName** as the identifier field in all tag structs (no `ResourceID`/`Id`).
- **TAG-003 (MUST)** File naming under `service/{domain}`: `tag_set.go`, `tag_create.go`, `tag_delete.go`, `tag_get.go`, `tag_list.go` (and corresponding `*_test.go`).
- **TAG-004 (MUST)** Standardized CRUD method set for each tag type.
- **TAG-005 (MUST NOT)** Keep deprecated `Update*Tag` methods; remove and document as **BREAKING CHANGE**.
- **TAG-006 (MUST)** Enforce YANG constraints: tag name ≤ 32 chars; test format `test-{tag-type}-tag-yyyymmdd-hhmmss`.
- **TAG-007 (MUST)** Use `internal/testutil` tag testing framework patterns.
- **TAG-008 (MUST)** Domain validations: non‑nil client/config; non‑empty `TagName`; domain‑specific rules.
- **TAG-009 (SHOULD)** Capture domain‑specific considerations (site/wlan/rf) in validations.
- **TAG-010 (MUST)** Document breaking changes in commits.
- **TAG-011 (SHOULD)** Centralized, contextual error handling for tag ops.
- **TAG-012 (MUST)** Live WNC integration testing with env validation.
- **TAG-013 (MUST)** Use **UnifiedTagIntegrationTestSuite** where applicable.
- **TAG-014 (SHOULD)** Comprehensive **status code** validation for all tag ops.
- **UTF-011/012 (MUST)** Naming/Docs for live tests; document counts, lint status, and formatting improvements in Recent Changes when relevant.

---

### Notes on resolved conflicts

- **Builders location**: While centralized builders under `internal/builders` are preferred, the repository currently uses certain **package‑local builders** (e.g., `buildAPCfgApTagData`) referenced directly in tests. This instruction **permits** local builders when **pure and test‑covered**, and recommends migration over time.
- **Radio operations API**: Public APIs and new tests SHOULD use **`core.RadioBand`**. Legacy slot‑based tests are acceptable short‑term; prefer migration.
- **Defaults**: Runtime defaults live in `internal/validation`. Any `internal/helpers/tag` defaults must be thin aliases. Tests should construct data using `internal/testutil` constants and assert behavior against `internal/validation` defaults.

---

description: "Go Library Development Instructions"
applyTo: "**/\*.go,**/go.mod,\*\*/go.sum"

---

# Go **Library** Development Instructions

## Scope & Metadata

- **Last Updated**: 2025-08-23
- **Precedence**: 1. `copilot-instructions.md` (Global) → 2. `go.instructions.md` (Community) → 3. `go-lib-umatare5.instructions.md` (This)
- **Compatibility**: Depends on Go specifications (cross‑platform).
- **Style Base**: [Effective Go](https://go.dev/doc/effective_go) / [Go Code Review Comments](https://go.dev/wiki/CodeReviewComments)
- **Goal**: Service‑accessor‑centric stable API, Functional Options, context‑first networking, idiomatic Go, and error wrapping.

## Recent Changes

- **2025-08-23**: Test naming standardization and CRUD testing guidance added (this file). Harmonized with current repo test patterns and your latest code samples.
- **2025-08-23**: Tag management standardization and TagSetOper pattern (status‑code centric tests; semantic test names; unified tag integration suite).
- **2025-08-22**: Helper consolidation (`internal/helpers/*`) and centralized routes (`internal/restconf/routes`). RPC plumbing standardized via `client.PostRPC` and `BuildRPCURL`. Directory migrated to `service/*` layout.

> **Conflict policy**: When instructions conflict with your **latest provided code/tests**, **code wins**. Rules below are updated to reflect that.

---

## 1. Architecture

- **PK-001 (MUST)** Expose **service accessors** from a **single root client** (e.g., `Client.AP()`, `Client.WLAN()`). All services live under `service/`.
- **PK-002 (MUST)** Use **Functional Options** in primary constructors: `NewClient(controller, token string, opts ...Option)`.
- **PK-003 (MUST)** Keep **functionality‑based packages** under `service/` (e.g., `service/ap`, `service/wlan`); put shared code under `internal/`.
- **PK-004 (SHOULD)** Use **thin service receivers** that delegate transport to shared HTTP helpers.
- **PK-005 (MUST NOT)** Pull cross‑service domain logic into a single service. Share via `internal/*` only.
- **PK-006 (MUST)** Resolve **cross‑service data dependencies internally** with private helpers; keep the public API decoupled.
- **PK-007 (SHOULD)** Provide **type‑safe integration interfaces** (normalize params internally; attach contextual errors).
- **PK-008 (MAY)** Offer **direct service constructors** (e.g., `NewService(*core.Client)`) primarily for tests/advanced use.

## 2. Coding Style

- **STY-001 (MUST)** Prefer readability and maintainability.
- **STY-002 (SHOULD)** Functions \~20–40 lines; refactor if >50 without cause.
- **STY-003 (SHOULD)** Split functions to avoid deep nesting; use early returns.

## 3. Dependencies & Injection

- **DI-001 (MUST)** Inject externals via **Options** (`WithHTTPClient`, `WithLogger`, `WithTimeout`, `WithInsecureSkipVerify`).
- **DI-002 (MUST NOT)** No singletons/mutable globals. Keep state on `Client`/receivers.
- **DI-003 (SHOULD)** Small interfaces for clocks/retries.
- **DI-004 (MUST)** Prefer stdlib first; justify any new dep.

## 4. Public API

- **API-001 (MUST)** Keep the public surface **minimal and stable**.
- **API-002 (MUST, v1+)** Follow **SemVer**; provide migration notes for breaks.
- **API-003 (SHOULD)** Place **executable examples** under `examples/`.
- **API-004 (SHOULD)** For filtered reads prefer `Get*By<Filter>` (e.g., `...ByEthernetMAC`, `...ByWlanID`, `...ByLocation`). **Exception**: if a single primary identifier (e.g., AP MAC) is canonical, you **may omit** `By<Primary>` (e.g., `GetApTags(ctx, apMAC)`). Be **internally consistent**.
- **API-005 (MAY)** For field‑projections use `*Only` (e.g., `GetGlobalOperApHistoryOnly`). Document projection semantics.
- **API-006 (SHOULD)** For write/admin ops use **imperative verbs** (`Enable*`, `Disable*`, `Assign*`, `Set*`, `Clear*`) with param order `(ctx, primaryIdentifier, qualifiers...)`.
- **API-007 (SHOULD)** Config root patterns: `GetCfg` (full), `Get<Subset>Cfg` (subresource), `GetCfg<Subset>Only` (field projection of root). Keep these distinct.
- **API-008 (SHOULD)** Read operation params order `(ctx, primaryIdentifier, qualifiers...)`.
- **API-009 (MAY)** Use `*All` suffix for unfiltered full‑dataset variants (e.g., `GetOperCapwapDataAll`).
- **API-010 (SHOULD)** Tag assignment API: `AssignSiteTag`, `AssignPolicyTag`, `AssignRFTag` → `(ctx, apMAC, tag)` plus internal multi‑tag orchestrator.
- **API-011 (SHOULD)** Prefer **domain‑typed qualifiers** (e.g., `core.RadioBand`) over raw ints for public APIs. _Legacy slot‑based tests are tolerated but SHOULD migrate._

## 5. Context / Concurrency

- **CTX-001 (MUST)** Thread `context.Context` through all network ops; honor deadlines/cancellation.

## 6. Errors

- **ERR-001 (MUST)** When propagating underlying errors, wrap with `%w` and include actionable context (operation/endpoint/identifier).
- **ERR-002 (MAY)** For local validation/sentinel errors, return typed/sentinel errors without wrapping; document and reuse in tests.
- **ERR-003 (MUST)** If using `fmt.Errorf` with an underlying error, the format string **must contain `%w`**; otherwise don’t pass the underlying error.

## 7. Logging

- **LOG-001 (SHOULD)** Use structured logging only if injected (`WithLogger`). **Never** create globals inside the library.

## 8. Configuration

- **CFG-001 (MUST)** Accept config via constructor args + Options; **do not** read env/files from inside the library.
- **CFG-002 (SHOULD)** Validate options on construction; fail fast.
- **CFG-003 (SHOULD)** Provide safe defaults for timeouts/headers/TLS.

## 9. Serialization / Wire Compatibility

- **SER-001 (MUST)** Keep JSON/YANG tags accurate (e.g., `json:"field_name,omitempty"`).
- **SER-002 (MUST)** Evolve models **gradually**; prefer additive fields; compose specializations where needed.
- **SER-003 (SHOULD)** Align JSON tags with **YANG** names and hierarchy; apply `omitempty` per API.
- **SER-004 (SHOULD)** For GETs, return ergonomic inner models when outer envelopes are transport/meta only; document deviations and be consistent within a family.

## 10. HTTP Client / Transport Hygiene

- **NET-001 (MUST)** Reuse one `http.Client`; **close `resp.Body`** on all paths.
- **NET-002 (MUST)** Set explicit timeouts; honor context.
- **NET-003 (SHOULD)** Use bounded retries/backoff; never infinite.
- **NET-004 (MUST)** Set required headers; never log credentials.
- **NET-005 (SHOULD)** Use shared generic helpers (`core.Get[T]`, `core.Post[T]`, `core.Put[T]`, `core.PutVoid`, `core.Delete[T]`, `core.DeleteVoid`).
- **NET-006 (SHOULD)** Use the client’s **RPC helper** for YANG RPCs (e.g., `s.client.PostRPC(ctx, <RPCConstant>, payload)`). **Do not** POST RPCs via data endpoints.

## 11. Testing

### 11.A Core principles

- **TS-001 (MUST)** Use **table‑driven tests** and `t.Run` subtests.
- **TS-002 (MUST)** Cover success/failure paths; avoid time‑fragile tests.
- **TS-003 (SHOULD)** Use fakes/mocks via small interfaces; avoid heavy frameworks.
- **TS-004 (SHOULD)** Guard integration tests via build tags/env/`testing.Short()`; skip gracefully. **Set/admin ops run sequentially**; read‑only GETs may run in parallel.
- **TS-005 (MUST)** Focus on critical paths; target \~95% coverage (guideline).
- **TS-006 (SHOULD)** Ensure `examples/` compile in CI.
- **TS-007 (SHOULD)** Centralize fixtures/helpers under `internal/testutil`.
- **TS-008 (MUST)** Adopt the **4‑pattern** structure per service test file: (1) Unit/Struct+JSON, (2) Table‑driven Method, (3) Fail‑fast, (4) Integration (guarded). Use standardized runners from `internal/testutil`.
- **TS-009 (SHOULD)** Use a **DataCollector** with `sync.Mutex` for concurrent verification.
- **TS-010 (SHOULD)** Optimize concurrency via `sync.WaitGroup` with bounded parallelism for CI stability.
- **TS-011 (MUST)** Make error collection **non‑blocking**; distinguish expected API errors (e.g., 404) from fatal ones.
- **TS-012 (SHOULD)** Prefer shared parallel runners: `RunParallel*Validation`, `RunParallelIntegrationValidation`, `RunSetOperationStructureTests`, etc.
- **TS-013 (SHOULD)** Use generic test cases like `GenericFunctionValidationCase[T]`.
- **TS-014 (MAY)** For internal helpers, allow exact string equality assertions; for public API, prefer sentinels or substring matches.
- **TS-015 (SHOULD)** Name subtests with **domain/scenario** (`site/valid`, `policy/empty`).
- **TS-015a (MUST)** Test function names reflect **what** is tested (e.g., _StatusCode/Integration/Validation_), not **how** (e.g., _TableDriven_).
- **TS-015b (MUST)** Live integration test names use `test-{service}-{operation}-{yyyymmdd}-{hhmmss}` (≤31 chars) for IOS‑XE compatibility.
- **TS-015c (SHOULD)** Keep naming consistent across services.
- **TS-016 (SHOULD)** Validate **nil client**, **canceled context**, and **nil context** paths.
- **TS-017 (SHOULD)** Use `SetupOptionalClient` / `SetupRequiredClient` helpers; skip if unmet.
- **TS-018 (SHOULD)** For **set/admin ops**, run integration sequentially and log safety notes; treat non‑existent resources as expected errors where applicable.
- **TS-019 (MAY)** Delineate test files with numbered headers (e.g., `// 1-1`, `// 1-2`).
- **TS-020 (SHOULD)** Use clear local constants or `testutil` (priorities, IDs, MACs, radio bands, locations).
- **TS-021 (MAY)** In struct validation, comparing **expected type names** as strings is acceptable via helpers.
- **TS-022 (SHOULD)** Prefer sample factory helpers (`CreateSample*`) from `internal/testutil`.
- **TS-023 (MAY)** Use underscore‑segmented function names to group by area (e.g., `Test_ApGetCfg_Methods`).
- **TS-024 (SHOULD)** In GET integrations, group tests by **Basic/Filtered/FieldsOnly**; parallel OK.
- **TS-025 (SHOULD)** For `By<Filter>` methods, include invalid/edge filter cases and assert non‑fatal handling.
- **TS-026 (SHOULD)** For set/admin methods with numeric/enum qualifiers (slots/priorities/**RadioBand**), include invalid/out‑of‑range cases.
- **TS-027 (SHOULD)** Tag assignment tests must cover invalid MAC format, empty/invalid tags, and “at least one tag required”; verify normalization before transport.
- **TS-028 (SHOULD)** Use centralized skip messages (`SkipIntegrationShortMode`, etc.).
- **TS-029 (SHOULD)** Use centralized timeouts (`testutil.TestOperationTimeout`).
- **TS-030 (MAY)** Use standardized expectation strings from `testutil`.
- **TS-031 (SHOULD)** **Disambiguate MAC types**: use `TestAPMac` (Radio/WTP) for admin/oper; use `TestAPEthernetMac` for cfg when required.
- **TS-032 (MAY)** Provide BSSID samples in dotted/colon forms if accepted; normalize or document.
- **TS-033 (SHOULD)** Prefer **domain‑typed** test constants (e.g., `core.RadioBand`); migrate legacy slot tests.

### 11.B Test naming & organization (additions)

- **Files**: name as `[target_file]_test.go` (e.g., `service.go` → `service_test.go`).
- **Service test functions**: `Test_[ServiceName]Service_[TestType]` (e.g., `TestClientService_Integration`, `TestSiteService_Integration`).
- **CRUD test functions**: `Test_[ServiceName]Service_[Operation]_[TestType]` (e.g., `TestAPService_GetOper_Methods`, `TestRFTagService_CreateRFTag_StatusCode`).

### 11.C Service (`service.go`) test flow

- **FailFast** first: verify early returns for abnormal paths (**NilClient / CtxCancel**).
- **Integration**: run CRUD against a real controller (Create → Get → Set → Delete → List).

### 11.D Standard CRUD test focus (set|get|run)\_\*.go

- **Structural**: struct models are correct.
- **URLConstruction**: URL builders produce expected paths.
- **Methods**: functions behave as expected (happy & error paths).

### 11.E Tag services (`tag_service.go`) — site/rf/wlan shared

- **FailFast** first (NilClient / CtxCancel), then **Integration** CRUD (Create → Get → Set → Delete → List).

### 11.F Tag CRUD files `tag_(set|get|list|create|delete).go`

- **Structural** / **URLConstruction** / **Methods** checks as above.

## 12. Test Consolidation Framework

- **TC-001 (MUST)** Use standardized suites from `internal/testutil` (`CompleteServiceTestSuite`, `MinimalServiceTestSuite`, `SimpleConfigurationTestSuite`, `GetOperTestSuite`).
- **TC-002 (MUST)** Choose suite by service capability; see repo patterns.
- **TC-003 (MUST)** Organize per service: `service_test.go`, `get_cfg_test.go`, `get_oper_test.go`, plus operation‑specific tests.
- **TC-004 (SHOULD)** Apply unified pattern directories (patterns/validation/execution) provided by `internal/testutil`.
- **TC-005 (MUST)** Classify expected vs fatal errors consistently (e.g., 404 for unsupported endpoints).
- **TC-006 (SHOULD)** Integration grouping and execution model as in §11.A.
- **TC-007 (SHOULD)** Centralize test data and factories.
- **TC-008 (MUST)** Use JSON/struct/method/context validation helpers per framework.

## 13. Function Design & Receiver Pattern

- **FN-001 (MUST)** Use **service receiver methods** when HTTP I/O or `s.client` is needed, the method is exported, or logic is stateful/complex.
- **FN-002 (MUST)** Use **pure functions** for parameter‑only, side‑effect‑free logic (validation/transforms/builders).
- **FN-003 (MUST)** **File placement**:

  - Service receiver methods → `<feature>.go` under `service/<domain>`.
  - Pure helpers → **centralized** in `internal/helpers/<domain>.go`.
  - Builders → **prefer** `internal/builders/<domain>.go`.
  - **Compatibility note (updated)**: Where current code keeps builder functions **inside the service package** (e.g., `buildAPCfgApTagData` used by tests), this is **allowed** if the builder is pure, covered by tests, and has no transport logic. New code SHOULD prefer centralized builders.

- **FN-004 (SHOULD)** Decision priority: need for HTTP → need for `s.client` → public API → statefulness → testability.
- **FN-005 (SHOULD)** Validation/builders that depend on external services → receiver methods; otherwise keep pure.
- **FN-006 (MUST)** Fail fast on **nil client**; cover with tests.
- **FN-007 (SHOULD)** Return promptly on **canceled contexts**.
- **FN-008 (SHOULD)** Early range/format validation for numeric/index qualifiers.
- **FN-009 (SHOULD)** Normalize primary identifiers (e.g., AP MAC) **before** composing URLs.
- **FN-010 (SHOULD)** Internal orchestrators may accept pointer qualifiers to signal optionality; validate nil clearly.

## 14. Package Structure & Namespace Management

- **PKG-001 (MUST)** Organize by domain under `service/*`; shared packages under `internal/*`:

  - `internal/restconf` — builders & protocol; `internal/restconf/routes` — endpoint constants.
  - `internal/core` — transport helpers, generic HTTP, RPC facade, URL builders.
  - `internal/helpers` — centralized, pure helpers (ap, site, wlan, mdns, dot11, rfid, rogue, tag, ...).
  - `internal/builders` — domain builders (e.g., `ap.go`).
  - `internal/validation` — format/range validators, normalizers, runtime defaults.
  - `internal/errors` — cross‑domain error strings; service‑specific errors may live under `service/<domain>/errors.go`.
  - `internal/model/*` — wire models with docs.
  - `internal/testutil` — unified test framework.

- **PKG-002 (MUST)** If package names collide with stdlib, use **import aliases** and document boundaries.
- **PKG-003 (SHOULD)** Split broad namespaces by domain.
- **PKG-004 (MUST NOT)** Avoid deep trees; keep `internal/` depth ≤ 2–3.
- **PKG-005 (SHOULD)** Intent‑driven test filenames: `service_test.go`, `get_cfg_test.go`, `get_oper_test.go`, `set_*_test.go`, `tag_*.go`, `*_helper_test.go`.
- **PKG-006 (SHOULD)** Make file scopes obvious; don’t mix unrelated responsibilities.

## 15. Design Principles in Practice

- **DP-001 (MUST)** Apply **SRP** strictly.
- **DP-002 (SHOULD)** Leverage **DI** to maximize testability; prefer pure functions when mocks aren’t needed.
- **DP-003 (SHOULD)** Prefer **early returns**.
- **DP-004 (SHOULD)** Use predicate helpers `is*/has*` for readability.

## 16. Template Constants & Indirection

- **TMP-001 (SHOULD)** Remove “template constants” that add indirection; prefer direct `fmt.Sprintf` for one‑offs.
- **TMP-002 (MUST)** Convert all usages before deletion; update imports and keep formatting consistent.

## 17. Function Ordering & File Organization

- **ORD-001 (MUST)** Order within files: (1) **Public API** → (2) **Core internal** → (3) **Validation** → (4) **Builders/transformers** → (5) **Error helpers**.
- **ORD-002 (SHOULD)** Keep related functions adjacent; align names & params.
- **ORD-003 (MUST)** Apply consistently across files.

## 18. Constants & Separation of Concerns

- **CONST-001 (MUST)** Separate by domain: errors (`errors.go`), payload/wire, endpoints (`internal/restconf/routes`), defaults (`internal/validation`).
- **CONST-002 (SHOULD)** Group by **usage pattern**, not alphabet.
- **CONST-003 (MUST)** Treat `internal/validation` as the **source of truth for runtime defaults**. If `internal/helpers/tag` exposes tag defaults, keep them as **thin aliases only**. Tests should prefer `internal/testutil` for constructing data; assert defaulting behavior against `internal/validation`.

## 19. Endpoint & RPC URL Construction

- **END-001 (SHOULD)** Define `<Feature>BasePath` using `restconf.YANGModelPrefix`; compose leaf endpoints from it.
- **END-002 (MUST)** Build field‑projection URLs via centralized helpers (e.g., `buildCfgFieldsURL` / generic `buildFieldsURL`).
- **END-003 (SHOULD)** For path/query construction with parameter encoding use `core.BuildQueryURL(base, id)`; **normalize identifiers first**.
- **RPC-001 (SHOULD)** Name RPC constants with action suffixes (e.g., `SetAPSlotAdminStateRPC`); colocate with endpoints; RPC payloads end with `RPCPayload`/`RPCInput`.

## 20. Helper Function Patterns & Naming

- **HELP-001 (SHOULD)** Intent‑revealing names: `*All`, `Get*By<Filter>`, primary‑implicit filtered reads, `*Only` for fields; builders `build*URL`; validation `is*/has*`; lookup `find*By*`.
- **HELP-002 (MUST)** Centralized helpers in `internal/helpers/<domain>.go` are **pure** and **stateless**.

## 21. Code Quality Metrics & Validation

- **QM-001 (SHOULD)** Track DRY, cohesion, cycle elimination, testability.
- **QM-002 (MUST)** Validate via compile, tests, lint, and updated docs.
- **QM-003 (MUST)** Detect and consolidate function duplication.
- **QM-004 (SHOULD)** Group helpers by role (builders/filters/finders/predicates).

## 22. Builder Semantics & Defaults

- **BLD-001 (MUST)** Builders are **pure**, deterministically fill missing optional fields using centralized defaults.
- **BLD-002 (MUST)** Test **all‑set**, **partial‑empty**, and **all‑empty** cases (e.g., `buildAPCfgApTagData`).
- **BLD-003 (SHOULD)** Keep builders small/composable; never embed transport logic.

## 23. Package Documentation (GoDoc)

- **DOC-001 (MUST)** Each top‑level package has a package comment with **Overview, Main Features, Usage Example, Known Limitations, Error Handling, Requirements**.
- **DOC-002 (MUST)** List only **implemented** API in **Main Features**.
- **DOC-003 (MUST)** Document unsupported/removed endpoints under **Known Limitations**, including typical **404** behavior.
- **DOC-004 (MUST)** Error Handling section must note that **404 may be expected** for unimplemented endpoints and how tests treat them (non‑fatal vs fatal).
- **DOC-005 (MUST)** Requirements: min controller/version (e.g., IOS‑XE **17.12+**), RESTCONF enabled, auth.
- **DOC-006 (SHOULD)** Minimal example using root client and service accessor with `context.Context`.
- **DOC-007 (SHOULD)** Safety notes for admin/set operations.
- **DOC-008 (MUST)** Use canonical qualifier names (e.g., `EthernetMAC`, `WlanID`, `Location`) consistently; when historical/idiomatic names exist (e.g., `MAC`, `WtpMAC`), keep the package internally consistent and document mapping.
- **DOC-009 (MUST)** Doc method names must match exported code; deprecations noted explicitly.
- **DOC-010 (SHOULD)** Keep functionality lists synchronized with code.
- **DOC-011 (SHOULD)** Examples compile in CI.
- **DOC-012 (MAY)** Include short error‑handling snippets (404/cancellation).
- **DOC-013 (SHOULD)** Document relationships between `*All` and filtered variants.
- **DOC-014 (SHOULD)** For radio admin, document **`core.RadioBand`** usage and mapping.
- **DOC-015 (SHOULD)** Keep Known Limitations aligned with integration results and controller versions.
- **DOC-016 (MUST)** **Do not** contradict exported methods in Known Limitations; explain version‑specific behavior instead.

## 24. Endpoint Development Best Practices (New Endpoint Guide)

- **EDP-001 (SHOULD)** Define clear **Goal** and **Deliverables** (`service.go`, `service_test.go`, `doc.go`) and **Success criteria** (unit/integration passing, verified on real WNC, lint clean).
- **EDP-002 (MUST)** Provide **technical specs upfront**: module names, target RESTCONF paths, resource types (cfg/oper/global‑oper), key filters, RPC vs data classification.
- **EDP-003 (SHOULD)** Phased plan: Foundation → Basic Methods → Filters → RPC → Tests & Docs.
- **EDP-004 (MUST)** Distinguish **data** (`/restconf/data/`) vs **RPC** (`/restconf/operations/`); use proper plumbing (`PostRPC`, `BuildRPCURL`).
- **EDP-005 (MUST)** **Identifier selection (AP domain)**: admin RPCs use **Radio MAC** (`ap-mac`); config queries use the YANG key; always validate/normalize.
- **EDP-006 (SHOULD)** Keep a **verification checklist** (build/tests/lint + real environment checks, MAC type correctness).
- **EDP-007 (MAY)** Provide **sanitized live commands** using env vars (no secrets).
- **EDP-008 (SHOULD)** Capture minimal notes tying filters/RPC inputs to YANG definitions.
- **EDP-009 (MUST)** For filters, prefer RESTCONF list syntax `/list=key` via helpers; never inline unescaped input.
- **EDP-010 (MUST)** Integration‑test RPC operations against a live controller; expect **HTTP 204 No Content** on success; log safety notes.
- **EDP-011 (SHOULD)** Keep `doc.go` limitations in sync with observed behavior.

### Endpoint Instruction Template

_(Omitted here for brevity; use the template block from the Best Practices section when creating a new endpoint service.)_

## 26. Tag Management Standardization (TagSetOper)

- **TAG-001 (MUST)** Use **TagSetOper** terminology; avoid temporal labels.
- **TAG-002 (MUST)** Use **TagName** as the identifier field in all tag structs (no `ResourceID`/`Id`).
- **TAG-003 (MUST)** File naming under `service/{domain}`: `tag_set.go`, `tag_create.go`, `tag_delete.go`, `tag_get.go`, `tag_list.go` (and corresponding `*_test.go`).
- **TAG-004 (MUST)** Standardized CRUD method set for each tag type.
- **TAG-005 (MUST NOT)** Keep deprecated `Update*Tag` methods; remove and document as **BREAKING CHANGE**.
- **TAG-006 (MUST)** Enforce YANG constraints: tag name ≤ 32 chars; test format `test-{tag-type}-tag-yyyymmdd-hhmmss`.
- **TAG-007 (MUST)** Use `internal/testutil` tag testing framework patterns.
- **TAG-008 (MUST)** Domain validations: non‑nil client/config; non‑empty `TagName`; domain‑specific rules.
- **TAG-009 (SHOULD)** Capture domain‑specific considerations (site/wlan/rf) in validations.
- **TAG-010 (MUST)** Document breaking changes in commits.
- **TAG-011 (SHOULD)** Centralized, contextual error handling for tag ops.
- **TAG-012 (MUST)** Live WNC integration testing with env validation.
- **TAG-013 (MUST)** Use **UnifiedTagIntegrationTestSuite** where applicable.
- **TAG-014 (SHOULD)** Comprehensive **status code** validation for all tag ops.
- **UTF-011/012 (MUST)** Naming/Docs for live tests; document counts, lint status, and formatting improvements in Recent Changes when relevant.

---

### Notes on resolved conflicts

- **Builders location**: While centralized builders under `internal/builders` are preferred, the repository currently uses certain **package‑local builders** (e.g., `buildAPCfgApTagData`) referenced directly in tests. This instruction **permits** local builders when **pure and test‑covered**, and recommends migration over time.
- **Radio operations API**: Public APIs and new tests SHOULD use **`core.RadioBand`**. Legacy slot‑based tests are acceptable short‑term; prefer migration.
- **Defaults**: Runtime defaults live in `internal/validation`. Any `internal/helpers/tag` defaults must be thin aliases. Tests should construct data using `internal/testutil` constants and assert behavior against `internal/validation` defaults.
