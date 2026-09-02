package core

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type RevocationOptions struct {
	RevocationEndpoint string
	ClientId           string
	Token              string
}

func Revoke(client *http.Client, options *RevocationOptions) error {
	return RevokeContext(context.Background(), client, options)
}

// RevokeContext is like Revoke but binds the request to ctx.
func RevokeContext(ctx context.Context, client *http.Client, options *RevocationOptions) error {
	values := url.Values{
		"client_id": {options.ClientId},
		"token":     {options.Token},
	}
	request, createRequestErr := http.NewRequestWithContext(ctx, "POST", options.RevocationEndpoint, strings.NewReader(values.Encode()))

	if createRequestErr != nil {
		return createRequestErr
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, fetchErr := client.Do(request)

	if fetchErr != nil {
		return fetchErr
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("revocation error, status code: %d", response.StatusCode)
	}

	return nil
}
