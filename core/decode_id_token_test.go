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

	// Decoding populates the raw claims, clear them to compare the modeled
	// fields against the literal expectation.
	assert.NotNil(t, idTokenClaims.rawClaims)
	idTokenClaims.rawClaims = nil
	assert.Equal(t, idTokenClaims, testClaims)
}

func TestDecodeIdTokenShouldSupportGettingAnyClaimFromRawClaims(t *testing.T) {
	now := time.Now()
	testClaims := map[string]any{
		"iss":          "1234567890",
		"sub":          "1234567890",
		"aud":          "1234567890",
		"exp":          now.Add(time.Hour).Unix(),
		"iat":          now.Unix(),
		"family_name":  "Doe",
		"given_name":   "John",
		"custom_claim": "custom_value",
	}

	idToken, _, generateError := generateRsaSigningTestTokenAndCorrespondJwks(testClaims)
	assert.Nil(t, generateError)

	idTokenClaims, decodeIdTokenErr := DecodeIdToken(idToken)
	assert.Nil(t, decodeIdTokenErr)

	// Standard profile claims are populated as modeled fields.
	assert.Equal(t, "Doe", idTokenClaims.FamilyName)
	assert.Equal(t, "John", idTokenClaims.GivenName)

	// Claims that are not modeled as struct fields are accessible via GetClaim.
	customClaim, ok := idTokenClaims.GetClaim("custom_claim")
	assert.True(t, ok)
	assert.Equal(t, "custom_value", customClaim)

	// Modeled claims are accessible via GetClaim as well.
	familyName, ok := idTokenClaims.GetClaim("family_name")
	assert.True(t, ok)
	assert.Equal(t, "Doe", familyName)

	// Absent claims are reported as not present.
	_, ok = idTokenClaims.GetClaim("nonexistent_claim")
	assert.False(t, ok)
}

func TestDecodeIdTokenShouldReturnErrorWhenTokenIsInvalid(t *testing.T) {
	idToken := "invalid_token"

	_, decodeIdTokenErr := DecodeIdToken(idToken)

	assert.NotNil(t, decodeIdTokenErr)
}
