package core

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserInfoResponseUnmarshalJsonShouldPopulateStandardProfileClaims(t *testing.T) {
	rawUserInfoResponse := `{
		"sub": "sub",
		"name": "name",
		"family_name": "Doe",
		"given_name": "John",
		"middle_name": "Middle",
		"nickname": "Johnny",
		"preferred_username": "johnny",
		"profile": "https://example.com/johnny",
		"website": "https://example.com",
		"gender": "male",
		"birthdate": "2000-01-01",
		"zoneinfo": "Europe/Paris",
		"locale": "en-US",
		"created_at": 1700000000000,
		"updated_at": 1700000001000,
		"custom_claim": "custom_value"
	}`

	var userInfoResponse UserInfoResponse
	unmarshalErr := json.Unmarshal([]byte(rawUserInfoResponse), &userInfoResponse)
	assert.Nil(t, unmarshalErr)

	assert.Equal(t, "sub", userInfoResponse.Sub)
	assert.Equal(t, "name", userInfoResponse.Name)
	assert.Equal(t, "Doe", userInfoResponse.FamilyName)
	assert.Equal(t, "John", userInfoResponse.GivenName)
	assert.Equal(t, "Middle", userInfoResponse.MiddleName)
	assert.Equal(t, "Johnny", userInfoResponse.Nickname)
	assert.Equal(t, "johnny", userInfoResponse.PreferredUsername)
	assert.Equal(t, "https://example.com/johnny", userInfoResponse.Profile)
	assert.Equal(t, "https://example.com", userInfoResponse.Website)
	assert.Equal(t, "male", userInfoResponse.Gender)
	assert.Equal(t, "2000-01-01", userInfoResponse.Birthdate)
	assert.Equal(t, "Europe/Paris", userInfoResponse.Zoneinfo)
	assert.Equal(t, "en-US", userInfoResponse.Locale)
	assert.Equal(t, int64(1700000000000), userInfoResponse.CreatedAt)
	assert.Equal(t, int64(1700000001000), userInfoResponse.UpdatedAt)
}

func TestUserInfoResponseUnmarshalJsonShouldLeaveAbsentClaimsAsZeroValues(t *testing.T) {
	rawUserInfoResponse := `{"sub": "sub"}`

	var userInfoResponse UserInfoResponse
	unmarshalErr := json.Unmarshal([]byte(rawUserInfoResponse), &userInfoResponse)
	assert.Nil(t, unmarshalErr)

	assert.Equal(t, "sub", userInfoResponse.Sub)
	assert.Empty(t, userInfoResponse.FamilyName)
	assert.Empty(t, userInfoResponse.GivenName)
	assert.Empty(t, userInfoResponse.Nickname)
	assert.Empty(t, userInfoResponse.Locale)
	assert.Zero(t, userInfoResponse.CreatedAt)
	assert.Zero(t, userInfoResponse.UpdatedAt)

	_, ok := userInfoResponse.GetClaim("family_name")
	assert.False(t, ok)
}

func TestUserInfoResponseGetClaimShouldReturnAnyClaim(t *testing.T) {
	rawUserInfoResponse := `{"sub": "sub", "custom_claim": "custom_value", "custom_data": {"level": 1}}`

	var userInfoResponse UserInfoResponse
	unmarshalErr := json.Unmarshal([]byte(rawUserInfoResponse), &userInfoResponse)
	assert.Nil(t, unmarshalErr)

	// Claims that are not modeled as struct fields are accessible via GetClaim.
	customClaim, ok := userInfoResponse.GetClaim("custom_claim")
	assert.True(t, ok)
	assert.Equal(t, "custom_value", customClaim)

	// Modeled claims are accessible via GetClaim as well.
	sub, ok := userInfoResponse.GetClaim("sub")
	assert.True(t, ok)
	assert.Equal(t, "sub", sub)

	customData, ok := userInfoResponse.GetClaim("custom_data")
	assert.True(t, ok)
	assert.Equal(t, map[string]any{"level": float64(1)}, customData)

	// Absent claims are reported as not present.
	_, ok = userInfoResponse.GetClaim("nonexistent_claim")
	assert.False(t, ok)
}

func TestIdTokenClaimsUnmarshalJsonShouldPopulateStandardProfileClaims(t *testing.T) {
	rawIdTokenClaims := `{
		"iss": "iss",
		"sub": "sub",
		"aud": "aud",
		"family_name": "Doe",
		"given_name": "John",
		"middle_name": "Middle",
		"nickname": "Johnny",
		"preferred_username": "johnny",
		"profile": "https://example.com/johnny",
		"website": "https://example.com",
		"gender": "male",
		"birthdate": "2000-01-01",
		"zoneinfo": "Europe/Paris",
		"locale": "en-US",
		"created_at": 1700000000000,
		"updated_at": 1700000001000,
		"custom_claim": "custom_value"
	}`

	var idTokenClaims IdTokenClaims
	unmarshalErr := json.Unmarshal([]byte(rawIdTokenClaims), &idTokenClaims)
	assert.Nil(t, unmarshalErr)

	assert.Equal(t, "iss", idTokenClaims.Iss)
	assert.Equal(t, "sub", idTokenClaims.Sub)
	assert.Equal(t, "aud", idTokenClaims.Aud)
	assert.Equal(t, "Doe", idTokenClaims.FamilyName)
	assert.Equal(t, "John", idTokenClaims.GivenName)
	assert.Equal(t, "Middle", idTokenClaims.MiddleName)
	assert.Equal(t, "Johnny", idTokenClaims.Nickname)
	assert.Equal(t, "johnny", idTokenClaims.PreferredUsername)
	assert.Equal(t, "https://example.com/johnny", idTokenClaims.Profile)
	assert.Equal(t, "https://example.com", idTokenClaims.Website)
	assert.Equal(t, "male", idTokenClaims.Gender)
	assert.Equal(t, "2000-01-01", idTokenClaims.Birthdate)
	assert.Equal(t, "Europe/Paris", idTokenClaims.Zoneinfo)
	assert.Equal(t, "en-US", idTokenClaims.Locale)
	assert.Equal(t, int64(1700000000000), idTokenClaims.CreatedAt)
	assert.Equal(t, int64(1700000001000), idTokenClaims.UpdatedAt)
}

func TestIdTokenClaimsUnmarshalJsonShouldLeaveAbsentClaimsAsZeroValues(t *testing.T) {
	rawIdTokenClaims := `{"iss": "iss", "sub": "sub"}`

	var idTokenClaims IdTokenClaims
	unmarshalErr := json.Unmarshal([]byte(rawIdTokenClaims), &idTokenClaims)
	assert.Nil(t, unmarshalErr)

	assert.Equal(t, "iss", idTokenClaims.Iss)
	assert.Equal(t, "sub", idTokenClaims.Sub)
	assert.Empty(t, idTokenClaims.FamilyName)
	assert.Empty(t, idTokenClaims.GivenName)
	assert.Empty(t, idTokenClaims.Nickname)
	assert.Empty(t, idTokenClaims.Locale)
	assert.Zero(t, idTokenClaims.CreatedAt)
	assert.Zero(t, idTokenClaims.UpdatedAt)

	_, ok := idTokenClaims.GetClaim("family_name")
	assert.False(t, ok)
}

func TestIdTokenClaimsGetClaimShouldReturnAnyClaim(t *testing.T) {
	rawIdTokenClaims := `{"iss": "iss", "sub": "sub", "custom_claim": "custom_value"}`

	var idTokenClaims IdTokenClaims
	unmarshalErr := json.Unmarshal([]byte(rawIdTokenClaims), &idTokenClaims)
	assert.Nil(t, unmarshalErr)

	// Claims that are not modeled as struct fields are accessible via GetClaim.
	customClaim, ok := idTokenClaims.GetClaim("custom_claim")
	assert.True(t, ok)
	assert.Equal(t, "custom_value", customClaim)

	// Modeled claims are accessible via GetClaim as well.
	iss, ok := idTokenClaims.GetClaim("iss")
	assert.True(t, ok)
	assert.Equal(t, "iss", iss)

	// Absent claims are reported as not present.
	_, ok = idTokenClaims.GetClaim("nonexistent_claim")
	assert.False(t, ok)
}
