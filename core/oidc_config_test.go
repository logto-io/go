package core

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jarcoal/httpmock"
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
	config, err := FetchOidcConfig(client, endpoint)
	if err != nil {
		t.Fatalf(err.Error())
	}
	var expectedConfig OidcConfigResponse

	unmarshalErr := json.Unmarshal([]byte(mockResponse), &expectedConfig)

	if unmarshalErr != nil {
		t.Fatalf(unmarshalErr.Error())
	}

	if !cmp.Equal(config, expectedConfig) {
		t.Fatalf("config did not match expected config")
	}
}
