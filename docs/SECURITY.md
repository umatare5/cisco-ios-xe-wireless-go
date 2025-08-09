# 🔐 Security

This document provides an overview of security practices for using this library.

## 🔒 TLS Verification

Strict certificate validation is enforced unless you explicitly opt out via option.

```go
client, err := wnc.NewClient("wnc.example.com", token)

insecureClient, err := wnc.NewClient(
    "wnc-dev.local", token,
    wnc.WithInsecureSkipVerify(true), // LAB ONLY
)
```

> [!CAUTION]
> The `wnc.WithInsecureSkipVerify(true)` option disables TLS certificate verification. This should only be used in development environments or when connecting to controllers with self-signed certificates. **Never use this option in production environments** as it compromises security.

## 🔑 Token Handling

### ✅ Recommended

1. **Environment Variables**: Store tokens in environment variables, never in source code:

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

2. **Token Generation**: Use Base64 encoding for username:password combinations:

   ```bash
    # Generate token manually (ad-hoc only)
   echo -n "admin:your-secure-password" | base64
   # Output: YWRtaW46eW91ci1zZWN1cmUtcGFzc3dvcmQ=

    # Prefer central secret store, not ad-hoc scripts
   ```

3. **Token Rotation**: Regenerate tokens regularly and update environment variables:

   ```bash
   # Automated token refresh script
   NEW_TOKEN=$(echo -n "admin:$NEW_PASSWORD" | base64)
   export WNC_ACCESS_TOKEN="$NEW_TOKEN"
   ```

4. **Secure Storage**: Use OS / Vault stores

   ```bash
   # Example with macOS Keychain
   PASSWORD=$(security find-generic-password -a admin -s wnc-password -w)
   TOKEN=$(echo -n "admin:$PASSWORD" | base64)
   export WNC_ACCESS_TOKEN="$TOKEN"

   # Example with HashiCorp Vault
   TOKEN=$(vault kv get -field=token secret/wnc/credentials)
   export WNC_ACCESS_TOKEN="$TOKEN"
   ```

5. **Context & Timeouts**: Always bound requests

   ```go
   ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
   defer cancel()

    apData, err := client.AP().GetOper(ctx)
   ```

### ❌ Avoid

- Hardcoding tokens
- Committing `.env` with tokens
- Reusing prod tokens in dev / staging
- Logging Authorization headers
- Sharing tokens between individuals

## 🌐 Network & Access

| Control       | Recommendation                         |
| ------------- | -------------------------------------- |
| Transport     | HTTPS only (default)                   |
| Port          | 443 (RESTCONF)                         |
| Segmentation  | Restrict controller to mgmt VLAN / VPN |
| Accounts      | Least privilege service accounts       |
| Rate limiting | Enforce on controller / upstream proxy |
| Auditing      | Periodic review of auth logs           |

### 📝 Logging

```go
import (
    "log/slog"
    "os"
)

// Security-focused logging configuration
logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
    Level: slog.LevelInfo, // Avoid debug level in production
    AddSource: true,
}))

client, err := wnc.NewClient(
    os.Getenv("WNC_CONTROLLER"),
    os.Getenv("WNC_ACCESS_TOKEN"),
    wnc.WithLogger(logger),
)
```

## 🏭 Environment Isolation

```go
dev, _ := wnc.NewClient("wnc-dev.local", os.Getenv("WNC_DEV_TOKEN"), wnc.WithInsecureSkipVerify(true), wnc.WithTimeout(5*time.Second))
staging, _ := wnc.NewClient("wnc-staging.local", os.Getenv("WNC_STAGING_TOKEN"), wnc.WithTimeout(15*time.Second))
prod, _ := wnc.NewClient("wnc-prod.local", os.Getenv("WNC_PROD_TOKEN"), wnc.WithTimeout(30*time.Second))
_, _, _ = dev, staging, prod
```

### 📊 Monitoring Focus

| Area    | Metric / Signal                        |
| ------- | -------------------------------------- |
| Auth    | Failed vs success ratio                |
| Volume  | Requests per service (AP, Client, RRM) |
| Latency | P95 request duration                   |
| TLS     | Handshake failures                     |

### 🔧 Error Handling

```go
apData, err := client.AP().GetOper(ctx)
if err != nil {
    // Log detailed errors securely (not to end users)
    logger.Error("API request failed", "error", err, "endpoint", "ap-oper")

    // Return generic error to end users
    return nil, fmt.Errorf("failed to retrieve access point data")
}
```

## 🛡️ Checklist

### ✅ Pre‑Deployment

- [ ] TLS certificate verification enabled (`InsecureSkipVerify: false`)
- [ ] Authentication tokens stored in secure credential management
- [ ] No hardcoded credentials in source code
- [ ] Environment-specific configurations separated
- [ ] Logging configured with appropriate security levels
- [ ] Network access restricted to necessary endpoints
- [ ] Service accounts configured with minimal privileges
- [ ] Context timeouts configured for all API calls

### 🔍 Regular Review

- [ ] Rotate authentication tokens quarterly
- [ ] Review API access logs monthly
- [ ] Audit user permissions quarterly
- [ ] Update dependency versions regularly
- [ ] Monitor for security advisories
- [ ] Test backup authentication mechanisms
- [ ] Validate network security controls

## 🚨 Incident Response

### Auth Compromise

1. **Immediate Actions**:

   - Revoke compromised tokens on controller
   - Generate new authentication credentials
   - Update environment variables/credential stores
   - Restart affected applications

2. **Investigate**: Review logs, correlate IPs, timeline events

### Network Breach

1. **Immediate Actions**:

   - Isolate affected controllers from network
   - Review firewall rules and network segmentation
   - Check for lateral movement attempts

2. **Recovery**: Re-issue certs, validate ACLs, enhance detection

## 📖 References

- [Go Security Best Practices](https://go.dev/security/)
- [RESTCONF Security Best Practices](https://tools.ietf.org/html/rfc8040#section-2.5)
