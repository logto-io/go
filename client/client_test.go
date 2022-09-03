package client

import (
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/google/go-cmp/cmp"
	"github.com/logto-io/go/core"
	"gopkg.in/square/go-jose.v2"
)

func TestGetAccessTokenShouldReturnAccessTokenAccessTokenInTokenMap(t *testing.T) {
	wantedAccessToken := AccessToken{
		Token:     "accessToken",
		Scope:     "openid",
		ExpiresAt: time.Now().Unix() + 60,
	}

	logtoClient := NewLogtoClient(
		&LogtoConfig{
			Resources: []string{""},
		},
		&TestStorage{
			data: map[string]string{
				StorageKeyIdToken: "id token",
			},
		},
	)

	logtoClient.accessTokenMap["@"] = wantedAccessToken

	gotAccessToken, getAccessTokenErr := logtoClient.GetAccessToken("")

	if getAccessTokenErr != nil {
		t.Fatal(getAccessTokenErr)
	}

	if !cmp.Equal(gotAccessToken, wantedAccessToken) {
		t.Fatalf("Expected Access Token : %v\nActual Access Token : %v", wantedAccessToken, gotAccessToken)
	}
}

func TestGetAccessTokenShouldReturnNotAuthenticatedErrIfNoIdTokenAvailable(t *testing.T) {
	logtoClient := NewLogtoClient(
		&LogtoConfig{
			Resources: []string{""},
		},
		&TestStorage{
			data: map[string]string{
				StorageKeyIdToken: "",
			},
		},
	)

	_, gotErr := logtoClient.GetAccessToken("")

	wantedErr := errors.New("not authenticated")

	if !cmp.Equal(gotErr.Error(), wantedErr.Error()) {
		t.Fatalf("Expected Error : %s\nActual Error : %s", wantedErr.Error(), gotErr.Error())
	}
}

func TestGetAccessTokenShouldReturnFetchedAccessTokenAndUpdateLocalAccessTokenIfLocalAccessTokenIsExpired(t *testing.T) {
	testAccessToken := "refreshed access token"
	testRefreshToken := "refresh token"
	testIdToken := "id token"

	testRefreshTokenResponse := core.RefreshTokenResponse{
		AccessToken:  testAccessToken,
		RefreshToken: testRefreshToken,
		IdToken:      testIdToken,
		Scope:        "openid",
		ExpireIn:     3600,
	}

	var logtoClientSpy *LogtoClient
	patchesForFetchOidcConfig := gomonkey.ApplyPrivateMethod(logtoClientSpy, "fetchOidcConfig", func(_ *LogtoClient) (core.OidcConfigResponse, error) {
		return core.OidcConfigResponse{}, nil
	})
	defer patchesForFetchOidcConfig.Reset()

	patchesForCreateRemoteJwks := gomonkey.ApplyPrivateMethod(logtoClientSpy, "createRemoteJwks", func(_ *LogtoClient) (*jose.JSONWebKeySet, error) {
		return &jose.JSONWebKeySet{}, nil
	})
	defer patchesForCreateRemoteJwks.Reset()

	patchesForFetchTokenByRefreshToken := gomonkey.ApplyFunc(
		core.FetchTokenByRefreshToken,
		func(client *http.Client, options *core.FetchTokenByRefreshTokenOptions) (core.RefreshTokenResponse, error) {
			return testRefreshTokenResponse, nil
		},
	)
	defer patchesForFetchTokenByRefreshToken.Reset()

	patchesForVerifyIdToken := gomonkey.ApplyFunc(
		core.VerifyIdToken,
		func(idToken, clientId, issuer string, jwks *jose.JSONWebKeySet) error {
			return nil
		},
	)
	defer patchesForVerifyIdToken.Reset()

	logtoClient := NewLogtoClient(
		&LogtoConfig{
			Resources: []string{""},
		},
		&TestStorage{
			data: map[string]string{
				StorageKeyIdToken:      "id token",
				StorageKeyRefreshToken: "refresh token",
			},
		},
	)

	expiredAccessToken := AccessToken{
		Token:     "expired access token",
		Scope:     "openid",
		ExpiresAt: time.Now().Unix() - 60,
	}

	logtoClient.accessTokenMap["@"] = expiredAccessToken

	gotAccessToken, getAccessTokenErr := logtoClient.GetAccessToken("")

	if getAccessTokenErr != nil {
		t.Fatal(getAccessTokenErr)
	}

	if gotAccessToken.Token != testAccessToken {
		t.Fatalf("Expected access token : %s\nActual access token : %s", testAccessToken, gotAccessToken.Token)
	}

	updatedAccessToken := logtoClient.accessTokenMap["@"]

	if updatedAccessToken.Token != testAccessToken {
		t.Fatalf("Local access token have not been updated, want : %s\ngot : %s", testAccessToken, updatedAccessToken.Token)
	}
}
