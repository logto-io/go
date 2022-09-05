package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateSignOutUriShouldGenerateCorrectUri(t *testing.T) {
	testSignOutUri := "https://example.com/logout?id_token_hint=idToken&post_logout_redirect_uri=https://example.com/callback"
	signOutUri, generateSignOutUriErr := GenerateSignOutUri(&SignOutUriGenerationOptions{
		EndSessionEndpoint:    "https://example.com/logout",
		IdToken:               "idToken",
		PostLogoutRedirectUri: "https://example.com/callback",
	})

	assert.Nil(t, generateSignOutUriErr)
	assert.Equal(t, testSignOutUri, signOutUri)
}

func TestGenerateSignOutUriShouldGenerateCorrectUriWithoutPostLogoutRedirectUri(t *testing.T) {
	testSignOutUri := "https://example.com/logout?id_token_hint=idToken"

	signOutUri, generateSignOutUriErr := GenerateSignOutUri(&SignOutUriGenerationOptions{
		EndSessionEndpoint: "https://example.com/logout",
		IdToken:            "idToken",
	})

	assert.Nil(t, generateSignOutUriErr)
	assert.Equal(t, testSignOutUri, signOutUri)
}
