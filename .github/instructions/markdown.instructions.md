---
description: Markdown Documentation Instructions
applyTo: "**/*.md"
---

# 📝 GitHub Copilot Agent – Markdown Instructions

> Applies to all `*.md`. If any rule conflicts with `general.instructions.md`, the general file wins.

Goal: Produce **clear, concise, lint‑clean** documentation that accelerates developer comprehension with minimal cognitive load.

Favor collapsible `<details>` for secondary / verbose samples (logs, long coverage tables). Primary task flows remain visible without expansion.

---

## ✨ Core Rules

| Aspect                | Rule                                                                  |
| --------------------- | --------------------------------------------------------------------- |
| Single Title          | Exactly one H1 at top; all other sections start H2.                   |
| Emojis                | Prefix H1/H2 (optional for H3) with a fitting emoji for scannability. |
| Lint                  | Must pass `markdownlint-cli2` (no inline disable unless justified).   |
| Whitespace            | Blank line before/after headings, lists, tables, code fences.         |
| Line Length           | Prefer ≤120 chars; wrap naturally (no forced hard breaks).            |
| Tone                  | Direct, active voice, minimal filler.                                 |
| Audience              | Go + network engineers; avoid marketing language.                     |
| File Names / Commands | Always in backticks.                                                  |

---

## 🔤 Formatting & Semantics

- Use `**bold**` for emphasis, never for decoration.
- Use backticks for identifiers, paths, env vars, commands.
- Use tables when presenting option matrices or comparisons (>2 columns).
- Prefer ordered lists for sequences; unordered for concepts.
- Keep paragraphs short (≤4 lines visually).

---

## 📦 Code & Examples

| Element       | Requirement                                                         |
| ------------- | ------------------------------------------------------------------- |
| Fenced Blocks | Always specify language: `go`, `bash`, `json`, `text`, `yaml`.      |
| One Concept   | Each block demonstrates exactly one focused idea.                   |
| Placeholders  | Use `<value>` or `<MODULE>`; document once if reused.               |
| Long Output   | Wrap inside `<details><summary>Show output</summary>...</details>`. |

Example:

```bash
# Explicit controller must be provided by user
# export WNC_CONTROLLER=<controller-hostname>
export WNC_ACCESS_TOKEN="<base64token>"
```

---

## 🚨 Alerts & Notes

Use GitHub alert syntax sparingly:

> [!NOTE]
> Clarifies nuances or cross-refs.
>
> [!WARNING]
> Risk, security, or destructive action.

Do not overuse. Prefer at most one alert per major section.

---

## 📄 Structure Template

Recommended outline (adapt as needed, keep lean):

1. Title / brief tagline
2. Purpose / scope
3. Quick start or primary action
4. Core concepts
5. Deeper usage / patterns
6. Reference (tables, env vars, exit codes)
7. Troubleshooting / FAQs (optional)
8. Appendix / extended samples (collapsible)

---

## 🔗 Links & Cross-References

- Use relative links for intra-repo docs.
- External links: plain Markdown; avoid bare URLs.
- Group related links at end of section if >2.

---

## 🧪 Quality Checklist (Per Edit)

- [ ] One H1 only
- [ ] Lint passes
- [ ] No stale references / renamed files
- [ ] Commands verified (where feasible) or marked as examples
- [ ] Consistent terminology (e.g. “controller”, not mixed with “WLC” unless defined)

---

## 🌐 RESTCONF / YANG Mentions

When referencing YANG:

| Element   | Style                                   |
| --------- | --------------------------------------- |
| Module    | `Cisco-IOS-XE-wireless-<domain>-<type>` |
| Path      | Backticks, no trailing slash            |
| Data Tree | Provide minimal path snippet only       |

Avoid speculative modules—only list what exists.

---

## 🧩 Reusability Patterns

- Use small snippet anchors (repeat by reference not duplication).
- Prefer a single canonical example for auth export.
- Collapse verbose sample outputs.

---

## 🔒 Security Hygiene

- Redact secrets: use `<token>` not real values.
- Warn (alert) before showing destructive or privileged operations.
- Never embed live credential values.

---

## ♻️ Maintenance Guidelines

| Action                   | Frequency                         | Notes                        |
| ------------------------ | --------------------------------- | ---------------------------- |
| Verify links             | Quarterly                         | Broken links removed/updated |
| Refresh examples         | When API/service semantics change | Keep minimal                 |
| Remove outdated sections | Immediately                       | Avoid historical clutter     |

---

## 🛠 Troubleshooting Section (Optional)

Use Q/A bullets. Each answer ≤6 lines. Link to source or spec when helpful.

---

## ✅ DO / ❌ AVOID

| ✅ DO                     | ❌ AVOID                        |
| ------------------------- | ------------------------------- |
| Concise actionable steps  | Narrative storytelling          |
| Tables for dense config   | Inline comma-separated lists    |
| Collapsible long examples | Flooding page with raw logs     |
| Relative links            | Hard-coded absolute GitHub URLs |
| Explicit environment vars | Implicit assumptions            |

---

End of Markdown Instructions.
