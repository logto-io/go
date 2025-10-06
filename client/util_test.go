package client

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"net/http"
	"testing"

	"github.com/go-jose/go-jose/v4"
	"github.com/go-jose/go-jose/v4/jwt"
	"github.com/stretchr/testify/assert"
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

	token, buildTokenError := builder.Claims(claims).Serialize()

	if buildTokenError != nil {
		return "", buildTokenError
	}

	return token, nil
}

func TestGetForwaredRequestUrlShouldReturnXForwardedIfPresent(t *testing.T) {
	req, err := http.NewRequest("GET", "http://example.com/path?query=1", nil)
	assert.Nil(t, err)
	// Ensure RequestURI is set like in real servers
	req.RequestURI = "/path?query=1"

	req.Header.Add("X-Forwarded-Host", "forwarded.example.com")
	req.Header.Add("X-Forwarded-Url", "/forwarded-path?query=2")
	req.Header.Add("X-Forwarded-Proto", "https")

	url := getForwaredRequestUrl(req)

	assert.Equal(t, "https://forwarded.example.com/forwarded-path?query=2", url)
}

func TestGetForwaredRequestHostShouldFallbackToRequestHost(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://example.com/", nil)
	req.RequestURI = "/"

	host := getForwaredRequestHost(req)

	assert.Equal(t, "example.com", host)

	req.Header.Add("X-Forwarded-Host", "proxied.example.com")
	host = getForwaredRequestHost(req)
	assert.Equal(t, "proxied.example.com", host)
}

func TestGetForwaredRequestRequestUriShouldFallbackToRequestURI(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://example.com/api", nil)
	req.RequestURI = "/api"

	uri := getForwaredRequestRequestUri(req)
	assert.Equal(t, "/api", uri)

	req.Header.Add("X-Forwarded-Url", "/proxied/api?x=1")
	uri = getForwaredRequestRequestUri(req)
	assert.Equal(t, "/proxied/api?x=1", uri)
}
