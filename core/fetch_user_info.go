package core

import "net/http"

func FetchUserInfo(client *http.Client, userInfoEndpoint, accessToken string) (UserInfoResponse, error) {
	response, requestErr := client.Get(userInfoEndpoint)

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
