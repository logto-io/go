package core

import (
	"net/http"
	"net/url"
)

type FetchTokenByAuthorizationCodeOptions struct {
	tokenEndpoint string
	code          string
	codeVerifier  string
	clientId      string
	redirectUri   string
	resource      string
}

func FetchTokenByAuthorizationCode(client *http.Client, options *FetchTokenByAuthorizationCodeOptions) (CodeTokenResponse, error) {
	values := url.Values{
		"client_id":     {options.clientId},
		"redirect_uri":  {options.redirectUri},
		"code_verifier": {options.codeVerifier},
		"code":          {options.code},
		"grant_type":    {"authorization_code"},
	}

	if options.resource != "" {
		values.Add("resource", options.resource)
	}

	response, requestErr := client.PostForm(options.tokenEndpoint, values)

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
	tokenEndpoint string
	clientId      string
	refreshToken  string
	resource      string
	scope         string
}

func FetchTokenByRefreshToken(client *http.Client, options *FetchTokenByRefreshTokenOptions) (RefreshTokenResponse, error) {
	values := url.Values{
		"client_id":     {options.clientId},
		"refresh_token": {options.refreshToken},
		"grant_type":    {"refresh_token"},
	}

	if options.resource != "" {
		values.Add("resource", options.resource)
	}

	if options.scope != "" {
		values.Add("scope", options.scope)
	}

	response, requestErr := client.PostForm(options.tokenEndpoint, values)

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
