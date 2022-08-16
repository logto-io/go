package core

import (
	"net/http"
)

func FetchOidcConfig(client *http.Client, endpoint string) (OidcConfigResponse, error) {
	response, fetchErr := client.Get(endpoint)

	if fetchErr != nil {
		return OidcConfigResponse{}, fetchErr
	}

	defer response.Body.Close()

	var config OidcConfigResponse
	err := parseDataFromResponse(response, &config)

	if err != nil {
		return OidcConfigResponse{}, err
	}

	return config, nil
}
