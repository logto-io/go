package client

import (
	"net/http"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/logto-io/go/core"
	"github.com/stretchr/testify/assert"
)

func TestSignOutShouldSignOutUserSuccessfully(t *testing.T) {
	testEndSessionEndpoint := "https://example.com/end-session"
	testSignOutUri := "testSignOutUri"

	var logtoClientSpy *LogtoClient

	patchesForFetchOidcConfig := gomonkey.ApplyPrivateMethod(logtoClientSpy, "fetchOidcConfig", func(_ *LogtoClient) (core.OidcConfigResponse, error) {
		oidcConfig := core.OidcConfigResponse{
			EndSessionEndpoint: testEndSessionEndpoint,
		}
		return oidcConfig, nil
	})
	defer patchesForFetchOidcConfig.Reset()

	patchesForRevoke := gomonkey.ApplyFunc(core.Revoke, func(client *http.Client, options *core.RevocationOptions) error {
		return nil
	})
	defer patchesForRevoke.Reset()

	patchesForGenerateSignOutUri := gomonkey.ApplyFunc(core.GenerateSignOutUri, func(option *core.SignOutUriGenerationOptions) (string, error) {
		return testSignOutUri, nil
	})
	defer patchesForGenerateSignOutUri.Reset()

	storage := &TestStorage{
		data: map[string]string{
			StorageKeyIdToken:        "idToken",
			StorageKeyAccessTokenMap: `{"key": {"token": "abc"}}`,
			StorageKeyRefreshToken:   "refreshToken",
		},
	}

	logtoClient := NewLogtoClient(&LogtoConfig{}, storage)

	logtoClient.accessTokenMap = map[string]AccessToken{
		"key": {Token: "abc"},
	}

	signOutUri, signOutErr := logtoClient.SignOut("https://example.com/home")
	assert.Nil(t, signOutErr)
	assert.Equal(t, testSignOutUri, signOutUri)

	assert.Equal(t, 0, len(logtoClient.accessTokenMap))
	assert.Equal(t, "", storage.GetItem(StorageKeyIdToken))
	assert.Equal(t, "", storage.GetItem(StorageKeyRefreshToken))
	assert.Equal(t, "", storage.GetItem(StorageKeyAccessTokenMap))
}
