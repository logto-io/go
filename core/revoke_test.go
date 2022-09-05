package core

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestRevokeShouldRevokeSuccessfully(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	revocationEndpoint := "http://example.com/oidc/revoke"
	mockResponse := `{}`

	httpmock.RegisterResponder(
		"POST",
		revocationEndpoint,
		httpmock.NewStringResponder(200, mockResponse),
	)

	client := &http.Client{}
	options := &RevocationOptions{
		RevocationEndpoint: revocationEndpoint,
		ClientId:           "clientId",
		Token:              "token",
	}

	revokeErr := Revoke(client, options)

	assert.Nil(t, revokeErr)
}
