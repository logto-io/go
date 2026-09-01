package core

import "net/url"

type SignOutUriGenerationOptions struct {
	EndSessionEndpoint    string
	ClientId              string
	PostLogoutRedirectUri string
}

func GenerateSignOutUri(option *SignOutUriGenerationOptions) (string, error) {
	uri, parseUrlErr := url.Parse(option.EndSessionEndpoint)
	if parseUrlErr != nil {
		return "", parseUrlErr
	}

	queries := uri.Query()

	queries.Add("client_id", option.ClientId)

	if option.PostLogoutRedirectUri != "" {
		queries.Add("post_logout_redirect_uri", option.PostLogoutRedirectUri)
	}

	uri.RawQuery = queries.Encode()
	return uri.String(), nil
}
