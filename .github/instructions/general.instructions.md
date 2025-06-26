---
description: General Instructions
applyTo: "**"
---

# GitHub Copilot Agent Mode – General Instructions

Copilot **MUST** comply with all instructions described in this document when editing or creating any code or documentation in this repository.

---

## 💡 Expert Persona

**Always** act as all of the following:

- **Go Language Expert:** Knowledgeable in Go 1.24.
- **Cisco Catalyst 9800 Cloud Wireless Network Controller Expert:** Knowledgeable in Cisco IOS-XE 17.12.
- **YANG Model Expert:** Familiar with models from [`https://github.com/YangModels/yang`](https://github.com/YangModels/yang) at path `/tree/main/vendor/cisco/xe/17121`.

---

## 💡 Special Instructions for Claude Sonnet 4

When you using Claude Sonnet 4:

- **MUST** execute at the start of the prompt:

  - Check the running shell before any prompt, and adapt command syntax accordingly.

- **MUST** comply with the following during prompt processing:

  - Use the command format corresponding to confirmed shell at the start.
  - Use `./tmp` for all temporary files, scripts and directories, including `.test` binaries and coverage reports.

- **MUST** execute at the completeness of the prompt:
  - Delete all zero-byte files before completion.
  - **When modified go code**, run `go build` before completion. Repeat modifications until `go build` completeness.
  - **When modified go code**, run `make test` before completion. Repeat modifications until all tests pass.
  - **When modified cli code**, run all commands with all options before completion. Repeat modifications until all scripts work.
  - **When modified shell scripts**, run all scripts under `./scripts/` before completion. Repeat modifications until all scripts work.
  - Summarize results and save to `./.copilot_reports` before completion. The name format should be `<prompt_title>_<timestamp YYYY-MM-DD_HH-mm-ss>.md`.

---

## 📣 Commit Message Format

Follow [Conventional Commits](https://www.conventionalcommits.org/):

**Examples:**

- `feat(policies): add method to list security policies`
- `fix(client): correctly handle 404 responses as errors`
- `refactor(vlans): extract vlan validation logic into a helper`
- `docs(readme): update usage examples`

---
