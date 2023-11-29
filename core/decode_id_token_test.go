package core

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDecodeIdTokenShouldGetExpectedIdTokenClaims(t *testing.T) {
	now := time.Now()
	testClaims := IdTokenClaims{
		Sub:      "1234567890",
		Aud:      "1234567890",
		Exp:      now.Add(time.Hour).Unix(),
		Iat:      now.Unix(),
		Iss:      "1234567890",
		AtHash:   "1234567890",
		Username: "1234567890",
		Name:     "1234567890",
		Picture:  "picture",
	}

	idToken, _, generateError := generateRsaSigningTestTokenAndCorrespondJwks(testClaims)
	assert.Nil(t, generateError)

	idTokenClaims, decodeIdTokenErr := DecodeIdToken(idToken)
	assert.Nil(t, decodeIdTokenErr)

	assert.Equal(t, idTokenClaims, testClaims)
}

func TestDecodeIdTokenShouldReturnErrorWhenTokenIsInvalid(t *testing.T) {
	idToken := "invalid_token"

	_, decodeIdTokenErr := DecodeIdToken(idToken)

	assert.NotNil(t, decodeIdTokenErr)
}
