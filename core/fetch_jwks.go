package core

import "net/http"

type JwksResponse struct {
	Keys []map[string]string `json:"keys"`
}

func FetchJwks(client *http.Client, jwksUri string) (JwksResponse, error) {
	response, requestErr := client.Get(jwksUri)

	if requestErr != nil {
		return JwksResponse{}, requestErr
	}

	defer response.Body.Close()

	var jwksResponse JwksResponse
	err := parseDataFromResponse(response, &jwksResponse)

	if err != nil {
		return JwksResponse{}, err
	}

	return jwksResponse, nil
}
