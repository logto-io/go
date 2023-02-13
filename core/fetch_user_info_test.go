package core

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestFetchUserInfo(t *testing.T) {
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
		`"identities": {"google": {"id": 1}}` +
		`}`

	httpmock.RegisterResponder(
		"GET",
		userInfoEndpoint,
		httpmock.NewStringResponder(200, mockResponse),
	)

	client := &http.Client{}

	userInfoResponse, fetchError := FetchUserInfo(client, userInfoEndpoint, "dummyAccessToken")
	assert.Nil(t, fetchError)

	var testUserInfoResponse UserInfoResponse
	unmarshalErr := json.Unmarshal([]byte(mockResponse), &testUserInfoResponse)
	assert.Nil(t, unmarshalErr)

	assert.Equal(t, testUserInfoResponse, userInfoResponse)
}
