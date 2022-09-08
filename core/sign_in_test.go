package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateSignInUriShouldGenerateCorrectUri(t *testing.T) {
	testSignInUri := "https://example.com/authorize?client_id=clientId&code_challenge=codeChallenge&code_challenge_method=S256&prompt=login&redirect_uri=https://example.com/callback&resource=resourceA&resource=resourceB&response_type=code&scope=openid offline_access profile&state=state"

	signInUri, generateSignInUriErr := GenerateSignInUri(&SignInUriGenerationOptions{
		AuthorizationEndpoint: "https://example.com/authorize",
		ClientId:              "clientId",
		RedirectUri:           "https://example.com/callback",
		CodeChallenge:         "codeChallenge",
		State:                 "state",
		Scopes:                []string{"openid", "offline_access", "profile"},
		Resources:             []string{"resourceA", "resourceB"},
		Prompt:                "login",
	})

	assert.Nil(t, generateSignInUriErr)
	assert.Equal(t, testSignInUri, signInUri)
}

func TestGenerateSignInUriShouldContainReservedScopesByDefault(t *testing.T) {
	testSignInUri := "https://example.com/authorize?client_id=clientId&code_challenge=codeChallenge&code_challenge_method=S256&prompt=consent&redirect_uri=https://example.com/callback&resource=resourceA&resource=resourceB&response_type=code&scope=openid offline_access profile&state=state"

	signInUri, generateSignInUriErr := GenerateSignInUri(&SignInUriGenerationOptions{
		AuthorizationEndpoint: "https://example.com/authorize",
		ClientId:              "clientId",
		RedirectUri:           "https://example.com/callback",
		CodeChallenge:         "codeChallenge",
		State:                 "state",
		Resources:             []string{"resourceA", "resourceB"},
		Prompt:                "consent",
	})

	assert.Nil(t, generateSignInUriErr)
	assert.Equal(t, testSignInUri, signInUri)
}

func TestGenerateSignInUriShouldContainReservedScopesAndExtraScopes(t *testing.T) {
	testSignInUri := "https://example.com/authorize?client_id=clientId&code_challenge=codeChallenge&code_challenge_method=S256&prompt=consent&redirect_uri=https://example.com/callback&resource=resourceA&resource=resourceB&response_type=code&scope=openid offline_access profile extra_scope&state=state"

	signInUri, generateSignInUriErr := GenerateSignInUri(&SignInUriGenerationOptions{
		AuthorizationEndpoint: "https://example.com/authorize",
		ClientId:              "clientId",
		RedirectUri:           "https://example.com/callback",
		CodeChallenge:         "codeChallenge",
		State:                 "state",
		Scopes:                []string{"openid", "offline_access", "profile", "extra_scope"},
		Resources:             []string{"resourceA", "resourceB"},
		Prompt:                "consent",
	})

	assert.Nil(t, generateSignInUriErr)
	assert.Equal(t, testSignInUri, signInUri)
}

func TestGenerateSignInUriShouldGenerateUriWithConsentAsThePromptValue(t *testing.T) {
	testSignInUri := "https://example.com/authorize?client_id=clientId&code_challenge=codeChallenge&code_challenge_method=S256&prompt=consent&redirect_uri=https://example.com/callback&resource=resourceA&resource=resourceB&response_type=code&scope=openid offline_access profile&state=state"

	signInUri, generateSignInUriErr := GenerateSignInUri(&SignInUriGenerationOptions{
		AuthorizationEndpoint: "https://example.com/authorize",
		ClientId:              "clientId",
		RedirectUri:           "https://example.com/callback",
		CodeChallenge:         "codeChallenge",
		State:                 "state",
		Scopes:                []string{"openid", "offline_access", "profile"},
		Resources:             []string{"resourceA", "resourceB"},
	})

	assert.Nil(t, generateSignInUriErr)
	assert.Equal(t, testSignInUri, signInUri)
}

func TestGenerateSignInUriShouldNotContainResourcesIfNoResourcesAreProvided(t *testing.T) {
	testSignInUri := "https://example.com/authorize?client_id=clientId&code_challenge=codeChallenge&code_challenge_method=S256&prompt=consent&redirect_uri=https://example.com/callback&response_type=code&scope=openid offline_access profile&state=state"

	signInUri, generateSignInUriErr := GenerateSignInUri(&SignInUriGenerationOptions{
		AuthorizationEndpoint: "https://example.com/authorize",
		ClientId:              "clientId",
		RedirectUri:           "https://example.com/callback",
		CodeChallenge:         "codeChallenge",
		State:                 "state",
		Scopes:                []string{"openid", "offline_access", "profile"},
	})

	assert.Nil(t, generateSignInUriErr)
	assert.Equal(t, testSignInUri, signInUri)
}
