# 📋 Scripts Reference

Scripts for collecting and analyzing YANG models from Cisco Wireless Network Controller.

- **Comprehensive YANG Model Operations**: Interact with the Cisco WNC RESTCONF API to discover available YANG models, retrieve their full definitions, and query real-time operational data using simple shell scripts.
- **Flexible Command-Line Interface**: Supports configuration via both command-line flags and environment variables, with options for various output formats to suit different workflows.
- **Modular and Robust Scripts**: Built upon a shared library for common functions like authentication, HTTP requests, and validation, ensuring the scripts are reliable and easy to maintain.

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

| Name                            | Desciription                         |
| ------------------------------- | ------------------------------------ |
| `list_yang_models.sh`           | Discover available YANG models.      |
| `get_yang_model_details.sh`     | Retrieve complete model definitions. |
| `get_yang_statement_details.sh` | Query operational data via YANG.     |

## 📋 `list_yang_models.sh` - YANG Model Discovery

Discovers and lists all available Cisco wireless YANG models from the WNC controller.

### Features

- Lists all Cisco wireless YANG models with their revisions
- Supports both HTTP and HTTPS protocols
- Pretty-formatted output with clear categorization
- Environment variable support for credentials
- TLS certificate verification bypass option

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

## 📖 `get_yang_model_details.sh` - YANG Model Definition Retrieval

Retrieves complete YANG model definitions including structure, types, and documentation.

### Features

- Fetches complete YANG module definitions
- Multiple output formats (pretty, json, raw)
- Support for specific model revisions
- Verbose debugging mode
- Input validation for model names and revisions

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

## 🔍 `get_yang_statement_details.sh` - Operational Data Querying

Queries real-time operational data from the WNC controller using YANG model paths.

### Features

- Retrieves live operational data from WNC
- JSON and pretty-formatted output
- Configurable YANG model and identifier
- Real-time wireless network status
- Support for all operational YANG models

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

| Problem                   | Solution                                             |
| ------------------------- | ---------------------------------------------------- |
| `curl: command not found` | Install curl: `brew install curl`                    |
| `jq: command not found`   | Install jq: `brew install jq`                        |
| "Failed to fetch data"    | Check controller hostname, network, auth token       |
| TLS certificate errors    | Use `-k` flag to skip verification                   |
| "Invalid YANG model"      | Ensure format: `Cisco-IOS-XE-wireless-*-(oper\|cfg)` |
