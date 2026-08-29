# Repository Instructions

> [!IMPORTANT]
> Read [README.md](README.md) for project overview and architecture.

## Tech Stack

- Go 1.27+ (see [go.mod](go.mod))
- **No third-party dependencies** — `go.mod` carries no `require` block and there is no `go.sum`, in production and in test code
- [`golangci-lint`](https://golangci-lint.run/) — the authority for lint, formatting, and suppressions; exits non-zero on any finding (see [.golangci.yml](.golangci.yml))
- [`goreleaser`](https://goreleaser.com/) v2 — cross-platform release builds (see [.goreleaser.yml](.goreleaser.yml))
- [`octocov`](https://github.com/k1LoW/octocov) — coverage gate (see [.octocov.yml](.octocov.yml))

## Repository Structure

- `wnc.go` — Public entry point; `NewClient(host, token, opts...)`, 42 service accessors, and the untyped `GetData`/`PostData`/`PutData`/`PatchData`/`DeleteData`/`PostRPC`/`Request` escape hatches, of which `Request` alone returns a `*Response` carrying the status
- `service/` — 32 service packages, one per YANG feature area. **No service imports another**; the invariant holds across all 32 and nothing enforces it, because `depguard` runs `list-mode: lax`
- `internal/core/` — Request execution and the five error sentinels ([errors.go](internal/core/errors.go)); a service package never defines its own sentinel, and `IsNotFoundError` is the intended predicate
- `internal/restconf/` — URL builders reached as `s.Client().RESTCONFBuilder()`; `routes/` is the single source for endpoint strings
- `internal/transport/` — HTTP client and request builder; the token is a client-lifetime field, not request-scoped
- `internal/{errors,service,validation}/` — Message templates, the service base type and MAC helpers, and input validation
- `pkg/testutil/` — `NewMockServer` with functional options; the mock server both tests and consumers use
- `tests/` — `contract/`, `integration/` (`//go:build integration`, read-only), `scenario/` (`//go:build scenario`, **writes to a controller**), and shared `testutil/`
- `example/` — Seven runnable programs; the only place the public API is exercised as a consumer would

## Setup and Commands

Install required tools (one-time):

- `make deps` installs `golangci-lint`, `goreleaser`, `gotestsum` and `gitleaks`
- `markdownlint-cli2` comes from npm and is not installed here — `make lint` skips Markdown when it is absent
- `make pre-commit-install` symlinks [.githooks/pre-commit](.githooks/pre-commit) into `.git/hooks/`, which every worktree shares

Make targets ([Makefile](Makefile), documented in [docs/MAKE_REFERENCE.md](docs/MAKE_REFERENCE.md)):

- `make build` — Verify compilation across the module
- `make lint` — `golangci-lint run` + `go mod tidy`
- `make test-unit` — Run unit tests with coverage
- `make test-integration` — Run `tests/integration/` against a live controller; needs `WNC_CONTROLLER` and `WNC_ACCESS_TOKEN`

## Code Style

- The git hook scans the staged index with `gitleaks` and refuses a commit on `main` (see [.githooks/pre-commit](.githooks/pre-commit)).
- Format and lint are CI's, not the hook's — [.github/workflows/go-test-fmt.yml](.github/workflows/go-test-fmt.yml) pins its own `golangci-lint`, so the hook keeps no second copy of the version.
- [.golangci.yml](.golangci.yml) owns lint, formatting, initialism casing, and suppressions — add a suppression there with its reason rather than inline.
- Run `make lint` and `make test-unit` before declaring work finished; the hook fires only at commit time.
- Every service package carries its package comment in `doc.go` alone, listing RESTCONF endpoints and YANG references with the trains each was verified against — see [service/wlan/doc.go](service/wlan/doc.go).

## Testing Instructions

- Run `make test-unit` before committing.
- Place unit tests next to the code under test (`*_test.go`) and name them `Test{Service}{Tier}_{Category}_{Scenario}`.
- **Base a mock payload on a real controller response, never on a YANG model.** 38 files under `service/` carry `Live: IOS-XE` annotations recording the release each shape came from; capturing a new one is a maintainer step.
- Assertions use the hand-rolled helpers in `internal/testutil/`, not a third-party library.
- Simulate a release gap by mocking `404`, not by branching on a version.
- Coverage threshold is enforced by [.octocov.yml](.octocov.yml); see [docs/TESTING.md](docs/TESTING.md) for the four tiers.

## Commits and PRs

- Use [Conventional Commits](https://www.conventionalcommits.org/) (`feat:`, `fix:`, `chore(deps):`, etc.).
- Sign off commits with `Signed-off-by:` (DCO).
- Open PRs against `main`. CI runs format and lint, test and build, coverage, actionlint, and CodeQL.
- The public surface follows a schema Cisco owns and revises per IOS-XE train, so a surface change is an ordinary consequence of tracking the vendor — name the train it was verified against, in `doc.go` and in the release notes.

## Domain Knowledge

### Verifying Values

- **A YANG model is a design document, not the implementation.** Units, ranges, enum spellings, and even the presence of a leaf can differ on a live controller, so confirm every value against a RESTCONF response from a real WNC before relying on it.
- **A configuration leaf missing from a response means its default is in force, not that nothing set it.** The default is often `true`, so decoding an omitted boolean as `false` inverts the reading. Appending `?with-defaults=report-all` returns the omitted leaves, and a controller that rejects the parameter answers `400`.
- **A body-less read is a successful read of a node holding nothing.** `core.Get` returns `T` at its zero value — a nil slice for a list node — rather than an error, which is what `tests/contract` pins.
- **A read-modify-write must route a carried value through the default builder, never into `""`.** Because the read omits a leaf holding its default, writing the omitted value back as an empty string is what the controller rejects.
- **For a replacing `PUT` the write type must cover every leaf the read type declares.** A leaf the write type omits is deleted from the entry, and no test compares the two types — see the repair in `db757c7`.
- **An optional container must be a pointer with `omitempty`.** A non-pointer struct always marshals, so the payload carries `"list": null` and the controller answers `400`.
- **A secret-bearing leaf is a named type, and the redaction has to sit on the field type and on the entry.** `wlan.Secret` redacts through `String` and `LogValue`, which covers every `fmt` verb but `%#v`. `slog.JSONHandler` renders a struct value through `json.Marshal`, which honors `json.Marshaler` and ignores `fmt.Stringer`, so `WlanCfgEntry.LogValue` is what keeps the key out of a record built from the whole entry, and `Secret.MarshalJSON` is what keeps it out of the containers above the entry, which have no `LogValue` of their own.
- **Tag writes, the AP write RPCs, the client deauthentication RPC and the controller reload and save-config RPCs are the only places this SDK mutates a controller.** The save is the one whose effect no read here can observe, because no release in scope exposes the startup datastore over RESTCONF. The three tag list key leaves declare the pattern `[!-~]([ -~]*[!-~])?` and no length on 17.12, 17.15 and 17.18 alike, while the controller caps a name at 32 characters on every kind, measured on 17.12.8, so `validation.ValidateTagName` enforces both and each tag service calls it.
- **An AP write RPC's input is a mandatory choice, and the accessor's `By` suffix names the arm it fills.** Both arms carry `omitempty`, so the arm the caller did not name is absent rather than sent empty. The slot-admin RPC's `band` selects a radio type and not a frequency, so the caller supplies the slot and the `radio-type` leaf and the controller arbitrates the pair.

### RESTCONF Access Patterns

GET a collection:

```bash
curl -k -H "Authorization: Basic $WNC_ACCESS_TOKEN" \
        -H "Accept: application/yang-data+json" \
        "https://$WNC_CONTROLLER/restconf/data/Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/capwap-data"
```

GET a single entry by list key (MAC address):

```bash
curl -k -H "Authorization: Basic $WNC_ACCESS_TOKEN" \
        -H "Accept: application/yang-data+json" \
        "https://$WNC_CONTROLLER/restconf/data/Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/capwap-data=00:11:22:33:44:55"
```

POST an RPC operation (`/restconf/operations/`):

```bash
curl -k -X POST \
        -H "Authorization: Basic $WNC_ACCESS_TOKEN" \
        -H "Content-Type: application/yang-data+json" \
        -d '{"input": {"ap-name": "AP-NAME"}}' \
        "https://$WNC_CONTROLLER/restconf/operations/Cisco-IOS-XE-wireless-access-point-cmd-rpc:ap-reset"
```

## References

- [CONTRIBUTING.md](CONTRIBUTING.md) — Contribution workflow and the local gate order
- [docs/TESTING.md](docs/TESTING.md) — The four test tiers and how to run each
- [docs/SECURITY.md](docs/SECURITY.md) — Vulnerability reporting and credential handling
- [docs/MAKE_REFERENCE.md](docs/MAKE_REFERENCE.md) — Every Make target
- [docs/SCRIPT_REFERENCE.md](docs/SCRIPT_REFERENCE.md) — Every script behind a Make target
