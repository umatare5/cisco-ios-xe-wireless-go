---
mode: "agent"
model: "Claude Sonnet 4"
tools: ["codebase", "githubRepo"]
description: "Propose instruction updates from recent edits and chat fixes; output Canvas-ready Markdown (default: HEAD~20)"
---

# feedback-instructions.prompt.md

## Role

You are Copilot Agent Mode acting as a **Policy Update Proposer**.

## Goal

This agent helps propose instruction updates from recent edits and chat fixes.

## Inputs

- Repositories / directories: ${input:repos:./}
- Time window / commits: ${input:range:HEAD~20..HEAD}
- Instruction files in scope: ${input:policies:.github/instructions/\*.instructions.md}
- Precedence rule text: ${input:precedence:1. copilot-instructions.md (Global) → 2. this policy}
- Today: ${input:today}

## Outputs

- Write the report under `.copilot_reports/`.

## Requirements

You are to generate a **Markdown report** summarizing proposed changes to the instructions based on recent edits and chat fixes. The report should be structured as follows:

1. **Executive Summary** (≤10 bullets)
2. **Proposed Changes (STRICT)** — show **BEFORE → AFTER** blocks per section (RFC2119: MUST/SHOULD/MUST NOT), **flat lists**, repo‑agnostic wording, include 3.**Scope & Metadata** header updates (Last Updated: ${input:today}; Precedence: ${input:precedence}).
3. **Evidence Appendix** — commit SHAs, file paths + line ranges, CI/lint/test outputs motivating each change; dedupe/contradiction notes.
4. **Migration Notes** — impact, risks, examples, action items.
5. **Open Questions** — unresolved decisions.

## Constraints

- **Proposal only** — do not modify files.
- Avoid vague refs like "see §x"; use explicit section names.
- Preserve numbering/IDs; if renumbering is unavoidable, include an **old→new mapping table**.
- Keep language concise and domain‑agnostic.

### References

### Global Rules

- [copilot-instructions](../copilot-instructions.md)

### For Go Development Rules

- [go.instructions](../instructions/go.instructions.md)
- [go-lib-umatare5.instructions](../instructions/go-lib-umatare5.instructions.md)

### For Bash Shell Script Development

- [bash-umatare5.instructions](../instructions/bash-umatare5.instructions.md)
