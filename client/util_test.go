package client

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

func TestGetOriginRequestUrlShouldReturnCorrectUrl(t *testing.T) {
	testRequestUrl := "http://example.com/good"
	request, createRequestErr := http.NewRequest("GET", testRequestUrl, nil)
	assert.Nil(t, createRequestErr)
	// specify request uri before really perform a request
	request.RequestURI = "/good"

	originUrl := GetOriginRequestUrl(request)

	assert.Equal(t, testRequestUrl, originUrl)
}

func TestGetOriginRequestUrlShouldReturnCorrectUrlWithForwardProtoConfig(t *testing.T) {
	testRequestUri := "example.com/good"
	request, createRequestErr := http.NewRequest("GET", "http://"+testRequestUri, nil)
	assert.Nil(t, createRequestErr)
	// specify request uri before really perform a request
	request.RequestURI = "/good"

	request.Header.Add("X-Forwarded-Proto", "https")

	originUrl := GetOriginRequestUrl(request)

	assert.Equal(t, "https://"+testRequestUri, originUrl)
}

func TestGetOriginRequestUrlShouldReturnCorrectUrlWithTlsConnection(t *testing.T) {
	testRequestUri := "example.com/good"
	request, createRequestErr := http.NewRequest("GET", "http://"+testRequestUri, nil)
	assert.Nil(t, createRequestErr)
	// specify request uri before really perform a request
	request.RequestURI = "/good"

	request.TLS = &tls.ConnectionState{}

	originUrl := GetOriginRequestUrl(request)

	assert.Equal(t, "https://"+testRequestUri, originUrl)
}

func TestBuildAccessTokenKeyShouldBuildCorrectly(t *testing.T) {
	tests := []struct {
		scopes   []string
		resource string
		result   string
	}{
		{[]string{}, "", "@"},
		{[]string{}, "http://api.example.com", "@http://api.example.com"},
		{[]string{"read", "write"}, "", "read write@"},
		{[]string{"read", "write"}, "http://api.example.com", "read write@http://api.example.com"},
		{[]string{"write", "read"}, "http://api.example.com", "read write@http://api.example.com"},
		{[]string{"read"}, "http://api.example.com", "read@http://api.example.com"},
	}

	for _, test := range tests {
		accessTokenKey := buildAccessTokenKey(test.scopes, test.resource, "")
		assert.Equal(t, test.result, accessTokenKey)
	}
}

func TestGetResourceFromAccessTokenShouldGetResourceCorrectly(t *testing.T) {
	testResource := "example.com"
	testAccessToken, createTokenError := createTestToken(testResource)
	assert.Nil(t, createTokenError)

	resource := getResourceFromAccessToken(testAccessToken)

	assert.Equal(t, testResource, resource)
}

func TestGetResourceFromAccessTokenShouldReturnEmptyStringIfResourceMissing(t *testing.T) {
	testResource := ""

	testAccessToken, createTokenErr := createTestToken(testResource)
	assert.Nil(t, createTokenErr)

	resource := getResourceFromAccessToken(testAccessToken)

	assert.Equal(t, testResource, resource)
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
