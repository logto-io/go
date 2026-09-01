<p align="center">
  <a href="https://logto.io" target="_blank" align="center" alt="Logto Logo">
    <picture>
      <source width="200" media="(prefers-color-scheme: dark)" srcset="https://github.com/logto-io/.github/raw/master/profile/logto-lockup-brand-on-dark.svg">
      <source width="200" media="(prefers-color-scheme: light)" srcset="https://github.com/logto-io/.github/raw/master/profile/logto-lockup-brand-on-light.svg">
      <img width="200" src="https://github.com/logto-io/.github/raw/master/profile/logto-lockup-brand-on-light.svg" alt="Logto logo">
    </picture>
  </a>
</p>

# Logto Go SDKs
[![Build Status](https://github.com/logto-io/go/actions/workflows/main.yml/badge.svg)](https://github.com/logto-io/go/actions/workflows/main.yml)
[![Codecov](https://img.shields.io/codecov/c/github/logto-io/go)](https://app.codecov.io/gh/logto-io/go?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/logto-io/go)](https://goreportcard.com/report/github.com/logto-io/go)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/logto-io/go)](https://pkg.go.dev/github.com/logto-io/go)

The repo for SDKs and working samples written in Go.

Check out the [Go SDK tutorial](https://docs.logto.io/quick-starts/go) or [Go reference](https://pkg.go.dev/github.com/logto-io/go/v2) for more information.

## Installation

To install Logto Go SDK, use `go get`.

For core package:

```bash
go get github.com/logto-io/go/v2/core
```

For client package:

```bash
go get github.com/logto-io/go/v2/client
```

To update Logto Go SDK to the latest version, use:
```bash
go get -u github.com/logto-io/go/v2/core
go get -u github.com/logto-io/go/v2/client
```

## Packages

| Name   | Description                          |
| ------ | ------------------------------------ |
| core   | Logto SDK core package               |
| client | Logto client built upon the `core` package |

## User info and ID token claims

`client.FetchUserInfo` and `client.GetIdTokenClaims` return `core.UserInfoResponse` and `core.IdTokenClaims` respectively. Besides basic claims such as `sub`, `name`, `email`, both types model the [standard claims](https://openid.net/specs/openid-connect-core-1_0.html#StandardClaims) included in the `profile` scope: `family_name`, `given_name`, `middle_name`, `nickname`, `preferred_username`, `profile`, `website`, `gender`, `birthdate`, `zoneinfo`, `locale`, as well as `created_at` and `updated_at`. Logto only returns these standard claims when their values are not empty, so absent claims are left as zero values.

To access claims that are not modeled as struct fields, e.g. custom claims, use the `GetClaim` method:

```go
userInfo, err := logtoClient.FetchUserInfo()
if err != nil {
	// Handle error
}

if value, ok := userInfo.GetClaim("custom_claim"); ok {
	// Use the claim value
}

idTokenClaims, err := logtoClient.GetIdTokenClaims()
if err != nil {
	// Handle error
}

if value, ok := idTokenClaims.GetClaim("custom_claim"); ok {
	// Use the claim value
}
```

## Error handling

SDK functions return structured errors that can be inspected with `errors.Is` and `errors.As` to turn failures into user actions.

The most common scenario: the stored refresh token has expired or been revoked, so the session cannot be renewed. `LogtoClient` methods (e.g. `GetAccessToken`, `FetchUserInfo`) report this as `client.ErrNotAuthenticated`, the same error returned when the user has never signed in:

```go
import (
	"errors"

	"github.com/logto-io/go/v2/client"
)

userInfo, err := logtoClient.FetchUserInfo()
if err != nil {
	if errors.Is(err, client.ErrNotAuthenticated) {
		// No valid session, redirect the user to sign in again.
	}
	// Handle other errors.
}
```

For everything else, any non-200 response from the Logto server is a `*core.ResponseError` carrying the HTTP status code, the parsed OIDC and Logto error fields, and the raw body:

```go
import "github.com/logto-io/go/v2/core"

var responseError *core.ResponseError
if errors.As(err, &responseError) {
	log.Printf(
		"logto request failed: status=%d, code=%s, description=%s",
		responseError.StatusCode,
		responseError.Code,
		responseError.ErrorDescription,
	)
}
```

## Running behind a reverse proxy

By default, `HandleSignInCallback` reconstructs the callback URI from the incoming request, inferring the scheme from the TLS state (or the `X-Forwarded-Proto` header) and the host from the `Host` header.

If your application runs behind reverse proxies or multi-layer gateways (e.g., Cloudflare -> ALB -> Nginx, or Firebase Hosting -> Cloud Run), the scheme and host seen by your application may differ from the public address, causing sign-in callbacks to fail with a "callback uri not match redirect uri" error. Trusting these headers also exposes the application to Host header injection.

To avoid this, set the optional `BaseUrl` in `LogtoConfig` to the public base URL of your application. It must be an absolute URL with a scheme and may include a path prefix:

```go
logtoConfig := &client.LogtoConfig{
	Endpoint:  "<your-logto-endpoint>",
	AppId:     "<your-application-id>",
	AppSecret: "<your-application-secret>",
	// The public base URL of your application
	BaseUrl:   "https://my-app.com",
}
```

When `BaseUrl` is set, the callback URI is constructed by directly appending the incoming request URI (path and query) to it, and the request's `Host` and `X-Forwarded-Proto` headers are no longer used. Make sure the concatenated result matches the public callback URL registered as the redirect URI — for example, if a proxy strips a path prefix before forwarding requests to your application, include that prefix in `BaseUrl`.

## Resources

[![Website](https://img.shields.io/badge/website-logto.io-8262F8.svg)](https://logto.io/)
[![Docs](https://img.shields.io/badge/docs-logto.io-green.svg)](https://docs.logto.io/)
[![Discord](https://img.shields.io/discord/965845662535147551?logo=discord&logoColor=ffffff&color=7389D8&cacheSeconds=600)](https://discord.gg/UEPaF3j5e6)
