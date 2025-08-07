# 🔐 Security Guide

This document outlines security best practices for the `cisco-ios-xe-wireless-go` Go library when connecting to Cisco Catalyst 9800 Wireless Network Controllers.

> [!WARNING]
> This guide covers security-critical configurations. Always follow your organization's security policies and compliance requirements.

## 🛡️ Core Security Features

### 🔒 TLS Certificate Verification

> [!CAUTION]
> The library enforces strict TLS certificate verification by default. Only disable in controlled development environments.

```go
import wnc "github.com/umatare5/cisco-ios-xe-wireless-go"

// ✅ Secure connection (default)
client, err := wnc.NewClient("wnc.example.com", "YWRtaW46cGFzc3dvcmQ=")

// ⚠️ Skip verification only for development/testing
client, err := wnc.NewClient("wnc-dev.local", "dGVzdDp0ZXN0",
    wnc.WithInsecureSkipVerify(true))
```

> [!WARNING]
> Use `WithInsecureSkipVerify(true)` **only in development environments** with self-signed certificates. **Never in production.**

### 🔑 Basic Authentication

The library uses HTTP Basic Authentication with Base64-encoded credentials:

```go
// Generate authentication token
// echo -n "admin:your-secure-password" | base64
token := "YWRtaW46eW91ci1zZWN1cmUtcGFzc3dvcmQ="

client, err := wnc.NewClient("wnc.example.com", token)
```

## 🏭 Production Security

### 🌍 Environment Variables

Store credentials in environment variables, never in source code:

```go
import (
    "os"
    wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
)

client, err := wnc.NewClient(
    os.Getenv("WNC_CONTROLLER"),
    os.Getenv("WNC_ACCESS_TOKEN"),
)
```

Set environment variables securely:

```bash
# Recommended environment variable names
export WNC_CONTROLLER="wnc-prod.company.com"
export WNC_ACCESS_TOKEN="$(echo -n 'admin:secure-password' | base64)"
```

### ⏱️ Request Timeouts

Always use context with timeouts to prevent hanging requests:

```go
import (
    "context"
    "time"
)

// Configure timeout at client level
client, err := wnc.NewClient("wnc.example.com", token,
    wnc.WithTimeout(30*time.Second))

// Use context for individual requests
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

apData, err := client.AP().Oper(ctx)
```

### 📊 Structured Logging

Configure secure logging with proper log levels:

```go
import (
    "log/slog"
    "os"
)

// Production logging configuration
logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
    Level: slog.LevelInfo, // Avoid debug level in production
    AddSource: false,      // Disable source info in production
}))

client, err := wnc.NewClient("wnc.example.com", token,
    wnc.WithLogger(logger))
```

## 🔐 Credential Management

### ✅ Best Practices

1. **Secure Credential Stores**:

   ```bash
   # macOS Keychain
   security add-generic-password -a admin -s wnc-prod -w "secure-password"
   PASSWORD=$(security find-generic-password -a admin -s wnc-prod -w)
   TOKEN=$(echo -n "admin:$PASSWORD" | base64)

   # HashiCorp Vault
   vault kv put secret/wnc/prod username=admin password=secure-password
   TOKEN=$(vault kv get -field=password secret/wnc/prod |
           awk -v user="$(vault kv get -field=username secret/wnc/prod)"
           'BEGIN{printf "%s:%s", user, $0}' | base64)
   ```

2. **Environment-Specific Tokens**:

   ```go
   // Development
   devClient, err := wnc.NewClient(
       "wnc-dev.local",
       os.Getenv("WNC_DEV_TOKEN"),
       wnc.WithInsecureSkipVerify(true), // Only acceptable in dev
       wnc.WithTimeout(5*time.Second))

   // Production
   prodClient, err := wnc.NewClient(
       "wnc-prod.company.com",
       os.Getenv("WNC_PROD_TOKEN"),
       wnc.WithTimeout(30*time.Second))
   ```

3. **Token Rotation**:

   ```bash
   # Automated token refresh script
   #!/bin/bash
   NEW_PASSWORD=$(generate-secure-password)
   NEW_TOKEN=$(echo -n "admin:$NEW_PASSWORD" | base64)

   # Update WNC password first
   update-wnc-password "$NEW_PASSWORD"

   # Update environment variable
   export WNC_ACCESS_TOKEN="$NEW_TOKEN"
   ```

### ❌ Security Anti-Patterns

- ❌ Never hardcode tokens in source code
- ❌ Don't store tokens in version control
- ❌ Avoid logging authentication tokens
- ❌ Don't share tokens between environments
- ❌ Never commit configuration files with tokens

## 🌐 Network Security

### 🔥 Firewall Configuration

Configure network access controls:

```bash
# Required ports and protocols
# HTTPS (RESTCONF API): TCP/443
# Controller management interface only

# Example firewall rules (iptables)
iptables -A OUTPUT -p tcp --dport 443 -d WNC_IP -j ACCEPT
iptables -A OUTPUT -p tcp --dport 443 -j DROP
```

### 🛡️ TLS Configuration

The library enforces strong TLS settings:

```go
// TLS transport configuration (internal)
transport := &http.Transport{
    TLSClientConfig: &tls.Config{
        InsecureSkipVerify: false, // Default: strict verification
        MinVersion:         tls.VersionTLS12,
    },
    TLSHandshakeTimeout: 10 * time.Second,
}
```

### 🚨 Error Handling

Implement secure error handling that doesn't leak sensitive information:

```go
import (
    "fmt"
    "log/slog"
)

apData, err := client.AP().Oper(ctx)
if err != nil {
    // ✅ Log detailed errors securely (server-side only)
    logger.Error("API request failed",
        "operation", "ap-oper",
        "error", err.Error())

    // ✅ Return generic error to clients
    return nil, fmt.Errorf("failed to retrieve access point data")
}
```

## 📋 Security Checklist

### 🛡️ Deployment Checklist

- [ ] TLS verification enabled (`InsecureSkipVerify: false`)
- [ ] Authentication tokens stored in secure credential management
- [ ] No hardcoded credentials in source code or configuration files
- [ ] Environment-specific configurations properly separated
- [ ] Logging configured with appropriate security levels
- [ ] Network access restricted to necessary endpoints only
- [ ] Service accounts configured with minimal required privileges
- [ ] Context timeouts configured for all API operations
- [ ] Error handling prevents information leakage

### 🔍 Regular Security Maintenance

- [ ] Rotate authentication tokens quarterly
- [ ] Review API access logs monthly
- [ ] Audit user permissions quarterly
- [ ] Update library dependencies regularly
- [ ] Monitor security advisories for Cisco IOS-XE
- [ ] Test backup authentication mechanisms
- [ ] Validate network security controls annually

## 🚨 Incident Response

### 🔑 Authentication Compromise

1. **Immediate Actions**:
   - Revoke compromised tokens on WNC
   - Generate new authentication credentials
   - Update credential stores and environment variables
   - Restart affected applications with new tokens

2. **Investigation**:
   - Review WNC access logs for suspicious activity
   - Check for unauthorized configuration changes
   - Validate API access patterns and usage

### 🌐 Network Security Incident

1. **Containment**:
   - Isolate affected controllers from network if necessary
   - Review firewall rules and network segmentation
   - Monitor for lateral movement attempts

2. **Recovery**:
   - Validate controller configurations
   - Update network security policies
   - Implement additional monitoring as needed

## 📚 Security Resources

- [Cisco Catalyst 9800 Security Configuration Guide](https://www.cisco.com/c/en/us/support/wireless/catalyst-9800-series-wireless-controllers/products-installation-and-configuration-guides-list.html)
- [RESTCONF Security Considerations (RFC 8040)](https://tools.ietf.org/html/rfc8040#section-2.5)
- [Go Security Best Practices](https://go.dev/security/)
- [OWASP API Security Top 10](https://owasp.org/www-project-api-security/)

---

**See Also:** [API Reference](API_REFERENCE.md) | [Testing Guide](TESTING.md) | [Main README](../README.md)

## 🔑 Authentication Token Security

### ✅ Best Practices

1. **Environment Variables**: Store tokens in environment variables, never in source code:

   ```go
   import (
       "os"
       wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
   )

   config := wnc.Config{
       Controller:  os.Getenv("WNC_CONTROLLER"),
       AccessToken: os.Getenv("WNC_ACCESS_TOKEN"),
   }
   ```

2. **Token Generation**: Use Base64 encoding for username:password combinations:

   ```bash
   # Generate token manually (not recommended for automation)
   echo -n "admin:your-secure-password" | base64
   # Output: YWRtaW46eW91ci1zZWN1cmUtcGFzc3dvcmQ=

   # Better: Use the wnc CLI tool for token generation
   wnc generate token -u admin -p "$SECURE_PASSWORD"
   ```

3. **Token Rotation**: Regenerate tokens regularly and update environment variables:

   ```bash
   # Automated token refresh script
   NEW_TOKEN=$(echo -n "admin:$NEW_PASSWORD" | base64)
   export WNC_ACCESS_TOKEN="$NEW_TOKEN"
   ```

4. **Secure Credential Management**: Use secure credential stores:

   ```bash
   # Example with macOS Keychain
   PASSWORD=$(security find-generic-password -a admin -s wnc-password -w)
   TOKEN=$(echo -n "admin:$PASSWORD" | base64)
   export WNC_ACCESS_TOKEN="$TOKEN"

   # Example with HashiCorp Vault
   TOKEN=$(vault kv get -field=token secret/wnc/credentials)
   export WNC_ACCESS_TOKEN="$TOKEN"
   ```

5. **Context-Aware Requests**: Always use context with timeouts:

   ```go
   ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
   defer cancel()

   apData, err := client.GetApOper(ctx)
   ```

### ❌ What to Avoid

- Never hardcode authentication tokens in source code
- Don't store tokens in version control systems
- Avoid logging authentication tokens in application logs
- Don't share tokens between environments (dev/staging/prod)
- Never commit configuration files containing tokens

## 🌐 Network Security

### 🔥 Firewall Considerations

- **HTTPS Traffic**: Port 443 (default for RESTCONF API)
- **Outbound Connections**: From client application to controller management interfaces
- **Network Segmentation**: Consider VPN or dedicated management networks for production
- **Controller Access**: Ensure controllers are on secure, isolated management networks

### 🚪 Access Control

- **Dedicated Service Accounts**: Use service accounts with minimal required privileges
- **Read-Only Access**: Implement read-only access where possible (this library currently supports GET operations only)
- **API Rate Limiting**: Monitor and implement rate limiting on controller side
- **Regular Audits**: Audit user permissions and API access patterns

### 📝 Logging and Monitoring

Configure structured logging for security monitoring:

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

config := wnc.Config{
    Controller:  "wnc.example.com",
    AccessToken: os.Getenv("WNC_ACCESS_TOKEN"),
    Logger:      logger,
}
```

## 🏭 Production Deployment

### 🏗️ Environment Isolation

```go
// Development environment
devConfig := wnc.Config{
    Controller:         "wnc-dev.local",
    AccessToken:        os.Getenv("WNC_DEV_TOKEN"),
    InsecureSkipVerify: true, // Only acceptable in dev
    Timeout:            5 * time.Second,
}

// Staging environment
stagingConfig := wnc.Config{
    Controller:  "wnc-staging.company.com",
    AccessToken: os.Getenv("WNC_STAGING_TOKEN"),
    Timeout:     15 * time.Second,
}

// Production environment
prodConfig := wnc.Config{
    Controller:  "wnc-prod.company.com",
    AccessToken: os.Getenv("WNC_PROD_TOKEN"),
    Timeout:     30 * time.Second,
    // InsecureSkipVerify: false (never set to true in production)
}
```

### 📊 Security Monitoring

1. **Request Monitoring**: Track API call patterns and volumes
2. **Authentication Failures**: Monitor and alert on authentication errors
3. **Network Anomalies**: Watch for unusual traffic patterns
4. **Configuration Changes**: Log all configuration modifications

### 🔧 Error Handling

Implement secure error handling that doesn't leak sensitive information:

```go
apData, err := client.GetApOper(ctx)
if err != nil {
    // Log detailed errors securely (not to end users)
    logger.Error("API request failed", "error", err, "endpoint", "ap-oper")

    // Return generic error to end users
    return nil, fmt.Errorf("failed to retrieve access point data")
}
```

## 🛡️ Security Checklist

### ✅ Pre-Deployment Checklist

- [ ] TLS certificate verification enabled (`InsecureSkipVerify: false`)
- [ ] Authentication tokens stored in secure credential management
- [ ] No hardcoded credentials in source code
- [ ] Environment-specific configurations separated
- [ ] Logging configured with appropriate security levels
- [ ] Network access restricted to necessary endpoints
- [ ] Service accounts configured with minimal privileges
- [ ] Context timeouts configured for all API calls

### 🔍 Regular Security Reviews

- [ ] Rotate authentication tokens quarterly
- [ ] Review API access logs monthly
- [ ] Audit user permissions quarterly
- [ ] Update dependency versions regularly
- [ ] Monitor for security advisories
- [ ] Test backup authentication mechanisms
- [ ] Validate network security controls

## 🚨 Incident Response

### Authentication Compromise

1. **Immediate Actions**:

   - Revoke compromised tokens on controller
   - Generate new authentication credentials
   - Update environment variables/credential stores
   - Restart affected applications

2. **Investigation**:
   - Review API access logs for suspicious activity
   - Check for unauthorized configuration changes
   - Validate network access patterns

### Network Security Breach

1. **Immediate Actions**:

   - Isolate affected controllers from network
   - Review firewall rules and network segmentation
   - Check for lateral movement attempts

2. **Recovery**:
   - Validate controller configurations
   - Update network security policies
   - Implement additional monitoring

## 📖 Additional Resources

- [Cisco Catalyst 9800 Security Configuration Guide](https://www.cisco.com/c/en/us/support/wireless/catalyst-9800-series-wireless-controllers/products-installation-and-configuration-guides-list.html)
- [RESTCONF Security Best Practices](https://tools.ietf.org/html/rfc8040#section-2.5)
- [Go Security Best Practices](https://go.dev/security/)

---

**Back to:** [API Reference](API_REFERENCE.md) | [Main README](../README.md)
