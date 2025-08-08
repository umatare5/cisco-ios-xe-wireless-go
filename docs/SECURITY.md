# 🔐 Security Considerations

This document outlines security best practices when using the `cisco-ios-xe-wireless-go` Go library for Cisco Catalyst 9800 Wireless Network Controllers.

## 🔒 TLS Certificate Verification

By default, the library enforces strict TLS certificate verification:

```go
// Secure connection (default)
config := wnc.Config{
    Controller:  "wnc.example.com",
    AccessToken: "YWRtaW46eW91ci1wYXNzd29yZA==",
    // InsecureSkipVerify: false (default)
}
client, err := wnc.NewClient(config)

// Skip verification only for development/testing
config := wnc.Config{
    Controller:         "wnc-dev.local",
    AccessToken:        "YWRtaW46ZGV2LXBhc3N3b3Jk",
    InsecureSkipVerify: true, // Only for development
}
client, err := wnc.NewClient(config)
```

> [!WARNING]
> The `InsecureSkipVerify: true` option disables TLS certificate verification and should only be used in trusted development environments with self-signed certificates. **Never use this option in production environments** as it compromises security.

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
