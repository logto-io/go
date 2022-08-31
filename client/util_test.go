package client

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"net/http"
	"testing"

	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

func TestGetOriginRequestUrlShouldReturnCorrectUrl(t *testing.T) {
	requestUrl := "http://example.com/good"
	request, createRequestErr := http.NewRequest("GET", requestUrl, nil)
	if createRequestErr != nil {
		t.Fatalf(createRequestErr.Error())
	}

	// specify request uri before really perform a request
	request.RequestURI = "/good"

	originUrl := GetOriginRequestUrl(request)

	if originUrl != requestUrl {
		t.Fatalf("Expected URL: %v\nActual URL: %v", requestUrl, originUrl)
	}
}

func TestGetOriginRequestUrlShouldReturnCorrectUrlWithForwardProtoConfig(t *testing.T) {
	requestUri := "example.com/good"
	request, createRequestErr := http.NewRequest("GET", "http://"+requestUri, nil)

	if createRequestErr != nil {
		t.Fatalf(createRequestErr.Error())
	}

	// specify request uri before really perform a request
	request.RequestURI = "/good"

	request.Header.Add("X-Forwarded-Proto", "https")

	if createRequestErr != nil {
		t.Fatalf(createRequestErr.Error())
	}

	originUrl := GetOriginRequestUrl(request)

	expectedUrl := "https://" + requestUri

	if originUrl != expectedUrl {
		t.Fatalf("Expected URL: %v\nActual URL: %v", expectedUrl, originUrl)
	}
}

func TestGetOriginRequestUrlShouldReturnCorrectUrlWithTlsConnection(t *testing.T) {
	requestUri := "example.com/good"
	request, createRequestErr := http.NewRequest("GET", "http://"+requestUri, nil)
	if createRequestErr != nil {
		t.Fatalf(createRequestErr.Error())
	}

	// specify request uri before really perform a request
	request.RequestURI = "/good"

	request.TLS = &tls.ConnectionState{}

	originUrl := GetOriginRequestUrl(request)

	expectedUrl := "https://" + requestUri

	if originUrl != expectedUrl {
		t.Fatalf("Expected URL: %v\nActual URL: %v", expectedUrl, originUrl)
	}
}

func TestCreateHttpClient(t *testing.T) {
	logtoConfig := &LogtoConfig{
		AppId:     "AppId",
		AppSecret: "AppSecret",
	}

	logtoHttpClient := createHttpClient(logtoConfig)

	if logtoHttpClient.Transport == http.DefaultTransport {
		t.Fatal("Create HTTP client for Logto failed")
	}
}

func TestBuildAccessTokenKeyShouldBuildCorrectly(t *testing.T) {
	tests := []struct {
		scopes   []string
		resource string
		want     string
	}{
		{[]string{}, "", "@"},
		{[]string{}, "http://api.example.com", "@http://api.example.com"},
		{[]string{"read", "write"}, "", "read write@"},
		{[]string{"read", "write"}, "http://api.example.com", "read write@http://api.example.com"},
		{[]string{"write", "read"}, "http://api.example.com", "read write@http://api.example.com"},
		{[]string{"read"}, "http://api.example.com", "read@http://api.example.com"},
	}

	for _, test := range tests {
		if got := buildAccessTokenKey(test.scopes, test.resource); got != test.want {
			t.Fatalf("Expected Access Token Key: %v\nActual Access Token Key: %v", test.want, got)
		}
	}
}

func TestGetResourceFromAccessTokenShouldGetResourceCorrectly(t *testing.T) {
	expectedResource := "example.com"
	testAccessToken, createTokenError := createTestToken(expectedResource)

	if createTokenError != nil {
		t.Fatal(createTokenError)
	}

	resource := getResourceFromAccessToken(testAccessToken)

	if resource != expectedResource {
		t.Fatalf("Expected Resource: %v\nActual Resource: %v", expectedResource, resource)
	}
}

func TestGetResourceFromAccessTokenShouldReturnEmptyStringIfResourceMissing(t *testing.T) {
	expectedResource := ""

	testAccessToken, createTokenErr := createTestToken(expectedResource)
	if createTokenErr != nil {
		t.Fatal(createTokenErr)
	}

	resource := getResourceFromAccessToken(testAccessToken)

	if resource != expectedResource {
		t.Fatalf("Expected Resource: %v\nActual Resource: %v", expectedResource, resource)
	}
}

func createTestToken(resource string) (string, error) {
	rsaPrivateKey, generateKeyError := rsa.GenerateKey(rand.Reader, 2048)

	if generateKeyError != nil {
		return "", generateKeyError
	}

	signingKey := jose.SigningKey{Algorithm: jose.RS256, Key: rsaPrivateKey}

	signingKeyOptions := jose.SignerOptions{}
	signingKeyOptions.WithType("JWT")

	rsaSigner, createSignerError := jose.NewSigner(signingKey, &signingKeyOptions)
	if createSignerError != nil {
		return "", createSignerError
	}

	builder := jwt.Signed(rsaSigner)

	type TestClaims struct {
		Sub string `json:"sub"`
		Aud string `json:"aud,omitempty"`
	}

	claims := TestClaims{
		Sub: "sub",
		Aud: resource,
	}

	token, buildTokenError := builder.Claims(claims).CompactSerialize()

	if buildTokenError != nil {
		return "", buildTokenError
	}

	return token, nil
}
