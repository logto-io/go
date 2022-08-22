package core

import (
	"fmt"
	"net/http"
	"net/url"
)

type RevocationOptions struct {
	RevocationEndpoint string
	ClientId           string
	Token              string
}

func Revoke(client *http.Client, options *RevocationOptions) error {
	values := url.Values{
		"client_id": {options.ClientId},
		"token":     {options.Token},
	}
	response, fetchErr := client.PostForm(options.RevocationEndpoint, values)

	if fetchErr != nil {
		return fetchErr
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	return nil
}
