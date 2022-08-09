package core

import (
	"fmt"
	"testing"
)

func TestVerifyAndParseCodeFromCallbackUriShouldGetCorrectCode(t *testing.T) {
	state := "state"
	code := "code"
	callbackUri := fmt.Sprintf("http://localhost:8080/callback?code=%s&state=%s", code, state)
	redirectUri := "http://localhost:8080/callback"

	parsedCode, err := VerifyAndParseCodeFromCallbackUri(callbackUri, redirectUri, state)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if parsedCode != code {
		t.Fatalf("parsed code not match")
	}
}

func TestVerifyAndParseCodeFromCallbackUriShouldGetCorrectCodeFromCustomSchemaCallbackUri(t *testing.T) {
	state := "state"
	code := "code"
	callbackUri := fmt.Sprintf("io.logto://localhost:8080/callback?code=%s&state=%s", code, state)
	redirectUri := "io.logto://localhost:8080/callback"

	parsedCode, err := VerifyAndParseCodeFromCallbackUri(callbackUri, redirectUri, state)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if parsedCode != code {
		t.Fatalf("Expected code: %s, but got: %s", code, parsedCode)
	}
}

func TestVerifyAndParseCodeFromCallbackUriShouldReturnErrorIfCallbackUriNotMatchRedirectUri(t *testing.T) {
	state := "state"
	code := "code"
	callbackUri := fmt.Sprintf("http://localhost:8080/callback?code=%s&state=%s", code, state)
	redirectUri := "http://localhost:8080/callback2"

	_, err := VerifyAndParseCodeFromCallbackUri(callbackUri, redirectUri, state)

	if err == nil || err.Error() != ErrCallbackUriNotMatchRedirectUri.Error() {
		t.Fatalf("Expected error: %s, but got: %s", ErrCallbackUriNotMatchRedirectUri.Error(), err.Error())
	}
}

func TestVerifyAndParseCodeFromCallbackUriShouldReturnErrorIfStateNotMatch(t *testing.T) {
	state := "state"
	code := "code"
	callbackUri := fmt.Sprintf("http://localhost:8080/callback?code=%s&state=%s", code, state)
	redirectUri := "http://localhost:8080/callback"

	_, err := VerifyAndParseCodeFromCallbackUri(callbackUri, redirectUri, "state2")

	if err == nil || err.Error() != ErrStateNotMatch.Error() {
		t.Fatalf("Expected error: %s, but got: %s", ErrStateNotMatch.Error(), err.Error())
	}
}

func TestVerifyAndParseCodeFromCallbackUriShouldReturnErrorIfCallbackUriHasError(t *testing.T) {
	errorInUri := "intent uri error"
	errorDescriptionInUri := "intent uri error description"
	callbackUri := fmt.Sprintf("http://localhost:8080/callback?error=%s&error_description=%s", errorInUri, errorDescriptionInUri)
	redirectUri := "http://localhost:8080/callback"

	_, err := VerifyAndParseCodeFromCallbackUri(callbackUri, redirectUri, "state")

	if err == nil || err.Error() != fmt.Sprintf("%s: %s", errorInUri, errorDescriptionInUri) {
		t.Fatalf("Expected error: %s, but got: %s", fmt.Sprintf("%s: %s", errorInUri, errorDescriptionInUri), err.Error())
	}
}

func TestVerifyAndParseCodeFromCallbackUriShouldReturnErrorIfCodeNotFoundInCallbackUri(t *testing.T) {
	state := "state"
	callbackUri := fmt.Sprintf("http://localhost:8080/callback?state=%s", state)
	redirectUri := "http://localhost:8080/callback"

	_, err := VerifyAndParseCodeFromCallbackUri(callbackUri, redirectUri, state)

	if err == nil || err.Error() != ErrCodeNotFoundInCallbackUri.Error() {
		t.Fatalf("Expected error: %s, but got: %s", ErrCodeNotFoundInCallbackUri.Error(), err.Error())
	}
}
