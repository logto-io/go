package client

import (
	"encoding/json"
	"net/url"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/logto-io/go/core"
)

func TestSignInShouldReturnSignInUriCorrectly(t *testing.T) {
	testAppId := "appId"
	testAppSecret := "appSecret"
	testAuthEndpoint := "https://example.com/auth"
	testRedirectUri := "https://myapp.com/sign-in-callback"

	var logtoClientSpy *LogtoClient

	patches := gomonkey.ApplyPrivateMethod(logtoClientSpy, "fetchOidcConfig", func(_ *LogtoClient) (core.OidcConfigResponse, error) {
		oidcConfig := core.OidcConfigResponse{
			AuthorizationEndpoint: testAuthEndpoint,
		}
		return oidcConfig, nil
	})

	defer patches.Reset()

	logtoConfig := &LogtoConfig{
		Endpoint:  "https://example.com",
		AppId:     testAppId,
		AppSecret: testAppSecret,
	}

	storage := &TestStorage{
		data: make(map[string]string),
	}

	logtoClient := NewLogtoClient(logtoConfig, storage)

	signInUri, signInErr := logtoClient.SignIn(testRedirectUri)

	if signInErr != nil {
		t.Fatal(signInErr)
	}

	parsedUri, parseUriErr := url.Parse(signInUri)

	if parseUriErr != nil {
		t.Fatal(parseUriErr)
	}

	gotAuthEndpoint := parsedUri.Scheme + "://" + parsedUri.Host + parsedUri.Path

	if gotAuthEndpoint != testAuthEndpoint {
		t.Fatalf("Expected auth endpoint : %v\nActual auth endpoint : %v", testAuthEndpoint, gotAuthEndpoint)
	}

	signInContext := SignInContext{}
	parseSignInContextErr := json.Unmarshal([]byte(storage.GetItem(StorageKeySignInContext)), &signInContext)

	if parseSignInContextErr != nil {
		t.Fatal(parseSignInContextErr)
	}

	redirectUriInQuery := parsedUri.Query().Get("redirect_uri")

	if redirectUriInQuery != testRedirectUri {
		t.Fatalf("Expected redirect uri: %v\nActual redirect uri: %v", testRedirectUri, redirectUriInQuery)
	}

	codeChallengeInQuery := parsedUri.Query().Get("code_challenge")

	if signInContext.CodeChallenge != codeChallengeInQuery {
		t.Fatal("code challenge in uri does not match the one in sign-in context")
	}

	stateInQuery := parsedUri.Query().Get("state")

	if signInContext.State != stateInQuery {
		t.Fatal("state in uri does not match the one in sign-in context")
	}
}
