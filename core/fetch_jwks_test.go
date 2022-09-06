package core

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
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
	assert.Nil(t, fetchError)

	var testJwks JwksResponse
	unmarshalErr := json.Unmarshal([]byte(mockResponse), &testJwks)
	assert.Nil(t, unmarshalErr)

	assert.Equal(t, testJwks, jwks)
}
