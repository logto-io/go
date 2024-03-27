package core

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestVerifyIdToken(t *testing.T) {
	issuer := "issuer.logto.io"
	audience := "audience.logto.io"

	now := time.Now()

	idTokenClaims := IdTokenClaims{
		Sub:      "1234567890",
		Aud:      audience,
		Exp:      now.Add(time.Hour).Unix(),
		Iat:      now.Unix(),
		Iss:      issuer,
		AtHash:   "1234567890",
		Username: "1234567890",
		Name:     "1234567890",
		Picture:  "1234567890",
	}

	idToken, jwks, generateError := generateRsaSigningTestTokenAndCorrespondJwks(idTokenClaims)

	assert.Nil(t, generateError)

	verifyIdTokenError := VerifyIdToken(idToken, audience, issuer, &jwks)

	assert.Nil(t, verifyIdTokenError)
}

func TestVerifyIdTokenShouldSupportES512FormatJwks(t *testing.T) {
	issuer := "issuer.logto.io"
	audience := "audience.logto.io"

	now := time.Now()

	idTokenClaims := IdTokenClaims{
		Sub:      "1234567890",
		Aud:      audience,
		Exp:      now.Add(time.Hour).Unix(),
		Iat:      now.Unix(),
		Iss:      issuer,
		AtHash:   "1234567890",
		Username: "1234567890",
		Name:     "1234567890",
		Picture:  "1234567890",
	}

	idToken, jwks, generateError := generateEcdsaSigningTestTokenAndCorrespondJwks(idTokenClaims)
	assert.Nil(t, generateError)

	verifyIdTokenError := VerifyIdToken(idToken, audience, issuer, &jwks)

	assert.Nil(t, verifyIdTokenError)
}

func TestVerifyIdTokenShouldReturnErrorIfIssuedAtIsInThePast(t *testing.T) {
	now := time.Now()
	issuedAt := now.Unix() - ISSUED_AT_RESTRICTIONS - 10

	idTokenClaims := IdTokenClaims{
		Sub:      "1234567890",
		Aud:      "audience.logto.io",
		Exp:      now.Add(time.Hour).Unix(),
		Iat:      issuedAt,
		Iss:      "issuer.logto.io",
		AtHash:   "1234567890",
		Username: "1234567890",
		Name:     "1234567890",
		Picture:  "1234567890",
	}

	idToken, jwks, generateError := generateRsaSigningTestTokenAndCorrespondJwks(idTokenClaims)
	assert.Nil(t, generateError)

	verifyIdTokenError := VerifyIdToken(idToken, "audience.logto.io", "issuer.logto.io", &jwks)

	assert.Equal(t, ErrTokenIssuedInThePast, verifyIdTokenError)
}

func TestVerifyIdTokenShouldReturnErrorIfIssuedAtIsInTheFuture(t *testing.T) {
	now := time.Now()
	issuedAt := now.Unix() + ISSUED_AT_RESTRICTIONS + 10

	idTokenClaims := IdTokenClaims{
		Sub:      "1234567890",
		Aud:      "audience.logto.io",
		Exp:      now.Add(time.Hour).Unix(),
		Iat:      issuedAt,
		Iss:      "issuer.logto.io",
		AtHash:   "1234567890",
		Username: "1234567890",
		Name:     "1234567890",
		Picture:  "1234567890",
	}

	idToken, jwks, generateError := generateRsaSigningTestTokenAndCorrespondJwks(idTokenClaims)
	assert.Nil(t, generateError)

	verifyIdTokenError := VerifyIdToken(idToken, "audience.logto.io", "issuer.logto.io", &jwks)

	assert.Equal(t, ErrTokenIssuedInTheFuture, verifyIdTokenError)
}

func TestVerifyIdTokenShouldReturnErrorIfExpired(t *testing.T) {
	now := time.Now()

	idTokenClaims := IdTokenClaims{
		Sub:      "1234567890",
		Aud:      "audience.logto.io",
		Exp:      now.Unix() - 10,
		Iat:      now.Unix(),
		Iss:      "issuer.logto.io",
		AtHash:   "1234567890",
		Username: "1234567890",
		Name:     "1234567890",
		Picture:  "1234567890",
	}

	idToken, jwks, generateError := generateRsaSigningTestTokenAndCorrespondJwks(idTokenClaims)
	assert.Nil(t, generateError)

	verifyIdTokenError := VerifyIdToken(idToken, "audience.logto.io", "issuer.logto.io", &jwks)

	assert.Equal(t, ErrTokenExpired, verifyIdTokenError)
}

func TestVerifyIdTokenShouldReturnErrorIfIssuerIsNotMatched(t *testing.T) {
	issuer := "issuer.logto.io"

	now := time.Now()

	idTokenClaims := IdTokenClaims{
		Sub:      "1234567890",
		Aud:      "audience.logto.io",
		Exp:      now.Add(time.Hour).Unix(),
		Iat:      now.Unix(),
		Iss:      issuer,
		AtHash:   "1234567890",
		Username: "1234567890",
		Name:     "1234567890",
		Picture:  "1234567890",
	}

	idToken, jwks, generateError := generateRsaSigningTestTokenAndCorrespondJwks(idTokenClaims)
	assert.Nil(t, generateError)

	anotherIssuer := "issuer.another.io"

	verifyIdTokenError := VerifyIdToken(idToken, "audience.logto.io", anotherIssuer, &jwks)

	assert.Equal(t, ErrTokenIssuerNotMatch, verifyIdTokenError)
}

func TestVerifyIdTokenShouldReturnErrorIfAudienceIsNotMatched(t *testing.T) {
	audience := "audience.logto.io"

	now := time.Now()

	idTokenClaims := IdTokenClaims{
		Sub:      "1234567890",
		Aud:      audience,
		Exp:      now.Add(time.Hour).Unix(),
		Iat:      now.Unix(),
		Iss:      "issuer.logto.io",
		AtHash:   "1234567890",
		Username: "1234567890",
		Name:     "1234567890",
		Picture:  "1234567890",
	}

	idToken, jwks, generateError := generateRsaSigningTestTokenAndCorrespondJwks(idTokenClaims)
	assert.Nil(t, generateError)

	anotherAudience := "audience.another.io"

	verifyIdTokenError := VerifyIdToken(idToken, anotherAudience, "issuer.logto.io", &jwks)

	assert.Equal(t, ErrTokenAudienceNotMatch, verifyIdTokenError)
}
