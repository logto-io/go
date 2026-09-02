package core

import (
	"context"
	"net/http"
)

type JwksResponse struct {
	Keys []map[string]string `json:"keys"`
}

func FetchJwks(client *http.Client, jwksUri string) (JwksResponse, error) {
	return FetchJwksContext(context.Background(), client, jwksUri)
}

// FetchJwksContext is like FetchJwks but binds the request to ctx.
func FetchJwksContext(ctx context.Context, client *http.Client, jwksUri string) (JwksResponse, error) {
	request, createRequestErr := http.NewRequestWithContext(ctx, "GET", jwksUri, nil)

	if createRequestErr != nil {
		return JwksResponse{}, createRequestErr
	}

	response, requestErr := client.Do(request)

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
