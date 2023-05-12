package core

import "net/http"

func FetchUserInfo(userInfoEndpoint, accessToken string) (UserInfoResponse, error) {
	client := &http.Client{}

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
