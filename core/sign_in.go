package core

import (
	"net/url"
	"slices"
	"strings"
)

type SignInUriGenerationOptions struct {
	AuthorizationEndpoint string
	ClientId              string
	RedirectUri           string
	CodeChallenge         string
	State                 string
	Scopes                []string
	Resources             []string
	Prompt                string
}

func GenerateSignInUri(option *SignInUriGenerationOptions) (string, error) {
	uri, parseUrlErr := url.Parse(option.AuthorizationEndpoint)

	if parseUrlErr != nil {
		return "", parseUrlErr
	}

	queries := uri.Query()

	queries.Add("client_id", option.ClientId)
	queries.Add("redirect_uri", option.RedirectUri)
	queries.Add("code_challenge", option.CodeChallenge)
	queries.Add("code_challenge_method", "S256")
	queries.Add("state", option.State)

	scopes := option.Scopes
	for _, defaultScope := range DefaultScopes {
		scopes = AppendIfNotExisted(scopes, defaultScope)
	}

	queries.Add("scope", strings.Join(scopes, " "))

	resources := option.Resources
	if slices.Contains(scopes, UserScopeOrganizations) {
		resources = AppendIfNotExisted(resources, ReservedResourceOrganization)
	}

	if len(resources) != 0 {
		for _, resource := range resources {
			queries.Add("resource", resource)
		}
	}

	queries.Add("response_type", "code")

	if option.Prompt != "" {
		queries.Add("prompt", option.Prompt)
	} else {
		queries.Add("prompt", "consent")
	}

	unescapedQueries, unescapeQueryErr := url.QueryUnescape(queries.Encode())

	if unescapeQueryErr != nil {
		return "", unescapeQueryErr
	}

	return uri.String() + "?" + unescapedQueries, nil
}
