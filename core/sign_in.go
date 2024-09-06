package core

import (
	"net/url"
	"slices"
	"strings"
)

type DirectSignInOptions struct {
	Method string
	Target string
}

type SignInUriGenerationOptions struct {
	AuthorizationEndpoint string
	ClientId              string
	RedirectUri           string
	CodeChallenge         string
	State                 string
	Scopes                []string
	Resources             []string
	Prompt                string
	LoginHint             string
	FirstScreen           string
	Identifiers           []string
	DirectSignIn          *DirectSignInOptions
	ExtraParams           map[string]string
	IncludeReservedScopes *bool
}

func GenerateSignInUri(option *SignInUriGenerationOptions) (string, error) {
	uri, parseUrlErr := url.Parse(option.AuthorizationEndpoint)

	if parseUrlErr != nil {
		return "", parseUrlErr
	}

	queries := uri.Query()

	queries.Add(QueryKeyClientID, option.ClientId)
	queries.Add(QueryKeyRedirectURI, option.RedirectUri)
	queries.Add(QueryKeyCodeChallenge, option.CodeChallenge)
	queries.Add(QueryKeyCodeChallengeMethod, "S256")
	queries.Add(QueryKeyState, option.State)

	scopes := option.Scopes
	if option.IncludeReservedScopes == nil || *option.IncludeReservedScopes {
		for _, defaultScope := range DefaultScopes {
			scopes = AppendIfNotExisted(scopes, defaultScope)
		}
	}
	if len(scopes) != 0 {
		queries.Add(QueryKeyScope, strings.Join(scopes, " "))
	}

	resources := option.Resources
	if slices.Contains(scopes, UserScopeOrganizations) {
		resources = AppendIfNotExisted(resources, ReservedResourceOrganization)
	}

	if len(resources) != 0 {
		for _, resource := range resources {
			queries.Add(QueryKeyResource, resource)
		}
	}

	queries.Add(QueryKeyResponseType, ResponseTypeCode)

	if option.Prompt != "" {
		queries.Add(QueryKeyPrompt, option.Prompt)
	} else {
		queries.Add(QueryKeyPrompt, PromptConsent)
	}

	if option.LoginHint != "" {
		queries.Add(QueryKeyLoginHint, option.LoginHint)
	}

	if option.FirstScreen != "" {
		queries.Add(QueryKeyFirstScreen, option.FirstScreen)
	}

	if len(option.Identifiers) != 0 {
		queries.Add(QueryKeyIdentifiers, strings.Join(option.Identifiers, " "))
	}

	if option.DirectSignIn != nil {
		directSignInValue := option.DirectSignIn.Method + ":" + option.DirectSignIn.Target
		queries.Add(QueryKeyDirectSignIn, directSignInValue)
	}

	if len(option.ExtraParams) != 0 {
		for key, value := range option.ExtraParams {
			queries.Add(key, value)
		}
	}

	unescapedQueries, unescapeQueryErr := url.QueryUnescape(queries.Encode())

	if unescapeQueryErr != nil {
		return "", unescapeQueryErr
	}

	return uri.String() + "?" + unescapedQueries, nil
}
