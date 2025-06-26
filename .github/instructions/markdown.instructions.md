---
description: Markdown Documentation Instructions
applyTo: "**/*.md"
---

# GitHub Copilot Agent Mode – Markdown Documentation Instructions

Copilot **MUST** comply with all instructions described in this document when editing or creating any Markdown files in this repository.

However, when there are conflicts between this document and `general.instructions.md`, **ALWAYS** prioritize the instructions in `general.instructions.md`.

---

## 🎯 Primary Goal

Create clear, comprehensive documentation that enables developers to successfully use this repository.

---

## 🛠️ Documentation Practices & Style

Apply the following style and best practices to all Markdown files:

- **Style & Linting:**
  Ensure all scripts pass `markdownlint-cli2`.

- **Headings & Emojis:**
  Use a single H1 (`#`) for the main document title. Use H2 (`##`) and H3 (`###`) for major sections and sub-sections, respectively. Prefix headings with a relevant emoji to visually represent the content.

- **Text Formatting:**
  Use `**bold**` text to emphasize key terms, warnings, or important notes. Use `` `inline code` `` for file paths, variable names, commands, and other code-related terms.

- **Consistent Structure:**
  Maintain a logical and consistent document structure across all files. A typical structure includes an introduction, setup/usage sections, detailed guides, and references.

- **Code Blocks:**
  Fence all multi-line code snippets with triple backticks (`` ` ``) and specify the language (e.g., `go`, `bash`, `json`, `text`) for proper syntax highlighting.

- **Tables:**
  Use Markdown tables to present structured data, such as configuration options, environment variables, or command references, for clarity and easy comparison.

- **Lists:**
  Use hyphens (`-`) for unordered lists and sequential numbers (`1.`, `2.`) for ordered lists, especially for step-by-step instructions.

- **Alerts for Emphasis:**
  Use GitHub-flavored Markdown alerts (e.g., `> [!CAUTION]`, `> [!WARNING]`) to draw attention to critical information like security risks or important usage notes.

- **Collapsible Sections:**
  Use `<details>` and `<summary>` tags to enclose lengthy, non-critical content like sample outputs or verbose examples, improving overall document readability.

- **Links:**
  Use standard Markdown syntax `[link text](url)` for all hyperlinks. For links to other documents in the repository, use relative paths.

- **Readability:**
  Write in clear, concise language. Use short paragraphs and ensure sufficient whitespace between elements like sections, lists, and code blocks to improve readability.
