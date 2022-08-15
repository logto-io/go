package core

import "testing"

func TestGenerateSignOutUriShouldGenerateCorrectUri(t *testing.T) {
	expectedSignOutUri := "https://example.com/logout?id_token_hint=idToken&post_logout_redirect_uri=https://example.com/callback"
	signOutUri, err := GenerateSignOutUri(&SignOutUriGenerationOptions{
		EndSessionEndpoint:    "https://example.com/logout",
		IdToken:               "idToken",
		PostLogoutRedirectUri: "https://example.com/callback",
	})

	if err != nil {
		t.Fatal(err)
	}

	if signOutUri != expectedSignOutUri {
		t.Errorf("Expected %s, got %s", expectedSignOutUri, signOutUri)
	}
}

func TestGenerateSignOutUriShouldGenerateCorrectUriWithoutPostLogoutRedirectUri(t *testing.T) {
	expectedSignOutUri := "https://example.com/logout?id_token_hint=idToken"

	signOutUri, err := GenerateSignOutUri(&SignOutUriGenerationOptions{
		EndSessionEndpoint: "https://example.com/logout",
		IdToken:            "idToken",
	})
	if err != nil {
		t.Fatal(err)
	}

	if signOutUri != expectedSignOutUri {
		t.Errorf("Expected %s, got %s", expectedSignOutUri, signOutUri)
	}
}
