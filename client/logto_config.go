package client

import (
	"slices"
	"strings"

	"github.com/logto-io/go/v2/core"
)

type LogtoConfig struct {
	Endpoint              string
	AppId                 string
	AppSecret             string
	Scopes                []string
	Resources             []string
	Prompt                string
	IncludeReservedScopes *bool
	// BaseUrl is the public base URL of your application, e.g. "https://my-app.com".
	// It must be an absolute URL with a scheme and may include a path prefix.
	//
	// When set, `HandleSignInCallback` constructs the callback URI from BaseUrl and
	// the incoming request path, instead of inferring the scheme and host from the
	// request (`Host` and `X-Forwarded-Proto` headers). Set it when your application
	// runs behind reverse proxies or multi-layer gateways, where the inferred values
	// may not match the public address and the headers cannot be trusted.
	//
	// When empty, the callback URI is inferred from the incoming request as before.
	BaseUrl string
}

/**
 * Normalize the Logto client configuration per the following rules:
 *
 * - Add default scopes (`openid`, `offline_access` and `profile`) if not provided.
 * - Add `ReservedResource.Organization` to resources if `UserScope.Organizations` is included in scopes.
 * - Trim the trailing slash of `BaseUrl` so it can be safely concatenated with a request path.
 */
func (logtoConfig *LogtoConfig) normalized() {
	logtoConfig.BaseUrl = strings.TrimSuffix(logtoConfig.BaseUrl, "/")

	includeReservedScopes := logtoConfig.IncludeReservedScopes
	if includeReservedScopes == nil || *includeReservedScopes {
		for _, defaultScope := range core.DefaultScopes {
			logtoConfig.Scopes = core.AppendIfNotExisted(logtoConfig.Scopes, defaultScope)
		}
	}

	if slices.Contains(logtoConfig.Scopes, core.UserScopeOrganizations) {
		logtoConfig.Resources = core.AppendIfNotExisted(logtoConfig.Resources, core.ReservedResourceOrganization)
	}
}
