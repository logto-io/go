package core

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerifyAndParseCodeFromCallbackUriShouldGetCorrectCode(t *testing.T) {
	state := "state"
	testCode := "code"
	callbackUri := fmt.Sprintf("http://localhost:8080/callback?code=%s&state=%s", testCode, state)
	redirectUri := "http://localhost:8080/callback"

	parsedCode, err := VerifyAndParseCodeFromCallbackUri(callbackUri, redirectUri, state)

	assert.Nil(t, err)
	assert.Equal(t, testCode, parsedCode)
}

func TestVerifyAndParseCodeFromCallbackUriShouldGetCorrectCodeFromCustomSchemaCallbackUri(t *testing.T) {
	state := "state"
	testCode := "code"
	callbackUri := fmt.Sprintf("io.logto://localhost:8080/callback?code=%s&state=%s", testCode, state)
	redirectUri := "io.logto://localhost:8080/callback"

	parsedCode, err := VerifyAndParseCodeFromCallbackUri(callbackUri, redirectUri, state)

	assert.Nil(t, err)
	assert.Equal(t, testCode, parsedCode)
}

func TestVerifyAndParseCodeFromCallbackUriShouldReturnErrorIfCallbackUriNotMatchRedirectUri(t *testing.T) {
	state := "state"
	code := "code"
	callbackUri := fmt.Sprintf("http://localhost:8080/callback?code=%s&state=%s", code, state)
	redirectUri := "http://localhost:8080/callback2"

	_, err := VerifyAndParseCodeFromCallbackUri(callbackUri, redirectUri, state)

	assert.Equal(t, ErrCallbackUriNotMatchRedirectUri, err)
}

func TestVerifyAndParseCodeFromCallbackUriShouldReturnErrorIfStateNotMatch(t *testing.T) {
	state := "state"
	code := "code"
	callbackUri := fmt.Sprintf("http://localhost:8080/callback?code=%s&state=%s", code, state)
	redirectUri := "http://localhost:8080/callback"

	_, err := VerifyAndParseCodeFromCallbackUri(callbackUri, redirectUri, "state2")

	assert.Equal(t, ErrStateNotMatch, err)
}

func TestVerifyAndParseCodeFromCallbackUriShouldReturnErrorIfCallbackUriHasError(t *testing.T) {
	errorInUri := "intent uri error"
	errorDescriptionInUri := "intent uri error description"
	callbackUri := fmt.Sprintf("http://localhost:8080/callback?error=%s&error_description=%s", errorInUri, errorDescriptionInUri)
	redirectUri := "http://localhost:8080/callback"

	_, err := VerifyAndParseCodeFromCallbackUri(callbackUri, redirectUri, "state")

	assert.EqualError(t, err, fmt.Sprintf("%s: %s", errorInUri, errorDescriptionInUri))
}

func TestVerifyAndParseCodeFromCallbackUriShouldReturnErrorIfCodeNotFoundInCallbackUri(t *testing.T) {
	state := "state"
	callbackUri := fmt.Sprintf("http://localhost:8080/callback?state=%s", state)
	redirectUri := "http://localhost:8080/callback"

	_, err := VerifyAndParseCodeFromCallbackUri(callbackUri, redirectUri, state)

	assert.Equal(t, ErrCodeNotFoundInCallbackUri, err)
}
