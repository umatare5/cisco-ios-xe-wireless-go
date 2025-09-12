# 🤝 Contribution Guide

Thank you for your interest in contributing to the **Cisco Catalyst 9800 WNC Go library**! This document explains how you can get involved, the development workflow, and our release process.

> [!WARNING]
> This library is under **active development**. I'll make the breaking changes until `v1.0.0`. Please create an issue before contributing to avoid duplicate work.
>
> - The remaining tasks to reach `v1.0.0` are tracked in **[Milestone: 1.0.0](https://github.com/umatare5/cisco-ios-xe-wireless-go/milestone/1)**.

## 💡 How to Contribute

I welcome all kinds of contributions, including:

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

I provide `make` commands and helper scripts for building, testing, and debugging this library.
The helper scripts use `curl` to access WNC directly, so they have **no dependency on Go**.

> **Note:** Integration tests require access to a live Cisco Catalyst 9800 WNC instance.
> Set `WNC_ACCESS_TOKEN` and `WNC_CONTROLLER` before running them.

### Prerequisites

Before running the build and test commands, you must install dependencies and set up pre-commit hooks:

```bash
make deps              # Install build and test dependencies
make pre-commit-install # Set up pre-commit hooks for code quality
```

### Quick Build & Tests

```bash
export WNC_CONTROLLER="<controller-host-or-ip>"
export WNC_ACCESS_TOKEN="<base64-username:password>"

make lint                    # Static analysis
make test-unit               # Run unit tests (runs lint first)
make test-integration        # Test live connection to WNC
make test-unit-coverage      # Check unit test coverage
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

They use `curl` to access WNC, so they are independent of Go. For detailed usage, see **[MAKE_REFERENCE.md](./docs/MAKE_REFERENCE.md)**.

---

## Change Review Process _(Maintainers Only)_

> [!Note]
>
> This section is for maintainers. Contributors do not need to perform these steps.

GitHub Actions cannot access a live WNC. Therefore, the maintainers have to run the integration and scenario tests manually to validate the behavior.

- Ensure you have access to a development Cisco C9800 WNC (RESTCONF enabled) and export the required env vars:

  ```bash
  export WNC_CONTROLLER="<controller-host-or-ip>"
  export WNC_ACCESS_TOKEN="<base64-username:password>"
  ```

- Run unit and integration tests as follows:

  ```bash
  make test-unit
  make test-integration
  ```

- Run scenario tests as follows:

  ```bash
  # Run AP Admin State Change and AP Radio State Change Test
  go test ./tests/scenario/ap/ -tags=scenario -run TestAPAdminStateScenario -v
  go test ./tests/scenario/ap/ -tags=scenario -run TestAPRadioStateScenario -v

  # Run Site Tag, Policy Tag and RF Tag Test
  go test ./tests/scenario/tag/ -tags=scenario -run TestSiteTagScenario -v
  go test ./tests/scenario/tag/ -tags=scenario -run TestPolicyTagScenario -v
  go test ./tests/scenario/tag/ -tags=scenario -run TestRFTagScenario -v
  ```

- Run destructive operations as follows:

  ```bash
  # Reload AP
  go run ./example/reload_ap/main.go

  # Reload WNC
  go run ./example/reload_controller/main.go
  ```

- Generate and commit coverage reports:

  ```bash
  make test-unit-coverage # writes coverage/report.html and coverage/report.out
  octocov badge coverage --out docs/assets/coverage.svg # generates coverage badge
  ```

- Commit coverage artifacts (CI will build the badge):

  - Commit the updated files:
    - `coverage/report.out` (coverprofile for CI)
    - `coverage/report.html` (human-readable report)
    - `docs/assets/coverage.svg` (coverage badge)

- In the PR description, mention the resulting total coverage and, if helpful, link to `coverage/report.html`.

Notes:

- CI cannot access a WNC instance; manual execution is required to validate integration behavior.
- Reviewers therefore must have a functional WNC development environment to complete reviews.

## 🚀 Release Process _(Maintainers Only)_

> [!Note]
>
> This section is for maintainers. Contributors do not need to perform these steps.

To release a new version:

- **Update the version** in the `VERSION` file.
- **Submit a pull request** with the updated `VERSION` file.

Once merged, GitHub Actions will automatically release the new version using [Release Workflow](https://github.com/umatare5/cisco-ios-xe-wireless-go/actions/workflows/go-release.yml).
