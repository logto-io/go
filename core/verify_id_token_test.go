package core

import (
	"testing"
	"time"
)

func TestVerifyIdToken(t *testing.T) {
	issuer := "issuer.logto.io"
	audience := "audience.logto.io"

	now := time.Now()

	idTokenClaims := IdTokenClaims{
		Sub:       "1234567890",
		Aud:       audience,
		Exp:       now.Add(time.Hour).Unix(),
		Iat:       now.Unix(),
		Iss:       issuer,
		AtHash:    "1234567890",
		Username:  "1234567890",
		Name:      "1234567890",
		Avatar:    "1234567890",
		RoleNames: []string{"1234567890"},
	}

	idToken, jwks, generateError := generateRsaSigningTestTokenAndCorrespondJwks(idTokenClaims)

	if generateError != nil {
		t.Fatalf(generateError.Error())
	}

	verifyIdTokenError := VerifyIdToken(idToken, audience, issuer, jwks)

	if verifyIdTokenError != nil {
		t.Fatalf(verifyIdTokenError.Error())
	}
}

func TestVerifyIdTokenShouldSupportES512FormatJwks(t *testing.T) {
	issuer := "issuer.logto.io"
	audience := "audience.logto.io"

	now := time.Now()

	idTokenClaims := IdTokenClaims{
		Sub:       "1234567890",
		Aud:       audience,
		Exp:       now.Add(time.Hour).Unix(),
		Iat:       now.Unix(),
		Iss:       issuer,
		AtHash:    "1234567890",
		Username:  "1234567890",
		Name:      "1234567890",
		Avatar:    "1234567890",
		RoleNames: []string{"1234567890"},
	}

	idToken, jwks, generateError := generateEcdsaSigningTestTokenAndCorrespondJwks(idTokenClaims)

	if generateError != nil {
		t.Fatalf(generateError.Error())
	}

	verifyIdTokenError := VerifyIdToken(idToken, audience, issuer, jwks)

	if verifyIdTokenError != nil {
		t.Fatalf(verifyIdTokenError.Error())
	}
}

func TestVerifyIdTokenShouldReturnErrorIfIssuedAtIsInThePast(t *testing.T) {
	now := time.Now()
	issuedAt := now.Unix() - ISSUED_AT_RESTRICTIONS - 10

	idTokenClaims := IdTokenClaims{
		Sub:       "1234567890",
		Aud:       "audience.logto.io",
		Exp:       now.Add(time.Hour).Unix(),
		Iat:       issuedAt,
		Iss:       "issuer.logto.io",
		AtHash:    "1234567890",
		Username:  "1234567890",
		Name:      "1234567890",
		Avatar:    "1234567890",
		RoleNames: []string{"1234567890"},
	}

	idToken, jwks, generateError := generateRsaSigningTestTokenAndCorrespondJwks(idTokenClaims)

	if generateError != nil {
		t.Fatalf(generateError.Error())
	}

	verifyIdTokenError := VerifyIdToken(idToken, "audience.logto.io", "issuer.logto.io", jwks)

	if verifyIdTokenError == nil || verifyIdTokenError.Error() != ErrTokenIssuedInThePast.Error() {
		t.Fatalf("Expected error when issuedAt is in the past")
	}
}

func TestVerifyIdTokenShouldReturnErrorIfIssuedAtIsInTheFuture(t *testing.T) {
	now := time.Now()
	issuedAt := now.Unix() + ISSUED_AT_RESTRICTIONS + 10

	idTokenClaims := IdTokenClaims{
		Sub:       "1234567890",
		Aud:       "audience.logto.io",
		Exp:       now.Add(time.Hour).Unix(),
		Iat:       issuedAt,
		Iss:       "issuer.logto.io",
		AtHash:    "1234567890",
		Username:  "1234567890",
		Name:      "1234567890",
		Avatar:    "1234567890",
		RoleNames: []string{"1234567890"},
	}

	idToken, jwks, generateError := generateRsaSigningTestTokenAndCorrespondJwks(idTokenClaims)

	if generateError != nil {
		t.Fatalf(generateError.Error())
	}

	verifyIdTokenError := VerifyIdToken(idToken, "audience.logto.io", "issuer.logto.io", jwks)

	if verifyIdTokenError == nil || verifyIdTokenError.Error() != ErrTokenIssuedInTheFuture.Error() {
		t.Fatalf("Expected error when issuedAt is in the future")
	}
}

func TestVerifyIdTokenShouldReturnErrorIfExpired(t *testing.T) {
	now := time.Now()

	idTokenClaims := IdTokenClaims{
		Sub:       "1234567890",
		Aud:       "audience.logto.io",
		Exp:       now.Unix() - 10,
		Iat:       now.Unix(),
		Iss:       "issuer.logto.io",
		AtHash:    "1234567890",
		Username:  "1234567890",
		Name:      "1234567890",
		Avatar:    "1234567890",
		RoleNames: []string{"1234567890"},
	}

	idToken, jwks, generateError := generateRsaSigningTestTokenAndCorrespondJwks(idTokenClaims)

	if generateError != nil {
		t.Fatalf(generateError.Error())
	}

	verifyIdTokenError := VerifyIdToken(idToken, "audience.logto.io", "issuer.logto.io", jwks)

	if verifyIdTokenError == nil || verifyIdTokenError.Error() != ErrTokenExpired.Error() {
		t.Fatalf("Expected error when expired")
	}
}

func TestVerifyIdTokenShouldReturnErrorIfIssuerIsNotMatched(t *testing.T) {
	issuer := "issuer.logto.io"

	now := time.Now()

	idTokenClaims := IdTokenClaims{
		Sub:       "1234567890",
		Aud:       "audience.logto.io",
		Exp:       now.Add(time.Hour).Unix(),
		Iat:       now.Unix(),
		Iss:       issuer,
		AtHash:    "1234567890",
		Username:  "1234567890",
		Name:      "1234567890",
		Avatar:    "1234567890",
		RoleNames: []string{"1234567890"},
	}

	idToken, jwks, generateError := generateRsaSigningTestTokenAndCorrespondJwks(idTokenClaims)

	if generateError != nil {
		t.Fatalf(generateError.Error())
	}

	anotherIssuer := "issuer.another.io"

	verifyIdTokenError := VerifyIdToken(idToken, "audience.logto.io", anotherIssuer, jwks)

	if verifyIdTokenError == nil || verifyIdTokenError.Error() != ErrTokenIssuerNotMatch.Error() {
		t.Fatalf("Expected error when issuer is not matched")
	}
}

func TestVerifyIdTokenShouldReturnErrorIfAudienceIsNotMatched(t *testing.T) {
	audience := "audience.logto.io"

	now := time.Now()

	idTokenClaims := IdTokenClaims{
		Sub:       "1234567890",
		Aud:       audience,
		Exp:       now.Add(time.Hour).Unix(),
		Iat:       now.Unix(),
		Iss:       "issuer.logto.io",
		AtHash:    "1234567890",
		Username:  "1234567890",
		Name:      "1234567890",
		Avatar:    "1234567890",
		RoleNames: []string{"1234567890"},
	}

	idToken, jwks, generateError := generateRsaSigningTestTokenAndCorrespondJwks(idTokenClaims)

	if generateError != nil {
		t.Fatalf(generateError.Error())
	}

	anotherAudience := "audience.another.io"

	verifyIdTokenError := VerifyIdToken(idToken, anotherAudience, "issuer.logto.io", jwks)

	if verifyIdTokenError == nil || verifyIdTokenError.Error() != ErrTokenAudienceNotMatch.Error() {
		t.Fatalf("Expected error when audience is not matched")
	}
}
