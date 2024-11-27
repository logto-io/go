package client

import (
	"github.com/logto-io/go/v2/core"
	"golang.org/x/exp/slices"
)

type LogtoConfig struct {
	Endpoint              string
	AppId                 string
	AppSecret             string
	Scopes                []string
	Resources             []string
	Prompt                string
	IncludeReservedScopes *bool
}

/**
 * Normalize the Logto client configuration per the following rules:
 *
 * - Add default scopes (`openid`, `offline_access` and `profile`) if not provided.
 * - Add `ReservedResource.Organization` to resources if `UserScope.Organizations` is included in scopes.
 */
func (logtoConfig *LogtoConfig) normalized() {
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
