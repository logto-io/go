package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateSignOutUriShouldGenerateCorrectUri(t *testing.T) {
	testSignOutUri := "https://example.com/logout?client_id=clientId&post_logout_redirect_uri=https%3A%2F%2Fexample.com%2Fcallback"
	signOutUri, generateSignOutUriErr := GenerateSignOutUri(&SignOutUriGenerationOptions{
		EndSessionEndpoint:    "https://example.com/logout",
		ClientId:              "clientId",
		PostLogoutRedirectUri: "https://example.com/callback",
	})

	assert.Nil(t, generateSignOutUriErr)
	assert.Equal(t, testSignOutUri, signOutUri)
}

func TestGenerateSignOutUriShouldGenerateCorrectUriWithoutPostLogoutRedirectUri(t *testing.T) {
	testSignOutUri := "https://example.com/logout?client_id=clientId"

	signOutUri, generateSignOutUriErr := GenerateSignOutUri(&SignOutUriGenerationOptions{
		EndSessionEndpoint: "https://example.com/logout",
		ClientId:           "clientId",
	})

	assert.Nil(t, generateSignOutUriErr)
	assert.Equal(t, testSignOutUri, signOutUri)
}
