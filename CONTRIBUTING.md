# 🤝 Contribution Guide

Thank you for your interest in contributing to the **Cisco Catalyst 9800 WNC Go SDK**! This document explains how you can get involved, the development workflow, and our release process.

> [!WARNING]
> This SDK is under **active development**. I'll make the breaking changes until `v1.0.0`. If you give the contribution to this repo, please create an issue before to avoid duplicate work. The remaining tasks to reach `v1.0.0` are tracked in **[Milestone: 1.0.0](https://github.com/umatare5/cisco-ios-xe-wireless-go/milestone/1)**.

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

This SDK includes **unit, contract, integration and scenario tests** to ensure reliability and compatibility with Cisco Catalyst 9800 controllers.

- **Unit tests** run without any external dependencies.
- **Contract tests** in [tests/contract](./tests/contract) hold every decode type to the route that reads it, and a violation fails `make test-unit`.
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

### Versioning Rules

This SDK is pre-`v1.0.0`, so a **MINOR release may ship breaking changes**. A PATCH release never changes the exported API.

| Change | Bump | Release notes must name |
| :--- | :--- | :--- |
| Added exported symbol | MINOR | The new symbols |
| Added variadic parameter | MINOR | The broken forms and the recovery path |
| Changed or removed symbol | MINOR | Every affected symbol |
| No exported API change | PATCH | Nothing |

Adding a variadic parameter is source-breaking for consumer-defined interfaces and method values, so it is never a PATCH.

Mark every breaking change with a Conventional Commits `!`. The release notes group those commits under **Breaking Changes**.

Before `v1.0.0` this SDK ships no deprecation window for an accessor. An accessor is removed in the release that repairs the resource it reads — whether it never decoded the node it names, or it merely duplicates the accessor that does — and the release notes name every removed symbol.

An escape hatch is the exception, because a consumer may have no replacement to move to. One is marked `Deprecated:` naming its removal release, and that release is a ceiling: removing it earlier breaks a promise a published tag already made.

The hatch this SDK ships is the untyped request methods on the root client: `GetData`, the four `*Data` verb methods, `PostRPC`, and `Request` for a method or a root they cannot express. It exists because the controller's schema moves between releases, so a node or an operation with no typed accessor still has to be reachable without waiting for one. `core.Client`'s transport methods are unexported, so these are the only route to the wire that is not a typed accessor. A further hatch is added against a named consumer need, never on the argument that the typed surface is incomplete.

### Toolchain Requirement

The `go` directive in [go.mod](./go.mod) and the `go_version` inputs under [.github/workflows](./.github/workflows) move together. Raising the directive lifts the toolchain floor of every consumer, so it ships as a MINOR release with the new floor named in the release notes.

## 🔢 Typing Rules for Wire Values

RESTCONF on this platform encodes **per leaf, not per container**: one response body carries bare numbers and quoted strings side by side. Type every field from a measured response — a YANG model states the value space, not the encoding.

| Wire form | Go type | Measured example |
| :--- | :--- | :--- |
| Quoted number: `decimal64`, 64-bit counter | `string`, or `*string` when omittable | `latitude`, `iapp-unconnected-client` |
| Bare number: `uint32` and narrower | `int` or `uintN`, `*T` when omittable | `adhoc-count`, `rogue-ap-mld-link-count` |
| Enumeration | `string`, or a named `string` type | `reboot-reason` |

- A Go `uint64` is not evidence of a quoted wire form. Two `uint64` siblings in one struct can differ, so retype a leaf only when that leaf was measured.
- Do not add `,string` to a numeric field. It accepts the quoted form and then rejects the bare one, so it cannot carry a leaf whose encoding differs between releases.
- Do not introduce a third numeric convention — `json.Number`, a custom `UnmarshalJSON`, or a quoted-tolerant named type — for a single leaf.
- Pointerize any leaf whose absence carries meaning and give it `,omitempty`. A configuration leaf omitted from a plain read means its default is in force, which is often `true`, so a value-typed field publishes a fabricated and sometimes inverted reading.

## 📖 Reference

### Module Charter

A service wraps the `Cisco-IOS-XE-wireless-*` models for one feature area. Two non-wireless models are in charter, and both live in [`service/controller`](./service/controller):

- `Cisco-IOS-XE-rpc` — controller reload ([internal/restconf/routes/controller.go](./internal/restconf/routes/controller.go#L13))
- `Cisco-IOS-XE-device-hardware-oper` — the controller's own boot instant ([internal/restconf/routes/controller.go](./internal/restconf/routes/controller.go#L25)). In charter on the same ground as the reload RPC: the subject is the controller as a system, not its wireless namespace. The route answers on every supported release.

Any other non-wireless model needs a decision recorded here before a route is added.

### Adding New Service

A new service must satisfy the six conventions below. The 45 accessors this release removes trace to two of them: the duplicate-accessor clause in §3 and the envelope rule in §4.

#### 1. Service shape and facade wiring

- Declare `type Service struct { service.BaseService }` and `func NewService(client *core.Client) Service` — a value type, not a pointer. See [service/rogue/service.go](./service/rogue/service.go#L12-L20) and [internal/service/base.go](./internal/service/base.go#L12-L29).
- Expose the service from the facade as exactly one accessor, `func (c *Client) X() x.Service`. See [wnc.go](./wnc.go).
- Add the service row to the **Supported Services** table in [README.md](./README.md#supported-services).
- Write a `doc.go` whose `RESTCONF Endpoints:` list matches the service's route file, and whose `YANG References:` list names only the modules those routes use.

#### 2. Route constants

- Declare every path as a constant in `internal/restconf/routes/<service>.go`, built from `RESTCONFDataPath` or `RESTCONFOperationsPath`. See [internal/restconf/routes/base.go](./internal/restconf/routes/base.go#L4-L9).
- Never assemble a path at the call site. One route, one constant, one accessor.
- Controller-level, non-wireless modules belong in `service/controller` alone, per the Module Charter above.

#### 3. Method naming and signature

- Name reads `GetOperational`, `GetConfig`, `List<Subtree>` and `Get<Subtree>ByMAC`. See [service/wlan/service.go](./service/wlan/service.go#L27-L35) and [service/client/service.go](./service/client/service.go).
- Every read method takes `opts ...core.GetOption` after `ctx` and forwards them, including a method that delegates to another read rather than to `core.Get`. See [service/rf/tag_service.go](./service/rf/tag_service.go#L27-L32) for the direct form and [service/site/tag_service.go](./service/site/tag_service.go#L51-L56) for the delegating one.
- Never add a second accessor to a route that already has one. A duplicate pair leaves consumers no way to tell which twin decodes.
- A keyed read is a different resource, not a duplicate: `BuildQueryURL` appends `=<key>`, so `List<X>` beside `Get<X>ByMAC` is the correct pair.

#### 4. Decode types

- A GET of container `X` answers exactly one module-qualified top-level key naming `X`. Every accessor decodes that envelope.
- Name the envelope after its module and node in camel case, give it exactly one field, and tag that field with the module-qualified node the route ends in. A `Cisco-IOS-XE-wireless-*` model yields `CiscoIOSXEWireless<Module><Node>` — see [service/ap/global_oper.go](./service/ap/global_oper.go#L22-L24) and [service/wlan/cfg.go](./service/wlan/cfg.go#L18-L20) — and a non-wireless model in charter carries no `Wireless`, as in [service/controller/oper.go](./service/controller/oper.go#L12-L14).
- **Qualify only that outermost tag.** A same-module child takes the bare node name; a module prefix below the top level never matches, so the field stays at its zero value while its siblings decode.
- Annotate each leaf comment `(Live: IOS-XE <version>)` when the leaf was seen in a controller response, and `(YANG: IOS-XE <version>)` otherwise. A YANG model is a design document, not the implementation.
- The envelope rules above are enforced by [tests/contract](./tests/contract); a violation fails `make test-unit`. The annotation rule is not gated — it is reviewed.

#### 5. Leaf typing

- Type every leaf by the rules in [Typing Rules for Wire Values](#-typing-rules-for-wire-values). Do not restate them here.

#### 6. Keyed reads and their sentinel

- A `Get<X>ByMAC` method rejects an empty key with `core.ErrResourceNotFound`, then validates and normalizes the key through the shared MAC helper in [internal/validation](./internal/validation), then builds the URL with `BuildQueryURL`. See [service/geolocation/service.go](./service/geolocation/service.go).
- Return the helper's error unchanged. Use that one sentinel and do not invent a second one for the same condition — `core.ErrInvalidConfiguration` reports a bad client configuration, not a bad list key.

#### Tests and release

- Put unit tests in `service/<service>/service_test.go`, package `<service>_test`, driving the mock server from [pkg/testutil](./pkg/testutil), and name them `Test<Service>ServiceUnit_<Area>_<Case>`. See [docs/TESTING.md](./docs/TESTING.md).
- Add the service's live-controller table to `tests/integration/<service>_service_test.go`, and add a scenario test under `tests/scenario/<service>/` only when the operation changes controller state.
- Release by updating the `VERSION` file, per [Versioning Rules](#versioning-rules).

### Adding New Function to an Existing Service

- [f595bb: feat/ap: add IoT firmware information support](https://github.com/umatare5/cisco-ios-xe-wireless-go/pull/27/commits/f595bbf830802703dce950ba42df3ee411d00b9a).

### HTTP Transport

- `NewTransport` is tuned for a single controller and is not a copy of `http.DefaultTransport`: HTTP/2 is off and both the handshake and header budgets are 5 s. No proxy is used by default — `Proxy` is unset and no environment variable is read. Pass `WithProxy(http.ProxyFromEnvironment)` to route through one.
