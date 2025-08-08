# 📋 Scripts Reference

Unified reference for helper scripts under `scripts/` that assist YANG model discovery and operational data inspection. Each script is idempotent and read‑only.

| Script                          | Purpose                                           | Core Operation            |
| ------------------------------- | ------------------------------------------------- | ------------------------- |
| `get_yang_models.sh`            | List available YANG modules on controller         | RESTCONF modules listing  |
| `get_yang_model_details.sh`     | Fetch a single model definition                   | RESTCONF module retrieval |
| `get_yang_statement_details.sh` | Fetch live data for specific model root/statement | RESTCONF data path query  |

> [!NOTE]
> All scripts honor environment variables: `WNC_CONTROLLER`, `WNC_ACCESS_TOKEN`. CLI flags override env values when both supplied.

## 🗂️ Directory Structure

High‑level layout (supporting libraries collapsed for brevity).

```text
scripts/
├── get_yang_models.sh
├── get_yang_model_details.sh
├── get_yang_statement_details.sh
└── lib/
  ├── bootstrap.sh
  ├── core/
  ├── network/
  ├── yang/
  └── ... (utilities)
```

## 🚀 Scripts Overview

| Script                          | Description                       | Output Formats                                |
| ------------------------------- | --------------------------------- | --------------------------------------------- |
| `get_yang_models.sh`            | List module names + revisions     | Pretty text                                   |
| `get_yang_model_details.sh`     | Retrieve single module definition | JSON (default), XML (`-f xml`), raw (`--raw`) |
| `get_yang_statement_details.sh` | Retrieve live data subtree        | JSON (default), XML (`-f xml`)                |

## 📋 `get_yang_models.sh` – Module Discovery

Lists all available Cisco wireless YANG models (module + revision) exposed by the controller RESTCONF modules endpoint.

### Key Features

| Capability          | Notes                                     |
| ------------------- | ----------------------------------------- |
| Protocol selection  | `--protocol http\|https` (default: https) |
| TLS skip            | `--insecure` for lab/self‑signed only     |
| Verbose diagnostics | `--verbose` verbose logging               |
| Color control       | `--no-color` disable ANSI                 |

### Usage

```bash
./scripts/get_yang_models.sh [OPTIONS]
```

### Flags

| Flag               | Description        | Default | Example               |
| ------------------ | ------------------ | ------- | --------------------- |
| `-c, --controller` | Controller host/IP | (env)   | `-c wnc1.example.com` |
| `-t, --token`      | Base64 auth token  | (env)   | `-t "dXNlcjpwYXNz"`   |
| `-p, --protocol`   | `http` or `https`  | `https` | `-p http`             |
| `-k, --insecure`   | Skip TLS verify    | `false` | `-k`                  |
| `-v, --verbose`    | Verbose logs       | `false` | `-v`                  |
| `--no-color`       | Disable color      | `false` | `--no-color`          |
| `-h, --help`       | Help text          | -       | `-h`                  |

### Examples

```bash
# Basic usage with environment variables
export WNC_CONTROLLER="wnc1.example.com"
export WNC_ACCESS_TOKEN="your-token-here"
./scripts/get_yang_models.sh -k

# Explicit controller and token
./scripts/get_yang_models.sh -c wnc1.example.com -t "dXNlcjpwYXNzd29yZA==" -k

# Using HTTP instead of HTTPS
./scripts/get_yang_models.sh -p http -c 192.168.1.100
```

<details>
<summary>Sample Output</summary>

```bash
$ ./scripts/get_yang_models.sh -k

Configuration:
=============
Protocol: https
Controller: wnc1.example.internal
Output Format: pretty

Available YANG Models (Cisco Wireless):
======================================
Cisco-IOS-XE-wireless-access-point-cfg-rpc/2023-07-01
Cisco-IOS-XE-wireless-access-point-cmd-rpc/2023-07-20
Cisco-IOS-XE-wireless-access-point-oper/2023-08-01
Cisco-IOS-XE-wireless-actions-rpc/2022-11-01
Cisco-IOS-XE-wireless-afc-cloud-oper/2023-07-10
Cisco-IOS-XE-wireless-afc-oper/2023-07-10
Cisco-IOS-XE-wireless-ap-cfg/2023-08-01
Cisco-IOS-XE-wireless-ap-types/2023-08-01
Cisco-IOS-XE-wireless-awips-oper/2023-08-01
Cisco-IOS-XE-wireless-ble-ltx-oper/2023-08-01
Cisco-IOS-XE-wireless-client-global-oper/2023-08-01
Cisco-IOS-XE-wireless-client-oper/2023-08-01
Cisco-IOS-XE-wireless-enum-types/2023-08-01
Cisco-IOS-XE-wireless-general-cfg/2023-08-01
Cisco-IOS-XE-wireless-general-oper/2023-08-01
<snip>

Operation completed successfully.
```

</details>

## 📖 `get_yang_model_details.sh` – Model Definition

Fetch full module definition (metadata + body) for a specific YANG module.

### Key Features

| Capability         | Notes                                          |
| ------------------ | ---------------------------------------------- |
| Format control     | `-f json\|xml` or `--raw` passthrough          |
| Revision selection | Provide specific revision or default to newest |
| Verbosity          | `-v` shows request metadata                    |
| TLS handling       | `-k` to skip verification (lab)                |

### Usage

```bash
./scripts/get_yang_model_details.sh [OPTIONS] <MODEL>
```

### Flags

| Flag               | Description        | Default | Example      |
| ------------------ | ------------------ | ------- | ------------ |
| `-c, --controller` | Controller host/IP | (env)   | `-c wnc1`    |
| `-t, --token`      | Base64 token       | (env)   | `-t YWRt...` |
| `-p, --protocol`   | Protocol           | `https` | `-p http`    |
| `-f, --format`     | Output format      | `json`  | `-f xml`     |
| `-r, --raw`        | Raw passthrough    | off     | `-r`         |
| `-k, --insecure`   | Skip TLS verify    | false   | `-k`         |
| `-v, --verbose`    | Verbose logs       | false   | `-v`         |
| `--no-color`       | Disable color      | false   | `--no-color` |
| `-h, --help`       | Help               | -       | `-h`         |

### Examples

```bash
# Get default access point operational model
./scripts/get_yang_model_details.sh -c wnc1.example.com Cisco-IOS-XE-wireless-access-point-oper -k

# Get specific model with custom revision
./scripts/get_yang_model_details.sh -c wnc1.example.com Cisco-IOS-XE-wireless-wlan-cfg -k

# Get raw YANG output for processing
./scripts/get_yang_model_details.sh -c wnc1.example.com -r -k > model.yang

# Verbose debugging mode
./scripts/get_yang_model_details.sh -c wnc1.example.com -v Cisco-IOS-XE-wireless-general-oper -k
```

<details>
<summary>Sample Output</summary>

```bash
$ ./scripts/get_yang_model_details.sh -c wnc1.example.internal -f pretty -k

Fetching YANG model details from: https://wnc1.example.internal/restconf/tailf/modules/Cisco-IOS-XE-wireless-access-point-oper/2023-08-01
Protocol: https
Controller: wnc1.example.internal
YANG Model: Cisco-IOS-XE-wireless-access-point-oper
Revision: 2023-08-01
Output Format: pretty
Insecure mode: --insecure
Verbose mode: false

YANG Model Details:
==================
Raw YANG Module Definition:
---------------------------
module Cisco-IOS-XE-wireless-access-point-oper {
  yang-version 1.1;
  namespace "http://cisco.com/ns/yang/Cisco-IOS-XE-wireless-access-point-oper";
  prefix wireless-access-point-oper;

  import Cisco-IOS-XE-event-history-types {
    prefix event-history-types;
  }
  import Cisco-IOS-XE-wireless-ap-types {
    prefix ap-types;
  }
  import Cisco-IOS-XE-wireless-enum-types {
    prefix wireless-enum-types;
  }
  <snip>

Operation completed successfully.
```

</details>

## 🔍 `get_yang_statement_details.sh` – Data Query

Fetches a live data subtree for a given model + statement (root node) via RESTCONF `data` path.

### Key Features

| Capability          | Notes                                         |
| ------------------- | --------------------------------------------- |
| Statement targeting | Provide `<MODEL> <STATEMENT>` positional args |
| Formats             | JSON(default) or XML (`-f xml`)               |
| Verbose             | `-v` for HTTP metadata                        |
| TLS                 | `-k` for lab/self-signed                      |

### Usage

```bash
./scripts/get_yang_statement_details.sh [OPTIONS] <MODEL> <STATEMENT>
```

### Flags

| Flag               | Description         | Default |
| ------------------ | ------------------- | ------- |
| `-c, --controller` | Controller host/IP  | (env)   |
| `-t, --token`      | Base64 token        | (env)   |
| `-p, --protocol`   | Protocol http/https | https   |
| `-f, --format`     | json or xml         | json    |
| `-k, --insecure`   | Skip TLS verify     | false   |
| `-v, --verbose`    | Verbose logs        | false   |
| `--no-color`       | Disable color       | false   |
| `-h, --help`       | Help                | -       |

### Examples

```bash
# Get access point operational data (default)
./scripts/get_yang_statement_details.sh -c wnc1.example.com Cisco-IOS-XE-wireless-access-point-oper access-point-oper-data -k

# Get client operational data
./scripts/get_yang_statement_details.sh -c wnc1.example.com Cisco-IOS-XE-wireless-client-oper client-oper-data -k

# Get JSON output for processing
./scripts/get_yang_statement_details.sh -c wnc1.example.com -f json Cisco-IOS-XE-wireless-general-oper general-oper-data -k > general.json

# Get general wireless operational status
./scripts/get_yang_statement_details.sh -c wnc1.example.com -v Cisco-IOS-XE-wireless-general-oper general-oper-data -k
```

<details>
<summary>Sample Output</summary>

```bash
$ ./scripts/get_yang_statement_details.sh -c wnc1.example.internal -f json -k

Fetching YANG statement details from: https://wnc1.example.internal/restconf/data/Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data
Protocol: https
Controller: wnc1.example.internal
YANG Model: Cisco-IOS-XE-wireless-access-point-oper
Identifier: access-point-oper-data
Output Format: json
Insecure mode: --insecure
Verbose mode: false

{
  "Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data": {
    "ap-radio-neighbor": [
      {
        "ap-mac": "28:ac:9e:bb:3c:80",
        "slot-id": 0,
        "bssid": "08:10:86:bf:07:e3",
        "ssid": "aterm-b5acbb-g",
        "rssi": -61,
        "channel": 0,
        "primary-channel": 4,
        "last-update-rcvd": "2025-06-25T14:35:59.306155+00:00"
      }
    ],
    "country-list": [
      {
        "country-code": "JP",
        "regulatory-domain": "JP"
      }
    ]
    <snip>
  }
}

Operation completed successfully.
```

</details>

## 📚 Appendix

<details>
<summary>Common Wireless YANG Modules</summary>

| Category               | Examples                                                              |
| ---------------------- | --------------------------------------------------------------------- |
| Operational (`-oper`)  | access-point-oper, client-oper, general-oper, rrm-oper, mobility-oper |
| Configuration (`-cfg`) | ap-cfg, wlan-cfg, rf-cfg, site-cfg, general-cfg                       |

</details>

<details>
<summary>Troubleshooting</summary>

| Problem                   | Resolution                                     |
| ------------------------- | ---------------------------------------------- |
| `curl: command not found` | Install: `brew install curl`                   |
| `jq: command not found`   | Install: `brew install jq`                     |
| Empty list                | Verify token & role privileges                 |
| TLS errors                | Use `-k` only in lab; fix cert chain otherwise |
| 404 model                 | Confirm spelling & revision availability       |

</details>

---

**Back to:** [API Reference](API_REFERENCE.md) | [Security](SECURITY.md)
