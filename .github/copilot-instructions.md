# GitHub Copilot Agent Mode – General Instructions (Global)

## Scope & Metadata

- **Last Updated**: 2025-08-10
- **Precedence**: Highest in this repository
- **Goals**:

  - **Primary Goal**: Contribute **only** to the SDK/library in this repo.
  - Keep public interfaces **clear, small, and stable**; maximize **readability/maintainability**.
  - Prefer **minimal diffs** (avoid unnecessary churn).
  - Defaults are secure (no secret logging; TLS is safe by default).

- **Non‑Goals**:

  - Creating a new `main` package or standalone apps.
  - Introducing editor‑external, ad‑hoc lint rules or style changes.
  - Emitting or persisting secrets/test credentials.

## 0. Normative Keywords

- NK-001 (**MUST**) Interpret **MUST / MUST NOT / SHOULD / SHOULD NOT / MAY** per RFC 2119/8174.

## 1. Repository Purpose & Scope

- GP-001 (**MUST**) Treat this repository as a **Go client SDK** for the **RESTCONF API** of **Cisco Catalyst C9800 Wireless LAN Controller (IOS‑XE 17.12)**.
- GP-002 (**MUST NOT**) Assume standalone application usage. This library is intended to be **imported** by other Go programs.

## 2. Precedence & Applicability

- GA-001 (**MUST**) When editing/generating code in this repository, Copilot **must follow** this document.
- GA-002 (**MUST**) In this repository, **this file (`copilot-instructions.md`) has the highest precedence** over any other instruction set. **On conflict, always prioritize this file**.
- GA-003 (**MUST**) Lint/format rules follow **editor/workspace settings only** (see §5).
- GA-004 (**MUST**) There is **no separate review instruction**. Review behavior is defined by this file as well.

## 3. Expert Personas (for AI edits/reviews)

- EP-001 (**MUST**) Act as a **Go 1.24 expert**.
- EP-002 (**MUST**) Act as a **domain expert for Cisco Catalyst 9800 / IOS‑XE 17.12** (controller behavior & RESTCONF basics).
- EP-003 (**SHOULD**) Be familiar with **YANG models** and behave accordingly (e.g., YangModels vendor/cisco/xe/17121) for IOS‑XE 17.12 contexts.

## 4. Security & Privacy

- SP-001 (**MUST NOT**) Log tokens or credentials. **MUST** mask secrets (e.g., `${TOKEN:0:6}...`).
- SP-002 (**MUST**) TLS is secure by default. **MUST NOT** default to insecure modes (allow explicitly for dev only).

## 5. Editor‑Driven Tooling (single source of truth)

- ED-001 (**MUST**) Lint/format/type checks follow repository settings (e.g., `.golangci.yml`, `.editorconfig`, `.vscode`, `.markdownlint.json`, `.markdownlintignore`, `.shellcheckrc`).
- ED-002 (**MUST NOT**) Do not add flags/rules or inline disables that are not configured.
- ED-003 (**SHOULD**) When reality conflicts with rules, propose a **minimal settings PR** instead of local overrides.

## 6. Coding Principles (Basics)

- GC-001 (**MUST**) Apply **KISS/DRY** and keep code quality high.
- GC-002 (**MUST**) Avoid magic numbers; **use named constants** proactively.
- GC-003 (**MUST**) Use **predicate helpers** (e.g., `is_*` / `has_*`) to improve readability.

## 7. Coding Principles (Conditionals)

- CF-001 (**MUST**) Prefer predicate helpers in conditions.
- CF-002 (**MUST**) Prefer **early returns** inside branches to keep logic simple and fast.

## 8. Coding Principles (Loops)

- LP-001 (**MUST**) In loops, prefer **early exits** (`return` / `break` / `continue`) to avoid deep nesting and keep logic simple and fast.

## 9. Working Directory / Temp Files

- WD-001 (**MUST**) Place all temporary artifacts (work files, coverage, test binaries, etc.) **under `./tmp`**.
- WD-002 (**MUST**) Before completion, delete **zero‑byte files** (**exception**: keep `.keep`).

## 10. Model‑Aware Execution Workflow (when shell execution is available)

- WF-001 (**MUST**) Before actions: **always launch and use `bash`** (no shell detection/adaptation).
- WF-002 (**MUST**) After editing Go code: run `make build` and fix until it passes.
- WF-003 (**MUST**) After editing Go code: run `make test-unit` and fix until green.
- WF-004 (**MUST**) After editing Go code: run `make test-integration` and fix until green (skip safely if env not set).
- WF-005 (**MUST**) After editing **shell scripts**: execute the target scripts under `./scripts/` with their documented options and fix until they succeed. Also ensure **all impacted `make` targets** run successfully.
- WF-006 (**MUST**) On completion: summarize actions/results into `./.copilot_reports/<prompt_title>_<YYYY-MM-DD_HH-mm-ss>.md`.

## 11. Tests / Quality Gate (for AI reviewers)

- QG-001 (**MUST**) Keep CI green. Do not merge code that violates configured lint/tests. **AI reviewer** must check CI/Lint/Test status on PRs and mark issues as **\[BLOCKER]** when unmet (see §12).

## 12. Change Scope & Tone (for AI reviewers)

- CS-001 (**MUST**) Focus on the **diff**; propose wide refactors only with explicit request/label (e.g., `allow-wide`).
- CS-002 (**SHOULD**) Tag comments with **\[BLOCKER] / \[MAJOR] / \[MINOR (Nit)] / \[QUESTION] / \[PRAISE]**.
- CS-003 (**SHOULD**) Structure comments as “TL;DR → Evidence (rule/proof) → Minimal‑diff proposal”.

## 13. Quick Checklist (before completion)

- QC-001 (**MUST**, **v1.0.0+**) API changes comply with SemVer and docs are updated.
- QC-002 (**MUST**) Follow **`./README.md`** for baseline requirements.
- QC-003 (**MUST**) Follow **`./CONTRIBUTING.md`** for quality requirements.
- QC-004 (**MUST**) Follow **`./docs/SECURITY.md`** for security requirements.
- QC-005 (**MUST**) Run required **Make targets** per `./CONTRIBUTING.md`.
- QC-006 (**MUST**) Lint/format are clean per editor settings (no ad‑hoc flags/inline disables).
- QC-007 (**MUST**) Temp artifacts under `./tmp`, zero‑byte files removed, and report written to `./.copilot_reports/`.
