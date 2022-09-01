package client

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/logto-io/go/core"
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
			StorageKeySignInContext: `{"RedirectUri":"redirectUri","CodeVerifier":"codeVerifier","CodeChallenge":"codeChallenge","State":"state"}`,
		},
	}

	logtoClient := NewLogtoClient(&LogtoConfig{
		Resources: []string{""},
	}, storage)

	signInCallbackRequest, createSignInCallbackRequestErr := http.NewRequest("GET", "https://example.com/sign-in-callback", nil)

	if createSignInCallbackRequestErr != nil {
		t.Fatal(createSignInCallbackRequestErr)
	}

	handleSignInErr := logtoClient.HandleSignInCallback(signInCallbackRequest)

	if handleSignInErr != nil {
		t.Fatal(handleSignInErr)
	}

	if storage.GetItem(StorageKeySignInContext) != "" {
		t.Fatal("sign-in context must be cleared after handle sign-in callback")
	}

	gotIdToken := logtoClient.GetIdToken()
	if gotIdToken != testIdToken {
		t.Fatalf("Expected id token: %s\nActual id token: %v", testIdToken, gotIdToken)
	}

	gotAccessToken, _ := logtoClient.GetAccessToken("")
	if gotAccessToken.Token != testAccessToken {
		fmt.Println(logtoClient.accessTokenMap)
		t.Fatalf("Expected access token: %s\nActual access token: %v", testAccessToken, gotAccessToken.Token)
	}

	gotRefreshToken := logtoClient.GetRefreshToken()
	if gotRefreshToken != testRefreshToken {
		t.Fatalf("Expected refresh token: %s\nActual refresh token: %v", testRefreshToken, gotRefreshToken)
	}

	if !logtoClient.IsAuthenticated() {
		t.Fatal("Expected authenticate state: authenticated\nActual authenticate state: not authenticated")
	}
}
