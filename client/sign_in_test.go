package client

import (
	"encoding/json"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/logto-io/go/core"
	"github.com/stretchr/testify/assert"
)

func TestSignInShouldReturnSignInUriCorrectly(t *testing.T) {
	testAppId := "appId"
	testAppSecret := "appSecret"
	testAuthEndpoint := "https://example.com/auth"
	testRedirectUri := "https://myapp.com/sign-in-callback"
	testState := "testState"
	testCodeVerifier := "testCodeVerifier"
	testCodeChallenge := "testCodeChallenge"
	testSignInUri := "testSignInUri"

	var logtoClientSpy *LogtoClient

	patchesForFetchOidcConfig := gomonkey.ApplyPrivateMethod(logtoClientSpy, "fetchOidcConfig", func(_ *LogtoClient) (core.OidcConfigResponse, error) {
		oidcConfig := core.OidcConfigResponse{
			AuthorizationEndpoint: testAuthEndpoint,
		}
		return oidcConfig, nil
	})
	defer patchesForFetchOidcConfig.Reset()

	patchesForGenerateState := gomonkey.ApplyFunc(core.GenerateState, func() string {
		return testState
	})
	defer patchesForGenerateState.Reset()

	patchesForGenerateCodeVerifier := gomonkey.ApplyFunc(core.GenerateCodeVerifier, func() string {
		return testCodeVerifier
	})
	defer patchesForGenerateCodeVerifier.Reset()

	patchesForGenerateCodeChallenge := gomonkey.ApplyFunc(core.GenerateCodeChallenge, func(codeVerifier string) string {
		return testCodeChallenge
	})
	defer patchesForGenerateCodeChallenge.Reset()

	patchesForGenerateSignInUri := gomonkey.ApplyFunc(core.GenerateSignInUri, func(option *core.SignInUriGenerationOptions) (string, error) {
		return testSignInUri, nil
	})
	defer patchesForGenerateSignInUri.Reset()

	logtoConfig := &LogtoConfig{
		Endpoint:  "https://example.com",
		AppId:     testAppId,
		AppSecret: testAppSecret,
	}

	storage := &TestStorage{
		data: make(map[string]string),
	}

	logtoClient := NewLogtoClient(logtoConfig, storage)

	signInUri, signInErr := logtoClient.SignInWithRedirectUri(testRedirectUri)
	assert.Nil(t, signInErr)
	assert.Equal(t, testSignInUri, signInUri)

	signInSession := SignInSession{}
	parseSignInSessionErr := json.Unmarshal([]byte(storage.GetItem(StorageKeySignInSession)), &signInSession)

	assert.Nil(t, parseSignInSessionErr)
	assert.Equal(t, testRedirectUri, signInSession.RedirectUri)
	assert.Equal(t, testCodeChallenge, signInSession.CodeChallenge)
	assert.Equal(t, testState, signInSession.State)
	assert.Equal(t, testCodeVerifier, signInSession.CodeVerifier)
}
