package client

import (
	"net/http"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/logto-io/go/core"
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

	gotSignOutUri, signOutErr := logtoClient.SignOut("https://example.com/home")

	if signOutErr != nil {
		t.Fatal(signOutErr)
	}

	if gotSignOutUri != testSignOutUri {
		t.Fatalf("Expected sign-out URI : %v\nActual sign-out URI : %v", testSignOutUri, gotSignOutUri)
	}

	if storage.GetItem(StorageKeyIdToken) != "" {
		t.Fatal("id token has not cleared")
	}

	if storage.GetItem(StorageKeyRefreshToken) != "" {
		t.Fatal("refresh token has not cleared")
	}

	if storage.GetItem(StorageKeyAccessTokenMap) != "" {
		t.Fatal("persisted access token map token has not cleared")
	}

	if len(logtoClient.accessTokenMap) != 0 {
		t.Fatal("persisted access token map token has not cleared")
	}
}
