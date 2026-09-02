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

## Resources

[![Website](https://img.shields.io/badge/website-logto.io-8262F8.svg)](https://logto.io/)
[![Docs](https://img.shields.io/badge/docs-logto.io-green.svg)](https://docs.logto.io/)
[![Discord](https://img.shields.io/discord/965845662535147551?logo=discord&logoColor=ffffff&color=7389D8&cacheSeconds=600)](https://discord.gg/UEPaF3j5e6)
