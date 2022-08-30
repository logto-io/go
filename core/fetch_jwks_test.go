package core

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jarcoal/httpmock"
)

func TestFetchJwks(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	jwksEndpoint := "http://example.com/oidc/jwks"
	mockResponse := `{` +
		`"keys": [{` +
		`"kty": "RSA",` +
		`"kid": "someKid",` +
		`"e": "someE",` +
		`"n": "someN"` +
		`}]` +
		`}`

	httpmock.RegisterResponder(
		"GET",
		jwksEndpoint,
		httpmock.NewStringResponder(200, mockResponse),
	)

	client := &http.Client{}

	jwks, fetchError := FetchJwks(client, jwksEndpoint)
	if fetchError != nil {
		t.Fatalf(fetchError.Error())
	}

	var expectedJwks JwksResponse
	unmarshalErr := json.Unmarshal([]byte(mockResponse), &expectedJwks)

	if unmarshalErr != nil {
		t.Fatalf(unmarshalErr.Error())
	}

	if !cmp.Equal(jwks, expectedJwks) {
		t.Fatalf("jwks does not match expected result")
	}
}
