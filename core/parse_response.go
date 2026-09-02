package core

import (
	"encoding/json"
	"io"
	"net/http"
)

func parseDataFromResponse(response *http.Response, dest interface{}) error {
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		responseError := &ResponseError{
			StatusCode: response.StatusCode,
			RawBody:    string(body),
		}
		// The error fields stay empty when the body is not a JSON object.
		_ = json.Unmarshal(body, responseError)
		return responseError
	}

	return json.Unmarshal(body, &dest)
}
