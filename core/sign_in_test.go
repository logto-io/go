package core

import (
	"testing"
)

func TestGenerateSignInUriShouldGenerateCorrectUri(t *testing.T) {
	expectedSignInUri := "https://example.com/authorize?client_id=clientId&code_challenge=codeChallenge&code_challenge_method=S256&prompt=consent&redirect_uri=https://example.com/callback&resource=resourceA&resource=resourceB&response_type=code&scope=openid offline_access profile&state=state"

	signInUri, err := GenerateSignInUri(&SignInUriGenerationOptions{
		AuthorizationEndpoint: "https://example.com/authorize",
		ClientId:              "clientId",
		RedirectUri:           "https://example.com/callback",
		CodeChallenge:         "codeChallenge",
		State:                 "state",
		Scopes:                []string{"openid", "offline_access", "profile"},
		Resources:             []string{"resourceA", "resourceB"},
		Prompt:                "consent",
	})

	if err != nil {
		t.Fatal(err)
	}

	if signInUri != expectedSignInUri {
		t.Errorf("Expected %s, got %s", expectedSignInUri, signInUri)
	}
}

func TestGenerateSignInUriShouldContainReservedScopesByDefault(t *testing.T) {
	expectedSignInUri := "https://example.com/authorize?client_id=clientId&code_challenge=codeChallenge&code_challenge_method=S256&prompt=consent&redirect_uri=https://example.com/callback&resource=resourceA&resource=resourceB&response_type=code&scope=openid offline_access profile&state=state"

	signInUri, err := GenerateSignInUri(&SignInUriGenerationOptions{
		AuthorizationEndpoint: "https://example.com/authorize",
		ClientId:              "clientId",
		RedirectUri:           "https://example.com/callback",
		CodeChallenge:         "codeChallenge",
		State:                 "state",
		Resources:             []string{"resourceA", "resourceB"},
		Prompt:                "consent",
	})

	if err != nil {
		t.Fatal(err)
	}

	if signInUri != expectedSignInUri {
		t.Errorf("Expected %s, got %s", expectedSignInUri, signInUri)
	}
}

func TestGenerateSignInUriShouldContainReservedScopesAndExtraScopes(t *testing.T) {
	expectedSignInUri := "https://example.com/authorize?client_id=clientId&code_challenge=codeChallenge&code_challenge_method=S256&prompt=consent&redirect_uri=https://example.com/callback&resource=resourceA&resource=resourceB&response_type=code&scope=openid offline_access profile extra_scope&state=state"

	signInUri, err := GenerateSignInUri(&SignInUriGenerationOptions{
		AuthorizationEndpoint: "https://example.com/authorize",
		ClientId:              "clientId",
		RedirectUri:           "https://example.com/callback",
		CodeChallenge:         "codeChallenge",
		State:                 "state",
		Scopes:                []string{"openid", "offline_access", "profile", "extra_scope"},
		Resources:             []string{"resourceA", "resourceB"},
		Prompt:                "consent",
	})

	if err != nil {
		t.Fatal(err)
	}

	if signInUri != expectedSignInUri {
		t.Errorf("Expected %s, got %s", expectedSignInUri, signInUri)
	}
}

func TestGenerateSignInUriShouldGenerateUriWithConsentAsThePromptValue(t *testing.T) {
	expectedSignInUri := "https://example.com/authorize?client_id=clientId&code_challenge=codeChallenge&code_challenge_method=S256&prompt=consent&redirect_uri=https://example.com/callback&resource=resourceA&resource=resourceB&response_type=code&scope=openid offline_access profile&state=state"

	signInUri, err := GenerateSignInUri(&SignInUriGenerationOptions{
		AuthorizationEndpoint: "https://example.com/authorize",
		ClientId:              "clientId",
		RedirectUri:           "https://example.com/callback",
		CodeChallenge:         "codeChallenge",
		State:                 "state",
		Scopes:                []string{"openid", "offline_access", "profile"},
		Resources:             []string{"resourceA", "resourceB"},
	})

	if err != nil {
		t.Fatal(err)
	}

	if signInUri != expectedSignInUri {
		t.Errorf("Expected %s, got %s", expectedSignInUri, signInUri)
	}
}

func TestGenerateSignInUriShouldNotContainsResourcesIfNoResourcesAreProvided(t *testing.T) {
	expectedSignInUri := "https://example.com/authorize?client_id=clientId&code_challenge=codeChallenge&code_challenge_method=S256&prompt=consent&redirect_uri=https://example.com/callback&response_type=code&scope=openid offline_access profile&state=state"

	signInUri, err := GenerateSignInUri(&SignInUriGenerationOptions{
		AuthorizationEndpoint: "https://example.com/authorize",
		ClientId:              "clientId",
		RedirectUri:           "https://example.com/callback",
		CodeChallenge:         "codeChallenge",
		State:                 "state",
		Scopes:                []string{"openid", "offline_access", "profile"},
	})

	if err != nil {
		t.Fatal(err)
	}

	if signInUri != expectedSignInUri {
		t.Errorf("Expected %s, got %s", expectedSignInUri, signInUri)
	}
}
