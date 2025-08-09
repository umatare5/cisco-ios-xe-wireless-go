# 🤝 Contribution Guide

Thank you for your interest in contributing to the **Cisco Catalyst 9800 WNC Go library**!
This document explains how you can get involved, the development workflow, and our release process.

## 💡 How to Contribute

We welcome all kinds of contributions, including:

- 🐞 Bug reports
- 📄 Documentation improvements
- 💡 Feature requests
- 🛠 Code contributions (new features, bug fixes, refactoring)

**Before you start coding:**

1. Check the [Issues](https://github.com/umatare5/cisco-ios-xe-wireless-go/issues) to avoid duplicate work.
2. Open a new issue if your change is significant or affects functionality.
3. Fork this repository and create a feature branch from `main`.
4. Follow the [Development](#️-development) and [Testing](#-testing) guidelines below.
5. Submit a pull request to the `main` branch.

## 🛠️ Development

We provide `make` commands and helper scripts for building, testing, and debugging this library.
The helper scripts use `curl` to access WNC directly, so they have **no dependency on Go**.

> **Note:** Integration tests require access to a live Cisco Catalyst 9800 WNC instance.
> Set `WNC_ACCESS_TOKEN` and `WNC_CONTROLLER` before running them.

### Quick Build & Tests

```bash
export WNC_ACCESS_TOKEN=your-wnc-access-token
export WNC_CONTROLLER=wnc1.example.internal

make lint             # Static analysis
make test-unit        # Run unit tests (runs lint first)
make test-integration # Test live connection to WNC
make test-coverage    # Check code coverage
```

## 🧪 Testing

This library includes **comprehensive unit and integration tests** to ensure reliability and compatibility with Cisco Catalyst 9800 controllers.

- **Unit tests** run without any external dependencies.
- **Integration tests** require a live WNC instance and valid credentials.

For detailed testing instructions, see **[TESTING.md](./docs/TESTING.md)**.

## 📜 Scripts

This repository contains useful debugging and development scripts in the `scripts/` directory.

These scripts are particularly helpful for:

- Exploring new API endpoints quickly
- Debugging API responses without building the Go library

They use `curl` to access WNC, so they are independent of Go.
For detailed usage, see **[MAKE_REFERENCE.md](./docs/MAKE_REFERENCE.md)**.

---

## Review Process

Because the WNC is not reachable from CI, automated pipelines cannot run integration tests or produce coverage.
PR reviewers are responsible for executing tests and refreshing the coverage report locally.

Reviewer checklist:

- Ensure you have access to a development Cisco C9800 WNC (RESTCONF enabled) and export the required env vars:

```bash
export WNC_CONTROLLER=<controller-host-or-ip>
export WNC_ACCESS_TOKEN=<base64-username:password>
```

- Run tests and generate coverage outputs:

```bash
make test-coverage        # produces tmp/coverage.out
make test-coverage-report # writes coverage/report.html and coverage/report.out
```

- Commit coverage artifacts (CI will build the badge):

  - Commit the updated files:
    - `coverage/report.out` (coverprofile for CI)
    - `coverage/report.html` (human-readable report)
  - CI generates `docs/assets/coverage.svg` from `coverage/report.out` automatically.

- In the PR description, mention the resulting total coverage and, if helpful, link to `coverage/report.html`.

Notes:

- CI cannot access a WNC instance; manual execution is required to validate integration behavior and coverage.
- Reviewers therefore must have a functional WNC development environment to complete reviews.

## 🚀 Release Process _(Maintainers Only)_

_This section is for maintainers. Contributors do not need to perform these steps._

To release a new version:

- **Update the version** in the `VERSION` file.
- **Submit a pull request** with the updated `VERSION` file.

Once merged, GitHub Actions will automatically:

- **Create and push a new tag** via [Tagging Workflow](https://github.com/umatare5/cisco-ios-xe-wireless-go/actions/workflows/tagging.yml).
- **Release the new version** via [Release Workflow](https://github.com/umatare5/cisco-ios-xe-wireless-go/actions/workflows/go-release.yml).
