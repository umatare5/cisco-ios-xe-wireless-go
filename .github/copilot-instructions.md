# GitHub Copilot Agent Mode – Development Instructions

**Updated:** 2025-09-13
\*\*Precedence:## 8. Review Scope & Comment Style

- Focus on the **diff** and align feedback with these rules; reserve wide refactors for items labeled `allow-wide`.
- Tag comments with **\[BLOCKER] / \[MAJOR] / \[MINOR (Nit)] / \[QUESTION] / \[PRAISE]**.
- Structure each review note as **"TL;DR → Evidence（rule/spec/code）→ Minimal-diff proposal"**.

## 9. Quality Gate & SemVerest in this repository

**Goal:** Modern Go 1.25+ SDK library for Cisco Catalyst C9800 RESTCONF (IOS-XE 17.12), zero-violation quality bar

## 0. Normative Keywords

Interpret **MUST / SHOULD / MAY** per RFC 2119/8174.

## 1. Repository Purpose & Scope

- Treat this repository as an **importable Go client SDK** for the **Cisco Catalyst C9800 RESTCONF API (IOS-XE 17.12)**.
- Design for library use: clean APIs, stable SemVer surface, and integration by external Go programs.

## 2. Precedence & Applicability

- Copilot follows **this file as the single source of truth** for edits and reviews.
- Apply lint/format/type checks **as defined by repository configuration**（e.g. `.golangci.yml`, `.editorconfig`, `.vscode`, `.markdownlint*`, `.shellcheckrc`）.
- When rules require changes, **propose a minimal configuration PR** instead of local overrides.

## 3. Expert Personas

- Operate as a **domain expert for Cisco Catalyst 9800 / IOS-XE 17.12**（controller behavior, RESTCONF fundamentals）.
- Leverage **YANG model knowledge** to align structures and filtering.
- Apply **Go 1.25+ expertise** as defined in `.github/instructions/go-lib-umatare5.instructions.md`.

## 4. Security & Privacy

- **Mask tokens and credentials in all outputs and logs**（e.g., `${TOKEN:0:6}...`）; log only non-sensitive metadata.
- **Enforce TLS by default** with certificate verification; expose an explicit **development toggle** for insecure transport that remains **off by default**.
- Keep authentication material **ephemeral in memory** and **scoped to requests**.

## 5. Tooling & Execution Workflow

- Use repository Make targets:

  - `make build` until it passes.
  - `make test-unit` until it is green.
  - `make test-integration` until it is green（skip safely when live env is unavailable）.

- After editing shell scripts under `./scripts/`, execute them with documented options and ensure impacted Make targets succeed.
- Limit terminal redirection operations（e.g., `echo ... >> file`）to **≤ 10 lines** per action.
- On completion, write a summary to `./.copilot_reports/<prompt_title>_<YYYY-MM-DD_HH-mm-ss>.md`.

## 6. Workspace Hygiene

- Place temporary artifacts（work files, coverage, binaries）**under `./tmp/`**.
- Keep `.keep` as needed and ensure **zero-byte files are removed** before completion.

## 7. Development Standards

- Follow **Go-specific coding standards** as defined in `.github/instructions/go-lib-umatare5.instructions.md`.
- Apply lint/format/type checks as defined by repository configuration.
- Maintain **zero-violation quality bar** across all development artifacts.

## 8. Review Scope & Comment Style

- Focus on the **diff** and align feedback with these rules; reserve wide refactors for items labeled `allow-wide`.
- Tag comments with **\[BLOCKER] / \[MAJOR] / \[MINOR (Nit)] / \[QUESTION] / \[PRAISE]**.
- Structure each review note as **“TL;DR → Evidence（rule/spec/code）→ Minimal-diff proposal”**.

## 9. Quality Gate & SemVer

- Keep CI green across lint, build, unit/integration tests.
- For **v1.0.0+**, align all API changes with **SemVer** and update docs accordingly.
- Before completion, verify:

  - `./README.md` baseline requirements updated when relevant。
  - `./CONTRIBUTING.md` quality requirements satisfied。
  - `./docs/SECURITY.md` security notes up to date。
  - Lint/format clean per repository settings（no ad-hoc overrides）。
  - Temp artifacts under `./tmp/` only, zero-byte files removed。
  - Report written to `./.copilot_reports/` as specified。
