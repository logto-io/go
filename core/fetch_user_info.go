package core

import "net/http"

// FetchUserInfoWithClient fetches user info using a custom HTTP client.
// This function allows you to use a custom HTTP client for observability,
// tracing, or other custom configurations.
func FetchUserInfoWithClient(client *http.Client, userInfoEndpoint, accessToken string) (UserInfoResponse, error) {
	request, createRequestErr := http.NewRequest("GET", userInfoEndpoint, nil)

	if createRequestErr != nil {
		return UserInfoResponse{}, createRequestErr
	}

	request.Header.Add("Authorization", "Bearer "+accessToken)

	response, requestErr := client.Do(request)

	if requestErr != nil {
		return UserInfoResponse{}, requestErr
	}

	defer response.Body.Close()

	var userInfoResponse UserInfoResponse
	err := parseDataFromResponse(response, &userInfoResponse)

	if err != nil {
		return UserInfoResponse{}, err
	}

	return userInfoResponse, nil
}

// FetchUserInfo fetches user info using the default HTTP client.
// Deprecated: Use FetchUserInfoWithClient instead for better flexibility and observability support.
func FetchUserInfo(userInfoEndpoint, accessToken string) (UserInfoResponse, error) {
	return FetchUserInfoWithClient(&http.Client{}, userInfoEndpoint, accessToken)
}
