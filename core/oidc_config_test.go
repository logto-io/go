package core

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestFetchOidcConfig(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	endpoint := "http://example.com/oidc/.well-known/openid-configuration"
	mockResponse := `{` +
		`"authorization_endpoint": "http://example.com/oidc/authorize",` +
		`"token_endpoint": "http://example.com/oidc/token",` +
		`"end_session_endpoint": "http://example.com/oidc/logout",` +
		`"revocation_endpoint": "http://example.com/oidc/revoke",` +
		`"jwks_uri": "http://example.com/oidc/jwks",` +
		`"issuer": "http://example.com/oidc"` +
		`}`

	httpmock.RegisterResponder(
		"GET",
		endpoint,
		httpmock.NewStringResponder(200, mockResponse),
	)

	client := &http.Client{}
	oidcConfig, fetchOidcConfigErr := FetchOidcConfig(client, endpoint)
	assert.Nil(t, fetchOidcConfigErr)

	var testOidcConfig OidcConfigResponse
	unmarshalErr := json.Unmarshal([]byte(mockResponse), &testOidcConfig)
	assert.Nil(t, unmarshalErr)

	assert.Equal(t, testOidcConfig, oidcConfig)
}
