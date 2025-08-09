# 🔐 Security

This document provides an overview of security practices for using this library.

## 🛡️ Checklist

This section lists essential security checks for pre‑deployment and ongoing review.

### ✅ Pre‑Deployment

- [ ] Enable TLS certificate verification and set `InsecureSkipVerify` to false. → See [TLS Verification](#tls-verification)
- [ ] Store authentication tokens in a secure credential manager. → See [Secure Storage](#token-storage)
- [ ] Ensure no credentials are hardcoded in source code. → See [Environment Variables](#token-env)
- [ ] Separate configurations for dev, staging, and production. → See [Environment Isolation](#environment-isolation)
- [ ] Configure logging to exclude secrets and use appropriate levels. → See [Logging](#logging)
- [ ] Restrict network access to only required endpoints. → See [Network & Access](#network-access)
- [ ] Use service accounts with minimal permissions. → See [Network & Access](#network-access)
- [ ] Apply timeouts to all API requests using contexts. → See [Context & Timeouts](#context-timeouts)

### 🔍 Periodic Review

- [ ] Rotate authentication tokens on a regular schedule. → See [Token Rotation](#token-rotation)
- [ ] Review API access logs monthly for anomalies. → See [Logging](#logging)
- [ ] Audit user permissions to maintain least‑privilege. → See [Network & Access](#network-access)
- [ ] Update dependencies and toolchain to current versions. → See [Token Handling](#token-handling)
- [ ] Monitor for upstream security advisories and CVEs. → See [References](#references)
- [ ] Test backup or fallback authentication mechanisms. → See [Token Handling](#token-handling)
- [ ] Validate firewall rules, ACLs, and rate limits are enforced. → See [Network & Access](#network-access)

## 🔒 TLS Verification <a id="tls-verification"></a>

Strict certificate validation is enforced unless you explicitly opt out via option.

```go
client, err := wnc.NewClient("wnc1.example.internal", token)

insecureClient, err := wnc.NewClient(
    "wnc1.example.internal", token,
    wnc.WithInsecureSkipVerify(true), // LAB ONLY
)
```

> [!CAUTION]
> The `wnc.WithInsecureSkipVerify(true)` option disables TLS certificate verification. This should only be used in development environments or when connecting to controllers with self-signed certificates. **Never use this option in production environments** as it compromises security.

## 🔑 Token Handling <a id="token-handling"></a>

Handle authentication tokens securely with isolated storage, periodic rotation, and no exposure in logs.

### ✅ Recommended

1. **Environment Variables**: Store tokens in environment variables, never in source code: <a id="token-env"></a>

   ```go
   import (
       "os"
       wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
   )

   client, err := wnc.NewClient(
       os.Getenv("WNC_CONTROLLER"),
       os.Getenv("WNC_ACCESS_TOKEN"),
       wnc.WithTimeout(30*time.Second),
   )
   ```

2. **Token Generation**: Use Base64 encoding for username:password combinations: <a id="token-generation"></a>

   ```bash
    # Generate token manually (ad-hoc only)
   echo -n "admin:your-secure-password" | base64
   # Output: YWRtaW46eW91ci1zZWN1cmUtcGFzc3dvcmQ=

    # Prefer central secret store, not ad-hoc scripts
   ```

3. **Token Rotation**: Regenerate tokens regularly and update environment variables: <a id="token-rotation"></a>

   ```bash
   # Automated token refresh script
   NEW_TOKEN=$(echo -n "admin:$NEW_PASSWORD" | base64)
   export WNC_ACCESS_TOKEN="$NEW_TOKEN"
   ```

4. **Secure Storage**: Use OS / Vault stores <a id="token-storage"></a>

   ```bash
   # Example with macOS Keychain
   PASSWORD=$(security find-generic-password -a admin -s wnc-password -w)
   TOKEN=$(echo -n "admin:$PASSWORD" | base64)
   export WNC_ACCESS_TOKEN="$TOKEN"

   # Example with HashiCorp Vault
   TOKEN=$(vault kv get -field=token secret/wnc/credentials)
   export WNC_ACCESS_TOKEN="$TOKEN"
   ```

5. **Context & Timeouts**: Always bound requests <a id="context-timeouts"></a>

   ```go
   ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
   defer cancel()

    apData, err := client.AP().GetOper(ctx)
   ```

### ❌ Avoid

Avoid these practices to reduce exposure, preserve accountability, and prevent secret leakage.

- Hardcoding tokens — Exposes credentials and prevents safe rotation.
- Committing `.env` with tokens — Leaks secrets through VCS and CI artifacts.
- Reusing prod tokens in dev or staging — Increases blast radius across environments.
- Logging Authorization headers — Risks credential disclosure in logs.
- Sharing tokens between individuals — Breaks accountability and auditability.

## 🌐 Network & Access <a id="network-access"></a>

This section defines network controls and access policies to protect the controller and data.

| Control       | Recommendation                                  |
| ------------- | ----------------------------------------------- |
| Transport     | Use HTTPS for all requests.                     |
| Port          | Expose only port 443 for RESTCONF.              |
| Segmentation  | Limit controller access to a mgmt VLAN or VPN.  |
| Accounts      | Use least‑privilege service accounts.           |
| Rate limiting | Apply rate limits on the controller or a proxy. |
| Auditing      | Review authentication logs regularly.           |

## 📝 Logging <a id="logging"></a>

Prefer structured logs, exclude secrets, and log only necessary context for operations and audits.

```go
import (
   "log/slog"
   "os"
   "time"
   wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
)

logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
   Level: slog.LevelInfo,
}))

client, err := wnc.NewClient(
   "wnc1.example.internal",
   os.Getenv("WNC_ACCESS_TOKEN"),
   wnc.WithLogger(logger),
   wnc.WithTimeout(30*time.Second),
)
```

## 🏭 Environment Isolation <a id="environment-isolation"></a>

Use separate clients and credentials for dev, staging, and prod to limit blast radius and enforce tailored timeouts.

```go
dev, _ := wnc.NewClient("wnc1.example.internal", os.Getenv("WNC_DEV_TOKEN"), wnc.WithInsecureSkipVerify(true), wnc.WithTimeout(5*time.Second))
staging, _ := wnc.NewClient("wnc1.example.internal", os.Getenv("WNC_STAGING_TOKEN"), wnc.WithTimeout(15*time.Second))
prod, _ := wnc.NewClient("wnc1.example.internal", os.Getenv("WNC_PROD_TOKEN"), wnc.WithTimeout(30*time.Second))
_, _, _ = dev, staging, prod
```

### 📊 Monitoring

Monitor authentication, request volume, latency, and TLS signals to detect issues early and guide response.

| Area    | Metric / Signal                                                 |
| ------- | --------------------------------------------------------------- |
| Auth    | Track failed versus successful authentications.                 |
| Volume  | Monitor request volume per service such as AP, Client, and RRM. |
| Latency | Watch the 95th percentile request duration.                     |
| TLS     | Alert on TLS handshake failures.                                |

### 🔧 Error Handling

Log actionable context for operators, return generic messages to users, and prevent any secret leakage.

```go
apData, err := client.AP().GetOper(ctx)
if err != nil {
    // Log detailed errors securely (not to end users)
    logger.Error("API request failed", "error", err, "endpoint", "ap-oper")

    // Return generic error to end users
    return nil, fmt.Errorf("failed to retrieve access point data")
}
```

## 📖 References <a id="references"></a>

- [Go Security Best Practices](https://go.dev/security/)
- [RESTCONF Security Best Practices](https://tools.ietf.org/html/rfc8040#section-2.5)
