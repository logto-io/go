package core

import "fmt"

// ResponseError represents a non-200 response returned by the Logto server.
//
// When the response body is a JSON object, the OIDC and Logto error fields
// are parsed into the corresponding struct fields. Otherwise only StatusCode
// and RawBody are populated.
//
// Use errors.As to inspect the details of a failed request:
//
//	var responseError *core.ResponseError
//	if errors.As(err, &responseError) {
//		log.Printf("status: %d, code: %s", responseError.StatusCode, responseError.Code)
//	}
//
// When using LogtoClient, prefer the sentinel errors in the client package for
// well-known scenarios: e.g. a rejected refresh token is reported as
// ErrNotAuthenticated.
type ResponseError struct {
	// StatusCode is the HTTP status code of the response.
	StatusCode int `json:"-"`
	// Code is the Logto error code, e.g. "oidc.invalid_grant".
	Code string `json:"code"`
	// Message is the human-readable message of the Logto error.
	Message string `json:"message"`
	// ErrorCode is the OIDC error code, e.g. "invalid_grant".
	ErrorCode string `json:"error"`
	// ErrorDescription is the human-readable description of the OIDC error.
	ErrorDescription string `json:"error_description"`
	// ErrorUri is a URI to a web page with more information about the error.
	ErrorUri string `json:"error_uri"`
	// RawBody is the raw response body.
	RawBody string `json:"-"`
}

func (responseError *ResponseError) Error() string {
	return fmt.Sprintf("unexpected status code: %d, response body: %s", responseError.StatusCode, responseError.RawBody)
}
