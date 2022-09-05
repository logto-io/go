package client

import (
	"encoding/json"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/logto-io/go/core"
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

	var logtoClientSpy *LogtoClient

	patchesForFetchOidcConfig := gomonkey.ApplyPrivateMethod(logtoClientSpy, "fetchOidcConfig", func(_ *LogtoClient) (core.OidcConfigResponse, error) {
		oidcConfig := core.OidcConfigResponse{
			AuthorizationEndpoint: testAuthEndpoint,
		}
		return oidcConfig, nil
	})
	defer patchesForFetchOidcConfig.Reset()

	logtoConfig := &LogtoConfig{
		Endpoint:  "https://example.com",
		AppId:     testAppId,
		AppSecret: testAppSecret,
	}

	storage := &TestStorage{
		data: make(map[string]string),
	}

	logtoClient := NewLogtoClient(logtoConfig, storage)

	gotSignInUri, signInErr := logtoClient.SignIn(testRedirectUri)

	if signInErr != nil {
		t.Fatal(signInErr)
	}

	if gotSignInUri != testSignInUri {
		t.Fatalf("Expected sign-in URI : %v\nActual sign-in URI : %v", testAuthEndpoint, gotSignInUri)
	}

	signInContext := SignInContext{}
	parseSignInContextErr := json.Unmarshal([]byte(storage.GetItem(StorageKeySignInContext)), &signInContext)

	if parseSignInContextErr != nil {
		t.Fatal(parseSignInContextErr)
	}

	if signInContext.RedirectUri != testRedirectUri {
		t.Fatalf("Expected redirect uri: %v\nActual redirect uri: %v", testRedirectUri, signInContext.RedirectUri)
	}

	if signInContext.CodeChallenge != testCodeChallenge {
		t.Fatal("code challenge in uri does not match the one in sign-in context")
	}

	if signInContext.State != testState {
		t.Fatal("state in uri does not match the one in sign-in context")
	}

	if signInContext.CodeVerifier != testCodeVerifier {
		t.Fatal("code verifier in sign-in context does not match the test code verifier")
	}
}
