package core

import (
	"context"
	"net/http"
)

func FetchOidcConfig(client *http.Client, endpoint string) (OidcConfigResponse, error) {
	return FetchOidcConfigContext(context.Background(), client, endpoint)
}

// FetchOidcConfigContext is like FetchOidcConfig but binds the request to ctx.
func FetchOidcConfigContext(ctx context.Context, client *http.Client, endpoint string) (OidcConfigResponse, error) {
	request, createRequestErr := http.NewRequestWithContext(ctx, "GET", endpoint, nil)

	if createRequestErr != nil {
		return OidcConfigResponse{}, createRequestErr
	}

	response, fetchErr := client.Do(request)

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
