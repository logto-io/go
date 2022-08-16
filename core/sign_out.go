package core

import "net/url"

type SignOutUriGenerationOptions struct {
	EndSessionEndpoint    string
	IdToken               string
	PostLogoutRedirectUri string
}

func GenerateSignOutUri(option *SignOutUriGenerationOptions) (string, error) {
	uri, parseUrlErr := url.Parse(option.EndSessionEndpoint)
	if parseUrlErr != nil {
		return "", parseUrlErr
	}

	queries := uri.Query()

	queries.Add("id_token_hint", option.IdToken)

	if option.PostLogoutRedirectUri != "" {
		queries.Add("post_logout_redirect_uri", option.PostLogoutRedirectUri)
	}

	unescapedQueries, unescapeQueryErr := url.QueryUnescape(queries.Encode())
	if unescapeQueryErr != nil {
		return "", unescapeQueryErr
	}

	return uri.String() + "?" + unescapedQueries, nil
}
