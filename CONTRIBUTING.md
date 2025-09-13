# 🤝 Contribution Guide

Thank you for your interest in contributing to the **Cisco Catalyst 9800 WNC Go SDK**! This document explains how you can get involved, the development workflow, and our release process.

> [!WARNING]
> This SDK is under **active development**. I'll make the breaking changes until `v1.0.0`. Please create an issue before contributing to avoid duplicate work. The remaining tasks to reach `v1.0.0` are tracked in **[Milestone: 1.0.0](https://github.com/umatare5/cisco-ios-xe-wireless-go/milestone/1)**.

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

I provide `make` commands and helper scripts for building, testing, and debugging this SDK.

### Prerequisites

Before running the build and test commands, you must install dependencies and set up pre-commit hooks:

```bash
make deps                # Install build and test dependencies
make pre-commit-install  # Set up pre-commit hooks for code quality
```

### Quick Build & Tests

```bash
export WNC_CONTROLLER="<controller-host-or-ip>"
export WNC_ACCESS_TOKEN="<base64-username:password>"
export WNC_AP_MAC_ADDR="<test-ap-radio-mac-address>"     # Pick a MAC Address from ./examples/list_ap.go result.
export WNC_CLIENT_MAC_ADDR="<test-client-mac-address>"   # Pick a MAC Address from ./examples/list_clients.go result.
export WNC_AP_WLAN_BSSID="<test-ap-wlan-bssid>"          # Pick a BSSID from ./examples/list_wlan.go result.
export WNC_AP_NEIGHBOR_BSSID="<test-ap-neighbor-bssid>"  # Pick a BSSID from ./examples/list_neighbors.go result.

make lint                # Static analysis
make test-unit           # Run unit tests
make test-unit-coverage  # Check unit test coverage
make test-integration    # Run integration test using live WNC
```

## 🧪 Testing

This SDK includes **unit, integration and scenario tests** to ensure reliability and compatibility with Cisco Catalyst 9800 controllers.

- **Unit tests** run without any external dependencies.
- **Integration tests** require a live WNC instance and valid credentials.
- **Scenario tests** perform end-to-end operations on a live WNC and may modify its state.

For detailed testing instructions, see **[TESTING.md](./docs/TESTING.md)**.

## 📜 Scripts

This repository contains useful debugging and development scripts in the `scripts/` directory.

They use `curl` to access WNC, so they are independent of Go. For detailed usage, see **[SCRIPT_REFERENCE.md](./docs/SCRIPT_REFERENCE.md)**.

## ♻️ Change Review Process: For Maintainers

> [!Note]
>
> This section is for maintainers. Contributors do not need to perform these steps.

GitHub Actions cannot access a live WNC. Reviewers therefore must have a functional WNC development environment to complete reviews.

### Verify the Changes using a Live WNC

Ensure you have access to a development Cisco C9800 WNC that enabled RESTCONF and export the required env vars.

#### 1. Run the Unit Tests

Run unit tests as follows:

```bash
make test-unit
```

#### 2. Run the Integration Tests

Run integration tests as follows:

```bash
make test-integration
```

#### 3. Run the Scenario Tests

Run scenario tests as follows:

```bash
# Run AP Admin State Change and AP Radio State Change Test
go test ./tests/scenario/ap/ -tags=scenario -run AdminStateManagement -v
go test ./tests/scenario/ap/ -tags=scenario -run RadioStateManagement -v

# Run RF Tag, Site Tag and Policy Tag Test
go test ./tests/scenario/rf/ -tags=scenario -run TagLifecycleManagement -v
go test ./tests/scenario/site/ -tags=scenario -run TagLifecycleManagement -v
go test ./tests/scenario/wlan/ -tags=scenario -run TagLifecycleManagement -v
```

#### 4. Run the Example Application

Run the example application listed in the [README.md](../README.md#-usecases) **Usecases** section.

> [!Warning]
>
> `example/reload_ap` and `example/reload_controller` will reboot the AP and controller. This causes downtime.

#### 5. Generate Coverage Reports and Badge

Generate and commit coverage reports:

```bash
make test-unit-coverage # writes coverage/report.html and coverage/report.out
octocov badge coverage --out docs/assets/coverage.svg # generates coverage badge
```

Commit coverage artifacts and badge:

- `coverage/report.out` - coverprofile for CI
- `coverage/report.html` - human-readable report
- `docs/assets/coverage.svg` - coverage badge

#### 6. Push the Changes

Push the coverage artifacts and badge to the PR.

## 🚀 Release Process: For Maintainers

> [!Note]
>
> This section is for maintainers. Contributors do not need to perform these steps.

To release a new version:

- **Update the version** in the `VERSION` file.
- **Submit a pull request** with the updated `VERSION` file.

Once merged, GitHub Actions will automatically release the new version using [Release Workflow](https://github.com/umatare5/cisco-ios-xe-wireless-go/actions/workflows/go-release.yml).
