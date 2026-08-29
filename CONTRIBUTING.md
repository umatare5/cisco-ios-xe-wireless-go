# 🤝 Contribution Guide

Thank you for your interest in contributing to the **Cisco Catalyst 9800 WNC Go SDK**!

This document explains how you can get involved, the development workflow, and our release process.

> [!WARNING]
> This SDK is under **active development**, so expect breaking changes. Open an issue before contributing.

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

`make deps` installs Go tooling only, so install [pre-commit](https://pre-commit.com/#install) yourself first. `make pre-commit-install` then registers the hooks declared in [.pre-commit-config.yaml](./.pre-commit-config.yaml).

Every commit is gated on `golangci-lint`, `gitleaks`, `markdownlint-cli2` and a guard that keeps `main` free of direct commits. Run `make pre-commit-test` to check the whole tree without committing.

### Quick Build & Tests

```bash
export WNC_CONTROLLER="<controller-host-or-ip>"
export WNC_ACCESS_TOKEN="<base64-username:password>"
export WNC_AP_MAC_ADDR="<test-ap-radio-mac-address>"     # Pick a MAC Address from ./example/list_aps result.
export WNC_CLIENT_MAC_ADDR="<test-client-mac-address>"   # Pick a MAC Address from ./example/list_clients result.
export WNC_AP_WLAN_BSSID="<test-ap-wlan-bssid>"          # Pick a BSSID from ./example/list_wlans result.
export WNC_AP_NEIGHBOR_BSSID="<test-ap-neighbor-bssid>"  # Pick a BSSID from ./example/list_neighbors result.

make lint                # Static analysis
make test-unit           # Run unit tests
make test-unit-coverage  # Check unit test coverage
make test-integration    # Run integration test using live WNC
```

## 🧪 Testing

This SDK includes **unit, contract, integration and scenario tests** for reliability against Cisco Catalyst 9800 controllers.

- **Unit tests** run without any external dependencies.
- **Contract tests** hold every decode type to the route that reads it, and fail `make test-unit`.
- **Integration tests** require a live WNC instance and valid credentials.
- **Scenario tests** perform end-to-end operations on a live WNC and may modify its state.

For detailed testing instructions, see **[TESTING.md](./docs/TESTING.md)**.

## 📜 Scripts

Every Make target above runs one script from the `scripts/` directory.

Each takes no arguments and wraps the Go toolchain, so `make` is the only entry point you need. For what each one does, see **[SCRIPT_REFERENCE.md](./docs/SCRIPT_REFERENCE.md)**.

## ♻️ Change Review Process: For Maintainers

> [!Note]
>
> This section is for maintainers. Contributors do not need to perform these steps.

GitHub Actions cannot access a live WNC, so a reviewer needs a working WNC development environment.

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
> **About Destructive Examples**
>
> - `example/reset_ap` and `example/reload_controller` will reboot the AP and controller. This causes downtime.
> - `example/save_config` overwrites the startup configuration, which cannot be undone.

#### 5. Generate Coverage Reports and Badge

Generate and commit coverage reports:

```bash
make test-unit-coverage # writes coverage/report.html and coverage/report.out
octocov badge coverage --out docs/assets/coverage.svg # generates coverage badge
```

Commit coverage artifacts and badge:

- `coverage/report.out` - coverprofile `.octocov.yml` reads to build the badge
- `coverage/report.html` - human-readable report
- `docs/assets/coverage.svg` - coverage badge

#### 6. Push the Changes

Push the coverage artifacts and badge to the PR.

## 🚀 Release Process: For Maintainers

> [!Note]
>
> This section is for maintainers. Contributors do not need to perform these steps.

To release a new version:

- **Update the version** in the `VERSION` file and in `internal/version.Version` (`internal/version/version.go`) — `version_test.go` fails unless both move in the same commit.
- **Submit a pull request** with the updated `VERSION` file.

Once merged, GitHub Actions will automatically release the new version using [Release Workflow](https://github.com/umatare5/cisco-ios-xe-wireless-go/actions/workflows/go-release.yml).

### Versioning Rules

While this SDK is on a `0.x` line, a **MINOR release may ship breaking changes**. A PATCH release may add to the exported API but never changes or removes what is already there.

| Change                    | Bump  | Release notes must name                |
| :------------------------ | :---- | :------------------------------------- |
| Added exported symbol     | PATCH | The new symbols                        |
| Added variadic parameter  | MINOR | The broken forms and the recovery path |
| Changed or removed symbol | MINOR | Every affected symbol                  |
| No exported API change    | PATCH | Nothing                                |

### Toolchain Requirement

The `go` directive in [go.mod](./go.mod) and the `go_version` inputs under [.github/workflows](./.github/workflows) move together. Raising the directive lifts the toolchain floor of every consumer, so it ships as a MINOR release with the new floor named in the release notes.

## 🔢 Typing Rules for Wire Values

This platform encodes **per leaf, not per container**: one body carries bare numbers and quoted strings side by side.

Type every field from a measured response, never from the YANG model.

| Wire form                                  | Go type                               |
| :----------------------------------------- | :------------------------------------ |
| Quoted number: `decimal64`, 64-bit counter | `string`, or `*string` when omittable |
| Bare number: `uint32` and narrower         | `int` or `uintN`, `*T` when omittable |
| Enumeration                                | `string`, or a named `string` type    |

- Retype a leaf only when that leaf was measured. Two `uint64` siblings in one struct can differ.
- One leaf that will not decode fails the whole read: `encoding/json` refuses an out-of-range integer and both decode paths discard the partly-filled value, at `internal/core/envelope.go:37-40` on a read and `internal/core/request.go:206-208` on a write response. Narrowing a numeric type therefore risks a collection, not a field.
- Do not add `,string` to a numeric field, and do not introduce a third numeric convention for a single leaf.
- Pointerize any leaf whose absence carries meaning, and give it `,omitempty`.
- An omitted configuration leaf means its default is in force, and that default is often `true`.

## 📖 Reference

### Adding New Service

Copy the shape of an existing service, then satisfy each rule below. [tests/contract](./tests/contract) gates the envelope and route rules.

- **Routes**: one constant per path, and no second accessor on a route that already has one.
- **Envelope**: one field per accessor, tagged with the node the route ends in.
- **Qualification**: qualify only that outermost tag — a prefix below it never matches.
- **Read options**: every read forwards `opts ...core.GetOption`, delegating reads included.
- **Leaf types**: follow [Typing Rules for Wire Values](#-typing-rules-for-wire-values).
- **Leaf comments**: `(Live: IOS-XE <version>)` when measured, `(YANG: IOS-XE <version>)` otherwise.
- **Keyed reads**: reject an empty key with `core.ErrResourceNotFound`, then normalize it.
- **Registration**: one facade accessor, a `doc.go`, and a row in [README.md](./README.md#supported-services).
- **Tests**: name unit tests `Test<Service>ServiceUnit_<Area>_<Case>`.
- **Scenario tests**: only when the operation changes controller state.

### Adding New Function to an Existing Service

- [f595bb: feat/ap: add IoT firmware information support](https://github.com/umatare5/cisco-ios-xe-wireless-go/pull/27/commits/f595bbf830802703dce950ba42df3ee411d00b9a).

### HTTP Transport

`NewTransport` tunes the transport for a single controller rather than copying `http.DefaultTransport`: HTTP/2 is off, and the handshake and header budgets are 5 s each. It uses no proxy by default, so pass `WithProxy(http.ProxyFromEnvironment)` to route through one.

Those two budgets are `DefaultTLSHandshakeTimeout` and `DefaultResponseHeaderTimeout`, and `WithTimeout` lifts neither: it bounds the whole request. Raise them with `WithTLSHandshakeTimeout` and `WithResponseHeaderTimeout`, one option each — every option mutates one shared transport in place, so a single option setting all three would overwrite a smaller budget an earlier option had set.
