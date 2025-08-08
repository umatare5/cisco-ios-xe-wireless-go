# 🤝 CONTRIBUTING

Thank you for your interest in contributing to the **Cisco Catalyst 9800 WNC Go library**!
This document explains how you can get involved, the development workflow, and our release process.

## 💡 How to Contribute

We welcome all kinds of contributions, including:

- 🐞 Bug reports
- 📄 Documentation improvements
- 💡 Feature requests
- 🛠 Code contributions (new features, bug fixes, refactoring)

**Before you start coding:**

1. Check the [Issues](../../issues) to avoid duplicate work.
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
export WNC_CONTROLLER=your-wnc-hostname

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
For detailed usage, see **[SCRIPT_REFERENCE.md](./docs/SCRIPT_REFERENCE.md)**.

---

## 🚀 Release Process _(Maintainers Only)_

_This section is for maintainers. Contributors do not need to perform these steps._

To release a new version:

1. **Update the version** in the `VERSION` file.
2. **Submit a pull request** with the updated `VERSION` file.

Once merged, GitHub Actions will automatically:

- **Create and push a new tag** via [Tagging Workflow](https://github.com/umatare5/cisco-ios-xe-wireless-go/actions/workflows/tagging.yml).
- **Release the new version** via [Release Workflow](https://github.com/umatare5/cisco-ios-xe-wireless-go/actions/workflows/go-release.yml).
