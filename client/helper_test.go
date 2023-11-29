package client

import (
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/jarcoal/httpmock"
	"github.com/logto-io/go/core"
	"github.com/stretchr/testify/assert"
	"gopkg.in/square/go-jose.v2"
)

var mockedValidJwks string = `{` +
	`"keys": [` +
	`{` +
	`"kty": "RSA",` +
	`"use": "sig",` +
	`"kid": "tb5ZqyJ5nC3MixPnk0YWAVLAfFPIFFrFoC7zPASOeBQ",` +
	`"e": "AQAB",` +
	`"n": "zW2x7tnxYwvzLIDJqZiu_0FOteREp7J2QimYlYSWkMQmkz72N81UZ9NMzgXGf3Hq7ENERuj6v61_PdVmG5p7X3pawQcvvj6Pggh1PNtcx33Hg0sY_8Iki4VHwKsinTWYRCC6T463V3EjPKR9AWJFMYwZu8vV6kfurH83DD-lo7don5SzHZXjOCsIyAqRMTvqOi3OvZjAzFXRopaGIhPCJTk8-VnCcL71ZRcMgz03_8QEPM1jcX7qX4aMg8fJ9k0nr1WgQK_ttaCw3ydz6AqMRI-M-AZU_KVhgAc_lKtfPEvG852v7BFtJj3TtDZTUfPxXkH1cg11qYpT5cOw1Il0g6PB0Xfpck7_U3gyiG58u6_NvGkzjSXv7385_ws1bmfwKcDWN3RySi7cI8f0TAOAarA2rqh10qPg_8w5dpFA2P20HZgy8JFxP7VDSsDyV02hBs7MAC7xIhiwlX3vjQbiWDpJclWu3psA2FHeJulODwfHQ-pEBHsJDhZv6U0-nmf3JQ3Wrb0m9NUe6I2SQfPyzZMq2q3nUFD0EHJquhEcd8zxTcmWxt6Xug4sZwm27uBdTuixJ3cJZdDXjKwB8uv99zgZT1RmwMWOipZnMkVkkLLJpmfjvlRp4nLIXa7Ar7gXiB8WAynqrbMpC_92eCNCFvI7boepZvhFpOzvW0mYnZM"` +
	`},` +
	`{` +
	`"kty": "RSA",` +
	`"use": "sig",` +
	`"kid": "igun4zI3Hs-O4vqncDpnoE5Ysn-ayBePCZ89dKqxdLA",` +
	`"e": "AQAB",` +
	`"n": "8kaNkiwPhmT2v8YVcoSJ72bG_9kJwvtStGUxz6RVRMEzXXZG70igKSI3-V9LkBoVdPeN8CLX3adv1vX5R18EhNUUFn88_RpdKOi0DKPbq1Rs_-obWtabtbGhMy7uo5BMVdmfBXjDmDJU11W44_FOjW3d21KjsozUR2LJkENfWS8A8yhbeGVzl1fv062DO86Dwfy5Zknph7irs3ioPc4PBrMhX55TuEE77841Xx4qIJKfKjFQy9--XNAm3qOJyMCQG_eL_tcCcU0Okl_5tz6PSEAzQ5G3WYA4DgHsRrk9FlvI7LyGVtObH70aPX3YK0Ggh6ieQFnomj6wcN6II6dgslgMr7a4fNblS0j3_k4twbhn-LFb39cFKZekPrxT4LO58kTtNxT4tnPOQxdjy96p4ES6YBYPerpIUMHp1nyL3yY_Uxub_gGKRq4fxLK3QrJBFigwg7isoglxd1lGZaaVVQwz5LxjeBut2sFkIrWTzae_hFZHI9D8v_WwzhbwmIJzfP64aqBV-Rloj83e4BKXItBqwBQrcJtAvo1IYLekBC6HON2vczECEPLOyx47BOry1M4kxZfu342Ya_8gSqRxKQKVuKh88aNHFTUqn4jn2MFErDdjXdlqjtZGti1vWqHAz6hQ2ZDoUEqjFCSKWJ5QmHMdc3AMJpKLNkuQWIFJcm8"` +
	`}` +
	`]` +
	`}`

type TestStorage struct {
	data map[string]string
}

func (storage *TestStorage) GetItem(key string) string {
	return storage.data[key]
}

func (storage *TestStorage) SetItem(key, value string) {
	storage.data[key] = value
}

func TestFetchOidcConfigShouldReturnExpectedOidcConfig(t *testing.T) {
	testOidcConfig := core.OidcConfigResponse{
		AuthorizationEndpoint: "testAuthorizationEndpoint",
	}

	patches := gomonkey.ApplyFunc(core.FetchOidcConfig, func(client *http.Client, endpoint string) (core.OidcConfigResponse, error) {
		return testOidcConfig, nil
	})
	defer patches.Reset()

	logtClient := NewLogtoClient(&LogtoConfig{
		Endpoint: "https://example.com",
	}, &TestStorage{})

	oidcConfig, fetchOidcConfigError := logtClient.fetchOidcConfig()
	assert.Nil(t, fetchOidcConfigError)
	assert.Equal(t, testOidcConfig, oidcConfig)
}

func TestLoadAccessTokenMapShouldLoadAccessTokenMapFromStorage(t *testing.T) {
	expiresAt := time.Now().Unix() + 60

	logtoConfig := &LogtoConfig{}

	testStorage := &TestStorage{
		data: map[string]string{
			StorageKeyIdToken:        "ThisIdTokenMakeIsAuthenticatedReturnTrue",
			StorageKeyAccessTokenMap: `{"@testResource":{"token":"token","scope":"scope","expiresAt":` + strconv.FormatInt(expiresAt, 10) + `}}`,
		},
	}

	testAccessToken := AccessToken{
		Token:     "token",
		Scope:     "scope",
		ExpiresAt: expiresAt,
	}

	logtoClient := NewLogtoClient(logtoConfig, testStorage)

	logtoClient.loadAccessTokenMap()

	accessToken := logtoClient.accessTokenMap["@testResource"]

	assert.Equal(t, testAccessToken, accessToken)
}

func TestPersistAccessTokenMapShouldSaveAccessTokenMapDataToStorage(t *testing.T) {
	logtoConfig := &LogtoConfig{}
	testStorage := &TestStorage{
		data: make(map[string]string),
	}

	logtoClient := NewLogtoClient(logtoConfig, testStorage)

	logtoClient.accessTokenMap["accessTokenKey"] = AccessToken{
		Token:     "token",
		Scope:     "scope",
		ExpiresAt: 1000,
	}

	testAccessTokenMapContent := `{"accessTokenKey":{"token":"token","scope":"scope","expiresAt":1000}}`

	logtoClient.persistAccessTokenMap()

	accessTokenMapContent := testStorage.GetItem(StorageKeyAccessTokenMap)

	assert.Equal(t, testAccessTokenMapContent, accessTokenMapContent)
}

func TestCreateRemoteJwksShouldCreateExpectedJwks(t *testing.T) {

	logtoClient := NewLogtoClient(&LogtoConfig{}, &TestStorage{})

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	jwksUri := "http://example.com/oidc/jwks"

	httpmock.RegisterResponder(
		"GET",
		jwksUri,
		httpmock.NewStringResponder(200, mockedValidJwks),
	)

	client := &http.Client{}

	logtoClient.httpClient = client

	jwks, createRemoteJwksErr := logtoClient.createRemoteJwks(jwksUri)
	assert.Nil(t, createRemoteJwksErr)
	assert.Equal(t, 2, len(jwks.Keys))
}

func TestVerifyAndSaveTokenResponseShouldSaveToken(t *testing.T) {
	var logtoClientSpy *LogtoClient
	patchesForCreateRemoteJwks := gomonkey.ApplyPrivateMethod(logtoClientSpy, "createRemoteJwks", func(_ *LogtoClient) (*jose.JSONWebKeySet, error) {
		return &jose.JSONWebKeySet{}, nil
	})
	defer patchesForCreateRemoteJwks.Reset()

	patchesForVerifyIdToken := gomonkey.ApplyFunc(core.VerifyIdToken, func(idToken, clientId, issuer string, jwks *jose.JSONWebKeySet) error {
		return nil
	})

	defer patchesForVerifyIdToken.Reset()

	logtoClient := NewLogtoClient(&LogtoConfig{}, &TestStorage{
		data: map[string]string{},
	})

	testIdToken := "idToken"
	testRefreshToken := "refreshToken"
	testAccessTokenKey := "@"
	testAccessToken := AccessToken{Token: "testToken", Scope: "read", ExpiresAt: time.Now().Unix() + 60}

	verifyAndSaveTokenResponseErr := logtoClient.verifyAndSaveTokenResponse(testIdToken, testRefreshToken, testAccessTokenKey, testAccessToken, &core.OidcConfigResponse{})
	assert.Nil(t, verifyAndSaveTokenResponseErr)

	assert.Equal(t, testIdToken, logtoClient.GetIdToken())
	assert.Equal(t, testRefreshToken, logtoClient.GetRefreshToken())

	accessToken, getAccessTokenErr := logtoClient.GetAccessToken("")
	assert.Nil(t, getAccessTokenErr)
	assert.Equal(t, testAccessToken, accessToken)
}
