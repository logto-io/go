package core

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jarcoal/httpmock"
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
		RedirectUri:   "redirectUri",
		Resource:      "resource",
	}

	token, fetchError := FetchTokenByAuthorizationCode(client, options)
	if fetchError != nil {
		t.Fatalf(fetchError.Error())
	}

	var expectedToken CodeTokenResponse
	unmarshalErr := json.Unmarshal([]byte(mockResponse), &expectedToken)

	if unmarshalErr != nil {
		t.Fatalf(unmarshalErr.Error())
	}

	if !cmp.Equal(token, expectedToken) {
		t.Fatalf("token does not match expected result")
	}
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
		TokenEndpoint: tokenEndpoint,
		ClientId:      "clientId",
		RefreshToken:  "refresh_token",
		Resource:      "resource",
		Scopes:        []string{"openid", "offline_access"},
	}

	token, fetchError := FetchTokenByRefreshToken(client, options)
	if fetchError != nil {
		t.Fatalf(fetchError.Error())
	}

	var expectedToken RefreshTokenResponse
	unmarshalErr := json.Unmarshal([]byte(mockResponse), &expectedToken)

	if unmarshalErr != nil {
		t.Fatalf(unmarshalErr.Error())
	}

	if !cmp.Equal(token, expectedToken) {
		t.Fatalf("token does not match expected result")
	}
}
