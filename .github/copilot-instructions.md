# GitHub Copilot Agent Mode – Development Instructions

**Updated:** 2025-09-14
\*\*Precedence:## 8. Review Scope & Comment Style

- Focus on the **diff** and align feedback with these rules; reserve wide refactors for items labeled `allow-wide`.
- Tag comments with **\[BLOCKER] / \[MAJOR] / \[MINOR (Nit)] / \[QUESTION] / \[PRAISE]**.
- Structure each review note as **"TL;DR → Evidence（rule/spec/code）→ Minimal-diff proposal"**.

## 9. Quality Gate & SemVerest in this repository

**Goal:** Modern Go 1.25+ SDK library for Cisco Catalyst C9800 RESTCONF (IOS-XE 17.12), zero-violation quality bar

## 0. Normative Keywords

Interpret **MUST / SHOULD / MAY** per RFC 2119/8174.

---

## 0.1. Core Development Principles

- **CDP-001 (MUST)** **Prioritize library consistency above all else.** All implementations, naming conventions, error handling patterns, and API designs must maintain absolute consistency throughout the codebase.
- **CDP-002 (MUST)** **Prioritize human readability, maintainability, and comprehensibility above all else.** Code must be self-documenting and optimized for long-term maintenance by development teams.
- **CDP-003 (MUST)** **Conduct deep impact analysis for all changes and implement fundamental solutions.** Surface-level fixes are prohibited; address root causes and consider downstream effects.
- **CDP-004 (MUST)** **Base all work on verified facts and concrete evidence. Speculation and assumptions are strictly prohibited.** Validate implementation details, API behavior, and system constraints before proceeding.
- **CDP-005 (MUST)** **Ask clarifying questions immediately when uncertainties arise. Independent assumptions and continued work without confirmation are strictly prohibited.** Halt progress and seek explicit guidance.
- **CDP-006 (MUST)** **For time-constrained work, document progress comprehensively and provide clear handoff instructions.** Include current state, next steps, and unresolved issues for session continuity.
- **CDP-007 (MUST)** **Create .bak backup files before editing any existing files in the codebase.** This applies to all file types including Go source files (.go), shell scripts (.sh). Preserve original state for rollback capability and change tracking.

## 1. Repository Purpose & Scope

- **RPS-001 (MUST)** Treat this repository as an **importable Go client SDK** for the **Cisco Catalyst C9800 RESTCONF API (IOS-XE 17.12)**.
- **RPS-002 (MUST)** Design for library use: clean APIs, stable SemVer surface, and integration by external Go programs.

## 2. Precedence & Applicability

- **PRA-001 (MUST)** Copilot follows **this file as the single source of truth** for edits and reviews.
- **RPA-002 (MUST)** Apply lint/format/type checks **as defined by repository configuration**（e.g. `.golangci.yml`, `.editorconfig`, `.vscode`, `.markdownlint*`, `.shellcheckrc`）.
- **RPS-003 (MUST)** When rules require changes, **propose a minimal configuration PR** instead of local overrides.

## 3. Expert Personas

- **EXP-001 (MUST)** Act as a **Cisco Catalyst 9800 Wireless Network Controller IOS-XE 17.12 to 17.18 expert**.
- **EXP-002 (MUST)** Act as a **Cisco Catalyst 9800 Wireless Network Controller's RESTCONF and the YANG model expert**.
- **EXP-003 (MUST)** Act as a **Go 1.25+ expert**.
- **EXP-004 (MUST)** Act as a **modern shell script expert**.

## 4. Security & Privacy

- **SEC-001 (MUST)** **Mask tokens and credentials in all outputs and logs**（e.g., `${TOKEN:0:6}...`）; log only non-sensitive metadata.
- **SEC-002 (MUST)** Keep authentication material **ephemeral in memory** and **scoped to requests**.

## 5. Tooling & Execution Workflow

- **TEW-001 (MUST)** Use repository Make targets:

  - `make build` until it passes.
  - `make test-unit` until it is green.
  - `make test-integration` until it is green（skip safely when live env is unavailable）.

- **TEW-002 (MUST)** **Create .bak backups before modifying any existing files in the codebase.** Preserve original state for rollback capability and change tracking.
- **TEW-003 (MUST)** After editing Go source files (.go), shell scripts (.sh) files in the codebase, execute relevant validation steps and ensure impacted Make targets succeed.
- **TEW-004 (MUST)** Limit terminal redirection operations（e.g., `echo ... >> file`）to **≤ 20 lines** per action.
- **TEW-005 (MUST)** On completion, write a summary to `./.copilot_reports/<YYYY-MM-DD_HH-mm-ss>_<prompt_title>.md`.

## 6. Workspace Hygiene

- **WSH-001 (MUST)** Place all temporary artifacts（work files, coverage, binaries）**under `./tmp/`**.
- **WSH-002 (MUST)** Keep `.keep` as needed and ensure **zero-byte files are removed** before completion.

## 7. Development Standards

- **DEV-001 (MUST)** When editing go code, follow **Go public coding standards** as defined in [./instructions/go.instructions.md](./instructions/go.instructions.md) primary.
- **DEV-002 (MUST)** When editing go code, follow **Go private coding standards** as defined in [./instructions/go-lib-umatare5.instructions.md](./instructions/go-lib-umatare5.instructions.md) secondary.
- **DEV-003 (MUST)** When editing bash shell scripts, follow **Bash shell scripting private standards** as defined in [./instructions/bash-umatare5.instructions.md](./instructions/bash-umatare5.instructions.md).
- **DEV-004 (MUST)** Apply lint/format/type checks as defined by repository configuration.
- **DEV-005 (MUST)** Maintain **zero-violation quality bar** across all development artifacts.

## 8. Review Scope & Comment Style

- **REV-001 (MUST)** Focus on the **diff** and align feedback with these rules; reserve wide refactors for items labeled `allow-wide`.
- **REV-002 (MUST)** Tag comments with **\[BLOCKER] / \[MAJOR] / \[MINOR (Nit)] / \[QUESTION] / \[PRAISE]**.
- **REV-003 (MUST)** Structure each review note as **“TL;DR → Evidence（rule/spec/code）→ Minimal-diff proposal”**.

## 9. Quality Gate & SemVer

- **QGA-001 (MUST)** Keep CI green across lint, build, unit/integration tests.
- **QGA-002 (MUST)** For **v1.0.0+**, align all API changes with **SemVer** and update docs accordingly.
- **QGA-003 (MUST)** Before completion, verify:

  - `../README.md` baseline requirements updated when relevant.
  - `../CONTRIBUTING.md` quality requirements satisfied.
  - `../docs/SECURITY.md` security notes up to date.
  - `../docs/TESTING.md` test instructions accurate.
  - Lint/format clean per repository settings（no ad-hoc overrides）.
  - Temp artifacts under `./tmp/` only, zero-byte files removed.
  - Report written to `./.copilot_reports/` as specified.
