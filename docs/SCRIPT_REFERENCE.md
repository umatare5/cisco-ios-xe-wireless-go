# 📋 Scripts Reference

Helper scripts for YANG model discovery and inspection against a Catalyst 9800 controller.

| Feature | Summary |
|---------|---------|
| Model Discovery | List wireless YANG modules + revisions |
| Model Detail | Fetch raw module text |
| Statement Data | Retrieve operational subtree data |
| Output Styles | `pretty`, `json`, `raw` (where applicable) |
| Banners | Unified colored headers (auto disabled with `--no-color`) |
| Env Handling | Missing required vars → hard error (no silent skip) |

## 🗂️ Directory Structure

The `./scripts/` directory contains utility scripts for interacting with Cisco WNC RESTCONF APIs:

```text
scripts/
├── list_yang_models.sh           # Discover available YANG models
├── get_yang_model_details.sh     # Retrieve YANG model definitions
├── get_yang_statement_details.sh # Query operational data via YANG
└── lib/                          # Shared library functions
    ├── authentication.sh         # Authentication handling
    ├── common.sh                 # Common utilities and constants
    ├── dependencies.sh           # System dependency checking
    ├── file_utils.sh             # File management utilities
    ├── http_client.sh            # HTTP/HTTPS request handling
    ├── output_formatter.sh       # Output formatting functions
    └── validation.sh             # Input validation functions
```

## 🚀 Scripts Overview

| Script | Description |
|--------|-------------|
| `list_yang_models.sh` | Enumerate available YANG modules |
| `get_yang_model_details.sh` | Download module definition |
| `get_yang_statement_details.sh` | Fetch operational data subtree |

## 📋 `list_yang_models.sh` - YANG Model Discovery

Discovers and lists all available Cisco wireless YANG models from the WNC controller.

### Features

* Lists wireless YANG models + revisions (filtered to wireless namespaces)
* HTTPS default; optional `-k` to skip TLS verification (dev only)
* Hard fails if `WNC_CONTROLLER` or `WNC_ACCESS_TOKEN` missing (or flags not provided)
* Pretty or minimal output

### Usage

```bash
./scripts/list_yang_models.sh [OPTIONS]
```

### Flags

| Flag               | Description                | Example                     |
| ------------------ | -------------------------- | --------------------------- |
| `-c, --controller` | WNC controller hostname/IP | `-c wnc1.example.com`       |
| `-t, --token`      | Basic auth token           | `-t "dXNlcjpwYXNzd29yZA=="` |
| `-p, --protocol`   | Protocol (http/https)      | `-p https`                  |
| `-k, --insecure`   | Skip TLS verification      | `-k`                        |
| `-h, --help`       | Show help                  | `-h`                        |

### Common Usage Patterns

```bash
# Basic usage with environment variables
export WNC_CONTROLLER="wnc1.example.com"
export WNC_ACCESS_TOKEN="your-token-here"
./scripts/list_yang_models.sh -k

# Explicit controller and token
./scripts/list_yang_models.sh -c wnc1.example.com -t "dXNlcjpwYXNzd29yZA==" -k

# Using HTTP instead of HTTPS
./scripts/list_yang_models.sh -p http -c 192.168.1.100
```

<details>
<summary>Example of the result</summary>

```bash
$ ./scripts/list_yang_models.sh -k

Configuration:
=============
Protocol: https
Controller: <controller-hostname>
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

## 📖 `get_yang_model_details.sh` - YANG Model Definition Retrieval

Retrieves complete YANG model definitions including structure, types, and documentation.

### Features

* Fetch full module text at specific revision
* Output: `pretty` (annotated), `raw` (verbatim), `json` (wrapped)
* Validates model and revision format

### Usage

```bash
./scripts/get_yang_model_details.sh [OPTIONS]
```

### Flags

| Flag               | Description                 | Default                                   | Example               |
| ------------------ | --------------------------- | ----------------------------------------- | --------------------- |
| `-c, --controller` | WNC controller hostname/IP  | -                                         | `-c wnc1.example.com` |
| `-t, --token`      | Basic auth token            | -                                         | `-t "token..."`       |
| `-p, --protocol`   | Protocol (http/https)       | `https`                                   | `-p https`            |
| `-m, --model`      | YANG model name             | `Cisco-IOS-XE-wireless-access-point-oper` | `-m "model-name"`     |
| `-r, --revision`   | Model revision (YYYY-MM-DD) | `2023-08-01`                              | `-r "2023-03-01"`     |
| `-k, --insecure`   | Skip TLS verification       | -                                         | `-k`                  |
| `-v, --verbose`    | Verbose output              | -                                         | `-v`                  |
| `-f, --format`     | Output format               | `pretty`                                  | `-f json`             |
| `-h, --help`       | Show help                   | -                                         | `-h`                  |

### Common Usage Patterns

```bash
# Get default access point operational model
./scripts/get_yang_model_details.sh -c wnc1.example.com -k

# Get specific model with custom revision
./scripts/get_yang_model_details.sh -c wnc1.example.com \
  -m "Cisco-IOS-XE-wireless-wlan-cfg" -r "2023-03-01" -k

# Get raw YANG output for processing
./scripts/get_yang_model_details.sh -c wnc1.example.com \
  -f raw -k > model.yang

# Verbose debugging mode
./scripts/get_yang_model_details.sh -c wnc1.example.com -v -k
```

<details>
<summary>Example of the result</summary>

```bash
$ ./scripts/get_yang_model_details.sh -c <controller-hostname> -f pretty -k

Fetching YANG model details from: https://<controller-hostname>/restconf/tailf/modules/Cisco-IOS-XE-wireless-access-point-oper/2023-08-01
Protocol: https
Controller: <controller-hostname>
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

## 🔍 `get_yang_statement_details.sh` - Operational Data Querying

Queries real-time operational data from the WNC controller using YANG model paths.

### Features

* Calls RESTCONF data path for chosen model root identifier
* Output: `pretty` (with formatting) or `json`
* Useful for validating live controller state

### Usage

```bash
./scripts/get_yang_statement_details.sh [OPTIONS]
```

### Flags

| Flag               | Description                | Default                                   | Example                 |
| ------------------ | -------------------------- | ----------------------------------------- | ----------------------- |
| `-c, --controller` | WNC controller hostname/IP | -                                         | `-c wnc1.example.com`   |
| `-t, --token`      | Basic auth token           | -                                         | `-t "token..."`         |
| `-p, --protocol`   | Protocol (http/https)      | `https`                                   | `-p https`              |
| `-m, --model`      | YANG model name            | `Cisco-IOS-XE-wireless-access-point-oper` | `-m "model-name"`       |
| `-i, --id`         | YANG model identifier      | `access-point-oper-data`                  | `-i "client-oper-data"` |
| `-k, --insecure`   | Skip TLS verification      | -                                         | `-k`                    |
| `-v, --verbose`    | Verbose output             | -                                         | `-v`                    |
| `-f, --format`     | Output format              | `pretty`                                  | `-f json`               |
| `-h, --help`       | Show help                  | -                                         | `-h`                    |

### Common Usage Patterns

```bash
# Get access point operational data (default)
./scripts/get_yang_statement_details.sh -c wnc1.example.com -k

# Get client operational data
./scripts/get_yang_statement_details.sh -c wnc1.example.com \
  -m "Cisco-IOS-XE-wireless-client-oper" -i "client-oper-data" -k

# Get JSON output for processing
./scripts/get_yang_statement_details.sh -c wnc1.example.com \
  -f json -k > ap_data.json

# Get general wireless operational status
./scripts/get_yang_statement_details.sh -c wnc1.example.com \
  -m "Cisco-IOS-XE-wireless-general-oper" -i "general-oper-data" -v -k
```

<details>
<summary>Example of the result</summary>

```bash
$ ./scripts/get_yang_statement_details.sh -c <controller-hostname> -f json -k

Fetching YANG statement details from: https://<controller-hostname>/restconf/data/Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data
Protocol: https
Controller: <controller-hostname>
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

## ⚙️ Environment Variables

| Variable | Purpose | Required |
|----------|---------|----------|
| `WNC_CONTROLLER` | Controller host/IP | Yes |
| `WNC_ACCESS_TOKEN` | Base64 `user:pass` | Yes |

Missing any required variable (and no equivalent flag) → exit status 1.

## 📚️ Appendix

### 📚 Common YANG Models

#### Operational Data (`-oper`)

| Model                                     | Description                       |
| ----------------------------------------- | --------------------------------- |
| `Cisco-IOS-XE-wireless-access-point-oper` | Access point operational data     |
| `Cisco-IOS-XE-wireless-client-oper`       | Client operational data           |
| `Cisco-IOS-XE-wireless-general-oper`      | General wireless operational data |
| `Cisco-IOS-XE-wireless-mobility-oper`     | Mobility operational data         |
| `Cisco-IOS-XE-wireless-rrm-oper`          | Radio Resource Management data    |

#### Configuration Data (`-cfg`)

| Model                               | Description                    |
| ----------------------------------- | ------------------------------ |
| `Cisco-IOS-XE-wireless-wlan-cfg`    | WLAN configuration             |
| `Cisco-IOS-XE-wireless-ap-cfg`      | Access point configuration     |
| `Cisco-IOS-XE-wireless-site-cfg`    | Site configuration             |
| `Cisco-IOS-XE-wireless-rf-cfg`      | RF profile configuration       |
| `Cisco-IOS-XE-wireless-general-cfg` | General wireless configuration |

#### Model Revisions

| Revision     | Description             |
| ------------ | ----------------------- |
| `2022-11-01` | Earlier stable release  |
| `2023-03-01` | Spring 2023 features    |
| `2023-07-01` | Summer 2023 features    |
| `2023-08-01` | Latest stable (default) |

## 🔥 Troubleshooting

| Issue | Action |
|-------|--------|
| Missing env vars | Export `WNC_CONTROLLER` / `WNC_ACCESS_TOKEN` or pass flags |
| TLS errors | Confirm CA trust; dev only: add `-k` |
| Empty list | Verify controller version (must be IOS‑XE 17.12) |
| Invalid model | Use exact module name from discovery output |
| JSON parse issues | Use `-f raw` then post‑process |
