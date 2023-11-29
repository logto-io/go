package core

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestFetchTokenByAuthorizationCode(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	tokenEndpoint := "http://example.com/oidc/token"
	mockResponse := `{` +
		`"access_token": "access_token",` +
		`"refresh_token": "refresh_token",` +
		`"id_token": "id_token",` +
		`"scope": "openid offline_access",` +
		`"expires_in": 3600` +
		`}`

	httpmock.RegisterResponder(
		"POST",
		tokenEndpoint,
		httpmock.NewStringResponder(200, mockResponse),
	)

	client := &http.Client{}
	options := &FetchTokenByAuthorizationCodeOptions{
		TokenEndpoint: tokenEndpoint,
		Code:          "code",
		CodeVerifier:  "codeVerifier",
		ClientId:      "clientId",
		ClientSecret:  "clientSecret",
		RedirectUri:   "redirectUri",
		Resource:      "resource",
	}

	token, fetchError := FetchTokenByAuthorizationCode(client, options)
	assert.Nil(t, fetchError)

	var testToken CodeTokenResponse
	unmarshalErr := json.Unmarshal([]byte(mockResponse), &testToken)
	assert.Nil(t, unmarshalErr)

	assert.Equal(t, testToken, token)
}

func TestFetchTokenByRefreshToken(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	tokenEndpoint := "http://example.com/oidc/token"
	mockResponse := `{` +
		`"access_token": "access_token",` +
		`"refresh_token": "refresh_token",` +
		`"id_token": "id_token",` +
		`"scope": "openid offline_access",` +
		`"expires_in": 3600` +
		`}`

	httpmock.RegisterResponder(
		"POST",
		tokenEndpoint,
		httpmock.NewStringResponder(200, mockResponse),
	)

	client := &http.Client{}
	options := &FetchTokenByRefreshTokenOptions{
		TokenEndpoint:  tokenEndpoint,
		ClientId:       "clientId",
		ClientSecret:   "clientSecret",
		RefreshToken:   "refresh_token",
		Resource:       "resource",
		Scopes:         []string{"openid", "offline_access"},
		OrganizationId: "organizationId",
	}

	token, fetchError := FetchTokenByRefreshToken(client, options)
	assert.Nil(t, fetchError)

	var testToken RefreshTokenResponse
	unmarshalErr := json.Unmarshal([]byte(mockResponse), &testToken)
	assert.Nil(t, unmarshalErr)

	assert.Equal(t, testToken, token)
}
