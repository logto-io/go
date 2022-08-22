package core

import (
	"net/http"
	"net/url"
	"strings"
)

type FetchTokenByAuthorizationCodeOptions struct {
	TokenEndpoint string
	Code          string
	CodeVerifier  string
	ClientId      string
	RedirectUri   string
	Resource      string
}

func FetchTokenByAuthorizationCode(client *http.Client, options *FetchTokenByAuthorizationCodeOptions) (CodeTokenResponse, error) {
	values := url.Values{
		"client_id":     {options.ClientId},
		"redirect_uri":  {options.RedirectUri},
		"code_verifier": {options.CodeVerifier},
		"code":          {options.Code},
		"grant_type":    {"authorization_code"},
	}

	if options.Resource != "" {
		values.Add("resource", options.Resource)
	}

	response, requestErr := client.PostForm(options.TokenEndpoint, values)

	if requestErr != nil {
		return CodeTokenResponse{}, requestErr
	}

	defer response.Body.Close()

	var codeTokenResponse CodeTokenResponse
	err := parseDataFromResponse(response, &codeTokenResponse)

	if err != nil {
		return CodeTokenResponse{}, err
	}

	return codeTokenResponse, nil
}

type FetchTokenByRefreshTokenOptions struct {
	TokenEndpoint string
	ClientId      string
	RefreshToken  string
	Resource      string
	Scopes        []string
}

func FetchTokenByRefreshToken(client *http.Client, options *FetchTokenByRefreshTokenOptions) (RefreshTokenResponse, error) {
	values := url.Values{
		"client_id":     {options.ClientId},
		"refresh_token": {options.RefreshToken},
		"grant_type":    {"refresh_token"},
	}

	if options.Resource != "" {
		values.Add("resource", options.Resource)
	}

	if len(options.Scopes) > 0 {
		values.Add("scope", strings.Join(options.Scopes, " "))
	}

	response, requestErr := client.PostForm(options.TokenEndpoint, values)

	if requestErr != nil {
		return RefreshTokenResponse{}, requestErr
	}

	defer response.Body.Close()

	var refreshTokenResponse RefreshTokenResponse
	err := parseDataFromResponse(response, &refreshTokenResponse)

	if err != nil {
		return RefreshTokenResponse{}, err
	}

	return refreshTokenResponse, nil
}
