package core

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestDecodeIdTokenShouldGetExpectedIdTokenClaims(t *testing.T) {
	now := time.Now()
	var expectedClaims = IdTokenClaims{
		Sub:       "1234567890",
		Aud:       "1234567890",
		Exp:       now.Add(time.Hour).Unix(),
		Iat:       now.Unix(),
		Iss:       "1234567890",
		AtHash:    "1234567890",
		Username:  "1234567890",
		Name:      "1234567890",
		Avatar:    "1234567890",
		RoleNames: []string{"1234567890"},
	}

	idToken, _, generateError := generateRsaSigningTestTokenAndCorrespondJwks(expectedClaims)

	if generateError != nil {
		t.Fatalf("Generate Test Token Error")
	}

	idTokenClaims, err := DecodeIdToken(idToken)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !cmp.Equal(expectedClaims, idTokenClaims) {
		t.Fatalf("Expected Claims: %v\nActual Claims: %v", expectedClaims, idTokenClaims)
	}
}

func TestDecodeIdTokenShouldReturnErrorWhenTokenIsInvalid(t *testing.T) {
	idToken := "invalid_token"

	_, err := DecodeIdToken(idToken)

	if err == nil {
		t.Fatalf("Expected error when decoding invalid token")
	}
}
