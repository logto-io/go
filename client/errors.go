package client

import "errors"

// invalidGrantErrorCode is the OIDC error code returned by the token endpoint
// when the presented grant (e.g. the refresh token) is invalid (RFC 6749 5.2).
const invalidGrantErrorCode = "invalid_grant"

var (
	// ErrNotAuthenticated is returned when no valid user session is available:
	// the user has not signed in, no refresh token is stored, or the stored
	// refresh token has been rejected by the server (expired or revoked).
	// Use errors.Is to detect it and prompt the user to sign in again.
	ErrNotAuthenticated            = errors.New("not authenticated")
	ErrUnacknowledgedResourceFound = errors.New("unacknowledged resource found")
	ErrMissingScopeOrganizations   = errors.New("missing 'urn:logto:scope:organizations' scope")
	ErrMissingOrganizationId       = errors.New("missing organization id")
)
