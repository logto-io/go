package core

import (
	"fmt"
	"net/http"
)

func FetchUserInfo(client *http.Client, userInfoEndpoint, accessToken string) (UserInfoResponse, error) {
	req, err := http.NewRequest("GET", userInfoEndpoint, nil)

	if err != nil {
		return UserInfoResponse{}, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	response, requestErr := client.Do(req)

	if requestErr != nil {
		return UserInfoResponse{}, requestErr
	}

	defer response.Body.Close()

	var userInfoResponse UserInfoResponse
	err = parseDataFromResponse(response, &userInfoResponse)

	if err != nil {
		return UserInfoResponse{}, err
	}

	return userInfoResponse, nil
}
