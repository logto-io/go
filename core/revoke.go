package core

import (
	"fmt"
	"net/http"
	"net/url"
)

type RevocationOptions struct {
	revocationEndpoint string
	clientId           string
	token              string
}

func Revoke(client *http.Client, options *RevocationOptions) error {
	values := url.Values{
		"client_id": {options.clientId},
		"token":     {options.token},
	}
	response, fetchErr := client.PostForm(options.revocationEndpoint, values)

	if fetchErr != nil {
		return fetchErr
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	return nil
}
