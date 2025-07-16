package core

import (
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestFetchUserInfoWithClient(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	userInfoEndpoint := "http://example.com/oidc/jwks"
	mockResponse := `{` +
		`"sub": "sub",` +
		`"name": "name",` +
		`"username": "username",` +
		`"picture": "picture",` +
		`"email": "email@gmail.com",` +
		`"email_verified": true,` +
		`"phone_number": "12345678",` +
		`"phone_number_verified": true,` +
		`"custom_data": {"level": 1},` +
		`"identities": {"google": {"id": 1}},` +
		`"roles": ["role1", "role2"],` +
		`"organizations": ["org1"],` +
		`"organization_roles": ["viewer", "editor"],` +
		`"organization_data": [{"id": "org1", "name": "org1Name", "description": "org1Desc"}]` +
		`}`

	httpmock.RegisterResponder(
		"GET",
		userInfoEndpoint,
		httpmock.NewStringResponder(200, mockResponse),
	)

	// Test with custom HTTP client
	customClient := &http.Client{
		Timeout: 10 * time.Second,
	}

	userInfoResponse, fetchError := FetchUserInfoWithClient(customClient, userInfoEndpoint, "accessToken")
	assert.Nil(t, fetchError)

	var testUserInfoResponse UserInfoResponse
	unmarshalErr := json.Unmarshal([]byte(mockResponse), &testUserInfoResponse)
	assert.Nil(t, unmarshalErr)

	assert.Equal(t, testUserInfoResponse, userInfoResponse)
}

// TestFetchUserInfoDeprecated tests that the deprecated function delegates to the new one
func TestFetchUserInfoDeprecated(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	userInfoEndpoint := "http://example.com/oidc/jwks"
	mockResponse := `{` +
		`"sub": "sub",` +
		`"name": "name",` +
		`"username": "username"` +
		`}`

	httpmock.RegisterResponder(
		"GET",
		userInfoEndpoint,
		httpmock.NewStringResponder(200, mockResponse),
	)

	// Test that deprecated function still works by calling the new implementation
	userInfoResponse, fetchError := FetchUserInfo(userInfoEndpoint, "accessToken")
	assert.Nil(t, fetchError)
	assert.Equal(t, "sub", userInfoResponse.Sub)
	assert.Equal(t, "name", userInfoResponse.Name)
	assert.Equal(t, "username", userInfoResponse.Username)
}
