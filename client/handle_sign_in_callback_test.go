package client

import (
	"net/http"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/logto-io/go/core"
	"github.com/stretchr/testify/assert"
	"gopkg.in/square/go-jose.v2"
)

func TestHandleSignInCallbackShouldHandleCallbackCorrectly(t *testing.T) {
	testAccessToken := "access token"
	testRefreshToken := "refresh token"
	testIdToken := "id token"

	testCodeTokenResponse := core.CodeTokenResponse{
		AccessToken:  testAccessToken,
		RefreshToken: testRefreshToken,
		IdToken:      testIdToken,
		Scope:        "openid",
		ExpireIn:     3600,
	}

	patchesForVerifyAndParseCodeFromCallbackUri := gomonkey.ApplyFunc(
		core.VerifyAndParseCodeFromCallbackUri,
		func(callbackUri, redirectUri, state string) (string, error) {
			return "code", nil
		},
	)
	defer patchesForVerifyAndParseCodeFromCallbackUri.Reset()

	patchesForFetchTokenByAuthorizationCode := gomonkey.ApplyFunc(
		core.FetchTokenByAuthorizationCode,
		func(client *http.Client, options *core.FetchTokenByAuthorizationCodeOptions) (core.CodeTokenResponse, error) {
			return testCodeTokenResponse, nil
		},
	)
	defer patchesForFetchTokenByAuthorizationCode.Reset()

	patchesForVerifyIdToken := gomonkey.ApplyFunc(
		core.VerifyIdToken,
		func(idToken, clientId, issuer string, jwks *jose.JSONWebKeySet) error {
			return nil
		},
	)
	defer patchesForVerifyIdToken.Reset()

	var logtoClientSpy *LogtoClient
	patchesForFetchOidcConfig := gomonkey.ApplyPrivateMethod(logtoClientSpy, "fetchOidcConfig", func(_ *LogtoClient) (core.OidcConfigResponse, error) {
		return core.OidcConfigResponse{}, nil
	})
	defer patchesForFetchOidcConfig.Reset()

	patchesForCreateRemoteJwks := gomonkey.ApplyPrivateMethod(logtoClientSpy, "createRemoteJwks", func(_ *LogtoClient) (*jose.JSONWebKeySet, error) {
		return &jose.JSONWebKeySet{}, nil
	})
	defer patchesForCreateRemoteJwks.Reset()

	storage := &TestStorage{
		data: map[string]string{
			StorageKeySignInSession: `{"RedirectUri":"redirectUri","CodeVerifier":"codeVerifier","CodeChallenge":"codeChallenge","State":"state"}`,
		},
	}

	logtoClient := NewLogtoClient(&LogtoConfig{}, storage)

	signInCallbackRequest, createSignInCallbackRequestErr := http.NewRequest("GET", "https://example.com/sign-in-callback", nil)
	assert.Nil(t, createSignInCallbackRequestErr)

	handleSignInErr := logtoClient.HandleSignInCallback(signInCallbackRequest)
	assert.Nil(t, handleSignInErr)

	assert.Equal(t, "", storage.GetItem(StorageKeySignInSession))
	assert.Equal(t, testIdToken, logtoClient.GetIdToken())
	assert.Equal(t, testRefreshToken, logtoClient.GetRefreshToken())
	assert.True(t, logtoClient.IsAuthenticated())

	accessToken, getAccessTokenErr := logtoClient.GetAccessToken("")
	assert.Nil(t, getAccessTokenErr)
	assert.Equal(t, testAccessToken, accessToken.Token)
}
