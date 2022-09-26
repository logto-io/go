package client

import (
	"errors"
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
		&LogtoConfig{},
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
		&LogtoConfig{},
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
		&LogtoConfig{},
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

func TestGetIdTokenClaimsShouldReturnIdTokenClaimsCorrectly(t *testing.T) {
	testIdTokenClaims := core.IdTokenClaims{
		Sub: "testSub",
	}

	var logtoClientSpy *LogtoClient
	patchesForIsAuthenticated := gomonkey.ApplyPrivateMethod(logtoClientSpy, "IsAuthenticated", func(_ *LogtoClient) bool {
		return true
	})
	defer patchesForIsAuthenticated.Reset()

	patchesForGetIdToken := gomonkey.ApplyPrivateMethod(logtoClientSpy, "GetIdToken", func(_ *LogtoClient) string {
		return "idToken"
	})
	defer patchesForGetIdToken.Reset()

	patchesForDecodeIdToken := gomonkey.ApplyFunc(core.DecodeIdToken, func(token string) (core.IdTokenClaims, error) {
		return testIdTokenClaims, nil
	})
	defer patchesForDecodeIdToken.Reset()

	logtoClient := NewLogtoClient(&LogtoConfig{}, &TestStorage{})

	idTokenClaims, getIdTokenClaimsErr := logtoClient.GetIdTokenClaims()

	assert.Nil(t, getIdTokenClaimsErr)
	assert.Equal(t, testIdTokenClaims, idTokenClaims)
}

func TestGetIdTokenClaimsShouldReturnNotAuthenticatedErrorIfUserIsNotAuthenticated(t *testing.T) {
	var logtoClientSpy *LogtoClient
	patchesForIsAuthenticated := gomonkey.ApplyPrivateMethod(logtoClientSpy, "IsAuthenticated", func(_ *LogtoClient) bool {
		return false
	})
	defer patchesForIsAuthenticated.Reset()

	logtoClient := NewLogtoClient(&LogtoConfig{}, &TestStorage{})

	_, getIdTokenClaimsErr := logtoClient.GetIdTokenClaims()

	assert.Equal(t, ErrNotAuthenticated, getIdTokenClaimsErr)
}

func TestFetchUserInfoShouldReturnCorrectUserInfoResponse(t *testing.T) {
	testUserInfoResponse := core.UserInfoResponse{
		Sub:      "sub",
		Username: "username",
	}

	var logtoClientSpy *LogtoClient
	patchesForIsAuthenticated := gomonkey.ApplyPrivateMethod(logtoClientSpy, "IsAuthenticated", func(_ *LogtoClient) bool {
		return true
	})
	defer patchesForIsAuthenticated.Reset()

	patchesForFetchOidcConfig := gomonkey.ApplyPrivateMethod(logtoClientSpy, "fetchOidcConfig", func(_ *LogtoClient) (core.OidcConfigResponse, error) {
		return core.OidcConfigResponse{}, nil
	})
	defer patchesForFetchOidcConfig.Reset()

	patchesForGetAccessToken := gomonkey.ApplyPrivateMethod(logtoClientSpy, "GetAccessToken", func(_ *LogtoClient) (AccessToken, error) {
		return AccessToken{}, nil
	})
	defer patchesForGetAccessToken.Reset()

	patchesCoreFetchUserInfo := gomonkey.ApplyFunc(core.FetchUserInfo, func(client *http.Client, userInfoEndpoint, accessToken string) (core.UserInfoResponse, error) {
		return testUserInfoResponse, nil
	})
	defer patchesCoreFetchUserInfo.Reset()

	logtoClient := NewLogtoClient(&LogtoConfig{}, &TestStorage{})

	userInfoResponse, fetchUserInfoResponseErr := logtoClient.FetchUserInfo()

	assert.Nil(t, fetchUserInfoResponseErr)
	assert.Equal(t, testUserInfoResponse, userInfoResponse)
}

func TestFetchUserInfoShouldReturnNotAuthenticatedErrorIfUserIsNotAuthenticated(t *testing.T) {
	var logtoClientSpy *LogtoClient
	patchesForIsAuthenticated := gomonkey.ApplyPrivateMethod(logtoClientSpy, "IsAuthenticated", func(_ *LogtoClient) bool {
		return false
	})
	defer patchesForIsAuthenticated.Reset()

	logtoClient := NewLogtoClient(&LogtoConfig{}, &TestStorage{})

	_, fetchUserInfoErr := logtoClient.FetchUserInfo()

	assert.Equal(t, ErrNotAuthenticated, fetchUserInfoErr)
}

func TestFetchUserInfoShouldReturnErrorIfFetchOidcConfigFailed(t *testing.T) {
	testFetchOidcConfigErr := errors.New("fetch oidc config error")

	var logtoClientSpy *LogtoClient
	patchesForIsAuthenticated := gomonkey.ApplyPrivateMethod(logtoClientSpy, "IsAuthenticated", func(_ *LogtoClient) bool {
		return true
	})
	defer patchesForIsAuthenticated.Reset()

	patchesForFetchOidcConfig := gomonkey.ApplyPrivateMethod(logtoClientSpy, "fetchOidcConfig", func(_ *LogtoClient) (core.OidcConfigResponse, error) {
		return core.OidcConfigResponse{}, testFetchOidcConfigErr
	})
	defer patchesForFetchOidcConfig.Reset()

	logtoClient := NewLogtoClient(&LogtoConfig{}, &TestStorage{})

	_, fetchUserInfoErr := logtoClient.FetchUserInfo()

	assert.Equal(t, testFetchOidcConfigErr, fetchUserInfoErr)
}

func TestFetchUserInfoShouldReturnErrorIfGetAccessTokenFailed(t *testing.T) {
	testGetAccessTokenErr := errors.New("get access token error")

	var logtoClientSpy *LogtoClient
	patchesForIsAuthenticated := gomonkey.ApplyPrivateMethod(logtoClientSpy, "IsAuthenticated", func(_ *LogtoClient) bool {
		return true
	})
	defer patchesForIsAuthenticated.Reset()

	patchesForFetchOidcConfig := gomonkey.ApplyPrivateMethod(logtoClientSpy, "fetchOidcConfig", func(_ *LogtoClient) (core.OidcConfigResponse, error) {
		return core.OidcConfigResponse{}, nil
	})
	defer patchesForFetchOidcConfig.Reset()

	patchesForGetAccessToken := gomonkey.ApplyPrivateMethod(logtoClientSpy, "GetAccessToken", func(_ *LogtoClient) (AccessToken, error) {
		return AccessToken{}, testGetAccessTokenErr
	})
	defer patchesForGetAccessToken.Reset()

	logtoClient := NewLogtoClient(&LogtoConfig{}, &TestStorage{})

	_, fetchUserInfoErr := logtoClient.FetchUserInfo()

	assert.Equal(t, testGetAccessTokenErr, fetchUserInfoErr)
}
