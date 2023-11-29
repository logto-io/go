package core

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func parseDataFromResponse(response *http.Response, dest interface{}) error {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d, response body: %s", response.StatusCode, body)
	}

	return json.Unmarshal(body, &dest)
}
