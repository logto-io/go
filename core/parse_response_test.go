package core

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func buildTestResponse(statusCode int, body string) *http.Response {
	return &http.Response{
		StatusCode: statusCode,
		Body:       io.NopCloser(strings.NewReader(body)),
	}
}

func TestParseDataFromResponseShouldParseDataOnOkResponse(t *testing.T) {
	var dest struct {
		Foo string `json:"foo"`
	}

	err := parseDataFromResponse(buildTestResponse(http.StatusOK, `{"foo":"bar"}`), &dest)

	assert.Nil(t, err)
	assert.Equal(t, "bar", dest.Foo)
}

func TestParseDataFromResponseShouldReturnResponseErrorWithParsedErrorBody(t *testing.T) {
	mockResponseBody := `{"code":"oidc.invalid_grant","message":"Grant request is invalid.","error":"invalid_grant","error_description":"grant request is invalid","error_uri":"https://openid.sh/debug/invalid_grant"}`

	var dest interface{}
	err := parseDataFromResponse(buildTestResponse(http.StatusBadRequest, mockResponseBody), &dest)

	var responseError *ResponseError
	assert.True(t, errors.As(err, &responseError))
	assert.Equal(t, http.StatusBadRequest, responseError.StatusCode)
	assert.Equal(t, "oidc.invalid_grant", responseError.Code)
	assert.Equal(t, "Grant request is invalid.", responseError.Message)
	assert.Equal(t, "invalid_grant", responseError.ErrorCode)
	assert.Equal(t, "grant request is invalid", responseError.ErrorDescription)
	assert.Equal(t, "https://openid.sh/debug/invalid_grant", responseError.ErrorUri)
	assert.Equal(t, mockResponseBody, responseError.RawBody)
	assert.Equal(t, "unexpected status code: 400, response body: "+mockResponseBody, err.Error())
}

func TestParseDataFromResponseShouldReturnResponseErrorOnNonJsonBody(t *testing.T) {
	mockResponseBody := "Bad Gateway"

	var dest interface{}
	err := parseDataFromResponse(buildTestResponse(http.StatusBadGateway, mockResponseBody), &dest)

	var responseError *ResponseError
	assert.True(t, errors.As(err, &responseError))
	assert.Equal(t, http.StatusBadGateway, responseError.StatusCode)
	assert.Empty(t, responseError.Code)
	assert.Empty(t, responseError.Message)
	assert.Empty(t, responseError.ErrorCode)
	assert.Empty(t, responseError.ErrorDescription)
	assert.Empty(t, responseError.ErrorUri)
	assert.Equal(t, mockResponseBody, responseError.RawBody)
	assert.Equal(t, "unexpected status code: 502, response body: "+mockResponseBody, err.Error())
}
