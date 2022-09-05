package client

import (
	"net/http"
	"testing"
	"time"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/logto-io/go/core"
	"github.com/stretchr/testify/assert"
	"gopkg.in/square/go-jose.v2"
)

func TestGetAccessTokenShouldReturnAccessTokenAccessTokenInTokenMap(t *testing.T) {
	testAccessToken := AccessToken{
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

	logtoClient.accessTokenMap["@"] = testAccessToken

	accessToken, getAccessTokenErr := logtoClient.GetAccessToken("")

	assert.Nil(t, getAccessTokenErr)
	assert.Equal(t, testAccessToken, accessToken)
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

	_, getAccessTokenErr := logtoClient.GetAccessToken("")

	assert.Equal(t, ErrNotAuthenticated, getAccessTokenErr)
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

	accessToken, getAccessTokenErr := logtoClient.GetAccessToken("")

	assert.Nil(t, getAccessTokenErr)
	assert.Equal(t, testAccessToken, accessToken.Token)
	assert.Equal(t, testAccessToken, logtoClient.accessTokenMap["@"].Token)
}
